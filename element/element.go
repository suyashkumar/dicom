package element

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/suyashkumar/dicom/dicomtag"
	"github.com/suyashkumar/dicom/frame"
)

// Constants
const (
	UnknownGroupName = "Unknown Group"
	PrivateGroupName = "Private Data"
)

// Element represents a single DICOM element. Use NewElement() to create a
// element denovo. Avoid creating a struct manually, because setting the VR
// field is a bit tricky.
type Element struct {
	// Tag is a pair of <group, element>. See tags.go for possible values.
	Tag dicomtag.Tag

	// List of values in the element. Their types depends on value
	// representation (VR) of the Tag; Cf. tag.go.
	//
	// If Tag==TagPixelData, len(Value)==1, and Value[0] is PixelDataInfo.
	// Else if Tag==TagItem, each Value[i] is a *Element.
	//    a value's Tag can be any (including TagItem, which represents a nested Item)
	// Else if VR=="SQ", Value[i] is a *Element, with Tag=TagItem.
	// Else if VR=="OW", "OB", then len(Value)==1, and Value[0] is []byte.
	// Else if VR=="LT", or "UT", then len(Value)==1, and Value[0] is string
	// Else if VR=="DA", then len(Value)==1, and Value[0] is string. Use ParseDate() to parse the date string.
	// Else if VR=="US", Value[] is a list of uint16s
	// Else if VR=="UL", Value[] is a list of uint32s
	// Else if VR=="SS", Value[] is a list of int16s
	// Else if VR=="SL", Value[] is a list of int32s
	// Else if VR=="FL", Value[] is a list of float32s
	// Else if VR=="FD", Value[] is a list of float64s
	// Else if VR=="AT", Value[] is a list of Tag's.
	// Else, Value[] is a list of strings.
	//
	// Note: Use GetVRKind() to map VR string to the go representation of
	// VR.
	Value []interface{} // Value Multiplicity PS 3.5 6.4

	// Note: the following fields are not interesting to most people, but
	// are filled for completeness.  You can ignore them.

	// VR defines the encoding of Value[] in two-letter alphabets, e.g.,
	// "AE", "UL". See P3.5 6.2. This field MUST be set.
	//
	// dicom.ReadElement() will fill this field with the VR of the tag,
	// either read from input stream (for explicit repl), or from the dicom
	// tag table (for implicit decl). This field need not be set in
	// WriteElement().
	//
	// Note: In a conformant DICOM file, the VR value of an element is
	// determined by its Tag, so this field is redundant.  This field is
	// still required because a non-conformant file with with explicitVR
	// encoding may have an element with VR that's different from the
	// standard's. In such case, this library honors the VR value found in
	// the file, and this field stores the VR used for parsing Values[].
	VR string

	// UndefinedLength is true if, in the DICOM file, the element is encoded
	// as having undefined length, and is delimited by end-sequence or
	// end-item element.  This flag is meaningful only if VR=="SQ" or
	// VR=="NA". Feel free to ignore this field if you don't understand what
	// this means.  It's one of the pointless complexities in the DICOM
	// standard.
	UndefinedLength bool
}

// NewElement creates a new Element with the given tag and values. The type of
// each each value must match the VR (value representation) of the tag (see
// tag_definition.go).
func NewElement(tag dicomtag.Tag, values ...interface{}) (*Element, error) {
	ti, err := dicomtag.Find(tag)
	if err != nil {
		return nil, err
	}
	e := Element{
		Tag:   tag,
		VR:    ti.VR,
		Value: make([]interface{}, len(values)),
	}
	vrKind := dicomtag.GetVRKind(tag, ti.VR)
	for i, v := range values {
		var ok bool
		switch vrKind {
		case dicomtag.VRStringList, dicomtag.VRDate:
			_, ok = v.(string)
		case dicomtag.VRBytes:
			_, ok = v.([]byte)
		case dicomtag.VRUInt16List:
			_, ok = v.(uint16)
		case dicomtag.VRUInt32List:
			_, ok = v.(uint32)
		case dicomtag.VRInt16List:
			_, ok = v.(int16)
		case dicomtag.VRInt32List:
			_, ok = v.(int32)
		case dicomtag.VRFloat32List:
			_, ok = v.(float32)
		case dicomtag.VRFloat64List:
			_, ok = v.(float64)
		case dicomtag.VRPixelData:
			_, ok = v.(PixelDataInfo)
		case dicomtag.VRTagList:
			_, ok = v.(dicomtag.Tag)
		case dicomtag.VRSequence:
			var subelem *Element
			subelem, ok = v.(*Element)
			if ok {
				ok = (subelem.Tag == dicomtag.Item)
			}
		case dicomtag.VRItem:
			_, ok = v.(*Element)
		}
		if !ok {
			return nil, fmt.Errorf("%v: wrong payload type for NewElement: expect %v, but found %v", dicomtag.DebugString(tag), vrKind, v)
		}
		e.Value[i] = v
	}
	return &e, nil
}

