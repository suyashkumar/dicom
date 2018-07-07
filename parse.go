package dicom

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
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
}

// parser implements Parser
type parser struct {
	decoder        *dicomio.Decoder
	parsedElements *DataSet
	op             ParseOptions
	frameChannel   chan [][]int
	file           *os.File // may be populated if coming from file
}

func NewParser(in io.Reader, bytesToRead int64, frameChannel chan [][]int) (Parser, error) {
	buffer := dicomio.NewDecoder(in, bytesToRead, binary.LittleEndian, dicomio.ExplicitVR)
	metaElems := ParseFileHeader(buffer)
	if buffer.Error() != nil {
		return nil, buffer.Error()
	}
	parsedElements := &DataSet{Elements: metaElems}
	p := parser{
		decoder:        buffer,
		parsedElements: parsedElements,
		frameChannel:   frameChannel,
	}

	return &p, nil
}

func NewParserFromBytes(data []byte, frameChannel chan [][]int) (Parser, error) {
	return NewParser(bytes.NewBuffer(data), int64(len(data)), frameChannel)
}

func NewParserFromFile(path string, frameChannel chan [][]int) (Parser, error) {
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
			fmt.Println("BREAKED")
			fmt.Println(p.decoder.Error())
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
		// P3.5, A.4 describes the format. Currently we only support an encapsulated image format.
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
			image.Encapsulated = true
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
				image.EncapsulatedFrames = append(image.EncapsulatedFrames, chunk)
			}
			data = append(data, image)
		} else {
			// Assume we're reading Native data since we have a defined value length as per Part 5 Sec A.4 of DICOM spec.
			// We need Elements that have been already parsed (rows, cols, etc) to parse frames out of Native Pixel data
			if p.parsedElements == nil {
				p.decoder.SetError(errors.New("dicom.ReadElement: parsedData is nil, must exist to parse Native pixel data"))
				return nil // TODO(suyash) investigate error handling in this library
			}

			image, _, err := readNativeFrames(p.decoder, p.parsedElements)

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
				item := p.ParseNext(options)
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
				item := p.ParseNext(options)
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
				subelem := p.ParseNext(options)
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
				subelem := p.ParseNext(options)
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
