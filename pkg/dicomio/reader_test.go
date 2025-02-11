package dicomio

// Test generated using Keploy
import (
    "bytes"
    "bufio"
    "encoding/binary"
    "testing"
    "io"
    "math"
)

func TestReader_Skip_ExceedsLimit(t *testing.T) {
    input := []byte("testdata")
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))

    err := reader.Skip(10)
    if err != ErrorInsufficientBytesLeft {
        t.Fatalf("expected ErrorInsufficientBytesLeft, got %v", err)
    }
}


// Test generated using Keploy
func TestReader_Read_SmallerBuffer(t *testing.T) {
    input := []byte("testdata")
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))
    
    buffer := make([]byte, 10) // Larger than input
    
    n, err := reader.Read(buffer)
    if err != nil && err != io.EOF {
        t.Fatalf("unexpected error: %v", err)
    }
    
    if n != len(input) {
        t.Errorf("expected to read %d bytes, got %d", len(input), n)
    }
    
    if string(buffer[:n]) != "testdata" {
        t.Errorf("expected buffer to contain 'testdata', got '%s'", string(buffer[:n]))
    }
}


// Test generated using Keploy
func TestReader_Skip_WithinLimit(t *testing.T) {
    input := []byte("testdata")
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))
    
    err := reader.Skip(4)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    
    buffer := make([]byte, 4)
    n, err := reader.Read(buffer)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    
    if n != 4 {
        t.Errorf("expected to read 4 bytes, got %d", n)
    }
    
    if string(buffer) != "data" {
        t.Errorf("expected buffer to contain 'data', got '%s'", string(buffer))
    }
}


// Test generated using Keploy
func TestReader_ReadUInt16_LittleEndian(t *testing.T) {
    input := []byte{0x01, 0x02}
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))
    
    value, err := reader.ReadUInt16()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    
    if value != 0x0201 {
        t.Errorf("expected value to be 0x0201, got 0x%04x", value)
    }
}


// Test generated using Keploy
func TestReader_ReadUInt32_LittleEndian(t *testing.T) {
    input := []byte{0x01, 0x02, 0x03, 0x04}
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))

    value, err := reader.ReadUInt32()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if value != 0x04030201 {
        t.Errorf("expected value to be 0x04030201, got 0x%08x", value)
    }
}


// Test generated using Keploy
func TestReader_ReadFloat32_LittleEndian(t *testing.T) {
    input := make([]byte, 4)
    binary.LittleEndian.PutUint32(input, math.Float32bits(3.14))
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))

    value, err := reader.ReadFloat32()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if value != 3.14 {
        t.Errorf("expected value to be 3.14, got %f", value)
    }
}


// Test generated using Keploy
func TestReader_ReadString(t *testing.T) {
    input := []byte("Hello, DICOM!")
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))

    s, err := reader.ReadString(uint32(len(input)))
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if s != "Hello, DICOM!" {
        t.Errorf("expected string 'Hello, DICOM!', got '%s'", s)
    }
}


// Test generated using Keploy
func TestReader_ReadFloat64_LittleEndian(t *testing.T) {
    input := make([]byte, 8)
    binary.LittleEndian.PutUint64(input, math.Float64bits(3.1415926535))
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))

    value, err := reader.ReadFloat64()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if value != 3.1415926535 {
        t.Errorf("expected value to be 3.1415926535, got %f", value)
    }
}


// Test generated using Keploy
func TestReader_ReadInt16_LittleEndian(t *testing.T) {
    input := []byte{0xFF, 0x7F} // 0x7FFF -> 32767
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))

    value, err := reader.ReadInt16()
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if value != 32767 {
        t.Errorf("expected value to be 32767, got %d", value)
    }
}


// Test generated using Keploy
func TestReader_Peek_CorrectBytes(t *testing.T) {
    input := []byte("testdata")
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))

    peeked, err := reader.Peek(4)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if string(peeked) != "test" {
        t.Errorf("expected 'test', got '%s'", string(peeked))
    }

    // Ensure the reader position has not advanced
    buffer := make([]byte, 4)
    n, err := reader.Read(buffer)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if n != 4 || string(buffer) != "test" {
        t.Errorf("expected to read 'test', got '%s'", string(buffer))
    }
}


// Test generated using Keploy
func TestReader_PushPopLimit(t *testing.T) {
    input := []byte("testdata")
    reader := NewReader(bufio.NewReader(bytes.NewReader(input)), binary.LittleEndian, int64(len(input)))

    err := reader.PushLimit(4) // Limit to the first 4 bytes
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    buffer := make([]byte, 4)
    n, err := reader.Read(buffer)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if n != 4 || string(buffer) != "test" {
        t.Errorf("expected to read 'test', got '%s'", string(buffer))
    }

    reader.PopLimit() // Restore original limit

    buffer = make([]byte, 4)
    n, err = reader.Read(buffer)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if n != 4 || string(buffer) != "data" {
        t.Errorf("expected to read 'data', got '%s'", string(buffer))
    }
}
