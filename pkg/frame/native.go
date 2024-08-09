package frame

import (
	"errors"
	"fmt"
	"image"
	"image/color"

	"golang.org/x/exp/constraints"
)

var ErrUnsupportedSamplesPerPixel = errors.New("unsupported samples per pixel")

// INativeFrame is an interface representation of NativeFrame[I]'s capabilities,
// and offers a way to use a NativeFrame _without_ requiring propogation of
// type parameters. This allows for some more ergonomic signatures, though
// NativeFrame[I] can be used directly as well for those who prefer it.
type INativeFrame interface {
	// Rows returns the number of rows in this frame (which is the max y
	// dimension).
	Rows() int
	// Cols returns the number of columns in this frame (which is the max x
	// dimension).
	Cols() int
	// SamplesPerPixel returns the number of samples per pixel in this frame.
	SamplesPerPixel() int
	// BitsPerSample returns the bits per sample in this frame.
	BitsPerSample() int
	// GetPixel returns the samples (as a slice) for the pixel at (x, y).
	// The coordinate system of the image starts with (0, 0) in the upper left
	// corner of the image, with X increasing to the right, and Y increasing
	// down:
	//
	//  0 -------▶ X
	//  |
	//  |
	//  ▼
	//  Y
	GetPixel(x, y int) ([]int, error)
	// RawDataSlice will return the underlying data slice, which will be of type
	// []I. Based on BitsPerSample, you can anticipate what type of slice you'll
	// get, and type assert as needed:
	//  BitsPerSample    Slice
	//  8                []uint8
	//  16               []uint16
	//  32               []uint32
	RawDataSlice() any
	// Equals returns true if this INativeFrame exactly equals the provided
	// INativeFrame. This checks every pixel value, so may be expensive.
	// In the future we may compute a one time hash during construction to make
	// this less expensive in the future if called multiple time.
	Equals(frame INativeFrame) bool
	CommonFrame
}

// NativeFrame represents a native image frame
type NativeFrame[I constraints.Integer] struct {
	// RawData is a slice of pixel values. For each pixel, each sample for the
	// pixel is unrolled per pixel. For example, consider 2 pixels that have 3
	// samples per Pixel: [[1,2,3], [4,5,6]]. This would be unrolled like
	// [1,2,3,4,5,6]. The pixels themselves are arranged in row order, so the
	// first row of pixels would be unrolled in order, followed by the next row,
	// and so on in this flattened array.
	// A flattened slice is used instead of a nested 2D slice because there is
	// significant overhead to creating nested slices in Go discussed here:
	// https://github.com/suyashkumar/dicom/issues/161#issuecomment-2143627792.
	RawData                 []I
	InternalSamplesPerPixel int
	InternalRows            int
	InternalCols            int
	InternalBitsPerSample   int
}

func NewNativeFrame[I constraints.Integer](bitsPerSample, rows, cols, pixelsPerFrame, samplesPerPixel int) *NativeFrame[I] {
	return &NativeFrame[I]{
		InternalBitsPerSample:   bitsPerSample,
		InternalRows:            rows,
		InternalCols:            cols,
		RawData:                 make([]I, pixelsPerFrame*samplesPerPixel),
		InternalSamplesPerPixel: samplesPerPixel,
	}
}

// Rows returns the number of rows in this frame (which is the max y dimension).
func (n *NativeFrame[I]) Rows() int { return n.InternalRows }

// Cols returns the number of columns in this frame (which is the max x
// dimension).
func (n *NativeFrame[I]) Cols() int { return n.InternalCols }

// BitsPerSample returns the bits per sample.
func (n *NativeFrame[I]) BitsPerSample() int { return n.InternalBitsPerSample }

// SamplesPerPixel returns the samples per pixel.
func (n *NativeFrame[I]) SamplesPerPixel() int { return n.InternalSamplesPerPixel }

// GetPixel returns the samples (as a slice) for the pixel at (x, y).
// The coordinate system of the image starts with (0, 0) in the upper left
// corner of the image, with X increasing to the right, and Y increasing
// down:
//
//	0 -------▶ X
//	|
//	|
//	▼
//	Y
func (n *NativeFrame[I]) GetPixel(x, y int) ([]int, error) {
	if x < 0 || y < 0 || x >= n.InternalCols || y >= n.InternalRows {
		return nil, fmt.Errorf("provided zero-indexed coordinate (%v, %v) is out of bounds for this frame which has dimension %v x %v", x, y, n.InternalCols, n.InternalRows)
	}
	pixelIdx := (x * n.InternalSamplesPerPixel) + (y * (n.Cols() * n.InternalSamplesPerPixel))
	vals := make([]int, n.InternalSamplesPerPixel)
	for i := 0; i < n.InternalSamplesPerPixel; i++ {
		vals[i] = int(n.RawData[pixelIdx+i])
	}
	return vals, nil
}

// GetSample returns a specific sample inside a pixel at (x, y).
func (n *NativeFrame[I]) GetSample(x, y, sampleIdx int) int {
	dataSampleIdx := (x * n.InternalSamplesPerPixel) + (y * (n.Cols() * n.InternalSamplesPerPixel)) + sampleIdx
	return int(n.RawData[dataSampleIdx])
}

// RawDataSlice will return the underlying data slice, which will be of type
// []I. Based on BitsPerSample, you can anticipate what type of slice you'll
// get, and type assert as needed:
//
//	BitsPerSample    Slice
//	8                []uint8
//	16               []uint16
//	32               []uint32
func (n *NativeFrame[I]) RawDataSlice() any { return n.RawData }

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
	if n.InternalSamplesPerPixel != 1 {
		return nil, fmt.Errorf("GetImage(): unexpected InternalSamplesPerPixel got %v, expected 1 since only grayscale images are supported for now %w", n.InternalSamplesPerPixel, ErrUnsupportedSamplesPerPixel)
	}
	i := image.NewGray16(image.Rect(0, 0, n.Cols(), n.Rows()))
	for idx := 0; idx < len(n.RawData); idx++ {
		i.SetGray16(idx%n.Cols(), idx/n.Cols(), color.Gray16{Y: uint16(n.RawData[idx])}) // for now, assume we're not overflowing uint16, assume gray image, we can check BitsAllocated if we want to be conservative.
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
		return false // in reality, this is kind of an error, unless folks are implementing this interface themselves.
	}

	// TODO: check this using the interface only, which might be more expensive.
	for sampleIdx, sample := range n.RawData {
		if sample != rawTarget.RawData[sampleIdx] {
			return false
		}
	}

	return true
}
