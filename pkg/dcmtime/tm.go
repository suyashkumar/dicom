package dcmtime

import (
	"fmt"
	"time"
)

// TM holds data for a parsed DICOM time (TM value).
type TM struct {
	// Value as a native go time.Time value.
	Time time.Time
	// The precision with this value was stored. For instance, a DA value with a
	// precision of Precision.Hours ONLY stored the hours value.
	Precision PrecisionLevel
}

// String converts internal time.Time value to dicom TM string, truncating the output
// to the DA value's Precision.
//
// NOTE: Time zones are ignored in this operation, as TM does not support encoding them.
// Make sure values are converted to UTC before passing if that is the desired output.
func (tm TM) String() string {
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

// NewTM creates new TM value from a given time.Time.
//
// precision is the last element to be included in the DICOM .String() value.
// Precision.Full will include all possible values.
func NewTM(timeVal time.Time, precision PrecisionLevel) TM {
	return TM{
		Time:      timeVal,
		Precision: precision,
	}
}
