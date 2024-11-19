package frame

import (
	"bytes"
	"image"
	"image/jpeg"
)

// EncapsulatedFrame represents an encapsulated image frame
type EncapsulatedFrame struct {
	// Data is a collection of bytes representing a JPEG encoded image frame
	Data []byte
}

// IsEncapsulated indicates if the frame is encapsulated or not.
func (e *EncapsulatedFrame) IsEncapsulated() bool { return true }

// GetEncapsulatedFrame returns an EncapsulatedFrame from this frame.
func (e *EncapsulatedFrame) GetEncapsulatedFrame() (*EncapsulatedFrame, error) {
	return e, nil
}

// GetNativeFrame returns ErrorFrameTypeNotPresent, because this struct does
// not hold a NativeFrame.
func (e *EncapsulatedFrame) GetNativeFrame() (INativeFrame, error) {
	return nil, ErrorFrameTypeNotPresent
}

// GetImage returns a Go image.Image from the underlying frame.
func (e *EncapsulatedFrame) GetImage() (image.Image, error) {
	// Decoding the Data to only re-encode it as a JPEG *without* modifications
	// is very inefficient. If all you want to do is write the JPEG to disk,
	// you should fetch the EncapsulatedFrame and grab the []byte Data from
	// there.
	return jpeg.Decode(bytes.NewReader(e.Data))
}

// Equals returns true if this frame equals the provided target frame, otherwise
// false.
func (e *EncapsulatedFrame) Equals(target *EncapsulatedFrame) bool {
	if !bytes.Equal(e.Data, target.Data) {
		return false
	}
	return true
}
