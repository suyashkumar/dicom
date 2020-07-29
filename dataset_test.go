package dicom

import (
	"fmt"
	"testing"

	"github.com/suyashkumar/dicom/pkg/tag"
)

func makeSequenceElement(tg tag.Tag, items [][]*Element) *Element {
	sequenceItems := make([]*SequenceItemValue, 0, len(items))
	for _, item := range items {
		sequenceItems = append(sequenceItems, &SequenceItemValue{elements: item})
	}

	return &Element{
		Tag:                 tg,
		ValueRepresentation: tag.VRSequence,
		Value: &sequencesValue{
			value: sequenceItems,
		},
	}
}

func TestDataset_FindElementByTag(t *testing.T) {
	data := Dataset{
		Elements: []*Element{
			{
				Tag:                 tag.Rows,
				ValueRepresentation: tag.VRInt32List,
				Value: &intsValue{
					value: []int{100},
				},
			},
			{
				Tag:                 tag.Columns,
				ValueRepresentation: tag.VRInt32List,
				Value: &intsValue{
					value: []int{200},
				},
			},
		},
	}

	elem, err := data.FindElementByTag(tag.Rows)
	if err != nil {
		t.Errorf("FindElementByTag(%v): unexpected err: %v", tag.Rows, err)
	}

	if rows := MustGetInts(elem.Value)[0]; rows != 100 {
		t.Errorf("FindElementByTag(%v): want: %v, got: %v", tag.Rows, 100, rows)
	}
}

func ExampleDataset_FlatIterator() {
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

	data := Dataset{
		Elements: []*Element{
			{
				Tag:                 tag.Rows,
				ValueRepresentation: tag.VRInt32List,
				Value: &intsValue{
					value: []int{100},
				},
			},
			{
				Tag:                 tag.Columns,
				ValueRepresentation: tag.VRInt32List,
				Value: &intsValue{
					value: []int{200},
				},
			},
			makeSequenceElement(tag.AddOtherSequence, nestedData),
		},
	}

	for elem := range data.FlatIterator() {
		fmt.Println(elem.Tag)
	}

	// Note the output below includes all three leaf elements __as well as__ the sequence element's tag

	// Unordered output:
	// (0028,0010)
	// (0028,0011)
	// (0010,0010)
	// (0046,0102)
}

func ExampleDataset_String() {
	d := Dataset{
		Elements: []*Element{
			{
				Tag:                    tag.Rows,
				ValueRepresentation:    tag.VRInt32List,
				RawValueRepresentation: "UL",
				Value: &intsValue{
					value: []int{100},
				},
			},
			{
				Tag:                    tag.Columns,
				ValueRepresentation:    tag.VRInt32List,
				RawValueRepresentation: "UL",
				Value: &intsValue{
					value: []int{200},
				},
			},
		},
	}

	fmt.Println(d.String())

	// Output:
	// [
	//   Tag: (0028,0010)
	//   Tag Name: Rows
	//   VR: VRInt32List
	//   VR Raw: UL
	//   VL: 0
	//   Value: &{[100]}
	// ]
	//
	// [
	//   Tag: (0028,0011)
	//   Tag Name: Columns
	//   VR: VRInt32List
	//   VR Raw: UL
	//   VL: 0
	//   Value: &{[200]}
	// ]

}
