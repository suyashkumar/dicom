package dcmtest_test

import (
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/dcmtest"
	"testing"
)

func MyDatasetFunc(dataset dicom.Dataset) error {
	return nil
}

func ExampleWalkIncludedFS() {
	// For this example we will make a mock *testing.T value. Normally this would be
	// passed into your testing function automatically by `go test`.
	t := new(testing.T)

	// Each file in dcmtest.IncludedFS will be run in it's own subtest automatically.
	dcmtest.WalkIncludedFS(t, func(t *testing.T, tc dcmtest.FSTestCase, setupErr error) {
		// On a setup error, we want to report it. setupErr is not reported to t
		// automatically, and is left to the caller to handle.
		if setupErr != nil {
			t.Fatal("setup error:", setupErr)
		}

		// Pass the pared dataset to the function we actually want to test
		err := MyDatasetFunc(tc.Dataset)
		if err != nil {
			t.Error("MyDatasetFunc error:", err)
		}
	})
}
