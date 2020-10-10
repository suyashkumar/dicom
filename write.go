package dicom

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/suyashkumar/dicom/pkg/uid"

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/tag"
)

var (
	// ErrorUnimplemented is for not yet finished things.
	ErrorUnimplemented = errors.New("this functionality is not yet implemented")
	// ErrorMismatchValueTypeAndVR is for when there's a discrepency betweeen the ValueType and what the VR specifies.
	ErrorMismatchValueTypeAndVR = errors.New("ValueType does not match the VR required")
	// ErrorUnexpectedValueType indicates an unexpected value type was seen.
	ErrorUnexpectedValueType = errors.New("Unexpected ValueType")
)

// TODO(suyashkumar): consider adding an element-by-element write API.

// Write will write the input DICOM dataset to the provided io.Writer as a complete DICOM (including any header
// information if available).
func Write(out io.Writer, ds Dataset, opts ...WriteOption) error {
	optSet := toOptSet(opts...)
	w := dicomio.NewWriter(out, nil, false)
	var metaElems []*Element
	for _, elem := range ds.Elements {
		if elem.Tag.Group == tag.MetadataGroup {
			metaElems = append(metaElems, elem)
		}
	}

	err := writeFileHeader(w, &ds, metaElems, *optSet)
	if err != nil {
		return err
	}

	endian, implicit, err := ds.TransferSyntax()
	if (err != nil && err != ErrorElementNotFound) || (err == ErrorElementNotFound && !optSet.defaultMissingTransferSyntax) {
		return err
	}

	if err == ErrorElementNotFound && optSet.defaultMissingTransferSyntax {
		w.SetTransferSynax(binary.LittleEndian, true)
	} else {
		w.SetTransferSynax(endian, implicit)
	}

	for _, elem := range ds.Elements {
		if elem.Tag.Group != tag.MetadataGroup {
			err = writeElement(w, elem, *optSet)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
// TransferSyntax should not raise an error, and instead the default
// LittleEndian Implicit transfer syntax should be used and written out as a
// Metadata element in the Dataset.
func DefaultMissingTransferSyntax() WriteOption {
	return func(set *writeOptSet) {
		set.defaultMissingTransferSyntax = true
	}
}

// writeOptSet represents the flattened option set after all WriteOptions have been applied.
type writeOptSet struct {
	skipVRVerification           bool
	skipValueTypeVerification    bool
	defaultMissingTransferSyntax bool
}

func toOptSet(opts ...WriteOption) *writeOptSet {
	optSet := &writeOptSet{}
	for _, opt := range opts {
		opt(optSet)
	}
	return optSet
}

func writeFileHeader(w dicomio.Writer, ds *Dataset, metaElems []*Element, opts writeOptSet) error {
	// File headers are always written in littleEndian explicit
	w.SetTransferSynax(binary.LittleEndian, false)

	metaBytes := &bytes.Buffer{}
	subWriter := dicomio.NewWriter(metaBytes, binary.LittleEndian, false)
	tagsUsed := make(map[tag.Tag]bool)
	tagsUsed[tag.FileMetaInformationGroupLength] = true

	err := writeMetaElem(subWriter, tag.FileMetaInformationVersion, ds, &tagsUsed, opts)
	if err != nil && err != ErrorElementNotFound {
		return err
	}
	err = writeMetaElem(subWriter, tag.MediaStorageSOPClassUID, ds, &tagsUsed, opts)
	if err != nil && err != ErrorElementNotFound {
		return err
	}
	err = writeMetaElem(subWriter, tag.MediaStorageSOPInstanceUID, ds, &tagsUsed, opts)
	if err != nil && err != ErrorElementNotFound {
		return err
	}
	err = writeMetaElem(subWriter, tag.TransferSyntaxUID, ds, &tagsUsed, opts)
	if err == ErrorElementNotFound && opts.defaultMissingTransferSyntax {
		// Write the default transfer syntax
		if err = writeElement(subWriter, mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}), opts); err != nil {
			return err
		}
		tagsUsed[tag.TransferSyntaxUID] = true
	} else if err != nil {
		return err
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

	w.WriteZeros(128)
	w.WriteString("DICM")
	lengthElem, err := NewElement(tag.FileMetaInformationGroupLength, []int{len(metaBytes.Bytes())})
	if err != nil {
		return err
	}

	err = writeElement(w, lengthElem, opts)
	if err != nil {
		return err
	}
	w.WriteBytes(metaBytes.Bytes())

	return nil
}

func writeElement(w dicomio.Writer, elem *Element, opts writeOptSet) error {
	vr := elem.RawValueRepresentation
	if !opts.skipVRVerification {
		var err error
		vr, err = verifyVROrDefault(elem.Tag, elem.RawValueRepresentation)
		if err != nil {
			return err
		}
	}
	if !opts.skipValueTypeVerification {
		err := verifyValueType(elem.Tag, elem.Value, vr)
		if err != nil {
			return err
		}
	}

	// writeValue to subwriter
	bo, implicit := w.GetTransferSyntax()
	data := &bytes.Buffer{}
	subWriter := dicomio.NewWriter(data, bo, implicit)
	err := writeValue(subWriter, elem.Tag, elem.Value, elem.Value.ValueType(), vr, elem.ValueLength, opts)
	if err != nil {
		return err
	}

	length := uint32(len(data.Bytes()))
	if elem.ValueLength == tag.VLUndefinedLength {
		length = tag.VLUndefinedLength
	}

	err = encodeElementHeader(w, elem.Tag, vr, length)
	if err != nil {
		return err
	}

	// Write the bytes to the original writer
	w.WriteBytes(data.Bytes())
	return nil
}

func writeMetaElem(w dicomio.Writer, t tag.Tag, ds *Dataset, tagsUsed *map[tag.Tag]bool, optSet writeOptSet) error {
	elem, err := ds.FindElementByTag(t)
	if err != nil {
		return err
	}
	err = writeElement(w, elem, optSet)
	if err != nil {
		return err
	}
	(*tagsUsed)[t] = true
	return nil
}

func verifyVROrDefault(t tag.Tag, vr string) (string, error) {
	tagInfo, err := tag.Find(t)
	if err != nil {
		return "UN", nil
	}
	if vr == "" {
		return tagInfo.VR, nil
	}
	if tagInfo.VR != vr {
		return "", fmt.Errorf("ERROR dicomio.veryifyElement: VR mismatch for tag %v. Element.VR=%v, but DICOM standard defines VR to be %v",
			tag.DebugString(t), vr, tagInfo.VR)
	}

	return vr, nil
}

func verifyValueType(t tag.Tag, value Value, vr string) error {
	valueType := value.ValueType()
	var ok bool
	switch vr {
	case "US", "UL", "SL", "SS":
		ok = valueType == Ints
	case "SQ":
		ok = valueType == Sequences
	case "NA":
		ok = valueType == SequenceItem
	case "OW", "OB":
		if t == tag.PixelData {
			ok = valueType == PixelData
		} else {
			ok = valueType == Bytes
		}
	case "FL", "FD":
		ok = valueType == Floats
	case "AT":
		fallthrough
	default:
		ok = valueType == Strings
	}

	if !ok {
		return fmt.Errorf("ValueType does not match the specified type in the VR")
	}
	return nil
}

func writeTag(w dicomio.Writer, t tag.Tag, vl uint32) error {
	if vl%2 != 0 && vl != tag.VLUndefinedLength {
		return fmt.Errorf("ERROR dicomio.writeTag: Value Length must be even, but for Tag=%v, ValueLength=%v",
			tag.DebugString(t), vl)
	}
	w.WriteUInt16(t.Group)
	w.WriteUInt16(t.Element)
	return nil
}

// TODO vl is a pointer so that we can alter the data in the original element
// so that later, we can check if elem.VL == tag.VLUndefinedLength
// Need to find a better solution for this
func writeVRVL(w dicomio.Writer, t tag.Tag, vr string, vl *uint32) error {
	// Rectify Undefined Length VL
	if *vl == 0xffff {
		// TODO: Ask suyash if it's okay to alter the actual element passed in
		// Another option (1) is to make a copy of elem passed in insetad of taking
		// a pointer element in writeElement
		// Option (2) is to just pass through vl and vr
		undefined := tag.VLUndefinedLength
		vl = &undefined
	}

	if len(vr) != 2 && *vl != tag.VLUndefinedLength && t != tag.SequenceDelimitationItem {
		return fmt.Errorf("ERROR dicomio.writeVRVL: Value Representation must be of length 2, e.g. 'UN'. For tag=%v, it was RawValueRepresentation=%v",
			tag.DebugString(t), vr)
	}

	// Write VR then VL
	_, implicit := w.GetTransferSyntax()
	if t.Group == tag.GROUP_ItemSeq {
		implicit = true
	}
	if !implicit { // Explicit
		w.WriteString(vr)
		switch vr {
		case "NA", "OB", "OD", "OF", "OL", "OW", "SQ", "UN", "UC", "UR", "UT":
			w.WriteZeros(2)
			w.WriteUInt32(*vl)
		default:
			w.WriteUInt16(uint16(*vl))
		}
	} else {
		w.WriteUInt32(*vl)
	}
	return nil
}

func writeRawItem(w dicomio.Writer, data []byte) {
	length := uint32(len(data))
	writeTag(w, tag.Item, length)
	writeVRVL(w, tag.Item, "NA", &length)
	w.WriteBytes(data)
}

func writeBasicOffsetTable(w dicomio.Writer, offsets []uint32) {
	byteOrder, implicit := w.GetTransferSyntax()
	data := &bytes.Buffer{}
	subWriter := dicomio.NewWriter(data, byteOrder, implicit)
	for _, offset := range offsets {
		subWriter.WriteUInt32(offset)
	}
	writeRawItem(w, data.Bytes())
}

func encodeElementHeader(w dicomio.Writer, t tag.Tag, vr string, vl uint32) error {
	err := writeTag(w, t, vl)
	if err != nil {
		return err
	}
	err = writeVRVL(w, t, vr, &vl)
	if err != nil {
		return err
	}
	return nil
}

func writeValue(w dicomio.Writer, t tag.Tag, value Value, valueType ValueType, vr string, vl uint32, opts writeOptSet) error {
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

func writeStrings(w dicomio.Writer, values []string, vr string) error {
	s := ""
	for i, substr := range values {
		if i > 0 {
			s += "\\"
		}
		s += substr
	}
	w.WriteString(s)
	if len(s)%2 == 1 {
		switch vr {
		case "DT", "LO", "LT", "PN", "SH", "ST", "UT", "DS", "CS", "TM", "IS", "UN":
			w.WriteString(" ") // http://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_6.2
		default:
			w.WriteByte(0)
		}
	}
	return nil
}

func writeBytes(w dicomio.Writer, values []byte, vr string) error {
	var err error
	switch vr {
	case "OW":
		err = writeOtherWordString(w, values)
	case "OB":
		err = writeOtherByteString(w, values)
	default:
		return ErrorMismatchValueTypeAndVR
	}
	if err != nil {
		return err
	}
	return nil
}

func writeInts(w dicomio.Writer, values []int, vr string) error {
	for _, value := range values {
		switch vr {
		case "US", "SS":
			w.WriteUInt16(uint16(value))
		case "UL", "SL":
			w.WriteUInt32(uint32(value))
		default:
			return ErrorMismatchValueTypeAndVR
		}
	}
	return nil
}

func writeFloats(w dicomio.Writer, v Value, vr string) error {
	if v.ValueType() != Floats {
		return ErrorUnexpectedValueType
	}
	floats := MustGetFloats(v)
	for _, fl := range floats {
		switch vr {
		case "FL":
			// NOTE: this is a conversion from float64 -> float32 which may lead to a loss in precision. The assumption
			// is that the value sitting in the float64 was originally at float32 precision if the VR is FL for this
			// element. We will need to revisit this. Maybe we can detect if there will be a loss of precision and if so
			// indicate an error or warning.
			err := w.WriteFloat32(float32(fl))
			if err != nil {
				return err
			}
		case "FD":
			err := w.WriteFloat64(fl)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func writePixelData(w dicomio.Writer, t tag.Tag, value Value, vr string, vl uint32) error {
	image := MustGetPixelDataInfo(value)
	if vl == tag.VLUndefinedLength {
		writeBasicOffsetTable(w, image.Offsets)
		for _, frame := range image.Frames {
			writeRawItem(w, frame.EncapsulatedData.Data)
		}
		err := encodeElementHeader(w, tag.SequenceDelimitationItem, "", 0)
		if err != nil {
			return err
		}
	} else {
		numFrames := len(image.Frames)
		numPixels := len(image.Frames[0].NativeData.Data)
		numValues := len(image.Frames[0].NativeData.Data[0])
		length := numFrames * numPixels * numValues * image.Frames[0].NativeData.BitsPerSample / 8 // length in bytes

		buf := new(bytes.Buffer)
		buf.Grow(length)
		for frame := 0; frame < numFrames; frame++ {
			for pixel := 0; pixel < numPixels; pixel++ {
				for value := 0; value < numValues; value++ {
					if image.Frames[frame].NativeData.BitsPerSample == 8 {
						binary.Write(buf, binary.LittleEndian, uint8(image.Frames[frame].NativeData.Data[pixel][value]))
					} else if image.Frames[frame].NativeData.BitsPerSample == 16 {
						binary.Write(buf, binary.LittleEndian, uint16(image.Frames[frame].NativeData.Data[pixel][value]))
					}
				}
			}
		}
		w.WriteBytes(buf.Bytes())
	}
	return nil
}

// TODO implement
func writeSequence(w dicomio.Writer, t tag.Tag, values []*SequenceItemValue, vr string, vl uint32, opts writeOptSet) error {
	return ErrorUnimplemented
}

// TODO implement
func writeSequenceItem(w dicomio.Writer, t tag.Tag, values []*Element, vr string, vl uint32, opts writeOptSet) error {
	return ErrorUnimplemented
}

func writeOtherWordString(w dicomio.Writer, data []byte) error {
	if len(data)%2 != 0 {
		return ErrorOWRequiresEvenVL
	}
	bo, _ := w.GetTransferSyntax()
	r, err := dicomio.NewReader(bufio.NewReader(bytes.NewBuffer(data)), bo, int64(len(data)))
	if err != nil {
		return err
	}
	for i := 0; i < int(len(data)/2); i++ {
		v, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		w.WriteUInt16(v)
	}
	return nil
}

func writeOtherByteString(w dicomio.Writer, data []byte) error {
	w.WriteBytes(data)
	if len(data)%2 == 1 {
		w.WriteByte(0)
	}
	return nil
}
