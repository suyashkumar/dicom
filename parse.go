// Package dicom provides tools to read and work with DICOM binary files.
package dicom

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"os"

	"github.com/suyashkumar/dicom/pkg/charset"
	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/element"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/suyashkumar/dicom/pkg/uid"
)

const (
	magicWord = "DICM"
)

var (
	// ErrorMagicWord indicates that the DICM magic word was not found at the correct location in the dicom.
	ErrorMagicWord = errors.New("error, DICM magic word not found in correct location")
	// ErrorMetaElementGroupLength indicates that the MetaElementGroupLength was not found where expected.
	ErrorMetaElementGroupLength = errors.New("MetaElementGroupLength tag not found where expected")
)

// Parser represents a DICOM parser, initialized on a particular stream of DICOM bytes.
// It is able to parse out DICOM elements from the raw binary data and return a Dataset of elements.
// More information about these concepts can be found at http://dicom.nema.org/medical/dicom/current/output/chtml/part05/chapter_7.html.
//
// Check out the godoc for the element package to learn more about how these are represented in this library,
// and tools for working with these constructs.
type Parser interface {
	// Parse reads and parses DICOM binary data from this Parser and returns a Dataset of parsed Elements.
	Parse() (element.Dataset, error)
}

type parser struct {
	reader  dicomio.Reader
	dataset element.Dataset
	// file is optional, might be populated if reading from an underlying file
	file         *os.File
	frameChannel chan<- *frame.Frame
}

// NewParser creates a new Parser for the provided io.Reader and will read up to bytesToRead. frameChannel is optional,
// and if supplied, image frame.Frames are sent to the channel as they are parsed--useful for streaming applications.
func NewParser(in io.Reader, bytesToRead int64, frameChannel chan<- *frame.Frame) (Parser, error) {
	reader, err := dicomio.NewReader(bufio.NewReader(in), binary.LittleEndian, bytesToRead)
	if err != nil {
		return nil, err
	}

	p := parser{
		reader:       reader,
		frameChannel: frameChannel,
	}

	elems, err := p.readHeader()
	if err != nil {
		return nil, err
	}

	p.dataset = element.Dataset{Elements: elems}

	return &p, nil
}

// readHeader reads the DICOM magic header and group two metadata elements.
func (p *parser) readHeader() ([]*element.Element, error) {
	// Must read as LittleEndian explicit VR
	err := p.reader.Skip(128) // skip preamble
	if err != nil {
		log.Println("skip er")
		return nil, err
	}

	// Check DICOM magic word
	if s, err := p.reader.ReadString(4); err != nil || s != magicWord {
		return nil, ErrorMagicWord
	}

	// Read the length of the metadata elements: (0002,0000) MetaElementGroupLength
	maybeMetaLen, err := readElement(p.reader, nil, nil)
	if err != nil {
		log.Println("read element err")
		return nil, err
	}

	if maybeMetaLen.Tag != tag.FileMetaInformationGroupLength || maybeMetaLen.Value.ValueType() != element.Ints {
		return nil, ErrorMetaElementGroupLength
	}

	metaLen := maybeMetaLen.Value.GetValue().([]int)[0]

	metaElems := []*element.Element{maybeMetaLen} // TODO: maybe set capacity to a reasonable initial size

	// Read the metadata elements
	err = p.reader.PushLimit(int64(metaLen))
	if err != nil {
		return nil, err
	}
	defer p.reader.PopLimit()
	for !p.reader.IsLimitExhausted() {
		elem, err := readElement(p.reader, nil, nil)
		if err != nil {
			// TODO: see if we can skip over malformed elements somehow
			log.Println("read element err")

			return nil, err
		}
		log.Printf("Metadata Element: %s\n", elem)
		metaElems = append(metaElems, elem)
	}
	return metaElems, nil
}

func (p *parser) Parse() (element.Dataset, error) {
	// Determine and set the transfer syntax based on the metadata elements parsed so far.
	ts, err := p.dataset.FindElementByTag(tag.TransferSyntaxUID)
	if err != nil {
		// proceed with a default?
		log.Println("WARN: could not find transfer syntax uid in metadata, proceeding with little endian implicit")
	}
	bo, implicit, err := uid.ParseTransferSyntaxUID(element.MustGetString(ts.Value))
	if err != nil {
		// proceed with a default?
		log.Println("WARN: could not parse transfer syntax uid in metadata, proceeding with little endian implicit")
	}
	p.reader.SetTransferSyntax(bo, implicit)
	for !p.reader.IsLimitExhausted() {
		// TODO: avoid silent looping
		elem, err := readElement(p.reader, &p.dataset, p.frameChannel)
		if err != nil {
			// TODO: tolerate some kinds of errors and continue parsing
			return element.Dataset{}, err
		}

		log.Println("Read tag: ", elem.Tag)

		// TODO: add dicom options to only keep track of certain tags

		if elem.Tag == tag.SpecificCharacterSet {
			encodingNames := element.MustGetStrings(elem.Value)
			cs, err := charset.ParseSpecificCharacterSet(encodingNames)
			if err != nil {
				// unable to parse character set, hard error
				// TODO: add option continue, even if unable to parse
				return p.dataset, err
			}
			p.reader.SetCodingSystem(cs)
		}

		p.dataset.Elements = append(p.dataset.Elements, elem)

	}

	if p.frameChannel != nil {
		close(p.frameChannel)
	}
	return p.dataset, nil
}
