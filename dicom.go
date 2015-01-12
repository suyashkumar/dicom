package dicom

import (
	"bytes"
	"fmt"
)

type DicomFile struct {
	NumberOfItems int
	Elements      []DicomElement
	PixelBuffer   []byte
}

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
	metaElem := ReadDataElement(buffer, false)
	metaLength := int(metaElem.Value.(uint32))
	appendDataElement(file, metaElem)

	// Read meta tags
	start := buffer.Len()
	for start-buffer.Len() < metaLength {
		elem := ReadDataElement(buffer, false)
		appendDataElement(file, elem)
	}

	return file, nil
}

// Append a dataElement to the DicomFile
func appendDataElement(file *DicomFile, elem *DicomElement) {
	file.Elements = append(file.Elements, *elem)
}
