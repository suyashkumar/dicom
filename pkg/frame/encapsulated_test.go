package frame

import (
    "testing"
)


// Test generated using Keploy
func TestIsEncapsulated_AlwaysReturnsTrue(t *testing.T) {
    frame := &EncapsulatedFrame{}
    if !frame.IsEncapsulated() {
        t.Errorf("Expected IsEncapsulated to return true, but got false")
    }
}

// Test generated using Keploy
func TestGetEncapsulatedFrame_ReturnsSelfAndNoError(t *testing.T) {
    frame := &EncapsulatedFrame{Data: []byte{1, 2, 3}}
    result, err := frame.GetEncapsulatedFrame()
    if err != nil {
        t.Errorf("Expected no error, but got %v", err)
    }
    if result != frame {
        t.Errorf("Expected result to be the same as the original frame, but got a different instance")
    }
}


// Test generated using Keploy
func TestGetNativeFrame_ReturnsErrorFrameTypeNotPresent(t *testing.T) {
    frame := &EncapsulatedFrame{}
    result, err := frame.GetNativeFrame()
    if result != nil {
        t.Errorf("Expected result to be nil, but got %v", result)
    }
    if err != ErrorFrameTypeNotPresent {
        t.Errorf("Expected error to be ErrorFrameTypeNotPresent, but got %v", err)
    }
}


// Test generated using Keploy
func TestEquals_IdenticalFrames_ReturnsTrue(t *testing.T) {
    data := []byte{1, 2, 3}
    frame1 := &EncapsulatedFrame{Data: data}
    frame2 := &EncapsulatedFrame{Data: data}
    if !frame1.Equals(frame2) {
        t.Errorf("Expected Equals to return true for identical frames, but got false")
    }
}


// Test generated using Keploy
func TestEquals_DifferentFrames_ReturnsFalse(t *testing.T) {
    frame1 := &EncapsulatedFrame{Data: []byte{1, 2, 3}}
    frame2 := &EncapsulatedFrame{Data: []byte{4, 5, 6}}
    if frame1.Equals(frame2) {
        t.Errorf("Expected Equals to return false for different frames, but got true")
    }
}

