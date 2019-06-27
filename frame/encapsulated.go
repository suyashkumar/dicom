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

func (e *EncapsulatedFrame) IsEncapsulated() bool { return true }

func (e *EncapsulatedFrame) GetEncapsulatedFrame() (*EncapsulatedFrame, error) {
	return e, nil
}

func (e *EncapsulatedFrame) GetNativeFrame() (*NativeFrame, error) {
	return nil, ErrorFrameTypeNotPresent
}

func (e *EncapsulatedFrame) GetImage() (image.Image, error) {
	// Decoding the data to only rencode it as a JPEG *without* modifications
	// is very inefficient. If all you want to do is write the JPEG to disk,
	// you should fetch the EncapsulatedFrame and grab the []byte Data from there.
	return jpeg.Decode(bytes.NewReader(e.Data))
}
