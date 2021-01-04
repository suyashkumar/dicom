package dcmtime

import (
	"fmt"
	"time"
)

// Time holds data for a parsed DICOM time (TM) value.
type Time struct {
	// Time native go time.Time value.
	Time time.Time
	// Precision with which the raw TM value was stored. For instance, a Date value
	// with a precision of PrecisionHour ONLY stored the Hour.
	Precision PrecisionLevel
}

// DCM converts internal time.Time value to dicom TM string, truncating the output
// to the DA value's Precision.
//
// NOTE: Time zones are ignored in this operation, as TM does not support encoding them.
// Make sure values are converted to UTC before passing if that is the desired output.
func (tm Time) DCM() string {
	tmVal := fmt.Sprintf("%02d", tm.Time.Hour())
	if !isIncluded(PrecisionMinutes, tm.Precision) {
		return tmVal
	}

	tmVal += fmt.Sprintf("%02d", tm.Time.Minute())
	if !isIncluded(PrecisionSeconds, tm.Precision) {
		return tmVal
	}

	tmVal += fmt.Sprintf("%02d", tm.Time.Second())
	if !isIncluded(PrecisionMS1, tm.Precision) {
		return tmVal
	}

	tmVal += "." + truncateMilliseconds(tm.Time.Nanosecond(), tm.Precision)

	return tmVal
}

// String implements fmt.Stringer.
func (tm Time) String() string {
	return tm.DCM()
}
