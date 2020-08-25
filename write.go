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
	// parse WriteOption options

	// if SkipVRVerification is FALSE
		// call verifyVR()

	// FOR DEBUGGING
		// doassert(elem.vr != nil, vr)

	// Write the file header with meta elements

	// Write the rest of the elements with writeElement
		// for _, elem := range ds {
		// 	writeElem(elem)
		// }

	return ErrorUnimplemented
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

func writeElement(out io.Writer, elem *Elemet) error {
	// make Writer struct

	w := dicomio.NewWriter(out)

	// writeTag
	err := writeTag(elem.Tag)

	// writeVRVL

	// writeValue

	return nil
}

func verifyVR(elem *Element) error {
	return nil
}

func writeTag(w *Writer, tag *tag.Tag) error {
	// see encodeElementHeader
	return nil
}

func writeVRVL() error {
	// see encodeElementHeader
	switch stuff {
		case "SQ":
			// write sq Tag
	}

	return nil
}

func writeValue() error {
	return nil
}
