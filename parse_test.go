package dicom_test

import (
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/suyashkumar/dicom/pkg/frame"

	"github.com/suyashkumar/dicom"
)

// TestParse is an end-to-end sanity check over DICOMs in testfiles/. Currently it only checks that no error is returned
// when parsing the files.
func TestParse(t *testing.T) {
	files, err := ioutil.ReadDir("./testfiles")
	if err != nil {
		t.Fatalf("unable to read testfiles/: %v", err)
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".dcm") {
			t.Run(f.Name(), func(t *testing.T) {
				dcm, err := os.Open("./testfiles/" + f.Name())
				if err != nil {
					t.Errorf("Unable to open %s. Error: %v", f.Name(), err)
				}
				defer dcm.Close()
				info, err := dcm.Stat()
				if err != nil {
					t.Errorf("Unable to stat %s. Error: %v", f.Name(), err)
				}
				_, err = dicom.Parse(dcm, info.Size(), nil)
				if err != nil {
					t.Errorf("dicom.Parse(%s) unexpected error: %v", f.Name(), err)
				}
			})
		}
	}
}

// BenchmarkParse runs sanity benchmarks over the sample files in testfiles.
func BenchmarkParse(b *testing.B) {
	files, err := ioutil.ReadDir("./testfiles")
	if err != nil {
		b.Fatalf("unable to read testfiles/: %v", err)
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".dcm") {
			b.Run(f.Name(), func(b *testing.B) {
				dcm, err := os.Open("./testfiles/" + f.Name())
				if err != nil {
					b.Errorf("Unable to open %s. Error: %v", f.Name(), err)
				}
				defer dcm.Close()

				data, err := ioutil.ReadAll(dcm)
				if err != nil {
					b.Errorf("Unable to read file into memory for benchmark: %v", err)
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					_, _ = dicom.Parse(bytes.NewBuffer(data), int64(len(data)), nil)
				}
			})
		}
	}
}

func Example_readFile() {
	// Here, we pass nil for frameChan because we don't wish to receive
	// streaming image frames for this example.
	// See also: dicom.Parse, which uses a more generic io.Reader API.
	dataset, _ := dicom.ParseFile("testfiles/1.dcm", nil)

	// Dataset will nicely print the DICOM dataset data out of the box.
	fmt.Println(dataset)

	// Dataset is also JSON serializable out of the box.
	j, _ := json.Marshal(dataset)
	fmt.Println(j)
}

func Example_streamingFrames() {
	frameChan := make(chan *frame.Frame)

	// Go routine to handle streaming frames as they are parsed. This may be
	// useful when parsing a large DICOM with many frames from a network source,
	// where image frames can start to be processed before the entire DICOM
	// is parsed (or even read from storage).
	go func() {
		for fr := range frameChan {
			fmt.Println(fr)
		}
	}()

	dataset, _ := dicom.ParseFile("testfiles/1.dcm", frameChan)
	fmt.Println(dataset) // The full dataset
}
