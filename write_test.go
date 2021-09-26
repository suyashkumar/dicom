package dicom

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"os"
	"testing"

	"github.com/suyashkumar/dicom/pkg/vrraw"

	"github.com/suyashkumar/dicom/pkg/frame"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/suyashkumar/dicom/pkg/uid"
)

// TestWrite tests the write package by ensuring that it is consistent with the
// Parse implementation. In particular, it is tested by writing out known
// collections of Element and reading them back in using the Parse API and
// ensuing the read in collection is equal to the initial collection.
//
// This also serves to test that the Parse implementation is consistent with the
// Write implementation (e.g. it kinda goes both ways and covers Parse too).
func TestWrite(t *testing.T) {
	cases := []struct {
		name          string
		dataset       Dataset
		extraElems    []*Element
		expectedError error
		opts          []WriteOption
		cmpOpts       []cmp.Option
	}{
		{
			name: "basic types",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
				mustNewElement(tag.Rows, []int{128}),
				mustNewElement(tag.FloatingPointValue, []float64{128.10}),
				mustNewElement(tag.DimensionIndexPointer, []int{32, 36950}),
				mustNewElement(tag.RedPaletteColorLookupTableData, []byte{0x1, 0x2, 0x3, 0x4}),
			}},
			expectedError: nil,
		},
		{
			name: "private tag",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				// We need to use an Explicit transfer syntax here or all data will be
				// read in with "UN".
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ExplicitVRLittleEndian}),
				mustNewPrivateElement(tag.Tag{0x0003, 0x0010}, vrraw.ShortText, []string{"some data"}),
			}},
			expectedError: nil,
		},
		{
			name: "sequence (2 Items with 2 values each)",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
				makeSequenceElement(tag.AddOtherSequence, [][]*Element{
					// Item 1.
					{
						{
							Tag:                    tag.PatientName,
							ValueRepresentation:    tag.VRStringList,
							RawValueRepresentation: "PN",
							Value: &stringsValue{
								value: []string{"Bob", "Jones"},
							},
						},
						{
							Tag:                    tag.Rows,
							ValueRepresentation:    tag.VRUInt16List,
							RawValueRepresentation: "US",
							Value: &intsValue{
								value: []int{100},
							},
						},
					},
					// Item 2.
					{
						{
							Tag:                    tag.PatientName,
							ValueRepresentation:    tag.VRStringList,
							RawValueRepresentation: "PN",
							Value: &stringsValue{
								value: []string{"Bob", "Jones"},
							},
						},
						{
							Tag:                    tag.Rows,
							ValueRepresentation:    tag.VRUInt16List,
							RawValueRepresentation: "US",
							Value: &intsValue{
								value: []int{100},
							},
						},
					},
				}),
			}},
			expectedError: nil,
		},
		{
			name: "sequence (2 Items with 2 values each) - skip vr verification",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
				makeSequenceElement(tag.AddOtherSequence, [][]*Element{
					// Item 1.
					{
						{
							Tag:                    tag.PatientName,
							ValueRepresentation:    tag.VRStringList,
							RawValueRepresentation: "PN",
							Value: &stringsValue{
								value: []string{"Bob", "Jones"},
							},
						},
						{
							Tag:                    tag.Rows,
							ValueRepresentation:    tag.VRUInt16List,
							RawValueRepresentation: "US",
							Value: &intsValue{
								value: []int{100},
							},
						},
					},
					// Item 2.
					{
						{
							Tag:                    tag.PatientName,
							ValueRepresentation:    tag.VRStringList,
							RawValueRepresentation: "PN",
							Value: &stringsValue{
								value: []string{"Bob", "Jones"},
							},
						},
						{
							Tag:                    tag.Rows,
							ValueRepresentation:    tag.VRUInt16List,
							RawValueRepresentation: "US",
							Value: &intsValue{
								value: []int{100},
							},
						},
					},
				}),
			}},
			expectedError: nil,
			opts:          []WriteOption{SkipVRVerification()},
		},
		{
			name: "nested sequences",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
				makeSequenceElement(tag.AddOtherSequence, [][]*Element{
					// Item 1.
					{
						{
							Tag:                    tag.PatientName,
							ValueRepresentation:    tag.VRStringList,
							RawValueRepresentation: "PN",
							Value: &stringsValue{
								value: []string{"Bob", "Jones"},
							},
						},
						// Nested Sequence.
						makeSequenceElement(tag.AnatomicRegionSequence, [][]*Element{
							{
								{
									Tag:                    tag.PatientName,
									ValueRepresentation:    tag.VRStringList,
									RawValueRepresentation: "PN",
									Value: &stringsValue{
										value: []string{"Bob", "Jones"},
									},
								},
							},
						}),
					},
				}),
			}},
			expectedError: nil,
		},
		{
			name: "nested sequences - without VR verification",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
				makeSequenceElement(tag.AddOtherSequence, [][]*Element{
					// Item 1.
					{
						{
							Tag:                    tag.PatientName,
							ValueRepresentation:    tag.VRStringList,
							RawValueRepresentation: "PN",
							Value: &stringsValue{
								value: []string{"Bob", "Jones"},
							},
						},
						// Nested Sequence.
						makeSequenceElement(tag.AnatomicRegionSequence, [][]*Element{
							{
								{
									Tag:                    tag.PatientName,
									ValueRepresentation:    tag.VRStringList,
									RawValueRepresentation: "PN",
									Value: &stringsValue{
										value: []string{"Bob", "Jones"},
									},
								},
							},
						}),
					},
				}),
			}},
			expectedError: nil,
			opts:          []WriteOption{SkipVRVerification()},
		},
		{
			name: "without transfer syntax",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
				mustNewElement(tag.Rows, []int{128}),
				mustNewElement(tag.FloatingPointValue, []float64{128.10}),
			}},
			expectedError: ErrorElementNotFound,
		},
		{
			name: "without transfer syntax with DefaultMissingTransferSyntax",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
				mustNewElement(tag.Rows, []int{128}),
				mustNewElement(tag.FloatingPointValue, []float64{128.10}),
			}},
			// This gets inserted if DefaultMissingTransferSyntax is provided:
			extraElems:    []*Element{mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian})},
			expectedError: nil,
			opts:          []WriteOption{DefaultMissingTransferSyntax()},
		},
		{
			name: "native PixelData: 8bit",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.Rows, []int{2}),
				mustNewElement(tag.Columns, []int{2}),
				mustNewElement(tag.BitsAllocated, []int{8}),
				mustNewElement(tag.NumberOfFrames, []string{"1"}),
				mustNewElement(tag.SamplesPerPixel, []int{1}),
				mustNewElement(tag.PixelData, PixelDataInfo{
					IsEncapsulated: false,
					Frames: []frame.Frame{
						{
							Encapsulated: false,
							NativeData: frame.NativeFrame{
								BitsPerSample: 8,
								Rows:          2,
								Cols:          2,
								Data:          [][]int{{1}, {2}, {3}, {4}},
							},
						},
					},
				}),
				mustNewElement(tag.FloatingPointValue, []float64{128.10}),
				mustNewElement(tag.DimensionIndexPointer, []int{32, 36950}),
			}},
			expectedError: nil,
		},
		{
			name: "native PixelData: 16bit",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.Rows, []int{2}),
				mustNewElement(tag.Columns, []int{2}),
				mustNewElement(tag.BitsAllocated, []int{16}),
				mustNewElement(tag.NumberOfFrames, []string{"1"}),
				mustNewElement(tag.SamplesPerPixel, []int{1}),
				mustNewElement(tag.PixelData, PixelDataInfo{
					IsEncapsulated: false,
					Frames: []frame.Frame{
						{
							Encapsulated: false,
							NativeData: frame.NativeFrame{
								BitsPerSample: 16,
								Rows:          2,
								Cols:          2,
								Data:          [][]int{{1}, {2}, {3}, {4}},
							},
						},
					},
				}),
			}},
			expectedError: nil,
		},
		{
			name: "native PixelData: 32bit",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.Rows, []int{2}),
				mustNewElement(tag.Columns, []int{2}),
				mustNewElement(tag.BitsAllocated, []int{32}),
				mustNewElement(tag.NumberOfFrames, []string{"1"}),
				mustNewElement(tag.SamplesPerPixel, []int{1}),
				mustNewElement(tag.PixelData, PixelDataInfo{
					IsEncapsulated: false,
					Frames: []frame.Frame{
						{
							Encapsulated: false,
							NativeData: frame.NativeFrame{
								BitsPerSample: 32,
								Rows:          2,
								Cols:          2,
								Data:          [][]int{{1}, {2}, {3}, {4}},
							},
						},
					},
				}),
			}},
			expectedError: nil,
		},
		{
			name: "native PixelData: 2 SamplesPerPixel, 2 frames",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.Rows, []int{2}),
				mustNewElement(tag.Columns, []int{2}),
				mustNewElement(tag.BitsAllocated, []int{32}),
				mustNewElement(tag.NumberOfFrames, []string{"2"}),
				mustNewElement(tag.SamplesPerPixel, []int{2}),
				mustNewElement(tag.PixelData, PixelDataInfo{
					IsEncapsulated: false,
					Frames: []frame.Frame{
						{
							Encapsulated: false,
							NativeData: frame.NativeFrame{
								BitsPerSample: 32,
								Rows:          2,
								Cols:          2,
								Data:          [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}},
							},
						},
						{
							Encapsulated: false,
							NativeData: frame.NativeFrame{
								BitsPerSample: 32,
								Rows:          2,
								Cols:          2,
								Data:          [][]int{{5, 1}, {2, 2}, {3, 3}, {4, 5}},
							},
						},
					},
				}),
			}},
			expectedError: nil,
		},
		{
			name: "encapsulated PixelData",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.BitsAllocated, []int{8}),
				setUndefinedLength(mustNewElement(tag.PixelData, PixelDataInfo{
					IsEncapsulated: true,
					Frames: []frame.Frame{
						{
							Encapsulated:     true,
							EncapsulatedData: frame.EncapsulatedFrame{Data: []byte{1, 2, 3, 4}},
						},
					},
				})),
				mustNewElement(tag.FloatingPointValue, []float64{128.10}),
				mustNewElement(tag.DimensionIndexPointer, []int{32, 36950}),
			}},
			expectedError: nil,
		},
		{
			name: "encapsulated PixelData: multiframe",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(tag.BitsAllocated, []int{8}),
				setUndefinedLength(mustNewElement(tag.PixelData, PixelDataInfo{
					IsEncapsulated: true,
					Frames: []frame.Frame{
						{
							Encapsulated:     true,
							EncapsulatedData: frame.EncapsulatedFrame{Data: []byte{1, 2, 3, 4}},
						},
						{
							Encapsulated:     true,
							EncapsulatedData: frame.EncapsulatedFrame{Data: []byte{1, 2, 3, 8}},
						},
					},
				})),
				mustNewElement(tag.FloatingPointValue, []float64{128.10}),
				mustNewElement(tag.DimensionIndexPointer, []int{32, 36950}),
			}},
			expectedError: nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := ioutil.TempFile("", "write_test.dcm")
			if err != nil {
				t.Fatalf("Unexpected error when creating tempfile: %v", err)
			}
			if err = Write(file, tc.dataset, tc.opts...); err != tc.expectedError {
				t.Errorf("Write(%v): unexpected error. got: %v, want: %v", tc.dataset, err, tc.expectedError)
			}
			file.Close()

			// Read the data back in and check for equality to the tc.dataset:
			if tc.expectedError == nil {
				f, err := os.Open(file.Name())
				if err != nil {
					t.Fatalf("Unexpected error opening file %s: %v", file.Name(), err)
				}
				info, err := f.Stat()
				if err != nil {
					t.Fatalf("Unexpected error state file: %s: %v", file.Name(), err)
				}

				readDS, err := Parse(f, info.Size(), nil)
				if err != nil {
					t.Errorf("Parse of written file, unexpected error: %v", err)
				}

				wantElems := append(tc.dataset.Elements, tc.extraElems...)

				cmpOpts := []cmp.Option{
					cmp.AllowUnexported(allValues...),
					cmpopts.IgnoreFields(Element{}, "ValueLength"),
					cmpopts.IgnoreSliceElements(func(e *Element) bool { return e.Tag == tag.FileMetaInformationGroupLength }),
					cmpopts.SortSlices(func(x, y *Element) bool { return x.Tag.Compare(y.Tag) == 1 }),
				}
				cmpOpts = append(cmpOpts, tc.cmpOpts...)

				if diff := cmp.Diff(
					wantElems,
					readDS.Elements,
					cmpOpts...,
				); diff != "" {
					t.Errorf("Reading back written dataset led to unexpected diff from source data: %s", diff)
				}
			}
		})
	}
}

