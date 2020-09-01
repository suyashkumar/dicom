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

	// writeVRVL
	err = writeVRVL(w, elem)
	if err != nil {
		return err
	}

// 	// writeValue
// 	err = writeValue(w, elem.Value)
// 	if err != nil {
// 		return err
// 	}
//
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
	// Get the tag info
	tagInfo, err := tag.Find(elem.Tag)
	if err != nil {
		 return "UN", nil	// TODO: double check with Suyash that this is still how this should be implemented
	}
	if elem.RawValueRepresentation == "" {
		return tagInfo.VR, nil
	}
	if tagInfo.VR != elem.RawValueRepresentation {
		return "", fmt.Errorf("ERROR dicomio.veryifyElement: VR mismatch for tag %s. Element.VR=%v, but DICOM standard defines VR to be %v", elem.Tag, elem.RawValueRepresentation, tagInfo.VR)
	}

	return elem.RawValueRepresentation, nil
}

func writeTag(w dicomio.Writer, elem *Element) error {
	if elem.ValueLength % 2 != 0 {
		return fmt.Errorf("ERROR dicomio.writeTag: Value Length must be even, but for Tag=%s, ValueLength=%v", elem.Tag, elem.ValueLength)
	}
	w.WriteUInt16(elem.Tag.Group)
	w.WriteUInt16(elem.Tag.Element)
	return nil
}

func writeVRVL(w dicomio.Writer, elem *Element) error {
	if len(elem.RawValueRepresentation) != 2 {
		return fmt.Errorf("ERROR dicomio.writeVRVL: Value Representation must be of length 2, e.g. 'UN'. For tag=%s, it was RawValueRepresentation=%v", elem.Tag, elem.RawValueRepresentation)
	}

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

// func writeValue(w dicomio.Writer, value Value) error {
// 	return nil
// }
