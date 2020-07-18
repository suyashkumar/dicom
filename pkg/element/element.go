package element

import (
	"fmt"

	"github.com/suyashkumar/dicom/pkg/frame"
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

func (e *Element) String() string {
	return fmt.Sprintf("Tag:%s\nVR:%s\nValue:%s\n\n", e.Tag.String(), e.ValueRepresentation.String(), e.Value.String())
}

type Value interface {
	// All types that can be a "Value" for an element will implement this empty method, similar to how protocol buffers
	// implement "oneof" in Go
	isElementValue()
	ValueType() ValueType
	GetValue() interface{} // TODO: rename to Get to read cleaner
	String() string
}

type ValueType int

// Possible ValueTypes
const (
	Strings ValueType = iota
	Bytes
	Ints
	PixelData
	SequenceItem
	Sequences
)

// Begin definitions of Values:

// BytesValue represents a Value of []byte.
type BytesValue struct {
	Value []byte
}

func (b *BytesValue) isElementValue()       {}
func (b *BytesValue) ValueType() ValueType  { return Bytes }
func (b *BytesValue) GetValue() interface{} { return b.Value }
func (b *BytesValue) String() string {
	return fmt.Sprintf("%v", b.Value)
}

// StringsValue represents a Value of []string.
type StringsValue struct {
	Value []string
}

func (s *StringsValue) isElementValue()       {}
func (s *StringsValue) ValueType() ValueType  { return Strings }
func (s *StringsValue) GetValue() interface{} { return s.Value }
func (s *StringsValue) String() string {
	return fmt.Sprintf("%v", s.Value)
}

// IntsValue represents a Value of []int.
type IntsValue struct {
	Value []int
}

func (s *IntsValue) isElementValue()       {}
func (s *IntsValue) ValueType() ValueType  { return Ints }
func (s *IntsValue) GetValue() interface{} { return s.Value }
func (s *IntsValue) String() string {
	return fmt.Sprintf("%v", s.Value)
}

type SequenceItemValue struct {
	Elements []*Element
}

func (s *SequenceItemValue) isElementValue()       {}
func (s *SequenceItemValue) ValueType() ValueType  { return SequenceItem }
func (s *SequenceItemValue) GetValue() interface{} { return s.Elements }
func (s *SequenceItemValue) String() string {
	// TODO: consider adding more sophisticated formatting
	return fmt.Sprintf("%+v", s.Elements)
}

// SequencesValue represents a set of items in a DICOM sequence.
type SequencesValue struct {
	Value []*SequenceItemValue
}

func (s *SequencesValue) isElementValue()       {}
func (s *SequencesValue) ValueType() ValueType  { return Sequences }
func (s *SequencesValue) GetValue() interface{} { return s.Value }
func (s *SequencesValue) String() string {
	// TODO: consider adding more sophisticated formatting
	return fmt.Sprintf("%+v", s.Value)
}

type PixelDataInfo struct {
	Frames         []frame.Frame // Frames
	IsEncapsulated bool
	Offsets        []uint32 // BasicOffsetTable
}

// PixelDataValue represents DICOM PixelData
type PixelDataValue struct {
	PixelDataInfo
}

func (e *PixelDataValue) isElementValue()       {}
func (e *PixelDataValue) ValueType() ValueType  { return PixelData }
func (e *PixelDataValue) GetValue() interface{} { return e.PixelDataInfo }
func (e *PixelDataValue) String() string {
	// TODO: consider adding more sophisticated formatting
	return ""
}

func MustGetInt(v Value) int {
	return v.GetValue().([]int)[0]
}

func MustGetInts(v Value) []int {
	return v.GetValue().([]int)
}

func MustGetString(v Value) string {
	return v.GetValue().([]string)[0]
}

func MustGetStrings(v Value) []string {
	return v.GetValue().([]string)
}
