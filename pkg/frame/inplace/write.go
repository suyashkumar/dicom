package inplace

import (
	"bytes"
	"encoding/binary"
)

// WriteUnprocessedValueData write the value of Dicom image directly to
// byte array of PixelData.UnprocessedValueData
func WriteUnprocessedValueData(info PixelDataMetadata, unprocessedValueData []byte,
	rowIndex, colIndex int, frameIndex int, sampleIndex int, newValue int) error {
	pixelsPerFrame := info.Rows * info.Cols
	bytesAllocated := info.BitsAllocated / 8
	buf := &bytes.Buffer{}
	buf.Grow(bytesAllocated)
	switch bytesAllocated {
	case 1:
		if err := binary.Write(buf, binary.LittleEndian, uint8(newValue)); err != nil {
			return err
		}
	case 2:
		if err := binary.Write(buf, binary.LittleEndian, uint16(newValue)); err != nil {
			return err
		}
	case 4:
		if err := binary.Write(buf, binary.LittleEndian, uint32(newValue)); err != nil {
			return err
		}
	}
	paintedPixel := buf.Bytes()
	offset := bytesAllocated * info.SamplesPerPixel * frameIndex * pixelsPerFrame
	offset += info.SamplesPerPixel * bytesAllocated * (rowIndex*info.Cols + colIndex)

	offset += bytesAllocated * sampleIndex
	for k := 0; k < bytesAllocated; k++ {
		unprocessedValueData[offset+k] = paintedPixel[k]
	}
	return nil
}
