package dicom

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"math/rand"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/suyashkumar/dicom/pkg/vrraw"

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/frame"

	"github.com/google/go-cmp/cmp"

	"github.com/suyashkumar/dicom/pkg/tag"
)

func TestReadTag(t *testing.T) {
	cases := []struct {
		name    string
		data    []byte
		wantTag tag.Tag
		wantErr error
	}{
		{
			name:    "basic",
			data:    buildTagData(t, tag.Rows),
			wantTag: tag.Rows,
			wantErr: nil,
		},
		{
			name:    "custom",
			data:    buildTagData(t, tag.Tag{0x0011, 0x0010}),
			wantTag: tag.Tag{0x0011, 0x0010},
			wantErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := bytes.NewBuffer(tc.data)

			r := &reader{
				rawReader: dicomio.NewReader(bufio.NewReader(data), binary.LittleEndian, int64(data.Len())),
			}
			gotTag, err := r.readTag()
			if err != tc.wantErr {
				t.Errorf("TestReadTag: unexpected err. got: %v, want: %v", err, tc.wantErr)
			}

			if !gotTag.Equals(tc.wantTag) {
				t.Errorf("TestReadTag: unexpected output. got: %v, want: %v", gotTag, tc.wantTag)
			}
		})
	}

}

func TestReadFloat_float64(t *testing.T) {
	cases := []struct {
		name        string
		floats      []float64
		VR          string
		want        Value
		expectedErr error
	}{
		{
			name:        "float64",
			floats:      []float64{20.1, 32.22},
			VR:          vrraw.FloatingPointDouble,
			want:        &floatsValue{value: []float64{20.1, 32.22}},
			expectedErr: nil,
		},
		{
			name:        "float64 with wrong VR",
			floats:      []float64{20.1, 32.22},
			VR:          "XX",
			want:        nil,
			expectedErr: errorUnableToParseFloat,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := bytes.Buffer{}
			for _, fl := range tc.floats {
				if err := binary.Write(&data, binary.LittleEndian, fl); err != nil {
					t.Errorf("TestReadFloat: Unable to setup test buffer")
				}
			}

			r := &reader{
				rawReader: dicomio.NewReader(bufio.NewReader(&data), binary.LittleEndian, int64(data.Len())),
			}
			got, err := r.readFloat(tag.Tag{}, tc.VR, uint32(data.Len()))
			if err != tc.expectedErr {
				t.Fatalf("readFloat(r, tg, %s, %d) got unexpected error: got: %v, want: %v", tc.VR, data.Len(), err, tc.expectedErr)
			}
			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(floatsValue{})); diff != "" {
				t.Errorf("readFloat(r, tg, %s, %d) unexpected diff: %s", tc.VR, data.Len(), diff)
			}
		})
	}
}

func TestReadFloat_float32(t *testing.T) {
	cases := []struct {
		name        string
		floats      []float32
		VR          string
		want        Value
		expectedErr error
	}{
		{
			name:        "float32",
			floats:      []float32{20.1001, 32.22},
			VR:          vrraw.FloatingPointSingle,
			want:        &floatsValue{value: []float64{20.1001, 32.22}},
			expectedErr: nil,
		},
		{
			name:        "float32 with wrong VR",
			floats:      []float32{20.1001, 32.22},
			VR:          "XX",
			want:        nil,
			expectedErr: errorUnableToParseFloat,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := bytes.Buffer{}
			for _, fl := range tc.floats {
				if err := binary.Write(&data, binary.LittleEndian, fl); err != nil {
					t.Errorf("TestReadFloat: Unable to setup test buffer")
				}
			}

			r := &reader{
				rawReader: dicomio.NewReader(bufio.NewReader(&data), binary.LittleEndian, int64(data.Len())),
			}
			got, err := r.readFloat(tag.Tag{}, tc.VR, uint32(data.Len()))
			if err != tc.expectedErr {
				t.Fatalf("readFloat(r, tg, %s, %d) got unexpected error: got: %v, want: %v", tc.VR, data.Len(), err, tc.expectedErr)
			}
			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(floatsValue{})); diff != "" {
				t.Errorf("readFloat(r, tg, %s, %d) unexpected diff: %s", tc.VR, data.Len(), diff)
			}
		})
	}
}

