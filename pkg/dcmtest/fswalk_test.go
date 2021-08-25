package dcmtest_test

import (
	"github.com/suyashkumar/dicom/pkg/dcmtest"
	"io/fs"
	"strings"
	"testing"
)

// TestDCMTestFiles tests that our files are being embedded correctly.
func TestDCMTestFiles(t *testing.T) {
	dcmCount := 0

	err := fs.WalkDir(dcmtest.IncludedFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			t.Errorf("error walking directory path '%v': %v", path, err)
			return err
		}

		t.Log("FOUND:", d.Name())

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".dcm") {
			t.Errorf("non-dcm file found: '%v'", path)
		}

		dcmCount++

		return nil
	})

	if err != nil {
		t.Errorf("error walking test file tree: %v", err)
	}

	// Check that we have at least one .dcm file.
	if dcmCount == 0 {
		t.Errorf("no dcm files embedded")
	}
}

// TestWalkDCMFiles tests our test helper.
func TestWalkDCMFiles(t *testing.T) {
	testsCount := 0

	dcmtest.WalkIncludedFS(t, func(t *testing.T, tc dcmtest.FSTestCase, setupErr error) {
		t.Log("FILE:", tc.DCMPath)

		if setupErr != nil {
			t.Fatalf("error with entry: %v", setupErr)
		}

		if !strings.HasSuffix(tc.DCMPath, ".dcm") {
			t.Errorf("non-dcm file found: '%v'", tc.DCMPath)
		}

		testsCount++
	})

	// Check that we have at least one .dcm file.
	if testsCount == 0 {
		t.Errorf("no dcm files embedded")
	}
}
