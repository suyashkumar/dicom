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

// DCM converts internal time.Time value to dicom TM string, truncating the output
// to the DA value's Precision.
//
// NOTE: Time zones are ignored in this operation, as TM does not support encoding them.
// Make sure values are converted to UTC before passing if that is the desired output.
func (tm Time) DCM() string {
	builder := new(strings.Builder)

	builder.WriteString(fmt.Sprintf("%02d", tm.Time.Hour()))
	if !isIncluded(PrecisionMinutes, tm.Precision) {
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf("%02d", tm.Time.Minute()))
	if !isIncluded(PrecisionSeconds, tm.Precision) {
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf("%02d", tm.Time.Second()))
	if !isIncluded(PrecisionMS1, tm.Precision) {
		return builder.String()
	}

	builder.WriteRune('.')
	builder.WriteString(truncateMilliseconds(tm.Time.Nanosecond(), tm.Precision))

	return builder.String()
}

// String implements fmt.Stringer.
func (tm Time) String() string {
	builder := new(strings.Builder)

	builder.WriteString(fmt.Sprintf("%02d", tm.Time.Hour()))
	if !isIncluded(PrecisionMinutes, tm.Precision) {
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf(":%02d", tm.Time.Minute()))
	if !isIncluded(PrecisionSeconds, tm.Precision) {
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf(":%02d", tm.Time.Second()))
	if !isIncluded(PrecisionMS1, tm.Precision) {
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
	precisionOut = updatePrecision(precisionIn, hours, PrecisionHours, false)

	minutes, err = extractDurationInfo(matches, groupIndexes.Minutes, false)
	if err != nil {
		return hours, minutes, seconds, nanos, precisionOut, err
	}
	precisionOut = updatePrecision(precisionOut, minutes, PrecisionMinutes, false)

	seconds, err = extractDurationInfo(matches, groupIndexes.Seconds, false)
	if err != nil {
		return hours, minutes, seconds, nanos, precisionOut, err
	}
	precisionOut = updatePrecision(precisionOut, seconds, PrecisionSeconds, false)

	nanos, err = extractDurationInfo(matches, groupIndexes.Fractal, true)
	if err != nil {
		return hours, minutes, seconds, nanos, precisionOut, err
	} else if nanos.Present {
		precisionOut = nanos.FractalPrecision
	}

	return hours, minutes, seconds, nanos, precisionOut, err
}

// ParseTime converts DICOM TM (time) value to time.Time as UTC with PrecisionLevel.
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
