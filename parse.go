// Package dicom provides a set of tools to read, write, and generally
// work with DICOM (https://dicom.nema.org/) medical image files in Go.
//
// dicom.Parse and dicom.Write provide the core functionality to read and write
// DICOM Datasets. This package provides Go data structures that represent
// DICOM concepts (for example, dicom.Dataset and dicom.Element). These
// structures will pretty-print by default and are JSON serializable out of the
// box.
//
// This package provides some advanced functionality as well, including:
// streaming image frames to an output channel, reading elements one-by-one
// (like an iterator pattern), flat iteration over nested elements in a Dataset,
// and more.
//
// General usage is simple.
// Check out the package examples below and some function specific examples.
//
// It may also be helpful to take a look at the example cmd/dicomutil program,
// which is a CLI built around this library to save out image frames from DICOMs
// and print out metadata to STDOUT.
package dicom

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"os"

	"github.com/suyashkumar/dicom/pkg/charset"
	"github.com/suyashkumar/dicom/pkg/debug"
	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/suyashkumar/dicom/pkg/uid"
)

const (
	magicWord = "DICM"
)

var (
	// ErrorMagicWord indicates that the magic word was not found in the correct
	// location in the DICOM.
	ErrorMagicWord = errors.New("error, DICM magic word not found in correct location")
	// ErrorMetaElementGroupLength indicates that the MetaElementGroupLength
	// was not found where expected in the metadata.
	ErrorMetaElementGroupLength = errors.New("MetaElementGroupLength tag not found where expected")
	// ErrorEndOfDICOM indicates to the callers of Parser.Next() that the DICOM
	// has been fully parsed. Users using one of the other Parse APIs should not
	// need to use this.
	ErrorEndOfDICOM = errors.New("this indicates to the caller of Next() that the DICOM has been fully parsed")

	// ErrorMismatchPixelDataLength indicates that the size calculated from DICOM mismatch the VL.
	ErrorMismatchPixelDataLength = errors.New("the size calculated from DICOM elements and the PixelData element's VL are mismatched")
)

func Parse(in io.Reader, bytesToRead int64, frameChan chan *frame.Frame, opts ...ParseOption) (Dataset, error) {
	return parseInternal(in, bytesToRead, frameChan, opts...)
}

func ParseUntilEOF(in io.Reader, frameChan chan *frame.Frame, opts ...ParseOption) (Dataset, error) {
	return parseInternal(in, dicomio.LimitReadUntilEOF, frameChan, opts...)
}

// Parse parses the entire DICOM at the input io.Reader into a Dataset of DICOM Elements. Use this if you are
// looking to parse the DICOM all at once, instead of element-by-element.
func parseInternal(in io.Reader, bytesToRead int64, frameChan chan *frame.Frame, opts ...ParseOption) (Dataset, error) {
	p, err := NewParser(in, bytesToRead, frameChan, opts...)
	if err != nil {
		return Dataset{}, err
	}

	for !p.reader.rawReader.IsLimitExhausted() {
		_, err := p.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// exiting on EOF
				err = nil
				break
			}

			// exit on error
			return p.dataset, err
		}
	}

	// Close the frameChannel if needed
	if p.frameChannel != nil {
		close(p.frameChannel)
	}
	return p.dataset, nil
}

// ParseFile parses the entire DICOM at the given filepath. See dicom.Parse as
// well for a more generic io.Reader based API.
func ParseFile(filepath string, frameChan chan *frame.Frame, opts ...ParseOption) (Dataset, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return Dataset{}, err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return Dataset{}, err
	}

	return Parse(f, info.Size(), frameChan, opts...)
}

