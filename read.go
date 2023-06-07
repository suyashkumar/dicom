package dicom

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"unicode"

	"github.com/suyashkumar/dicom/pkg/debug"
	"github.com/suyashkumar/dicom/pkg/vrraw"

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
)

var (
	// ErrorOWRequiresEvenVL indicates that an element with VR=OW had a not even
	// value length which is not allowed.
	ErrorOWRequiresEvenVL = errors.New("vr of OW requires even value length")
	// ErrorUnsupportedVR indicates that this VR is not supported.
	ErrorUnsupportedVR = errors.New("unsupported VR")
	// ErrorUnsupportedBitsAllocated indicates that the BitsAllocated in the
	// NativeFrame PixelData is unsupported. In this situation, the rest of the
	// dataset returned is still valid.
	ErrorUnsupportedBitsAllocated = errors.New("unsupported BitsAllocated")
	errorUnableToParseFloat       = errors.New("unable to parse float type")
	ErrorExpectedEvenLength       = errors.New("field length is not even, in violation of DICOM spec")
)

// reader is responsible for mid-level dicom parsing capabilities, like
// reading tags, VRs, and elements from the low-level dicomio.Reader dicom data.
// TODO(suyashkumar): consider revisiting naming of this struct "reader" as it
// interplays with the rawReader dicomio.Reader. We could consider combining
// them, or embedding the dicomio.Reader struct into reader.
type reader struct {
	rawReader dicomio.Reader
	opts      parseOptSet
}

func (r *reader) readTag() (*tag.Tag, error) {
	group, gerr := r.rawReader.ReadUInt16()
	element, eerr := r.rawReader.ReadUInt16()

	if gerr == nil && eerr == nil {
		return &tag.Tag{Group: group, Element: element}, nil
	}
	return nil, fmt.Errorf("error reading tag: %v %v", gerr, eerr)
}

// TODO: Parsed VR should be an enum. Will require refactors of tag pkg.
func (r *reader) readVR(isImplicit bool, t tag.Tag) (string, error) {
	if isImplicit {
		if entry, err := tag.Find(t); err == nil {
			return entry.VR, nil
		}
		return tag.UnknownVR, nil
	}

	// Explicit Transfer Syntax, read 2 byte VR:
	return r.rawReader.ReadString(2)

}

func (r *reader) readVL(isImplicit bool, t tag.Tag, vr string) (uint32, error) {
	if isImplicit {
		return r.rawReader.ReadUInt32()
	}

	// Explicit Transfer Syntax
	// More details here: https://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_7.1.2
	switch vr {
	// TODO: Parsed VR should be an enum. Will require refactors of tag pkg.
	case "NA", vrraw.OtherByte, vrraw.OtherDouble, vrraw.OtherFloat,
		vrraw.OtherLong, vrraw.OtherWord, vrraw.Sequence, vrraw.Unknown,
		vrraw.UnlimitedCharacters, vrraw.UniversalResourceIdentifier,
		vrraw.UnlimitedText:
		_ = r.rawReader.Skip(2) // ignore two reserved bytes (0000H)
		vl, err := r.rawReader.ReadUInt32()
		if err != nil {
			return 0, err
		}

		if vl == tag.VLUndefinedLength &&
			(vr == vrraw.UnlimitedCharacters ||
				vr == vrraw.UniversalResourceIdentifier ||
				vr == vrraw.UnlimitedText) {
			return 0, errors.New("UC, UR and UT may not have an Undefined Length, i.e.,a Value Length of FFFFFFFFH")
		}
		return vl, nil
	default:
		vl16, err := r.rawReader.ReadUInt16()
		if err != nil {
			return 0, err
		}
		vl := uint32(vl16)
		// Rectify Undefined Length VL
		if vl == 0xffff {
			vl = tag.VLUndefinedLength
		}
		return vl, nil
	}
}

