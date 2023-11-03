// Package inplace contains code for handling UnprocessedValueData
package inplace

import (
	"encoding/binary"
)

// ReadUnprocessedValueData read the value of Dicom image directly to
// byte array of PixelData.UnprocessedValueData
func ReadUnprocessedValueData(info PixelDataMetadata, unprocessedValueData []byte,
	bo binary.ByteOrder, frameIndex int) ([][]int, error) {

	pixelsPerFrame := info.Rows * info.Cols
	bytesAllocated := info.BitsAllocated / 8
	offset := bytesAllocated * info.SamplesPerPixel * frameIndex * pixelsPerFrame
	samplesPerPixel := info.SamplesPerPixel

	re := make([][]int, pixelsPerFrame)
	for i := 0; i < pixelsPerFrame; i++ {
		re[i] = make([]int, bytesAllocated)
		for j := 0; j < samplesPerPixel; j++ {
			pointOffset := offset + info.SamplesPerPixel*bytesAllocated*i + j*samplesPerPixel
			switch bytesAllocated {
			case 1:
				re[i][j] = int(unprocessedValueData[pointOffset])
			case 2:
				re[i][j] = int(bo.Uint16(unprocessedValueData[pointOffset : pointOffset+bytesAllocated]))
			case 4:
				re[i][j] = int(bo.Uint32(unprocessedValueData[pointOffset : pointOffset+bytesAllocated]))
			}
		}
	}

	return nil, nil
}
