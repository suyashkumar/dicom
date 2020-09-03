package dicom

import (
	"errors"
	"fmt"
	"io"
	"encoding/binary"
	"bytes"

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/tag"
)

var ErrorUnimplemented = errors.New("this functionality is not yet implemented")
var ErrorMismatchValueTypeAndVR = errors.New("ValueType does not match what VR required") // TODO make this error description better

// TODO(suyashkumar): consider adding an element-by-element write API.

// WriteOption represents an option that can be passed to WriteDataset. Later options will override previous options if
// applicable.
type WriteOption func(*writeOptSet)

// Write will write the input DICOM dataset to the provided io.Writer as a complete DICOM (including any header
// information if available).
func Write(out io.Writer, ds *Dataset, opts ...WriteOption) error {
	// make Writer struct
	w := dicomio.NewWriter(out, nil, false)
	var metaElems []*Element
	for _, elem := range ds.Elements {
		if elem.Tag.Group == tag.MetadataGroup {
			metaElems = append(metaElems, elem)
		}
	}

	// Write the file header with meta elements
	err := writeFileHeader(w, ds, metaElems, opts...)
	if err != nil {
		return err
	}

	// // set correct TransferSyntax
	// endian, implicit, err := ds.TransferSyntax()
	// if err != nil {
	// 	return err
	// }
	// w.SetTransferSynax(endian, implicit)	// TODO: either expand this or make this function
	//
	// // Write the rest of the elements with writeElement
	// for _, elem := range ds.Elements {
	// 	if elem.Tag != tag.MetadataGroup {
	// 		err = writeElement(w, elem, opts...)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }

	return nil
}

// SkipVRVerification returns a WriteOption that skips VR verification.
func SkipVRVerification() WriteOption {
	return func(set *writeOptSet) {
		set.skipVRVerification = true
	}
}

func SkipValueTypeVerification() WriteOption {
	return func(set *writeOptSet) {
		set.skipValueTypeVerification = true
	}
}

// writeOptSet represents the flattened option set after all WriteOptions have been applied.
type writeOptSet struct {
	skipVRVerification bool
	skipValueTypeVerification bool
}

func toOptSet(opts ...WriteOption) *writeOptSet {
	optSet := &writeOptSet{}
	for _, opt := range opts {
		opt(optSet)
	}
	return optSet
}

func writeFileHeader(w dicomio.Writer, ds *Dataset, metaElems []*Element, opts ...WriteOption) error {
	w.SetTransferSynax(binary.LittleEndian, false) // TODO: either expand this or make this function

	subWriter := dicomio.NewWriter(&bytes.Buffer{}, binary.LittleEndian, false)
	tagsUsed := make(map[tag.Tag]bool)
	tagsUsed[tag.FileMetaInformationGroupLength] = true

	writeMetaElem(w, tag.FileMetaInformationVersion, ds, &tagsUsed, opts...)
	// writeMetaElem(tag.MediaStorageSOPClassUID)
	// writeMetaElem(tag.MediaStorageSOPInstanceUID)
	// writeMetaElem(tag.TransferSyntaxUID)
	// writeMetaElem(tag.ImplementationClassUID)
	// writeMetaElem(tag.ImplementationVersionName)

	for _, elem := range metaElems {
		if elem.Tag.Group == tag.MetadataGroup {
			if _, ok := tagsUsed[elem.Tag]; !ok {
				err := writeElement(subWriter, elem, opts...)
				if err != nil {
					return err
				}
			}
		}
	}

	metaBytes := subWriter.Bytes()
	w.WriteZeros(128)
	w.WriteString("DICM")
	lengthElem, err := newElement(tag.FileMetaInformationGroupLength, uint32(len(metaBytes)))
	if err != nil {
		return err
	}

	err = writeElement(w, lengthElem, opts...) // TODO write metaelementgrouplength tag
	if err != nil {
		return err
	}
	w.WriteBytes(metaBytes)

	return nil
}

func writeElement(w dicomio.Writer, elem *Element, opts ...WriteOption) error {
	// parse WriteOption options
	options := toOptSet(opts...)
	vr := elem.RawValueRepresentation
	// SkipVRVerification
	if !options.skipVRVerification {
		vr, err := verifyVR(elem.Tag, elem.RawValueRepresentation, elem.ValueLength)
		if err != nil {
			return nil
		}
	}
	if !options.skipValueTypeVerification {
		err := verifyValueType()
		if err != nil {
			return err
		}
	}

	// writeValue to subwriter
	subWriter := NewWriter()
	err = writeValue(subWriter, elem, vr)
	if err != nil {
		return err
	}

	err := encodeElementHeader(w, elem.Tag, vr, elem.ValueLength) // TODO add error catches for rest of places encodeelemheader is called
	if err != nil {
		return err
	}

	// Write the bytes to the original writer
	w.WriteBytes(subWriter.Bytes())

	return ErrorUnimplemented
}

