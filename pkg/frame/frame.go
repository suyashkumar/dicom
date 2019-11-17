package frame

import (
	"errors"
	"image"
)

// ErrorFrameTypeNotPresent is returned when the user asked to Get an underlying GetNativeFrame or
// GetEncapsulatedFrame that is not contained in that particular CommonFrame.
var ErrorFrameTypeNotPresent = errors.New("The frame type you requested is not present in this CommonFrame")

// CommonFrame represents a harmonized DICOM Frame with a consistent interface (harmonized across
// Native and Encapsulated frames), however users still have the ability to fetch underlying Native
// or Encapsulated frame constructs.
type CommonFrame interface {
	// GetImage gets this frame as an image.Image. Beware that the underlying frame may perform
	// some default rendering and conversions. Operate on the raw NativeFrame or EncapsulatedFrame
	// if you need to do some custom rendering work or want the data from the dicom.
	GetImage() (image.Image, error)
	// Encapsulated indicates if the underlying Frame is an EncapsulatedFrame
	IsEncapsulated() bool
	// GetNativeFrame attempts to get the underlying NativeFrame (or returns an error)
	GetNativeFrame() (*NativeFrame, error)
	// GetEncapsulatedFrame attempts to get the underlying EncapsulatedFrame (or returns an error)
	GetEncapsulatedFrame() (*EncapsulatedFrame, error)
}

// Frame wraps a single encapsulated or native image frame
// TODO: deprecate this old intermediate representation in favor of CommonFrame once happy and solid with API
type Frame struct {
	Encapsulated     bool
	EncapsulatedData EncapsulatedFrame
	NativeData       NativeFrame
}

func (f *Frame) IsEncapsulated() bool { return f.Encapsulated }

func (f *Frame) GetNativeFrame() (*NativeFrame, error) {
	if f.Encapsulated {
		return f.EncapsulatedData.GetNativeFrame()
	} else {
		return f.NativeData.GetNativeFrame()
	}
}

func (f *Frame) GetEncapsulatedFrame() (*EncapsulatedFrame, error) {
	if f.Encapsulated {
		return f.EncapsulatedData.GetEncapsulatedFrame()
	} else {
		return f.NativeData.GetEncapsulatedFrame()
	}

}

func (f *Frame) GetImage() (image.Image, error) {
	if f.Encapsulated {
		return f.EncapsulatedData.GetImage()
	} else {
		return f.NativeData.GetImage()
	}
}
