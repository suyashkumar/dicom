package dicom

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/suyashkumar/dicom/pkg/tag"
)

// parse_internal_test.go holds tests that must exist in the dicom package (as
// opposed to dicom_test), in order to access internal entities.

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

func TestParse_InfersMissingTransferSyntax(t *testing.T) {
	dsWithMissingTS := Dataset{Elements: []*Element{
		mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
		mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
		mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
		mustNewElement(tag.Rows, []int{128}),
		mustNewElement(tag.FloatingPointValue, []float64{128.10}),
		mustNewElement(tag.DimensionIndexPointer, []int{32, 36950}),
		mustNewElement(tag.RedPaletteColorLookupTableData, make([]byte, 200)),
	}}

	cases := []struct {
		name                   string
		overrideTransferSyntax string
	}{
		{
			name:                   "Little Endian Implicit",
			overrideTransferSyntax: "1.2.840.10008.1.2",
		},
		{
			name:                   "Little Endian Explicit",
			overrideTransferSyntax: "1.2.840.10008.1.2.1",
		},
		{
			name:                   "Big Endian Explicit",
			overrideTransferSyntax: "1.2.840.10008.1.2.2",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Write out Dataset with OverrideMissingTransferSyntax option _and_
			// the testSkipWritingTransferSyntax to ensure no Transfer Syntax
			// element is written to the test dicom. The test later verifies
			// that no Transfer Syntax element was written to the metadata.
			writtenDICOM := &bytes.Buffer{}
			if err := Write(writtenDICOM, dsWithMissingTS, OverrideMissingTransferSyntax(tc.overrideTransferSyntax), testSkipWritingTransferSyntax()); err != nil {
				t.Errorf("Write(OverrideMissingTransferSyntax(%v)) returned unexpected error: %v", tc.overrideTransferSyntax, err)
			}

			parsedDS, err := ParseUntilEOF(writtenDICOM, nil)
			if err != nil {
				t.Fatalf("ParseUntilEOF returned unexpected error when reading written dataset back in: %v", err)
			}
			_, err = parsedDS.FindElementByTag(tag.TransferSyntaxUID)
			if !errors.Is(err, ErrorElementNotFound) {
				t.Fatalf("expected test dicom dataset to be missing explicit TransferSyntaxUID tag, but found one. got: %v, want: ErrorElementNotFound", err)
			}
		})
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