func writeMetaElem(w dicomio.Writer, t tag.Tag, ds *Dataset, tagsUsed *map[tag.Tag]bool, opts ...WriteOption) error {
		elem, err := ds.FindElementByTag(t)
		if err != nil {
			return err
		}
		err = writeElement(w, elem, opts...)
		if err != nil {
			return err
		}
		(*tagsUsed)[t] = true
		return nil
}

func verifyVR(t tag.Tag, vr string, vl uint32) (string, error) {
	// TODO rectify the vr and vl as either pass through variables, or altering
	// the actual element data

	// Get the tag info
	tagInfo, err := tag.Find(t)
	if err != nil {
		 return "UN", nil	// TODO: double check with Suyash that this is still how this should be implemented
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

func verifyValueType(t tag.Tag, value Value, valueType ValueType, vr string) error {
	var ok bool
	switch vr {
		case "US", "UL", "SL", "SS":
			_, ok = value.([]int)
			ok = ok && (valueType == Ints)
		case "SQ":
			_, ok = value.([]*SequenceItemValue)
			ok = ok && (valueType == Sequences)
		case "NA":
			_, ok = value.([]*Element)
			ok = ok && (valueType == SequenceItem)
		case "OW", "OB":
			if t == tag.PixelData {
				_, ok = value.(PixelDataInfo)
				ok = ok && (valueType == PixelData)
			} else {
				_, ok = value.([]byte)
				ok = ok && (valueType == Bytes)
			}
		case "FL", "FD": // TODO floats?
			return ErrorUnimplemented
		case "AT":
			fallthrough
		default:
			_, ok = value.([]string)
			ok = ok && (valueType == Strings)
		}

		if !ok {
			return fmt.Errorf("ValueType does not match the specified type in the VR")
		}
		return nil
}

func writeTag(w dicomio.Writer, t tag.Tag, vl uint32) error {
	if vl % 2 != 0 {
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
	if vl == 0xffff {
		// TODO: Ask suyash if it's okay to alter the actual element passed in
		// Another option (1) is to make a copy of elem passed in insetad of taking
		// a pointer element in writeElement
		// Option (2) is to just pass through vl and vr
		vl = &(tag.VLUndefinedLength)
	}

	if len(vr) != 2 && *vl != tag.VLUndefinedLength {
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
	writeTag(w, tag.Item, uint32(len(data)))
	writeVRVL(w, tag.Item, "NA", &uint32(len(data)))
	w.WriteBytes(data)
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

func writeValue(w dicomio.Writer, value Value, valueType ValueType, vr string, opts ...WriteOption) error {
		// TODO figure out what I'm doing about the Undefined length error that gets thrown in some of these states
		// TODO what about floats?
		switch valueType {
		case Strings:
			return writeStrings(w, value.([]string), vr) // TODO this is writeStringValue
		case Bytes:
			return writeBytes(w, value.([]byte), vr)
		case Ints:
			return writeInts(w, value.([]int), vr)
		case PixelData:
			return writePixelData(w, elem.Tag, MustGetPixelDataInfo(value), vr, elem.ValueLength)
		case SequenceItem:
			return writeSequenceItem()
		case Sequences:
			return writeSequence(w, t, value.([]*SequenceItemValue), vr, elem.ValueLength, opts...)
		default:
			return fmt.Errorf("ValueType not supported")
		}
		// TODO figure out if this actual has to be here
		return fmt.Errorf("Something went real bad, this should never be reached")
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
		if len(s) % 2 == 1 {
			switch vr {
				case "DT", "LO", "LT", "PN", "SH", "ST", "UT":
					w.WriteString(" ") // http://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_6.2
				default:
					w.WriteByte(0)
			}
		}
		return nil
}

func writeBytes(w dicomio.Writer, values []byte, vr string) error {
	var err error
	if len(values) != 1 {
		return fmt.Errorf("Expect a single value but found %v", values)
	}
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
			w.WriteUInt16(value.(uint16)) // TODO verify that there's no reason this cast would fail
		case "UL", "SL":
			w.WriteUInt32(value.(uint32))
		default:
			return ErrorMismatchValueTypeAndVR
		}
	}
	return nil
}

// w, tag.Tag, elem.Value, vr, vl
func writePixelData(w dicomio.Writer, t tag.Tag, image PixelDataInfo, vr string, vl uint32) error {
	if vl == tag.VLUndefinedLength {
		encodeElementHeader(w, t, vr, vl)
		image.writeBasicOffsetTable(w)
		for _, frame := range image.Frames {
				writeRawItem(w, frame.EncapsulatedData.Data)
		}
		writeTag(w, tag.SequenceDelimitationItem, 0)
		writeVRVL(w, tag.SequenceDelimitationItem, "", &0) // TODO figure out how to make a pointer to 0
	} else {
			numFrames := len(image.Frames)
			numPixels := len(image.Frames[0].NativeData.Data)
			numValues := len(image.Frames[0].NativeData.Data[0])
			length := numFrames * numPixels * numValues * image.Frames[0].NativeData.BitsPerSample / 8 // length in bytes

			// encodeElementHeader(e, elem.Tag, vr, uint32(length))
			encodeElementHeader(w, t, vr, uint32(length))
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

// TODO investigate why these functions work in original write
// TODO trace back how we got to casting to Element
func writeSequence(w dicomio.Writer, t tag.Tag, values []*SequenceItemValue, vr string, vl uint32, opts ...WriteOption) error {
	// if vl == tag.VLUndefinedLength {
	// 		encodeElementHeader(w, t, vr, tag.VLUndefinedLength)
	// 		err := writeSubSequence(w, values, opts...)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		encodeElementHeader(w, tag.SequenceDelimitationItem, "", 0)
	// } else {
	// 	bo, implicit := w.TransferSyntax()
	// 	subWriter := dicomio.NewWriter(&bytes.Buffer{}, bo, implicit)
	// 	err := writeSubSequence(subWriter, values, opts...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	bytes := subWriter.Bytes()
	// 	encodeElementHeader(w, t, vr, uint32s(len(bytes)))
	// 	w.WriteBytes(bytes)
	// }
	return ErrorUnimplemented
}

// TODO investigate why these functions work in original write
// TODO trace back how we got to casting to Element
func writeSequenceItem(w dicomio.Writer, t tag.Tag, values []*Element, vr string, vl uint32, opts ...WriteOption) error {
	// if vl == tag.VLUndefinedLength {
	// 		encodeElementHeader(w, t, vr, tag.VLUndefinedLength)
	// 		err := writeSubSequence(w, values, opts...)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		encodeElementHeader(w, tag.ItemDelimitationItem, "", 0)
	// } else {
	// 	bo, implicit := w.TransferSyntax()
	// 	subWriter := dicomio.NewWriter(&bytes.Buffer{}, bo, implicit)
	// 	err := writeSubSequence(subWriter, values, opts...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	bytes := subWriter.Bytes()
	// 	encodeElementHeader(w, t, vr, uint32(len(bytes)))
	// 	w.WriteBytes(bytes)
	// }
	return ErrorUnimplemented
}

// func writeSubSequence(w dicomio.Writer, values []*SequenceItemValue, opts ...WriteOption) error {
// 	for _, value := range values {
// 			subelem, ok := value.(*Element)
// 			if !ok || subelem.Tag != tag.Item {
// 				return fmt.Errorf("SQ element must be an Item, and Item values must be dicom.Element, but found %v", value)
// 			}
// 			err := writeElement(w, subelem, opts...)
// 			if err != nil {
// 				return err
// 			}
// 	}
// 	return nil
// }

// w, value Value, tag.Tag, vr string, vl uint32
func writeValuePlaceHolder(w dicomio.Writer, value Value, t tag.Tag, vr string, vl uint32) error {
	vrkind := tag.GetVRKind(t, vr)
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

	// TODO figure out how to loop through elem.Value.GetValue() interface
	// var err error
	// for _, value := range elem.Value.GetValue() {
	// 	switch vr {
	// 	case "US", "SS":
	// 		v, ok := value.(uint16)
	// 		err = dissectValue(w, v, ok, "uint16")
	// 	case "UL", "SL":
	// 		v, ok := value.(uint32)
	// 		err = dissectValue(w, v, ok, "uint32")
	// 	case "FL":
	// 		v, ok := value.(float32)
	// 		err = dissectValue(w, v, ok, "float32")
	// 	case "FD":
	// 		v, ok := value.(float64)
	// 		err = dissectValue(w, v, ok, "float64")
	// 	case "OW", "OB":
	// 		if len(elem.Value.GetValue()) != 1 {
	// 			return fmt.Errorf("%v: expect a single value but found %v",
	// 				  tag.DebugString(elem.Tag), elem.Value.GetValue())
	// 		}
	// 		if elem.ValueType != Bytes {
	// 			return fmt.Errorf("%v: expect a binary string, but found %v",
 	// 					tag.DebugString(elem.Tag), elem.Value.GetValue())
	// 		}
	//
	// 		if vr == "OW" {
	// 			err = writeOtherWordString(w, elem.Value.GetValue().([]byte))
	// 		} else if vr == "OB" {
	// 			err = writeOtherByteString(w, elem.Value.GetValue().([]byte))
	// 		}
	// 	case "AT", "NA":
	// 		fallthrough
	// 	default:
	// 		err = writeStringValue(w, elem.Value, vr)
	// 	}
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// return nil
}

func writeOtherWordString(w dicomio.Writer, data []byte) error {
	if len(data) % 2 != 0 {
		return ErrorOWRequiresEvenVL
	}
	r, err := dicomio.NewReader(bytes.NewBuffer(data), w.bo, int64(len(bytes)))
	if err != nil {
		return err
	}
	for i := 0; i < int(len(bytes) / 2); i++ {
		v := r.ReadUInt16()
		w.WriteUInt16(v)
	}
	return nil
}

func writeOtherByteString(w dicomio.Writer, data[]byte) error {
	w.WriteBytes(data)
	if len(data) % 2 == 1 {
		w.WriteByte(0)
	}
	return nil
}
