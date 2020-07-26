package dicom

import (
	"encoding/json"
	"fmt"
	"log"

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
	// Strings represents an underlying value of []string
	Strings ValueType = iota
	// Bytes represents an underlying value of []byte
	Bytes
	// Ints represents an underlying value of []int
	Ints
	// PixelData represents an underlying value of PixelDataInfo
	PixelData
	// SequenceItem represents an underlying value of []*Element
	SequenceItem
	// Sequences represents an underlying value of []SequenceItem
	Sequences
)

// Begin definitions of Values:

// bytesValue represents a value of []byte.
type bytesValue struct {
	value []byte `json:"value"`
}

func (b *bytesValue) isElementValue()       {}
func (b *bytesValue) ValueType() ValueType  { return Bytes }
func (b *bytesValue) GetValue() interface{} { return b.value }
func (b *bytesValue) String() string {
	return fmt.Sprintf("%v", b.value)
}
func (b *bytesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.value)
}

// stringsValue represents a value of []string.
type stringsValue struct {
	value []string `json:"value"`
}

func (s *stringsValue) isElementValue()       {}
func (s *stringsValue) ValueType() ValueType  { return Strings }
func (s *stringsValue) GetValue() interface{} { return s.value }
func (s *stringsValue) String() string {
	return fmt.Sprintf("%v", s.value)
}
func (s *stringsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

// intsValue represents a value of []int.
type intsValue struct {
	value []int `json:"value"`
}

func (s *intsValue) isElementValue()       {}
func (s *intsValue) ValueType() ValueType  { return Ints }
func (s *intsValue) GetValue() interface{} { return s.value }
func (s *intsValue) String() string {
	return fmt.Sprintf("%v", s.value)
}
func (s *intsValue) MarshalJSON() ([]byte, error) {
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

// sequencesValue represents a set of items in a DICOM sequence.
type sequencesValue struct {
	value []*SequenceItemValue `json:"sequenceItems"`
}

func (s *sequencesValue) isElementValue()       {}
func (s *sequencesValue) ValueType() ValueType  { return Sequences }
func (s *sequencesValue) GetValue() interface{} { return s.value }
func (s *sequencesValue) String() string {
	// TODO: consider adding more sophisticated formatting
	return fmt.Sprintf("%+v", s.value)
}
func (s *sequencesValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.value)
}

type PixelDataInfo struct {
	Frames         []frame.Frame // Frames
	IsEncapsulated bool          `json:"isEncapsulated"`
	Offsets        []uint32      // BasicOffsetTable
}

// pixelDataValue represents DICOM PixelData
type pixelDataValue struct {
	PixelDataInfo
}

func (e *pixelDataValue) isElementValue()       {}
func (e *pixelDataValue) ValueType() ValueType  { return PixelData }
func (e *pixelDataValue) GetValue() interface{} { return e.PixelDataInfo }
func (e *pixelDataValue) String() string {
	// TODO: consider adding more sophisticated formatting
	return ""
}
func (s *pixelDataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.PixelDataInfo)
}

func MustGetInts(v Value) []int {
	if v.ValueType() != Ints {
		log.Panicf("MustGetInts expected ValueType of Ints, got: %v", v.ValueType())
	}
	return v.GetValue().([]int)
}

func MustGetStrings(v Value) []string {
	if v.ValueType() != Strings {
		log.Panicf("MustGetStrings expected ValueType of Strings, got: %v", v.ValueType())
	}
	return v.GetValue().([]string)
}

func MustGetBytes(v Value) []byte {
	if v.ValueType() != Bytes {
		log.Panicf("MustGetBytes expected ValueType of Bytes, got: %v", v.ValueType())
	}
	return v.GetValue().([]byte)
}

func MustGetPixelDataInfo(v Value) PixelDataInfo {
	if v.ValueType() != PixelData {
		log.Panicf("MustGetPixelDataInfo expected ValueType of PixelData, got: %v", v.ValueType())
	}
	return v.GetValue().(PixelDataInfo)
}
