package dicom

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"slices"

	"github.com/suyashkumar/dicom/pkg/vrraw"

	"github.com/suyashkumar/dicom/pkg/uid"

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/tag"
)

var (
	// ErrorUnimplemented is for not yet finished things.
	ErrorUnimplemented = errors.New("this functionality is not yet implemented")
	// ErrorMismatchValueTypeAndVR is for when there's a discrepancy between the ValueType and what the VR specifies.
	ErrorMismatchValueTypeAndVR = errors.New("ValueType does not match the VR required")
	// ErrorUnexpectedValueType indicates an unexpected value type was seen.
	ErrorUnexpectedValueType = errors.New("Unexpected ValueType")
	// ErrorUnsupportedBitsPerSample indicates that the BitsPerSample in this
	// Dataset is not supported when unpacking native PixelData.
	ErrorUnsupportedBitsPerSample          = errors.New("unsupported BitsPerSample value")
	errorDeflatedTransferSyntaxUnsupported = errors.New("deflated explicit vr little endian transfer syntax not yet support on write (https://github.com/suyashkumar/dicom/issues/323)")
)

// Writer is a struct that allows element-by element writing to a DICOM writer.
type Writer struct {
	writer *dicomio.Writer
	optSet *writeOptSet
}

// NewWriter returns a new Writer, that points to the provided io.Writer.
func NewWriter(out io.Writer, opts ...WriteOption) (*Writer, error) {
	optSet := toWriteOptSet(opts...)
	if err := optSet.validate(); err != nil {
		return nil, err
	}
	w := dicomio.NewWriter(out, nil, false)

	return &Writer{
		writer: w,
		optSet: optSet,
	}, nil
}

// SetTransferSyntax sets the transfer syntax for the underlying dicomio.Writer.
func (w *Writer) SetTransferSyntax(bo binary.ByteOrder, implicit bool) {
	w.writer.SetTransferSyntax(bo, implicit)
}

