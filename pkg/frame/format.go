package frame

import (
	"errors"
	"sync"
	"sync/atomic"
)

// ErrFormat indicates that decoding encountered an unknown format.
var ErrFormat = errors.New("image: unknown format")

// sniff determines the format of frame's TransferSyntax.
func sniff(frame EncapsulatedFrame) format {
	formats, _ := atomicFormats.Load().([]format)
	for _, f := range formats {
		if f.transferSyntax == frame.TransferSyntax {
			return f
		}
	}
	return format{}
}

// A format holds an image format's name and how to decode it.
type format struct {
	transferSyntax string
	decode         func(EncapsulatedFrame) (NativeFrame, error)
}

// Formats is the list of registered formats.
var (
	formatsMu     sync.Mutex
	atomicFormats atomic.Value
)

// RegisterFormat registers an image format for use by Decode.
// Decode is the function that decodes the encoded image.
func RegisterFormat(transferSyntax string, decode func(EncapsulatedFrame) (NativeFrame, error)) {
	formatsMu.Lock()
	formats, _ := atomicFormats.Load().([]format)
	atomicFormats.Store(append(formats, format{transferSyntax, decode}))
	formatsMu.Unlock()
}
