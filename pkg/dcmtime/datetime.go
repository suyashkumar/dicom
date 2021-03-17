package dcmtime

import (
	"strings"
	"time"
)

// Datetime holds data for a parsed DICOM datetime (DT) value.
type Datetime struct {
	// Time is the underlying time.Time value.
	Time time.Time
	// Precision with which this value was stored. For instance, a DT value with a
	// precision of PrecisionYear ONLY stored the year.
	Precision PrecisionLevel
	// NoOffset: if true, offset information was not specifically included in the
	// original DT string, and will not be rendered with DCM()
	NoOffset bool
}

// Year returns the underlying Time.Year(). Since a DICOM DT value must contain a year,
// presence is not reported.
func (dt Datetime) Year() int {
	return dt.Time.Year()
}

// Month returns the underlying Time.Month(), and a boolean indicating whether the
// original DICOM value included the month.
func (dt Datetime) Month() (month time.Month, ok bool) {
	return dt.Time.Month(), hasPrecision(PrecisionMonth, dt.Precision)
}

// Day returns the underlying time.Month, and a boolean indicating whether the
func (dt Datetime) Day() (month int, ok bool) {
	return dt.Time.Day(), hasPrecision(PrecisionDay, dt.Precision)
}

// Hour returns the underlying Time.Hour(). Since a DICOM TM value must contain an hour,
// presence is not reported.
func (dt Datetime) Hour() (hour int, ok bool) {
	return dt.Time.Hour(), hasPrecision(PrecisionHours, dt.Precision)
}

// Minute returns the underlying Time.Minute(), and a boolean indicating whether the
// original DICOM value included minutes.
func (dt Datetime) Minute() (minute int, ok bool) {
	return dt.Time.Minute(), hasPrecision(PrecisionMinutes, dt.Precision)
}

// Second returns the underlying Time.Second(), and a boolean indicating whether the
// original DICOM value included seconds.
func (dt Datetime) Second() (second int, ok bool) {
	return dt.Time.Second(), hasPrecision(PrecisionSeconds, dt.Precision)
}

// Nanosecond returns the underlying Time.Nanosecond(), and a boolean indicating whether
// the original DICOM value included any fractal seconds.
func (dt Datetime) Nanosecond() (second int, ok bool) {
	return dt.Time.Nanosecond(), hasPrecision(PrecisionMS1, dt.Precision)
}

// Location returns the underlying Time.Location(), and a boolean indicating whether
// the original DICOM value included any timezone offset.
func (dt Datetime) Location() (location *time.Location, ok bool) {
	return dt.Time.Location(), !dt.NoOffset
}

// DCM converts time.Time value to dicom DT string. Values are truncated to the
// DT.Precision value.
//
// If NoOffset is true, no offset will be encoded.
func (dt Datetime) DCM() string {
	builder := strings.Builder{}

	// We start by using the existing DA and TM formatters, since the bulk of datetime
	// is just those two formats slammed together.
	builder.WriteString(Date{Time: dt.Time, Precision: dt.Precision}.DCM())

	// Check that at lead
	if hasPrecision(PrecisionHours, dt.Precision) {
		builder.WriteString(Time{Time: dt.Time, Precision: dt.Precision}.DCM())
	}

	// If we are not rendering the offset, return the current value
	if dt.NoOffset {
		return builder.String()
	}

	// Write the timezone info.
	builder.WriteString(dt.Time.Format("-0700"))

	return builder.String()
}

// String implements fmt.Stringer.
func (dt Datetime) String() string {
	builder := strings.Builder{}

	// We start by using the existing DA and TM formatters, since the bulk of datetime
	// is just those two formats slammed together.
	builder.WriteString(Date{Time: dt.Time, Precision: dt.Precision}.String())

	// Check that at lead
	if hasPrecision(PrecisionHours, dt.Precision) {
		builder.WriteRune(' ')
		builder.WriteString(Time{Time: dt.Time, Precision: dt.Precision}.String())
	}

	// If we are not rendering the offset, return the current value
	if dt.NoOffset {
		return builder.String()
	}

	// Write the timezone info.
	builder.WriteString(dt.Time.Format(" -07:00"))

	return builder.String()
}

// ParseDatetime converts DICOM DT (datetime) value to time.Time, PrecisionLevel, and
// offset presence as UTC.
func ParseDatetime(dtString string) (Datetime, error) {
	matches := dtRegex.FindStringSubmatch(dtString)
	if len(matches) == 0 {
		return Datetime{}, ErrParseDT
	}

	year, month, day, precision, err := extractDate(matches, PrecisionFull, false)
	if err != nil {
		return Datetime{}, err
	}

	hours, minutes, seconds, nanos, precision, err := extractTime(
		matches, precision, false,
	)
	if err != nil {
		return Datetime{}, err
	}

	var hasOffset bool

	offsetHours, err := extractDurationInfo(matches, dtRegexGroupOffsetHours, false)
	if err != nil {
		return Datetime{}, ErrParseDT
	}
	// If hours are not present, there is either no offset or the regex will fail,
	// so we only need to check this here.
	if offsetHours.PresentInSource {
		hasOffset = true
	}

	offsetMinutes, err := extractDurationInfo(matches, dtRegexGroupOffsetMinutes, false)
	if err != nil {
		return Datetime{}, ErrParseDT
	}

	offsetSign := 1
	// If the zone sign is '-', then we need to multiply the offset by -1
	if matches[dtRegexGroupOffsetSign] == "-" {
		offsetSign = -1
	}

	// Get zone offset in seconds.
	offset := (offsetHours.Value*3600 + offsetMinutes.Value*60) * offsetSign

	// Return a new time.Time with the given values and UTC encoding.
	parsed := time.Date(
		year.Value,
		time.Month(month.Value),
		day.Value,
		hours.Value,
		minutes.Value,
		seconds.Value,
		nanos.Value,
		time.FixedZone("", offset),
	)

	return Datetime{
		Time:      parsed,
		Precision: precision,
		NoOffset:  hasOffset,
	}, nil
}
