package inplace

import (
	"bytes"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/suyashkumar/dicom/pkg/uid"
)

func TestWriteUnprocessedValueData(t *testing.T) {
	originDataset := dicom.Dataset{
		Elements: []*dicom.Element{
			mustNewElement(t, tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
			mustNewElement(t, tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
			mustNewElement(t, tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
			mustNewElement(t, tag.Rows, []int{2}),
			mustNewElement(t, tag.Columns, []int{2}),
			mustNewElement(t, tag.NumberOfFrames, []string{"2"}),
			mustNewElement(t, tag.BitsAllocated, []int{16}),
			mustNewElement(t, tag.SamplesPerPixel, []int{2}),

			mustNewElement(t, tag.PixelData, dicom.PixelDataInfo{
				IsEncapsulated: false,
				Frames: []*frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 16,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 2}},
						},
					},
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 16,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 5}},
						},
					},
				},
			}),
		},
	}
	var buffer bytes.Buffer
	if err := dicom.Write(io.Writer(&buffer), originDataset); err != nil {
		t.Errorf("write from dataset to DICOM bytes: %v", err)
	}
	rawData := buffer.Bytes()

	for _, tc := range []struct {
		sampleIndex, rowIndex, colIndex, frameIndex int
		expectedPixelInfo                           dicom.PixelDataInfo
	}{
		{
			sampleIndex: 0, colIndex: 1, rowIndex: 0, frameIndex: 0,
			expectedPixelInfo: dicom.PixelDataInfo{
				Frames: []*frame.Frame{
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {100, 2}, {1, 2}, {3, 2}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 5}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
				},
			},
		},
		{
			sampleIndex: 1, colIndex: 1, rowIndex: 0, frameIndex: 0,
			expectedPixelInfo: dicom.PixelDataInfo{
				Frames: []*frame.Frame{
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 100}, {1, 2}, {3, 2}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 5}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
				},
			},
		},
		{
			sampleIndex: 1, colIndex: 0, rowIndex: 1, frameIndex: 0,
			expectedPixelInfo: dicom.PixelDataInfo{
				Frames: []*frame.Frame{
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 100}, {3, 2}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 5}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
				},
			},
		},
		{
			sampleIndex: 1, colIndex: 0, rowIndex: 1, frameIndex: 1,
			expectedPixelInfo: dicom.PixelDataInfo{
				Frames: []*frame.Frame{
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 2}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 100}, {3, 5}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
				},
			},
		},
	} {
		var err error
		dataset, metadata, pixelDataInfo := setupUnprocessedPixelData(t, rawData)
		// in-place write then converts to bytes
		if err = WriteUnprocessedValueData(metadata, pixelDataInfo.UnprocessedValueData,
			tc.rowIndex, tc.colIndex, tc.frameIndex, tc.sampleIndex, 100); err != nil {
			t.Errorf("WriteUnprocessedValueData: %v", err)
		}
		var buffer2 bytes.Buffer
		if err = dicom.Write(io.Writer(&buffer2), dataset); err != nil {
			t.Errorf("Write new DICOM obj to file: %v", err)
		}
		outputBytes := buffer2.Bytes()
		assertPixelDataInfo(t, outputBytes, tc.expectedPixelInfo)
	}
}