func TestReadOWBytes(t *testing.T) {
	cases := []struct {
		name        string
		bytes       []byte
		VR          string
		want        Value
		expectedErr error
	}{
		{
			name:        "OW VR with even-number bytes",
			bytes:       []byte{0x1, 0x2, 0x3, 0x4},
			VR:          vrraw.OtherWord,
			want:        &bytesValue{value: []byte{0x1, 0x2, 0x3, 0x4}},
			expectedErr: nil,
		},
		{
			name:        "UN VR even-number bytes",
			bytes:       []byte{0x1, 0x2, 0x3, 0x4},
			VR:          vrraw.Unknown,
			want:        &bytesValue{value: []byte{0x1, 0x2, 0x3, 0x4}},
			expectedErr: nil,
		},
		{
			name:        "error on odd-number bytes",
			bytes:       []byte{0x1, 0x2, 0x3},
			VR:          vrraw.OtherWord,
			want:        nil,
			expectedErr: ErrorOWRequiresEvenVL,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := bytes.Buffer{}
			if err := binary.Write(&data, binary.LittleEndian, tc.bytes); err != nil {
				t.Errorf("TestReadOWBytes: Unable to setup test buffer")
			}

			r := &reader{rawReader: dicomio.NewReader(bufio.NewReader(&data), binary.LittleEndian, int64(data.Len()))}
			got, err := r.readBytes(tag.Tag{}, tc.VR, uint32(data.Len()))
			if err != tc.expectedErr {
				t.Fatalf("readBytes(r, tg, %s, %d) got unexpected error: got: %v, want: %v", tc.VR, data.Len(), err, tc.expectedErr)
			}
			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(bytesValue{})); diff != "" {
				t.Errorf("readBytes(r, tg, %s, %d) unexpected diff: %s", tc.VR, data.Len(), diff)
			}
		})
	}
}

