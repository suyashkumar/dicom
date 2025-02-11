package uid

import (
    "testing"
    "encoding/binary"
)


// Test generated using Keploy
func TestCanonicalTransferSyntaxUID_ValidStandardUID(t *testing.T) {
    validUIDs := []string{
        ImplicitVRLittleEndian,
        ExplicitVRLittleEndian,
        ExplicitVRBigEndian,
        DeflatedExplicitVRLittleEndian,
    }

    for _, uid := range validUIDs {
        result, err := CanonicalTransferSyntaxUID(uid)
        if err != nil {
            t.Errorf("Expected no error for UID %v, got %v", uid, err)
        }
        if result != uid {
            t.Errorf("Expected %v, got %v", uid, result)
        }
    }
}

// Test generated using Keploy
func TestCanonicalTransferSyntaxUID_InvalidUID(t *testing.T) {
    invalidUID := "1.2.3.4.5.6.7.8.9"
    _, err := CanonicalTransferSyntaxUID(invalidUID)
    if err == nil {
        t.Errorf("Expected an error for invalid UID %v, but got none", invalidUID)
    }
}


// Test generated using Keploy
func TestParseTransferSyntaxUID_ValidUID(t *testing.T) {
    testCases := []struct {
        uid       string
        expectedBO binary.ByteOrder
        expectedVR bool
    }{
        {ImplicitVRLittleEndian, binary.LittleEndian, true},
        {ExplicitVRLittleEndian, binary.LittleEndian, false},
        {ExplicitVRBigEndian, binary.BigEndian, false},
        {DeflatedExplicitVRLittleEndian, binary.LittleEndian, false},
    }

    for _, tc := range testCases {
        bo, vr, err := ParseTransferSyntaxUID(tc.uid)
        if err != nil {
            t.Errorf("Expected no error for UID %v, got %v", tc.uid, err)
        }
        if bo != tc.expectedBO {
            t.Errorf("Expected byte order %v for UID %v, got %v", tc.expectedBO, tc.uid, bo)
        }
        if vr != tc.expectedVR {
            t.Errorf("Expected VR %v for UID %v, got %v", tc.expectedVR, tc.uid, vr)
        }
    }
}


// Test generated using Keploy
func TestParseTransferSyntaxUID_InvalidUID(t *testing.T) {
    invalidUID := "1.2.3.4.5.6.7.8.9"
    _, _, err := ParseTransferSyntaxUID(invalidUID)
    if err == nil {
        t.Errorf("Expected an error for invalid UID %v, but got none", invalidUID)
    }
}


// Test generated using Keploy
func TestParseTransferSyntaxUID_NonTransferSyntaxUID(t *testing.T) {
    nonTransferSyntaxUID := "1.2.840.10008.1.1" // Example of a valid but non-transfer syntax UID
    _, _, err := ParseTransferSyntaxUID(nonTransferSyntaxUID)
    if err == nil {
        t.Errorf("Expected an error for non-transfer syntax UID %v, but got none", nonTransferSyntaxUID)
    }
}

