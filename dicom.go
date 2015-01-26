package dicom

import (
	"encoding/binary"
	"errors"
)

type DicomFile struct {
	NumberOfItems int
	Elements      []DicomElement
	PixelBuffer   interface{}
}

// Errors
var (
	ErrIllegalTag  = errors.New("Illegal tag found in PixelData")
	ErrTagNotFound = errors.New("Could not find tag in dicom dictionary")
	ErrBrokenFile  = errors.New("Invalid DICOM file")
	ErrOddLength   = errors.New("Encountered odd length Value Length")
)

const (
	magic_word                = "DICM"
	implicit_vr_little_endian = "1.2.840.10008.1.2"
	explicit_vr_little_endian = "1.2.840.10008.1.2.1"
	explicit_vr_big_endian    = "1.2.840.10008.1.2.2"
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
	metaElem := buffer.readDataElement(p)
	metaLength := int(metaElem.Value.(uint32))
	file.appendDataElement(metaElem)

	// Read meta tags
	start := buffer.Len()
	for start-buffer.Len() < metaLength {
		elem := buffer.readDataElement(p)
		file.appendDataElement(elem)
	}

	// read endianness and explicit VR
	endianess, implicit, err := file.getTransferSyntax()
	if err != nil {
		return nil, ErrBrokenFile
	}

	// modify buffer according to new TransferSyntaxUID
	buffer.bo = endianess
	buffer.implicit = implicit

	startedPixelData := false

	// Start with image meta data
	for buffer.Len() != 0 {

		elem := buffer.readDataElement(p)
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
			file.PixelBuffer = elem.Value
			elem.Value = "..."
		}

	}

	return file, nil
}

// Append a dataElement to the DicomFile
func (file *DicomFile) appendDataElement(elem *DicomElement) {
	file.Elements = append(file.Elements, *elem)
}

// Finds the SyntaxTrasnferUID and returns the endianess and implicit VR for the file
func (file *DicomFile) getTransferSyntax() (binary.ByteOrder, bool, error) {

	var err error

	elem, err := file.LookupElement("TransferSyntaxUID")
	if err != nil {
		return nil, true, err
	}

	ts := elem.Value.([]string)[0]

	// defaults are explicit VR, little endian
	switch ts {
	case implicit_vr_little_endian:
		return binary.LittleEndian, true, nil
	case explicit_vr_little_endian:
		return binary.LittleEndian, false, nil
	case explicit_vr_big_endian:
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
