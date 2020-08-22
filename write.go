package dicom

import (
	"errors"
	"io"
)

var ErrorUnimplemented = errors.New("this functionality is not yet implemented")

// TODO(suyashkumar): do we want to have a Writer struct that mirrors the Parser struct?

// WriteOption represents an option that can be passed to WriteDataset. Later options will override previous options if
// applicable.
type WriteOption func(*writeOptSet)

// WriteDataset will write the input DICOM dataset to the provided io.Writer as a complete DICOM (including any header
// information if available).
func WriteDataset(out io.Writer, ds *Dataset, opts ...WriteOption) error {
	return ErrorUnimplemented
}

// SkipVRVerification returns a WriteOption that skips VR verification.
func SkipVRVerification() WriteOption {
	return func(set *writeOptSet) {
		set.skipVRVerification = true
	}
}

// writeOptSet represents the flattened option set after all WriteOptions have been applied.
type writeOptSet struct {
	skipVRVerification bool
}

func toOptSet(opts ...WriteOption) *writeOptSet {
	optSet := &writeOptSet{}
	for _, opt := range opts {
		opt(optSet)
	}
	return optSet
}
