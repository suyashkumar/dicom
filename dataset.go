package dicom

import (
	"errors"
	"fmt"
	"strings"

	"github.com/suyashkumar/dicom/pkg/tag"
)

var ErrorElementNotFound = errors.New("element not found")

type Dataset struct {
	Elements []*Element `json:"elements"`
}

// FindElementByTag searches through the dataset and returns a pointer to the matching element.
// It DOES NOT search within Sequences as well.
func (d *Dataset) FindElementByTag(tag tag.Tag) (*Element, error) {
	for _, e := range d.Elements {
		if e.Tag == tag {
			return e, nil
		}
	}
	return nil, ErrorElementNotFound
}

// FindElementByTagNested searches through the dataset and returns a pointer to the matching element.
// This call searches through a flat representation of the dataset, including within sequences.
func (d *Dataset) FindElementByTagNested(tag tag.Tag) (*Element, error) {
	for e := range d.FlatIterator() {
		if e.Tag == tag {
			return e, nil
		}
	}
	return nil, ErrorElementNotFound
}

// FlatIterator returns a channel upon which every element in this Dataset will be sent,
// including elements nested inside sequences.
// Note that the sequence element itself is sent on the channel in addition to the child elements in the sequence.
// TODO(suyashkumar): decide if the sequence element itself should be sent or not
func (d *Dataset) FlatIterator() <-chan *Element {
	elemChan := make(chan *Element)
	go func() {
		flatElementsIterator(d.Elements, elemChan)
		close(elemChan)
	}()
	return elemChan
}

func flatElementsIterator(elems []*Element, elemChan chan<- *Element) {
	for _, elem := range elems {
		if elem.Value.ValueType() == Sequences {
			elemChan <- elem
			for _, seqItem := range elem.Value.GetValue().([]*SequenceItemValue) {
				flatElementsIterator(seqItem.elements, elemChan)
			}
			continue
		}
		elemChan <- elem
	}
}

// String returns a representation of this dataset as a string
func (d *Dataset) String() string {
	var b strings.Builder
	b.Grow(len(d.Elements) * 100) // Underestimate of the size of the final string in an attempt to limit buffer copying
	for elem := range d.flatIteratorWithLevel() {
		tabs := buildTabs(elem.l)
		var tagName string
		if tagInfo, err := tag.Find(elem.e.Tag); err == nil {
			tagName = tagInfo.Name
		}

		b.WriteString(fmt.Sprintf("%s[\n", tabs))
		b.WriteString(fmt.Sprintf("%s  Tag: %s\n", tabs, elem.e.Tag))
		b.WriteString(fmt.Sprintf("%s  Tag Name: %s\n", tabs, tagName))
		b.WriteString(fmt.Sprintf("%s  VR: %s\n", tabs, elem.e.ValueRepresentation))
		b.WriteString(fmt.Sprintf("%s  VR Raw: %s\n", tabs, elem.e.RawValueRepresentation))
		b.WriteString(fmt.Sprintf("%s  VL: %d\n", tabs, elem.e.ValueLength))
		b.WriteString(fmt.Sprintf("%s  Value: %d\n", tabs, elem.e.Value))
		b.WriteString(fmt.Sprintf("%s]\n\n", tabs))
	}
	return b.String()
}

type elementWithLevel struct {
	e *Element
	// l represents the nesting level of the Element
	l uint
}

func (d *Dataset) flatIteratorWithLevel() <-chan *elementWithLevel {
	elemChan := make(chan *elementWithLevel)
	go func() {
		flatElementsIteratorWithLevel(d.Elements, 0, elemChan)
		close(elemChan)
	}()
	return elemChan
}

func flatElementsIteratorWithLevel(elems []*Element, level uint, eWithLevelChan chan<- *elementWithLevel) {
	for _, elem := range elems {
		if elem.Value.ValueType() == Sequences {
			eWithLevelChan <- &elementWithLevel{elem, level}
			for _, seqItem := range elem.Value.GetValue().([]*SequenceItemValue) {
				flatElementsIteratorWithLevel(seqItem.elements, level+1, eWithLevelChan)
			}
			continue
		}
		eWithLevelChan <- &elementWithLevel{elem, level}
	}
}

func buildTabs(number uint) string {
	var b strings.Builder
	b.Grow(int(number))
	for i := 0; i < int(number); i++ {
		b.WriteString("\t")
	}
	return b.String()
}
