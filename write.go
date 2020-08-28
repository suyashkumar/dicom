package dicom

import (
	"errors"
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
// 	options := toOptSet(opts...)
// 	// SkipVRVerification
// 	if !options.SkipVRVerification {
// 		err := verifyVR(elem)
// 		if err != nil {
// 			return nil
// 		}
// 	}
//
// 	// writeTag
// 	err = writeTag(w, elem.Tag)
// 	if err != nil {
// 		return nil
// 	}
//
// 	// writeVRVL
// 	err = writeVRVL(w, elem.ValueRepresentation, elem.ValueLength)
// 	if err != nil {
// 		return err
// 	}
//
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

// func verifyVR(elem *Element) error {
// 	return nil
// }
//
// func writeTag(w dicomio.Writer, tag *tag.Tag) error {
// 	// see encodeElementHeader
// 	return nil
// }
//
// func writeVRVL(w dicomio.Writer, vr string, vl int32) error {
// 	// see encodeElementHeader
// 	switch stuff {
// 		case "SQ":
// 			// write sq Tag
// 	}
//
// 	return nil
// }
//
// func writeValue(w dicomio.Writer, value Value) error {
// 	return nil
// }
