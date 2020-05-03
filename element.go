package dicom

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

// BytesValue represents a value of []byte.
type BytesValue struct {
	value []byte
}

func (b *BytesValue) isElementValue()       {}
func (b *BytesValue) ValueType() ValueType  { return Bytes }
func (b *BytesValue) GetValue() interface{} { return b.value }
func (b *BytesValue) String() string {
	return fmt.Sprintf("%v", b.value)
}

// StringsValue represents a value of []string.
type StringsValue struct {
	value []string
}

func (s *StringsValue) isElementValue()       {}
func (s *StringsValue) ValueType() ValueType  { return Strings }
func (s *StringsValue) GetValue() interface{} { return s.value }
func (s *StringsValue) String() string {
	return fmt.Sprintf("%v", s.value)
}

// IntsValue represents a value of []int.
type IntsValue struct {
	value []int
}

func (s *IntsValue) isElementValue()       {}
func (s *IntsValue) ValueType() ValueType  { return Ints }
func (s *IntsValue) GetValue() interface{} { return s.value }
func (s *IntsValue) String() string {
	return fmt.Sprintf("%v", s.value)
}

type SequenceItemValue struct {
	elements []*Element
}

func (s *SequenceItemValue) isElementValue()       {}
func (s *SequenceItemValue) ValueType() ValueType  { return SequenceItem }
func (s *SequenceItemValue) GetValue() interface{} { return s.elements }
func (s *SequenceItemValue) String() string {
	// TODO: consider adding more sophisticated formatting
	return fmt.Sprintf("%+v", s.elements)
}

// SequencesValue represents a set of items in a DICOM sequence.
type SequencesValue struct {
	value []*SequenceItemValue
}

func (s *SequencesValue) isElementValue()       {}
func (s *SequencesValue) ValueType() ValueType  { return Sequences }
func (s *SequencesValue) GetValue() interface{} { return s.value }
func (s *SequencesValue) String() string {
	// TODO: consider adding more sophisticated formatting
	return fmt.Sprintf("%+v", s.value)
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
