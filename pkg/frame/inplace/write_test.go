package inplace

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
	require.NoError(t, dicom.Write(io.Writer(&buffer), originDataset))
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
		dataset, err := dicom.Parse(bytes.NewReader(rawData), int64(len(rawData)), nil, dicom.SkipProcessingPixelDataValue())
		require.NoError(t, err)
		metadata, err := GetPixelDataMetadata(&dataset)
		require.NoError(t, err)
		pixelElement, err := dataset.FindElementByTag(tag.PixelData)
		require.NoError(t, err)
		pixelDataInfo := dicom.MustGetPixelDataInfo(pixelElement.Value)
		assert.Equal(t, pixelDataInfo.IntentionallyUnprocessed, true)
		assert.Equal(t, pixelDataInfo.IsEncapsulated, false)
		// Sanity check if the dataset is suitable for in-place read-write
		require.NoError(t, IsSafeForUnprocessedValueDataHandling(metadata, pixelDataInfo.UnprocessedValueData))
		// in-place write then converts to bytes
		require.NoError(t, WriteUnprocessedValueData(metadata, pixelDataInfo.UnprocessedValueData, tc.rowIndex, tc.colIndex, tc.frameIndex, tc.sampleIndex, 100))
		var buffer2 bytes.Buffer
		require.NoError(t, dicom.Write(io.Writer(&buffer2), dataset))
		outputBytes := buffer2.Bytes()
		assertPixelDataInfo(t, outputBytes, tc.expectedPixelInfo)
	}
}

func assertPixelDataInfo(t testing.TB, rawData []byte, expectedPixelDataInfo dicom.PixelDataInfo) {
	dataset2, err := dicom.Parse(bytes.NewReader(rawData), int64(len(rawData)), nil)
	require.NoError(t, err)
	pixelElement, err := dataset2.FindElementByTag(tag.PixelData)
	require.NoError(t, err)
	pixelDataInfo := dicom.MustGetPixelDataInfo(pixelElement.Value)
	require.Equal(t, pixelDataInfo, expectedPixelDataInfo)
}
