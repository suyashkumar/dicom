package dcmtime_test

import (
	"github.com/suyashkumar/dicom/pkg/dcmtime"
	"testing"
)

// precisionRange defines the (inclusive) minimum and maximum precision to be expected.
type precisionRange struct {
	Min dcmtime.PrecisionLevel
	Max dcmtime.PrecisionLevel
}

// Contains returns true if a value falls within the given range (inclusive).
func (pRange precisionRange) Contains(val dcmtime.PrecisionLevel) bool {
	if val > pRange.Min {
		return false
	}

	if val < pRange.Max {
		return false
	}
	// If this value falls within the omit range, it is false.
	return true
}

// precisionChecker defines an interface for types that can check their precision.
type precisionChecker interface {
	HasPrecision(check dcmtime.PrecisionLevel) bool
}

// checkHasPrecision checks that we get the expected results from a type with
// hasPrecisionOmits
func checkHasPrecision(t *testing.T, value precisionChecker, expectedRange precisionRange, omits precisionRange) {
	for p := dcmtime.PrecisionFull; p <= dcmtime.PrecisionYear; p++ {
		expected := expectedRange.Contains(p) && !omits.Contains(p)
		result := value.HasPrecision(p)

		t.Logf("Has Precision %v: %v", p, result)

		if result != expected {
			t.Errorf(
				"expected value '%v' HasPrecision() to be '%v' for precision '%v', got '%v'",
				value,
				expected,
				p,
				result,
			)
		}
	}
}
