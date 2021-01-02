package dcmtime

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func extractDurationInfo(subMatches []string, index int, fractal bool) (int, error) {
	if len(subMatches) <= index {
		return 0, errors.New("not enough sub-matches")
	}

	valueStr := subMatches[index]
	// If there was no match for the specific subgroup, this value is 0.
	if valueStr == "" {
		valueStr = "0"
	}

	// If this is a fractal seconds value, we need to pad it out to nanoseconds for go.
	if fractal {
		// The fractal value can be any length, with truncation implying a loss of
		// precision, we need to add trailing zeros to the hundred-millionth place to
		// get our nano-seconds.
		missingPlaces := 9 - len(valueStr)
		valueStr = valueStr + strings.Repeat("0", missingPlaces)
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}

	return value, nil
}

// Checks if a regex result has matches
func hasMatches(matches []string, original string) bool {
	// There must be at least one full match
	if len(matches) < 1 {
		return false
	}

	// If the full match is not the entire original string, then it is not a match.
	if matches[0] != original {
		return false
	}

	// Otherwise it's good.
	return true
}

// Converts DICOM DA (date) value to time.Time as UTC.
//
// If allowNEMA is true, the old NEMA 300 specification of YYYY.MM.DD will also be
// checked if parsing as YYYYMMDD fails.
func DAToTime(daString string, allowNEMA bool) (time.Time, error) {
	var matches []string

	// Iterate over our two regexes and attempt them.
	var matchesFound bool
	for _, thisRegex := range [2]*regexp.Regexp{daRegex, daRegexNema} {
		matches = thisRegex.FindStringSubmatch(daString)
		if hasMatches(matches, daString) {
			// If we find a match, mark it as found and, break out.
			matchesFound = true
			break
		} else if !allowNEMA {
			// If we are not allowing parsing based on the NEMA format, break without
			// marking we have found a match
			break
		}
	}

	if !matchesFound {
		return time.Time{}, ErrParseDate
	}

	year, err := extractDurationInfo(matches, daRegexYear, false)
	if err != nil {
		return time.Time{}, ErrParseDate
	}

	month, err := extractDurationInfo(matches, daRegexMonth, false)
	if err != nil {
		return time.Time{}, ErrParseDate
	}

	day, err := extractDurationInfo(matches, daRegexDay, false)
	if err != nil {
		return time.Time{}, ErrParseDate
	}

	// Return a new time.Time with the given values and UTC encoding.
	return time.Date(
		year, time.Month(month), day, 0, 0, 0, 0, time.UTC,
	), nil
}

// Converts DICOM DT (datetime) value to time.Time as UTC.
func DTToTime(dtString string) (time.Time, error) {
	matches := dtRegex.FindStringSubmatch(dtString)
	if !hasMatches(matches, dtString) {
		return time.Time{}, ErrParseDatetime
	}

	year, err := extractDurationInfo(matches, dtRegexYear, false)
	if err != nil {
		return time.Time{}, ErrParseDatetime
	}

	month, err := extractDurationInfo(matches, dtRegexMonth, false)
	if err != nil {
		return time.Time{}, ErrParseDatetime
	}

	day, err := extractDurationInfo(matches, dtRegexDay, false)
	if err != nil {
		return time.Time{}, ErrParseDatetime
	}

	hours, err := extractDurationInfo(matches, dtRegexHours, false)
	if err != nil {
		return time.Time{}, ErrParseDatetime
	}

	minutes, err := extractDurationInfo(matches, dtRegexMinutes, false)
	if err != nil {
		return time.Time{}, ErrParseDatetime
	}

	seconds, err := extractDurationInfo(matches, dtRegexSeconds, false)
	if err != nil {
		return time.Time{}, ErrParseDatetime
	}

	nanos, err := extractDurationInfo(matches, dtRegexFractal, true)
	if err != nil {
		return time.Time{}, ErrParseDatetime
	}

	offsetHours, err := extractDurationInfo(matches, dtRegexOffsetHours, false)
	if err != nil {
		return time.Time{}, ErrParseDatetime
	}

	offsetMinutes, err := extractDurationInfo(matches, dtRegexOffsetMinutes, false)
	if err != nil {
		return time.Time{}, ErrParseDatetime
	}

	// If the zone sign is '-', then we need to multiply the offset by -1
	var offsetSign int
	switch matches[dtRegexOffsetSign] {
	case "-":
		offsetSign = -1
	default:
		offsetSign = 1
	}

	// Get zone offset in seconds.
	offset := (offsetHours * 3600 + offsetMinutes * 60) * offsetSign

	// Return a new time.Time with the given values and UTC encoding.
	return time.Date(
		year,
		time.Month(month),
		day,
		hours,
		minutes,
		seconds,
		nanos,
		time.FixedZone("", offset),
	), nil
}

// Converts DICOM TM (time) value to time.Time as UTC.
func TMToTime(tmString string) (time.Time, error) {
	matches := tmRegex.FindStringSubmatch(tmString)
	// If no full match is found, return an error
	if !hasMatches(matches, tmString) {
		return time.Time{}, ErrParseTime
	}

	hours, err := extractDurationInfo(matches, tmRegexHours, false)
	if err != nil {
		return time.Time{}, ErrParseTime
	}

	minutes, err := extractDurationInfo(matches, tmRegexMinutes, false)
	if err != nil {
		return time.Time{}, ErrParseTime
	}

	seconds, err := extractDurationInfo(matches, tmRegexSeconds, false)
	if err != nil {
		return time.Time{}, ErrParseTime
	}

	nanos, err := extractDurationInfo(matches, tmRegexFractal, true)
	if err != nil {
		return time.Time{}, ErrParseTime
	}

	// Return a new time.Time with the given values and UTC encoding.
	return time.Date(
		0, 0, 0, hours, minutes, seconds, nanos, time.UTC,
	), nil
}