func TestVerifyVR(t *testing.T) {
	cases := []struct {
		name    string
		tg      tag.Tag
		inVR    string
		wantVR  string
		wantErr bool
		opts    writeOptSet
	}{
		{
			name:    "wrong vr",
			tg:      tag.FileMetaInformationGroupLength,
			inVR:    "OB",
			wantVR:  "",
			wantErr: true,
		},
		{
			name:    "no vr",
			tg:      tag.FileMetaInformationGroupLength,
			inVR:    "",
			wantVR:  "UL",
			wantErr: false,
		},
		{
			name: "made up tag",
			tg: tag.Tag{
				Group:   0x9999,
				Element: 0x9999,
			},
			inVR:    "",
			wantVR:  "UN",
			wantErr: false,
		},
		{
			name: "private element",
			tg: tag.Tag{
				Group:   0x0003,
				Element: 0x0010,
			},
			inVR:    "DA",
			wantVR:  "DA",
			wantErr: false,
		},
		{
			name:    "skip validation - wrong vr",
			tg:      tag.PatientName,
			inVR:    "DS",
			wantVR:  "DS",
			wantErr: false,
			opts: writeOptSet{
				skipVRVerification: true,
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			vr, err := verifyVROrDefault(tc.tg, tc.inVR, tc.opts)
			if (err != nil && !tc.wantErr) || (err == nil && tc.wantErr) {
				t.Errorf("verifyVROrDefault(%v, %v), got err: %v but want err: %v", tc.tg, tc.inVR, err, tc.wantErr)
			}
			if vr != tc.wantVR {
				t.Errorf("verifyVROrDefault(%v, %v): unexpected vr. got: %v, want: %v", tc.tg, tc.inVR, vr, tc.wantVR)
			}
		})
	}
}

