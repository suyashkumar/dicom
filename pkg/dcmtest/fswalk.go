package dcmtest

import (
	"fmt"
	"github.com/suyashkumar/dicom"
	"io/fs"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// walkFS backs WalkAndTestFS and WalkAndBenchFS.
func walkFS(
	tb testing.TB,
	dcmFS fs.FS,
	testFunc func(t *testing.T, tc FSTestCase, setupErr error),
	benchFunc func(t *testing.B, tc FSTestCase, setupErr error),
) {
	// Walk the testdata directory.
	_ = fs.WalkDir(dcmFS, ".", func(path string, d fs.DirEntry, err error) error {
		// Only test "*.dcm" files.
		if d.IsDir() || strings.ToLower(filepath.Ext(d.Name())) != ".dcm" {
			return nil
		}

		// For each non-directory file that ends in '.dcm', run a test or bench.
		switch runner := tb.(type) {
		case *testing.T:
			runner.Run(path, func(t *testing.T) {
				// Set up our test case.
				tc, walkErr := newWalkTestCase(t, dcmFS, path)
				// Call the supplied testing function.
				testFunc(t, tc, walkErr)
			})
		case *testing.B:
			runner.Run(path, func(b *testing.B) {
				// Set up our test case.
				tc, walkErr := newWalkTestCase(b, dcmFS, path)
				// Call the supplied testing function.
				benchFunc(b, tc, walkErr)
			})
		default:
			tb.Fatalf(
				"runner was not *testing.T or *testing.B, got unsupported type: %v",
				reflect.TypeOf(tb),
			)
		}

		return nil
	})
}

// WalkFS walks the test *.dcm files found in dcmFS and runs testFunc against each file.
//
// Each entry ending in ".dcm" will be run in it's own subtest named after it's path
//
// Setup errors will not be logged to t automatically. testFunc must decide whether and
// how to log them itself.
//
// This package comes with a limited set of testdata stored in IncludedFS, if you
// wish to run tests against these files, you can use WalkIncludedFS instead.
//
// Subtests are not run in parallel by default. Users must call t.Parallel in their
// test functions to enable parallel testing.
func WalkFS(
	t *testing.T,
	dcmFS fs.FS,
	testFunc func(t *testing.T, tc FSTestCase, setupErr error),
) {
	walkFS(t, dcmFS, testFunc, nil)
}

// BenchFS is as WalkFS but takes *testing.B instead of *testing.T for benchmarking.
func BenchFS(
	b *testing.B,
	dcmFS fs.FS,
	benchFunc func(b *testing.B, tc FSTestCase, setupErr error),
) {
	walkFS(b, dcmFS, nil, benchFunc)
}

// FSTestCase is a testcase for a single dicom file.
type FSTestCase struct {
	// dcmFs holds our test filesystem.
	dcmFs fs.FS

	// DCMPath is the full source file path of DCMFile. This value will always be
	// populated, even if testFunc is passed a non-nil setupErr.
	DCMPath string

	// DCMStat is the result of calling fs.File.Info() on the DCMPath.
	DCMStat fs.FileInfo

	// Dataset is the DICOM dataset parsed from the fs.File at DCMPath.
	Dataset dicom.Dataset
}

// OpenDCMFile opens a new fs.File handler using DCMPath and the test fs.FS.
func (fsCase FSTestCase) OpenDCMFile() (fs.File, error) {
	return fsCase.dcmFs.Open(fsCase.DCMPath)
}

// newWalkTestCase sets up a new test case by parsing the source file, then re-writing
// and re-parsing the dataset.
func newWalkTestCase(tb testing.TB, dcmFS fs.FS, sourcePath string) (FSTestCase, error) {
	tc := FSTestCase{
		dcmFs:   dcmFS,
		DCMPath: sourcePath,
	}

	// Load the source file into the test case.
	file, err := tc.OpenDCMFile()
	if err != nil {
		return tc, fmt.Errorf("error opening file '%v': %w", sourcePath, err)
	}

	defer file.Close()

	tc.DCMStat, err = file.Stat()
	if err != nil {
		return tc, fmt.Errorf("error getting file stat for '%v': %w", sourcePath, err)
	}

	tc.Dataset, err = dicom.Parse(file, tc.DCMStat.Size(), nil)
	if err != nil {
		return tc, fmt.Errorf("error parsing dataset for '%v': %w", sourcePath, err)
	}

	return tc, nil
}

// WalkIncludedFS works as WalkFS, but runs against the included IncludedFS var.
func WalkIncludedFS(
	t *testing.T,
	testFunc func(t *testing.T, tc FSTestCase, setupErr error),
) {
	WalkFS(t, IncludedFS, testFunc)
}

// BenchIncludedFS works as BenchFS, but runs against the included IncludedFS var.
func BenchIncludedFS(
	b *testing.B,
	testFunc func(b *testing.B, tc FSTestCase, setupErr error),
) {
	BenchFS(b, IncludedFS, testFunc)
}
