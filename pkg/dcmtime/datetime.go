package dcmtime

import (
	"fmt"
	"time"
)

// Datetime holds data for a parsed DICOM datetime (DT) value.
type Datetime struct {
	// Time is a native go time.Time value.
	Time time.Time
	// Precision with this value was stored. For instance, a DT value with a
	// precision of Precision.Year ONLY stored the year.
	Precision PrecisionLevel
	// NoOffset: if true, offset information was not specifically included in the
	// original DT string, and will not be rendered with DCM()
	NoOffset bool
}

// DCM converts time.Time value to dicom DT string. Values are truncated to the
// DT.Precision value.
//
// If NoOffset is true, no offset will be encoded.
func (dt Datetime) DCM() string {
	// We start by using the existing DA and TM formatters, since the bulk of datetime
	// is just those two formats slammed together.
	dtVal := Date{Time: dt.Time, Precision: dt.Precision}.DCM()

	// Check that at lead
	if isIncluded(PrecisionHours, dt.Precision) {
		dtVal += Time{Time: dt.Time, Precision: dt.Precision}.DCM()
	}

	// If we are not rendering the offset, return the current value
	if dt.NoOffset {
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

// String implements fmt.Stringer.
func (dt Datetime) String() string {
	return dt.DCM()
}
