package dicom

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math/rand"
	"strconv"
	"testing"

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
			r, err := dicomio.NewReader(bufio.NewReader(data), binary.LittleEndian, int64(data.Len()))
			if err != nil {
				t.Errorf("TestReadTag: unable to create new dicomio.Reader")
			}

			gotTag, err := readTag(r)
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

			r, err := dicomio.NewReader(bufio.NewReader(&data), binary.LittleEndian, int64(data.Len()))
			if err != nil {
				t.Errorf("TestReadFloat: unable to create new dicomio.Reader")
			}

			got, err := readFloat(r, tag.Tag{}, tc.VR, uint32(data.Len()))
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

			r, err := dicomio.NewReader(bufio.NewReader(&data), binary.LittleEndian, int64(data.Len()))
			if err != nil {
				t.Errorf("TestReadFloat: unable to create new dicomio.Reader")
			}

			got, err := readFloat(r, tag.Tag{}, tc.VR, uint32(data.Len()))
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
			name:        "even-number bytes",
			bytes:       []byte{0x1, 0x2, 0x3, 0x4},
			VR:          vrraw.OtherWord,
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

			r, err := dicomio.NewReader(bufio.NewReader(&data), binary.LittleEndian, int64(data.Len()))
			if err != nil {
				t.Errorf("TestReadOWBytes: unable to create new dicomio.Reader")
			}

			got, err := readBytes(r, tag.Tag{}, tc.VR, uint32(data.Len()))
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
		expectedPixelData *PixelDataInfo
		expectedError     error
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
			expectedError:     io.ErrUnexpectedEOF,
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
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			dcmdata := bytes.Buffer{}
			for _, item := range tc.data {
				if err := binary.Write(&dcmdata, binary.LittleEndian, item); err != nil {
					t.Errorf("TestReadNativeFrames: Unable to setup test buffer")
				}
			}

			r, err := dicomio.NewReader(bufio.NewReader(&dcmdata), binary.LittleEndian, int64(dcmdata.Len()))
			if err != nil {
				t.Errorf("TestReadFloat: unable to create new dicomio.Reader")
			}

			pixelData, _, err := readNativeFrames(r, &tc.existingData, nil)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("TestReadNativeFrames(%v): did not get expected error. got: %v, want: %v", tc.data, err, tc.expectedError)
			}

			if diff := cmp.Diff(tc.expectedPixelData, pixelData); diff != "" {
				t.Errorf("TestReadNativeFrames(%v): unexpected diff: %v", tc.data, diff)
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
			dataset, r := buildReadNativeFramesInput(c.Rows, c.Cols, c.NumFrames, c.SamplesPerPixel, b)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _, _ = readNativeFrames(r, dataset, nil)
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
	r, err := dicomio.NewReader(bufio.NewReader(&dcmdata), binary.LittleEndian, int64(dcmdata.Len()))
	if err != nil {
		b.Errorf("buildReadNativeFramesInput: unable to create new dicomio.Reader")
	}

	return &dataset, r
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
