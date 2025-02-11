package uid

// Test generated using Keploy
import (
    "testing"
)

func TestMustLookup_InvalidUID_Panics(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Fatalf("Expected panic, but code did not panic")
        }
    }()

    MustLookup("1.2.840.invalid")
}


// Test generated using Keploy
func TestUIDString_ValidUID_ReturnsFormattedString(t *testing.T) {
    uid := "1.2.840.10008.1.1"
    expected := "1.2.840.10008.1.1[Verification SOP Class]"

    result := UIDString(uid)
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}


// Test generated using Keploy
func TestUIDString_UnknownUID_ReturnsFormattedStringWithUnknown(t *testing.T) {
    uid := "1.2.840.unknown"
    expected := "1.2.840.unknown[??]"

    result := UIDString(uid)
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}
