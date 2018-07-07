package dicom

import (
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/gradienthealth/go-dicom/dicomio"
	"github.com/gradienthealth/go-dicom/dicomtag"
	"github.com/gradienthealth/go-dicom/dicomuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testWriteDataElement(t *testing.T, bo binary.ByteOrder, implicit dicomio.IsImplicitVR) {
	// Encode two scalar elements.
	e := dicomio.NewBytesEncoder(bo, implicit)
	var values []interface{}
	values = append(values, string("FooHah"))
	WriteElement(e, &Element{
		Tag:   dicomtag.Tag{0x0018, 0x9755},
		Value: values})
	values = nil
	values = append(values, uint32(1234))
	values = append(values, uint32(2345))
	WriteElement(e, &Element{
		Tag:   dicomtag.Tag{0x0020, 0x9057},
		Value: values})
	data := e.Bytes()
	// Read them back.
	d := dicomio.NewBytesDecoder(data, bo, implicit)
	p := parser{
		decoder: d,
	}
	elem0 := p.ParseNext(ParseOptions{})

	require.NoError(t, p.DecoderError())
	tag := dicomtag.Tag{0x18, 0x9755}
	assert.Equal(t, elem0.Tag, tag)
	assert.Equal(t, len(elem0.Value), 1)
	assert.Equal(t, elem0.Value[0].(string), "FooHah")
	tag = dicomtag.Tag{Group: 0x20, Element: 0x9057}
	elem1 := p.ParseNext(ParseOptions{})
	require.NoError(t, p.DecoderError())
	assert.Equal(t, elem1.Tag, tag)
	assert.Equal(t, len(elem1.Value), 2)
	assert.Equal(t, elem1.Value[0].(uint32), uint32(1234))
	assert.Equal(t, elem1.Value[1].(uint32), uint32(2345))
	require.NoError(t, p.Finish())
}

func TestWriteDataElementImplicit(t *testing.T) {
	testWriteDataElement(t, binary.LittleEndian, dicomio.ImplicitVR)
}

func TestWriteDataElementExplicit(t *testing.T) {
	testWriteDataElement(t, binary.LittleEndian, dicomio.ExplicitVR)
}

func TestWriteDataElementBigEndianExplicit(t *testing.T) {
	testWriteDataElement(t, binary.BigEndian, dicomio.ExplicitVR)
}

func TestReadWriteFileHeader(t *testing.T) {
	e := dicomio.NewBytesEncoder(binary.LittleEndian, dicomio.ImplicitVR)
	WriteFileHeader(
		e,
		[]*Element{
			MustNewElement(dicomtag.TransferSyntaxUID, dicomuid.ImplicitVRLittleEndian),
			MustNewElement(dicomtag.MediaStorageSOPClassUID, "1.2.840.10008.5.1.4.1.1.1.2"),
			MustNewElement(dicomtag.MediaStorageSOPInstanceUID, "1.2.3.4.5.6.7"),
		})
	bytes := e.Bytes()
	d := dicomio.NewBytesDecoder(bytes, binary.LittleEndian, dicomio.ImplicitVR)
	p := parser{
		decoder: d,
	}
	elems := p.ParseFileHeader()
	require.NoError(t, d.Finish())
	elem, err := FindElementByTag(elems, dicomtag.TransferSyntaxUID)
	require.NoError(t, err)
	assert.Equalf(t, elem.MustGetString(), dicomuid.ImplicitVRLittleEndian,
		"Wrong element value %+v", elem)
	elem, err = FindElementByTag(elems, dicomtag.MediaStorageSOPClassUID)
	require.NoError(t, err)
	assert.Equal(t, elem.MustGetString(), "1.2.840.10008.5.1.4.1.1.1.2")
	elem, err = FindElementByTag(elems, dicomtag.MediaStorageSOPInstanceUID)
	require.NoError(t, err)
	assert.Equal(t, elem.MustGetString(), "1.2.3.4.5.6.7")
}

func TestNewElement(t *testing.T) {
	elem, err := NewElement(dicomtag.TriggerSamplePosition, uint32(10), uint32(11))
	require.NoError(t, err)
	require.Equal(t, elem.Tag, dicomtag.TriggerSamplePosition)
	require.Truef(t, reflect.DeepEqual(elem.MustGetUint32s(), []uint32{10, 11}),
		"Elem: %+v", elem)
	// Pass a wrong value type.
	elem, err = NewElement(dicomtag.TriggerSamplePosition, "foo")
	require.Error(t, err)
}