// MustNewElement is similar to NewElement, but it crashes the process on any
// error.
func MustNewElement(tag dicomtag.Tag, values ...interface{}) *Element {
	elem, err := NewElement(tag, values...)
	if err != nil {
		panic(fmt.Sprintf("Failed to create element with tag %v: %v", tag, err))
	}
	return elem
}

// GetUInt32 gets a uint32 value from an element.  It returns an error if the
// element contains zero or >1 values, or the value is not a uint32.
func (e *Element) GetUInt32() (uint32, error) {
	if len(e.Value) != 1 {
		return 0, fmt.Errorf("Found %decoder value(s) in getuint32 (expect 1): %v", len(e.Value), e)
	}
	v, ok := e.Value[0].(uint32)
	if !ok {
		return 0, fmt.Errorf("Uint32 value not found in %v", e)
	}
	return v, nil
}

// MustGetUInt32 is similar to GetUInt32, but panics on error.
func (e *Element) MustGetUInt32() uint32 {
	v, err := e.GetUInt32()
	if err != nil {
		panic(err)
	}
	return v
}

// GetUInt16 gets a uint16 value from an element.  It returns an error if the
// element contains zero or >1 values, or the value is not a uint16.
func (e *Element) GetUInt16() (uint16, error) {
	if len(e.Value) != 1 {
		return 0, fmt.Errorf("Found %decoder value(s) in getuint16 (expect 1): %v", len(e.Value), e)
	}
	v, ok := e.Value[0].(uint16)
	if !ok {
		return 0, fmt.Errorf("Uint16 value not found in %v", e)
	}
	return v, nil
}

// MustGetUInt16 is similar to GetUInt16, but panics on error.
func (e *Element) MustGetUInt16() uint16 {
	v, err := e.GetUInt16()
	if err != nil {
		panic(err)
	}
	return v
}

// GetString gets a string value from an element.  It returns an error if the
// element contains zero or >1 values, or the value is not a string.
func (e *Element) GetString() (string, error) {
	if len(e.Value) != 1 {
		return "", fmt.Errorf("Found %decoder value(s) in getstring (expect 1): %v", len(e.Value), e.String())
	}
	v, ok := e.Value[0].(string)
	if !ok {
		return "", fmt.Errorf("string value not found in %v", e)
	}
	return v, nil
}

// MustGetString is similar to GetString(), but panics on error.
//
// TODO(saito): Add other variants of MustGet<type>.
func (e *Element) MustGetString() string {
	v, err := e.GetString()
	if err != nil {
		panic(err)
	}
	return v
}

// GetStrings returns the list of strings stored in the elment. Returns an error
// if the VR of e.Tag is not a string.
func (e *Element) GetStrings() ([]string, error) {
	values := make([]string, 0, len(e.Value))
	for _, v := range e.Value {
		v, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("string value not found in %v", e.String())
		}
		values = append(values, v)
	}
	return values, nil
}

// MustGetStrings is similar to GetStrings, but crashes the process on error.
func (e *Element) MustGetStrings() []string {
	values, err := e.GetStrings()
	if err != nil {
		panic(err)
	}
	return values
}

// GetUint32s returns the list of uint32 values stored in the elment. Returns an
// error if the VR of e.Tag is not a uint32.
func (e *Element) GetUint32s() ([]uint32, error) {
	values := make([]uint32, 0, len(e.Value))
	for _, v := range e.Value {
		v, ok := v.(uint32)
		if !ok {
			return nil, fmt.Errorf("uint32 value not found in %v", e.String())
		}
		values = append(values, v)
	}
	return values, nil
}

