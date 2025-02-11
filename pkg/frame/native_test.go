package frame_test

import (
	"errors"
	"fmt"
	"image"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/suyashkumar/dicom/pkg/frame"
)

// point represents a 2D point for testing.
type point struct {
	x int
	y int
}

func TestNativeFrame_GetImage(t *testing.T) {
	cases := []struct {
		Name        string
		NativeFrame frame.NativeFrame[int]
		SetPoints   []point
	}{
		{
			Name: "Square",
			NativeFrame: frame.NativeFrame[int]{
				InternalRows:            2,
				InternalCols:            2,
				InternalSamplesPerPixel: 1,
				RawData:                 []int{0, 0, 1, 0},
			},
			SetPoints: []point{{0, 1}},
		},
		{
			Name: "Rectangle",
			NativeFrame: frame.NativeFrame[int]{
				InternalRows:            3,
				InternalCols:            2,
				InternalSamplesPerPixel: 1,
				RawData:                 []int{0, 0, 0, 0, 1, 0},
			},
			SetPoints: []point{{0, 2}},
		},
		{
			Name: "Rectangle - multiple points",
			NativeFrame: frame.NativeFrame[int]{
				InternalRows:            5,
				InternalCols:            3,
				InternalSamplesPerPixel: 1,
				RawData:                 []int{0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0},
			},
			SetPoints: []point{{1, 1}, {2, 1}, {1, 3}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			img, err := tc.NativeFrame.GetImage()
			if err != nil {
				t.Fatalf("GetImage(%v) got unexpected error: %v", tc.NativeFrame, err)
			}
			imgGray, ok := img.(*image.Gray16)
			if !ok {
				t.Fatalf("GetImage(%v) did not return an image convertible to Gray16", tc.NativeFrame)
			}

			// Check that all pixels are zero except at the
			// (ExpectedSetPointX, ExpectedSetPointY) point.
			for x := 0; x < tc.NativeFrame.Cols(); x++ {
				for y := 0; y < tc.NativeFrame.Rows(); y++ {
					currValue := imgGray.Gray16At(x, y).Y
					if within(point{x, y}, tc.SetPoints) {
						if currValue != 1 {
							t.Errorf("GetImage(%v), unexpected value at set point. got: %v, want: %v", tc.NativeFrame, currValue, 1)
						}
						continue
					}

					if currValue != 0 {
						t.Errorf("GetImage(%v), unexpected value outside of set point. At (%d, %d) got: %v, want: %v", tc.NativeFrame, x, y, currValue, 0)
					}

				}
			}
		})

	}
}

func TestNativeFrame_GetImage_Errors(t *testing.T) {
	cases := []struct {
		name        string
		nativeFrame frame.NativeFrame[int]
		wantErr     error
	}{
		{
			name: "InternalSamplesPerPixel is not 1",
			nativeFrame: frame.NativeFrame[int]{
				InternalSamplesPerPixel: 2,
			},
			wantErr: frame.ErrUnsupportedSamplesPerPixel,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.nativeFrame.GetImage()
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("GetImage unexpected error. got: %v, want: %v", err, tc.wantErr)
			}
		})
	}
}

func TestNativeFrame_SimpleHelpers(t *testing.T) {
	// Tests various helper methods on NativeFrame[I]. GetImage is tested in a
	// separate top-level test case.
	f := frame.NativeFrame[uint8]{
		RawData:                 []uint8{1, 2, 3, 4, 5, 6},
		InternalSamplesPerPixel: 2,
		InternalRows:            1,
		InternalCols:            3,
		InternalBitsPerSample:   8,
	}

	if got := f.Rows(); got != 1 {
		t.Errorf("Rows() unexpected value, got: %v, want: %v", got, 1)
	}
	if got := f.Cols(); got != 3 {
		t.Errorf("Cols() unexpected value, got: %v, want: %v", got, 3)
	}
	if got := f.SamplesPerPixel(); got != 2 {
		t.Errorf("SamplesPerPixel() unexpected value, got: %v, want: %v", got, 2)
	}
	if got := f.BitsPerSample(); got != 8 {
		t.Errorf("BitsPerSample() unexpected value, got: %v, want: %v", got, 8)
	}
}

func TestNativeFrame_GetPixel(t *testing.T) {
	f := frame.NativeFrame[uint8]{
		RawData:                 []uint8{1, 2, 3, 4, 5, 6, 7, 8},
		InternalSamplesPerPixel: 2,
		InternalRows:            2,
		InternalCols:            2,
		InternalBitsPerSample:   8,
	}
	cases := []struct {
		x    int
		y    int
		want []int
	}{
		{
			x:    0,
			y:    0,
			want: []int{1, 2},
		},
		{
			x:    1,
			y:    0,
			want: []int{3, 4},
		},
		{
			x:    0,
			y:    1,
			want: []int{5, 6},
		},
		{
			x:    1,
			y:    1,
			want: []int{7, 8},
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("x: %d, y: %d", tc.x, tc.y), func(t *testing.T) {
			got, err := f.GetPixel(tc.x, tc.y)
			if err != nil {
				t.Errorf("GetPixel(%d, %d) got unexpected error: %v", tc.x, tc.y, err)
			}
			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("GetPixel(%d, %d) got unexpected slice. diff: %v", tc.x, tc.y, diff)
			}
		})
	}
}

