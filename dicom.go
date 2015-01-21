package dicom

import (
	"encoding/binary"
	"errors"
)

type DicomFile struct {
	NumberOfItems int
	Elements      []DicomElement
	PixelBuffer   []uint16
}

// Errors
var (
	ErrIllegalTag      = errors.New("Illegal tag found in PixelData")
	ErrTagNotFound     = errors.New("Could not find tag in dicom dictionary")
	ErrWrongNumberSize = errors.New("Not a valid byte size for readNumber")
	ErrBrokenFile      = errors.New("Invalid DICOM file")
)

const (
	magic_word = "DICM"
)

// Parse a byte array, returns a DICOM file struct
func (p *Parser) Parse(buff []byte) (*DicomFile, error) {
	buffer := newDicomBuffer(buff)

	buffer.Next(128) // skip preamble

	// check for magic word
	if magicWord := string(buffer.Next(4)); magicWord != magic_word {
		return nil, ErrBrokenFile
	}

	file := &DicomFile{}

	// (0002,0000) MetaElementGroupLength
	metaElem := p.readDataElement(buffer, false)
	metaLength := int(metaElem.Value.(uint32))
	file.appendDataElement(metaElem)

	// Read meta tags
	start := buffer.Len()
	for start-buffer.Len() < metaLength {
		elem := p.readDataElement(buffer, false)
		file.appendDataElement(elem)
	}

	// read endianness and explicit VR
	endianess, implicit, err := file.getTransferSyntax()
	if err != nil {
		return nil, ErrBrokenFile
	}

	buffer.bo = endianess

	startedPixelData := false

	// Start with image meta data
	for buffer.Len() != 0 {

		elem := p.readDataElement(buffer, implicit)
		name := elem.Name
		file.appendDataElement(elem)

		if startedPixelData == true {

			// TODO: refactor this in separate function
			if name == "Item" {
				if len(elem.Value.([]byte)) == 4 {
					break // Skip Basic Offset Table
				} else {
					// TODO: concat multiple pixel data images
					elem.Value = "..."
				}
			} else if name == "SequenceDelimitationItem" {
				startedPixelData = false
			} else {
				panic(ErrIllegalTag)
			}
		}

		if name == "PixelData" {
			startedPixelData = true
			file.PixelBuffer = elem.Value.([]uint16)
			elem.Value = "..."
		}

	}

	return file, nil
}

// Append a dataElement to the DicomFile
func (file *DicomFile) appendDataElement(elem *DicomElement) {
	file.Elements = append(file.Elements, *elem)
}

// Finds the SyntaxTrasnferUID and returns the endianess and byteorder for the file
func (file *DicomFile) getTransferSyntax() (binary.ByteOrder, bool, error) {

	var err error

	elem, err := file.LookupElement("TransferSyntaxUID")
	if err != nil {
		return nil, true, err
	}

	ts := elem.Value.([]string)[0]

	// defaults are explicit VR, little endian
	switch ts {
	case "1.2.840.10008.1.2":
		return binary.LittleEndian, true, nil
	case "1.2.840.10008.1.2.1":
		return binary.LittleEndian, false, nil
	case "1.2.840.10008.1.2.2":
		return binary.BigEndian, false, nil
	}

	return binary.LittleEndian, false, nil

}

// Lookup a tag by name
func (file *DicomFile) LookupElement(name string) (*DicomElement, error) {

	for _, elem := range file.Elements {
		if elem.Name == name {
			return &elem, nil
		}
	}

	return nil, ErrTagNotFound
}
