package dicom

import (
	"encoding/json"
	"github.com/suyashkumar/dicom/pkg/frame"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/suyashkumar/dicom/pkg/tag"
)

func TestElement_MarshalJSON_NestedElements(t *testing.T) {
	nestedData := [][]*Element{
		{
			{
				Tag:                 tag.PatientName,
				ValueRepresentation: tag.VRString,
				Value: &stringsValue{
					value: []string{"Bob"},
				},
			},
		},
	}
	seqElement := makeSequenceElement(tag.AddOtherSequence, nestedData)

	want := `{"tag":{"Group":70,"Element":258},"VR":9,"rawVR":"SQ","valueLength":0,"value":[[{"tag":{"Group":16,"Element":16},"VR":2,"rawVR":"","valueLength":0,"value":["Bob"]}]]}`

	j, err := json.Marshal(seqElement)
	if err != nil {
		t.Errorf("unexpected error marshaling json: %v", err)
	}
	if string(j) != want {
		t.Errorf("json.Marshal(%v) produced incorrect output. got: %s, want: %s", seqElement, string(j), want)
	}
}

func TestElement_String(t *testing.T) {
	e := &Element{
		Tag:                    tag.Rows,
		ValueRepresentation:    tag.VRInt32List,
		RawValueRepresentation: "US",
		Value: &intsValue{
			value: []int{100},
		},
	}
	want := "[\n" +
		"  Tag: (0028,0010)\n" +
		"  Tag Name: Rows\n" +
		"  VR: VRInt32List\n" +
		"  VR Raw: US\n" +
		"  VL: 0\n" +
		"  Value: [100]\n" +
		"]\n\n"
	got := e.String()
	if want != got {
		t.Errorf("String(%v) unexpected diff. got:\n%s want:\n%s", e, got, want)
	}

}

func TestNewValue(t *testing.T) {
	cases := []struct {
		name      string
		data      interface{}
		wantValue Value
		wantError error
	}{
		{
			name:      "strings",
			data:      []string{"a", "b"},
			wantValue: &stringsValue{value: []string{"a", "b"}},
			wantError: nil,
		},
		{
			name:      "floats",
			data:      []float64{1.11, 1.22},
			wantValue: &floatsValue{value: []float64{1.11, 1.22}},
			wantError: nil,
		},
		{
			name:      "ints",
			data:      []int{1, 2},
			wantValue: &intsValue{value: []int{1, 2}},
			wantError: nil,
		},
		{
			name:      "bytes",
			data:      []byte{0x00, 0x01},
			wantValue: &bytesValue{value: []byte{0x00, 0x01}},
			wantError: nil,
		},
		{
			// TODO: maybe enhance this case
			name:      "PixelDataInfo",
			data:      PixelDataInfo{IsEncapsulated: true},
			wantValue: &pixelDataValue{PixelDataInfo{IsEncapsulated: true}},
			wantError: nil,
		},
		{
			name: "sequence",
			data: [][]*Element{
				{
					{
						Tag:                 tag.PatientName,
						ValueRepresentation: tag.VRString,
						Value: &stringsValue{
							value: []string{"Bob"},
						},
					},
				},
			},
			wantValue: &sequencesValue{value: []*SequenceItemValue{
				{
					elements: []*Element{
						{
							Tag:                 tag.PatientName,
							ValueRepresentation: tag.VRString,
							Value: &stringsValue{
								value: []string{"Bob"},
							},
						},
					},
				},
			}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v, err := NewValue(tc.data)
			if err != tc.wantError {
				t.Fatalf("NewValue(%v) returned unexpected error: %v", tc.data, err)
			}
			if diff := cmp.Diff(v, tc.wantValue, cmp.AllowUnexported(allValues...)); diff != "" {
				t.Errorf("NewValue(%v) unexpected value. diff: %v", tc.data, diff)
			}
		})
	}
}

