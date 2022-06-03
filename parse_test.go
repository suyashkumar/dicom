package dicom_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/suyashkumar/dicom/pkg/tag"

	"github.com/suyashkumar/dicom/pkg/frame"

	"github.com/suyashkumar/dicom"
)

// TestParse is an end-to-end sanity check over DICOMs in testdata/. Currently it only checks that no error is returned
// when parsing the files.
func TestParse(t *testing.T) {
	files, err := ioutil.ReadDir("./testdata")
	if err != nil {
		t.Fatalf("unable to read testdata/: %v", err)
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".dcm") {
			t.Run(f.Name(), func(t *testing.T) {
				dcm, err := os.Open("./testdata/" + f.Name())
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

// TestNewParserSkipMetadataReadOnNewParserInit tests that NewParser with the SkipMetadataReadOnNewParserInit option
// parses the specified dataset but not its header metadata.
func TestNewParserSkipMetadataReadOnNewParserInit(t *testing.T) {
	fStat, err := os.Stat("./testdata/1.dcm")
	if err != nil {
		t.Fatalf("Unable to stat %s. Error: %v", fStat.Name(), err)
	}

	f, err := os.Open("./testdata/1.dcm")
	if err != nil {
		t.Fatalf("Unable to open %s. Error: %v", f.Name(), err)
	}

	p, err := dicom.NewParser(f, fStat.Size(), nil, dicom.SkipMetadataReadOnNewParserInit())
	if err != nil {
		t.Fatalf("dicom.Parse(%s) unexpected error: %v", f.Name(), err)
	}

	metadata := p.GetMetadata()
	if len(metadata.Elements) > 0 {
		t.Fatalf("Found %d metadata elements despite SkipMetadataReadOnNewParserInit()", len(metadata.Elements))
	}
}

// TestNewParserSkipPixelData tests that NewParser with the SkipPixelData option
// parses the specified dataset but not its pixel data.
func TestNewParserSkipPixelData(t *testing.T) {
	for _, num := range []int{1, 2, 3, 4, 5} {
		f := fmt.Sprintf("./testdata/%d.dcm", num)
		ds, err := dicom.ParseFile(f, nil, dicom.SkipPixelData())
		if err != nil {
			t.Fatalf("dicom.Parse(%s) unexpected error: %v", f, err)
		}

		el, err := ds.FindElementByTag(tag.PixelData)
		if err != nil {
			t.Fatalf("dataset.FindElementByTag(%s) unexpected error: %v", tag.PixelData, err)
		}

		info, ok := el.Value.GetValue().(dicom.PixelDataInfo)
		if !ok {
			t.Fatalf("element.Value.GetValue() pixel data tag returns wrong value type: %T", el.Value.GetValue())
		}
		if len(info.Frames) == 0 {
			t.Fatalf("unexpected error: pixel data value has zero frames")
		}
		for i, frame := range info.Frames {
			if frame.Encapsulated {
				if len(frame.EncapsulatedData.Data) > 0 {
					t.Fatalf("frame %d encapsulated data is longer than zero bytes with SkipPixelData option enabled", i+1)
				}
			} else {
				if len(frame.NativeData.Data) > 0 {
					t.Fatalf("frame %d native data is longer than zero bytes with SkipPixelData option enabled", i+1)
				}
			}
		}
	}
}

// TestNewParserPixelData tests that NewParser with no options parses the
// specified dataset including its pixel data.
func TestNewParserPixelData(t *testing.T) {
	for _, num := range []int{1, 2, 3, 4, 5} {
		f := fmt.Sprintf("./testdata/%d.dcm", num)
		ds, err := dicom.ParseFile(f, nil)
		if err != nil {
			t.Fatalf("dicom.Parse(%s) unexpected error: %v", f, err)
		}

		el, err := ds.FindElementByTag(tag.PixelData)
		if err != nil {
			t.Fatalf("dataset.FindElementByTag(%s) unexpected error: %v", tag.PixelData, err)
		}

		info, ok := el.Value.GetValue().(dicom.PixelDataInfo)
		if !ok {
			t.Fatalf("element.Value.GetValue() pixel data tag returns wrong value type: %T", el.Value.GetValue())
		}
		if len(info.Frames) == 0 {
			t.Fatalf("unexpected error: pixel data value has zero frames")
		}
		for i, frame := range info.Frames {
			if frame.Encapsulated {
				if len(frame.EncapsulatedData.Data) == 0 {
					t.Fatalf("frame %d encapsulated data is empty with no options enabled", i+1)
				}
			} else {
				if len(frame.NativeData.Data) == 0 {
					t.Fatalf("frame %d native data is empty with no options enabled", i+1)
				}
			}
		}
	}
}

// BenchmarkParse runs sanity benchmarks over the sample files in testdata.
func BenchmarkParse(b *testing.B) {
	files, err := ioutil.ReadDir("./testdata")
	if err != nil {
		b.Fatalf("unable to read testdata/: %v", err)
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".dcm") {
			b.Run(f.Name(), func(b *testing.B) {
				dcm, err := os.Open("./testdata/" + f.Name())
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
	// See also: dicom.Parse, which uses a more generic io.Reader API.
	dataset, _ := dicom.ParseFile("testdata/1.dcm", nil)

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

	dataset, _ := dicom.ParseFile("testdata/1.dcm", frameChan)
	fmt.Println(dataset) // The full dataset
}

func Example_getImageFrames() {
	dataset, _ := dicom.ParseFile("testdata/1.dcm", nil)
	pixelDataElement, _ := dataset.FindElementByTag(tag.PixelData)
	pixelDataInfo := dicom.MustGetPixelDataInfo(pixelDataElement.Value)
	for i, fr := range pixelDataInfo.Frames {
		img, _ := fr.GetImage() // The Go image.Image for this frame
		f, _ := os.Create(fmt.Sprintf("image_%d.jpg", i))
		_ = jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
		_ = f.Close()
	}
}
