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

// writeOptSet represents the flattened option set after all WriteOptions have been applied.
type writeOptSet struct {
	skipVRVerification bool
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
		vr, err := verifyVR(elem)
		if err != nil {
			return nil
		}
	}

	// writeTag
	err := writeTag(w, elem)
	if err != nil {
		return nil
	}

	// writeValue to subwriter
	subWriter := NewWriter()
	err = writeValue(subWriter, elem, vr)
	if err != nil {
		return err
	}

	// writeVRVL with lv length as length of bytes in subwriter
	err = writeVRVL(w, elem)
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

func verifyVR(elem *Element) (string, error) {
	// TODO rectify the vr and vl as either pass through variables, or altering
	// the actual element data

	// Get the tag info
	tagInfo, err := tag.Find(elem.Tag)
	if err != nil {
		 return "UN", nil	// TODO: double check with Suyash that this is still how this should be implemented
	}
	if elem.RawValueRepresentation == "" {
		return tagInfo.VR, nil
	}
	if tagInfo.VR != elem.RawValueRepresentation {
		return "", fmt.Errorf("ERROR dicomio.veryifyElement: VR mismatch for tag %v. Element.VR=%v, but DICOM standard defines VR to be %v",
			 tag.DebugString(elem.Tag), elem.RawValueRepresentation, tagInfo.VR)
	}

	return elem.RawValueRepresentation, nil
}

func writeTag(w dicomio.Writer, elem *Element) error {
	if elem.ValueLength % 2 != 0 {
		return fmt.Errorf("ERROR dicomio.writeTag: Value Length must be even, but for Tag=%v, ValueLength=%v",
			tag.DebugString(elem.Tag), elem.ValueLength)
	}
	w.WriteUInt16(elem.Tag.Group)
	w.WriteUInt16(elem.Tag.Element)
	return nil
}

func writeVRVL(w dicomio.Writer, vr string, vl uint32, elem *Element) error {
	// Rectify Undefined Length VL
	if elem.ValueLength == 0xffff {
		// TODO: Ask suyash if it's okay to alter the actual element passed in
		// Another option (1) is to make a copy of elem passed in insetad of taking
		// a pointer element in writeElement
		// Option (2) is to just pass through vl and vr
		elem.ValueLength = tag.VLUndefinedLength
	}

	if len(elem.RawValueRepresentation) != 2 && elem.ValueLength != tag.VLUndefinedLength {
		return fmt.Errorf("ERROR dicomio.writeVRVL: Value Representation must be of length 2, e.g. 'UN'. For tag=%v, it was RawValueRepresentation=%v",
			 tag.DebugString(elem.Tag), elem.RawValueRepresentation)
	}

	// Write VR then VL
	_, implicit := w.GetTransferSyntax()
	if elem.Tag.Group == tag.GROUP_ItemSeq {
		implicit = true
	}
	if !implicit { // Explicit
		w.WriteString(elem.RawValueRepresentation)
		switch elem.RawValueRepresentation {
			case "NA", "OB", "OD", "OF", "OL", "OW", "SQ", "UN", "UC", "UR", "UT":
				w.WriteZeros(2)
				w.WriteUInt32(elem.ValueLength)
			default:
				w.WriteUInt16(uint16(elem.ValueLength))
		}
	} else {
		w.WriteUInt32(elem.ValueLength)
	}

	return nil
}

func writeRawItem(w dicomio.Writer, data []byte) {
	// TODO refactor writeTag, writeVRVL so that I can call them without an element
	// e.g. here need to call them
	writeTag()
	writeVRVL()
	w.WriteBytes(data)
}