func TestValue_IsEmpty(t *testing.T) {
	cases := []struct {
		name            string
		value           Value
		expectedIsEmpty bool
	}{
		// STRINGS VALUE
		{
			name:            "strings_none_empty",
			value:           &stringsValue{value: []string{"a", "b"}},
			expectedIsEmpty: false,
		},
		{
			name:            "strings_one_empty",
			value:           &stringsValue{value: []string{"a", ""}},
			expectedIsEmpty: false,
		},
		{
			name:            "strings_all_empty",
			value:           &stringsValue{value: []string{"", ""}},
			expectedIsEmpty: true,
		},
		{
			name:            "strings_single_empty",
			value:           &stringsValue{value: []string{""}},
			expectedIsEmpty: true,
		},

		// INTS VALUE
		{
			name:            "ints_none_empty",
			value:           &intsValue{value: []int{1, 2}},
			expectedIsEmpty: false,
		},
		{
			name:            "ints_one_empty",
			value:           &intsValue{value: []int{0, 2}},
			expectedIsEmpty: false,
		},
		{
			name:            "ints_all_empty",
			value:           &intsValue{value: []int{0, 0}},
			expectedIsEmpty: true,
		},
		{
			name:            "ints_single_empty",
			value:           &intsValue{value: []int{0}},
			expectedIsEmpty: true,
		},

		// FLOATS VALUE
		{
			name:            "floats_none_empty",
			value:           &floatsValue{value: []float64{1, 2}},
			expectedIsEmpty: false,
		},
		{
			name:            "floats_one_empty",
			value:           &floatsValue{value: []float64{0, 2}},
			expectedIsEmpty: false,
		},
		{
			name:            "floats_all_empty",
			value:           &floatsValue{value: []float64{0, 0}},
			expectedIsEmpty: true,
		},
		{
			name:            "floats_single_empty",
			value:           &floatsValue{value: []float64{0}},
			expectedIsEmpty: true,
		},

		// BYTES
		{
			name:            "bytes_not_empty",
			value:           &bytesValue{value: []byte{0x0, 0x1}},
			expectedIsEmpty: false,
		},
		{
			name:            "bytes_empty",
			value:           &bytesValue{value: []byte{}},
			expectedIsEmpty: true,
		},

		// PIXEL DATA
		{
			name:            "PixelData_empty",
			value:           &pixelDataValue{PixelDataInfo{IsEncapsulated: true}},
			expectedIsEmpty: true,
		},
		{
			name: "PixelData_not_empty",
			value: &pixelDataValue{
				PixelDataInfo{
					Frames: []frame.Frame{
						{},
					},
					IsEncapsulated: true,
				},
			},
			expectedIsEmpty: false,
		},

		// SEQUENCES
		{
			name:            "sequence_empty",
			expectedIsEmpty: true,
			value:           &sequencesValue{value: []*SequenceItemValue{}},
		},
		{
			name:            "sequence_empty_sub_elements",
			expectedIsEmpty: true,
			value: &sequencesValue{value: []*SequenceItemValue{
				{
					elements: []*Element{
						{
							Tag:                 tag.PatientName,
							ValueRepresentation: tag.VRString,
							Value: &stringsValue{
								value: []string{""},
							},
						},
					},
				},
			}},
		},
		{
			name:            "sequence_multiple_empty_sub_elements",
			expectedIsEmpty: true,
			value: &sequencesValue{value: []*SequenceItemValue{
				{
					elements: []*Element{
						{
							Tag:                 tag.PatientName,
							ValueRepresentation: tag.VRString,
							Value: &stringsValue{
								value: []string{""},
							},
						},
						{
							Tag:                 tag.SOPInstanceUID,
							ValueRepresentation: tag.VRString,
							Value: &stringsValue{
								value: []string{""},
							},
						},
					},
				},
			}},
		},
		{
			name:            "sequence_not_empty",
			expectedIsEmpty: false,
			value: &sequencesValue{value: []*SequenceItemValue{
				{
					elements: []*Element{
						{
							Tag:                 tag.PatientName,
							ValueRepresentation: tag.VRString,
							Value: &stringsValue{
								value: []string{"Bob"},
							},
						},
					},
				},
			}},
		},
		{
			name:            "sequence_one_empty_sub_elements",
			expectedIsEmpty: false,
			value: &sequencesValue{value: []*SequenceItemValue{
				{
					elements: []*Element{
						{
							Tag:                 tag.PatientName,
							ValueRepresentation: tag.VRString,
							Value: &stringsValue{
								value: []string{"bob"},
							},
						},
						{
							Tag:                 tag.SOPInstanceUID,
							ValueRepresentation: tag.VRString,
							Value: &stringsValue{
								value: []string{""},
							},
						},
					},
				},
			}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			isEmpty := tc.value.IsEmpty()
			if isEmpty != tc.expectedIsEmpty {
				t.Errorf(
					"Value.IsEmpty() returned %v, expected %v",
					isEmpty,
					tc.expectedIsEmpty,
				)
			}
		})
	}
}

func TestNewValue_UnexpectedType(t *testing.T) {
	data := 10
	_, err := NewValue(data)
	if err != ErrorUnexpectedDataType {
		t.Errorf("NewValue(%v) expected an error. got: %v, want: %v", data, err, ErrorUnexpectedDataType)
	}
}
