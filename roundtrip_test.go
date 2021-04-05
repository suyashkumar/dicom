package dicom_test

import (
	"bytes"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/dcmtest"
	"testing"
)

// TestRoundTrip tests that we can read a dicom from disk, write the dicom back to a
// buffer, then read it back in and yield two identical datasets.
func TestRoundTrip(t *testing.T) {
	dcmtest.WalkIncludedFS(t, func(t *testing.T, tc dcmtest.FSTestCase, setupErr error) {
		// make a new buffer to write to
		rtCase := RoundTripTestCase{
			FSTestCase:     tc,
		}

		buffer := bytes.NewBuffer(make([]byte, 0, tc.DCMStat.Size()))
		err := dicom.Write(
			buffer,
			rtCase.Dataset,
			dicom.SkipVRVerification(),
			dicom.SkipValueTypeVerification(),
		)
		if err != nil {
			t.Fatalf("error writing dataset: %v", err)
		}

		rtCase.WrittenDataset, err = dicom.Parse(buffer, int64(buffer.Len()), nil)
		if err != nil {
			t.Fatalf("error re-parsing written dataset: %v", err)
		}
	})
}

// RoundTripTestCase is a test case for TestRoundTrip.
type RoundTripTestCase struct {
	// Embed the test case passed in by dcmtest.WalkIncludedFS.
	dcmtest.FSTestCase
	// WrittenDataset is the dataset resulting from writing FSTestCase.Dataset to a
	// buffer and re-parsing it.
	WrittenDataset dicom.Dataset
}
