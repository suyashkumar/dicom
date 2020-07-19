package dicom

import (
	"errors"

	"github.com/suyashkumar/dicom/pkg/tag"
)

var ErrorElementNotFound = errors.New("element not found")

type Dataset struct {
	Elements []*Element
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
			for _, seqItem := range elem.Value.GetValue().([]*SequenceItemValue) {
				flatElementsIterator(seqItem.elements, elemChan)
			}
		}
		elemChan <- elem
	}
}
