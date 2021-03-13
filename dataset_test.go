package dicom

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func makeSequenceElement(tg tag.Tag, items [][]*Element) *Element {
	sequenceItems := make([]*SequenceItemValue, 0, len(items))
	for _, item := range items {
		sequenceItems = append(sequenceItems, &SequenceItemValue{elements: item})
	}

	return &Element{
		Tag:                    tg,
		ValueRepresentation:    tag.VRSequence,
		RawValueRepresentation: "SQ",
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
		t.Errorf("FindElementByTag(%v): got: %v, want: %v", tag.Rows, rows, 100)
	}
}

func TestDataset_FlatStatefulIterator(t *testing.T) {
	cases := []struct {
		name                 string
		dataset              Dataset
		expectedFlatElements []*Element
	}{
		{
			name: "flat dataset",
			dataset: Dataset{Elements: []*Element{
				mustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
				mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
			}},
			expectedFlatElements: []*Element{
				mustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
				mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
			},
		},
		{
			name: "nested dataset",
			dataset: Dataset{Elements: []*Element{
				makeSequenceElement(tag.AddOtherSequence, [][]*Element{
					// Item 1
					{
						mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
						// Nested Sequence.
						makeSequenceElement(tag.AnatomicRegionSequence, [][]*Element{
							{
								mustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
							},
						}),
					},
				}),
			}},
			expectedFlatElements: []*Element{
				// First, expect the entire SQ element
				makeSequenceElement(tag.AddOtherSequence, [][]*Element{
					// Item 1
					{
						mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
						// Nested Sequence.
						makeSequenceElement(tag.AnatomicRegionSequence, [][]*Element{
							{
								mustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
							},
						}),
					},
				}),
				// Then expect the inner elements
				mustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
				// Inner SQ element
				makeSequenceElement(tag.AnatomicRegionSequence, [][]*Element{
					{
						mustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
					},
				}),
				// Inner element of the inner SQ
				mustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var gotElems []*Element
			for iter := tc.dataset.FlatStatefulIterator(); iter.HasNext(); {
				gotElems = append(gotElems, iter.Next())
			}
			if diff := cmp.Diff(tc.expectedFlatElements, gotElems, cmp.AllowUnexported(allValues...)); diff != "" {
				t.Errorf("FlatStatefulIterator(%v) returned unexpected set of elements: %v", tc.dataset, diff)
			}
		})
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

func ExampleDataset_FlatStatefulIterator() {
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

	for iter := data.FlatStatefulIterator(); iter.HasNext(); {
		fmt.Println(iter.Next().Tag)
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
