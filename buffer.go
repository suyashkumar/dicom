package dicom

import (
	"bytes"
	"encoding/binary"
	"math"
)

type dicomBuffer struct {
	*bytes.Buffer
	bo binary.ByteOrder
}

func newDicomBuffer(b []byte) *dicomBuffer {
	return &dicomBuffer{
		bytes.NewBuffer(b),
		binary.LittleEndian,
	}
}

// Read the VR from the DICOM ditionary
// The VL is a 32-bit unsigned integer
func (p *Parser) readImplicit(buffer *dicomBuffer, elem *DicomElement) (string, uint32) {
	var vr string
	entry, err := p.getDictEntry(elem.Group, elem.Element)
	if err != nil {
		if err == ErrTagNotFound {
			vr = "UN"
			panic(err)
		}
	}
	vr = entry.vr

	vl := buffer.readUInt32()

	return vr, uint32(vl)
}

// The VR is represented by the next two consecutive bytes
// The VL depends on the VR value
func (p *Parser) readExplicit(buffer *dicomBuffer, elem *DicomElement) (string, uint32) {

	var vl uint32
	vr := string(buffer.Next(2))

	// long value representations
	if vr == "OB" || vr == "OF" || vr == "SQ" || vr == "OW" || vr == "UN" || vr == "UR" || vr == "UT" {
		buffer.Next(2) // ignore two bytes for "future use"
		vl = buffer.readUInt32()
	} else {
		vl = uint32(buffer.readUInt16())
	}

	return vr, vl
}

func (buffer *dicomBuffer) readNumber(vl uint32) (uint32, error) {
	chunk := buffer.Next(int(vl))

	switch vl {
	case 1:
		return uint32(chunk[0]), nil
	case 2:
		return uint32(buffer.bo.Uint16(chunk)), nil
	case 4:
		return uint32(buffer.bo.Uint32(chunk)), nil
	}

	return 0, ErrWrongNumberSize
}

// Read x consecutive bytes as a string
func (buffer *dicomBuffer) readString(vl uint32) string {
	chunk := buffer.Next(int(vl))
	return string(chunk)
}

// Read 8 consecutive bytes as a float32
func (buffer *dicomBuffer) readFloat() float32 {
	chunk := buffer.Next(8)
	b := buffer.bo.Uint32(chunk)

	return math.Float32frombits(b)
}

// Read a DICOM data element's tag value
// ie. (0002,0000)
func (p *Parser) readTag(buffer *dicomBuffer) *DicomElement {
	group := buffer.readHex()   // group
	element := buffer.readHex() // element

	var name string
	entry, err := p.getDictEntry(group, element)
	if err != nil {
		name = unknown_group_name
	} else {
		name = entry.name
	}

	return &DicomElement{
		Group:   group,
		Element: element,
		Name:    name,
	}
}

// Read 2 bytes as a hexadecimal value
func (buffer *dicomBuffer) readHex() int {
	val := buffer.readUInt16()
	return int(val)
}

// Read 4 bytes as an UInt32
func (buffer *dicomBuffer) readUInt32() uint32 {
	chunk := buffer.Next(4)
	return buffer.bo.Uint32(chunk)
}

// Read 2 bytes as an UInt16
func (buffer *dicomBuffer) readUInt16() uint16 {
	chunk := buffer.Next(2)        // 2-bytes chunk
	return buffer.bo.Uint16(chunk) // read as uint16
}

// Read x number of bytes as an array of UInt16 values
func (buffer *dicomBuffer) readUInt16Array(vl uint32) []uint16 {
	slice := make([]uint16, int(vl)/2)

	for i := 0; i < len(slice)-1; i++ {
		slice[i] = buffer.readUInt16()
	}

	return slice
}

// Read x number of bytes as an array of UInt8 values
func (buffer *dicomBuffer) readUInt8Array(vl uint32) []byte {
	chunk := buffer.Next(int(vl))
	return chunk
}
