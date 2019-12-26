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
var ErrorReadingMagicWord = errors.New("error reading the DICM magic word")

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
		if entry, err := tag.Find(t); err != nil {
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
		vl16, err := r.ReadUInt32()
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

func readRawValue(r dicomio.Reader, vl uint32) ([]byte, error) {
	return r.ReadN(vl)
}

func readValue(r dicomio.Reader, t tag.Tag, vr string, vl uint32, isImplicit bool) (Value, error) {
	// TODO: implement
	// TODO: consolidate VR enums and strings
	vrkind := tag.GetVRKind(t, vr)
	switch vrkind {
	case tag.VRBytes:
		return readBytes(r, t, vr, vl)
	case tag.VRString:
		return readString(r, t, vr, vl)
	case tag.VRDate:
		return readDate(r, t, vr, vl)
	case tag.VRUInt16List, tag.VRUInt32List, tag.VRInt16List, tag.VRInt32List:
		return readInt(r, t, vr, vl)

	}

	return nil, fmt.Errorf("unsure how to parse this VR")
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
	// TODO: use enum values for VR
	switch vr {
	case "US":
		val, err := r.ReadUInt16()
		return &IntsValue{value: []int{int(val)}}, err
	case "UL":
		val, err := r.ReadUInt32()
		return &IntsValue{value: []int{int(val)}}, err
	case "SL":
		val, err := r.ReadInt32()
		return &IntsValue{value: []int{int(val)}}, err
	case "SS":
		val, err := r.ReadInt16()
		return &IntsValue{value: []int{int(val)}}, err
	default:
		return &IntsValue{}, ErrorUnsupportedVR // TODO: consider a better error
	}

}

func readHeaderAndMetadata(r dicomio.Reader) ([]*Element, error) {
	// Assume Little Endian Explicit VR Transfer Syntax
	err := r.Skip(128) // Skip Preamble
	if err != nil {
		return []*Element{}, err // consider returning a more descriptive error
	}

	// Check for the magic word:
	if magic, err := r.ReadString(4); err != nil || magic != "DICM" {
		return []*Element{}, ErrorReadingMagicWord
	}

	// (0002,0000) MetaElementGroupLength
	metaElement, err := readElement(r)

	if err != nil {
		return []*Element{}, err
	}

	if metaElement.Tag != tag.FileMetaInformationGroupLength {
		return []*Element{}, err
	}

	metaElements := []*Element{metaElement}

	lim, ok := metaElement.Value.(*IntsValue)
	if !ok {
		return []*Element{}, errors.New("Issue parsing metaelementgrouplength value")
	}

	// Read the metadata elements
	r.PushLimit(uint(lim.value[0]))
	defer r.PopLimit()

	elem, readErr := readElement(r)
	for readErr == nil {
		metaElements = append(metaElements, elem)
		elem, readErr = readElement(r)
	}

	if readErr != dicomio.ErrorLimitReached {
		return metaElements, readErr
	}

	return metaElements, nil
}

func readElement(r dicomio.Reader) (*Element, error) {
	currTag, err := readTag(r)
	if err != nil {
		return nil, err
	}
	// TODO: add options to stop at tag, drop pixel data, etc here.
	// TODO: support other TransferSyntaxes (other than little endian explicit)

	vr, err := readVR(r, false, *currTag)
	vl, err := readVL(r, false, *currTag, vr)

	val, err := readValue(r, *currTag, vr, vl, false)

	return &Element{Tag: *currTag, RawValueRepresentation: vr, ValueLength: vl, Value: val}, nil

}

func readDICOM(r dicomio.Reader) ([]*Element, error) {
	dataSet, err := readHeaderAndMetadata(r)
	if err != nil {
		return []*Element{}, err
	}
	// TODO: set different transfer syntax here if needed in future, now that we have read metadata elements

	// Read all DICOM elements:
	elem, readErr := readElement(r)
	for readErr == nil {
		if elem.ValueRepresentation == tag.VRSequence {
			if elem.UndefinedLength {
				for {
					subElem, err := readElement(r)
					if err != nil {
						break
					}
					if subElem.Tag == tag.SequenceDelimitationItem {
						break
					}

					if subElem.Tag == tag.Item {
						//TODO: Parse Items

					} else {
						//TODO: validation error?
						return []*Element{}, nil
					}
					dataSet = append(dataSet, subElem)
				}
			} else {
				// TODO: SQ implementation with defined VL
				return dataSet, errors.New("SQ with defined VL not implemented yet")
			}
		}
		dataSet = append(dataSet, elem)
		elem, readErr = readElement(r)
	}

	return dataSet, nil
}
