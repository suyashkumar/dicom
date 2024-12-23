package dicom

import (
	"encoding/json"
	"testing"

	"github.com/codeninja55/dicom/pkg/frame"
	"github.com/google/go-cmp/cmp"

	"github.com/codeninja55/dicom/pkg/tag"
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

func TestNewValue_UnexpectedType(t *testing.T) {
	data := 10
	_, err := NewValue(data)
	if err != ErrorUnexpectedDataType {
		t.Errorf("NewValue(%v) expected an error. got: %v, want: %v", data, err, ErrorUnexpectedDataType)
	}
}

func TestElement_Equals(t *testing.T) {
	// Some general sanity checks below, not every possible case is included.
	cases := []struct {
		name      string
		a         *Element
		b         *Element
		wantEqual bool
	}{
		{
			name:      "EqualNilElements",
			a:         nil,
			b:         nil,
			wantEqual: true,
		},
		{
			name:      "UnequalNilElement",
			a:         nil,
			b:         mustNewElement(tag.FloatingPointValue, []float64{1.23}),
			wantEqual: false,
		},
		{
			name:      "EqualFloats",
			a:         mustNewElement(tag.FloatingPointValue, []float64{1.23, 4.40, 5.50}),
			b:         mustNewElement(tag.FloatingPointValue, []float64{1.23, 4.40, 5.50}),
			wantEqual: true,
		},
		{
			name:      "UnequalLenFloats",
			a:         mustNewElement(tag.FloatingPointValue, []float64{1.23, 4.40}),
			b:         mustNewElement(tag.FloatingPointValue, []float64{1.23, 4.40, 5.50}),
			wantEqual: false,
		},
		{
			name:      "UnequalFloats",
			a:         mustNewElement(tag.FloatingPointValue, []float64{1.23, 4.40, 10.1}),
			b:         mustNewElement(tag.FloatingPointValue, []float64{1.23, 4.40, 5.50}),
			wantEqual: false,
		},
		{
			name:      "EqualInts",
			a:         mustNewElement(tag.Rows, []int{1, 2, 3}),
			b:         mustNewElement(tag.Rows, []int{1, 2, 3}),
			wantEqual: true,
		},
		{
			name:      "UnequalInts",
			a:         mustNewElement(tag.Rows, []int{1, 2, 6}),
			b:         mustNewElement(tag.Rows, []int{1, 2, 3}),
			wantEqual: false,
		},
		{
			name:      "UnequalLenInts",
			a:         mustNewElement(tag.Rows, []int{1, 6}),
			b:         mustNewElement(tag.Rows, []int{1, 2, 3}),
			wantEqual: false,
		},
		{
			name:      "EqualBytes",
			a:         mustNewElement(tag.AirCounts, []byte{1, 2, 3}),
			b:         mustNewElement(tag.AirCounts, []byte{1, 2, 3}),
			wantEqual: true,
		},
		{
			name:      "UnequalBytes",
			a:         mustNewElement(tag.AirCounts, []byte{1, 2, 4}),
			b:         mustNewElement(tag.AirCounts, []byte{1, 2, 3}),
			wantEqual: false,
		},
		{
			name:      "UnequalLenBytes",
			a:         mustNewElement(tag.AirCounts, []byte{1, 2, 3}),
			b:         mustNewElement(tag.AirCounts, []byte{1, 2}),
			wantEqual: false,
		},
		{
			name:      "EqualStrings",
			a:         mustNewElement(tag.PatientName, []string{"John", "Smith"}),
			b:         mustNewElement(tag.PatientName, []string{"John", "Smith"}),
			wantEqual: true,
		},
		{
			name:      "UnequalStrings",
			a:         mustNewElement(tag.PatientName, []string{"John", "Doe"}),
			b:         mustNewElement(tag.PatientName, []string{"John", "Smith"}),
			wantEqual: false,
		},
		{
			name:      "UnequalLenStrings",
			a:         mustNewElement(tag.PatientName, []string{"John"}),
			b:         mustNewElement(tag.PatientName, []string{"John", "Smith"}),
			wantEqual: false,
		},
		{
			name: "EqualNativePixelData",
			a: mustNewElement(tag.PixelData, PixelDataInfo{
				IsEncapsulated: false,
				Frames: []*frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 8,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1}, {2}, {3}, {4}},
						},
					},
				},
			}),
			b: mustNewElement(tag.PixelData, PixelDataInfo{
				IsEncapsulated: false,
				Frames: []*frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 8,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1}, {2}, {3}, {4}},
						},
					},
				},
			}),
			wantEqual: true,
		},
		{
			name: "UnequalNativePixelData",
			a: mustNewElement(tag.PixelData, PixelDataInfo{
				IsEncapsulated: false,
				Frames: []*frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 8,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1}, {2}, {3}, {6}},
						},
					},
				},
			}),
			b: mustNewElement(tag.PixelData, PixelDataInfo{
				IsEncapsulated: false,
				Frames: []*frame.Frame{
					{
						Encapsulated: false,
						NativeData: frame.NativeFrame{
							BitsPerSample: 8,
							Rows:          2,
							Cols:          2,
							Data:          [][]int{{1}, {2}, {3}, {4}},
						},
					},
				},
			}),
			wantEqual: false,
		},
		{
			name: "EqualSequences",
			a: makeSequenceElement(tag.AddOtherSequence, [][]*Element{
				// Item 1.
				{
					{
						Tag:                    tag.PatientName,
						ValueRepresentation:    tag.VRStringList,
						RawValueRepresentation: "PN",
						Value:                  &stringsValue{value: []string{"Bob", "Jones"}},
					},
				},
				// Item 2.
				{
					{
						Tag:                    tag.PatientName,
						ValueRepresentation:    tag.VRStringList,
						RawValueRepresentation: "PN",
						Value:                  &stringsValue{value: []string{"Bob", "Jones"}},
					},
					{
						Tag:                    tag.Rows,
						ValueRepresentation:    tag.VRUInt16List,
						RawValueRepresentation: "US",
						Value:                  &intsValue{value: []int{100}},
					},
				},
			}),
			b: makeSequenceElement(tag.AddOtherSequence, [][]*Element{
				// Item 1.
				{
					{
						Tag:                    tag.PatientName,
						ValueRepresentation:    tag.VRStringList,
						RawValueRepresentation: "PN",
						Value:                  &stringsValue{value: []string{"Bob", "Jones"}},
					},
				},
				// Item 2.
				{
					{
						Tag:                    tag.PatientName,
						ValueRepresentation:    tag.VRStringList,
						RawValueRepresentation: "PN",
						Value:                  &stringsValue{value: []string{"Bob", "Jones"}},
					},
					{
						Tag:                    tag.Rows,
						ValueRepresentation:    tag.VRUInt16List,
						RawValueRepresentation: "US",
						Value:                  &intsValue{value: []int{100}},
					},
				},
			}),
			wantEqual: true,
		},
		{
			name: "UnequalSequences",
			a: makeSequenceElement(tag.AddOtherSequence, [][]*Element{
				// Item 1.
				{
					{
						Tag:                    tag.PatientName,
						ValueRepresentation:    tag.VRStringList,
						RawValueRepresentation: "PN",
						//	Smith instead of Jones below causes inequality.
						Value: &stringsValue{value: []string{"Bob", "Smith"}},
					},
				},
				// Item 2.
				{
					{
						Tag:                    tag.PatientName,
						ValueRepresentation:    tag.VRStringList,
						RawValueRepresentation: "PN",
						Value:                  &stringsValue{value: []string{"Bob", "Jones"}},
					},
					{
						Tag:                    tag.Rows,
						ValueRepresentation:    tag.VRUInt16List,
						RawValueRepresentation: "US",
						Value:                  &intsValue{value: []int{100}},
					},
				},
			}),
			b: makeSequenceElement(tag.AddOtherSequence, [][]*Element{
				// Item 1.
				{
					{
						Tag:                    tag.PatientName,
						ValueRepresentation:    tag.VRStringList,
						RawValueRepresentation: "PN",
						Value:                  &stringsValue{value: []string{"Bob", "Jones"}},
					},
				},
				// Item 2.
				{
					{
						Tag:                    tag.PatientName,
						ValueRepresentation:    tag.VRStringList,
						RawValueRepresentation: "PN",
						Value:                  &stringsValue{value: []string{"Bob", "Jones"}},
					},
					{
						Tag:                    tag.Rows,
						ValueRepresentation:    tag.VRUInt16List,
						RawValueRepresentation: "US",
						Value:                  &intsValue{value: []int{100}},
					},
				},
			}),
			wantEqual: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.a.Equals(tc.b) != tc.wantEqual {
				t.Errorf("Element.Equals(%v, %v) != %v", tc.a, tc.b, tc.wantEqual)
			}
		})
	}
}
