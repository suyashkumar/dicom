package dicomio_test

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/gradienthealth/dicom/dicomio"
	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	e := dicomio.NewBytesEncoder(binary.BigEndian, dicomio.UnknownVR)
	e.WriteByte(10)
	e.WriteByte(11)
	e.WriteUInt16(0x123)
	e.WriteUInt32(0x234)
	e.WriteZeros(12)
	e.WriteString("abcde")
	encoded := e.Bytes()
	d := dicomio.NewDecoder(
		bytes.NewBuffer(encoded), int64(len(encoded)),
		binary.BigEndian, dicomio.ImplicitVR)
	require.Equal(t, byte(10), d.ReadByte())
	require.Equal(t, byte(11), d.ReadByte())
	require.Equal(t, uint16(0x123), d.ReadUInt16())
	require.Equal(t, uint32(0x234), d.ReadUInt32())
	d.Skip(12)
	require.Equal(t, "abcde", d.ReadString(5))
	require.Equal(t, int64(0), d.Len())
	require.NoError(t, d.Error())

	// Read past the buffer. It should flag an error
	d.ReadByte()
	require.Error(t, d.Error())
}

func TestSkip(t *testing.T) {
	e := dicomio.NewBytesEncoder(binary.BigEndian, dicomio.UnknownVR)
	e.WriteString("abcdefghijk")
	encoded := e.Bytes()
	d := dicomio.NewBytesDecoder(encoded, binary.BigEndian, dicomio.UnknownVR)
	d.Skip(3)
	require.Equal(t, int64(8), d.Len())
	require.Equal(t, "defghijk", d.ReadString(8))
}

func TestPartialData(t *testing.T) {
	e := dicomio.NewBytesEncoder(binary.BigEndian, dicomio.UnknownVR)
	e.WriteByte(10)
	encoded := e.Bytes()
	// Read uint16, when there's only one byte in buffer.
	d := dicomio.NewDecoder(bytes.NewBuffer(encoded), int64(len(encoded)),
		binary.BigEndian, dicomio.ImplicitVR)
	if _ = d.ReadUInt16(); d.Error() == nil {
		t.Errorf("ReadUint16")
	}
}

func TestLimit(t *testing.T) {
	e := dicomio.NewBytesEncoder(binary.BigEndian, dicomio.UnknownVR)
	e.WriteByte(10)
	e.WriteByte(11)
	e.WriteByte(12)
	encoded := e.Bytes()
	// Allow reading only the first two bytes
	d := dicomio.NewDecoder(bytes.NewBuffer(encoded), int64(len(encoded)),
		binary.BigEndian, dicomio.ImplicitVR)
	require.Equal(t, int64(3), d.Len())
	d.PushLimit(2)
	require.Equal(t, int64(2), d.Len())
	v0, v1 := d.ReadByte(), d.ReadByte()
	require.Equal(t, int64(0), d.Len())
	_ = d.ReadByte()
	if v0 != 10 || v1 != 11 || d.Error() != io.EOF {
		t.Errorf("Limit: %v %v %v", v0, v1, d.Error())
	}
}
