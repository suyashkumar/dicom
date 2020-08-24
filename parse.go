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
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/suyashkumar/dicom/pkg/uid"
)

const (
	magicWord = "DICM"
)

var (
	ErrorMagicWord              = errors.New("error, DICM magic word not found in correct location")
	ErrorMetaElementGroupLength = errors.New("MetaElementGroupLength tag not found where expected")
	ErrorEndOfDICOM             = errors.New("this indicates to the caller of Next() that the DICOM has been fully parsed.")
)

// Parse parses the entire DICOM at the input io.Reader into a Dataset of DICOM Elements. Use this if you are
// looking to parse the DICOM all at once, instead of element-by-element.
func Parse(in io.Reader, bytesToRead int64, frameChan chan *frame.Frame) (Dataset, error) {
	p, err := NewParser(in, bytesToRead, frameChan)
	if err != nil {
		return Dataset{}, err
	}

	for !p.reader.IsLimitExhausted() {
		_, err := p.Next()
		if err != nil {
			return p.dataset, err
		}
	}

	// Close the frameChannel if needed
	if p.frameChannel != nil {
		close(p.frameChannel)
	}
	return p.dataset, nil
}

// Parser is a struct that allows a user to parse Elements from a DICOM element-by-element using Next(), which may be
// useful for some streaming processing applications. If you instead just want to parse the whole input DICOM at once,
// just use the dicom.Parse(...) method.
type Parser struct {
	reader   dicomio.Reader
	dataset  Dataset
	metadata Dataset
	// file is optional, might be populated if reading from an underlying file
	file         *os.File
	frameChannel chan *frame.Frame
}

// NewParser returns a new Parser that points to the provided io.Reader, with bytesToRead bytes left to read. NewParser
// will read the DICOM header and metadata as part of initialization.
//
// frameChannel is an optional channel (can be nil) upon which DICOM image frames will be sent as they are parsed (if
// provided).
func NewParser(in io.Reader, bytesToRead int64, frameChannel chan *frame.Frame) (*Parser, error) {
	reader, err := dicomio.NewReader(bufio.NewReader(in), binary.LittleEndian, bytesToRead)
	if err != nil {
		return nil, err
	}

	p := Parser{
		reader:       reader,
		frameChannel: frameChannel,
	}

	elems, err := p.readHeader()
	if err != nil {
		return nil, err
	}

	p.dataset = Dataset{Elements: elems}
	// TODO(suyashkumar): avoid storing the metadata pointers twice (though not that expensive)
	p.metadata = Dataset{Elements: elems}

	// Determine and set the transfer syntax based on the metadata elements parsed so far.
	ts, err := p.dataset.FindElementByTag(tag.TransferSyntaxUID)
	if err != nil {
		// proceed with a default?
		log.Println("WARN: could not find transfer syntax uid in metadata, proceeding with little endian implicit")
	}
	bo, implicit, err := uid.ParseTransferSyntaxUID(MustGetStrings(ts.Value)[0])
	if err != nil {
		// proceed with a default?
		log.Println("WARN: could not parse transfer syntax uid in metadata, proceeding with little endian implicit")
	}
	p.reader.SetTransferSyntax(bo, implicit)

	return &p, nil
}

// Next parses and returns the next top-level element from the DICOM this Parser points to.
func (p *Parser) Next() (*Element, error) {
	if p.reader.IsLimitExhausted() {
		// Close the frameChannel if needed
		if p.frameChannel != nil {
			close(p.frameChannel)
		}
		return nil, ErrorEndOfDICOM
	}
	elem, err := readElement(p.reader, &p.dataset, p.frameChannel)
	if err != nil {
		// TODO: tolerate some kinds of errors and continue parsing
		return nil, err
	}

	// TODO: add dicom options to only keep track of certain tags

	if elem.Tag == tag.SpecificCharacterSet {
		encodingNames := MustGetStrings(elem.Value)
		cs, err := charset.ParseSpecificCharacterSet(encodingNames)
		if err != nil {
			// unable to parse character set, hard error
			// TODO: add option continue, even if unable to parse
			return nil, err
		}
		p.reader.SetCodingSystem(cs)
	}

	p.dataset.Elements = append(p.dataset.Elements, elem)
	return elem, nil

}

func (p *Parser) GetMetadata() Dataset {
	return p.metadata
}

// readHeader reads the DICOM magic header and group two metadata elements.
func (p *Parser) readHeader() ([]*Element, error) {
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

	if maybeMetaLen.Tag != tag.FileMetaInformationGroupLength || maybeMetaLen.Value.ValueType() != Ints {
		return nil, ErrorMetaElementGroupLength
	}

	metaLen := maybeMetaLen.Value.GetValue().([]int)[0]

	metaElems := []*Element{maybeMetaLen} // TODO: maybe set capacity to a reasonable initial size

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
		// log.Printf("Metadata Element: %s\n", elem)
		metaElems = append(metaElems, elem)
	}
	return metaElems, nil
}
