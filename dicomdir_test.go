package dicom_test

import (
	"os"
	"testing"

	"github.com/gradienthealth/go-dicom"
)

func TestParseDICOMDIR(t *testing.T) {
	in, err := os.Open("examples/testdicomdir")
	if err != nil {
		t.Fatal(err)
	}
	defer in.Close()
	records, err := dicom.ParseDICOMDIR(in)
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 14 {
		t.Errorf("Found %d records: %v", len(records), records)
	}
	r := dicom.DirectoryRecord{Path: "1331/1332/1336"}
	if records[0] != r {
		t.Errorf("Mismatched records: %v %v", records[0], r)
	}
	r = dicom.DirectoryRecord{Path: "13410/13411/13415"}
	if records[4] != r {
		t.Errorf("Mismatched records: %v %v", records[4], r)
	}
	r = dicom.DirectoryRecord{Path: "13519/13525/13531"}
	if records[13] != r {
		t.Errorf("Mismatched records: %v %v", records[13], r)
	}
}
