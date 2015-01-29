package dicom

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type dicomBuffer struct {
	*bytes.Buffer
	bo       binary.ByteOrder
	implicit bool
}

// The default DicomBuffer reads a buffer with Little Endian byteorder
// and explicit VR
func newDicomBuffer(b []byte) *dicomBuffer {
	return &dicomBuffer{
		bytes.NewBuffer(b),
		binary.LittleEndian,
		false,
	}
}

// Read the VR from the DICOM ditionary
// The VL is a 32-bit unsigned integer
func (buffer *dicomBuffer) readImplicit(elem *DicomElement, p *Parser) (string, uint32) {

	var vr string

	entry, err := p.getDictEntry(elem.Group, elem.Element)
	if err != nil {
		vr = "UN"
	} else {
		vr = entry.vr
	}

	vl, err := decodeValueLength(buffer, vr)
	if err == ErrOddLength {
		fmt.Printf("WARN: attempted to read odd length VL for %+v\n", elem)
	}

	return vr, vl
}

// The VR is represented by the next two consecutive bytes
// The VL depends on the VR value
func (buffer *dicomBuffer) readExplicit(elem *DicomElement) (string, uint32) {
	vr := string(buffer.Next(2))
	vl, err := decodeValueLength(buffer, vr)
	if err == ErrOddLength {
		fmt.Printf("WARN: attempted to read odd length VL for %+v\n", elem)
	}

	return vr, vl
}

func decodeValueLength(buffer *dicomBuffer, vr string) (uint32, error) {

	var vl uint32

	// long value representations
	switch vr {
	case "OB", "OF", "SQ", "OW", "UN", "UT":
		buffer.Next(2) // ignore two bytes for "future use" (0000H)
		vl = buffer.readUInt32()
	default:
		vl = uint32(buffer.readUInt16())
	}

	// Rectify Undefined Length VL
	if vl == 0xffffffff {
		vl = 0
	}

	// Error when encountering odd length
	if vl > 0 && vl%2 != 0 {
		return 0, ErrOddLength
	}

	return vl, nil
}

// Read a DICOM data element's tag value
// ie. (0002,0000)
func (buffer *dicomBuffer) readTag(p *Parser) *DicomElement {
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

// Read x consecutive bytes as a string
func (buffer *dicomBuffer) readString(vl uint32) string {
	chunk := buffer.Next(int(vl))
	chunk = bytes.Trim(chunk, "\x00") // trim those pesky null bytes
	return string(chunk)
}

// Read 4 consecutive bytes as a float32
func (buffer *dicomBuffer) readFloat() (val float32) {
	buf := bytes.NewBuffer(buffer.Next(4))
	binary.Read(buf, buffer.bo, &val)
	return
}

// Read 8 consecutive bytes as a float64
func (buffer *dicomBuffer) readFloat64() (val float64) {
	buf := bytes.NewBuffer(buffer.Next(8))
	binary.Read(buf, buffer.bo, &val)
	return
}

// Read 2 bytes as a hexadecimal value
func (buffer *dicomBuffer) readHex() uint16 {
	return buffer.readUInt16()
}

// Read 4 bytes as an UInt32
func (buffer *dicomBuffer) readUInt32() (val uint32) {
	buf := bytes.NewBuffer(buffer.Next(4))
	binary.Read(buf, buffer.bo, &val)
	return
}

// Read 4 bytes as an int32
func (buffer *dicomBuffer) readInt32() (val int32) {
	buf := bytes.NewBuffer(buffer.Next(4))
	binary.Read(buf, buffer.bo, &val)
	return
}

// Read 2 bytes as an UInt16
func (buffer *dicomBuffer) readUInt16() (val uint16) {
	buf := bytes.NewBuffer(buffer.Next(2))
	binary.Read(buf, buffer.bo, &val)
	return
}

// Read 2 bytes as an int16
func (buffer *dicomBuffer) readInt16() (val int16) {
	buf := bytes.NewBuffer(buffer.Next(2))
	binary.Read(buf, buffer.bo, &val)
	return
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
