package dicom_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/suyashkumar/dicom/pkg/uid"

	"github.com/suyashkumar/dicom/pkg/frame"

	"github.com/suyashkumar/dicom"
)

// TestParse is an end-to-end sanity check over DICOMs in testdata/. Currently,
// it only checks that no error is returned when parsing the files.
func TestParse(t *testing.T) {
	files, err := os.ReadDir("./testdata")
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
	files, err := os.ReadDir("./testdata")
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

// TestParse_CEchoRQ is a test for parsing a CEchoRQ command. It checks that
// the command is parsed correctly and that the expected tags are present in
// the dataset.
func TestParse_CEchoRQ(t *testing.T) {
	commandBytes := []byte{
		// Command Group Length
		0x00, 0x00, //Tag Group
		0x00, 0x00, //Tag Element
		0x04, 0x00, 0x00, 0x00, //VL
		0x38, 0x00, 0x00, 0x00, //Value

		// AffectedSOPClassUID
		0x00, 0x00, //Tag Group
		0x02, 0x00, //Tag Element
		0x12, 0x00, 0x00, 0x00, //VL
		0x31, 0x2E, 0x32, 0x2E, 0x38, 0x34, 0x30, 0x2E, 0x31, //Value
		0x30, 0x30, 0x30, 0x38, 0x2E, 0x31, 0x2E, 0x31, 0x00, //Value

		//CommandField
		0x00, 0x00, //Tag Group
		0x00, 0x01, //Tag Element
		0x02, 0x00, 0x00, 0x00, //VL
		0x30, 0x00, //Value

		//MessageID
		0x00, 0x00, //Tag Group
		0x10, 0x01, //Tag Element
		0x02, 0x00, 0x00, 0x00, //VL
		0x01, 0x00, //Value

		//CommandDataSetType
		0x00, 0x00, //Tag Group
		0x00, 0x08, //Tag Element
		0x02, 0x00, 0x00, 0x00, //VL
		0x01, 0x01, //Value
	}

	ioReader := bytes.NewReader(commandBytes)
	dataset, err := dicom.Parse(ioReader, int64(len(commandBytes)), nil, dicom.SkipPixelData(), dicom.SkipMetadataReadOnNewParserInit())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	tags := []tag.Tag{
		{Group: 0x0000, Element: 0x0000},
		{Group: 0x0000, Element: 0x0002},
		{Group: 0x0000, Element: 0x0100},
		{Group: 0x0000, Element: 0x0110},
		{Group: 0x0000, Element: 0x0800},
	}
	for _, tag := range tags {
		_, err := dataset.FindElementByTag(tag)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
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

func TestParseFile_SkipPixelData(t *testing.T) {
	t.Run("WithSkipPixelData", func(t *testing.T) {
		runForEveryTestFile(t, func(t *testing.T, filename string) {
			dataset, err := dicom.ParseFile(filename, nil, dicom.SkipPixelData())
			if err != nil {
				t.Errorf("Unexpected error parsing dataset: %v, dataset: %v", err, dataset)
			}
			// If PixelData present in this DICOM, check if it's populated
			// correctly. The current test assumption is that if PixelData is
			// missing, it was not originally in the dicom (which we should
			// consider revisiting).
			el, err := dataset.FindElementByTag(tag.PixelData)
			if err == nil {
				pixelData := dicom.MustGetPixelDataInfo(el.Value)
				if !pixelData.IntentionallySkipped {
					t.Errorf("Expected pixelData.IntentionallySkipped=true, got false")
				}
				if got := len(pixelData.Frames); got != 0 {
					t.Errorf("unexpected frames length. got: %v, want: %v", got, 0)
				}
			}
		})
	})
	t.Run("WithNOSkipPixelData", func(t *testing.T) {
		runForEveryTestFile(t, func(t *testing.T, filename string) {
			dataset, err := dicom.ParseFile(filename, nil)
			if err != nil {
				t.Errorf("Unexpected error parsing dataset: %v, dataset: %v", err, dataset)
			}
			// If PixelData present in this DICOM, check if it's populated
			// correctly. The current test assumption is that if PixelData is
			// missing, it was not originally in the dicom (which we should
			// consider revisiting).
			el, err := dataset.FindElementByTag(tag.PixelData)
			if err == nil {
				pixelData := dicom.MustGetPixelDataInfo(el.Value)
				if pixelData.IntentionallySkipped {
					t.Errorf("Expected pixelData.IntentionallySkipped=false when SkipPixelData option not present, got true")
				}
				if len(pixelData.Frames) == 0 {
					t.Errorf("unexpected frames length when SkipPixelData=false. got: %v, want: >0", len(pixelData.Frames))
				}
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

func TestParseFile_AllowUnknownSpecificCharacterSet(t *testing.T) {
	file, err := os.CreateTemp("", "unknown_specific_character_set.dcm")
	if err != nil {
		t.Fatalf("Unexpected error when creating tempfile: %v", err)
	}

	transferSyntaxUIDElement, err := dicom.NewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian})
	if err != nil {
		t.Fatalf("Unexpected error when creating transfer syntax uid element: %v", err)
	}
	specificCharacterSetElement, err := dicom.NewElement(tag.SpecificCharacterSet, []string{"UNKNOWN"})
	if err != nil {
		t.Fatalf("Unexpected error when creating specific character set element: %v", err)
	}
	patientNameElement, err := dicom.NewElement(tag.PatientName, []string{"Bob", "Jones"})
	if err != nil {
		t.Fatalf("Unexpected error when creating patient name element: %v", err)
	}
	unknownCharacterSetDataset := dicom.Dataset{Elements: []*dicom.Element{
		transferSyntaxUIDElement,
		specificCharacterSetElement,
		patientNameElement,
	}}
	err = dicom.Write(file, unknownCharacterSetDataset)
	if err != nil {
		t.Errorf("Unexpected error writing dataset: %v", unknownCharacterSetDataset)
	}
	file.Close()

	t.Run("WithoutAllowUnknownSpecificCharacterSet", func(t *testing.T) {
		dataset, err := dicom.ParseFile(file.Name(), nil)
		if err == nil {
			t.Errorf("Expected error parsing dataset: %v", dataset)
		}
	})
	t.Run("WithAllowUnknownSpecificCharacterSet", func(t *testing.T) {
		dataset, err := dicom.ParseFile(file.Name(), nil, dicom.AllowUnknownSpecificCharacterSet())
		if err != nil {
			t.Errorf("Unexpected error parsing dataset: %v", dataset)
		}
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
			files, err := os.ReadDir("./testdata")
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

						data, err := io.ReadAll(dcm)
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
	files, err := os.ReadDir("./testdata")
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

				data, err := io.ReadAll(dcm)
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
	files, err := os.ReadDir("./testdata")
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
