package dicom

import (
<<<<<<< HEAD
	"encoding/binary"
	"errors"
)

type DicomFile struct {
	Elements []DicomElement
}

// Errors
var (
	ErrIllegalTag            = errors.New("Illegal tag found in PixelData")
	ErrTagNotFound           = errors.New("Could not find tag in dicom dictionary")
	ErrBrokenFile            = errors.New("Invalid DICOM file")
	ErrOddLength             = errors.New("Encountered odd length Value Length")
	ErrUndefLengthNotAllowed = errors.New("UC, UR and UT may not have an Undefined Length, i.e.,a Value Length of FFFFFFFFH.")
)

const (
	magic_word                = "DICM"
	implicit_vr_little_endian = "1.2.840.10008.1.2"
	explicit_vr_little_endian = "1.2.840.10008.1.2.1"
	explicit_vr_big_endian    = "1.2.840.10008.1.2.2"
)

// Parse a byte array, returns a DICOM file struct
func (p *Parser) Parse(buff []byte) (*DicomFile, <-chan DicomMessage) {

	c := make(chan DicomMessage)
	waitMsg := make(chan bool)

	buffer := newDicomBuffer(buff) //*di.Bytes)

	buffer.Next(128) // skip preamble
	buffer.p = +128

	// check for magic word
	if magicWord := string(buffer.Next(4)); magicWord != magic_word {
		panic(ErrBrokenFile)
	}
	file := &DicomFile{}

	go func() {

		// (0002,0000) MetaElementGroupLength
		metaElem := buffer.readDataElement(p)
		metaLength := int(metaElem.Value[0].(uint32))
		p.appendDataElement(file, metaElem)

		// Read meta tags
		start := buffer.Len()
		for start-buffer.Len() < metaLength {
			elem := buffer.readDataElement(p)
			p.appendDataElement(file, elem)
			c <- DicomMessage{elem, waitMsg}
			<-waitMsg
		}

		// read endianness and explicit VR
		endianess, implicit, err := file.getTransferSyntax()
		if err != nil {
			panic(ErrBrokenFile)
		}

		// modify buffer according to new TransferSyntaxUID
		buffer.bo = endianess
		buffer.implicit = implicit

		// Start with image meta data
		for buffer.Len() != 0 {

			elem := buffer.readDataElement(p)
			p.appendDataElement(file, elem)
			c <- DicomMessage{elem, waitMsg}
			<-waitMsg

			if elem.Vr == "SQ" {
				p.readItems(file, buffer, elem, c)
			}

			if elem.Name == "PixelData" {
				p.readPixelItems(file, buffer, elem, c)
				break
			}

		}

		close(c)
	}()

	return file, c
}

func (p *Parser) readItems(file *DicomFile, buffer *dicomBuffer, sq *DicomElement, c chan DicomMessage) (uint32, error) {
	waitMsg := make(chan bool)

	sq.IndentLevel++
	sqLength := sq.Vl

	if sqLength == 0 {
		return 0, nil
	}

	elem := buffer.readDataElement(p)
	elem.IndentLevel = sq.IndentLevel

	sqAcum := elem.elemLen
	itemLength := elem.Vl
	itemAcum := uint32(0)

	if elem.Name == "Item" {
		if elem.Vl > 0 {
			for buffer.Len() != 0 {

				p.appendDataElement(file, elem)
				c <- DicomMessage{elem, waitMsg}
				<-waitMsg

				if elem.Vr == "SQ" {
					l, _ := p.readItems(file, buffer, elem, c)
					sqAcum += l
				}

				if itemAcum == itemLength {
					break
				}

				if sqAcum == sqLength {
					break
				}

				elem = buffer.readDataElement(p)
				elem.IndentLevel = sq.IndentLevel
				if elem.Name == "Item" {
					itemLength = elem.Vl
				}
				itemAcum += elem.elemLen
				sqAcum += elem.elemLen

			}

		} else if elem.undefLen == true {
			//log.Println("____ ITEM UNDEF LEN ____")
			for buffer.Len() != 0 {

				if elem.Vr == "SQ" {
					p.readItems(file, buffer, elem, c)
				}

				if elem.Name == "SequenceDelimitationItem" {
					break
				}

				p.appendDataElement(file, elem)
				c <- DicomMessage{elem, waitMsg}
				<-waitMsg

				elem = buffer.readDataElement(p)
				elem.IndentLevel = sq.IndentLevel

			}
		} else {
			// ITEM 0 LEN
		}
	}

	return sqAcum, nil

}

func (p *Parser) readPixelItems(file *DicomFile, buffer *dicomBuffer, sq *DicomElement, c chan DicomMessage) {
	waitMsg := make(chan bool)

	elem := buffer.readDataElement(p)

	for buffer.Len() != 0 {
		if elem.Name == "Item" {
			elem.Value = append(elem.Value, buffer.readUInt8Array(elem.Vl))
		}
		p.appendDataElement(file, elem)
		c <- DicomMessage{elem, waitMsg}
		<-waitMsg
		elem = buffer.readDataElement(p)

	}
	p.appendDataElement(file, elem)
	c <- DicomMessage{elem, waitMsg}
	<-waitMsg

}

// Append a dataElement to the DicomFile
func (p *Parser) appendDataElement(file *DicomFile, elem *DicomElement) {
	file.Elements = append(file.Elements, *elem)

}

// Finds the SyntaxTrasnferUID and returns the endianess and implicit VR for the file
func (file *DicomFile) getTransferSyntax() (binary.ByteOrder, bool, error) {

	var err error

	elem, err := file.LookupElement("TransferSyntaxUID")
	if err != nil {
		return nil, true, err
	}

	ts := elem.Value[0].(string)

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
=======
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/grailbio/go-dicom/dicomio"
	"github.com/grailbio/go-dicom/dicomtag"
)

