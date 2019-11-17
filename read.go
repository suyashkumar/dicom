package dicom

import (
	"errors"
	"fmt"

	"github.com/suyashkumar/dicom/legacy/element"
	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func readTag(r dicomio.Reader) (*tag.Tag, error) {
	group, gerr := r.ReadUInt16()
	element, eerr := r.ReadUInt16()

	if gerr != nil && eerr != nil {
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

	// Explicit Transfer Syntax:
	return r.ReadString()

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

		if vl == element.VLUndefinedLength && (vr == "UC" || vr == "UR" || vr == "UT") {
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
			vl = element.VLUndefinedLength
		}
		return vl, nil
	}
}

func readRawValue(r dicomio.Reader, vl uint32) ([]byte, error) {
	return r.ReadN(vl)
}

func readValue(r dicomio.Reader, vr string, vl uint32, isImplicit bool) (Value, bool) {
	// TODO: implement
}

func readElement() {}

func readRawItem(r dicomio.Reader) ([]byte, error) {
	t, err := readTag(r)
	if err != nil {
		return []byte{}, err
	}

	vr, err := readVR(r, false, *t)
	if err != nil {
		return []byte{}, err
	}

	vl, err := readVL(r, false, *t, vr)
	if err != nil {
		return []byte{}, err
	}

}
