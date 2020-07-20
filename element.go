package dicom

import (
	"encoding/json"
	"fmt"

	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
)

type Element struct {
	Tag                    tag.Tag    `json:"tag"`
	ValueRepresentation    tag.VRKind `json:"VR"`
	RawValueRepresentation string     `json:"rawVR"`
	ValueLength            uint32     `json:"valueLength"`
	Value                  Value      `json:"value"`
}

func (e *Element) String() string {
	var tagName string
	if tagInfo, err := tag.Find(e.Tag); err == nil {
		tagName = tagInfo.Name
	}
	return fmt.Sprintf("[\n  Tag: %s\n  Tag Name: %s\n  VR: %s\n  VR Raw: %s\n  VL: %d\n  Value: %s\n]\n\n",
		e.Tag.String(),
		tagName,
		e.ValueRepresentation.String(),
		e.RawValueRepresentation,
		e.ValueLength,
		e.Value.String())
}

type Value interface {
	// All types that can be a "Value" for an element will implement this empty method, similar to how protocol buffers
	// implement "oneof" in Go
	isElementValue()
	ValueType() ValueType
	GetValue() interface{} // TODO: rename to Get to read cleaner
	String() string
	MarshalJSON() ([]byte, error)
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
	value []byte `json:"value"`
}

func (b *BytesValue) isElementValue()       {}
func (b *BytesValue) ValueType() ValueType  { return Bytes }
func (b *BytesValue) GetValue() interface{} { return b.value }
func (b *BytesValue) String() string {
	return fmt.Sprintf("%v", b.value)
}
func (b *BytesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.value)
}

// StringsValue represents a value of []string.
type StringsValue struct {
	value []string `json:"value"`
}

func (s *StringsValue) isElementValue()       {}
func (s *StringsValue) ValueType() ValueType  { return Strings }
func (s *StringsValue) GetValue() interface{} { return s.value }
func (s *StringsValue) String() string {
	return fmt.Sprintf("%v", s.value)
}
func (s *StringsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

// IntsValue represents a value of []int.
type IntsValue struct {
	value []int `json:"value"`
}

func (s *IntsValue) isElementValue()       {}
func (s *IntsValue) ValueType() ValueType  { return Ints }
func (s *IntsValue) GetValue() interface{} { return s.value }
func (s *IntsValue) String() string {
	return fmt.Sprintf("%v", s.value)
}
func (s *IntsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

type SequenceItemValue struct {
	elements []*Element `json:"elements"`
}

func (s *SequenceItemValue) isElementValue()       {}
func (s *SequenceItemValue) ValueType() ValueType  { return SequenceItem }
func (s *SequenceItemValue) GetValue() interface{} { return s.elements }
func (s *SequenceItemValue) String() string {
	// TODO: consider adding more sophisticated formatting
	return fmt.Sprintf("%+v", s.elements)
}
func (s *SequenceItemValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.elements)
}

// SequencesValue represents a set of items in a DICOM sequence.
type SequencesValue struct {
	value []*SequenceItemValue `json:"sequenceItems"`
}

func (s *SequencesValue) isElementValue()       {}
func (s *SequencesValue) ValueType() ValueType  { return Sequences }
func (s *SequencesValue) GetValue() interface{} { return s.value }
func (s *SequencesValue) String() string {
	// TODO: consider adding more sophisticated formatting
	return fmt.Sprintf("%+v", s.value)
}
func (s *SequencesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

type PixelDataInfo struct {
	Frames         []frame.Frame // Frames
	IsEncapsulated bool          `json:"isEncapsulated"`
	Offsets        []uint32      // BasicOffsetTable
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
func (s *PixelDataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.PixelDataInfo)
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
