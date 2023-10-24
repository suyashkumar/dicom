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

// TestParse is an end-to-end sanity check over DICOMs in testdata/. Currently,
// it only checks that no error is returned when parsing the files.
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

func TestParseUntilEOF(t *testing.T) {
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

				if err != nil {
					t.Errorf("Unable to stat %s. Error: %v", f.Name(), err)
				}
				_, err = dicom.ParseUntilEOF(dcm, nil)
				if err != nil {
					t.Errorf("dicom.Parse(%s) unexpected error: %v", f.Name(), err)
				}
			})
		}
	}
}

func TestParseFile_WithTagSpecificCharacterSetToEncodingNameConverter(t *testing.T) {
	aDicomFileHavingNonStandardSpecificCharacterSet := "./testdata/malformed/1.dcm"
	t.Run("WithCustomConverterParseSucceed", func(t *testing.T) {
		converter := func(origin string) string {
			substitutions := map[string]string{
				"ISO_2022_IR_6": "ISO 2022 IR 6",
			}
			if sub, ok := substitutions[origin]; ok {
				return sub
			}
			return origin
		}
		dataset, err := dicom.ParseFile(
			aDicomFileHavingNonStandardSpecificCharacterSet,
			nil,
			dicom.WithTagSpecificCharacterSetToEncodingNameConverter(converter))
		if err != nil {
			t.Fatalf("parsing dataset: %v", err)
		}
		charsetElem, err := dataset.FindElementByTag(tag.SpecificCharacterSet)
		if err != nil {
			t.Fatalf("find element after parsing: %v", err)
		}
		charset := dicom.MustGetStrings(charsetElem.Value)[0]
		if charset != "ISO 2022 IR 6" {
			t.Fatalf("unexpected charset after parsing: %v", charset)
		}
	})
	t.Run("WithNoCustomerConverterParseFail", func(t *testing.T) {
		expectedErrMsg := "ParseSpecificCharacterSet: Unknown character set"
		_, err := dicom.ParseFile(aDicomFileHavingNonStandardSpecificCharacterSet, nil)
		if err == nil {
			t.Fatalf("expect error %s but got empty one", expectedErrMsg)
		}
		if !strings.Contains(err.Error(), expectedErrMsg) {
			t.Fatalf("unexpected parsing error: %v", err)
		}
	})
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

func TestParseFile_SkipPixelData(t *testing.T) {
	t.Run("WithSkipPixelData", func(t *testing.T) {
		runForEveryTestFile(t, func(t *testing.T, filename string) {
			dataset, err := dicom.ParseFile(filename, nil, dicom.SkipPixelData())
			if err != nil {
				t.Errorf("Unexpected error parsing dataset: %v", dataset)
			}
			el, err := dataset.FindElementByTag(tag.PixelData)
			if err != nil {
				t.Errorf("Unexpected error when finding PixelData in Dataset: %v", err)
			}
			pixelData := dicom.MustGetPixelDataInfo(el.Value)
			if !pixelData.IntentionallySkipped {
				t.Errorf("Expected pixelData.IntentionallySkipped=true, got false")
			}
			if got := len(pixelData.Frames); got != 0 {
				t.Errorf("unexpected frames length. got: %v, want: %v", got, 0)
			}
		})
	})
	t.Run("WithNOSkipPixelData", func(t *testing.T) {
		runForEveryTestFile(t, func(t *testing.T, filename string) {
			dataset, err := dicom.ParseFile(filename, nil)
			if err != nil {
				t.Errorf("Unexpected error parsing dataset: %v", dataset)
			}
			el, err := dataset.FindElementByTag(tag.PixelData)
			if err != nil {
				t.Errorf("Unexpected error when finding PixelData in Dataset: %v", err)
			}
			pixelData := dicom.MustGetPixelDataInfo(el.Value)
			if pixelData.IntentionallySkipped {
				t.Errorf("Expected pixelData.IntentionallySkipped=false when SkipPixelData option not present, got true")
			}
			if len(pixelData.Frames) == 0 {
				t.Errorf("unexpected frames length when SkipPixelData=false. got: %v, want: >0", len(pixelData.Frames))
			}
		})
	})
}

func TestParseFile_SkipProcessingPixelDataValue(t *testing.T) {
	t.Run("WithSkipProcessingPixelDataValue", func(t *testing.T) {
		runForEveryTestFile(t, func(t *testing.T, filename string) {
			dataset, err := dicom.ParseFile(filename, nil, dicom.SkipProcessingPixelDataValue())
			if err != nil {
				t.Errorf("Unexpected error parsing dataset: %v", dataset)
			}
			el, err := dataset.FindElementByTag(tag.PixelData)
			if err != nil {
				t.Errorf("Unexpected error when finding PixelData in Dataset: %v", err)
			}
			pixelData := dicom.MustGetPixelDataInfo(el.Value)
			if !pixelData.IntentionallyUnprocessed {
				t.Errorf("Expected pixelData.IntentionallyUnprocessed=true, got false")
			}
			if got := len(pixelData.Frames); got != 0 {
				t.Errorf("unexpected frames length. got: %v, want: %v", got, 0)
			}
		})
	})
	t.Run("WithNOSkipProcessingPixelDataValue", func(t *testing.T) {
		runForEveryTestFile(t, func(t *testing.T, filename string) {
			dataset, err := dicom.ParseFile(filename, nil)
			if err != nil {
				t.Errorf("Unexpected error parsing dataset: %v", dataset)
			}
			el, err := dataset.FindElementByTag(tag.PixelData)
			if err != nil {
				t.Errorf("Unexpected error when finding PixelData in Dataset: %v", err)
			}
			pixelData := dicom.MustGetPixelDataInfo(el.Value)
			if pixelData.IntentionallyUnprocessed {
				t.Errorf("Expected pixelData.IntentionallyUnprocessed=false when TestParseFile_SkipProcessingPixelDataValue option not present, got true")
			}
			if len(pixelData.Frames) == 0 {
				t.Errorf("unexpected frames length when TestParseFile_SkipProcessingPixelDataValue=false. got: %v, want: >0", len(pixelData.Frames))
			}
		})
	})
	t.Run("WithAllowErrorMetaElementGroupLength", func(t *testing.T) {
		runForEveryTestFile(t, func(t *testing.T, filename string) {
			dataset, err := dicom.ParseFile(filename, nil, dicom.AllowMissingMetaElementGroupLength())
			if err != nil {
				t.Errorf("Unexpected error parsing dataset: %v", dataset)
			}
		})
	})
}

// BenchmarkParse runs sanity benchmarks over the sample files in testdata.
func BenchmarkParse(b *testing.B) {
	cases := []struct {
		name string
		opts []dicom.ParseOption
	}{
		{
			name: "NoOptions",
		},
		{
			name: "SkipPixelData",
			opts: []dicom.ParseOption{dicom.SkipPixelData()},
		},
		{
			name: "SkipProcessingPixelDataValue",
			opts: []dicom.ParseOption{dicom.SkipProcessingPixelDataValue()},
		},
	}
	for _, tc := range cases {
		b.Run(tc.name, func(b *testing.B) {
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
							_, _ = dicom.Parse(bytes.NewBuffer(data), int64(len(data)), nil, tc.opts...)
						}

					})
				}
			}
		})
	}
}

func BenchmarkParser_NextAPI(b *testing.B) {
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
					r := bytes.NewReader(data)
					p, _ := dicom.NewParser(r, int64(len(data)), nil)
					var err error
					for err == nil {
						_, err = p.Next()
					}
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

func runForEveryTestFile(t *testing.T, testFunc func(t *testing.T, filename string)) {
	files, err := ioutil.ReadDir("./testdata")
	if err != nil {
		t.Fatalf("unable to read testdata/: %v", err)
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".dcm") {
			fullName := "./testdata/" + f.Name()
			t.Run(fullName, func(t *testing.T) {
				testFunc(t, fullName)
			})
		}
	}
}
