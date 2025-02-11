package dicomio

import (
    "bytes"
    "encoding/binary"
    "testing"
    "math"
)


// Test generated using Keploy
func TestNewWriter_Initialization(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    byteOrder := binary.LittleEndian
    implicit := true

    writer := NewWriter(mockWriter, byteOrder, implicit)

    if writer.out != mockWriter {
        t.Errorf("Expected writer.out to be %v, got %v", mockWriter, writer.out)
    }
    if writer.bo != byteOrder {
        t.Errorf("Expected writer.bo to be %v, got %v", byteOrder, writer.bo)
    }
    if writer.implicit != implicit {
        t.Errorf("Expected writer.implicit to be %v, got %v", implicit, writer.implicit)
    }
}

// Test generated using Keploy
func TestWriteZeros_CorrectLength(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    writer := NewWriter(mockWriter, binary.LittleEndian, true)

    err := writer.WriteZeros(5)
    if err != nil {
        t.Errorf("WriteZeros returned an error: %v", err)
    }

    expected := []byte{0, 0, 0, 0, 0}
    if !bytes.Equal(mockWriter.Bytes(), expected) {
        t.Errorf("Expected %v, got %v", expected, mockWriter.Bytes())
    }
}


// Test generated using Keploy
func TestWriteString_CorrectOutput(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    writer := NewWriter(mockWriter, binary.LittleEndian, true)

    input := "Hello, DICOM!"
    err := writer.WriteString(input)
    if err != nil {
        t.Errorf("WriteString returned an error: %v", err)
    }

    if mockWriter.String() != input {
        t.Errorf("Expected %v, got %v", input, mockWriter.String())
    }
}


// Test generated using Keploy
func TestWriteUInt16_CorrectOutput(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    writer := NewWriter(mockWriter, binary.LittleEndian, true)

    input := uint16(42)
    err := writer.WriteUInt16(input)
    if err != nil {
        t.Errorf("WriteUInt16 returned an error: %v", err)
    }

    expected := []byte{42, 0} // LittleEndian representation of 42
    if !bytes.Equal(mockWriter.Bytes(), expected) {
        t.Errorf("Expected %v, got %v", expected, mockWriter.Bytes())
    }
}


// Test generated using Keploy
func TestSetTransferSyntax_UpdatesFields(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    writer := NewWriter(mockWriter, binary.LittleEndian, true)

    newByteOrder := binary.BigEndian
    newImplicit := false
    writer.SetTransferSyntax(newByteOrder, newImplicit)

    if writer.bo != newByteOrder {
        t.Errorf("Expected writer.bo to be %v, got %v", newByteOrder, writer.bo)
    }
    if writer.implicit != newImplicit {
        t.Errorf("Expected writer.implicit to be %v, got %v", newImplicit, writer.implicit)
    }
}


// Test generated using Keploy
func TestGetTransferSyntax_ReturnsCorrectValues(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    byteOrder := binary.LittleEndian
    implicit := true
    writer := NewWriter(mockWriter, byteOrder, implicit)

    returnedBO, returnedImplicit := writer.GetTransferSyntax()

    if returnedBO != byteOrder {
        t.Errorf("Expected byte order %v, got %v", byteOrder, returnedBO)
    }
    if returnedImplicit != implicit {
        t.Errorf("Expected implicit %v, got %v", implicit, returnedImplicit)
    }
}


// Test generated using Keploy
func TestWriteByte_CorrectOutput(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    writer := NewWriter(mockWriter, binary.LittleEndian, true)

    input := byte(0xAB)
    err := writer.WriteByte(input)
    if err != nil {
        t.Errorf("WriteByte returned an error: %v", err)
    }

    expected := []byte{0xAB}
    if !bytes.Equal(mockWriter.Bytes(), expected) {
        t.Errorf("Expected %v, got %v", expected, mockWriter.Bytes())
    }
}


// Test generated using Keploy
func TestWriteBytes_CorrectOutput(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    writer := NewWriter(mockWriter, binary.LittleEndian, true)

    input := []byte{0x01, 0x02, 0x03}
    err := writer.WriteBytes(input)
    if err != nil {
        t.Errorf("WriteBytes returned an error: %v", err)
    }

    if !bytes.Equal(mockWriter.Bytes(), input) {
        t.Errorf("Expected %v, got %v", input, mockWriter.Bytes())
    }
}


// Test generated using Keploy
func TestWriteFloat32_CorrectOutput(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    writer := NewWriter(mockWriter, binary.LittleEndian, true)

    input := float32(3.14)
    err := writer.WriteFloat32(input)
    if err != nil {
        t.Errorf("WriteFloat32 returned an error: %v", err)
    }

    expected := make([]byte, 4)
    binary.LittleEndian.PutUint32(expected, math.Float32bits(input))
    if !bytes.Equal(mockWriter.Bytes(), expected) {
        t.Errorf("Expected %v, got %v", expected, mockWriter.Bytes())
    }
}


// Test generated using Keploy
func TestWriteFloat64_CorrectOutput(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    writer := NewWriter(mockWriter, binary.LittleEndian, true)

    input := float64(3.14159265359)
    err := writer.WriteFloat64(input)
    if err != nil {
        t.Errorf("WriteFloat64 returned an error: %v", err)
    }

    expected := make([]byte, 8)
    binary.LittleEndian.PutUint64(expected, math.Float64bits(input))
    if !bytes.Equal(mockWriter.Bytes(), expected) {
        t.Errorf("Expected %v, got %v", expected, mockWriter.Bytes())
    }
}


// Test generated using Keploy
func TestWriteUInt32_CorrectOutput(t *testing.T) {
    mockWriter := &bytes.Buffer{}
    writer := NewWriter(mockWriter, binary.LittleEndian, true)

    input := uint32(123456789)
    err := writer.WriteUInt32(input)
    if err != nil {
        t.Errorf("WriteUInt32 returned an error: %v", err)
    }

    expected := make([]byte, 4)
    binary.LittleEndian.PutUint32(expected, input)
    if !bytes.Equal(mockWriter.Bytes(), expected) {
        t.Errorf("Expected %v, got %v", expected, mockWriter.Bytes())
    }
}