func (r *reader) readValue(t tag.Tag, vr string, vl uint32, isImplicit bool, d *Dataset, fc chan<- *frame.Frame) (Value, error) {
	vrkind := tag.GetVRKind(t, vr)
	// TODO: if we keep consistent function signature, consider a static map of VR to func?
	switch vrkind {
	case tag.VRBytes:
		return r.readBytes(t, vr, vl)
	case tag.VRString:
		return r.readString(t, vr, vl)
	case tag.VRDate:
		return r.readDate(t, vr, vl)
	case tag.VRUInt16List, tag.VRUInt32List, tag.VRInt16List, tag.VRInt32List, tag.VRTagList:
		return r.readInt(t, vr, vl)
	case tag.VRSequence:
		return r.readSequence(t, vr, vl, d)
	case tag.VRItem:
		return r.readSequenceItem(t, vr, vl, d)
	case tag.VRPixelData:
		return r.readPixelData(vl, d, fc)
	case tag.VRFloat32List, tag.VRFloat64List:
		return r.readFloat(t, vr, vl)
	default:
		return r.readString(t, vr, vl)
	}
}

// readHeader reads the DICOM magic header and group two metadata elements.
// This should only be called once per DICOM at the start of parsing.
func (r *reader) readHeader() ([]*Element, error) {
	// Check to see if magic word is at byte offset 128. If not, this is a
	// non-standard non-compliant DICOM. We try to read this DICOM in a
	// compatibility mode, where we rewind to position 0 and blindly attempt to
	// parse a Dataset (and do not parse metadata in the usual way).
	data, err := r.rawReader.Peek(128 + 4)
	if err != nil {
		return nil, err
	}
	if string(data[128:]) != magicWord {
		return nil, nil
	}

	err = r.rawReader.Skip(128 + 4) // skip preamble + magic word
	if err != nil {
		return nil, err
	}

	// Must read metadata as LittleEndian explicit VR
	// Read the length of the metadata elements: (0002,0000) MetaElementGroupLength
	maybeMetaLen, err := r.readElement(nil, nil)
	if err != nil {
		return nil, err
	}

	metaElems := []*Element{maybeMetaLen} // TODO: maybe set capacity to a reasonable initial size
	metaElementGroupLengthDefined := true
	if maybeMetaLen.Tag != tag.FileMetaInformationGroupLength || maybeMetaLen.Value.ValueType() != Ints {
		// MetaInformationGroupLength is not present or of the wrong value type.
		if !r.opts.allowMissingMetaElementGroupLength {
			return nil, ErrorMetaElementGroupLength
		}
		metaElementGroupLengthDefined = false
	}

	if metaElementGroupLengthDefined {
		metaLen := maybeMetaLen.Value.GetValue().([]int)[0]

		// Read the metadata elements
		err = r.rawReader.PushLimit(int64(metaLen))
		if err != nil {
			return nil, err
		}
		defer r.rawReader.PopLimit()
		for !r.rawReader.IsLimitExhausted() {
			elem, err := r.readElement(nil, nil)
			if err != nil {
				// TODO: see if we can skip over malformed elements somehow
				return nil, err
			}
			// log.Printf("Metadata Element: %s\n", elem)
			metaElems = append(metaElems, elem)
		}
	} else {
		// We cannot use the limit functionality
		debug.Log("Proceeding without metadata group length")
		for {
			// Lets peek into the tag field until we get to end-of-header
			tag_bytes, err := r.rawReader.Peek(4)
			if err != nil {
				return nil, ErrorMetaElementGroupLength
			}
			tg := tag.Tag{}
			buff := bytes.NewBuffer(tag_bytes)
			if err := binary.Read(buff, binary.LittleEndian, &tg.Group); err != nil {
				return nil, err
			}
			if err := binary.Read(buff, binary.LittleEndian, &tg.Element); err != nil {
				return nil, err
			}
			debug.Logf("header-tag: %v", tg)
			// Only read group 2 data
			if tg.Group != 0x0002 {
				break
			}
			elem, err := r.readElement(nil, nil)
			if err != nil {
				// TODO: see if we can skip over malformed elements somehow
				return nil, err
			}
			metaElems = append(metaElems, elem)
		}
	}
	return metaElems, nil
}

