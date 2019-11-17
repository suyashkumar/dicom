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

// intToInt64 takes a form of int and returns it as an int64
func intToInt64(i interface{}) (int64, error) {
	switch val := i.(type) {
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	}

	return 0, fmt.Errorf("unexpected value type, expected a form of int")
}

// GetInt retrieves some type of embedded int value (uint32, int16, etc) as an int64
func (e *Element) GetInt() (int64, error) {
	if len(e.Value) != 1 {
		return 0, fmt.Errorf("found %d value(s) in GetInt (expected 1)", len(e.Value))
	}

	return intToInt64(e.Value[0])

}

// GetInts returns a slice of int like values stored in the element.
func (e *Element) GetInts() ([]int64, error) {
	values := make([]int64, len(e.Value))

	for idx, v := range e.Value {
		i, err := intToInt64(v)
		if err != nil {
			return []int64{}, err
		}
		values[idx] = i
	}

	return values, nil
}

// MustGetInt retrieves some type of embedded int value as an int64. Panics if unable to do so.
func (e *Element) MustGetInt() int64 {
	i, err := e.GetInt()
	if err != nil {
		panic("Unable to get int in call to MustGetInt. Error: " + err.Error())
	}
	return i
}

// MustGetInts returns a slice of int like values stored in the element. Panics if unable to do so.
func (e *Element) MustGetInts() []int64 {
	vals, err := e.GetInts()
	if err != nil {
		panic(err)
	}
	return vals
}

// GetString gets a string value from an element.  It returns an error if the
// element contains zero or >1 values, or the value is not a string.
func (e *Element) GetString() (string, error) {
	if len(e.Value) != 1 {
		return "", fmt.Errorf("Found %d value(s) in getstring (expect 1): %v", len(e.Value), e.String())
	}
	v, ok := e.Value[0].(string)
	if !ok {
		return "", fmt.Errorf("string value not found in %v", e)
	}
	return v, nil
}

// MustGetString is similar to GetString(), but panics on error.
// TODO(saito): Add other variants of MustGet<type>.
func (e *Element) MustGetString() string {
	v, err := e.GetString()
	if err != nil {
		panic(err)
	}
	return v
}

// GetStrings returns the list of strings stored in the element. Returns an error
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
		s += fmt.Sprintf(" (#%d)[\n", len(e.Value))
		for _, v := range e.Value {
			s += elementString(v.(*Element), nestLevel+1) + "\n"
		}
		s += indent + " ]"
	} else {
		var sv string
		if len(e.Value) == 1 {
			sv = fmt.Sprintf("%v", e.Value)
		} else {
			sv = fmt.Sprintf("(%d)%v", len(e.Value), e.Value)
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
		if data.Frames[i].Encapsulated {
			csum := sha256.Sum256(data.Frames[i].EncapsulatedData.Data)
			s += fmt.Sprintf("%d:{size:%d, csum:%v, encapsulated:true}, ",
				i, len(data.Frames[i].EncapsulatedData.Data),
				base64.URLEncoding.EncodeToString(csum[:]))
		} else {
			s += fmt.Sprintf("%d:{size:%d, encapsulated: false}, ",
				i, len(data.Frames[i].NativeData.Data))
		}
	}

	return s + "]}"
}

// EndOfData is an pseudoelement to cause the caller to stop reading the input.
var EndOfData = &Element{Tag: dicomtag.Tag{Group: 0x7fff, Element: 0x7fff}}

const VLUndefinedLength uint32 = 0xffffffff

// FindByName finds an element with the given Element.Name in
// "elems" If not found, returns an error.
func FindByName(elems []*Element, name string) (*Element, error) {
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

// FindByTag finds an element with the given Element.Tag in
// "elems" If not found, returns an error.
// TODO: consider a map lookup table perhaps in DataSet
func FindByTag(elems []*Element, tag dicomtag.Tag) (*Element, error) {
	for _, elem := range elems {
		if elem.Tag == tag {
			return elem, nil
		}
	}
	return nil, fmt.Errorf("%s: element not found", dicomtag.DebugString(tag))
}
