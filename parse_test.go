package dicom_test

import (
	"testing"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/dcmtest"
)

// TestParse is an end-to-end sanity check over DICOMs in testdata/. Currently it only checks that no error is returned
// when parsing the files.
func TestParse(t *testing.T) {
	// This will walk all of our test data and try parsing the Dataset, so we can simply
	// report if we get passed an error.
	dcmtest.WalkIncludedFS(t, func(t *testing.T, tc dcmtest.FSTestCase, setupErr error) {
		if setupErr != nil {
			t.Fatalf("setup error: %v", setupErr)
		}
	})
}

// BenchmarkParse runs sanity benchmarks over the sample files in testdata.
func BenchmarkParse(b *testing.B) {
	// This will walk all of our test data and run a sub-test for each one.
	dcmtest.BenchIncludedFS(b, func(b *testing.B, tc dcmtest.FSTestCase, setupErr error) {
		// The test helper will have already parsed the file once, so we will use that
		// as validation that the dataset is good.
		if setupErr != nil {
			b.Fatalf("setup error: %v", setupErr)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = dicom.Parse(tc.DCMFile, tc.DCMStat.Size(), nil)
		}
	})
}
