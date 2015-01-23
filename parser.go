package dicom

import (
	"fmt"
	"os"
	"strings"
)

// Constants
const (
	pixeldata_group    = 0xFFFE
	unknown_group_name = "Dicom::Unknown"
)

// A DICOM element
type DicomElement struct {
	Group   int
	Element int
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
	fh, err := os.Open("dicom.dic")
	defer fh.Close()

	if err != nil {
		panic(err)
	}

	err = Dictionary(fh)(&p)

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

	// always read (PixelData) implicit
	if elem.Group == pixeldata_group {
		implicit = true
	}

	if implicit {
		vr, vl = buffer.readImplicit(elem, p)
	} else {
		vr, vl = buffer.readExplicit(elem)
	}

	// Double check the value of VL
	if vl == 0xffffffff {
		vl = 0
	}

	elem.Vr = vr
	elem.Vl = vl

	// data
	var data interface{}

	switch vr {
	case "US", "UL":
		if vl == 8 {
			data = buffer.readFloat()
		} else {
			data, _ = buffer.readNumber(vl)
		}
	case "OW":
		data = buffer.readUInt16Array(vl)
	case "OB", "NA":
		data = buffer.readUInt8Array(vl)
	case "OX":
		// TODO: work with the BitsAllocated tag
		data = buffer.readUInt16Array(vl)
	default:
		str := buffer.readString(vl)
		data = strings.Split(str, "\\")
	}

	elem.Value = data

	return elem
}
