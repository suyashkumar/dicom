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

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
)

var (
	ErrorOWRequiresEvenVL = errors.New("vr of OW requires even value length")
	ErrorUnsupportedVR    = errors.New("unsupported VR")
)

func readTag(r dicomio.Reader) (*tag.Tag, error) {
	group, gerr := r.ReadUInt16()
	element, eerr := r.ReadUInt16()

	if gerr == nil && eerr == nil {
		return &tag.Tag{Group: group, Element: element}, nil
	} else {
		return nil, fmt.Errorf("error reading tag: %v %v", gerr, eerr)
	}
}

// TODO: Parsed VR should be an enum. Will require refactors of tag pkg.
func readVR(r dicomio.Reader, isImplicit bool, t tag.Tag) (string, error) {
	if isImplicit {
		if entry, err := tag.Find(t); err == nil {
			return entry.VR, nil
		}
		return tag.UNKNOWN, nil
	}

	// Explicit Transfer Syntax, read 2 byte VR:
	return r.ReadString(2)

}

func readVL(r dicomio.Reader, isImplicit bool, t tag.Tag, vr string) (uint32, error) {
	if isImplicit {
		return r.ReadUInt32()
	}

	// Explicit Transfer Syntax
	// More details here: http://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_7.1.2
	switch vr {
	// TODO: Parsed VR should be an enum. Will require refactors of tag pkg.
	case "NA", "OB", "OD", "OF", "OL", "OW", "SQ", "UN", "UC", "UR", "UT":
		_ = r.Skip(2) // ignore two reserved bytes (0000H)
		vl, err := r.ReadUInt32()
		if err != nil {
			return 0, err
		}

		if vl == tag.VLUndefinedLength && (vr == "UC" || vr == "UR" || vr == "UT") {
			return 0, errors.New("UC, UR and UT may not have an Undefined Length, i.e.,a Value Length of FFFFFFFFH")
		}
		return vl, nil
	default:
		vl16, err := r.ReadUInt16()
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

func readValue(r dicomio.Reader, t tag.Tag, vr string, vl uint32, isImplicit bool, d *Dataset, fc chan<- *frame.Frame) (Value, error) {
	// TODO: implement
	vrkind := tag.GetVRKind(t, vr)
	// TODO: if we keep consistent function signature, consider a static map of VR to func?
	switch vrkind {
	case tag.VRBytes:
		return readBytes(r, t, vr, vl)
	case tag.VRString:
		return readString(r, t, vr, vl)
	case tag.VRDate:
		return readDate(r, t, vr, vl)
	case tag.VRUInt16List, tag.VRUInt32List, tag.VRInt16List, tag.VRInt32List:
		return readInt(r, t, vr, vl)
	case tag.VRSequence:
		return readSequence(r, t, vr, vl)
	case tag.VRItem:
		return readSequenceItem(r, t, vr, vl)
	case tag.VRPixelData:
		return readPixelData(r, t, vr, vl, d, fc)
	default:
		return readString(r, t, vr, vl)
	}

	return nil, fmt.Errorf("unsure how to parse this VR")
}

func readPixelData(r dicomio.Reader, t tag.Tag, vr string, vl uint32, d *Dataset, fc chan<- *frame.Frame) (Value,
	error) {
	if vl == tag.VLUndefinedLength {
		var image PixelDataInfo
		image.IsEncapsulated = true
		// The first Item in PixelData is the basic offset table. Skip this for now.
		// TODO: use basic offset table
		_, _, err := readRawItem(r)
		if err != nil {
			return nil, err
		}

		for !r.IsLimitExhausted() {
			data, endOfItems, err := readRawItem(r)
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
		return &pixelDataValue{PixelDataInfo: image}, nil
	}

	// Assume we're reading NativeData data since we have a defined value length as per Part 5 Sec A.4 of DICOM spec.
	// We need Elements that have been already parsed (rows, cols, etc) to parse frames out of NativeData Pixel data
	if d == nil {
		return nil, errors.New("the Dataset context cannot be nil in order to read Native PixelData")
	}

	i, _, err := readNativeFrames(r, d, fc)
	log.Println("Dataset Size: ", len(d.Elements))

	if err != nil {
		return nil, err
	}

	// TODO: avoid this copy
	return &pixelDataValue{PixelDataInfo: *i}, nil

}

// readNativeFrames reads NativeData frames from a Decoder based on already parsed pixel information
// that should be available in parsedData (elements like NumberOfFrames, rows, columns, etc)
func readNativeFrames(d dicomio.Reader, parsedData *Dataset, fc chan<- *frame.Frame) (pixelData *PixelDataInfo,
	bytesRead int, err error) {
	image := PixelDataInfo{
		IsEncapsulated: false,
	}

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

	// Parse the pixels:
	image.Frames = make([]frame.Frame, nFrames)
	for frameIdx := 0; frameIdx < nFrames; frameIdx++ {
		// Init current frame
		currentFrame := frame.Frame{
			Encapsulated: false,
			NativeData: frame.NativeFrame{
				BitsPerSample: bitsAllocated,
				Rows:          MustGetInts(rows.Value)[0],
				Cols:          MustGetInts(cols.Value)[0],
				Data:          make([][]int, int(pixelsPerFrame)),
			},
		}
		for pixel := 0; pixel < int(pixelsPerFrame); pixel++ {
			currentPixel := make([]int, samplesPerPixel)
			for value := 0; value < samplesPerPixel; value++ {
				if bitsAllocated == 8 {
					val, err := d.ReadUInt8()
					if err != nil {
						return nil, bytesRead, errors.New("")
					}
					currentPixel[value] = int(val)
				} else if bitsAllocated == 16 {
					val, err := d.ReadUInt16()
					if err != nil {
						return nil, bytesRead, errors.New("")
					}
					currentPixel[value] = int(val)
				}
			}
			currentFrame.NativeData.Data[pixel] = currentPixel
		}
		image.Frames[frameIdx] = currentFrame
		if fc != nil {
			fc <- &currentFrame // write the current frame to the frame channel
		}
	}

	bytesRead = (bitsAllocated / 8) * samplesPerPixel * pixelsPerFrame * nFrames

	return &image, bytesRead, nil
}

// readSequence reads a sequence element (VR = SQ) that contains a subset of Items. Each item contains
// a set of Elements.
// See http://dicom.nema.org/medical/dicom/current/output/chtml/part05/sect_7.5.2.html#table_7.5-1
func readSequence(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	var sequences sequencesValue

	if vl == tag.VLUndefinedLength {
		for {
			subElement, err := readElement(r, nil, nil)
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
				return nil, fmt.Errorf("non item found in sequence")
			}

			// Append the Item element's dataset of elements to this Sequence's sequencesValue.
			sequences.value = append(sequences.value, subElement.Value.(*SequenceItemValue))
		}
	} else {
		// Sequence of elements for a total of VL bytes
		err := r.PushLimit(int64(vl))
		if err != nil {
			return nil, err
		}
		for !r.IsLimitExhausted() {
			subElement, err := readElement(r, nil, nil)
			if err != nil {
				// TODO: option to ignore errors parsing subelements?
				return nil, err
			}

			// Append the Item element's dataset of elements to this Sequence's sequencesValue.
			sequences.value = append(sequences.value, subElement.Value.(*SequenceItemValue))
		}
		r.PopLimit()
	}

	return &sequences, nil
}

// readSequenceItem reads an item component of a sequence dicom element and returns an Element
// with a SequenceItem value.
func readSequenceItem(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	var sequenceItem SequenceItemValue

	// seqElements holds items read so far.
	// TODO: deduplicate with sequenceItem above
	var seqElements Dataset

	if vl == tag.VLUndefinedLength {
		for {
			subElem, err := readElement(r, &seqElements, nil)
			if err != nil {
				return nil, err
			}
			if subElem.Tag == tag.ItemDelimitationItem {
				break
			}
			// log.Println("readSequenceItem: tag: ", subElem.Tag)

			sequenceItem.elements = append(sequenceItem.elements, subElem)
			seqElements.Elements = append(seqElements.Elements, subElem)
		}
	} else {
		err := r.PushLimit(int64(vl))
		if err != nil {
			return nil, err
		}

		for !r.IsLimitExhausted() {
			subElem, err := readElement(r, &seqElements, nil)
			if err != nil {
				return nil, err
			}
			// log.Println("readSequenceItem: tag: ", subElem.Tag)

			sequenceItem.elements = append(sequenceItem.elements, subElem)
			seqElements.Elements = append(seqElements.Elements, subElem)
		}
		r.PopLimit()
	}

	return &sequenceItem, nil
}

func readBytes(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	// TODO: add special handling of PixelData
	if vr == "OB" {
		data := make([]byte, vl)
		_, err := io.ReadFull(r, data)
		return &bytesValue{value: data}, err
	} else if vr == "OW" {
		// OW -> stream of 16 bit words
		if vl%2 != 0 {
			return nil, ErrorOWRequiresEvenVL
		}
		data := make([]byte, vl)
		buf := bytes.NewBuffer(data)
		numWords := int(vl / 2)
		for i := 0; i < numWords; i++ {
			word, err := r.ReadUInt16()
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

func readString(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	str, err := r.ReadString(vl)
	onlySpaces := true
	for _, char := range str {
		if !unicode.IsSpace(char) {
			onlySpaces = false
		}
	}
	if !onlySpaces{
		// String may have '\0' suffix if its length is odd.
		str = strings.Trim(str, " \000")
	}

	// Split multiple strings
	strs := strings.Split(str, "\\")

	return &stringsValue{value: strs}, err
}

func readDate(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	rawDate, err := r.ReadString(vl)
	if err != nil {
		return nil, err
	}
	date := strings.Trim(rawDate, " \000")

	return &stringsValue{value: []string{date}}, nil

}

func readInt(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	// TODO: add other integer types here
	err := r.PushLimit(int64(vl))
	retVal := &intsValue{value: make([]int, 0, vl/2)}
	for !r.IsLimitExhausted() {
		switch vr {
		case "US":
			val, err := r.ReadUInt16()
			if err != nil {
				return nil, err
			}
			retVal.value = append(retVal.value, int(val))
			break
		case "UL":
			val, err := r.ReadUInt32()
			if err != nil {
				return nil, err
			}
			retVal.value = append(retVal.value, int(val))
			break
		case "SL":
			val, err := r.ReadInt32()
			if err != nil {
				return nil, err
			}
			retVal.value = append(retVal.value, int(val))
			break
		case "SS":
			val, err := r.ReadInt16()
			if err != nil {
				return nil, err
			}
			retVal.value = append(retVal.value, int(val))
			break
		default:
			return nil, errors.New("unable to parse integer type")
		}
	}
	r.PopLimit()
	return retVal, err
}

// readElement reads the next element. If the next element is a sequence element,
// it may result in a collection of Elements. It takes a pointer to the Dataset of
// elements read so far, since previously read elements may be needed to parse
// certain Elements (like native PixelData). If the Dataset is nil, it is
// treated as an empty Dataset.
func readElement(r dicomio.Reader, d *Dataset, fc chan<- *frame.Frame) (*Element, error) {
	t, err := readTag(r)
	if err != nil {
		return nil, err
	}
	// log.Println("readElement: readTag: ", t)

	vr, err := readVR(r, r.IsImplicit(), *t)
	if err != nil {
		return nil, err
	}

	vl, err := readVL(r, r.IsImplicit(), *t, vr)
	if err != nil {
		return nil, err
	}

	// log.Println("readElement: vr, vl", vr, vl)

	val, err := readValue(r, *t, vr, vl, r.IsImplicit(), d, fc)
	if err != nil {
		log.Println("error reading value ", err)
		return nil, err
	}

	return &Element{Tag: *t, ValueRepresentation: tag.GetVRKind(*t, vr), RawValueRepresentation: vr, ValueLength: vl, Value: val}, nil

}

// Read an Item object as raw bytes, useful when parsing encapsulated PixelData
func readRawItem(r dicomio.Reader) ([]byte, bool, error) {
	t, err := readTag(r)
	// Item is always encoded implicit. PS3.6 7.5
	vr, err := readVR(r, true, *t)
	vl, err := readVL(r, true, *t, vr)

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

	data := make([]byte, vl)
	_, err = io.ReadFull(r, data)
	if err != nil {
		log.Println(err)
		return nil, false, err
	}
	return data, false, nil
}
