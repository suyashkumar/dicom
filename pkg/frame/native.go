package frame

import (
	"image"
	"image/color"
)

// NativeFrame represents a native image frame
type NativeFrame struct {
	// Data is a slice of pixels, where each pixel can have multiple values
	Data          [][]int
	Rows          int
	Cols          int
	BitsPerSample int
}

// IsEncapsulated indicates if the frame is encapsulated or not.
func (n *NativeFrame) IsEncapsulated() bool { return false }

// GetNativeFrame returns a NativeFrame from this frame. If the underlying frame
// is not a NativeFrame, ErrorFrameTypeNotPresent will be returned.
func (n *NativeFrame) GetNativeFrame() (*NativeFrame, error) {
	return n, nil
}

// GetEncapsulatedFrame returns ErrorFrameTypeNotPresent, because this struct
// does not hold encapsulated frame data.
func (n *NativeFrame) GetEncapsulatedFrame() (*EncapsulatedFrame, error) {
	return nil, ErrorFrameTypeNotPresent
}

// GetImage returns an image.Image representation the frame, using default
// processing.
func (n *NativeFrame) GetImage() (image.Image, error) {
	// Determine the min and max pixel values
	minVal, maxVal := findMinMax(n)

	// Calculate the window width and level
	windowWidth := maxVal - minVal
	windowLevel := (maxVal + minVal) / 2

	i := image.NewGray(image.Rect(0, 0, n.Cols, n.Rows))
	for j := 0; j < len(n.Data); j++ {
		// Apply window width and level adjustment to pixel value
		adjustedPixel := applyWindowLevel(n.Data[j][0], windowWidth, windowLevel)

		// Map the adjusted pixel value to grayscale
		grayValue := uint8(adjustedPixel)
		i.SetGray(j%n.Cols, j/n.Cols, color.Gray{Y: grayValue})
	}
	return i, nil
}

// Equals returns true if this frame equals the provided target frame, otherwise
// false.
func (n *NativeFrame) Equals(target *NativeFrame) bool {
	if target == nil || n == nil {
		return n == target
	}
	if n.Rows != target.Rows ||
		n.Cols != target.Cols ||
		n.BitsPerSample != target.BitsPerSample {
		return false
	}
	for pixIdx, pix := range n.Data {
		for valIdx, val := range pix {
			if val != target.Data[pixIdx][valIdx] {
				return false
			}
		}
	}
	return true
}

// Finds the minimum and maximum pixel values
func findMinMax(frame *NativeFrame) (min, max int) {
	min = frame.Data[0][0]
	max = frame.Data[0][0]
	for _, dataRow := range frame.Data {
		val := dataRow[0]
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max
}

// applying the pixel values to the image
func applyWindowLevel(pixelValue, windowWidth, windowLevel int) int {
	// Calculate the lower and upper bounds of the window
	lowerBound := windowLevel - windowWidth/2
	upperBound := windowLevel + windowWidth/2

	// Clip pixel value to the window bounds
	if pixelValue < lowerBound {
		pixelValue = lowerBound
	}
	if pixelValue > upperBound {
		pixelValue = upperBound
	}

	// Scale the pixel value to 8-bit range (0-255)
	scaledValue := 255 * (pixelValue - lowerBound) / windowWidth

	return scaledValue
}
