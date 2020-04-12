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

type Element struct {
	Tag                    tag.Tag
	ValueRepresentation    tag.VRKind
	RawValueRepresentation string
	ValueLength            uint32
	Value                  Value
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
	Strings ValueType = iota
	Bytes
	Ints
	ElementPtrs
)

// Begin definitions of Values:

// BytesValue represents a value of []byte.
type BytesValue struct {
	value []byte
}

func (b *BytesValue) isElementValue()       {}
func (b *BytesValue) ValueType() ValueType  { return Bytes }
func (b *BytesValue) GetValue() interface{} { return b.value }

// StringsValue represents a value of []string.
type StringsValue struct {
	value []string
}

func (s *StringsValue) isElementValue()       {}
func (s *StringsValue) ValueType() ValueType  { return Strings }
func (s *StringsValue) GetValue() interface{} { return s.value }

// IntsValue represents a value of []int.
type IntsValue struct {
	value []int
}

func (s *IntsValue) isElementValue()       {}
func (s *IntsValue) ValueType() ValueType  { return Ints }
func (s *IntsValue) GetValue() interface{} { return s.value }

// ElementPtrsValue represents a slice of pointers to other Elements (in the case of sequences)
type ElementPtrsValue struct {
	value []*Element
}

func (e *ElementPtrsValue) isElementValue()       {}
func (e *ElementPtrsValue) ValueType() ValueType  { return ElementPtrs }
func (e *ElementPtrsValue) GetValue() interface{} { return e.value }
