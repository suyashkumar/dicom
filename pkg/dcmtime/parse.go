package dcmtime

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Parsed piece of DA, TM, or DT info.
type durationInfo struct {
	// The parsed int value
	Value int
	// Whether it was present in the string
	Present bool
	// Only used when parsing Fractal seconds: the level of precision used.
	FractalPrecision PrecisionLevel
}

func extractDurationInfo(subMatches []string, index int, fractal bool) (
	info durationInfo, err error,
) {
	if len(subMatches) <= index {
		return info, errors.New("not enough sub-matches")
	}

	valueStr := subMatches[index]
	// If there was no match for the specific subgroup, this value is 0.
	if valueStr == "" {
		valueStr = "0"
	} else {
		info.Present = true
	}

	// If this is a fractal seconds value, we need to pad it out to nanoseconds for go.
	if fractal {
		// The fractal value can be any length, with truncation implying a loss of
		// precision, we need to add trailing zeros to the hundred-millionth place to
		// get our nano-seconds.
		missingPlaces := 9 - len(valueStr)
		valueStr = valueStr + strings.Repeat("0", missingPlaces)
		info.FractalPrecision = Precision.Full - PrecisionLevel(missingPlaces-3)
	}

	info.Value, err = strconv.Atoi(valueStr)
	if err != nil {
		return info, fmt.Errorf("error parsing int: %w", err)
	}

	return info, nil
}

func updatePrecision(
	current PrecisionLevel,
	info durationInfo,
	infoLevel PrecisionLevel,
	levelIsFull bool,
) PrecisionLevel {
	if info.Present {
		if levelIsFull {
			return Precision.Full
		}
		return infoLevel
	}

	return current
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

// Holds group indexes for a given source types regex.
type dateRegexGroups struct {
	Year  int
	Month int
	Day   int
}

// Cached info on DA date regex groups.
var dateRegexGroupsDA = dateRegexGroups{
	Year:  daRegexGroupYear,
	Month: daRegexGroupMonth,
	Day:   daRegexGroupDay,
}

// Cached info on DT date regex groups.
var dateRegexGroupsDT = dateRegexGroups{
	Year:  dtRegexGroupYear,
	Month: dtRegexGroupMonth,
	Day:   dtRegexGroupDay,
}

// extracts the date info from a DA or DT regex match.
func extractDate(
	matches []string,
	precisionIn PrecisionLevel,
	isDA bool,
) (
	year durationInfo,
	month durationInfo,
	day durationInfo,
	precisionOut PrecisionLevel,
	err error,
) {
	regexGroups := dateRegexGroupsDA
	if !isDA {
		regexGroups = dateRegexGroupsDT
	}

	// The year MUST be present in all DA and DT values or the regex will not pass.
	year, err = extractDurationInfo(matches, regexGroups.Year, false)
	if err != nil {
		return year, month, day, precisionOut, err
	}
	precisionOut = updatePrecision(precisionIn, year, Precision.Year, false)

	month, err = extractDurationInfo(matches, regexGroups.Month, false)
	if err != nil {
		return year, month, day, precisionOut, err
	}
	// If there is no month, we need to set it to 1 instead of zero, otherwise the year
	// will roll back 1.
	if !month.Present {
		month.Value = 1
	}
	precisionOut = updatePrecision(precisionOut, month, Precision.Month, false)

	day, err = extractDurationInfo(matches, regexGroups.Day, false)
	if err != nil {
		return year, month, day, precisionOut, err
	}
	// If there is no day, we need to set it to 1 instead of zero, otherwise the month
	// will roll back 1, ie December -> November.
	if !day.Present {
		day.Value = 1
	}
	precisionOut = updatePrecision(precisionOut, day, Precision.Day, isDA)

	return year, month, day, precisionOut, err
}

// Converts DICOM DA (date) value to time.Time as UTC.
//
// If allowNEMA is true, the old NEMA 300 specification of YYYY.MM.DD will also be
// checked if parsing as YYYYMMDD fails.
func ParseDA(daString string, allowNEMA bool) (Date, error) {
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
		return Date{}, ErrParseDA
	}

	year, month, day, precision, err := extractDate(matches, Precision.Full, true)
	if err != nil {
		return Date{}, err
	}

	// Return a new time.Time with the given values and UTC encoding.
	parsed := time.Date(
		year.Value,
		time.Month(month.Value),
		day.Value,
		0,
		0,
		0,
		0,
		zeroTimezone,
	)

	return Date{
		Time:      parsed,
		Precision: precision,
	}, nil
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
	precisionOut = updatePrecision(precisionIn, hours, Precision.Hours, false)

	minutes, err = extractDurationInfo(matches, groupIndexes.Minutes, false)
	if err != nil {
		return hours, minutes, seconds, nanos, precisionOut, err
	}
	precisionOut = updatePrecision(precisionOut, minutes, Precision.Minutes, false)

	seconds, err = extractDurationInfo(matches, groupIndexes.Seconds, false)
	if err != nil {
		return hours, minutes, seconds, nanos, precisionOut, err
	}
	precisionOut = updatePrecision(precisionOut, seconds, Precision.Seconds, false)

	nanos, err = extractDurationInfo(matches, groupIndexes.Fractal, true)
	if err != nil {
		return hours, minutes, seconds, nanos, precisionOut, err
	} else if nanos.Present {
		precisionOut = nanos.FractalPrecision
	}

	return hours, minutes, seconds, nanos, precisionOut, err
}

// Converts DICOM TM (time) value to time.Time as UTC.
func ParseTM(tmString string) (Time, error) {
	matches := tmRegex.FindStringSubmatch(tmString)
	// If no full match is found, return an error
	if !hasMatches(matches, tmString) {
		return Time{}, ErrParseTM
	}

	hours, minutes, seconds, nanos, precision, err := extractTime(
		matches, Precision.Full, true,
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

// Converts DICOM DT (datetime) value to time.Time as UTC.
func ParseDT(dtString string) (Datetime, error) {
	matches := dtRegex.FindStringSubmatch(dtString)
	if !hasMatches(matches, dtString) {
		return Datetime{}, ErrParseDT
	}

	year, month, day, precision, err := extractDate(matches, Precision.Full, false)
	if err != nil {
		return Datetime{}, err
	}

	hours, minutes, seconds, nanos, precision, err := extractTime(
		matches, precision, false,
	)
	if err != nil {
		return Datetime{}, err
	}

	var hasOffet bool

	offsetHours, err := extractDurationInfo(matches, dtRegexGroupOffsetHours, false)
	if err != nil {
		return Datetime{}, ErrParseDT
	}
	// If hours are not present, there is either no offset or the regex will fail,
	// so we only need to check this here.
	if offsetHours.Present {
		hasOffet = true
	}

	offsetMinutes, err := extractDurationInfo(matches, dtRegexGroupOffsetMinutes, false)
	if err != nil {
		return Datetime{}, ErrParseDT
	}

	// If the zone sign is '-', then we need to multiply the offset by -1
	var offsetSign int
	switch matches[dtRegexGroupOffsetSign] {
	case "-":
		offsetSign = -1
	default:
		offsetSign = 1
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
		HasOffset: hasOffet,
	}, nil
}