// GoDICOMImplementationClassUIDPrefix defines the UID prefix for
// go-dicom. Provided by https://www.medicalconnections.co.uk/Free_UID
const GoDICOMImplementationClassUIDPrefix = "1.2.826.0.1.3680043.9.7133"

var GoDICOMImplementationClassUID = GoDICOMImplementationClassUIDPrefix + ".1.1"

const GoDICOMImplementationVersionName = "GODICOM_1_1"

// DataSet represents contents of one DICOM file.
type DataSet struct {
	// Elements in the file, in order of appearance.
	//
	// Note: unlike pydicom, Elements also contains meta elements (those
	// with Tag.Group==2).
	Elements []*Element
}

func doassert(cond bool, values ...interface{}) {
	if !cond {
		var s string
		for _, value := range values {
			s += fmt.Sprintf("%v ", value)
		}
		panic(s)
	}
}

// ReadOptions defines how DataSets and Elements are parsed.
type ReadOptions struct {
	// DropPixelData will cause the parser to skip the PixelData element
	// (bulk images) in ReadDataSet.
	DropPixelData bool

	// ReturnTags is a whitelist of tags to return.
	ReturnTags []dicomtag.Tag

	// StopAtag defines a tag at which when read (or a tag with a greater
	// value than it is read), the program will stop parsing the dicom file.
	StopAtTag *dicomtag.Tag
}

// ReadDataSetInBytes is a shorthand for ReadDataSet(bytes.NewBuffer(data), len(data)).
//
// On parse error, this function may return a non-nil dataset and a non-nil
// error. In such case, the dataset will contain parts of the file that are
// parsable, and error will show the first error found by the parser.
func ReadDataSetInBytes(data []byte, options ReadOptions) (*DataSet, error) {
	return ReadDataSet(bytes.NewBuffer(data), int64(len(data)), options)
}

// ReadDataSetFromFile parses file cotents into dicom.DataSet. It is a thin
// wrapper around ReadDataSet.
//
// On parse error, this function may return a non-nil dataset and a non-nil
// error. In such case, the dataset will contain parts of the file that are
// parsable, and error will show the first error found by the parser.
func ReadDataSetFromFile(path string, options ReadOptions) (*DataSet, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	st, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return ReadDataSet(file, st.Size(), options)
}

// ReadDataSet reads a DICOM file from "io", up to "bytes".
//
// On parse error, this function may return a non-nil dataset and a non-nil
// error. In such case, the dataset will contain parts of the file that are
// parsable, and error will show the first error found by the parser.
func ReadDataSet(in io.Reader, bytes int64, options ReadOptions) (*DataSet, error) {
	buffer := dicomio.NewDecoder(in, bytes, binary.LittleEndian, dicomio.ExplicitVR)
	metaElems := ParseFileHeader(buffer)
	if buffer.Error() != nil {
		return nil, buffer.Error()
	}
	file := &DataSet{Elements: metaElems}

	// Change the transfer syntax for the rest of the file.
	endian, implicit, err := getTransferSyntax(file)
	if err != nil {
		return nil, err
	}
	buffer.PushTransferSyntax(endian, implicit)
	defer buffer.PopTransferSyntax()

	// Read the list of elements.
	for buffer.Len() > 0 {
		startLen := buffer.Len()
		elem := ReadElement(buffer, options)
		if buffer.Len() >= startLen { // Avoid silent infinite looping.
			log.Panicf("ReadElement failed to consume data: %d %d: %v", startLen, buffer.Len(), buffer.Error())
		}
		if elem == endOfDataElement {
			// element is a pixel data and was dropped by options
			break
		}
		if elem == nil {
			// Parse error.
			continue
		}
		if elem.Tag == dicomtag.SpecificCharacterSet {
			// Set the []byte -> string decoder for the rest of the
			// file.  It's sad that SpecificCharacterSet isn't part
			// of metadata, but is part of regular attrs, so we need
			// to watch out for multiple occurrences of this type of
			// elements.
			encodingNames, err := elem.GetStrings()
			if err != nil {
				buffer.SetError(err)
			} else {
				// TODO(saito) SpecificCharacterSet may appear in a
				// middle of a SQ or NA.  In such case, the charset seem
				// to be scoped inside the SQ or NA. So we need to make
				// the charset a stack.
				cs, err := dicomio.ParseSpecificCharacterSet(encodingNames)
				if err != nil {
					buffer.SetError(err)
				} else {
					buffer.SetCodingSystem(cs)
				}
			}
		}
		if options.ReturnTags == nil || (options.ReturnTags != nil && tagInList(elem.Tag, options.ReturnTags)) {
			file.Elements = append(file.Elements, elem)
		}
	}
	return file, buffer.Error()
}

func getTransferSyntax(ds *DataSet) (bo binary.ByteOrder, implicit dicomio.IsImplicitVR, err error) {
	elem, err := ds.FindElementByTag(dicomtag.TransferSyntaxUID)
	if err != nil {
		return nil, dicomio.UnknownVR, err
	}
	transferSyntaxUID, err := elem.GetString()
	if err != nil {
		return nil, dicomio.UnknownVR, err
	}
	return dicomio.ParseTransferSyntaxUID(transferSyntaxUID)
}

// FindElementByName finds an element from the dataset given the element name,
// such as "PatientName".
func (f *DataSet) FindElementByName(name string) (*Element, error) {
	return FindElementByName(f.Elements, name)
}

// FindElementByTag finds an element from the dataset given its tag, such as
// Tag{0x0010, 0x0010}.
func (f *DataSet) FindElementByTag(tag dicomtag.Tag) (*Element, error) {
	return FindElementByTag(f.Elements, tag)
>>>>>>> 3bad039... Initial commit
}
