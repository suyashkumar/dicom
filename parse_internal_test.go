package dicom

import (
	"os"
	"strings"
	"testing"
)

// TestParseUntilEOFConformsToParse runs both the dicom.ParseUntilEOF and the dicom.Parse APIs against each
// testdata file and ensures the outputs are the same.
// This test lives in parse_internal_test.go because this test cannot live in the dicom_test package, due
// to some dependencies on internal valuesets for diffing.
func TestParseUntilEOFConformsToParse(t *testing.T) {
	files, err := os.ReadDir("./testdata")
	if err != nil {
		t.Fatalf("unable to read testdata/: %v", err)
	}
	for _, f := range files {
		f := f
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".dcm") {
			t.Run(f.Name(), func(t *testing.T) {
				t.Parallel()
				// Read dataset with ParseUntilEOF
				dcm := readTestdataFile(t, f.Name())
				parse_eof_dataset, err := ParseUntilEOF(dcm, nil)
				if err != nil {
					t.Errorf("dicom.ParseUntilEOF(%s) unexpected error: %v", f.Name(), err)
				}

				// Read dataset with Parse
				dcm2 := readTestdataFile(t, f.Name())
				info, err := dcm2.Stat()
				if err != nil {
					t.Errorf("Unable to stat %s. Error: %v", f.Name(), err)
				}
				parse_dataset, err := Parse(dcm2, info.Size(), nil)
				if err != nil {
					t.Errorf("dicom.Parse(%s) unexpected error: %v", f.Name(), err)
				}

				// Ensure dataset read from ParseUntilEOF and Parse are the same.
				if !parse_dataset.Equals(&parse_eof_dataset) {
					t.Errorf("dicom.Parse and dicom.ParseUntilEOF do not result in the same dataset.\nParse Dataset: %v\n\n\nParse EOF Dataset: %v", parse_dataset, parse_eof_dataset)
				}

			})
		}
	}
}

func readTestdataFile(t *testing.T, name string) *os.File {
	dcm, err := os.Open("./testdata/" + name)
	if err != nil {
		t.Errorf("Unable to open %s. Error: %v", name, err)
	}
	t.Cleanup(func() { dcm.Close() })
	return dcm
}
