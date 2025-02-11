package tag

import (
    "testing"
)


// Test generated using Keploy
func TestVRKindString_ValidValues(t *testing.T) {
    testCases := []struct {
        input    VRKind
        expected string
    }{
        {0, "VRStringList"},
        {1, "VRBytes"},
        {2, "VRString"},
        {3, "VRUInt16List"},
        {4, "VRUInt32List"},
        {5, "VRInt16List"},
        {6, "VRInt32List"},
        {7, "VRFloat32List"},
        {8, "VRFloat64List"},
        {9, "VRSequence"},
        {10, "VRItem"},
        {11, "VRTagList"},
        {12, "VRDate"},
        {13, "VRPixelData"},
    }

    for _, tc := range testCases {
        result := tc.input.String()
        if result != tc.expected {
            t.Errorf("Expected %v, got %v", tc.expected, result)
        }
    }
}

// Test generated using Keploy
func TestVRKindString_InvalidValues(t *testing.T) {
    testCases := []struct {
        input    VRKind
        expected string
    }{
        {-1, "VRKind(-1)"},
        {14, "VRKind(14)"},
        {100, "VRKind(100)"},
    }

    for _, tc := range testCases {
        result := tc.input.String()
        if result != tc.expected {
            t.Errorf("Expected %v, got %v", tc.expected, result)
        }
    }
}

