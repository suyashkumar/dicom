package dicom

import (
	"encoding/json"
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
		wantCount int
		wantError error
	}{
		{
			name:      "strings",
			data:      []string{"a", "b"},
			wantValue: &stringsValue{value: []string{"a", "b"}},
			wantCount: 2,
			wantError: nil,
		},
		{
			name:      "floats",
			data:      []float64{1.11, 1.22},
			wantCount: 2,
			wantValue: &floatsValue{value: []float64{1.11, 1.22}},
			wantError: nil,
		},
		{
			name:      "ints",
			data:      []int{1, 2},
			wantCount: 2,
			wantValue: &intsValue{value: []int{1, 2}},
			wantError: nil,
		},
		{
			name:      "bytes",
			data:      []byte{0x00, 0x01},
			wantValue: &bytesValue{value: []byte{0x00, 0x01}},
			wantCount: 1,
			wantError: nil,
		},
		{
			// TODO: maybe enhance this case
			name:      "PixelDataInfo",
			data:      PixelDataInfo{IsEncapsulated: true},
			wantValue: &pixelDataValue{PixelDataInfo{IsEncapsulated: true}},
			wantCount: 0,
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
			wantCount: 1,
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
			if v.ValueCount() != tc.wantCount {
				t.Errorf("Value.ValueCount() expected count %v, got %v", tc.wantCount, v.ValueCount())
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
