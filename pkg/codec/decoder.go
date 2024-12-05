// Copyright (c) 2023 Segmed Inc
package codec

/*
#include <stdlib.h>
*/
import "C"
import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"image/color"
	"strings"
	"unsafe"

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/frame"
)

// GetNativePixelData reads a JPEG-compressed DICOM image, decompresses the JPEG data (i. e. conversion to a
// native DICOM transfer syntax) and returns the converted image
func GetNativePixelData(e frame.EncapsulatedFrame) ([]byte, error) {
	var (
		dataOut *byte
		outLen  int64
		comp    = TransferSyntaxUIDToCompressionMethod[e.TransferSyntax]
	)
	var (
		dataIn        = (*byte)(&e.Data[0])
		length        = int64(len(e.Data))
		bitsAllocated = e.BitsAllocated
	)
	switch comp {
	case JPEG2000Compression:
		dataOut = J2K_Decode(dataIn, length, e.Cols, e.Rows, &outLen)
	case JPEGLSCompression:
		dataOut = JPEGLS_Decode(dataIn, length, bitsAllocated, e.Cols, e.Rows, e.SamplesPerPixel, e.PlanarConfiguration, &outLen)
	case JPEGCompression:
		ybr := 0
		if strings.HasPrefix(e.PhotometricInterpretation, "YBR_") {
			ybr = 1
		}
		dataOut = JPEG_Decode(dataIn, length, bitsAllocated, e.Cols, e.Rows, e.SamplesPerPixel, e.PlanarConfiguration, ybr, &outLen)
	default:
		return nil, errors.New("unsupported transfer syntax")
	}
	if outLen == 0 {
		return nil, errors.New("failed to decode DICOM")
	}
	defer C.free(unsafe.Pointer(dataOut))
	return makeSliceFromByte(dataOut, uint32(outLen)), nil
}

// GetStdImage returns the converted image in Go standard format
func GetStdImage(e frame.EncapsulatedFrame) (image.Image, error) {
	var (
		upLeft      = image.Point{0, 0}
		lowRight    = image.Point{e.Cols, e.Rows}
		img8        = image.NewRGBA(image.Rectangle{upLeft, lowRight})
		img16       = image.NewRGBA64(image.Rectangle{upLeft, lowRight})
		cols        = e.Cols
		a           = 0xffff
		pixelValues = make([]int, e.SamplesPerPixel)
	)
	decodedBytes, err := GetNativePixelData(e)
	if err != nil {
		return nil, fmt.Errorf("decode pixeldata: %w", err)
	}
	buf := bytes.NewBuffer(decodedBytes)
	r := dicomio.NewReader(bufio.NewReader(buf), binary.LittleEndian, int64(buf.Len()))
	oneByte := e.BitsAllocated == 8
	for i := 0; i < e.Cols*e.Rows; i++ {
		var clr color.Color
		err := readPixelValues(e, *r, pixelValues)
		if err != nil {
			return nil, fmt.Errorf("read pixel values: %w", err)
		}
		switch e.SamplesPerPixel {
		case 1:
			r := pixelValues[0]
			if oneByte {
				clr = color.Gray{Y: uint8(r)}
			} else {
				clr = color.Gray16{Y: uint16(r)}
			}
		case 3:
			r, g, b := pixelValues[0], pixelValues[1], pixelValues[2]
			if oneByte {
				clr = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			} else {
				clr = color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}
			}
		case 4:
			r, g, b, a := pixelValues[0], pixelValues[1], pixelValues[2], pixelValues[3]
			if oneByte {
				clr = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			} else {
				clr = color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)}
			}
		}
		if oneByte {
			img8.Set(i%cols, i/cols, clr)
		} else {
			img16.Set(i%cols, i/cols, clr)
		}
	}
	if oneByte {
		return img8, nil
	}
	return img16, nil
}

func readPixelValues(e frame.EncapsulatedFrame, r dicomio.Reader, pixelValues []int) error {
	for sample := 0; sample < e.SamplesPerPixel; sample++ {
		if r.IsLimitExhausted() {
			break
		}
		switch e.BitsAllocated {
		case 8:
			v, err := r.ReadUInt8()
			if err != nil {
				return fmt.Errorf("read uint8 failed: %w", err)
			}
			pixelValues[sample] = int(v)
		case 16:
			v, err := r.ReadUInt16()
			if err != nil {
				return fmt.Errorf("read uint16 failed: %w", err)
			}
			pixelValues[sample] = int(v)
		case 32:
			v, err := r.ReadUInt32()
			if err != nil {
				return fmt.Errorf("read uint32 failed: %w", err)
			}
			pixelValues[sample] = int(v)
		default:
			return fmt.Errorf("bitsAllocated not supported: %d", e.BitsAllocated)
		}
	}
	return nil
}

// Decode reads a JPEG image from EncapsulatedFrame and returns it as an NativeFrame.
func Decode(e frame.EncapsulatedFrame) (frame.NativeFrame, error) {
	var nativeData = make([][]int, e.Cols*e.Rows)
	decodedBytes, err := GetNativePixelData(e)
	if err != nil {
		return frame.NativeFrame{}, fmt.Errorf("decode pixeldata: %w", err)
	}
	buf := bytes.NewBuffer(decodedBytes)
	r := dicomio.NewReader(bufio.NewReader(buf), binary.LittleEndian, int64(buf.Len()))
	var (
		pixelValues = make([]int, e.SamplesPerPixel)
	)
	for i := 0; i < e.Cols*e.Rows; i++ {
		err := readPixelValues(e, *r, pixelValues)
		if err != nil {
			return frame.NativeFrame{}, fmt.Errorf("read pixel values: %w", err)
		}
		switch e.SamplesPerPixel {
		case 1:
			r := pixelValues[0]
			nativeData[i] = []int{int(r)}
		case 3:
			r, g, b := pixelValues[0], pixelValues[1], pixelValues[2]
			nativeData[i] = []int{int(r), int(g), int(b)}
		case 4:
			r, g, b, a := pixelValues[0], pixelValues[1], pixelValues[2], pixelValues[3]
			nativeData[i] = []int{int(r), int(g), int(b), int(a)}
		}
	}
	nf := frame.NativeFrame{
		BitsPerSample: e.BitsAllocated,
		Cols:          e.Cols,
		Rows:          e.Rows,
		Data:          nativeData,
	}
	return nf, nil
}

func init() {
	for transferSyntax := range TransferSyntaxUIDToCompressionMethod {
		frame.RegisterFormat(transferSyntax, Decode)
	}
}
