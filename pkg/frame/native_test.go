package frame_test

import (
	"image"
	"testing"

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
				InternalRows: 2,
				InternalCols: 2,
				Data:         [][]int{{0}, {0}, {1}, {0}},
			},
			SetPoints: []point{{0, 1}},
		},
		{
			Name: "Rectangle",
			NativeFrame: frame.NativeFrame[int]{
				InternalRows: 3,
				InternalCols: 2,
				Data:         [][]int{{0}, {0}, {0}, {0}, {1}, {0}},
			},
			SetPoints: []point{{0, 2}},
		},
		{
			Name: "Rectangle - multiple points",
			NativeFrame: frame.NativeFrame[int]{
				InternalRows: 5,
				InternalCols: 3,
				Data:         [][]int{{0}, {0}, {0}, {0}, {1}, {1}, {0}, {0}, {0}, {0}, {1}, {0}, {0}, {0}, {0}},
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

// within returns true if pt is in the []point
func within(pt point, set []point) bool {
	for _, item := range set {
		if pt.x == item.x && pt.y == item.y {
			return true
		}
	}
	return false
}