func (r *reader) readPixelData(vl uint32, d *Dataset, fc chan<- *frame.Frame) (Value,
	error) {
	if vl == tag.VLUndefinedLength {
		var image PixelDataInfo
		image.IsEncapsulated = true
		// The first Item in PixelData is the basic offset table. Skip this for now.
		// TODO: use basic offset table
		_, _, err := r.readRawItem(true /*shouldSkip*/)
		if err != nil {
			return nil, err
		}

		for !r.rawReader.IsLimitExhausted() {
			data, endOfItems, err := r.readRawItem(r.opts.skipPixelData /*shouldSkip*/)
			if err != nil {
				break
			}

			if endOfItems {
				break
			}

			f := frame.Frame{
				Encapsulated: true,
				EncapsulatedData: frame.EncapsulatedFrame{
					Data: data,
				},
			}

			if fc != nil {
				fc <- &f
			}

			image.Frames = append(image.Frames, f)
		}
		image.IntentionallySkipped = r.opts.skipPixelData
		return &pixelDataValue{PixelDataInfo: image}, nil
	}

	if r.opts.skipPixelData {
		// If we're here, it means the VL isn't undefined length, so we should
		// be able to safely skip the native PixelData.
		if err := r.rawReader.Skip(int64(vl)); err != nil {
			return nil, err
		}
		return &pixelDataValue{PixelDataInfo{IntentionallySkipped: true}}, nil
	}

	if r.opts.skipProcessingPixelDataValue {
		val := &pixelDataValue{PixelDataInfo{IntentionallyUnprocessed: true}}
		val.PixelDataInfo.UnprocessedValueData = make([]byte, vl)
		_, err := io.ReadFull(r.rawReader, val.PixelDataInfo.UnprocessedValueData)
		return val, err
	}

	// Assume we're reading NativeData data since we have a defined value length as per Part 5 Sec A.4 of DICOM spec.
	// We need Elements that have been already parsed (rows, cols, etc) to parse frames out of NativeData Pixel data
	if d == nil {
		return nil, errors.New("the Dataset context cannot be nil in order to read Native PixelData")
	}

	i, _, err := r.readNativeFrames(d, fc, vl)

	if err != nil {
		return nil, err
	}

	// TODO: avoid this copy
	return &pixelDataValue{PixelDataInfo: *i}, nil

}

func getNthBit(data byte, n int) int {
	debug.Logf("mask: %0b", 1<<n)
	if (1 << n & uint8(data)) > 0 {
		return 1
	}
	return 0
}

func fillBufferSingleBitAllocated(pixelData []int, d dicomio.Reader, bo binary.ByteOrder) error {
	debug.Logf("len of pixeldata: %d", len(pixelData))
	if len(pixelData)%8 > 0 {
		return errors.New("when bitsAllocated is 1, we can't read a number of samples that is not a multiple of 8")
	}

	var currentByte byte
	for i := 0; i < len(pixelData)/8; i++ {
		rawData := make([]byte, 1)
		_, err := d.Read(rawData)
		if err != nil {
			return err
		}
		currentByte = rawData[0]
		debug.Logf("currentByte: %0b", currentByte)

		// Read in the 8 bits from the current byte.
		// Always treat the data as LittleEndian encoded.
		// This is what pydicom appears to do, and I can't get Go to properly
		// write out bytes literals in BigEndian, even using binary.Write
		// (in order to test what BigEndian might look like). We should consider
		// revisiting this more closely, and see if the most significant bit tag
		// should be used to determine the read order here.
		idx := 0
		for j := 7; j >= 0; j-- {
			pixelData[(8*i)+idx] = getNthBit(currentByte, j)
			debug.Logf("getbit #%d: %d", j, getNthBit(currentByte, j))
			idx++
		}

	}

	return nil
}

func makeErrorPixelData(reader io.Reader, vl uint32, fc chan<- *frame.Frame, parseErr error) (*PixelDataInfo, error) {
	data := make([]byte, vl)
	_, err := io.ReadFull(reader, data)
	if err != nil {
		return nil, fmt.Errorf("makeErrorPixelData: read pixelData: %w", err)
	}

	f := frame.Frame{
		EncapsulatedData: frame.EncapsulatedFrame{
			Data: data,
		},
	}

	if fc != nil {
		fc <- &f
	}
	image := PixelDataInfo{
		ParseErr: parseErr,
		Frames:   []frame.Frame{f},
	}
	return &image, nil
}

