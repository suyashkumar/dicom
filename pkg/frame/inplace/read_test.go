package inplace

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/suyashkumar/dicom/pkg/uid"
)

func mustNewElement(t *testing.T, tag tag.Tag, data interface{}) *dicom.Element {
	t.Helper()
	elem, err := dicom.NewElement(tag, data)
	if err != nil {
		t.Errorf("dicom.NewElement: %v", err)
	}
	return elem
}

func TestReadReadUnprocessedValueData(t *testing.T) {
	cases := []struct {
		Name         string
		existingData dicom.Dataset
	}{
		{
			Name: "5x5_1_frame_1_samples/pixel",
			existingData: dicom.Dataset{Elements: []*dicom.Element{
				mustNewElement(t, tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
				mustNewElement(t, tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
				mustNewElement(t, tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
				mustNewElement(t, tag.Rows, []int{5}),
				mustNewElement(t, tag.Columns, []int{5}),
				mustNewElement(t, tag.NumberOfFrames, []string{"1"}),
				mustNewElement(t, tag.BitsAllocated, []int{16}),
				mustNewElement(t, tag.SamplesPerPixel, []int{1}),
				mustNewElement(t, tag.PixelData, dicom.PixelDataInfo{
					IsEncapsulated: false,
					Frames: []*frame.Frame{
						{
							Encapsulated: false,
							NativeData: frame.NativeFrame{
								BitsPerSample: 16,
								Rows:          5,
								Cols:          5,
								Data: [][]int{
									{1}, {2}, {3}, {4}, {5},
									{0}, {0}, {0}, {0}, {0},
									{0}, {0}, {0}, {0}, {0},
									{0}, {0}, {0}, {0}, {0},
									{0}, {0}, {0}, {0}, {0},
								},
							},
						},
					},
				}),
			}},
		},
		{
			Name: "2x2, 3 frames, 1 samples/pixel",
			existingData: dicom.Dataset{
				Elements: []*dicom.Element{
					mustNewElement(t, tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
					mustNewElement(t, tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
					mustNewElement(t, tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
					mustNewElement(t, tag.Rows, []int{2}),
					mustNewElement(t, tag.Columns, []int{2}),
					mustNewElement(t, tag.NumberOfFrames, []string{"3"}),
					mustNewElement(t, tag.BitsAllocated, []int{16}),
					mustNewElement(t, tag.SamplesPerPixel, []int{1}),
					mustNewElement(t, tag.PixelData, dicom.PixelDataInfo{
						IsEncapsulated: false,
						Frames: []*frame.Frame{
							{
								Encapsulated: false,
								NativeData: frame.NativeFrame{
									BitsPerSample: 16,
									Rows:          2,
									Cols:          2,
									Data:          [][]int{{4}, {5}, {6}, {7}},
								},
							},
							{
								Encapsulated: false,
								NativeData: frame.NativeFrame{
									BitsPerSample: 16,
									Rows:          2,
									Cols:          2,
									Data:          [][]int{{1}, {2}, {3}, {2}},
								},
							},
							{
								Encapsulated: false,
								NativeData: frame.NativeFrame{
									BitsPerSample: 16,
									Rows:          2,
									Cols:          2,
									Data:          [][]int{{1}, {2}, {3}, {0}},
								},
							},
						},
					}),
				},
			},
		},
		{
			Name: "2x2, 2 frames, 2 samples/pixel",
			existingData: dicom.Dataset{
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
			},
		},
		{
			Name: "1x1, 3 frames, 3 samples/pixel, data bytes with padded 0",
			existingData: dicom.Dataset{
				Elements: []*dicom.Element{
					mustNewElement(t, tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"}),
					mustNewElement(t, tag.MediaStorageSOPInstanceUID, []string{"1.2.3.4.5.6.7"}),
					mustNewElement(t, tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian}),
					mustNewElement(t, tag.Rows, []int{1}),
					mustNewElement(t, tag.Columns, []int{1}),
					mustNewElement(t, tag.NumberOfFrames, []string{"3"}),
					mustNewElement(t, tag.BitsAllocated, []int{8}),
					mustNewElement(t, tag.SamplesPerPixel, []int{3}),

					mustNewElement(t, tag.PixelData, dicom.PixelDataInfo{
						IsEncapsulated: false,
						Frames: []*frame.Frame{
							{
								Encapsulated: false,
								NativeData: frame.NativeFrame{
									BitsPerSample: 8, Rows: 1, Cols: 1,
									Data: [][]int{{1, 2, 3}},
								},
							}, {
								Encapsulated: false,
								NativeData: frame.NativeFrame{
									BitsPerSample: 8, Rows: 1, Cols: 1,
									Data: [][]int{{4, 5, 6}},
								},
							}, {
								Encapsulated: false,
								NativeData: frame.NativeFrame{
									BitsPerSample: 8, Rows: 1, Cols: 1,
									Data: [][]int{{7, 8, 9}},
								},
							},
						},
					}),
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var filesOut bytes.Buffer
			var err error
			if err = dicom.Write(io.Writer(&filesOut), tc.existingData); err != nil {
				t.Errorf("Write DICOM obj to bytes: %v", err)
			}
			rawData := filesOut.Bytes()

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
			if pixelDataInfo.IsEncapsulated || !pixelDataInfo.IntentionallyUnprocessed {
				t.Errorf("pixelDataInfo should be IntentionallyUnprocessed: IsEncapsulated=%t IntentionallyUnprocessed%t",
					pixelDataInfo.IsEncapsulated, pixelDataInfo.IntentionallyUnprocessed)
			}

			if err = IsSafeForUnprocessedValueDataHandling(metadata, pixelDataInfo.UnprocessedValueData); err != nil {
				t.Errorf("IsSafeForUnprocessedValueDataHandling(%v)", err)
			}

			originPixelElement, err := tc.existingData.FindElementByTag(tag.PixelData)
			if err != nil {
				t.Errorf("Find tag.PixelData(%v)", err)
			}
			originPixelDataInfo := dicom.MustGetPixelDataInfo(originPixelElement.Value)

			for i := 0; i < metadata.Frames; i++ {
				originData := originPixelDataInfo.Frames[i].NativeData.Data
				inplaceData, err := ReadUnprocessedValueData(metadata, pixelDataInfo.UnprocessedValueData, i)
				if err != nil {
					t.Errorf("ReadUnprocessedValueData(%v)", err)
				}
				assertUnprocessedPixelData(t, originData, inplaceData, fmt.Sprint("frame:", i))
			}
		})
	}
}

func assertUnprocessedPixelData(t *testing.T, inplaceData [][]int, originData [][]int, msg string) {
	if len(inplaceData[0]) != len(originData) || len(inplaceData) != len(originData[0]) {
		t.Errorf("Mismatch data size: origin-data(%d-%d) inplace-data(%d-%d)",
			len(originData), len(originData[0]), len(inplaceData), len(inplaceData[0]))
	}
	tmp := make([][]int, len(originData))
	for i := 0; i < len(tmp); i++ {
		tmp[i] = make([]int, len(originData[0]))
		for j := 0; j < len(originData[0]); j++ {
			tmp[i][j] = inplaceData[j][i]
		}
	}
	if diff := cmp.Diff(originData, tmp, cmpopts.EquateErrors()); diff != "" {
		t.Errorf("assertUnprocessedPixelData(%v-%s): unexpected diff: %v", inplaceData, msg, diff)
	}
}
