package dcmtime

import (
	"fmt"
	"time"
)

// DT holds data for a parsed DICOM datetime (DT value).
type DT struct {
	// Value as a native go time.Time value.
	Time time.Time
	// The precision with this value was stored. For instance, a DT value with a
	// precision of Precision.Year ONLY stored the year.
	Precision PrecisionLevel
	// If true, offset information was specifically included in the DT string.
	IgnoreOffset bool
}

// String converts time.Time value to dicom TM string. Allows truncation through passing
// a truncation setting.
//
// If noOffset is true, no offset will be encoded.
func (dt DT) String() string {
	// We start by using the existing DA and TM formatters, since the bulk of datetime
	// is just those two formats slammed together.
	dtVal := NewDA(dt.Time, dt.Precision).String()
	// Check that at lead
	if isIncluded(Precision.Hours, dt.Precision) {
		dtVal += NewTM(dt.Time, dt.Precision).String()
	}

	// If we are not rendering the offset, return the current value
	if dt.IgnoreOffset {
		return dtVal
	}

	// Get the seconds offset from the zone.
	_, offset := dt.Time.Zone()

	// Deduce the offset sign, then invert the offset to positive if it is negative.
	offsetSign := "+"
	if offset < 0 {
		offsetSign = "-"
		offset *= -1
	}

	// Divide seconds by 60 to get minutes
	offsetMinutes := offset / 60
	// Hours equal minutes divided by 60
	offsetHours := offsetMinutes / 60
	// Set minutes to remainder of minutes divided by 60
	offsetMinutes %= 60

	// Add the offset string.
	dtVal += fmt.Sprintf("%v%02d%02d", offsetSign, offsetHours, offsetMinutes)
	return dtVal
}

// NewDT creates new DT value from a given time.Time.
//
// precision is the last element to be included in the DICOM .String() value.
// Precision.Full will include all possible values.
//
// if ignoreOffset is true, the offset will be not included in the DICOM .String()
// value.
func NewDT(timeVal time.Time, precision PrecisionLevel, ignoreOffset bool) DT {
	return DT{
		Time:         timeVal,
		Precision:    precision,
		IgnoreOffset: ignoreOffset,
	}
}