// MustGetUint32s is similar to GetUint32s, but crashes the process on error.
func (e *Element) MustGetUint32s() []uint32 {
	values, err := e.GetUint32s()
	if err != nil {
		panic(err)
	}
	return values
}

// GetUint16s returns the list of uint16 values stored in the elment. Returns an
// error if the VR of e.Tag is not a uint16.
func (e *Element) GetUint16s() ([]uint16, error) {
	values := make([]uint16, 0, len(e.Value))
	for _, v := range e.Value {
		v, ok := v.(uint16)
		if !ok {
			return nil, fmt.Errorf("uint16 value not found in %v", e.String())
		}
		values = append(values, v)
	}
	return values, nil
}

// MustGetUint16s is similar to GetUint16s, but crashes the process on error.
func (e *Element) MustGetUint16s() []uint16 {
	values, err := e.GetUint16s()
	if err != nil {
		panic(err)
	}
	return values
}

func elementString(e *Element, nestLevel int) string {
	// doassert(nestLevel < 10)
	indent := strings.Repeat(" ", nestLevel)
	s := indent
	sVl := ""
	if e.UndefinedLength {
		sVl = "u"
	}
	s = fmt.Sprintf("%s %s %s %s ", s, dicomtag.DebugString(e.Tag), e.VR, sVl)
	if e.VR == "SQ" || e.Tag == dicomtag.Item {
		s += fmt.Sprintf(" (#%decoder)[\n", len(e.Value))
		for _, v := range e.Value {
			s += elementString(v.(*Element), nestLevel+1) + "\n"
		}
		s += indent + " ]"
	} else {
		var sv string
		if len(e.Value) == 1 {
			sv = fmt.Sprintf("%v", e.Value)
		} else {
			sv = fmt.Sprintf("(%decoder)%v", len(e.Value), e.Value)
		}
		if len(sv) > 1024 {
			sv = sv[1:1024] + "(...)"
		}
		s += sv
	}
	return s
}

// Stringer
func (e *Element) String() string {
	return elementString(e, 0)
}

// PixelDataInfo is the Element.Value payload for PixelData element.
type PixelDataInfo struct {
	Offsets        []uint32      // BasicOffsetTable
	IsEncapsulated bool          // is the data encapsulated/jpeg encoded?
	Frames         []frame.Frame // Frames
}

func (data PixelDataInfo) String() string {
	s := fmt.Sprintf("image{offsets: %v, frames: [", data.Offsets)
	for i := 0; i < len(data.Frames); i++ {
		if data.Frames[i].IsEncapsulated {
			csum := sha256.Sum256(data.Frames[i].EncapsulatedData.Data)
			s += fmt.Sprintf("%decoder:{size:%decoder, csum:%v, encapsulated:true}, ",
				i, len(data.Frames[i].EncapsulatedData.Data),
				base64.URLEncoding.EncodeToString(csum[:]))
		} else {
			s += fmt.Sprintf("%decoder:{size:%decoder, encapsulated: false}, ",
				i, len(data.Frames[i].NativeData.Data))
		}
	}

	return s + "]}"
}

// EndOfData is an pseudoelement to cause the caller to stop reading the input.
var EndOfData = &Element{Tag: dicomtag.Tag{Group: 0x7fff, Element: 0x7fff}}

const VLUndefinedLength uint32 = 0xffffffff

// FindElementByName finds an element with the given Element.Name in
// "elems" If not found, returns an error.
func FindElementByName(elems []*Element, name string) (*Element, error) {
	t, err := dicomtag.FindByName(name)
	if err != nil {
		return nil, err
	}
	for _, elem := range elems {
		if elem.Tag == t.Tag {
			return elem, nil
		}
	}
	return nil, fmt.Errorf("Could not find element named '%s' in dicom file", name)
}

// FindElementByTag finds an element with the given Element.Tag in
// "elems" If not found, returns an error.
func FindElementByTag(elems []*Element, tag dicomtag.Tag) (*Element, error) {
	for _, elem := range elems {
		if elem.Tag == tag {
			return elem, nil
		}
	}
	return nil, fmt.Errorf("%s: element not found", dicomtag.DebugString(tag))
}