// readNativeFrames reads NativeData frames from a Decoder based on already parsed pixel information
// that should be available in parsedData (elements like NumberOfFrames, rows, columns, etc)
func (r *reader) readNativeFrames(parsedData *Dataset, fc chan<- *frame.Frame, vl uint32) (pixelData *PixelDataInfo,
	bytesToRead int, err error) {
	// Parse information from previously parsed attributes that are needed to parse NativeData Frames:
	rows, err := parsedData.FindElementByTag(tag.Rows)
	if err != nil {
		return nil, 0, err
	}

	cols, err := parsedData.FindElementByTag(tag.Columns)
	if err != nil {
		return nil, 0, err
	}

	nof, err := parsedData.FindElementByTag(tag.NumberOfFrames)
	nFrames := 0
	if err == nil {
		// No error, so parse number of frames
		nFrames, err = strconv.Atoi(MustGetStrings(nof.Value)[0]) // odd that number of frames is encoded as a string...
		if err != nil {
			return nil, 0, err
		}
	} else {
		// error fetching NumberOfFrames, so default to 1. TODO: revisit
		nFrames = 1
	}

	b, err := parsedData.FindElementByTag(tag.BitsAllocated)
	if err != nil {
		return nil, 0, err
	}
	bitsAllocated := MustGetInts(b.Value)[0]

	s, err := parsedData.FindElementByTag(tag.SamplesPerPixel)
	if err != nil {
		return nil, 0, err
	}
	samplesPerPixel := MustGetInts(s.Value)[0]

	pixelsPerFrame := MustGetInts(rows.Value)[0] * MustGetInts(cols.Value)[0]

	debug.Logf("readNativeFrames:\nRows: %d\nCols:%d\nFrames::%d\nBitsAlloc:%d\nSamplesPerPixel:%d", MustGetInts(rows.Value)[0], MustGetInts(cols.Value)[0], nFrames, bitsAllocated, samplesPerPixel)

	bytesAllocated := bitsAllocated / 8
	bytesToRead = bytesAllocated * samplesPerPixel * pixelsPerFrame * nFrames
	if bitsAllocated == 1 {
		bytesToRead = pixelsPerFrame * samplesPerPixel / 8 * nFrames
	}

	skipFinalPaddingByte := false
	if uint32(bytesToRead) != vl {
		switch {
		case uint32(bytesToRead) == vl-1 && vl%2 == 0:
			skipFinalPaddingByte = true
		case uint32(bytesToRead) == vl-1 && vl%2 != 0:
			return nil, 0, fmt.Errorf("vl=%d: %w", vl, ErrorExpectedEvenLength)
		default:
			// calculated bytesToRead and actual VL mismatch
			if !r.opts.allowMismatchPixelDataLength {
				return nil, 0, fmt.Errorf("expected_vl=%d actual_vl=%d %w", bytesToRead, vl, ErrorMismatchPixelDataLength)
			}
			image, err := makeErrorPixelData(r.rawReader, vl, fc, ErrorMismatchPixelDataLength)
			if err != nil {
				return nil, 0, err
			}
			return image, int(vl), nil
		}
	}

	// Parse the pixels:
	image := PixelDataInfo{
		IsEncapsulated: false,
	}
	image.Frames = make([]frame.Frame, nFrames)
	bo := r.rawReader.ByteOrder()
	pixelBuf := make([]byte, bytesAllocated)
	for frameIdx := 0; frameIdx < nFrames; frameIdx++ {
		// Init current frame
		currentFrame := frame.Frame{
			Encapsulated: false,
			NativeData: frame.NativeFrame{
				BitsPerSample: bitsAllocated,
				Rows:          MustGetInts(rows.Value)[0],
				Cols:          MustGetInts(cols.Value)[0],
				Data:          make([][]int, pixelsPerFrame),
			},
		}
		buf := make([]int, pixelsPerFrame*samplesPerPixel)
		if bitsAllocated == 1 {
			if err := fillBufferSingleBitAllocated(buf, r.rawReader, bo); err != nil {
				return nil, bytesToRead, err
			}
			for pixel := 0; pixel < pixelsPerFrame; pixel++ {
				for value := 0; value < samplesPerPixel; value++ {
					currentFrame.NativeData.Data[pixel] = buf[pixel*samplesPerPixel : (pixel+1)*samplesPerPixel]
				}
			}
		} else {
			for pixel := 0; pixel < pixelsPerFrame; pixel++ {
				for value := 0; value < samplesPerPixel; value++ {
					_, err := io.ReadFull(r.rawReader, pixelBuf)
					if err != nil {
						return nil, bytesToRead,
							fmt.Errorf("could not read uint%d from input: %w", bitsAllocated, err)
					}
					if bitsAllocated == 8 {
						buf[(pixel*samplesPerPixel)+value] = int(pixelBuf[0])
					} else if bitsAllocated == 16 {
						buf[(pixel*samplesPerPixel)+value] = int(bo.Uint16(pixelBuf))
					} else if bitsAllocated == 32 {
						buf[(pixel*samplesPerPixel)+value] = int(bo.Uint32(pixelBuf))
					} else {
						return nil, bytesToRead, fmt.Errorf("bitsAllocated=%d : %w", bitsAllocated, ErrorUnsupportedBitsAllocated)
					}
				}
				currentFrame.NativeData.Data[pixel] = buf[pixel*samplesPerPixel : (pixel+1)*samplesPerPixel]
			}
		}
		image.Frames[frameIdx] = currentFrame
		if fc != nil {
			fc <- &currentFrame // write the current frame to the frame channel
		}
	}
	if skipFinalPaddingByte {
		err := r.rawReader.Skip(1)
		if err != nil {
			return nil, bytesToRead, fmt.Errorf("could not read padding byte: %w", err)
		}
		bytesToRead++
	}
	return &image, bytesToRead, nil
}

