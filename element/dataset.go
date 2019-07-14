package element

import (
	"encoding/binary"

	"github.com/suyashkumar/dicom/dicomio"
	"github.com/suyashkumar/dicom/dicomtag"
)

// DataSet represents contents of one DICOM file.
type DataSet struct {
	// Elements in the file, in order of appearance.
	//
	// Note: unlike pydicom, Elements also contains meta elements (those
	// with Tag.Group==2).
	Elements []*Element
}

// FindByName finds an element from the dataset given the element name,
// such as "PatientName".
func (f *DataSet) FindElementByName(name string) (*Element, error) {
	return FindByName(f.Elements, name)
}

// FindByTag finds an element from the dataset given its tag, such as
// Tag{0x0010, 0x0010}.
func (f *DataSet) FindElementByTag(tag dicomtag.Tag) (*Element, error) {
	return FindByTag(f.Elements, tag)
}

func (ds *DataSet) TransferSyntax() (bo binary.ByteOrder, implicit dicomio.IsImplicitVR, err error) {
	elem, err := ds.FindElementByTag(dicomtag.TransferSyntaxUID)
	if err != nil {
		return nil, dicomio.UnknownVR, err
	}
	transferSyntaxUID, err := elem.GetString()
	if err != nil {
		return nil, dicomio.UnknownVR, err
	}
	return dicomio.ParseTransferSyntaxUID(transferSyntaxUID)
}
