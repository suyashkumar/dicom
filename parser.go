package dicom

import (
	"os"

	"github.com/suyashkumar/dicom/pkg/frame"
)

type ParseOptions struct {
	DropPixelData bool
}

// Parser represents an entity that can read and parse DICOMs
type Parser interface {
	Parse(options ParseOptions) (*Dataset, error)
	NextElement() (*Element, error)
}

type parser struct {
	dataset      Dataset
	opts         ParseOptions
	file         *os.File // may be populated if coming from file
	frameChannel chan *frame.Frame
}