// Parser is a struct that allows a user to parse Elements from a DICOM element-by-element using Next(), which may be
// useful for some streaming processing applications. If you instead just want to parse the whole input DICOM at once,
// just use the dicom.Parse(...) method.
type Parser struct {
	reader   *reader
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
func NewParser(in io.Reader, bytesToRead int64, frameChannel chan *frame.Frame, opts ...ParseOption) (*Parser, error) {
	optSet := toParseOptSet(opts...)
	p := Parser{
		reader: &reader{
			rawReader: dicomio.NewReader(bufio.NewReader(in), binary.LittleEndian, bytesToRead),
			opts:      optSet,
		},
		frameChannel: frameChannel,
	}

	elems := []*Element{}
	var err error
	if !optSet.skipMetadataReadOnNewParserInit {
		debug.Log("NewParser: readHeader")
		elems, err = p.reader.readHeader()
		if err != nil {
			return nil, err
		}
		debug.Log("NewParser: readHeader complete")
	}

	p.dataset = Dataset{Elements: elems}

	// TODO(suyashkumar): avoid storing the metadata pointers twice (though not that expensive)
	p.metadata = Dataset{Elements: elems}

	// Determine and set the transfer syntax based on the metadata elements parsed so far.
	// The default will be LittleEndian Implicit.
	var bo binary.ByteOrder = binary.LittleEndian
	implicit := true

	ts, err := p.dataset.FindElementByTag(tag.TransferSyntaxUID)
	if err != nil {
		debug.Log("WARN: could not find transfer syntax uid in metadata, proceeding with little endian implicit")
	} else {
		bo, implicit, err = uid.ParseTransferSyntaxUID(MustGetStrings(ts.Value)[0])
		if err != nil {
			// TODO(suyashkumar): should we attempt to parse with LittleEndian
			// Implicit here?
			debug.Log("WARN: could not parse transfer syntax uid in metadata")
		}
	}
	p.SetTransferSyntax(bo, implicit)

	return &p, nil
}

// Next parses and returns the next top-level element from the DICOM this Parser points to.
func (p *Parser) Next() (*Element, error) {
	if !p.reader.moreToRead() {
		// Close the frameChannel if needed
		if p.frameChannel != nil {
			close(p.frameChannel)
		}
		return nil, ErrorEndOfDICOM
	}
	elem, err := p.reader.readElement(&p.dataset, p.frameChannel)
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
		p.reader.rawReader.SetCodingSystem(cs)
	}

	p.dataset.Elements = append(p.dataset.Elements, elem)
	return elem, nil

}

// GetMetadata returns just the set of metadata elements that have been parsed
// so far.
func (p *Parser) GetMetadata() Dataset {
	return p.metadata
}

// SetTransferSyntax sets the transfer syntax for the underlying dicomio.Reader.
func (p *Parser) SetTransferSyntax(bo binary.ByteOrder, implicit bool) {
	p.reader.rawReader.SetTransferSyntax(bo, implicit)
}

// ParseOption represents an option that can be passed to NewParser.
type ParseOption func(*parseOptSet)

// parseOptSet represents the flattened option set after all ParseOptions have been applied.
type parseOptSet struct {
	skipMetadataReadOnNewParserInit bool
	allowMismatchPixelDataLength    bool
	skipPixelData                   bool
	skipProcessingPixelDataValue    bool
}

func toParseOptSet(opts ...ParseOption) parseOptSet {
	optSet := parseOptSet{}
	for _, opt := range opts {
		opt(&optSet)
	}
	return optSet
}

// AllowMismatchPixelDataLength allows parser to ignore an error when the length calculated from elements do not match with value length.
func AllowMismatchPixelDataLength() ParseOption {
	return func(set *parseOptSet) {
		set.allowMismatchPixelDataLength = true
	}
}

// SkipMetadataReadOnNewParserInit makes NewParser skip trying to parse metadata. This will make the Parser default to implicit little endian byte order.
// Any metatata tags found in the dataset will still be available when parsing.
func SkipMetadataReadOnNewParserInit() ParseOption {
	return func(set *parseOptSet) {
		set.skipMetadataReadOnNewParserInit = true
	}
}

// SkipPixelData skips reading data from the PixelData tag, wherever it appears
// (e.g. even if within an IconSequence). A PixelDataInfo will be added to the
// Dataset with the IntentionallySkipped property set to true, and no other
// data. Use this option if you don't need the PixelData value to be in the
// Dataset at all, and want to save both CPU and Memory. If you need the
// PixelData value in the Dataset (e.g. so it can be written out identically
// later) but _don't_ want to process/parse the value, see the
// SkipProcessingPixelDataValue option below.
func SkipPixelData() ParseOption {
	return func(set *parseOptSet) {
		set.skipPixelData = true
	}
}

// SkipProcessingPixelDataValue will attempt to skip processing the _value_
// of any PixelData elements. Unlike SkipPixelData(), this means the PixelData
// bytes will still be read into the Dataset, and can be written back out via
// this library's write functionality. But, if possible, the value will be read
// in as raw bytes with no further processing instead of being parsed. In the
// future, we may be able to extend this functionality to support on-demand
// processing of elements elsewhere in the library.
func SkipProcessingPixelDataValue() ParseOption {
	return func(set *parseOptSet) {
		set.skipProcessingPixelDataValue = true
	}
}
