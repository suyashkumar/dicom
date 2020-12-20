package dicom

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/frame"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	mock_dicomio "github.com/suyashkumar/dicom/mocks/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func TestReadTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	m := mock_dicomio.NewMockReader(ctrl)

	first := m.EXPECT().ReadUInt16().Return(uint16(0x7FE0), nil)
	m.EXPECT().ReadUInt16().Return(uint16(0x0010), nil).After(first)

	retTag, err := readTag(m)

	assert.NoError(err)
	fmt.Println(err)
	assert.Equal(tag.PixelData, *retTag)

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
			VR:          "FD",
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
			VR:          "FL",
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
			expectedError:     ErrorIncompleteRead,
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
