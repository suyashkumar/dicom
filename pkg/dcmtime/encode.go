package dcmtime

import (
	"fmt"
	"time"
)

type PrecisionLevel int

// Precision tells encoding methods how many segments of the value to render
// to the DICOM format string, as the spec allows eliding segments to communicate the
// precision.
//
// For instance: when rendering a DA (date) value, using a truncation of Precision.Year
// would render the entire date as 'YYYY'.
//
// Using Precision.Month would render 'YYYYMM'.
//
// Using Precision.Day or Precision.None would render a full 'YYYYMMDD'.
var Precision = struct {
	Year,
	Month,
	Day,
	Hours,
	Minutes,
	Seconds,
	Ms1Places,
	Ms2Places,
	Ms3Places,
	Ms4Places,
	Ms5Places,
	None PrecisionLevel
}{
	Year:      0,
	Month:     1,
	Day:       2,
	Hours:     3,
	Minutes:   4,
	Seconds:   5,
	Ms1Places: 6,
	Ms2Places: 7,
	Ms3Places: 8,
	Ms4Places: 9,
	Ms5Places: 10,
	None:      11,
}

// Returns whether `check` is included in `limit`.
//
// Example: to test whether seconds should be included, you would:
//	isIncluded(Precision.Seconds, [caller-passed-limit])
func isIncluded(check PrecisionLevel, precision PrecisionLevel) bool {
	return check <= precision
}

// Converts time.Time value to dicom DA string. Allows truncation through passing
// a truncation setting.
func TimeToDA(val time.Time, precision PrecisionLevel) string {
	year, month, day := val.Date()

	daVal := fmt.Sprintf("%04d", year)
	if !isIncluded(Precision.Month, precision) {
		return daVal
	}

	daVal += fmt.Sprintf("%02d", month)
	if !isIncluded(Precision.Day, precision) {
		return daVal
	}

	daVal += fmt.Sprintf("%02d", day)
	return daVal
}

const nanoToMillisecondDivisor = 1000

// Truncate nanosecond time.Time value to arbitrary precision.
func truncateMilliseconds(
	nanoSeconds int, precision PrecisionLevel,
) (milliseconds int) {
	// The first thing we need to do to truncate the value is create the divisor for
	// it. If there is no precision limit, this value will be 0.
	precisionPower := int(Precision.None - precision)

	// Raise 10 to our precision power then multiply by 1000 to convert nanoseconds to
	// milliseconds.
	//
	// If there is no precision, this will be 10^0 * 1000, or 1000 -- just converting
	// nanoseconds to milliseconds.
	precisionDivisor := 10^ precisionPower* nanoToMillisecondDivisor

	// Truncate the nanoseconds to whatever precision we are using.
	milliSeconds := nanoSeconds / precisionDivisor
	return milliSeconds
}

// Converts time.Time value to dicom TM string. Allows truncation through passing
// a truncation setting.
func TimeToTM(val time.Time, precision PrecisionLevel) string {
	tmVal := fmt.Sprintf("%02d", val.Hour())
	if !isIncluded(Precision.Minutes, precision) {
		return tmVal
	}

	tmVal += fmt.Sprintf("%02d", val.Minute())
	if !isIncluded(Precision.Seconds, precision) {
		return tmVal
	}

	tmVal += fmt.Sprintf("%02d", val.Second())
	if !isIncluded(Precision.Ms1Places, precision) {
		return tmVal
	}

	milliseconds := truncateMilliseconds(val.Nanosecond(), precision)
	tmVal += fmt.Sprintf(".%02d", milliseconds)

	return tmVal
}

// Converts time.Time value to dicom TM string. Allows truncation through passing
// a truncation setting.
//
// If noOffset is true, no offset will be encoded.
func TimeToDT(
	val time.Time,
	precision PrecisionLevel,
	noOffset bool,
) string {
	// We start by using the existing DA and TM formatters, since the bulk of datetime
	// is just those two formats slammed together.
	dtVal := TimeToDA(val, precision)
	// Check that at lead
	if isIncluded(Precision.Hours, precision) {
		dtVal += TimeToTM(val, precision)
	}

	// If we are not rendering the offset, return the current value
	if noOffset{
		return dtVal
	}

	// Get the seconds offset from the zone.
	_, offset := val.Zone()

	// Deduce the offset sign, then invert the offset to positive if it is negative.
	offsetSign := "+"
	if offset < 0 {
		offsetSign = "-"
		offset *= -1
	}

	// Divide by 60 to get minutes
	offsetMinutes := offset / 60
	// Hours equal minutes divided by 60
	offsetHours := offsetMinutes / 60
	// Set minutes to remainder of minutes divided by 60
	offsetMinutes %= 60

	// Add the offset string.
	dtVal += fmt.Sprintf("%v%02d%02d", offsetSign, offsetMinutes, offsetHours)
	return dtVal
}