func writeValue(w dicomio.Writer, elem *Element, vr string) error {
	// NOTE: vr is passed into the function instead of using elemnt.VR so that
	// the original data in elem isn't altered

	if elem.Tag == tag.PixelData {
		return writePixelData(w, elem)
	}
	if vr == "SQ" {
		return writeSequenceData() // TODO implement
	} else if vr == "NA" { // Item
		return writeItemData()
	} else {
		if elem.ValueLength == tag.VLUndefinedLength {
			return fmt.Errorf("ERROR writeValue: Undefined-length elemnt writing is not yet supported. Tag=%v, ValueRepresentation=%v, ValueLength=%v",
				 tag.DebugString(elem.Tag), elem.RawValueRepresentation, elem.ValueLength)
		}
		bo, implicit := w.GetTransferSyntax()
		subWriter := dicomio.NewWriter(&bytes.Buffer{}, bo, implicit)
		return writeGeneralData(w, elem, vr)
	}

	return nil
}

func writePixelData(w dicomio.Writer, elem *Element) error {
	image := MustGetPixelDataInfo(elem.Value)
	if elem.ValueLength == tag.VLUndefinedLength {
		writeTag(w, elem) // TODO move these calls to a encodeElementHeader function for readability
		writeVRVL(w, elem)
		image.writeBasicOffsetTable(w)
		for _, frame := range image.Frames {
				writeRawItem(w, frame.EncapsulatedData.Data)
		}
		writeTag(w, tag.SequenceDelimitationItem) // TODO make all writeTag calls take the tag
		writeVRVL(w, "", 0) // TODO vr is "" and vl is 0
	} else {
			numFrames := len(image.Frames)
			numPixels := len(image.Frames[0].NativeData.Data)
			numValues := len(image.Frames[0].NativeData.Data[0])
			length := numFrames * numPixels * numValues * image.Frames[0].NativeData.BitsPerSample / 8 // length in bytes

			// encodeElementHeader(e, elem.Tag, vr, uint32(length))
			writeTag(w, elem.Tag) // TODO make all writeTag calls take the tag
			writeVRVL(w, vr, uint32(length)) // TODO find vr
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
	}
	return nil
}

func writeSequenceData() error {return ErrorUnimplemented}

func writeItemData() error {return ErrorUnimplemented}

func writeGeneralData(w dicomio.Writer, elem *Element, vr string) error {
	// TODO figure out how to loop through elem.Value.GetValue() interface
	switch
	values :=

	var err error
	for _, value := range elem.Value.GetValue() {
		switch vr {
		case "US", "SS":
			v, ok := value.(uint16)
			err = dissectValue(w, v, ok, "uint16")
		case "UL", "SL":
			v, ok := value.(uint32)
			err = dissectValue(w, v, ok, "uint32")
		case "FL":
			v, ok := value.(float32)
			err = dissectValue(w, v, ok, "float32")
		case "FD":
			v, ok := value.(float64)
			err = dissectValue(w, v, ok, "float64")
		case "OW", "OB":
			if len(elem.Value.GetValue()) != 1 {
				return fmt.Errorf("%v: expect a single value but found %v",
					  tag.DebugString(elem.Tag), elem.Value.GetValue())
			}
			if elem.ValueType != Bytes {
				return fmt.Errorf("%v: expect a binary string, but found %v",
 						tag.DebugString(elem.Tag), elem.Value.GetValue())
			}

			if vr == "OW" {
				err = writeOtherWordString(w, elem.Value.GetValue().([]byte))
			} else if vr == "OB" {
				err = writeOtherByteString(w, elem.Value.GetValue().([]byte))
			}
		case "AT", "NA":
			fallthrough
		default:
			err = writeStringValue(w, elem.Value, vr)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func dissectValue(w dicomio.Writer, value interface{}, ok bool, dataType string) error {
	if !ok {
		return fmt.Errorf("ERROR expected %v, but found %T (%v)",
				dataType, value, value)
	}
	return w.Write(value)
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

func writeStringValue(w dicomio.Writer, value Value, vr string) error {
	if value.ValueType != Strings {
		return fmt.Errorf("Non-string value found")
	}
	s := ""
	for i, value := range value.GetValue() {
		substr, _ := value.(string)
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