func TestNativeFrame_RawDataSlice(t *testing.T) {
	f := frame.NativeFrame[uint8]{
		RawData:                 []uint8{1, 2, 3, 4, 5, 6, 7, 8},
		InternalSamplesPerPixel: 2,
		InternalRows:            2,
		InternalCols:            2,
		InternalBitsPerSample:   8,
	}

	sl := f.RawDataSlice()
	rd, ok := sl.([]uint8)
	if !ok {
		t.Errorf("RawDataSlice() should have returned a []uint8, but unable to type cast to []uint8")
	}
	if diff := cmp.Diff(rd, f.RawData); diff != "" {
		t.Errorf("RawDataSlice() got unexpected slice. diff: %v", diff)
	}
}

func TestNativeFrame_Equals(t *testing.T) {
	cases := []struct {
		name  string
		a     frame.NativeFrame[int]
		b     frame.NativeFrame[int]
		equal bool
	}{
		{
			name: "equal",
			a: frame.NativeFrame[int]{
				RawData:                 []int{1, 2, 3},
				InternalSamplesPerPixel: 2,
				InternalCols:            3,
				InternalRows:            4,
				InternalBitsPerSample:   64,
			},
			b: frame.NativeFrame[int]{
				RawData:                 []int{1, 2, 3},
				InternalSamplesPerPixel: 2,
				InternalCols:            3,
				InternalRows:            4,
				InternalBitsPerSample:   64,
			},
			equal: true,
		},
		{
			name: "mismatched data",
			a: frame.NativeFrame[int]{
				RawData: []int{1, 2, 3},
			},
			b: frame.NativeFrame[int]{
				RawData: []int{2, 2, 3},
			},
			equal: false,
		},
		{
			name: "mismatched BitsPerSample",
			a: frame.NativeFrame[int]{
				InternalBitsPerSample: 2,
			},
			b: frame.NativeFrame[int]{
				InternalBitsPerSample: 4,
			},
			equal: false,
		},
		{
			name: "mismatched SamplesPerPixel",
			a: frame.NativeFrame[int]{
				InternalSamplesPerPixel: 2,
			},
			b: frame.NativeFrame[int]{
				InternalSamplesPerPixel: 4,
			},
			equal: false,
		},
		{
			name: "mismatched Rows",
			a: frame.NativeFrame[int]{
				InternalRows: 2,
			},
			b: frame.NativeFrame[int]{
				InternalRows: 4,
			},
			equal: false,
		},
		{
			name: "mismatched Cols",
			a: frame.NativeFrame[int]{
				InternalCols: 2,
			},
			b: frame.NativeFrame[int]{
				InternalCols: 4,
			},
			equal: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.a.Equals(&tc.b)
			if got != tc.equal {
				t.Errorf("Equals(%+v, %+v) got unexpected value. got: %v, want: %v", tc.a, tc.b, got, tc.equal)
			}
		})
	}
}

// Test generated using Keploy
func TestNativeFrame_GetPixel_OutOfBounds(t *testing.T) {
    f := frame.NativeFrame[uint8]{
        RawData:                 []uint8{1, 2, 3, 4},
        InternalSamplesPerPixel: 2,
        InternalRows:            2,
        InternalCols:            2,
        InternalBitsPerSample:   8,
    }
    cases := []struct {
        x, y int
    }{
        {x: -1, y: 0},
        {x: 0, y: -1},
        {x: 2, y: 0},
        {x: 0, y: 2},
    }
    for _, tc := range cases {
        t.Run(fmt.Sprintf("x: %d, y: %d", tc.x, tc.y), func(t *testing.T) {
            _, err := f.GetPixel(tc.x, tc.y)
            if err == nil {
                t.Errorf("GetPixel(%d, %d) expected error but got none", tc.x, tc.y)
            }
        })
    }
}


// Test generated using Keploy
func TestNativeFrame_GetSample(t *testing.T) {
    f := frame.NativeFrame[uint8]{
        RawData:                 []uint8{1, 2, 3, 4, 5, 6},
        InternalSamplesPerPixel: 2,
        InternalRows:            1,
        InternalCols:            3,
        InternalBitsPerSample:   8,
    }
    cases := []struct {
        x, y, sampleIdx int
        want            int
    }{
        {x: 0, y: 0, sampleIdx: 0, want: 1},
        {x: 0, y: 0, sampleIdx: 1, want: 2},
        {x: 1, y: 0, sampleIdx: 0, want: 3},
        {x: 1, y: 0, sampleIdx: 1, want: 4},
        {x: 2, y: 0, sampleIdx: 0, want: 5},
        {x: 2, y: 0, sampleIdx: 1, want: 6},
    }
    for _, tc := range cases {
        t.Run(fmt.Sprintf("x: %d, y: %d, sampleIdx: %d", tc.x, tc.y, tc.sampleIdx), func(t *testing.T) {
            got := f.GetSample(tc.x, tc.y, tc.sampleIdx)
            if got != tc.want {
                t.Errorf("GetSample(%d, %d, %d) got: %d, want: %d", tc.x, tc.y, tc.sampleIdx, got, tc.want)
            }
        })
    }
}


// Test generated using Keploy
func TestNativeFrame_GetEncapsulatedFrame(t *testing.T) {
    f := frame.NativeFrame[int]{}
    ef, err := f.GetEncapsulatedFrame()
    if ef != nil {
        t.Errorf("GetEncapsulatedFrame() returned non-nil frame, expected nil")
    }
    if !errors.Is(err, frame.ErrorFrameTypeNotPresent) {
        t.Errorf("GetEncapsulatedFrame() unexpected error. got: %v, want: %v", err, frame.ErrorFrameTypeNotPresent)
    }
}


// within returns true if pt is in the []point
func within(pt point, set []point) bool {
	for _, item := range set {
		if pt.x == item.x && pt.y == item.y {
			return true
		}
	}
	return false
}
