package dicom

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gradienthealth/go-dicom/dicomio"
	"github.com/gradienthealth/go-dicom/dicomlog"
	"github.com/gradienthealth/go-dicom/dicomtag"
)

// GoDICOMImplementationClassUIDPrefix defines the UID prefix for
// go-dicom. Provided by https://www.medicalconnections.co.uk/Free_UID
const GoDICOMImplementationClassUIDPrefix = "1.2.826.0.1.3680043.9.7133"

var GoDICOMImplementationClassUID = GoDICOMImplementationClassUIDPrefix + ".1.1"

const GoDICOMImplementationVersionName = "GODICOM_1_1"

// Parser represents an entity that can read and parse DICOMs
type Parser interface {
	// Parse DICOM data
	Parse(options ParseOptions) (*DataSet, error)
	// ParseNext reads and parses the next element
	ParseNext(options ParseOptions) *Element
	// DecoderError fetches an error (if exists) from the dicomio.Decoder
	DecoderError() error // This should go away as we continue refactors
	// Finish should be called after manually parsing elements using ParseNext (instead of Parse)
	Finish() error // This should maybe go away as we continue refactors
}

// parser implements Parser
type parser struct {
	decoder        *dicomio.Decoder
	parsedElements *DataSet
	op             ParseOptions
	frameChannel   chan *Frame
	file           *os.File // may be populated if coming from file
}

// NewParser initializes and returns a new Parser
func NewParser(in io.Reader, bytesToRead int64, frameChannel chan *Frame) (Parser, error) {
	buffer := dicomio.NewDecoder(in, bytesToRead, binary.LittleEndian, dicomio.ExplicitVR)
	p := parser{
		decoder:      buffer,
		frameChannel: frameChannel,
	}

	metaElems := p.parseFileHeader()
	if buffer.Error() != nil {
		return nil, buffer.Error()
	}
	parsedElements := &DataSet{Elements: metaElems}
	p.parsedElements = parsedElements
	return &p, nil
}

// NewParserFromBytes initializes and returns a new Parser from []byte
func NewParserFromBytes(data []byte, frameChannel chan *Frame) (Parser, error) {
	return NewParser(bytes.NewBuffer(data), int64(len(data)), frameChannel)
}

// NewParserFromFile initializes and returns a new dicom Parser from a file path
func NewParserFromFile(path string, frameChannel chan *Frame) (Parser, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	st, err := file.Stat()
	if err != nil {
		return nil, err
	}
	p, err := NewParser(file, st.Size(), frameChannel)
	p.(*parser).file = file
	return p, err
}

func (p *parser) Parse(options ParseOptions) (*DataSet, error) {
	// Change the transfer syntax for the rest of the file.
	endian, implicit, err := getTransferSyntax(p.parsedElements)
	if err != nil {
		return nil, err
	}
	p.decoder.PushTransferSyntax(endian, implicit)
	defer p.decoder.PopTransferSyntax()

	// if reading from file, close the file after done parsing
	if p.file != nil {
		defer p.file.Close()
	}

	// Read the list of elements.
	for p.decoder.Len() > 0 {
		startLen := p.decoder.Len()
		elem := p.ParseNext(options)
		if p.decoder.Len() >= startLen { // Avoid silent infinite looping.
			panic(fmt.Sprintf("ReadElement failed to consume data: %decoder %decoder: %v", startLen, p.decoder.Len(), p.decoder.Error()))
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
				p.decoder.SetError(err)
			} else {
				// TODO(saito) SpecificCharacterSet may appear in a
				// middle of a SQ or NA.  In such case, the charset seem
				// to be scoped inside the SQ or NA. So we need to make
				// the charset a stack.
				cs, err := dicomio.ParseSpecificCharacterSet(encodingNames)
				if err != nil {
					p.decoder.SetError(err)
				} else {
					p.decoder.SetCodingSystem(cs)
				}
			}
		}
		if options.ReturnTags == nil || (options.ReturnTags != nil && tagInList(elem.Tag, options.ReturnTags)) {
			p.parsedElements.Elements = append(p.parsedElements.Elements, elem)
		}

	}
	return p.parsedElements, p.decoder.Error()
}

