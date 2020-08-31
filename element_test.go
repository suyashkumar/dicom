package dicom

import (
	"encoding/json"
	"testing"

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

	want := `{"tag":{"Group":70,"Element":258},"VR":9,"rawVR":"","valueLength":0,"value":[[{"tag":{"Group":16,"Element":16},"VR":2,"rawVR":"","valueLength":0,"value":["Bob"]}]]}`

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
	data := []string{"a", "b", "c"}
	v, err := NewValue(data)
	if err != nil {
		t.Fatalf("NewValue(%v) returned unexpected error: %v", data, err)
	}
	if v.ValueType() != Strings {
		t.Errorf("NewValue(%v) unexpected ValueType. got: %v, want: %v", data, v.ValueType(), Strings)
	}
	// TODO(suyashkumar): add tests for the remaining types
}

func TestNewValue_UnexpectedType(t *testing.T) {
	data := 10
	_, err := NewValue(data)
	if err != ErrorUnexpectedDataType {
		t.Errorf("NewValue(%v) expected an error. got: %v, want: %v", data, err, ErrorUnexpectedDataType)
	}
}
