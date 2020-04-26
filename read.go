package dicom

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/tag"
)

var ErrorOWRequiresEvenVL = errors.New("vr of OW requires even value length")
var ErrorUnsupportedVR = errors.New("unsupported VR")

func readTag(r dicomio.Reader) (*tag.Tag, error) {
	group, gerr := r.ReadUInt16()
	element, eerr := r.ReadUInt16()

	if gerr == nil && eerr == nil {
		return &tag.Tag{Group: group, Element: element}, nil
	} else {
		return nil, fmt.Errorf("error reading tag: %v %v", gerr, eerr)
	}
}

// TODO: Parsed VR should be an enum. Will require refactors of tag pkg.
func readVR(r dicomio.Reader, isImplicit bool, t tag.Tag) (string, error) {
	if isImplicit {
		if entry, err := tag.Find(t); err == nil {
			return entry.VR, nil
		}
		return tag.UNKNOWN, nil
	}

	// Explicit Transfer Syntax, read 2 byte VR:
	return r.ReadString(2)

}

func readVL(r dicomio.Reader, isImplicit bool, t tag.Tag, vr string) (uint32, error) {
	if isImplicit {
		return r.ReadUInt32()
	}

	// Explicit Transfer Syntax
	// More details here: http://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_7.1.2
	switch vr {
	// TODO: Parsed VR should be an enum. Will require refactors of tag pkg.
	case "NA", "OB", "OD", "OF", "OL", "OW", "SQ", "UN", "UC", "UR", "UT":
		_ = r.Skip(2) // ignore two reserved bytes (0000H)
		vl, err := r.ReadUInt32()
		if err != nil {
			return 0, err
		}

		if vl == tag.VLUndefinedLength && (vr == "UC" || vr == "UR" || vr == "UT") {
			return 0, errors.New("UC, UR and UT may not have an Undefined Length, i.e.,a Value Length of FFFFFFFFH")
		}
		return vl, nil
	default:
		vl16, err := r.ReadUInt16()
		if err != nil {
			return 0, err
		}
		vl := uint32(vl16)
		// Rectify Undefined Length VL
		if vl == 0xffff {
			vl = tag.VLUndefinedLength
		}
		return vl, nil
	}
}

func readValue(r dicomio.Reader, t tag.Tag, vr string, vl uint32, isImplicit bool) (Value, error) {
	// TODO: implement
	vrkind := tag.GetVRKind(t, vr)
	// TODO: if we keep consistent function signature, consider a static map of VR to func?
	switch vrkind {
	case tag.VRBytes:
		return readBytes(r, t, vr, vl)
	case tag.VRString:
		return readString(r, t, vr, vl)
	case tag.VRDate:
		return readDate(r, t, vr, vl)
	case tag.VRUInt16List, tag.VRUInt32List, tag.VRInt16List, tag.VRInt32List:
		return readInt(r, t, vr, vl)
	case tag.VRSequence:
		return readSequence(r, t, vr, vl)
	default:
		return readString(r, t, vr, vl)
	}

	return nil, fmt.Errorf("unsure how to parse this VR")
}

func readSequence(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	var subElements ElementPtrsValue

	if vl == tag.VLUndefinedLength {
		for {
			subElement, err := readElement(r)
			if err != nil {
				// Stop reading due to error
				return nil, err
			}
			if subElement.Tag == tag.SequenceDelimitationItem {
				// Stop reading
				break
			}
			if subElement.Tag != tag.Item {
				// This is an error, should be an Item!
				// TODO: use error var
				return nil, fmt.Errorf("non item found in sequence")
			}
			subElements.value = append(subElements.value, subElement)
		}
	} else {
		// Sequence of elements for a total of VL bytes
		r.PushLimit(int64(vl))
		for !r.IsLimitExhausted() {
			subElem, err := readElement(r)
			if err != nil {
				// TODO: option to ignore errors parsing subelements?
				return nil, err
			}
			subElements.value = append(subElements.value, subElem)
		}
		r.PopLimit()
	}

	return &subElements, nil
}

func readBytes(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	// TODO: add special handling of PixelData
	if vr == "OB" {
		data := make([]byte, vl)
		_, err := r.Read(data)
		return &BytesValue{value: data}, err
	} else if vr == "OW" {
		// OW -> stream of 16 bit words
		if vl%2 != 0 {
			return nil, ErrorOWRequiresEvenVL
		}
		data := make([]byte, vl)
		buf := bytes.NewBuffer(data)
		numWords := int(vl / 2)
		for i := 0; i < numWords; i++ {
			word, err := r.ReadUInt16()
			if err != nil {
				return nil, err
			}
			// TODO: support bytes.BigEndian byte ordering
			err = binary.Write(buf, binary.LittleEndian, word)
			if err != nil {
				return nil, err
			}
		}
		return &BytesValue{value: buf.Bytes()}, nil
	}

	return nil, ErrorUnsupportedVR
}

func readString(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	str, err := r.ReadString(vl)
	// String may have '\0' suffix if its length is odd.
	str = strings.Trim(str, " \000")

	// Split multiple strings
	strs := strings.Split(str, "\\")

	return &StringsValue{value: strs}, err
}

func readDate(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	rawDate, err := r.ReadString(vl)
	if err != nil {
		return nil, err
	}
	date := strings.Trim(rawDate, " \000")

	return &StringsValue{value: []string{date}}, nil

}

func readInt(r dicomio.Reader, t tag.Tag, vr string, vl uint32) (Value, error) {
	// TODO: add other integer types here
	switch vr {
	case "US":
		val, err := r.ReadUInt16()
		return &IntsValue{value: []int{int(val)}}, err
	case "UL":
		val, err := r.ReadUInt32()
		return &IntsValue{value: []int{int(val)}}, err
	}

	return nil, errors.New("could not parse integer type correctly")
}

// readElement reads the next element. If the next element is a sequence element,
// it may result in a collection of Elements.
func readElement(r dicomio.Reader) (*Element, error) {
	t, err := readTag(r)
	if err != nil {
		return nil, err
	}

	vr, err := readVR(r, false, *t)
	if err != nil {
		return nil, err
	}

	vl, err := readVL(r, false, *t, vr)
	if err != nil {
		return nil, err
	}

	val, err := readValue(r, *t, vr, vl, false)

	return &Element{Tag: *t, ValueRepresentation: tag.GetVRKind(*t, vr), ValueLength: vl, Value: val}, nil

}
