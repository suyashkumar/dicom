package dcmtime

import (
	"fmt"
	"strings"
	"time"
)

// Time holds data for a parsed DICOM time (TM) value.
type Time struct {
	// Time is the underlying time.Time value.
	Time time.Time
	// Precision with which the raw TM value was stored. For instance, a Date value
	// with a precision of PrecisionHours ONLY stored the Hour.
	Precision PrecisionLevel
}

// GetTime returns the Time field value for the Time. Included to support common
// interfaces with other dcmtime types.
func (tm Time) GetTime() time.Time {
	return tm.Time
}

// GetPrecision returns the Precision field value for the Time. Included to support
// common  interfaces with other dcmtime types.
func (tm Time) GetPrecision() PrecisionLevel {
	return tm.Precision
}

// tmPrecisionOmits is the range of precision values not relevant to Time.
var tmPrecisionOmits = precisionRange{
	Min: PrecisionYear,
	Max: PrecisionDay,
}

// HasPrecision returns whether this da value has a precision of AT LEAST 'check'.
//
// Will always Return false for PrecisionYear, PrecisionMonth, and PrecisionDay.
//
// Will return true for PrecisionFull if all possible values are present.
func (tm Time) HasPrecision(check PrecisionLevel) bool {
	return hasPrecisionOmits(check, tm.Precision, tmPrecisionOmits)
}

// Hour returns the underlying Time.Hour(). Since a DICOM TM value must contain an hour,
// 'ok' will always be true. 'ok' is included to form a common interface with Datetime.
func (tm Time) Hour() (hour int, ok bool) {
	return tm.Time.Hour(), true
}

// Minute returns the underlying Time.Minute(), and a boolean indicating whether the
// original DICOM value included minutes.
func (tm Time) Minute() (minute int, ok bool) {
	return tm.Time.Minute(), hasPrecision(PrecisionMinutes, tm.Precision)
}

// Second returns the underlying Time.Second(), and a boolean indicating whether the
// original DICOM value included seconds.
func (tm Time) Second() (second int, ok bool) {
	return tm.Time.Second(), hasPrecision(PrecisionSeconds, tm.Precision)
}

// Nanosecond returns the underlying Time.Nanosecond(), and a boolean indicating whether
// the original DICOM value included any fractal seconds.
func (tm Time) Nanosecond() (second int, ok bool) {
	return tm.Time.Nanosecond(), hasPrecision(PrecisionMS1, tm.Precision)
}

// Combine combines the Time with a Date value into a single Datetime value.
//
// The Date value must have a PrecisionLevel of PrecisionFull or the method will fail.
//
// If no location is given, time.FixedZone("", 0) will be used and NoOffset will be
// set to 'true'.
func (tm Time) Combine(da Date, location *time.Location) (Datetime, error) {
	return combineDateAndTime(da, tm, location)
}

// DCM converts internal time.Time value to dicom TM string, truncating the output
// to the DA value's Precision.
//
// NOTE: Time zones are ignored in this operation, as TM does not support encoding them.
// Make sure values are converted to UTC before passing if that is the desired output.
func (tm Time) DCM() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("%02d", tm.Time.Hour()))
	if !hasPrecision(PrecisionMinutes, tm.Precision) {
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf("%02d", tm.Time.Minute()))
	if !hasPrecision(PrecisionSeconds, tm.Precision) {
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf("%02d", tm.Time.Second()))
	if !hasPrecision(PrecisionMS1, tm.Precision) {
		return builder.String()
	}

	builder.WriteRune('.')
	builder.WriteString(truncateMilliseconds(tm.Time.Nanosecond(), tm.Precision))

	return builder.String()
}

// String implements fmt.Stringer.
func (tm Time) String() string {
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("%02d", tm.Time.Hour()))
	if !hasPrecision(PrecisionMinutes, tm.Precision) {
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf(":%02d", tm.Time.Minute()))
	if !hasPrecision(PrecisionSeconds, tm.Precision) {
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf(":%02d", tm.Time.Second()))
	if !hasPrecision(PrecisionMS1, tm.Precision) {
		return builder.String()
	}

	builder.WriteRune('.')
	builder.WriteString(truncateMilliseconds(tm.Time.Nanosecond(), tm.Precision))

	return builder.String()
}

// Holds group indexes for a given source types regex.
type timeRegexGroups struct {
	Hours   int
	Minutes int
	Seconds int
	Fractal int
}

// Cached info on TM time regex groups.
var timeRegexGroupsTM = timeRegexGroups{
	Hours:   tmRegexGroupHours,
	Minutes: tmRegexGroupMinutes,
	Seconds: tmRegexGroupSeconds,
	Fractal: tmRegexGroupFractal,
}

// Cached info on DT time regex groups.
var timeRegexGroupsDT = timeRegexGroups{
	Hours:   dtRegexGroupHours,
	Minutes: dtRegexGroupMinutes,
	Seconds: dtRegexGroupSeconds,
	Fractal: dtRegexGroupFractal,
}

func extractTime(
	matches []string,
	precisionIn PrecisionLevel,
	isTM bool,
) (
	hours durationInfo,
	minutes durationInfo,
	seconds durationInfo,
	nanos durationInfo,
	precisionOut PrecisionLevel,
	err error,
) {
	groupIndexes := timeRegexGroupsTM
	if !isTM {
		groupIndexes = timeRegexGroupsDT
	}

	hours, err = extractDurationInfo(matches, groupIndexes.Hours, false)
	if err != nil {
		return hours, minutes, seconds, nanos, precisionOut, err
	}
	precisionOut = updatePrecision(hours, precisionIn, PrecisionHours, false)

	minutes, err = extractDurationInfo(matches, groupIndexes.Minutes, false)
	if err != nil {
		return hours, minutes, seconds, nanos, precisionOut, err
	}
	precisionOut = updatePrecision(minutes, precisionOut, PrecisionMinutes, false)

	seconds, err = extractDurationInfo(matches, groupIndexes.Seconds, false)
	if err != nil {
		return hours, minutes, seconds, nanos, precisionOut, err
	}
	precisionOut = updatePrecision(seconds, precisionOut, PrecisionSeconds, false)

	nanos, err = extractDurationInfo(matches, groupIndexes.Fractal, true)
	if err != nil {
		return hours, minutes, seconds, nanos, precisionOut, err
	}

	if nanos.PresentInSource {
		precisionOut = nanos.FractalPrecision
	}

	return hours, minutes, seconds, nanos, precisionOut, err
}

// ParseTime converts DICOM TM (time) value to dcmtime.Time.
func ParseTime(tmString string) (Time, error) {
	matches := tmRegex.FindStringSubmatch(tmString)
	// If no full match is found, return an error
	if len(matches) == 0 {
		return Time{}, ErrParseTM
	}

	hours, minutes, seconds, nanos, precision, err := extractTime(
		matches, PrecisionFull, true,
	)
	if err != nil {
		return Time{}, err
	}

	// Return a new time.Time with the given values and UTC encoding.
	parsed := time.Date(
		// If these are set to 0, they try to roll back to the last
		// date that makes sense, so we are going to set all
		// of them to 1.
		1,
		1,
		1,
		hours.Value,
		minutes.Value,
		seconds.Value,
		nanos.Value,
		zeroTimezone,
	)

	return Time{
		Time:      parsed,
		Precision: precision,
	}, nil
}
