package dicom

import (
	"fmt"
	"strings"
)

// Constants
const (
	pixeldata_group    = 0xFFFE
	unknown_group_name = "Dicom::Unknown"
)

// A DICOM element
type DicomElement struct {
	Group   uint16
	Element uint16
	Name    string
	Vr      string
	Vl      uint32
	Value   interface{}
}

type Parser struct {
	dictionary [][]*dictEntry
}

// Return the tag as a string to use in the Dicom dictionary
func (e *DicomElement) getTag() string {
	return fmt.Sprintf("(%04X,%04X)", e.Group, e.Element)
}

// Create a new parser, with functional options for configuration
// http://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
func NewParser(options ...func(*Parser) error) (*Parser, error) {

	p := Parser{}

	// apply defaults
	err := Dictionary([]byte(dicomDictData))(&p)

	if err != nil {
		panic(err)
	}

	// override defaults
	for _, option := range options {
		err := option(&p)
		if err != nil {
			panic(err)
		}
	}

	return &p, nil
}

// Read a DICOM data element
func (buffer *dicomBuffer) readDataElement(p *Parser) *DicomElement {

	implicit := buffer.implicit
	elem := buffer.readTag(p)

	var vr string     // Value Representation
	var vl uint32 = 0 // Value Length

	// The elements for group 0xFFFE should be Encoded as Implicit VR.
	// DICOM Standard 09. PS 3.6 - Section 7.5: "Nesting of Data Sets"
	if elem.Group == pixeldata_group {
		implicit = true
	}

	if implicit {
		vr, vl = buffer.readImplicit(elem, p)
	} else {
		vr, vl = buffer.readExplicit(elem)
	}

	elem.Vr = vr
	elem.Vl = vl

	// data
	var data interface{}

	switch vr {
	case "AT":
		data = buffer.readUInt16Array(4)
	case "UL":
		data = buffer.readUInt32()
	case "SL":
		data = buffer.readInt32()
	case "US":
		data = buffer.readUInt16()
	case "SS":
		data = buffer.readInt16()
	case "FL":
		data = buffer.readFloat()
	case "FD":
		data = buffer.readFloat64()
	case "OW":
		data = buffer.readUInt16Array(vl)
	case "OB", "NA":
		data = buffer.Next(int(vl))
	case "OX":
		// TODO: work with the BitsAllocated tag
		data = buffer.readUInt16Array(vl)
	case "SQ":
		// TODO: imlement sequence read
		_ = buffer.Next(int(vl))
		data = "..."
	default:
		str := buffer.readString(vl)
		data = strings.Split(str, "\\")
	}

	elem.Value = data

	return elem
}
