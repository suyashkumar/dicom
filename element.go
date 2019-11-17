package dicom

import (
	"github.com/suyashkumar/dicom/pkg/tag"
)

type RawElement struct {
	Tag                    tag.Tag
	ValueRepresentation    tag.VRKind
	RawValueRepresentation string
	ValueLength            uint32
	UndefinedLength        bool
	Value                  []byte
}

type Value interface {
	// All types that can be a "Value" for an element will implement this empty method, similar to how protocol buffers
	// implement "oneof" in Go
	isElementValue()
	ValueType() ValueType
	GetValue() interface{}
}

type ValueType int

// Possible ValueTypes
const (
	String ValueType = iota
)

type Element struct {
	Tag                    tag.Tag
	ValueRepresentation    tag.VRKind
	RawValueRepresentation string
	ValueLength            uint32
	Value                  Value
}
