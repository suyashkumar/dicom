package dicom

import (
	"errors"
	"io"

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
	w := dicomio.NewWriter(out, nil)
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

	// set correct TransferSyntax
	endian, implicit, err := ds.TransferSyntax()
	if err != nil {
		return err
	}
	w.SetTransferSynax(endian, implicit)	// TODO: either expand this or make this function

	// Write the rest of the elements with writeElement
	for _, elem := range ds.Elements {
		if elem.Tag != tag.MetadataGroup {
			err = writeElement(w, elem, opts...)
			if err != nil {
				return err
			}
		}
	}

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

func writeFileHeader(w *dicomio.Writer, ds *Dataset, metaElems []*Element, opts ...WriteOption) error {
	w.SetTransferSynax(binary.LittleEndian, false) // TODO: either expand this or make this function

	subWriter := dicomio.NewWriter(&bytes.Buffer{}, binary.LittleEndian, false)
	tagsUsed := make(map[tag.Tag]bool)
	tagsUsed[tag.FileMetaInformationGroupLength] = true

	writeMetaElem := func(tag tag.Tag) error {
		elem, err := ds.FindElementByTag(tag)
		if err != nil {
			return err
		}
		tagsUsed[tag] = true // TODO this hsouldn't actually be changed until writeElement returns nil for error
		return writeElement(subWriter, elem, opts...)
	}

	writeMetaElem(dicomtag.FileMetaInformationVersion)
	writeMetaElem(dicomtag.MediaStorageSOPClassUID)
	writeMetaElem(dicomtag.MediaStorageSOPInstanceUID)
	writeMetaElem(dicomtag.TransferSyntaxUID)
	writeMetaElem(dicomtag.ImplementationClassUID)
	writeMetaElem(dicomtag.ImplementationVersionName)

	for _, elem := range metaElems {
		if elem.Tag.Group == dicomtag.MetadataGroup {
			if _, ok := tagsUsed[elem.Tag]; !ok {
				err = writeElement(subEncoder, elem, opts...)
				if err != nil {
					return err
				}
			}
		}
	}

	metaBytes := subWriter.Bytes()
	w.WriteZeros(128)
	w.WriteString("DICM")
	lengthElem  // TODO make lengthElem
	err = writeElement(w, lengthElem, opts...) // TODO write metaelementgrouplength tag
	if err != nil {
		return err
	}
	w.WriteBytes(metaBytes)

	return nil
}

func writeElement(w *dicomio.Writer, elem *Elemet, opts ...WriteOption) error {
	// parse WriteOption options
	options := toOptSet(opts...)
	// SkipVRVerification
	if !options.SkipVRVerification {
		err := verifyVR(elem)
		if err != nil {
			return nil
		}
	}

	// writeTag
	err = writeTag(w, elem.Tag)
	if err != nil {
		return nil
	}

	// writeVRVL
	err = writeVRVL(w, elem.ValueRepresentation, elem.ValueLength)
	if err != nil {
		return err
	}

	// writeValue
	err = writeValue(w, elem.Value)
	if err != nil {
		return err
	}

	return nil
}

// func writeMetaElem(w *Writer, tag tag.Tag, tagsUsed *map[tag.Tag]bool) error {
// 	tagInfo := tag.Find(tag)
// 	elem := Element{
// 		Tag: tag,
// 		ValueRepresentation: tag.GetVRKind(tagInfo.VR),
// 		RawValueRepresentation: tagInfo.VR,
//
//
// 	}
//
// 	if elem, err := element.FindByTag(metaElems, tag); err == nil {
// 			Element(subEncoder, elem, opts...)
// 		} else {
// 			subEncoder.SetErrorf("%v not found in metaelems: %v", dicomtag.DebugString(tag), err)
// 		}
// 		tagsUsed[tag] = true
// }

func verifyVR(elem *Element) error {
	return nil
}

func writeTag(w *Writer, tag *tag.Tag) error {
	// see encodeElementHeader
	return nil
}

func writeVRVL(w *Writer, vr string, vl int32) error {
	// see encodeElementHeader
	switch stuff {
		case "SQ":
			// write sq Tag
	}

	return nil
}

func writeValue(w *Writer, value Value) error {
	return nil
}
