package dicom

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/suyashkumar/dicom/dicomtag"
)

// DirectoryRecord contains info about one DICOM file mentioned in DICOMDIR.
type DirectoryRecord struct {
	Path string
	// TODO(saito): perhaps extract more fields
}

// ParseDICOMDIR parses contents of a "DICOMDIR" stored in "in".
//
// http://dicom.nema.org/medical/Dicom/2016b/output/chtml/part03/sect_F.2.2.2.html
func ParseDICOMDIR(in io.Reader) (recs []DirectoryRecord, err error) {
	bytes, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, err
	}
	p, err := NewParserFromBytes(bytes, nil)
	if err != nil {
		return nil, err
	}
	ds, err := p.Parse(ParseOptions{})
	if err != nil {
		return nil, err
	}
	seq, err := ds.FindElementByTag(dicomtag.DirectoryRecordSequence)
	if err != nil {
		return nil, err
	}
	for _, item := range seq.Value {
		path := ""
		for _, subvalue := range item.(*Element).Value {
			subelem := subvalue.(*Element)
			if subelem.Tag == dicomtag.ReferencedFileID {
				names, err := subelem.GetStrings()
				if err != nil {
					return nil, err
				}
				path = strings.Join(names, "/")
			}
		}
		if path != "" {
			recs = append(recs, DirectoryRecord{Path: path})
		}
	}
	return recs, nil
}
