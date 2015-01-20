package dicom

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

// Constants
const (
	pixeldata_group    = 0xFFFE
	unknown_group_name = "Dicom::Unknown"
)

var (
	byteorder = binary.LittleEndian
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
	return fmt.Sprintf("(%4x,%4x)", e.Group, e.Element)
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
func (p *Parser) readDataElement(buffer *DicomBuffer, implicit bool) *DicomElement {

	elem := p.readTag(buffer)

	var vr string     // Value Representation
	var vl uint32 = 0 // Value Length

	// always read (PixelData) implicit
	if elem.Group == pixeldata_group {
		implicit = true
	}

	if implicit {
		vr, vl = p.readImplicit(buffer, elem)
	} else {
		vr, vl = readExplicit(buffer, elem)
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
			data = readFloat(buffer)
		} else {
			data, _ = readNumber(buffer, vl)
		}
	case "OW":
		data = readUInt16Array(buffer, vl)
	case "OB", "NA":
		data = readUInt8Array(buffer, vl)
	case "OX":
		// TODO: work with the BitsAllocated tag
		data = readUInt16Array(buffer, vl)
	default:
		str := readString(buffer, vl)
		data = strings.Split(str, "\\")
	}

	elem.Value = data

	return elem
}