// readSequence reads a sequence element (VR = SQ) that contains a subset of Items. Each item contains
// a set of Elements.
// See https://dicom.nema.org/medical/dicom/current/output/chtml/part05/sect_7.5.2.html#table_7.5-1
func (r *reader) readSequence(t tag.Tag, vr string, vl uint32, d *Dataset) (Value, error) {
	var sequences sequencesValue

	seqElements := &Dataset{}
	if vl == tag.VLUndefinedLength {
		for {
			subElement, err := r.readElement(seqElements, nil)
			if err != nil {
				// Stop reading due to error
				log.Println("error reading subitem, ", err)
				return nil, err
			}
			if subElement.Tag == tag.SequenceDelimitationItem {
				// Stop reading
				break
			}
			if subElement.Tag != tag.Item || subElement.Value.ValueType() != SequenceItem {
				// This is an error, should be an Item!
				// TODO: use error var
				log.Println("Tag is ", subElement.Tag)
				return nil, fmt.Errorf("non item found in sequence")
			}

			// Append the Item element's dataset of elements to this Sequence's sequencesValue.
			sequences.value = append(sequences.value, subElement.Value.(*SequenceItemValue))
		}
	} else {
		// Sequence of elements for a total of VL bytes
		err := r.rawReader.PushLimit(int64(vl))
		if err != nil {
			return nil, err
		}
		for !r.rawReader.IsLimitExhausted() {
			subElement, err := r.readElement(seqElements, nil)
			if err != nil {
				// TODO: option to ignore errors parsing subelements?
				return nil, err
			}

			// Append the Item element's dataset of elements to this Sequence's sequencesValue.
			sequences.value = append(sequences.value, subElement.Value.(*SequenceItemValue))
		}
		r.rawReader.PopLimit()
	}

	return &sequences, nil
}