func TestReadNativeFrames(t *testing.T) {
	cases := []struct {
		Name              string
		existingData      Dataset
		data              []uint16
		dataBytes         []byte
		expectedPixelData *PixelDataInfo
		expectedError     error
		pixelVLOverride   uint32
		parseOptSet       parseOptSet
	}{
		{
			Name: "5x5, 1 frame, 1 samples/pixel",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{5}),
				mustNewElement(tag.Columns, []int{5}),
				mustNewElement(tag.NumberOfFrames, []string{"1"}),
				mustNewElement(tag.BitsAllocated, []int{16}),
				mustNewElement(tag.SamplesPerPixel, []int{1}),
			}},
			data: []uint16{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedPixelData: &PixelDataInfo{
				IsEncapsulated: false,
				Frames: []frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 16,
							Rows:          5,
							Cols:          5,
							Data:          [][]int{{1}, {2}, {3}, {4}, {5}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			Name: "2x2, 3 frames, 1 samples/pixel",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{2}),
				mustNewElement(tag.Columns, []int{2}),
				mustNewElement(tag.NumberOfFrames, []string{"3"}),
				mustNewElement(tag.BitsAllocated, []int{16}),
				mustNewElement(tag.SamplesPerPixel, []int{1}),
			}},
			data: []uint16{1, 2, 3, 2, 1, 2, 3, 2, 1, 2, 3, 0},
			expectedPixelData: &PixelDataInfo{
				IsEncapsulated: false,
				Frames: []frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 16,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1}, {2}, {3}, {2}},
						},
					},
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 16,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1}, {2}, {3}, {2}},
						},
					},
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 16,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1}, {2}, {3}, {0}},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			Name: "2x2, 2 frames, 2 samples/pixel",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{2}),
				mustNewElement(tag.Columns, []int{2}),
				mustNewElement(tag.NumberOfFrames, []string{"2"}),
				mustNewElement(tag.BitsAllocated, []int{16}),
				mustNewElement(tag.SamplesPerPixel, []int{2}),
			}},
			data: []uint16{1, 2, 3, 2, 1, 2, 3, 2, 1, 2, 3, 2, 1, 2, 3, 5},
			expectedPixelData: &PixelDataInfo{
				IsEncapsulated: false,
				Frames: []frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 16,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 2}},
						},
					},
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 16,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 5}},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			Name: "insufficient bytes, uint32",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{2}),
				mustNewElement(tag.Columns, []int{2}),
				mustNewElement(tag.NumberOfFrames, []string{"2"}),
				mustNewElement(tag.BitsAllocated, []int{32}),
				mustNewElement(tag.SamplesPerPixel, []int{2}),
			}},
			data:              []uint16{1, 2, 3, 2, 1, 2, 3, 2, 1, 2, 3, 2, 1, 2, 3},
			expectedPixelData: nil,
			expectedError:     ErrorMismatchPixelDataLength,
		},
		{
			Name: "redundant bytes, uint32",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{2}),
				mustNewElement(tag.Columns, []int{2}),
				mustNewElement(tag.NumberOfFrames, []string{"1"}),
				mustNewElement(tag.BitsAllocated, []int{32}),
				mustNewElement(tag.SamplesPerPixel, []int{2}),
			}},
			data:              []uint16{1, 2, 3, 2, 1, 2, 3, 2, 1, 2, 3, 2, 1, 2, 3, 2, 2},
			expectedPixelData: nil,
			expectedError:     ErrorMismatchPixelDataLength,
		},
		{
			Name: "redundant bytes, uint32 with allowing mismatch length",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{2}),
				mustNewElement(tag.Columns, []int{2}),
				mustNewElement(tag.NumberOfFrames, []string{"1"}),
				mustNewElement(tag.BitsAllocated, []int{32}),
				mustNewElement(tag.SamplesPerPixel, []int{2}),
			}},
			data: []uint16{1, 2, 3, 2, 1, 2, 3, 2, 1, 2, 3, 2, 1, 2, 3, 2, 2},
			expectedPixelData: &PixelDataInfo{
				ParseErr: ErrorMismatchPixelDataLength,
				Frames: []frame.Frame{
					{
						EncapsulatedData: frame.EncapsulatedFrame{
							Data: []byte{1, 0, 2, 0, 3, 0, 2, 0, 1, 0, 2, 0, 3, 0, 2, 0, 1, 0, 2, 0, 3, 0, 2, 0, 1, 0, 2, 0, 3, 0, 2, 0, 2, 0},
						},
					},
				},
			},
			parseOptSet:   parseOptSet{allowMismatchPixelDataLength: true},
			expectedError: nil,
		},
		{
			Name: "missing Columns",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{5}),
				mustNewElement(tag.NumberOfFrames, []string{"1"}),
				mustNewElement(tag.BitsAllocated, []int{16}),
				mustNewElement(tag.SamplesPerPixel, []int{1}),
			}},
			data:              []uint16{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedPixelData: nil,
			expectedError:     ErrorElementNotFound,
		},
		{
			Name: "unsupported BitsAllocated",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{5}),
				mustNewElement(tag.Columns, []int{2}),
				mustNewElement(tag.NumberOfFrames, []string{"1"}),
				mustNewElement(tag.BitsAllocated, []int{24}),
				mustNewElement(tag.SamplesPerPixel, []int{1}),
			}},
			data:              []uint16{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expectedPixelData: nil,
			expectedError:     ErrorUnsupportedBitsAllocated,
		},
		{
			Name: "3x3, 3 frames, 1 samples/pixel, data bytes with padded 0",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{3}),
				mustNewElement(tag.Columns, []int{3}),
				mustNewElement(tag.NumberOfFrames, []string{"3"}),
				mustNewElement(tag.BitsAllocated, []int{8}),
				mustNewElement(tag.SamplesPerPixel, []int{1}),
			}},
			dataBytes: []byte{11, 12, 13, 21, 22, 23, 31, 32, 33, 11, 12, 13, 21, 22, 23, 31, 32, 33, 11, 12, 13, 21, 22, 23, 31, 32, 33, 0}, // there is a 28th byte to make total value length even, as required by DICOM spec
			expectedPixelData: &PixelDataInfo{
				IsEncapsulated: false,
				Frames: []frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 8,
							Rows:          3,
							Cols:          3,
							Data:          [][]int{{11}, {12}, {13}, {21}, {22}, {23}, {31}, {32}, {33}},
						},
					},

					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 8,
							Rows:          3,
							Cols:          3,
							Data:          [][]int{{11}, {12}, {13}, {21}, {22}, {23}, {31}, {32}, {33}},
						},
					},
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 8,
							Rows:          3,
							Cols:          3,
							Data:          [][]int{{11}, {12}, {13}, {21}, {22}, {23}, {31}, {32}, {33}},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			Name: "1x1, 3 frames, 3 samples/pixel, data bytes with padded 0",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{1}),
				mustNewElement(tag.Columns, []int{1}),
				mustNewElement(tag.NumberOfFrames, []string{"3"}),
				mustNewElement(tag.BitsAllocated, []int{8}),
				mustNewElement(tag.SamplesPerPixel, []int{3}),
			}},
			dataBytes: []byte{1, 2, 3, 1, 2, 3, 1, 2, 3, 0}, // 10th byte to make total value length even
			expectedPixelData: &PixelDataInfo{
				IsEncapsulated: false,
				Frames: []frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 8,
							Rows:          1,
							Cols:          1,
							Data:          [][]int{{1, 2, 3}},
						},
					},
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 8,
							Rows:          1,
							Cols:          1,
							Data:          [][]int{{1, 2, 3}},
						},
					},
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 8,
							Rows:          1,
							Cols:          1,
							Data:          [][]int{{1, 2, 3}},
						},
					},
				},
			},
			expectedError: nil,
		},
		{
			Name: "1x1, 2 frames, 3 samples/pixel, bad pixel length",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{1}),
				mustNewElement(tag.Columns, []int{1}),
				mustNewElement(tag.NumberOfFrames, []string{"2"}),
				mustNewElement(tag.BitsAllocated, []int{8}),
				mustNewElement(tag.SamplesPerPixel, []int{3}),
			}},
			dataBytes:         []byte{1, 2, 3, 1, 2, 3},
			expectedPixelData: nil,
			pixelVLOverride:   7,
			expectedError:     ErrorExpectedEvenLength,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			dcmdata := bytes.Buffer{}
			var expectedBytes int

			if len(tc.data) == 0 {
				// writing byte-by-byte
				expectedBytes = len(tc.dataBytes)
				for _, item := range tc.dataBytes {
					if err := binary.Write(&dcmdata, binary.LittleEndian, item); err != nil {
						t.Errorf("TestReadNativeFrames: Unable to setup test buffer")
					}
				}
			} else {
				// writing 2 bytes (uint16) at a time
				expectedBytes = len(tc.data) * 2
				for _, item := range tc.data {
					if err := binary.Write(&dcmdata, binary.LittleEndian, item); err != nil {
						t.Errorf("TestReadNativeFrames: Unable to setup test buffer")
					}
				}
			}

			var vl uint32
			if tc.pixelVLOverride > 0 {
				vl = tc.pixelVLOverride
			} else {
				vl = uint32(dcmdata.Len())
			}

			r := &reader{
				rawReader: dicomio.NewReader(bufio.NewReader(&dcmdata), binary.LittleEndian, int64(dcmdata.Len())),
				opts:      tc.parseOptSet,
			}

			pixelData, bytesRead, err := r.readNativeFrames(&tc.existingData, nil, vl)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("TestReadNativeFrames(%v): did not get expected error. got: %v, want: %v", tc.data, err, tc.expectedError)
			}
			if err == nil && bytesRead != expectedBytes {
				t.Errorf("TestReadNativeFrames(%v): did not read expected number of bytes. got: %d, want: %d", tc.data, bytesRead, expectedBytes)
			}

			if diff := cmp.Diff(tc.expectedPixelData, pixelData, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("TestReadNativeFrames(%v): unexpected diff: %v", tc.data, diff)
			}
		})
	}
}