// writeDataset writes the provided DICOM dataset to the Writer, including headers if available.
func (w *Writer) writeDataset(ds Dataset) error {
	var metaElems []*Element
	for _, elem := range ds.Elements {
		if elem.Tag.Group == tag.MetadataGroup {
			metaElems = append(metaElems, elem)
		}
	}

	err := writeFileHeader(w.writer, &ds, metaElems, *w.optSet)
	if err != nil {
		return err
	}

	bo, implicit, err := ds.transferSyntax()
	if errors.Is(err, ErrorElementNotFound) && w.optSet.defaultMissingTransferSyntax {
		bo = binary.LittleEndian
		implicit = true
	} else if errors.Is(err, ErrorElementNotFound) && w.optSet.overrideMissingTransferSyntaxUID != "" {
		bo, implicit, err = uid.ParseTransferSyntaxUID(w.optSet.overrideMissingTransferSyntaxUID)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	w.writer.SetTransferSyntax(bo, implicit)

	for _, elem := range ds.Elements {
		if elem.Tag.Group != tag.MetadataGroup {
			err = writeElement(w.writer, elem, *w.optSet)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// WriteElement writes a single DICOM element to a Writer.
func (w *Writer) WriteElement(e *Element) error {
	return writeElement(w.writer, e, *w.optSet)
}

// Write will write the input DICOM dataset to the provided io.Writer as a complete DICOM (including any header
// information if available).
func Write(out io.Writer, ds Dataset, opts ...WriteOption) error {
	w, err := NewWriter(out, opts...)
	if err != nil {
		return err
	}
	return w.writeDataset(ds)
}

// WriteOption represents an option that can be passed to WriteDataset. Later options will override previous options if
// applicable.
type WriteOption func(*writeOptSet)

// SkipVRVerification returns a WriteOption that skips VR verification.
func SkipVRVerification() WriteOption {
	return func(set *writeOptSet) {
		set.skipVRVerification = true
	}
}

// SkipValueTypeVerification returns WriteOption function that skips checking ValueType
// for concurrency with VR and casting
func SkipValueTypeVerification() WriteOption {
	return func(set *writeOptSet) {
		set.skipValueTypeVerification = true
	}
}

// DefaultMissingTransferSyntax returns a WriteOption indicating that a missing
// transferSyntax should not raise an error, and instead the default
// LittleEndian Implicit transfer syntax should be used and written out as a
// Metadata element in the Dataset.
// TODO(suyashkumar): consider deprecating in favor of
// OverrideMissingTransferSyntax.
func DefaultMissingTransferSyntax() WriteOption {
	return func(set *writeOptSet) {
		set.defaultMissingTransferSyntax = true
	}
}

// OverrideMissingTransferSyntax indicates that if the Dataset to be written
// does _not_ have a transfer syntax UID in its metadata, the Dataset should
// be written out with the provided transfer syntax UID if possible. A
// transfer syntax uid element with the specified transfer syntax will be
// written to the metadata as well.
//
// If the Writer is unable to recognize or write the dataset using the provided
// transferSyntaxUID, an error will be returned at initialization time.
func OverrideMissingTransferSyntax(transferSyntaxUID string) WriteOption {
	return func(set *writeOptSet) {
		set.overrideMissingTransferSyntaxUID = transferSyntaxUID
	}
}

// skipWritingTransferSyntaxForTests is a test WriteOption that cause Write to skip
// writing the transfer syntax uid element in the DICOM metadata. When used in
// combination with OverrideMissingTransferSyntax, this can be used to set the
// TransferSyntax for the written dataset without writing the actual transfer
// syntax element to the metadata.
func skipWritingTransferSyntaxForTests() WriteOption {
	return func(set *writeOptSet) {
		set.skipWritingTransferSyntaxForTests = true
	}
}

// writeOptSet represents the flattened option set after all WriteOptions have been applied.
type writeOptSet struct {
	skipVRVerification                bool
	skipValueTypeVerification         bool
	defaultMissingTransferSyntax      bool
	overrideMissingTransferSyntaxUID  string
	skipWritingTransferSyntaxForTests bool
}

func (w *writeOptSet) validate() error {
	if w.overrideMissingTransferSyntaxUID != "" {
		if _, _, err := uid.ParseTransferSyntaxUID(w.overrideMissingTransferSyntaxUID); err != nil {
			return fmt.Errorf("unable to parse OverrideMissingTransferSyntax transfer syntax uid %v due to: %s", w.overrideMissingTransferSyntaxUID, err)
		}
	}
	return nil
}

func toWriteOptSet(opts ...WriteOption) *writeOptSet {
	optSet := &writeOptSet{}
	for _, opt := range opts {
		opt(optSet)
	}
	return optSet
}

func writeFileHeader(w *dicomio.Writer, ds *Dataset, metaElems []*Element, opts writeOptSet) error {
	// File headers are always written in littleEndian explicit
	w.SetTransferSyntax(binary.LittleEndian, false)

	metaBytes := &bytes.Buffer{}
	subWriter := dicomio.NewWriter(metaBytes, binary.LittleEndian, false)
	tagsUsed := make(map[tag.Tag]bool)
	tagsUsed[tag.FileMetaInformationGroupLength] = true

	err := writeMetaElem(subWriter, tag.FileMetaInformationVersion, ds, &tagsUsed, opts)
	if err != nil && !errors.Is(err, ErrorElementNotFound) {
		return err
	}
	err = writeMetaElem(subWriter, tag.MediaStorageSOPClassUID, ds, &tagsUsed, opts)
	if err != nil && !errors.Is(err, ErrorElementNotFound) {
		return err
	}
	err = writeMetaElem(subWriter, tag.MediaStorageSOPInstanceUID, ds, &tagsUsed, opts)
	if err != nil && !errors.Is(err, ErrorElementNotFound) {
		return err
	}

	if !opts.skipWritingTransferSyntaxForTests {
		err = writeMetaElem(subWriter, tag.TransferSyntaxUID, ds, &tagsUsed, opts)

		if errors.Is(err, ErrorElementNotFound) && opts.defaultMissingTransferSyntax {
			// Write the default transfer syntax
			if err = writeElement(subWriter, mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}), opts); err != nil {
				return err
			}
		} else if errors.Is(err, ErrorElementNotFound) && opts.overrideMissingTransferSyntaxUID != "" {
			// Write the override transfer syntax
			if err = writeElement(subWriter, mustNewElement(tag.TransferSyntaxUID, []string{opts.overrideMissingTransferSyntaxUID}), opts); err != nil {
				return err
			}
		} else if err != nil {
			// Return the error if none of the above conditions/overrides apply.
			return err
		}
	}

	for _, elem := range metaElems {
		if elem.Tag.Group == tag.MetadataGroup {
			if _, ok := tagsUsed[elem.Tag]; !ok {
				err = writeElement(subWriter, elem, opts)
				if err != nil {
					return err
				}
			}
		}
	}

	if err := w.WriteZeros(128); err != nil {
		return err
	}
	if err := w.WriteString(magicWord); err != nil {
		return err
	}
	lengthElem, err := NewElement(tag.FileMetaInformationGroupLength, []int{len(metaBytes.Bytes())})
	if err != nil {
		return err
	}

	err = writeElement(w, lengthElem, opts)
	if err != nil {
		return err
	}
	err = w.WriteBytes(metaBytes.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func writeElement(w *dicomio.Writer, elem *Element, opts writeOptSet) error {
	vr := elem.RawValueRepresentation
	var err error
	vr, err = verifyVROrDefault(elem.Tag, elem.RawValueRepresentation, opts)
	if err != nil {
		return err
	}

	if !opts.skipValueTypeVerification && elem.Value != nil {
		err := verifyValueType(elem.Tag, elem.Value, vr)
		if err != nil {
			return err
		}
	}

	length := elem.ValueLength
	var valueData = &bytes.Buffer{}
	if elem.Value != nil {
		bo, implicit := w.GetTransferSyntax()
		subWriter := dicomio.NewWriter(valueData, bo, implicit)
		err := writeValue(subWriter, elem.Tag, elem.Value, elem.Value.ValueType(), vr, elem.ValueLength, opts)
		if err != nil {
			return err
		}

		length = uint32(len(valueData.Bytes()))
		if elem.ValueLength == tag.VLUndefinedLength {
			length = tag.VLUndefinedLength
		}
	}

	err = encodeElementHeader(w, elem.Tag, vr, length)
	if err != nil {
		return err
	}

	if elem.Value != nil {
		// Write the bytes to the original writer
		err = w.WriteBytes(valueData.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func writeMetaElem(w *dicomio.Writer, t tag.Tag, ds *Dataset, tagsUsed *map[tag.Tag]bool, optSet writeOptSet) error {
	elem, err := ds.FindElementByTag(t)
	if err != nil {
		return err
	}
	if elem.Tag == tag.TransferSyntaxUID && MustGetStrings(elem.Value)[0] == uid.DeflatedExplicitVRLittleEndian {
		return errorDeflatedTransferSyntaxUnsupported
	}
	err = writeElement(w, elem, optSet)
	if err != nil {
		return err
	}
	(*tagsUsed)[t] = true
	return nil
}

func verifyVROrDefault(t tag.Tag, vr string, opts writeOptSet) (string, error) {
	// If our VR is not blank and we are skipping VF verification, nothing needs to be
	// done, so we can immediately return.
	if vr != "" && opts.skipVRVerification {
		return vr, nil
	}

	// Otherwise, get our tag info.
	tagInfo, err := tag.Find(t)
	if err != nil {
		// If we cannot find information about the tag and our VR is blank, we will use
		// "UN" (Unknown). Otherwise we will fallback to the caller's VR and trust that
		// they know more about this tag than we do. This could be a private tag, or a
		// tag from a newer version of the DICOM spec.
		if vr == "" {
			vr = vrraw.Unknown
		}
		return vr, nil
	}

	if vr == "" {
		// Otherwise if we did find it, and our VR is blank, we'll return the known vr
		// we just pulled.
		return tagInfo.VRs[0], nil
	}

	// Verify the VR on the way out if the caller wants it.
	if !opts.skipVRVerification && !slices.Contains(tagInfo.VRs, vr) {
		return "", fmt.Errorf("ERROR dicomio.verifyElement: VR mismatch for tag %v. Element.VR=%v, but DICOM standard defines VR to be %v",
			tag.DebugString(t), vr, tagInfo.VRs)
	}

	return vr, nil
}

func verifyValueType(t tag.Tag, value Value, vr string) error {
	valueType := value.ValueType()
	var ok bool
	switch vr {
	case vrraw.UnsignedShort, vrraw.UnsignedLong, vrraw.SignedLong, vrraw.SignedShort, vrraw.AttributeTag:
		ok = valueType == Ints
	case vrraw.Sequence:
		ok = valueType == Sequences
	case "NA":
		ok = valueType == SequenceItem
	case vrraw.OtherWord, vrraw.OtherByte:
		if t == tag.PixelData {
			ok = valueType == PixelData
		} else {
			ok = valueType == Bytes
		}
	case vrraw.Unknown:
		ok = valueType == Bytes || valueType == Sequences
	case vrraw.FloatingPointSingle, vrraw.FloatingPointDouble:
		ok = valueType == Floats
	default:
		ok = valueType == Strings
	}

	if !ok {
		return fmt.Errorf("ValueType does not match the specified type in the VR")
	}
	return nil
}

func writeTag(w *dicomio.Writer, t tag.Tag, vl uint32) error {
	if vl%2 != 0 && vl != tag.VLUndefinedLength {
		return fmt.Errorf("ERROR dicomio.writeTag: Value Length must be even, but for Tag=%v, ValueLength=%v",
			tag.DebugString(t), vl)
	}
	if err := w.WriteUInt16(t.Group); err != nil {
		return err
	}
	return w.WriteUInt16(t.Element)
}

func writeVRVL(w *dicomio.Writer, t tag.Tag, vr string, vl uint32) error {
	// Rectify Undefined Length VL
	if vl == 0xffff {
		vl = tag.VLUndefinedLength
	}

	if vr == vrraw.Sequence {
		// We are going to write these out with undefined length always.
		vl = tag.VLUndefinedLength
	}

	// We want to make sure there is any VR unless this is a Sequence delimiter.
	if len(vr) != 2 && vl != tag.VLUndefinedLength && t != tag.SequenceDelimitationItem && t != tag.ItemDelimitationItem {
		return fmt.Errorf("ERROR dicomio.writeVRVL: Value Representation must be of length 2, e.g. 'UN'. For tag=%v, it was RawValueRepresentation=%v",
			tag.DebugString(t), vr)
	}

	// Write VR then VL
	_, implicit := w.GetTransferSyntax()
	if t.Group == tag.GroupSeqItem {
		implicit = true
	}
	if !implicit { // Explicit
		if err := w.WriteString(vr); err != nil {
			return err
		}
		switch vr {
		case "NA", vrraw.OtherByte, vrraw.OtherDouble, vrraw.OtherFloat,
			vrraw.OtherLong, vrraw.OtherWord, vrraw.Sequence, vrraw.Unknown,
			vrraw.UnlimitedCharacters, vrraw.UniversalResourceIdentifier,
			vrraw.UnlimitedText:
			if err := w.WriteZeros(2); err != nil {
				return err
			}
			if err := w.WriteUInt32(vl); err != nil {
				return err
			}
		default:
			if err := w.WriteUInt16(uint16(vl)); err != nil {
				return err
			}
		}
	} else {
		if err := w.WriteUInt32(vl); err != nil {
			return err
		}
	}
	return nil
}

func writeRawItem(w *dicomio.Writer, data []byte) error {
	length := uint32(len(data))
	if err := writeTag(w, tag.Item, length); err != nil {
		return err
	}
	if err := writeVRVL(w, tag.Item, "NA", length); err != nil {
		return err
	}
	if err := w.WriteBytes(data); err != nil {
		return err
	}
	return nil
}

func writeBasicOffsetTable(w *dicomio.Writer, offsets []uint32) error {
	byteOrder, implicit := w.GetTransferSyntax()
	data := &bytes.Buffer{}
	subWriter := dicomio.NewWriter(data, byteOrder, implicit)
	for _, offset := range offsets {
		if err := subWriter.WriteUInt32(offset); err != nil {
			return err
		}
	}
	return writeRawItem(w, data.Bytes())
}

func encodeElementHeader(w *dicomio.Writer, t tag.Tag, vr string, vl uint32) error {
	err := writeTag(w, t, vl)
	if err != nil {
		return err
	}
	err = writeVRVL(w, t, vr, vl)
	if err != nil {
		return err
	}
	return nil
}

func writeValue(w *dicomio.Writer, t tag.Tag, value Value, valueType ValueType, vr string, vl uint32, opts writeOptSet) error {
	if vl == tag.VLUndefinedLength && valueType <= 2 { // strings, bytes or ints
		return fmt.Errorf("encoding undefined-length element not yet supported: %v", t)
	}

	v := value.GetValue()
	switch valueType {
	case Strings:
		return writeStrings(w, v.([]string), vr)
	case Bytes:
		return writeBytes(w, v.([]byte), vr)
	case Ints:
		return writeInts(w, v.([]int), vr)
	case PixelData:
		return writePixelData(w, t, value, vr, vl)
	case SequenceItem:
		return writeSequenceItem(w, t, v.([]*Element), vr, vl, opts)
	case Sequences:
		return writeSequence(w, t, v.([]*SequenceItemValue), vr, vl, opts)
	case Floats:
		return writeFloats(w, value, vr)
	default:
		return fmt.Errorf("ValueType not supported")
	}
}

func writeStrings(w *dicomio.Writer, values []string, vr string) error {
	s := ""
	for i, substr := range values {
		if i > 0 {
			s += "\\"
		}
		s += substr
	}
	if err := w.WriteString(s); err != nil {
		return err
	}
	if len(s)%2 == 1 {
		switch vr {
		case vrraw.DateTime, vrraw.LongString, vrraw.LongText, vrraw.PersonName,
			vrraw.ShortString, vrraw.ShortText, vrraw.UnlimitedText,
			vrraw.DecimalString, vrraw.CodeString, vrraw.Time, vrraw.ApplicationEntity,
			vrraw.IntegerString, vrraw.Unknown:
			if err := w.WriteString(" "); err != nil { // https://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_6.2
				return err
			}
		default:
			if err := w.WriteByte(0); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeBytes(w *dicomio.Writer, values []byte, vr string) error {
	var err error
	switch vr {
	case vrraw.OtherWord, vrraw.Unknown:
		err = writeOtherWordString(w, values)
	case vrraw.OtherByte:
		err = writeOtherByteString(w, values)
	default:
		return ErrorMismatchValueTypeAndVR
	}
	if err != nil {
		return err
	}
	return nil
}

func writeInts(w *dicomio.Writer, values []int, vr string) error {
	for _, value := range values {
		switch vr {
		// TODO(suyashkumar): consider additional validation of VR=AT elements.
		case vrraw.UnsignedShort, vrraw.SignedShort, vrraw.AttributeTag:
			if err := w.WriteUInt16(uint16(value)); err != nil {
				return err
			}
		case vrraw.UnsignedLong, vrraw.SignedLong:
			if err := w.WriteUInt32(uint32(value)); err != nil {
				return err
			}
		default:
			return ErrorMismatchValueTypeAndVR
		}
	}
	return nil
}

func writeFloats(w *dicomio.Writer, v Value, vr string) error {
	if v.ValueType() != Floats {
		return ErrorUnexpectedValueType
	}
	floats := MustGetFloats(v)
	for _, fl := range floats {
		switch vr {
		case vrraw.FloatingPointSingle:
			// NOTE: this is a conversion from float64 -> float32 which may lead to a loss in precision. The assumption
			// is that the value sitting in the float64 was originally at float32 precision if the VR is FL for this
			// element. We will need to revisit this. Maybe we can detect if there will be a loss of precision and if so
			// indicate an error or warning.
			err := w.WriteFloat32(float32(fl))
			if err != nil {
				return err
			}
		case vrraw.FloatingPointDouble:
			err := w.WriteFloat64(fl)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func writePixelData(w *dicomio.Writer, t tag.Tag, value Value, vr string, vl uint32) error {
	image := MustGetPixelDataInfo(value)

	if vl == tag.VLUndefinedLength {
		if err := writeBasicOffsetTable(w, image.Offsets); err != nil {
			return err
		}
		for _, frame := range image.Frames {
			if err := writeRawItem(w, frame.EncapsulatedData.Data); err != nil {
				return err
			}
		}
		err := encodeElementHeader(w, tag.SequenceDelimitationItem, "", 0)
		if err != nil {
			return err
		}
	} else {
		if image.IntentionallySkipped {
			return nil
		}
		// For now, IntentionallyUnprocessed will only happen for Native
		// PixelData.
		if image.IntentionallyUnprocessed {
			w.WriteBytes(image.UnprocessedValueData)
			return nil
		}
		numFrames := len(image.Frames)
		numPixels := image.Frames[0].NativeData.Rows() * image.Frames[0].NativeData.Cols()
		numValues := image.Frames[0].NativeData.SamplesPerPixel()
		// Total required buffer length in bytes:
		length := numFrames * numPixels * numValues * image.Frames[0].NativeData.BitsPerSample() / 8

		buf := &bytes.Buffer{}
		buf.Grow(length)
		bo, _ := w.GetTransferSyntax()
		for frame := 0; frame < numFrames; frame++ {
			currFrameData := image.Frames[frame].NativeData
			for pixel := 0; pixel < numPixels; pixel++ {
				for sampleIdx := 0; sampleIdx < currFrameData.SamplesPerPixel(); sampleIdx++ {
					switch image.Frames[frame].NativeData.BitsPerSample() {
					case 8:
						rawSlice, ok := currFrameData.RawDataSlice().([]uint8)
						if !ok {
							return fmt.Errorf("got frame with bitsAllocated=8 but can't assert RawDataSlice to []uint8")
						}
						if err := binary.Write(buf, bo, rawSlice[(pixel*currFrameData.SamplesPerPixel())+sampleIdx]); err != nil {
							return err
						}
					case 16:
						rawSlice, ok := currFrameData.RawDataSlice().([]uint16)
						if !ok {
							return fmt.Errorf("got frame with bitsAllocated=16 but can't assert RawDataSlice to []uint16")
						}
						if err := binary.Write(buf, bo, rawSlice[(pixel*currFrameData.SamplesPerPixel())+sampleIdx]); err != nil {
							return err
						}
					case 32:
						rawSlice, ok := currFrameData.RawDataSlice().([]uint32)
						if !ok {
							return fmt.Errorf("got frame with bitsAllocated=32 but can't assert RawDataSlice to []uint32")
						}
						if err := binary.Write(buf, bo, rawSlice[(pixel*currFrameData.SamplesPerPixel())+sampleIdx]); err != nil {
							return err
						}
					default:
						return ErrorUnsupportedBitsPerSample
					}
				}
			}
		}
		// If the byte length is not even, append 1 padding byte to make it even.
		// https://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_8.1.1
		if buf.Len()%2 != 0 {
			if err := buf.WriteByte(0); err != nil {
				return err
			}
		}
		if err := w.WriteBytes(buf.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

var sequenceDelimitationItem = &Element{
	Tag:         tag.SequenceDelimitationItem,
	ValueLength: 0, // This should be 00000000H in base32
}

func writeSequence(w *dicomio.Writer, t tag.Tag, values []*SequenceItemValue, vr string, vl uint32, opts writeOptSet) error {
	// We always write out sequences using the undefined length encoding.
	// Note: we currently don't validate that the length of the sequence matches
	// the VL if it's not undefined VL.
	// More details about the sequence structure can be found at:
	// https://dicom.nema.org/medical/dicom/current/output/chtml/part05/sect_7.5.html

	// Write out the items.
	for _, seqItem := range values {
		if err := writeSequenceItem(w, t, seqItem.elements, vr, vl, opts); err != nil {
			return err
		}
	}

	// Write Sequence Delimitation Item as implicit VR
	oldBO, oldImplicit := w.GetTransferSyntax()
	w.SetTransferSyntax(oldBO, true)
	if err := writeElement(w, sequenceDelimitationItem, opts); err != nil {
		return err
	}
	w.SetTransferSyntax(oldBO, oldImplicit) // Return TS to what it was before.

	return nil
}

var sequenceItemDelimitationItem = &Element{
	Tag:         tag.ItemDelimitationItem,
	ValueLength: 0, // This should be 00000000H in base32
}

var item = &Element{
	Tag:         tag.Item,
	ValueLength: tag.VLUndefinedLength,
}

func writeSequenceItem(w *dicomio.Writer, t tag.Tag, values []*Element, vr string, vl uint32, opts writeOptSet) error {
	// Write out item header.
	if err := writeElement(w, item, opts); err != nil {
		return err
	}

	// Write out nested Dataset elements.
	for _, elem := range values {
		if err := writeElement(w, elem, opts); err != nil {
			return err
		}
	}

	// Write ItemDelimitationItem.
	return writeElement(w, sequenceItemDelimitationItem, opts)
}

func writeOtherWordString(w *dicomio.Writer, data []byte) error {
	if len(data)%2 != 0 {
		return ErrorOWRequiresEvenVL
	}
	bo, _ := w.GetTransferSyntax()
	r := dicomio.NewReader(bufio.NewReader(bytes.NewBuffer(data)), bo, int64(len(data)))

	for i := 0; i < int(len(data)/2); i++ {
		v, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		if err := w.WriteUInt16(v); err != nil {
			return err
		}
	}
	return nil
}

func writeOtherByteString(w *dicomio.Writer, data []byte) error {
	if err := w.WriteBytes(data); err != nil {
		return err
	}
	if len(data)%2 == 1 {
		if err := w.WriteByte(0); err != nil {
			return err
		}
	}
	return nil
}