func (p *parser) ParseNext(options ParseOptions) *Element {
	tag := readTag(p.decoder)
	if tag == dicomtag.PixelData && options.DropPixelData {
		return endOfDataElement
	}

	// Return nil if the tag is greater than the StopAtTag if a StopAtTag is given
	if options.StopAtTag != nil && tag.Group >= options.StopAtTag.Group && tag.Element >= options.StopAtTag.Element {
		return endOfDataElement
	}
	// The elements for group 0xFFFE should be Encoded as Implicit VR.
	// DICOM Standard 09. PS 3.6 - Section 7.5: "Nesting of Data Sets"
	_, implicit := p.decoder.TransferSyntax()
	if tag.Group == itemSeqGroup {
		implicit = dicomio.ImplicitVR
	}
	var vr string // Value Representation
	var vl uint32 // Value Length
	if implicit == dicomio.ImplicitVR {
		vr, vl = readImplicit(p.decoder, tag)
	} else {
		doassert(implicit == dicomio.ExplicitVR, implicit)
		vr, vl = readExplicit(p.decoder, tag)
	}
	var data []interface{}

	elem := &Element{
		Tag:             tag,
		VR:              vr,
		UndefinedLength: vl == undefinedLength,
	}
	if vr == "UN" && vl == undefinedLength {
		// This combination appears in some file, but it's unclear what
		// to do. The standard, as always, is unclear. The best guess is
		// in PS3.5, 6.2.2, where it states that the combination of
		// vr=UN and undefined length is allowed, it refers to a section
		// of parsing "Data Elemets with Unknown Length". That reference
		// is specifically about element of type SQ, so I'm just
		// assuming <UN, undefinedlength> is the same as <SQ, undefined
		// length>.
		vr = "SQ"
	}
	if tag == dicomtag.PixelData {
		// P3.5, A.4 describes the format.
		//
		// PixelData is usually the last element in a DICOM file. When
		// the file stores N images, the elements that follow PixelData
		// are laid out in the following way:
		//
		// Item(BasicOffsetTable) Item(PixelDataInfo0) ... Item(PixelDataInfoM) SequenceDelimiterItem
		//
		// Item(BasicOffsetTable) is an Item element whose payload
		// encodes N uint32 values. Kth uint32 is the bytesize of the
		// Kth image. Item(PixelDataInfo*) are chunked sequences of bytes. I
		// presume that single PixelDataInfo item doesn't cross a image
		// boundary, but the spec isn't clear.
		//
		// The total byte size of Item(PixelDataInfo*) equal the total of
		// the bytesizes found in BasicOffsetTable.
		if vl == undefinedLength {
			var image PixelDataInfo
			image.IsEncapsulated = true
			image.Offsets = readBasicOffsetTable(p.decoder) // TODO(saito) Use the offset table.
			if len(image.Offsets) > 1 {
				dicomlog.Vprintf(1, "dicom.ReadElement: Multiple images not supported yet. Combining them into a byte sequence: %v", image.Offsets)
			}
			for p.decoder.Len() > 0 {
				chunk, endOfItems := readRawItem(p.decoder)
				if p.decoder.Error() != nil {
					break
				}
				if endOfItems {
					break
				}

				// Construct frame
				f := Frame{
					IsEncapsulated: true,
					EncapsulatedData: EncapsulatedFrame{
						Data: chunk,
					},
				}

				image.Frames = append(image.Frames, f)
				if p.frameChannel != nil {
					p.frameChannel <- &f // write frame to channel
				}
			}
			if p.frameChannel != nil {
				close(p.frameChannel)
			}
			data = append(data, image)
		} else {
			// Assume we're reading NativeData data since we have a defined value length as per Part 5 Sec A.4 of DICOM spec.
			// We need Elements that have been already parsed (rows, cols, etc) to parse frames out of NativeData Pixel data
			if p.parsedElements == nil {
				p.decoder.SetError(errors.New("dicom.ReadElement: parsedData is nil, must exist to parse NativeData pixel data"))
				return nil // TODO(suyash) investigate error handling in this library
			}

			image, _, err := readNativeFrames(p.decoder, p.parsedElements, p.frameChannel)

			if err != nil {
				p.decoder.SetError(err)
				dicomlog.Vprintf(1, "dicom.ReadElement: Error reading native frames")
				return nil
			}

			data = append(data, *image)
		}
	} else if vr == "SQ" {
		// Note: when reading subitems inside sequence or item, we ignore
		// DropPixelData and other shortcircuiting options. If we honored them, we'decoder
		// be unable to read the rest of the file.
		if vl == undefinedLength {
			// Format:
			//  Sequence := ItemSet* SequenceDelimitationItem
			//  ItemSet := Item Any* ItemDelimitationItem (when Item.VL is undefined) or
			//             Item Any*N                     (when Item.VL has a defined value)
			for {
				// Makes sure to return all sub elements even if the tag is not in the return tags list of options or is greater than the Stop At Tag
				item := p.ParseNext(ParseOptions{})
				if p.decoder.Error() != nil {
					break
				}
				if item.Tag == dicomtag.SequenceDelimitationItem {
					break
				}
				if item.Tag != dicomtag.Item {
					p.decoder.SetErrorf("dicom.ReadElement: Found non-Item element in seq w/ undefined length: %v", dicomtag.DebugString(item.Tag))
					break
				}
				data = append(data, item)
			}
		} else {
			// Format:
			//  Sequence := ItemSet*VL
			// See the above comment for the definition of ItemSet.
			p.decoder.PushLimit(int64(vl))
			for p.decoder.Len() > 0 {
				// Makes sure to return all sub elements even if the tag is not in the return tags list of options or is greater than the Stop At Tag
				item := p.ParseNext(ParseOptions{})
				if p.decoder.Error() != nil {
					break
				}
				if item.Tag != dicomtag.Item {
					p.decoder.SetErrorf("dicom.ReadElement: Found non-Item element in seq w/ undefined length: %v", dicomtag.DebugString(item.Tag))
					break
				}
				data = append(data, item)
			}
			p.decoder.PopLimit()
		}
	} else if tag == dicomtag.Item { // Item (component of SQ)
		if vl == undefinedLength {
			// Format: Item Any* ItemDelimitationItem
			for {
				// Makes sure to return all sub elements even if the tag is not in the return tags list of options or is greater than the Stop At Tag
				subelem := p.ParseNext(ParseOptions{})
				if p.decoder.Error() != nil {
					break
				}
				if subelem.Tag == dicomtag.ItemDelimitationItem {
					break
				}
				data = append(data, subelem)
			}
		} else {
			// Sequence of arbitary elements, for the  total of "vl" bytes.
			p.decoder.PushLimit(int64(vl))
			for p.decoder.Len() > 0 {
				// Makes sure to return all sub elements even if the tag is not in the return tags list of options or is greater than the Stop At Tag
				subelem := p.ParseNext(ParseOptions{})
				if p.decoder.Error() != nil {
					break
				}
				data = append(data, subelem)
			}
			p.decoder.PopLimit()
		}
	} else { // List of scalar
		if vl == undefinedLength {
			p.decoder.SetErrorf("dicom.ReadElement: Undefined length disallowed for VR=%s, tag %s", vr, dicomtag.DebugString(tag))
			return nil
		}
		p.decoder.PushLimit(int64(vl))
		defer p.decoder.PopLimit()
		if vr == "DA" {
			// TODO(saito) Maybe we should validate the date.
			date := strings.Trim(p.decoder.ReadString(int(vl)), " \000")
			data = []interface{}{date}
		} else if vr == "AT" {
			// (2byte group, 2byte elem)
			for p.decoder.Len() > 0 && p.decoder.Error() == nil {
				tag := dicomtag.Tag{p.decoder.ReadUInt16(), p.decoder.ReadUInt16()}
				data = append(data, tag)
			}
		} else if vr == "OW" {
			if vl%2 != 0 {
				p.decoder.SetErrorf("dicom.ReadElement: tag %v: OW requires even length, but found %v", dicomtag.DebugString(tag), vl)
			} else {
				n := int(vl / 2)
				e := dicomio.NewBytesEncoder(dicomio.NativeByteOrder, dicomio.UnknownVR)
				for i := 0; i < n; i++ {
					v := p.decoder.ReadUInt16()
					e.WriteUInt16(v)
				}
				doassert(e.Error() == nil, e.Error())
				// TODO(saito) Check that size is even. Byte swap??
				// TODO(saito) If OB's length is odd, is VL odd too? Need to check!
				data = append(data, e.Bytes())
			}
		} else if vr == "OB" {
			// TODO(saito) Check that size is even. Byte swap??
			// TODO(saito) If OB's length is odd, is VL odd too? Need to check!
			data = append(data, p.decoder.ReadBytes(int(vl)))
		} else if vr == "LT" || vr == "UT" {
			str := p.decoder.ReadString(int(vl))
			data = append(data, str)
		} else if vr == "UL" {
			for p.decoder.Len() > 0 && p.decoder.Error() == nil {
				data = append(data, p.decoder.ReadUInt32())
			}
		} else if vr == "SL" {
			for p.decoder.Len() > 0 && p.decoder.Error() == nil {
				data = append(data, p.decoder.ReadInt32())
			}
		} else if vr == "US" {
			for p.decoder.Len() > 0 && p.decoder.Error() == nil {
				data = append(data, p.decoder.ReadUInt16())
			}
		} else if vr == "SS" {
			for p.decoder.Len() > 0 && p.decoder.Error() == nil {
				data = append(data, p.decoder.ReadInt16())
			}
		} else if vr == "FL" {
			for p.decoder.Len() > 0 && p.decoder.Error() == nil {
				data = append(data, p.decoder.ReadFloat32())
			}
		} else if vr == "FD" {
			for p.decoder.Len() > 0 && p.decoder.Error() == nil {
				data = append(data, p.decoder.ReadFloat64())
			}
		} else {
			// List of strings, each delimited by '\\'.
			v := p.decoder.ReadString(int(vl))
			// String may have '\0' suffix if its length is odd.
			str := strings.Trim(v, " \000")
			if len(str) > 0 {
				for _, s := range strings.Split(str, "\\") {
					data = append(data, s)
				}
			}
		}
	}
	elem.Value = data
	return elem
}

