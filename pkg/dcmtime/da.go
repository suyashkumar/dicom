package dcmtime

import (
	"fmt"
	"time"
)

// DA holds data for a parsed DICOM date (DA value).
type DA struct {
	// Value as a native go time.Time value.
	Time time.Time
	// The precision with this value was stored. For instance, a DA value with a
	// precision of Precision.Year ONLY stored the year.
	Precision PrecisionLevel
}

// DCM converts time.Time value to dicom DA string. Values are truncated to the
// DA.Precision value.
//
// NOTE: Time zones are ignored in this operation, as DA does not support encoding them.
// Make sure values are converted to UTC before passing if that is the desired output.
func (da DA) DCM() string {
	year, month, day := da.Time.Date()

	daVal := fmt.Sprintf("%04d", year)
	if !isIncluded(Precision.Month, da.Precision) {
		return daVal
	}

	daVal += fmt.Sprintf("%02d", month)
	if !isIncluded(Precision.Day, da.Precision) {
		return daVal
	}

	daVal += fmt.Sprintf("%02d", day)
	return daVal
}

// String implements fmt.Stringer.
func (da DA) String() string {
	return da.DCM()
}

// NewDA creates new DA value from a given time.Time.
//
// precision is the last element to be included in the DICOM DA.String() value.
// Precision.Full will include all possible values.
func NewDA(timeVal time.Time, precision PrecisionLevel) DA {
	return DA{
		Time:      timeVal,
		Precision: precision,
	}
}
