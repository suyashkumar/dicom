// Package inplace contains code for handling UnprocessedValueData
package inplace

// ReadUnprocessedValueData read the value of Dicom image directly
// from the byte array of PixelData.UnprocessedValueData with given frame ID.
// This ease the memory usage of reading DICOM image.
func ReadUnprocessedValueData(info *PixelDataMetadata, unprocessedValueData []byte, frameIndex int) ([][]int, error) {
	pixelsPerFrame := info.Rows * info.Cols
	bytesAllocated := info.BitsAllocated / 8
	offset := frameIndex * pixelsPerFrame * info.SamplesPerPixel * bytesAllocated
	samplesPerPixel := info.SamplesPerPixel

	re := make([][]int, samplesPerPixel)
	for i := 0; i < samplesPerPixel; i++ {
		re[i] = make([]int, pixelsPerFrame)
		for j := 0; j < pixelsPerFrame; j++ {
			pointOffset := offset + j*info.SamplesPerPixel*bytesAllocated + i*bytesAllocated
			switch bytesAllocated {
			case 1:
				re[i][j] = int(unprocessedValueData[pointOffset])
			case 2:
				re[i][j] = int(info.Bo.Uint16(unprocessedValueData[pointOffset : pointOffset+bytesAllocated]))
			case 4:
				re[i][j] = int(info.Bo.Uint32(unprocessedValueData[pointOffset : pointOffset+bytesAllocated]))
			}
		}
	}
	return re, nil
}