func (p *parser) DecoderError() error {
	return p.decoder.Error()
}

func (p *parser) Finish() error {
	// if we've been reading from a file, close it
	if p.file != nil {
		p.file.Close()
	}
	return p.decoder.Finish()
}

// parseFileHeader consumes the DICOM magic header and metadata elements (whose
// elements with tag group==2) from a Dicom file. Errors are reported through
// decoder.Error().
func (p *parser) parseFileHeader() []*Element {
	p.decoder.PushTransferSyntax(binary.LittleEndian, dicomio.ExplicitVR)
	defer p.decoder.PopTransferSyntax()
	p.decoder.Skip(128) // skip preamble

	// check for magic word
	if s := p.decoder.ReadString(4); s != "DICM" {
		p.decoder.SetError(errors.New("Keyword 'DICM' not found in the header"))
		return nil
	}

	// (0002,0000) MetaElementGroupLength
	metaElem := p.ParseNext(ParseOptions{})
	if p.decoder.Error() != nil {
		return nil
	}
	if metaElem.Tag != dicomtag.FileMetaInformationGroupLength {
		p.decoder.SetErrorf("MetaElementGroupLength not found; insteadfound %s", metaElem.Tag.String())
	}
	metaLength, err := metaElem.GetUInt32()
	if err != nil {
		p.decoder.SetErrorf("Failed to read uint32 in MetaElementGroupLength: %v", err)
		return nil
	}
	if p.decoder.Len() <= 0 {
		p.decoder.SetErrorf("No data element found")
		return nil
	}
	metaElems := []*Element{metaElem}

	// Read meta tags
	p.decoder.PushLimit(int64(metaLength))
	defer p.decoder.PopLimit()
	for p.decoder.Len() > 0 {
		elem := p.ParseNext(ParseOptions{})
		if p.decoder.Error() != nil {
			break
		}
		metaElems = append(metaElems, elem)
		dicomlog.Vprintf(2, "dicom.parseFileHeader: Meta elem: %v, len %v", elem.String(), p.decoder.Len())
	}
	return metaElems
}

