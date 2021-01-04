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
	if !isIncluded(Precision.Minutes, tm.Precision) {
		return tmVal
	}

	tmVal += fmt.Sprintf("%02d", tm.Time.Minute())
	if !isIncluded(Precision.Seconds, tm.Precision) {
		return tmVal
	}

	tmVal += fmt.Sprintf("%02d", tm.Time.Second())
	if !isIncluded(Precision.MS1, tm.Precision) {
		return tmVal
	}

	tmVal += "." + truncateMilliseconds(tm.Time.Nanosecond(), tm.Precision)

	return tmVal
}

// String implements fmt.Stringer.
func (tm Time) String() string {
	return tm.DCM()
}

// NewTM creates new TM value from a given time.Time.
//
// precision is the last element to be included in the DICOM TM.DCM() value.
// Precision.Full will include all possible values.
func NewTM(timeVal time.Time, precision PrecisionLevel) Time {
	return Time{
		Time:      timeVal,
		Precision: precision,
	}
}
