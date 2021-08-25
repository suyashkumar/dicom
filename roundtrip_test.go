package dicom_test

import (
	"bytes"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/dcmtest"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
	"testing"
)

// TestRoundTrip tests that we can read a dicom from disk, write the dicom back to a
// buffer, then read it back in and yield two identical datasets.
func TestRoundTrip(t *testing.T) {
	dcmtest.WalkIncludedFS(t, func(t *testing.T, tc dcmtest.FSTestCase, setupErr error) {

		buffer := bytes.NewBuffer(make([]byte, 0, tc.DCMStat.Size()))
		err := dicom.Write(
			buffer,
			tc.Dataset,
			dicom.SkipVRVerification(),
			dicom.SkipValueTypeVerification(),
		)
		if err != nil {
			t.Fatalf("error writing dataset: %v", err)
		}

		writtenDataset, err := dicom.Parse(buffer, int64(buffer.Len()), nil)
		if err != nil {
			t.Fatalf("error re-parsing written dataset: %v", err)
		}

		// Compare the datasets and report discrepancies.
		checkDatasets(t, tc.Dataset, writtenDataset)
	})
}

// checkDatasets checks that all written element values match the original.
func checkDatasets(t *testing.T, original, written dicom.Dataset) {
	// If the number of elements in a dataset are mismatched, that's an immediate exit.
	if len(original.Elements) != len(written.Elements) {
		// TODO: add report of what elements are missing (or added) from/to the
		//   original.
		t.Errorf(
			"element count mis-match, original had %v, written has %v",
			len(original.Elements),
			len(written.Elements),
		)
	}

	// Iterate over both datasets.
	originalIter := original.FlatStatefulIterator()
	writtenIter := written.FlatStatefulIterator()

	// We'll use the original dataset for our
	for originalIter.HasNext() {
		originalElm := originalIter.Next()

		// Get the element name for errors. Wll try to use the known name.
		tagInfo, err := tag.Find(originalElm.Tag)
		if err != nil {
			// Default to the tag number if we don't know the tag.
			tagInfo.Name = originalElm.Tag.String()
		}

		if !writtenIter.HasNext() {
			t.Errorf(
				"could not get next element from written dataset on source"+
					" tag %v",
				tagInfo.Name,
			)
		}
		writtenElm := writtenIter.Next()

		// Run a sub-test for this element.
		t.Run(tagInfo.Name, func(t *testing.T) {
			checkElement(t, originalElm, writtenElm)
		})
	}
}

// checkElement checks that two elements are identical.
func checkElement(t *testing.T, original, written *dicom.Element) {
	t.Log("VR:", original.RawValueRepresentation)

	if original.Tag != written.Tag {
		t.Fatalf(
			"element tag mismatch. original is '%v', written is '%v'",
			original.Tag,
			written.Tag,
		)
	}

	if original.RawValueRepresentation != written.RawValueRepresentation {
		t.Errorf(
			"element vr mismatch. original is '%v', written is '%v'",
			original.RawValueRepresentation,
			written.RawValueRepresentation,
		)
	}

	if original.ValueRepresentation != written.ValueRepresentation {
		t.Errorf(
			"element vrkind mismatch. original is '%v', written is '%v'",
			original.ValueRepresentation,
			written.ValueRepresentation,
		)
	}

	if original.ValueLength != written.ValueLength {
		t.Errorf(
			"element value length mismatch. original is '%v', written is '%v'",
			original.ValueLength,
			written.ValueLength,
		)
	}

	// We need to handle SQ (sequence) and PixelData specially.
	switch original.ValueRepresentation {
	case tag.VRSequence:
		originalSeq := original.Value.GetValue().([]*dicom.SequenceItemValue)
		writtenSeq := written.Value.GetValue().([]*dicom.SequenceItemValue)
		checkSequence(t, originalSeq, writtenSeq)
	case tag.VRPixelData:
		originalData := original.Value.GetValue().(dicom.PixelDataInfo)
		writtenData := original.Value.GetValue().(dicom.PixelDataInfo)
		checkPixelData(t, originalData, writtenData)
	default:
		// Otherwise we will just compare the values directly.
		if !cmp.Equal(original.Value.GetValue(), written.Value.GetValue()) {
			t.Errorf(
				"original value does not equal written value. Diff: %v",
				cmp.Diff(original.Value.GetValue(), written.Value.GetValue()),
			)
		}
	}
}