func TestVerifyValueType(t *testing.T) {
	cases := []struct {
		name      string
		tg        tag.Tag
		value     Value
		vr        string
		wantError bool
	}{
		{
			name:      "valid",
			tg:        tag.FileMetaInformationGroupLength,
			value:     mustNewValue([]int{128}),
			vr:        "UL",
			wantError: false,
		},
		{
			name:      "invalid vr",
			tg:        tag.FileMetaInformationGroupLength,
			value:     mustNewValue([]int{128}),
			vr:        "NA",
			wantError: true,
		},
		{
			name:      "wrong valueType",
			tg:        tag.FileMetaInformationGroupLength,
			value:     mustNewValue([]string{"str"}),
			vr:        "UL",
			wantError: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := verifyValueType(tc.tg, tc.value, tc.vr)
			if (err != nil && !tc.wantError) || (err == nil && tc.wantError) {
				t.Errorf("verifyValueType(%v, %v, %v), got err: %v but want err: %v", tc.tg, tc.value, tc.vr, err, tc.wantError)
			}
		})
	}
}

func TestWriteFloats(t *testing.T) {
	// TODO: add additional cases
	cases := []struct {
		name         string
		value        Value
		vr           string
		expectedData []byte
		expectedErr  error
	}{
		{
			name:  "float64s",
			value: &floatsValue{value: []float64{20.1019, 21.212}},
			vr:    "FD",
			// TODO: improve test expectation
			expectedData: []byte{0x60, 0x76, 0x4f, 0x1e, 0x16, 0x1a, 0x34, 0x40, 0x83, 0xc0, 0xca, 0xa1, 0x45, 0x36, 0x35, 0x40},
			expectedErr:  nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			w := dicomio.NewWriter(&buf, binary.LittleEndian, false)
			err := writeFloats(w, tc.value, tc.vr)
			if err != tc.expectedErr {
				t.Errorf("writeFloats(%v, %s) returned unexpected err. got: %v, want: %v", tc.value, tc.vr, err, tc.expectedErr)
			}
			if diff := cmp.Diff(tc.expectedData, buf.Bytes()); diff != "" {
				t.Errorf("writeFloats(%v, %s) wrote unexpected data. diff: %s", tc.value, tc.vr, diff)
				t.Errorf("% x", buf.Bytes())
			}
		})
	}

}

