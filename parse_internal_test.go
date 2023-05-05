package dicom

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestParseUntilEOFConformsToParse runs both the dicom.ParseUntilEOF and the dicom.Parse APIs against each
// testdata file and ensures the outputs are the same.
// This test lives in parse_internal_test.go because this test cannot live in the dicom_test package, due
// to some dependencies on internal valuesets for diffing.
func TestParseUntilEOFConformsToParse(t *testing.T) {
	files, err := ioutil.ReadDir("./testdata")
	if err != nil {
		t.Fatalf("unable to read testdata/: %v", err)
	}
	for _, f := range files {
		f := f
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".dcm") {
			t.Run(f.Name(), func(t *testing.T) {
				t.Parallel()
				// Read dataset with ParseUntilEOF
				dcm, err := os.Open("./testdata/" + f.Name())
				if err != nil {
					t.Errorf("Unable to open %s. Error: %v", f.Name(), err)
				}
				defer dcm.Close()

				if err != nil {
					t.Errorf("Unable to stat %s. Error: %v", f.Name(), err)
				}
				parse_eof_dataset, err := ParseUntilEOF(dcm, nil)
				if err != nil {
					t.Errorf("dicom.ParseUntilEOF(%s) unexpected error: %v", f.Name(), err)
				}

				// Read dataset with Parse
				dcm2, err := os.Open("./testdata/" + f.Name())
				if err != nil {
					t.Errorf("Unable to open %s. Error: %v", f.Name(), err)
				}
				defer dcm2.Close()
				info, err := dcm2.Stat()
				if err != nil {
					t.Errorf("Unable to stat %s. Error: %v", f.Name(), err)
				}
				parse_dataset, err := Parse(dcm2, info.Size(), nil)
				if err != nil {
					t.Errorf("dicom.Parse(%s) unexpected error: %v", f.Name(), err)
				}

				// Ensure dataset read from ParseUntilEOF and Parse are the same.
				cmpOpts := []cmp.Option{
					cmp.AllowUnexported(allValues...),
				}
				if diff := cmp.Diff(parse_dataset, parse_eof_dataset, cmpOpts...); diff != "" {
					t.Errorf("dicom.Parse and dicom.ParseUntilEOF do not result in the same dataset. diff: %v", diff)
				}
			})
		}
	}
}

