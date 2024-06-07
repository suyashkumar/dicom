package inplace

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

// GetIntFromValue is similar to dicom.MustGetInts but without panic
func GetIntFromValue(v dicom.Value) (int, error) {
	if v.ValueType() != dicom.Ints {
		return 0, errors.New("value is not ints")
	}
	arr, ok := v.GetValue().([]int)
	if !ok {
		return 0, errors.New("actual value is not ints")
	}
	if len(arr) == 0 {
		return 0, errors.New("empty ints value")
	}
	return arr[0], nil
}

// GetStringFromValue is similar to dicom.MustGetStrings but without panic
func GetStringFromValue(v dicom.Value) (string, error) {
	if v.ValueType() != dicom.Strings {
		return "", errors.New("value is not ints")
	}
	arr, ok := v.GetValue().([]string)
	if !ok {
		return "", errors.New("actual value is not ints")
	}
	if len(arr) == 0 {
		return "", errors.New("empty ints value")
	}
	return arr[0], nil
}

// PixelDataMetadata is the metadata for tag.PixelData
type PixelDataMetadata struct {
	Rows                int
	Cols                int
	Frames              int
	SamplesPerPixel     int
	BitsAllocated       int
	PlanarConfiguration int
	Bo                  binary.ByteOrder
}

// GetPixelDataMetadata returns the pixel data metadata.
func GetPixelDataMetadata(ds *dicom.Dataset) (*PixelDataMetadata, error) {
	re := &PixelDataMetadata{}
	rows, err := ds.FindElementByTag(tag.Rows)
	if err != nil {
		return nil, fmt.Errorf("get Rows element: %w", err)
	}
	if re.Rows, err = GetIntFromValue(rows.Value); err != nil {
		return nil, fmt.Errorf("convert Rows element to int: %w", err)
	}

	cols, err := ds.FindElementByTag(tag.Columns)
	if err != nil {
		return nil, fmt.Errorf("get Columns element: %w", err)
	}
	if re.Cols, err = GetIntFromValue(cols.Value); err != nil {
		return nil, fmt.Errorf("convert Columns element to int: %w", err)
	}

	numberOfFrames, err := ds.FindElementByTag(tag.NumberOfFrames)
	if err != nil {
		re.Frames = 1
	} else {
		var framesStr string
		framesStr, err = GetStringFromValue(numberOfFrames.Value)
		if err != nil {
			return nil, fmt.Errorf("convert NumberOfFrames element to str: %w", err)
		}
		if re.Frames, err = strconv.Atoi(framesStr); err != nil {
			return nil, fmt.Errorf("convert NumberOfFrames to int: %w", err)
		}
	}

	samplesPerPixel, err := ds.FindElementByTag(tag.SamplesPerPixel)
	if err != nil {
		return nil, fmt.Errorf("get SamplesPerPixel element: %w", err)
	}
	if re.SamplesPerPixel, err = GetIntFromValue(samplesPerPixel.Value); err != nil {
		return nil, fmt.Errorf("convert SamplesPerPixel element to int: %w", err)
	}
	bitsAllocated, err := ds.FindElementByTag(tag.BitsAllocated)
	if err != nil {
		return nil, fmt.Errorf("get BitsAllocated element: %w", err)
	}
	if re.BitsAllocated, err = GetIntFromValue(bitsAllocated.Value); err != nil {
		return nil, fmt.Errorf("convert BitsAllocated element to int: %w", err)
	}
	re.Bo, _, err = ds.TransferSyntax()
	if err != nil {
		return nil, fmt.Errorf("get byteOrder: %w", err)
	}

	planarConfElement, err := ds.FindElementByTag(tag.PlanarConfiguration)
	if err != nil {
		re.PlanarConfiguration = 0
	} else {
		if re.PlanarConfiguration, err = GetIntFromValue(planarConfElement.Value); err != nil {
			return nil, fmt.Errorf("convert Rows element to int: %w", err)
		}
	}

	return re, nil
}

// IsSafeForUnprocessedValueDataHandling check if we can support in-place read-write
// from Pixeldata.UnprocessedValueData
// This avoids the case that we can not handle it, yet.
func IsSafeForUnprocessedValueDataHandling(info *PixelDataMetadata, unprocessedValueData []byte) error {
	// https://dicom.innolitics.com/ciods/enhanced-mr-image/enhanced-mr-image/00280006
	if info.PlanarConfiguration == 1 {
		return fmt.Errorf("unsupported PlanarConfiguration: %d", info.PlanarConfiguration)
	}
	// TODO: support for BitsAllocated == 1
	switch info.BitsAllocated {
	case 8, 16, 32:
	default: // bitsAllocated = 1 and other cases
		return fmt.Errorf("unsupported bit allocated: %d", info.BitsAllocated)
	}
	pixelsPerFrame := info.Rows * info.Cols
	bytesAllocated := info.BitsAllocated / 8
	expectedBytes := bytesAllocated * info.SamplesPerPixel * info.Frames * pixelsPerFrame
	// odd number of bytes.
	if expectedBytes%2 == 1 {
		expectedBytes += 1
	}
	if len(unprocessedValueData) != expectedBytes {
		return errors.New("mismatch data size")
	}
	return nil
}
