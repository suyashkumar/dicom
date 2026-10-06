package dicom

import (
	"github.com/suyashkumar/dicom/pkg/frame"
)

// EncapsulatedDecoder represents a function capable of decoding encapsulated frame bytes into an image.Image.
type EncapsulatedDecoder = frame.EncapsulatedDecoder

// ErrUnsupportedEncapsulatedCodec indicates that the encapsulated frame uses a transfer syntax
// or marker unsupported by the standard Go image/jpeg decoder.
var ErrUnsupportedEncapsulatedCodec = frame.ErrUnsupportedEncapsulatedCodec

// UnsupportedCodecError provides diagnostic details when an encapsulated frame cannot be decoded.
type UnsupportedCodecError = frame.UnsupportedCodecError

// RegisterEncapsulatedDecoder registers a custom decoder for the specified DICOM transfer syntax UID.
func RegisterEncapsulatedDecoder(transferSyntaxUID string, decoder EncapsulatedDecoder) {
	frame.RegisterEncapsulatedDecoder(transferSyntaxUID, decoder)
}

// GetEncapsulatedDecoder retrieves a registered decoder for the specified DICOM transfer syntax UID.
func GetEncapsulatedDecoder(transferSyntaxUID string) (EncapsulatedDecoder, bool) {
	return frame.GetEncapsulatedDecoder(transferSyntaxUID)
}

// UnregisterEncapsulatedDecoder removes any registered decoder for the specified DICOM transfer syntax UID.
func UnregisterEncapsulatedDecoder(transferSyntaxUID string) {
	frame.UnregisterEncapsulatedDecoder(transferSyntaxUID)
}

// ResetEncapsulatedDecoders removes all registered encapsulated decoders.
func ResetEncapsulatedDecoders() {
	frame.ResetEncapsulatedDecoders()
}
