// Copyright (c) 2023 Segmed Inc.
package codec

import "unsafe"

// CompressionMethod specifies the compression method for codec
type CompressionMethod int

const (
	// JPEG2000Compression represents the compression method for JPEG200
	JPEG2000Compression CompressionMethod = iota
	// JPEGLSCompression represents the compression method for JPEG Lossless
	JPEGLSCompression
	// JPEGCompression represents the compression method for JPEG
	JPEGCompression
)

// TransferSyntaxUIDToCompressionMethod is a map of transfer syntax UIDs to compression method
var TransferSyntaxUIDToCompressionMethod = map[string]CompressionMethod{
	"1.2.840.10008.1.2.4.90": JPEG2000Compression, // JPEG2000 Image Compression (Lossless Only)
	"1.2.840.10008.1.2.4.91": JPEG2000Compression, // JPEG2000 Image Compression

	"1.2.840.10008.1.2.4.80": JPEGLSCompression, // JPEG-LS Lossless
	"1.2.840.10008.1.2.4.81": JPEGLSCompression, // JPEG-LS Lossy (Near-Lossless)

	"1.2.840.10008.1.2.4.50": JPEGCompression, // JPEGProcess1TransferSyntax
	"1.2.840.10008.1.2.4.51": JPEGCompression, // JPEGProcess2_4TransferSyntax
	"1.2.840.10008.1.2.4.53": JPEGCompression, // JPEGProcess6_8TransferSyntax
	"1.2.840.10008.1.2.4.55": JPEGCompression, // JPEGProcess10_12TransferSyntax
	"1.2.840.10008.1.2.4.57": JPEGCompression, // JPEGProcess14TransferSyntax
	"1.2.840.10008.1.2.4.70": JPEGCompression, // JPEGProcess14SV1TransferSyntax
}

func makeSliceFromByte(ptr *byte, length uint32) []byte {
	slice := unsafe.Slice(ptr, length)
	var ret = make([]byte, length)
	for i, elem := range slice {
		ret[i] = byte(elem)
	}
	return ret
}
