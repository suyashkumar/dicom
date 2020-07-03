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

/*
type FlatDatasetIterator struct {
	ElementsPtr *[]*Element
}

// Next returns the next element when iterating through the flattened dataset (which means even recursively
func (f FlatDatasetIterator) Next() (*Element, error) {

}
*/
