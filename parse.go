package dicom

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"os"

	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/tag"
)

const (
	MagicWord = "DICM"
)

var (
	ErrorMagicWord              = errors.New("error, DICM magic word not found in correct location")
	ErrorMetaElementGroupLength = errors.New("MetaElementGroupLength tag not found where expected")
)

type Parser interface {
	// Parse DICOM data
	Parse() (Dataset, error)
}

type parser struct {
	reader  dicomio.Reader
	dataset Dataset
	// file is optional, might be populated if reading from an underlying file
	file *os.File
}

func NewParser(in io.Reader, bytesToRead int64) (Parser, error) {
	reader, err := dicomio.NewReader(bufio.NewReader(in), binary.LittleEndian, bytesToRead)
	if err != nil {
		return nil, err
	}

	p := parser{
		reader: reader,
	}

	elems, err := p.readHeader()
	if err != nil {
		return nil, err
	}

	p.dataset = Dataset{Elements: elems}

	return &p, nil
}

func NewParserFromBytes(data []byte) (Parser, error) {
	return nil, nil
}

// readHeader reads the DICOM magic header and group two metadata elements.
func (p *parser) readHeader() ([]*Element, error) {
	// Must read as LittleEndian explicit VR
	err := p.reader.Skip(128) // skip preamble
	if err != nil {
		log.Println("skip er")
		return nil, err
	}

	// Check DICOM magic word
	if s, err := p.reader.ReadString(4); err != nil || s != MagicWord {
		return nil, ErrorMagicWord
	}

	// Read the length of the metadata elements: (0002,0000) MetaElementGroupLength
	maybeMetaLen, err := readElement(p.reader)
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
	p.reader.PushLimit(int64(metaLen))
	defer p.reader.PopLimit()
	for !p.reader.IsLimitExhausted() {
		elem, err := readElement(p.reader)
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

func (p *parser) Parse() (Dataset, error) {
	for !p.reader.IsLimitExhausted() {
		// TODO: avoid silent looping
		elem, err := readElement(p.reader)
		if err != nil {
			// TODO: tolerate some kinds of errors and continue parsing
			return Dataset{}, err
		}

		// TODO: if we encounter a dicomtag.SpecificCharacterSet, update the reader to accommodate
		// TODO: add dicom options to only keep track of certain tags

		p.dataset.Elements = append(p.dataset.Elements, elem)
	}

	return p.dataset, nil
}
