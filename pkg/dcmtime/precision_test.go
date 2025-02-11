package dcmtime

import (
    "testing"
)


// Test generated using Keploy
func TestPrecisionLevelString_ValidCases(t *testing.T) {
    testCases := []struct {
        level    PrecisionLevel
        expected string
    }{
        {PrecisionFull, "FULL"},
        {PrecisionYear, "YEAR"},
        {PrecisionMonth, "MONTH"},
        {PrecisionDay, "DAY"},
        {PrecisionHours, "HOURS"},
        {PrecisionMinutes, "MINUTES"},
        {PrecisionSeconds, "SECONDS"},
        {PrecisionMS1, "MS1"},
        {PrecisionMS2, "MS2"},
        {PrecisionMS3, "MS3"},
        {PrecisionMS4, "MS4"},
        {PrecisionMS5, "MS5"},
    }

    for _, tc := range testCases {
        t.Run(tc.expected, func(t *testing.T) {
            result := tc.level.String()
            if result != tc.expected {
                t.Errorf("Expected %v, got %v", tc.expected, result)
            }
        })
    }
}

// Test generated using Keploy
func TestPrecisionLevelString_UndefinedCase(t *testing.T) {
    undefinedLevel := PrecisionLevel(999) // An undefined PrecisionLevel
    expected := "!{PRECISION}"
    result := undefinedLevel.String()
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}