func TestWriteUnprocessedValueData_BigEndian(t *testing.T) {
	originDataset := dicom.Dataset{
		Elements: []*dicom.Element{
			mustNewElement(t, tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
			mustNewElement(t, tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
			mustNewElement(t, tag.TransferSyntaxUID, []string{uid.ExplicitVRBigEndian}),
			mustNewElement(t, tag.Rows, []int{2}),
			mustNewElement(t, tag.Columns, []int{2}),
			mustNewElement(t, tag.NumberOfFrames, []string{"2"}),
			mustNewElement(t, tag.BitsAllocated, []int{16}),
			mustNewElement(t, tag.SamplesPerPixel, []int{2}),

			mustNewElement(t, tag.PixelData, dicom.PixelDataInfo{
				IsEncapsulated: false,
				Frames: []*frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 16,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 2}},
						},
					},
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 16,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 5}},
						},
					},
				},
			}),
		},
	}
	var buffer bytes.Buffer
	if err := dicom.Write(io.Writer(&buffer), originDataset); err != nil {
		t.Errorf("write from dataset to DICOM bytes: %v", err)
	}
	rawData := buffer.Bytes()

	for _, tc := range []struct {
		sampleIndex, rowIndex, colIndex, frameIndex int
		expectedPixelInfo                           dicom.PixelDataInfo
	}{
		{
			sampleIndex: 0, colIndex: 1, rowIndex: 0, frameIndex: 0,
			expectedPixelInfo: dicom.PixelDataInfo{
				Frames: []*frame.Frame{
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {100, 2}, {1, 2}, {3, 2}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 5}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
				},
			},
		},
		{
			sampleIndex: 1, colIndex: 1, rowIndex: 0, frameIndex: 0,
			expectedPixelInfo: dicom.PixelDataInfo{
				Frames: []*frame.Frame{
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 100}, {1, 2}, {3, 2}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 5}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
				},
			},
		},
		{
			sampleIndex: 1, colIndex: 0, rowIndex: 1, frameIndex: 0,
			expectedPixelInfo: dicom.PixelDataInfo{
				Frames: []*frame.Frame{
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 100}, {3, 2}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 5}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
				},
			},
		},
		{
			sampleIndex: 1, colIndex: 0, rowIndex: 1, frameIndex: 1,
			expectedPixelInfo: dicom.PixelDataInfo{
				Frames: []*frame.Frame{
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 2}, {3, 2}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
					{NativeData: frame.NativeFrame{Data: [][]int{{1, 2}, {3, 2}, {1, 100}, {3, 5}}, BitsPerSample: 16, Rows: 2, Cols: 2}},
				},
			},
		},
	} {
		dataset, metadata, pixelDataInfo := setupUnprocessedPixelData(t, rawData)
		var err error
		// in-place write then converts to bytes
		if err = WriteUnprocessedValueData(metadata, pixelDataInfo.UnprocessedValueData,
			tc.rowIndex, tc.colIndex, tc.frameIndex, tc.sampleIndex, 100); err != nil {
			t.Errorf("WriteUnprocessedValueData: %v", err)
		}
		var buffer2 bytes.Buffer
		if err = dicom.Write(io.Writer(&buffer2), dataset); err != nil {
			t.Errorf("Write new DICOM obj to file: %v", err)
		}
		outputBytes := buffer2.Bytes()
		assertPixelDataInfo(t, outputBytes, tc.expectedPixelInfo)
	}
}

func setupUnprocessedPixelData(t testing.TB, rawData []byte) (dicom.Dataset, *PixelDataMetadata, dicom.PixelDataInfo) {
	dataset, err := dicom.Parse(bytes.NewReader(rawData), int64(len(rawData)), nil, dicom.SkipProcessingPixelDataValue())
	if err != nil {
		t.Errorf("dicom.Parse: %v", err)
	}
	metadata, err := GetPixelDataMetadata(&dataset)
	if err != nil {
		t.Errorf("GetPixelDataMetadata: %v", err)
	}
	pixelElement, err := dataset.FindElementByTag(tag.PixelData)
	if err != nil {
		t.Errorf("tag.PixelData not found: %v", err)
	}
	pixelDataInfo := dicom.MustGetPixelDataInfo(pixelElement.Value)
	if !pixelDataInfo.IntentionallyUnprocessed || pixelDataInfo.IsEncapsulated {
		t.Errorf("unexpected pixelDataInfo: IntentionallyUnprocessed=%t IsEncapsulated=%t",
			pixelDataInfo.IntentionallyUnprocessed, pixelDataInfo.IsEncapsulated)
	}
	// Sanity check if the dataset is suitable for in-place read-write
	if err = IsSafeForUnprocessedValueDataHandling(metadata, pixelDataInfo.UnprocessedValueData); err != nil {
		t.Errorf("IsSafeForUnprocessedValueDataHandling(%v)", err)
	}
	return dataset, metadata, pixelDataInfo
}

func assertPixelDataInfo(t testing.TB, rawData []byte, expectedPixelDataInfo dicom.PixelDataInfo) {
	dataset2, err := dicom.Parse(bytes.NewReader(rawData), int64(len(rawData)), nil)
	if err != nil {
		t.Errorf("read new create dcm file: %v", err)
	}
	pixelElement, err := dataset2.FindElementByTag(tag.PixelData)
	if err != nil {
		t.Errorf("pixel data not found: %v", err)
	}
	pixelDataInfo := dicom.MustGetPixelDataInfo(pixelElement.Value)
	if diff := cmp.Diff(pixelDataInfo, expectedPixelDataInfo, cmpopts.EquateErrors()); diff != "" {
		t.Errorf("assert pixelDataInfo in dicom file(%v): unexpected diff: %v", pixelDataInfo, diff)
	}
}
