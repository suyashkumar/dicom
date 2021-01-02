package dcmtime

import (
	"fmt"
	"time"
)

type TruncationLimit int

// Truncation tells encoding methods how many segments of the value to render
// to the DICOM format string, as the spec allows eliding segments to communicate the
// precision.
//
// For instance: when rendering a DA (date) value, using a truncation of Truncation.Year
// would render the entire date as 'YYYY'.
//
// Using Truncation.Month would render 'YYYYMM'.
//
// Using truncation.Day or Truncation.None would render a full 'YYYYMMDD'.
var Truncation = struct {
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
	None TruncationLimit
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
//	isIncluded(Truncation.Seconds, [caller-passed-limit])
func isIncluded(check TruncationLimit, limit TruncationLimit) bool {
	return check <= limit
}

// Converts time.Time value to dicom DA string. Allows truncation through passing
// a truncation setting.
func TimeToDA(val time.Time, truncation TruncationLimit) string {
	year, month, day := val.Date()

	daVal := fmt.Sprintf("%04d", year)
	if !isIncluded(Truncation.Month, truncation) {
		return daVal
	}

	daVal += fmt.Sprintf("%02d", month)
	if !isIncluded(Truncation.Day, truncation) {
		return daVal
	}

	daVal += fmt.Sprintf("%02d", day)
	return daVal
}

const nanoToMillisecondDivisor = 1000

// Truncate nanosecond time.Time value to arbitrary precision.
func truncateMilliseconds(
	nanoSeconds int, truncation TruncationLimit,
) (milliseconds int) {
	// The first thing we need to do to truncate the value is create the divisor for
	// it. If there is no truncations, this value will be 0.
	truncatePower := int(Truncation.None - truncation)

	// Raise 10 to our truncation power then multiply by 1000 to convert nanoseconds to
	// milliseconds.
	//
	// If there is no truncation, this will be 10^0 * 1000, or 1000 -- just converting
	// nanoseconds to milliseconds.
	truncateDivisor := 10^truncatePower * nanoToMillisecondDivisor

	// Truncate the nanoseconds to whatever precision we are using.
	milliSeconds := nanoSeconds / truncateDivisor
	return milliSeconds
}

// Converts time.Time value to dicom TM string. Allows truncation through passing
// a truncation setting.
func TimeToTM(val time.Time, truncation TruncationLimit) string {
	tmVal := fmt.Sprintf("%02d", val.Hour())
	if !isIncluded(Truncation.Minutes, truncation) {
		return tmVal
	}

	tmVal += fmt.Sprintf("%02d", val.Minute())
	if !isIncluded(Truncation.Seconds, truncation) {
		return tmVal
	}

	tmVal += fmt.Sprintf("%02d", val.Second())
	if !isIncluded(Truncation.Ms1Places, truncation) {
		return tmVal
	}

	milliseconds := truncateMilliseconds(val.Nanosecond(), truncation)
	tmVal += fmt.Sprintf(".%02d", milliseconds)

	return tmVal
}

// Converts time.Time value to dicom TM string. Allows truncation through passing
// a truncation setting.
//
// If noOffset is true, no offset will be encoded.
func TimeToDT(
	val time.Time,
	truncation TruncationLimit,
	noOffset bool,
) string {
	// We start by using the existing DA and TM formatters, since the bulk of datetime
	// is just those two formats slammed together.
	dtVal := TimeToDA(val, truncation)
	// Check that at lead
	if isIncluded(Truncation.Hours, truncation) {
		dtVal += TimeToTM(val, truncation)
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