func TestReadPixelData_SkipPixelData(t *testing.T) {
	cases := []struct {
		name string
		vl   uint32
		data []byte
	}{
		{
			name: "NativePixelData",
			vl:   6,
			data: []byte{1, 2, 3, 4, 5, 6},
		},
		{
			name: "EncapsulatedPixelData",
			vl:   tag.VLUndefinedLength,
			data: makeEncapsulatedSequence(t),
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			opts := parseOptSet{skipPixelData: true}
			dcmdata := bytes.NewBuffer(tc.data)

			r := &reader{
				rawReader: dicomio.NewReader(bufio.NewReader(dcmdata), binary.LittleEndian, int64(dcmdata.Len())),
				opts:      opts,
			}
			val, err := r.readPixelData(tc.vl, &Dataset{}, nil)
			if err != nil {
				t.Errorf("unexpected error in readPixelData: %v", err)
			}
			pixelVal, ok := val.GetValue().(PixelDataInfo)
			if !ok {
				t.Errorf("Expected value to be of type PixelDataInfo")
			}
			if !pixelVal.IntentionallySkipped {
				t.Errorf("Expected PixelDataInfo to have IntentionallySkipped=true")
			}
		})
	}
}

func makeEncapsulatedSequence(t *testing.T) []byte {
	t.Helper()
	buf := &bytes.Buffer{}
	w := dicomio.NewWriter(buf, binary.LittleEndian, true)

	writePixelData(w, tag.PixelData, &pixelDataValue{PixelDataInfo{IsEncapsulated: true, Frames: []frame.Frame{
		{
			Encapsulated: true,
			EncapsulatedData: frame.EncapsulatedFrame{
				Data: []byte{1, 2, 3, 4},
			},
		},
	}}}, "", tag.VLUndefinedLength)

	return buf.Bytes()

}

