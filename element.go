package dicom

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/grailbio/go-dicom/dicomio"
	"github.com/grailbio/go-dicom/dicomtag"
	"v.io/x/lib/vlog"
)

// Constants
const (
	itemSeqGroup     = 0xFFFE
	unknownGroupName = "Unknown Group"
	privateGroupName = "Private Data"
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
		case dicomtag.VRStringList:
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
		case dicomtag.VRTagList:
			_, ok = v.(dicomtag.Tag)
		case dicomtag.VRSequence:
			subelem, ok := v.(*Element)
			if ok {
				ok = (subelem.Tag == dicomtag.Item)
			}
		case dicomtag.VRItem:
			_, ok = v.(*Element)
			// TODO(saito) I'm missing a few legit VR types here.
		}
		if !ok {
			return nil, fmt.Errorf("%v: wrong value type for NewElement: %v", dicomtag.DebugString(tag), v)
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
		vlog.Fatalf("Failed to create element with tag %v: %v", tag, err)
	}
	return elem
}

// GetUInt32 gets a uint32 value from an element.  It returns an error if the
// element contains zero or >1 values, or the value is not a uint32.
func (e *Element) GetUInt32() (uint32, error) {
	if len(e.Value) != 1 {
		return 0, fmt.Errorf("Found %d value(s) in getuint32 (expect 1): %v", len(e.Value), e)
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
		return 0, fmt.Errorf("Found %d value(s) in getuint16 (expect 1): %v", len(e.Value), e)
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
		return "", fmt.Errorf("Found %d value(s) in getstring (expect 1): %v", len(e.Value), e.String())
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
	doassert(nestLevel < 10)
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
		s += indent + "]"
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

// Read an Item object as raw bytes, w/o parsing them into DataElement. Used to
// parse pixel data.
func readRawItem(d *dicomio.Decoder) ([]byte, bool) {
	tag := readTag(d)
	// Item is always encoded implicit. PS3.6 7.5
	vr, vl := readImplicit(d, tag)
	if d.Error() != nil {
		return nil, true
	}
	if tag == dicomtag.SequenceDelimitationItem {
		if vl != 0 {
			d.SetErrorf("SequenceDelimitationItem's VL != 0: %v", vl)
		}
		return nil, true
	}
	if tag != dicomtag.Item {
		d.SetErrorf("Expect Item in pixeldata but found tag %v", dicomtag.DebugString(tag))
		return nil, false
	}
	if vl == undefinedLength {
		d.SetErrorf("Expect defined-length item in pixeldata")
		return nil, false
	}
	if vr != "NA" {
		d.SetErrorf("Expect NA item, but found %s", vr)
		return nil, true
	}
	return d.ReadBytes(int(vl)), false
}

// PixelDataInfo is the Element.Value payload for PixelData element.
type PixelDataInfo struct {
	Offsets []uint32 // BasicOffsetTable
	Frames  [][]byte // Parsed images
}

func (data PixelDataInfo) String() string {
	s := fmt.Sprintf("image{offsets: %v, frames: [", data.Offsets)
	for i := 0; i < len(data.Frames); i++ {
		csum := sha256.Sum256(data.Frames[i])
		s += fmt.Sprintf("%d:{size:%d, csum:%v} ",
			i, len(data.Frames[i]),
			base64.URLEncoding.EncodeToString(csum[:]))
	}
	return s + "]}"
}

// Read the basic offset table. This is the first Item object embedded inside
// PixelData element. P3.5 8.2. P3.5, A4 has a better example.
func readBasicOffsetTable(d *dicomio.Decoder) []uint32 {
	data, endOfData := readRawItem(d)
	if endOfData {
		d.SetErrorf("basic offset table not found")
	}
	if len(data) == 0 {
		return []uint32{0}
	}

	byteOrder, _ := d.TransferSyntax()
	// The payload of the item is sequence of uint32s, each representing the
	// byte size of an image that follows.
	subdecoder := dicomio.NewBytesDecoder(data, byteOrder, dicomio.ImplicitVR)
	var offsets []uint32
	for subdecoder.Len() > 0 && subdecoder.Error() == nil {
		offsets = append(offsets, subdecoder.ReadUInt32())
	}
	return offsets
}

// ParseFileHeader consumes the DICOM magic header and metadata elements (whose
// elements with tag group==2) from a Dicom file. Errors are reported through
// d.Error().
func ParseFileHeader(d *dicomio.Decoder) []*Element {
	d.PushTransferSyntax(binary.LittleEndian, dicomio.ExplicitVR)
	defer d.PopTransferSyntax()
	d.Skip(128) // skip preamble

	// check for magic word
	if s := d.ReadString(4); s != "DICM" {
		d.SetError(errors.New("Keyword 'DICM' not found in the header"))
		return nil
	}

	// (0002,0000) MetaElementGroupLength
	metaElem := ReadElement(d, ReadOptions{})
	if d.Error() != nil {
		return nil
	}
	if metaElem.Tag != dicomtag.FileMetaInformationGroupLength {
		d.SetErrorf("MetaElementGroupLength not found; insteadfound %s", metaElem.Tag.String())
	}
	metaLength, err := metaElem.GetUInt32()
	if err != nil {
		d.SetErrorf("Failed to read uint32 in MetaElementGroupLength: %v", err)
		return nil
	}
	if d.Len() <= 0 {
		d.SetErrorf("No data element found")
		return nil
	}
	metaElems := []*Element{metaElem}

	// Read meta tags
	d.PushLimit(int64(metaLength))
	defer d.PopLimit()
	for d.Len() > 0 {
		elem := ReadElement(d, ReadOptions{})
		if d.Error() != nil {
			break
		}
		metaElems = append(metaElems, elem)
		vlog.VI(2).Infof("Meta elem: %v, len %v", elem.String(), d.Len())
	}
	return metaElems
}

// ReadElement reads one DICOM data element. Errors are reported through
// d.Error(). The caller must check d.Error() before using the returned value.
func ReadElement(d *dicomio.Decoder, options ReadOptions) *Element {
	tag := readTag(d)
	if tag == dicomtag.PixelData && options.DropPixelData {
		return nil
	}
	// The elements for group 0xFFFE should be Encoded as Implicit VR.
	// DICOM Standard 09. PS 3.6 - Section 7.5: "Nesting of Data Sets"

	_, implicit := d.TransferSyntax()
	if tag.Group == itemSeqGroup {
		implicit = dicomio.ImplicitVR
	}
	var vr string // Value Representation
	var vl uint32 // Value Length
	if implicit == dicomio.ImplicitVR {
		vr, vl = readImplicit(d, tag)
	} else {
		doassert(implicit == dicomio.ExplicitVR, implicit)
		vr, vl = readExplicit(d, tag)
	}
	if d.Error() != nil {
		return nil
	}
	elem := &Element{
		Tag:             tag,
		VR:              vr,
		UndefinedLength: (vl == undefinedLength),
	}
	var data []interface{}

	if tag == dicomtag.PixelData {
		// P3.5, A.4 describes the format. Currently we only support an encapsulated image format.
		//
		// PixelData is usually the last element in a DICOM file. When
		// the file stores N images, the elements that follow PixelData
		// are laid out in the following way:
		//
		// Item(BasicOffsetTable) Item(PixelDataInfo0) ... Item(PixelDataInfoM) SequenceDelimiterItem
		//
		// Item(BasicOffsetTable) is an Item element whose payload
		// encodes N uint32 values. Kth uint32 is the bytesize of the
		// Kth image. Item(PixelDataInfo*) are chunked sequences of bytes. I
		// presume that single PixelDataInfo item doesn't cross a image
		// boundary, but the spec isn't clear.
		//
		// The total byte size of Item(PixelDataInfo*) equal the total of
		// the bytesizes found in BasicOffsetTable.
		if vl == undefinedLength {
			var image PixelDataInfo
			image.Offsets = readBasicOffsetTable(d) // TODO(saito) Use the offset table.
			if len(image.Offsets) > 1 {
				vlog.Infof("Warning: multiple images not supported yet. Combining them into a byte sequence: %v", image.Offsets)
			}
			for d.Len() > 0 {
				chunk, endOfItems := readRawItem(d)
				if d.Error() != nil {
					break
				}
				if endOfItems {
					break
				}
				image.Frames = append(image.Frames, chunk)
			}
			data = append(data, image)
		} else {
			vlog.Errorf("Warning: defined-length pixel data not supported: tag %v, VR=%v, VL=%v", tag.String(), vr, vl)
			var image PixelDataInfo
			image.Frames = append(image.Frames, d.ReadBytes(int(vl)))
			data = append(data, image)
		}
		// TODO(saito) handle multi-frame image.
	} else if vr == "SQ" { // Sequence
		if vl == undefinedLength {
			// Format:
			//  Sequence := ItemSet* SequenceDelimitationItem
			//  ItemSet := Item Any* ItemDelimitationItem (when Item.VL is undefined) or
			//             Item Any*N                     (when Item.VL has a defined value)
			for {
				item := ReadElement(d, options)
				if d.Error() != nil {
					break
				}
				if item.Tag == dicomtag.SequenceDelimitationItem {
					break
				}
				if item.Tag != dicomtag.Item {
					d.SetErrorf("Found non-Item element in seq w/ undefined length: %v", dicomtag.DebugString(item.Tag))
					break
				}
				data = append(data, item)
			}
		} else {
			// Format:
			//  Sequence := ItemSet*VL
			// See the above comment for the definition of ItemSet.
			d.PushLimit(int64(vl))
			defer d.PopLimit()
			for d.Len() > 0 {
				item := ReadElement(d, options)
				if d.Error() != nil {
					break
				}
				if item.Tag != dicomtag.Item {
					d.SetErrorf("Found non-Item element in seq w/ undefined length: %v", dicomtag.DebugString(item.Tag))
					break
				}
				data = append(data, item)
			}
		}
	} else if tag == dicomtag.Item { // Item (component of SQ)
		if vl == undefinedLength {
			// Format: Item Any* ItemDelimitationItem
			for {
				subelem := ReadElement(d, options)
				if d.Error() != nil {
					break
				}
				if subelem.Tag == dicomtag.ItemDelimitationItem {
					break
				}
				data = append(data, subelem)
			}
		} else {
			// Sequence of arbitary elements, for the  total of "vl" bytes.
			d.PushLimit(int64(vl))
			for d.Len() > 0 {
				subelem := ReadElement(d, options)
				if d.Error() != nil {
					break
				}
				data = append(data, subelem)
			}
			d.PopLimit()
		}
	} else { // List of scalar
		if vl == undefinedLength {
			d.SetErrorf("Undefined length disallowed for VR=%s, tag %s", vr, dicomtag.DebugString(tag))
			return nil
		}
		d.PushLimit(int64(vl))
		defer d.PopLimit()
		if vr == "DA" {
			// TODO(saito) Maybe we should validate the date.
			date := strings.Trim(d.ReadString(int(vl)), " \000")
			data = []interface{}{date}
		} else if vr == "AT" {
			// (2byte group, 2byte elem)
			for d.Len() > 0 && d.Error() == nil {
				tag := dicomtag.Tag{d.ReadUInt16(), d.ReadUInt16()}
				data = append(data, tag)
			}
		} else if vr == "OW" {
			if vl%2 != 0 {
				d.SetErrorf("%v: OW requires even length, but found %v", dicomtag.DebugString(tag), vl)
			} else {
				n := int(vl / 2)
				e := dicomio.NewBytesEncoder(dicomio.NativeByteOrder, dicomio.UnknownVR)
				for i := 0; i < n; i++ {
					v := d.ReadUInt16()
					e.WriteUInt16(v)
				}
				doassert(e.Error() == nil, e.Error())
				// TODO(saito) Check that size is even. Byte swap??
				// TODO(saito) If OB's length is odd, is VL odd too? Need to check!
				data = append(data, e.Bytes())
			}
		} else if vr == "OB" {
			// TODO(saito) Check that size is even. Byte swap??
			// TODO(saito) If OB's length is odd, is VL odd too? Need to check!
			data = append(data, d.ReadBytes(int(vl)))
		} else if vr == "LT" || vr == "UT" {
			str := d.ReadString(int(vl))
			data = append(data, str)
		} else if vr == "UL" {
			for d.Len() > 0 && d.Error() == nil {
				data = append(data, d.ReadUInt32())
			}
		} else if vr == "SL" {
			for d.Len() > 0 && d.Error() == nil {
				data = append(data, d.ReadInt32())
			}
		} else if vr == "US" {
			for d.Len() > 0 && d.Error() == nil {
				data = append(data, d.ReadUInt16())
			}
		} else if vr == "SS" {
			for d.Len() > 0 && d.Error() == nil {
				data = append(data, d.ReadInt16())
			}
		} else if vr == "FL" {
			for d.Len() > 0 && d.Error() == nil {
				data = append(data, d.ReadFloat32())
			}
		} else if vr == "FD" {
			for d.Len() > 0 && d.Error() == nil {
				data = append(data, d.ReadFloat64())
			}
		} else {
			// List of strings, each delimited by '\\'.
			v := d.ReadString(int(vl))
			// String may have '\0' suffix if its length is odd.
			str := strings.Trim(v, " \000")
			if len(str) > 0 {
				for _, s := range strings.Split(str, "\\") {
					data = append(data, s)
				}
			}
		}
	}
	elem.Value = data
	return elem
}

const undefinedLength uint32 = 0xffffffff

// Read a DICOM data element's tag value ie. (0002,0000) added Value
// Multiplicity PS 3.5 6.4
func readTag(buffer *dicomio.Decoder) dicomtag.Tag {
	group := buffer.ReadUInt16()   // group
	element := buffer.ReadUInt16() // element
	return dicomtag.Tag{group, element}
}

// Read the VR from the DICOM ditionary The VL is a 32-bit unsigned integer
func readImplicit(buffer *dicomio.Decoder, tag dicomtag.Tag) (string, uint32) {
	vr := "UN"
	if entry, err := dicomtag.Find(tag); err == nil {
		vr = entry.VR
	}

	vl := buffer.ReadUInt32()
	// Rectify Undefined Length VL
	if vl == 0xffffffff {
		vl = undefinedLength
	}
	// Error when encountering odd length
	if vl != undefinedLength && vl%2 != 0 {
		buffer.SetErrorf("Encountered odd length (vl=%v) when reading implicit VR '%v' for tag %s", vl, vr, dicomtag.DebugString(tag))
	}
	return vr, vl
}

// The VR is represented by the next two consecutive bytes
// The VL depends on the VR value
func readExplicit(buffer *dicomio.Decoder, tag dicomtag.Tag) (string, uint32) {
	vr := buffer.ReadString(2)
	var vl uint32
	if vr == "US" {
		vl = 2
	}
	// long value representations
	switch vr {
	case "NA", "OB", "OD", "OF", "OL", "OW", "SQ", "UN", "UC", "UR", "UT":
		buffer.Skip(2) // ignore two bytes for "future use" (0000H)
		vl = buffer.ReadUInt32()
		// Rectify Undefined Length VL
		if vl == 0xffffffff {
			switch vr {
			case "UC", "UR", "UT":
				buffer.SetError(errors.New("UC, UR and UT may not have an Undefined Length, i.e.,a Value Length of FFFFFFFFH"))
			default:
				vl = undefinedLength
			}
		}
	default:
		vl = uint32(buffer.ReadUInt16())
		// Rectify Undefined Length VL
		if vl == 0xffff {
			vl = undefinedLength
		}
	}
	if vl != undefinedLength && vl%2 != 0 {
		buffer.SetErrorf("Encountered odd length (vl=%v) when reading explicit VR %v for tag %s", vl, vr, dicomtag.DebugString(tag))
	}
	return vr, vl
}

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
