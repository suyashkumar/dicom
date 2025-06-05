// Package codec provides JPEG/J2K/JPEG-LS encoder and decoder
// Copyright (c) 2023 Segmed Inc.
package codec

// #cgo pkg-config: libopenjp2 charls
// #cgo darwin pkg-config: dcmtk
// #cgo linux LDFLAGS: -ldcmjpeg
import "C"