func TestWriteOtherWord(t *testing.T) {
	// TODO: add additional cases
	cases := []struct {
		name         string
		value        []byte
		vr           string
		expectedData []byte
		expectedErr  error
	}{
		{
			name:         "OtherWord",
			value:        []byte{0x1, 0x2, 0x3, 0x4},
			vr:           "OW",
			expectedData: []byte{0x1, 0x2, 0x3, 0x4},
			expectedErr:  nil,
		},
		{
			name:         "OtherBytes",
			value:        []byte{0x1, 0x2, 0x3, 0x4},
			vr:           "OB",
			expectedData: []byte{0x1, 0x2, 0x3, 0x4},
			expectedErr:  nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			w := dicomio.NewWriter(&buf, binary.LittleEndian, false)
			err := writeBytes(w, tc.value, tc.vr)
			if err != tc.expectedErr {
				t.Errorf("writeBytes(%v, %s) returned unexpected err. got: %v, want: %v", tc.value, tc.vr, err, tc.expectedErr)
			}
			if diff := cmp.Diff(tc.expectedData, buf.Bytes()); diff != "" {
				t.Errorf("writeBytes(%v, %s) wrote unexpected data. diff: %s", tc.value, tc.vr, diff)
				t.Errorf("% x", buf.Bytes())
			}
		})
	}

}

