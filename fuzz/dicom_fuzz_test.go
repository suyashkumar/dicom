package fuzz

import (
	"bytes"
	"testing"

	"github.com/amitbet/dicom"
)

func FuzzDICOMParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, test_file []byte) {
		reader := bytes.NewReader(test_file)
		_, _ = dicom.Parse(reader, int64(len(test_file)), nil)
	})
}
