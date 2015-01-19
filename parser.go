package dicom

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
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
func (p *Parser) readDataElement(buffer *bytes.Buffer, implicit bool) *DicomElement {

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

// Read the VR from the DICOM ditionary
// The VL is a 32-bit unsigned integer
func (p *Parser) readImplicit(buffer *bytes.Buffer, elem *DicomElement) (string, uint32) {
	var vr string
	entry, err := p.getDictEntry(elem.Group, elem.Element)
	if err != nil {
		if err == ErrTagNotFound {
			vr = "UN"
			panic(err)
		}
	}
	vr = entry.vr

	vl := readUInt32(buffer)

	return vr, uint32(vl)
}

// The VR is represented by the next two consecutive bytes
// The VL depends on the VR value
func readExplicit(buffer *bytes.Buffer, elem *DicomElement) (string, uint32) {

	var vl uint32
	vr := string(buffer.Next(2))

	// long value representations
	if vr == "OB" || vr == "OF" || vr == "SQ" || vr == "OW" || vr == "UN" || vr == "UR" || vr == "UT" {
		buffer.Next(2) // ignore two bytes for "future use"
		vl = readUInt32(buffer)
	} else {
		vl = uint32(readUInt16(buffer))
	}

	return vr, vl
}

func readNumber(buffer *bytes.Buffer, size uint32) (uint32, error) {
	chunk := buffer.Next(int(size))

	switch size {
	case 1:
		return uint32(chunk[0]), nil
	case 2:
		return uint32(byteorder.Uint16(chunk)), nil
	case 4:
		return uint32(byteorder.Uint32(chunk)), nil
	}

	return 0, ErrWrongNumberSize
}

// Read x consecutive bytes as a string
func readString(buffer *bytes.Buffer, size uint32) string {
	chunk := buffer.Next(int(size))
	return string(chunk)
}

// Read 8 consecutive bytes as a float32
func readFloat(buffer *bytes.Buffer) float32 {
	chunk := buffer.Next(8)
	b := byteorder.Uint32(chunk)

	return math.Float32frombits(b)
}

// Read a DICOM data element's tag value
// ie. (0002,0000)
func (p *Parser) readTag(buffer *bytes.Buffer) *DicomElement {
	group := readHex(buffer)   // group
	element := readHex(buffer) // element

	var name string
	entry, err := p.getDictEntry(group, element)
	if err != nil {
		if err == ErrTagNotFound {
			name = unknown_group_name
			panic(err)
		}
	}
	name = entry.name

	return &DicomElement{
		Group:   group,
		Element: element,
		Name:    name,
	}
}

// Read 2 bytes as a hexadecimal value
func readHex(buffer *bytes.Buffer) int {
	val := readUInt16(buffer)
	return int(val)
}

// Read 4 bytes as an UInt32
func readUInt32(buffer *bytes.Buffer) uint32 {
	chunk := buffer.Next(4)
	return byteorder.Uint32(chunk)
}

// Read 2 bytes as an UInt16
func readUInt16(buffer *bytes.Buffer) uint16 {
	chunk := buffer.Next(2)        // 2-bytes chunk
	return byteorder.Uint16(chunk) // read as uint16
}

// Read x number of bytes as an array of UInt16 values
func readUInt16Array(buffer *bytes.Buffer, vl uint32) []uint16 {
	slice := make([]uint16, int(vl)/2)

	for i := 0; i < len(slice)-1; i++ {
		slice[i] = readUInt16(buffer)
	}

	return slice
}

// Read x number of bytes as an array of UInt8 values
func readUInt8Array(buffer *bytes.Buffer, vl uint32) []byte {
	chunk := buffer.Next(int(vl))
	return chunk
}
