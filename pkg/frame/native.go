package frame

import (
	"image"
	"image/color"

	"golang.org/x/exp/constraints"
)

type INativeFrame interface {
	Rows() int
	Cols() int
	BitsPerSample() int
	GetPixel(x, y int) []int
	GetPixelAtIdx(idx int) []int
	RawDataSlice() any
	Equals(frame INativeFrame) bool
	CommonFrame
}

// NativeFrame represents a native image frame
type NativeFrame[I constraints.Integer] struct {
	// Data is a slice of pixels, where each pixel can have multiple values
	Data                  [][]I
	InternalRows          int
	InternalCols          int
	InternalBitsPerSample int
}

func NewNativeFrame[I constraints.Integer](bitsPerSample, rows, cols, pixelsPerFrame int) *NativeFrame[I] {
	return &NativeFrame[I]{
		InternalBitsPerSample: bitsPerSample,
		InternalRows:          rows,
		InternalCols:          cols,
		Data:                  make([][]I, pixelsPerFrame),
	}
}

func (n *NativeFrame[I]) Rows() int          { return n.InternalRows }
func (n *NativeFrame[I]) Cols() int          { return n.InternalCols }
func (n *NativeFrame[I]) BitsPerSample() int { return n.InternalBitsPerSample }
func (n *NativeFrame[I]) GetPixelAtIdx(idx int) []int {
	rawPixel := n.Data[idx]
	vals := make([]int, len(rawPixel))
	for i, val := range rawPixel {
		vals[i] = int(val)
	}
	return vals
}
func (n *NativeFrame[I]) GetPixel(x, y int) []int {
	dataIdx := x + (y * n.Cols())
	return n.GetPixelAtIdx(dataIdx)
}
func (n *NativeFrame[I]) RawDataSlice() any { return n.Data }

// IsEncapsulated indicates if the frame is encapsulated or not.
func (n *NativeFrame[I]) IsEncapsulated() bool { return false }

// GetNativeFrame returns a NativeFrame from this frame. If the underlying frame
// is not a NativeFrame, ErrorFrameTypeNotPresent will be returned.
func (n *NativeFrame[I]) GetNativeFrame() (INativeFrame, error) {
	return n, nil
}

// GetEncapsulatedFrame returns ErrorFrameTypeNotPresent, because this struct
// does not hold encapsulated frame Data.
func (n *NativeFrame[I]) GetEncapsulatedFrame() (*EncapsulatedFrame, error) {
	return nil, ErrorFrameTypeNotPresent
}

// GetImage returns an image.Image representation the frame, using default
// processing. This default processing is basic at the moment, and does not
// autoscale pixel values or use window width or level info.
func (n *NativeFrame[I]) GetImage() (image.Image, error) {
	i := image.NewGray16(image.Rect(0, 0, n.Cols(), n.Rows()))
	for j := 0; j < len(n.Data); j++ {
		i.SetGray16(j%n.Cols(), j/n.Cols(), color.Gray16{Y: uint16(n.Data[j][0])}) // for now, assume we're not overflowing uint16, assume gray image
	}
	return i, nil
}

// Equals returns true if this frame equals the provided target frame, otherwise
// false. This may be expensive.
func (n *NativeFrame[I]) Equals(target INativeFrame) bool {
	if target == nil || n == nil {
		return INativeFrame(n) == target
	}
	if n.Rows() != target.Rows() ||
		n.Cols() != target.Cols() ||
		n.BitsPerSample() != n.BitsPerSample() {
		return false
	}

	// If BitsPerSample are equal, we assume the target is of type
	// *NativeFrame[I]
	rawTarget, ok := target.(*NativeFrame[I])
	if !ok {

	}

	for pixIdx, pix := range n.Data {
		for valIdx, val := range pix {
			if val != rawTarget.Data[pixIdx][valIdx] {
				return false
			}
		}
	}
	return true
}