// readSequenceItem reads an item component of a sequence dicom element and returns an Element
// with a SequenceItem value.
func (r *reader) readSequenceItem(t tag.Tag, vr string, vl uint32, d *Dataset) (Value, error) {
	var sequenceItem SequenceItemValue

	// seqElements holds items read so farawReader.
	// TODO: deduplicate with sequenceItem above
	seqElements := Dataset{}

	if vl == tag.VLUndefinedLength {
		for {
			subElem, err := r.readElement(&seqElements, nil)
			if err != nil {
				return nil, err
			}
			if subElem.Tag == tag.ItemDelimitationItem {
				break
			}

			sequenceItem.elements = append(sequenceItem.elements, subElem)
			seqElements.Elements = append(seqElements.Elements, subElem)
		}
	} else {
		err := r.rawReader.PushLimit(int64(vl))
		if err != nil {
			return nil, err
		}

		for !r.rawReader.IsLimitExhausted() {
			subElem, err := r.readElement(&seqElements, nil)
			if err != nil {
				return nil, err
			}

			sequenceItem.elements = append(sequenceItem.elements, subElem)
			seqElements.Elements = append(seqElements.Elements, subElem)
		}
		r.rawReader.PopLimit()
	}

	return &sequenceItem, nil
}

func (r *reader) readBytes(t tag.Tag, vr string, vl uint32) (Value, error) {
	// TODO: add special handling of PixelData
	if vr == vrraw.OtherByte || vr == vrraw.Unknown {
		data := make([]byte, vl)
		_, err := io.ReadFull(r.rawReader, data)
		return &bytesValue{value: data}, err
	} else if vr == vrraw.OtherWord {
		// OW -> stream of 16 bit words
		if vl%2 != 0 {
			return nil, ErrorOWRequiresEvenVL
		}

		buf := bytes.NewBuffer(make([]byte, 0, vl))
		numWords := int(vl / 2)
		for i := 0; i < numWords; i++ {
			word, err := r.rawReader.ReadUInt16()
			if err != nil {
				return nil, err
			}
			// TODO: support bytes.BigEndian byte ordering
			err = binary.Write(buf, binary.LittleEndian, word)
			if err != nil {
				return nil, err
			}
		}
		return &bytesValue{value: buf.Bytes()}, nil
	}

	return nil, ErrorUnsupportedVR
}

func (r *reader) readString(t tag.Tag, vr string, vl uint32) (Value, error) {
	str, err := r.rawReader.ReadString(vl)
	onlySpaces := true
	for _, char := range str {
		if !unicode.IsSpace(char) {
			onlySpaces = false
		}
	}
	if !onlySpaces {
		// String may have '\0' suffix if its length is odd.
		str = strings.Trim(str, " \000")
	}

	// Split multiple strings
	strs := strings.Split(str, "\\")
	return &stringsValue{value: strs}, err
}

func (r *reader) readFloat(t tag.Tag, vr string, vl uint32) (Value, error) {
	err := r.rawReader.PushLimit(int64(vl))
	if err != nil {
		return nil, err
	}
	retVal := &floatsValue{value: make([]float64, 0, vl/2)}
	for !r.rawReader.IsLimitExhausted() {
		switch vr {
		case vrraw.FloatingPointSingle:
			val, err := r.rawReader.ReadFloat32()
			if err != nil {
				return nil, err
			}
			// TODO(suyashkumar): revisit this hack to prevent some internal representation issues upconverting from
			// float32 to float64. There is no loss of precision, but the value gets some additional significant digits
			// when using golang casting. This approach prevents those artifacts, but is less efficient.
			pval, err := strconv.ParseFloat(fmt.Sprint(val), 64)
			if err != nil {
				return nil, err
			}
			retVal.value = append(retVal.value, pval)
			break
		case vrraw.FloatingPointDouble:
			val, err := r.rawReader.ReadFloat64()
			if err != nil {
				return nil, err
			}
			retVal.value = append(retVal.value, val)
			break
		default:
			return nil, errorUnableToParseFloat
		}
	}
	r.rawReader.PopLimit()
	return retVal, nil
}

func (r *reader) readDate(t tag.Tag, vr string, vl uint32) (Value, error) {
	rawDate, err := r.rawReader.ReadString(vl)
	if err != nil {
		return nil, err
	}
	date := strings.Trim(rawDate, " \000")

	return &stringsValue{value: []string{date}}, nil

}

