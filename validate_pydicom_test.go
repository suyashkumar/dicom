// We're going to use a build tag here so that by default devs don't need to worry about
// setting up python.
//
// To run these tests pass the '-tags=pydicom' argument to go test.
//
// +build pydicom

package dicom_test

import (
	"context"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/dcmtest"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

// pydicomValidationScript is the path to the python script we are going to use to
// validate our writes.
const pydicomValidationScript = "./validate_pydicom.py"

// PythonInterpreter is the python interpreter command we are going to use. This
// allows devs to configure a virtual environment if they so choose.
var PythonInterpreter = os.Getenv("PYDICOM_TEST_INTERPRETER")

func init() {
	if PythonInterpreter == "" {
		PythonInterpreter = "python3"
	}
}

func TestPydicomValidation(t *testing.T) {
	dcmtest.WalkIncludedFS(t, func(t *testing.T, tc dcmtest.FSTestCase, setupErr error) {
		if setupErr != nil {
			t.Fatalf("setup err: %v", setupErr)
		}

		t.Parallel()

		sourceDCMPath, rewrittenDCMPath := setupTempFiles(t, tc)
		checkWithPydicom(t, sourceDCMPath, rewrittenDCMPath)
	})
}

// setupTempFiles sets up source and round-tripped dataset files to pass to the python script.
func setupTempFiles(t *testing.T, tc dcmtest.FSTestCase) (sourceDCMPath, rewrittenDCMPath string) {
	// Get a temp directory.
	tempDir := t.TempDir()

	// Create a temp file to write our dataset to.
	sourceDCMPath = filepath.Join(tempDir, "source_dataset.dcm")
	sourceDatasetFile, err := os.Create(sourceDCMPath)
	if err != nil {
		t.Fatalf("could not create temp file for source dataset write")
	}
	defer sourceDatasetFile.Close()

	// Since the file reader is already exhausted, we need to open it again
	sourceFile, err := dcmtest.IncludedFS.Open(tc.DCMPath)
	if err != nil {
		t.Fatalf("error opening source file: %v", err)
	}

	size, err := io.Copy(sourceDatasetFile, sourceFile)
	if err != nil {
		t.Fatalf("error writing source to temp file: %v", err)
	}
	if size != tc.DCMStat.Size() {
		t.Fatalf("written source file wrong size: expected %v, got %v", tc.DCMStat.Size(), size)
	}

	err = sourceDatasetFile.Close()
	if err != nil {
		t.Fatalf("error closing source data temp file: %v", err)
	}

	// Create a temp file to write our dataset to.
	rewrittenDCMPath = filepath.Join(tempDir, "re-written_dataset.dcm")
	writtenDatasetFile, err := os.Create(rewrittenDCMPath)
	if err != nil {
		t.Fatalf("could not create temp file for written dataset write")
	}
	defer writtenDatasetFile.Close()

	err = dicom.Write(
		writtenDatasetFile,
		tc.Dataset,
		dicom.SkipVRVerification(),
		dicom.SkipValueTypeVerification(),
	)
	if err != nil {
		t.Fatalf("error re-writing dataset to temp file: %v", err)
	}

	err = writtenDatasetFile.Close()
	if err != nil {
		t.Fatalf("error closing re-written dataset temp file: %v", err)
	}

	return sourceDCMPath, rewrittenDCMPath
}

// checkWithPydicom runs the chek_go_write.py script, which in turn uses pydicom to
// check if our written file:
//
//	1. Can be parsed.
//  2. Once parsed, has the same values as the original.
//
// The theory here is that we should use an existing, widely-used library to validate
// that our read/write roundtrip yielded a correctly formatted dicom.
func checkWithPydicom(t *testing.T, sourceDCMPath, rewrittenDCMPath string) {

	// Log what python interpreter we are using.
	t.Log("INTERPRETER:", PythonInterpreter)

	// Run the test script with a 5-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pyProcess := exec.CommandContext(
		ctx, PythonInterpreter, pydicomValidationScript, sourceDCMPath, rewrittenDCMPath,
	)

	// Run the process and get the output so we can log it to the test.
	output, err := pyProcess.CombinedOutput()
	if err != nil {
		t.Error("error running pydicom tests:", err)
	}

	t.Log("PYTHON OUTPUT:\n", string(output))
}
