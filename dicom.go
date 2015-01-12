package dicom

import (
	"bytes"
	"errors"
	"fmt"
)

type DicomFile struct {
	NumberOfItems int
	Elements      []DicomElement
	PixelBuffer   []uint16
}

// Errors
var (
	ErrIllegalTag = errors.New("Illegal tag found in PixelData")
)

// Parse a byte array, returns a DICOM file struct
func Parse(buff []byte) (*DicomFile, error) {
	buffer := bytes.NewBuffer(buff)

	buffer.Next(128) // skip preamble

	// check for magic word
	if magicWord := string(buffer.Next(4)); magicWord != "DICM" {
		return nil, fmt.Errorf("Invalid DICOM file")
	}

	file := &DicomFile{}

	// (0002,0000) MetaElementGroupLength
	metaElem := readDataElement(buffer, false)
	metaLength := int(metaElem.Value.(uint32))
	file.appendDataElement(metaElem)

	// Read meta tags
	start := buffer.Len()
	for start-buffer.Len() < metaLength {
		elem := readDataElement(buffer, false)
		file.appendDataElement(elem)
	}

	startedPixelData := false

	// Start with image meta data
	for buffer.Len() != 0 {

		elem := readDataElement(buffer, false)
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

// Lookup a tag by name
func (file *DicomFile) lookupElement(name string) (*DicomElement, error) {

	for _, elem := range file.Elements {
		if elem.Name == name {
			return &elem, nil
		}
	}

	return nil, ErrTagNotFound
}