func TestReadNativeFrames_OneBitAllocated(t *testing.T) {
	cases := []struct {
		Name              string
		existingData      Dataset
		data              []byte
		expectedPixelData *PixelDataInfo
		expectedError     error
		byteOrder         binary.ByteOrder
	}{
		{
			Name: "LittleEndian, 4x4, 1 frames, 1 samples/pixel",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{4}),
				mustNewElement(tag.Columns, []int{4}),
				mustNewElement(tag.NumberOfFrames, []string{"1"}),
				mustNewElement(tag.BitsAllocated, []int{1}),
				mustNewElement(tag.SamplesPerPixel, []int{1}),
			}},
			data: []byte{0b00010111, 0b10010111},
			expectedPixelData: &PixelDataInfo{
				IsEncapsulated: false,
				Frames: []frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 1,
							Rows:          4,
							Cols:          4,
							Data:          [][]int{{0}, {0}, {0}, {1}, {0}, {1}, {1}, {1}, {1}, {0}, {0}, {1}, {0}, {1}, {1}, {1}},
						},
					},
				},
			},
			expectedError: nil,
			byteOrder:     binary.LittleEndian,
		},
		{
			Name: "\"BigEndian\" (maybe), 4x4, 1 frames, 1 samples/pixel",
			existingData: Dataset{Elements: []*Element{
				mustNewElement(tag.Rows, []int{4}),
				mustNewElement(tag.Columns, []int{4}),
				mustNewElement(tag.NumberOfFrames, []string{"1"}),
				mustNewElement(tag.BitsAllocated, []int{1}),
				mustNewElement(tag.SamplesPerPixel, []int{1}),
			}},
			data: []byte{0b00010111, 0b10010111},
			expectedPixelData: &PixelDataInfo{
				IsEncapsulated: false,
				Frames: []frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 1,
							Rows:          4,
							Cols:          4,
							Data:          [][]int{{0}, {0}, {0}, {1}, {0}, {1}, {1}, {1}, {1}, {0}, {0}, {1}, {0}, {1}, {1}, {1}},
						},
					},
				},
			},
			expectedError: nil,
			// For some reason, this doesn't make a difference, even though we
			// don't change the order in which we read the bits in the
			// implementation. For some reason, binary.Write isn't changing the
			// order of the bits (or it is, and we somehow always read it back
			// in the correct order).
			byteOrder: binary.BigEndian,
		},
	}
	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			dcmdata := bytes.Buffer{}
			for _, item := range tc.data {
				if err := binary.Write(&dcmdata, tc.byteOrder, item); err != nil {
					t.Errorf("TestReadNativeFrames: Unable to setup test buffer")
				}
			}

			r := &reader{rawReader: dicomio.NewReader(bufio.NewReader(&dcmdata), tc.byteOrder, int64(dcmdata.Len()))}

			pixelData, _, err := r.readNativeFrames(&tc.existingData, nil, uint32(dcmdata.Len()))
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("TestReadNativeFrames(%v): did not get expected error. got: %v, want: %v", tc.data, err, tc.expectedError)
			}

			if diff := cmp.Diff(tc.expectedPixelData, pixelData); diff != "" {
				t.Errorf("TestReadNativeFrames(%v): unexpected diff: %v\ndata:%v", tc.data, diff, pixelData)
			}
		})
	}
}