func setUndefinedLength(e *Element) *Element {
	e.ValueLength = tag.VLUndefinedLength
	return e
}

// TestWriteElement tests a dataset written using writer.WriteElement can be parsed into an identical dataset using NewParser.
func TestWriteElement(t *testing.T) {
	writeDS := Dataset{Elements: []*Element{
		mustNewElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
		mustNewElement(tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
		mustNewElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
		mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
		mustNewElement(tag.Rows, []int{128}),
		mustNewElement(tag.FloatingPointValue, []float64{128.10}),
		mustNewElement(tag.DimensionIndexPointer, []int{32, 36950}),
		mustNewElement(tag.RedPaletteColorLookupTableData, []byte{0x1, 0x2, 0x3, 0x4}),
	}}

	buf := bytes.Buffer{}
	w := NewWriter(&buf)
	w.SetTransferSyntax(binary.LittleEndian, true)

	for _, e := range writeDS.Elements {
		err := w.WriteElement(e)
		if err != nil {
			t.Errorf("error in writing element %s: %s", e.String(), err.Error())
		}
	}

	p, err := NewParser(&buf, int64(buf.Len()), nil, SkipMetadataReadOnNewParserInit())
	if err != nil {
		t.Fatalf("failed to create parser: %v", err)
	}

	for _, writtenElem := range writeDS.Elements {
		readElem, err := p.Next()
		if err != nil {
			t.Errorf("error in reading element %s: %s", readElem.String(), err.Error())
		}

		if diff := cmp.Diff(writtenElem, readElem, cmp.AllowUnexported(allValues...), cmpopts.IgnoreFields(Element{}, "ValueLength")); diff != "" {
			t.Errorf("unexpected diff in element: %s", diff)
		}
	}
}
