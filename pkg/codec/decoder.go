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
	buf := bytes.NewBuffer(e.Data)
	r := dicomio.NewReader(bufio.NewReader(buf), binary.LittleEndian, int64(len(e.Data)))
	var (
		upLeft          = image.Point{0, 0}
		lowRight        = image.Point{e.Cols, e.Rows}
		img8            = image.NewRGBA(image.Rectangle{upLeft, lowRight})
		img16           = image.NewRGBA64(image.Rectangle{upLeft, lowRight})
		samplesPerPixel int
		values8         [3]uint8 // red, green, blue
		values16        [3]uint16
		cols            = e.Cols
		a               = 0xffff
		i               = 0
	)
	if e.BitsAllocated == 8 {
		for !r.IsLimitExhausted() {
			i++
			v, err := r.ReadUInt8()
			if err != nil {
				return nil, fmt.Errorf("read uint8 failed: %w", err)
			}
			values8[samplesPerPixel%3] = v
			samplesPerPixel++
			if samplesPerPixel < e.SamplesPerPixel {
				continue
			}
			var clr color.Color
			switch samplesPerPixel {
			case 1:
				clr = color.Gray{Y: values8[0]}
			case 3:
				r, g, b := values8[0], values8[1], values8[2]
				clr = color.RGBA{r, g, b, uint8(a)}
			}
			img8.Set(i%cols, i/cols, clr)
			samplesPerPixel = 0
		}
		return img8, nil
	}
	for !r.IsLimitExhausted() {
		i++
		v, err := r.ReadUInt16()
		if err != nil {
			return nil, fmt.Errorf("read uint16 failed: %w", err)
		}
		values16[samplesPerPixel%3] = v
		samplesPerPixel++
		if samplesPerPixel < e.SamplesPerPixel {
			continue
		}
		var clr color.Color
		switch samplesPerPixel {
		case 1:
			clr = color.Gray16{Y: values16[0]}
		case 3:
			r, g, b := values16[0], values16[1], values16[2]
			clr = color.RGBA64{r, g, b, uint16(a)}
		}
		img16.Set(i%cols, i/cols, clr)
		samplesPerPixel = 0
	}
	return img16, nil
}

// Decode reads a JPEG image from EncapsulatedFrame and returns it as an NativeFrame.
func Decode(e frame.EncapsulatedFrame) (frame.NativeFrame, error) {
	img, err := GetStdImage(e)
	if err != nil {
		return frame.NativeFrame{}, fmt.Errorf("get Go's std image: %w", err)
	}
	var nativeData [][]int
	for i := 0; i < e.Cols*e.Rows; i++ {
		clr := img.At(i%e.Cols, i/e.Cols)
		r, g, b, a := clr.RGBA()
		switch e.SamplesPerPixel {
		case 1:
			nativeData = append(nativeData, []int{int(r)})
		case 3:
			nativeData = append(nativeData, []int{int(r), int(g), int(b)})
		case 4:
			nativeData = append(nativeData, []int{int(r), int(g), int(b), int(a)})
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