func BenchmarkReadNativeFrames(b *testing.B) {
	cases := []struct {
		Name            string
		Rows            int
		Cols            int
		NumFrames       int
		SamplesPerPixel int
	}{
		{
			Name:            "10x10, 10 frames, 1 sample/pixel",
			Rows:            10,
			Cols:            10,
			NumFrames:       10,
			SamplesPerPixel: 1,
		},
		{
			Name:            "100x100, 10 frames, 1 sample/pixel",
			Rows:            100,
			Cols:            100,
			NumFrames:       10,
			SamplesPerPixel: 1,
		},
		{
			Name:            "512x512, 10 frames, 1 sample/pixel",
			Rows:            512,
			Cols:            512,
			NumFrames:       10,
			SamplesPerPixel: 1,
		},
		{
			Name:            "512x512, 10 frames, 5 sample/pixel",
			Rows:            512,
			Cols:            512,
			NumFrames:       10,
			SamplesPerPixel: 5,
		},
	}
	for _, c := range cases {
		b.Run(c.Name, func(b *testing.B) {
			dataset, rawReader := buildReadNativeFramesInput(c.Rows, c.Cols, c.NumFrames, c.SamplesPerPixel, b)
			r := &reader{rawReader: rawReader}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _, _ = r.readNativeFrames(dataset, nil, uint32(c.Rows*c.Cols*c.NumFrames))
			}
		})
	}
}

func buildReadNativeFramesInput(rows, cols, numFrames, samplesPerPixel int, b *testing.B) (*Dataset, dicomio.Reader) {
	b.Helper()
	dataset := Dataset{
		Elements: []*Element{
			mustNewElement(tag.Rows, []int{rows}),
			mustNewElement(tag.Columns, []int{cols}),
			mustNewElement(tag.NumberOfFrames, []string{strconv.Itoa(numFrames)}),
			mustNewElement(tag.BitsAllocated, []int{16}),
			mustNewElement(tag.SamplesPerPixel, []int{samplesPerPixel}),
		},
	}
	dcmdata := bytes.Buffer{}
	for fr := 0; fr < numFrames; fr++ {
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				for pxs := 0; pxs < samplesPerPixel; pxs++ {
					if err := binary.Write(&dcmdata, binary.LittleEndian, uint16(rand.Intn(100))); err != nil {
						b.Fatalf("TestReadNativeFrames: Unable to setup test buffer")
					}
				}
			}
		}
	}

	return &dataset, dicomio.NewReader(bufio.NewReader(&dcmdata), binary.LittleEndian, int64(dcmdata.Len()))
}

func buildTagData(t *testing.T, tg tag.Tag) []byte {
	t.Helper()
	data := bytes.Buffer{}
	if err := binary.Write(&data, binary.LittleEndian, tg.Group); err != nil {
		t.Errorf("buildTagData: Unable to setup test buffer: %v", err)
	}
	if err := binary.Write(&data, binary.LittleEndian, tg.Element); err != nil {
		t.Errorf("buildTagData: Unable to setup test buffer: %v", err)
	}

	return data.Bytes()
}
