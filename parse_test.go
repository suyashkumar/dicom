package dicom_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
	"image/jpeg"
	"io"
	"os"
	"testing"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/dcmtest"
)

// TestParse is an end-to-end sanity check over DICOMs in /pkg/dcmtest/data/. Currently
// it only checks that no error is returned when parsing the files.
func TestParse(t *testing.T) {
	// This will walk all of our test data and try parsing the Dataset, so we can simply
	// report if we get passed an error.
	dcmtest.WalkIncludedFS(t, func(t *testing.T, tc dcmtest.FSTestCase, setupErr error) {
		if setupErr != nil {
			t.Fatalf("setup error: %v", setupErr)
		}
	})
}

// BenchmarkParse runs sanity benchmarks over the sample files in /pkg/dcmtest/data/.
func BenchmarkParse(b *testing.B) {
	// This will walk all of our test data and run a sub-test for each one.
	dcmtest.BenchIncludedFS(b, func(b *testing.B, tc dcmtest.FSTestCase, setupErr error) {
		// The test helper will have already parsed the file once, so we will use that
		// as validation that the dataset is good.
		if setupErr != nil {
			b.Fatalf("setup error: %v", setupErr)
		}

		dcmFile, err := tc.OpenDCMFile()
		if err != nil {
			b.Fatalf("error opening dcm file: %v", err)
		}
		defer dcmFile.Close()

		// Extract the binary data.
		binData, err := io.ReadAll(dcmFile)
		if err != nil {
			b.Fatalf("error reading dcm file: %v", err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Read from a fresh buffer each iteration.
			buffer := bytes.NewBuffer(binData)

			// Parse the dataset.
			dataset, err := dicom.Parse(buffer, tc.DCMStat.Size(), nil)
			if err != nil {
				b.Fatalf("error parsing dataset on repetition %v: %v", i, err)
			}

			// Make sure we got sane results.
			if len(dataset.Elements) == 0 {
				b.Fatalf("no elements were parsed on repetition %v", i)
			}

			if len(dataset.Elements) != len(tc.Dataset.Elements) {
				b.Fatalf(
					"element count mismatch: expected %v, got %v on repetition %v",
					len(tc.Dataset.Elements),
					len(dataset.Elements),
					i,
				)
			}
		}
	})
}

func Example_readFile() {
	// See also: dicom.Parse, which uses a more generic io.Reader API.
	dataset, _ := dicom.ParseFile("./pkg/dcmtest/data/1.dcm", nil)

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

	dataset, _ := dicom.ParseFile("./pkg/dcmtest/data/1.dcm", frameChan)
	fmt.Println(dataset) // The full dataset
}

func Example_getImageFrames() {
	dataset, _ := dicom.ParseFile("./pkg/dcmtest/data/1.dcm", nil)
	pixelDataElement, _ := dataset.FindElementByTag(tag.PixelData)
	pixelDataInfo := dicom.MustGetPixelDataInfo(pixelDataElement.Value)
	for i, fr := range pixelDataInfo.Frames {
		img, _ := fr.GetImage() // The Go image.Image for this frame
		f, _ := os.Create(fmt.Sprintf("image_%d.jpg", i))
		_ = jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
		_ = f.Close()
	}
}