// checkSequence checks that two sequence elements are identical.
func checkSequence(t *testing.T, original, written []*dicom.SequenceItemValue) {
	// Because we are using a flat iterator in the main function, we don't need to
	// deeply compare each element within the sequences. Just that the sequences
	// have the same number of elements in the same order.
	if len(original) != len(written) {
		t.Fatalf(
			"sequence item count mismatch. original: %v, written: %v",
			len(original),
			len(written),
		)
	}

	for i, thisOriginal := range original {
		thisWritten := written[i]

		originalElements := thisOriginal.GetValue().([]*dicom.Element)
		writtenElements := thisWritten.GetValue().([]*dicom.Element)

		if len(originalElements) != len(writtenElements) {
			t.Fatalf(
				"element count mismatch for SequenceItemValue %v. original: %v, written: %v",
				i,
				len(originalElements),
				len(writtenElements),
			)
		}

		// Iterate over both sequences.
		for j, originalElm := range originalElements {
			writtenElm := writtenElements[j]

			// We just need to make sure the tags are equal to ensure ordering.
			if originalElm.Tag != writtenElm.Tag {
				t.Errorf(
					"SQ value %v, element %v tag mismatch. original: %v, written: %v",
					i,
					j,
					originalElm.Tag,
					writtenElm.Tag,
				)
			}
		}
	}
}

// checkPixelData checks that pixel data values match
func checkPixelData(t *testing.T, original, written dicom.PixelDataInfo) {
	if !original.IsEncapsulated == written.IsEncapsulated {
		t.Errorf(
			"PixelDataInfo.IsEncapsulated mismatch. original: %v written: %v",
			original.IsEncapsulated,
			written.IsEncapsulated,
		)
	}

	if len(original.Offsets) != len(written.Offsets) {
		t.Errorf(
			"PixelDataInfo.Offsets length mismatch. original: %v written: %v",
			len(original.Offsets),
			len(written.Offsets),
		)
	}

	if len(original.Frames) != len(written.Frames) {
		// If the frame count is wong, we don't need to compare the frames.
		t.Fatalf(
			"PixelDataInfo.Frame count mismatch. original: %v written: %v",
			len(original.Frames),
			len(written.Frames),
		)
	}

	for i, originalFrame := range original.Frames {
		writtenFrame := written.Frames[i]
		t.Run(fmt.Sprint("Frame", i), func(t *testing.T) {
			checkFrame(t, originalFrame, writtenFrame)
		})
	}
}

// checkFrame checks that the written frame matches the original frame.
func checkFrame(t *testing.T, original, written frame.Frame) {
	if original.Encapsulated != written.Encapsulated {
		t.Errorf(
			"Encapsulated mismatch. Original: %v Written: %v",
			original.Encapsulated,
			written.Encapsulated,
		)
	}

	if original.Encapsulated {
		originalData := original.EncapsulatedData.Data
		writtenData := written.EncapsulatedData.Data

		if len(originalData) != len(writtenData) {
			t.Fatalf(
				"oiginal data length %v does not match written lenfth %v",
				len(originalData),
				len(writtenData),
			)
		}
		return
	}

	if original.NativeData.BitsPerSample != written.NativeData.BitsPerSample {
		t.Errorf(
			"BitsPerSample mismatch. Original: %v Written: %v",
			original.NativeData.BitsPerSample,
			written.NativeData.BitsPerSample,
		)
	}

	if original.NativeData.Cols != written.NativeData.Cols {
		t.Errorf(
			"Cols mismatch. Original: %v Written: %v",
			original.NativeData.Cols,
			written.NativeData.Cols,
		)
	}

	if original.NativeData.Rows != written.NativeData.Rows {
		t.Errorf(
			"Rows mismatch. Original: %v Written: %v",
			original.NativeData.Rows,
			written.NativeData.Rows,
		)
	}

	if len(original.NativeData.Data) != len(written.NativeData.Data) {
		t.Fatalf(
			"Pixel count mismatch. Original: %v Written %v",
			len(original.NativeData.Data),
			len(written.NativeData.Data),
		)
	}

	// Check that all the pixel-values match.
	for i, originalPixel := range original.NativeData.Data {
		writtenPixel := written.NativeData.Data[i]

		if len(originalPixel) != len(writtenPixel) {
			t.Fatalf(
				"Pixel value count mismatch on pixel %v. oiginal: %v witten: %v",
				i,
				len(originalPixel),
				len(writtenPixel),
			)
		}

		for i2, originalValue := range originalPixel {
			writtenValue := writtenPixel[i2]

			if originalValue != writtenValue {
				t.Fatalf(
					"pixel value mismatch for pixel %v, value %v. Original: %v, Written: %v",
					i,
					i2,
					originalValue,
					writtenValue,
				)
			}
		}
	}
}