func (r *reader) readInt(t tag.Tag, vr string, vl uint32) (Value, error) {
	// TODO: add other integer types here
	err := r.rawReader.PushLimit(int64(vl))
	if err != nil {
		return nil, err
	}
	retVal := &intsValue{value: make([]int, 0, vl/2)}
	for !r.rawReader.IsLimitExhausted() {
		switch vr {
		case vrraw.UnsignedShort, vrraw.AttributeTag:
			val, err := r.rawReader.ReadUInt16()
			if err != nil {
				return nil, err
			}
			retVal.value = append(retVal.value, int(val))
			break
		case vrraw.UnsignedLong:
			val, err := r.rawReader.ReadUInt32()
			if err != nil {
				return nil, err
			}
			retVal.value = append(retVal.value, int(val))
			break
		case vrraw.SignedLong:
			val, err := r.rawReader.ReadInt32()
			if err != nil {
				return nil, err
			}
			retVal.value = append(retVal.value, int(val))
			break
		case vrraw.SignedShort:
			val, err := r.rawReader.ReadInt16()
			if err != nil {
				return nil, err
			}
			retVal.value = append(retVal.value, int(val))
			break
		default:
			return nil, errors.New("unable to parse integer type")
		}
	}
	r.rawReader.PopLimit()
	return retVal, err
}

// readElement reads the next element. If the next element is a sequence element,
// it may result in a collection of Elements. It takes a pointer to the Dataset of
// elements read so far, since previously read elements may be needed to parse
// certain Elements (like native PixelData). If the Dataset is nil, it is
// treated as an empty Dataset.
func (r *reader) readElement(d *Dataset, fc chan<- *frame.Frame) (*Element, error) {
	t, err := r.readTag()
	if err != nil {
		return nil, err
	}
	debug.Logf("readElement: tag: %s", t.String())

	readImplicit := r.rawReader.IsImplicit()
	if *t == tag.Item {
		// Always read implicit for item elements
		readImplicit = true
	}

	vr, err := r.readVR(readImplicit, *t)
	if err != nil {
		return nil, err
	}
	debug.Logf("readElement: vr: %s", vr)

	vl, err := r.readVL(readImplicit, *t, vr)
	if err != nil {
		return nil, err
	}
	debug.Logf("readElement: vl: %d", vl)

	val, err := r.readValue(*t, vr, vl, readImplicit, d, fc)
	if err != nil {
		log.Println("error reading value ", err)
		return nil, err
	}

	return &Element{Tag: *t, ValueRepresentation: tag.GetVRKind(*t, vr), RawValueRepresentation: vr, ValueLength: vl, Value: val}, nil

}

// Read an Item object as raw bytes, useful when parsing encapsulated PixelData.
// This returns the read raw item, an indication if this is the end of the set
// of items, and a possible errorawReader.
func (r *reader) readRawItem(shouldSkip bool) ([]byte, bool, error) {
	t, err := r.readTag()
	if err != nil {
		return nil, true, err
	}
	// Item is always encoded implicit. PS3.6 7.5
	vr, err := r.readVR(true, *t)
	if err != nil {
		return nil, true, err
	}
	vl, err := r.readVL(true, *t, vr)
	if err != nil {
		return nil, true, err
	}

	if *t == tag.SequenceDelimitationItem {
		if vl != 0 {
			log.Printf("SequenceDelimitationItem's VL != 0: %d", vl)
		}
		return nil, true, nil
	}
	if *t != tag.Item {
		log.Printf("Expect Item in pixeldata but found tag %s", tag.DebugString(*t))
		return nil, false, nil
	}
	if vl == tag.VLUndefinedLength {
		log.Println("Expect defined-length item in pixeldata")
		return nil, false, nil
	}
	if vr != "NA" {
		return nil, true, fmt.Errorf("readRawItem: expected VR=NA, got VR=%s", vr)
	}

	if shouldSkip {
		if err := r.rawReader.Skip(int64(vl)); err != nil {
			return nil, false, err
		}
	} else {
		data := make([]byte, vl)
		_, err = io.ReadFull(r.rawReader, data)
		if err != nil {
			log.Println(err)
			return nil, false, err
		}
		return data, false, nil
	}
	return nil, false, nil
}

// moreToRead returns true if there is more to read from the underlying dicom.
func (r *reader) moreToRead() bool {
	return !r.rawReader.IsLimitExhausted()
}