// DataSet represents contents of one DICOM file.
type DataSet struct {
	// Elements in the file, in order of appearance.
	//
	// Note: unlike pydicom, Elements also contains meta elements (those
	// with Tag.Group==2).
	Elements []*Element
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

// ParseOptions defines how DataSets and Elements are parsed.
type ParseOptions struct {
	// DropPixelData will cause the parser to skip the PixelData element
	// (bulk images) in ReadDataSet.
	DropPixelData bool

	// ReturnTags is a whitelist of tags to return.
	ReturnTags []dicomtag.Tag

	// StopAtag defines a tag at which when read (or a tag with a greater
	// value than it is read), the program will stop parsing the dicom file.
	StopAtTag *dicomtag.Tag
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

// readNativeFrames reads NativeData frames from a Decoder based on already parsed pixel information
// that should be available in parsedData (elements like NumberOfFrames, rows, columns, etc)
func readNativeFrames(d *dicomio.Decoder, parsedData *DataSet, frameChan chan *Frame) (pixelData *PixelDataInfo, bytesRead int, err error) {
	image := PixelDataInfo{
		IsEncapsulated: false,
	}

	// Parse information from previously parsed attributes that are needed to parse NativeData Frames:
	rows, err := parsedData.FindElementByTag(dicomtag.Rows)
	if err != nil {
		return nil, 0, err
	}

	cols, err := parsedData.FindElementByTag(dicomtag.Columns)
	if err != nil {
		return nil, 0, err
	}

	nof, err := parsedData.FindElementByTag(dicomtag.NumberOfFrames)
	nFrames := 0
	if err == nil {
		// No error, so parse number of frames
		nFrames, err = strconv.Atoi(nof.MustGetString()) // odd that number of frames is encoded as a string...
		if err != nil {
			dicomlog.Vprintf(1, "ERROR converting nof")
			return nil, 0, err
		}
	} else {
		// error fetching NumberOfFrames, so default to 1. TODO: revisit
		nFrames = 1
	}

	b, err := parsedData.FindElementByTag(dicomtag.BitsAllocated)
	if err != nil {
		dicomlog.Vprintf(1, "ERROR finding bits allocated.")
		return nil, 0, err
	}
	bitsAllocated := int(b.MustGetUInt16())

	s, err := parsedData.FindElementByTag(dicomtag.SamplesPerPixel)
	if err != nil {
		dicomlog.Vprintf(1, "ERROR finding samples per pixel")
	}
	samplesPerPixel := int(s.MustGetUInt16())

	pixelsPerFrame := int(rows.MustGetUInt16()) * int(cols.MustGetUInt16())

	dicomlog.Vprintf(1, "Image size: %decoder x %decoder", rows.MustGetUInt16(), cols.MustGetUInt16())
	dicomlog.Vprintf(1, "Pixels Per Frame: %decoder", pixelsPerFrame)
	dicomlog.Vprintf(1, "Number of frames %decoder", nFrames)

	// Parse the pixels:
	image.Frames = make([]Frame, nFrames)
	for frame := 0; frame < nFrames; frame++ {
		// Init current frame
		currentFrame := Frame{
			IsEncapsulated: false,
			NativeData: NativeFrame{
				BitsPerSample: bitsAllocated,
				Rows:          int(rows.MustGetUInt16()),
				Cols:          int(cols.MustGetUInt16()),
				Data:          make([][]int, int(pixelsPerFrame)),
			},
		}
		for pixel := 0; pixel < int(pixelsPerFrame); pixel++ {
			currentPixel := make([]int, samplesPerPixel)
			for value := 0; value < samplesPerPixel; value++ {
				if bitsAllocated == 8 {
					currentPixel[value] = int(d.ReadUInt8())
				} else if bitsAllocated == 16 {
					currentPixel[value] = int(d.ReadUInt16())
				}
			}
			currentFrame.NativeData.Data[pixel] = currentPixel
		}
		image.Frames[frame] = currentFrame
		if frameChan != nil {
			frameChan <- &currentFrame // write the current frame to the frameChan
		}
	}

	if frameChan != nil {
		close(frameChan) // close frameChan if it exists
	}

	bytesRead = (bitsAllocated / 8) * samplesPerPixel * pixelsPerFrame * nFrames

	return &image, bytesRead, nil
}
