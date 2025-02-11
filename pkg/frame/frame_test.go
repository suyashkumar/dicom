package frame

import (
    "testing"

)


// Test generated using Keploy
func TestGetNativeFrame_Encapsulated_ReturnsError(t *testing.T) {
    mockEncapsulatedFrame := &EncapsulatedFrame{}
    frame := &Frame{
        Encapsulated: true,
        EncapsulatedData: *mockEncapsulatedFrame,
    }

    _, err := frame.GetNativeFrame()
    if err == nil {
        t.Fatalf("Expected an error, got nil")
    }
    if err != ErrorFrameTypeNotPresent {
        t.Errorf("Expected error %v, got %v", ErrorFrameTypeNotPresent, err)
    }
}

// Test generated using Keploy
func TestEquals_DifferentEncapsulation_ReturnsFalse(t *testing.T) {
    frame1 := &Frame{
        Encapsulated: false,
    }
    frame2 := &Frame{
        Encapsulated: true,
    }

    if frame1.Equals(frame2) {
        t.Errorf("Expected frames to not be equal, but they are")
    }
}


// Test generated using Keploy
func TestIsEncapsulated_True(t *testing.T) {
    frame := &Frame{
        Encapsulated: true,
    }

    if !frame.IsEncapsulated() {
        t.Errorf("Expected IsEncapsulated to return true, but got false")
    }
}


// Test generated using Keploy
func TestGetEncapsulatedFrame_Encapsulated_ReturnsFrame(t *testing.T) {
    // Create a mock encapsulated frame
    mockEncapsulatedFrame := EncapsulatedFrame{
        Data: []byte{0x01, 0x02, 0x03},
    }

    frame := &Frame{
        Encapsulated:     true,
        EncapsulatedData: mockEncapsulatedFrame,
    }

    result, err := frame.GetEncapsulatedFrame()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    if !result.Equals(&mockEncapsulatedFrame) {
        t.Errorf("Expected frames to be equal, but they are not")
    }
}


// Test generated using Keploy
func TestEquals_DifferentData_ReturnsFalse(t *testing.T) {
    frame1 := &Frame{
        Encapsulated: true,
        EncapsulatedData: EncapsulatedFrame{
            Data: []byte{0x01, 0x02, 0x03},
        },
    }
    frame2 := &Frame{
        Encapsulated: true,
        EncapsulatedData: EncapsulatedFrame{
            Data: []byte{0x04, 0x05, 0x06},
        },
    }

    if frame1.Equals(frame2) {
        t.Errorf("Expected frames to not be equal, but they are")
    }
}


// Test generated using Keploy
func TestEquals_OneFrameNil_ReturnsFalse(t *testing.T) {
    frame1 := &Frame{
        Encapsulated: true,
    }
    var frame2 *Frame = nil

    if frame1.Equals(frame2) {
        t.Errorf("Expected frames to not be equal, but they are")
    }
}


// Test generated using Keploy
func TestGetImage_Encapsulated_Error(t *testing.T) {
    mockEncapsulatedFrame := &EncapsulatedFrame{
        Data: []byte{0x01, 0x02, 0x03},
    }
    frame := &Frame{
        Encapsulated:     true,
        EncapsulatedData: *mockEncapsulatedFrame,
    }

    _, err := frame.GetImage()
    if err == nil {
        t.Fatalf("Expected an error, got nil")
    }
}

