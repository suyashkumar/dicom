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

// Converts DICOM DA (date) value to time.Time as UTC.
//
// If allowNEMA is true, the old NEMA 300 specification of YYYY.MM.DD will also be
// checked if parsing as YYYYMMDD fails.
func ParseDA(daString string, allowNEMA bool) (DA, error) {
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
		return DA{}, ErrParseDA
	}

	var precision PrecisionLevel

	// The year has to be present, if it isn't, it will not pass the regex.
	year, err := extractDurationInfo(matches, daRegexYear, false)
	if err != nil {
		return DA{}, ErrParseDA
	}
	precision = updatePrecision(precision, year, Precision.Year, false)

	month, err := extractDurationInfo(matches, daRegexMonth, false)
	if err != nil {
		return DA{}, ErrParseDA
	}
	precision = updatePrecision(precision, month, Precision.Month, false)

	day, err := extractDurationInfo(matches, daRegexDay, false)
	if err != nil {
		return DA{}, ErrParseDA
	}
	precision = updatePrecision(precision, day, Precision.Day, true)

	// Return a new time.Time with the given values and UTC encoding.
	parsed := time.Date(
		year.Value,
		time.Month(month.Value),
		day.Value,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	return DA{
		Time:      parsed,
		Precision: precision,
	}, nil
}

// Converts DICOM DT (datetime) value to time.Time as UTC.
func ParseDT(dtString string) (DT, error) {
	matches := dtRegex.FindStringSubmatch(dtString)
	if !hasMatches(matches, dtString) {
		return DT{}, ErrParseDT
	}

	precision := Precision.Full

	year, err := extractDurationInfo(matches, dtRegexYear, false)
	if err != nil {
		return DT{}, ErrParseDT
	}
	precision = updatePrecision(precision, year, Precision.Year, false)

	month, err := extractDurationInfo(matches, dtRegexMonth, false)
	if err != nil {
		return DT{}, ErrParseDT
	}
	precision = updatePrecision(precision, month, Precision.Month, false)

	day, err := extractDurationInfo(matches, dtRegexDay, false)
	if err != nil {
		return DT{}, ErrParseDT
	}
	precision = updatePrecision(precision, day, Precision.Day, false)

	hours, err := extractDurationInfo(matches, dtRegexHours, false)
	if err != nil {
		return DT{}, ErrParseDT
	}
	precision = updatePrecision(precision, hours, Precision.Hours, false)

	minutes, err := extractDurationInfo(matches, dtRegexMinutes, false)
	if err != nil {
		return DT{}, ErrParseDT
	}
	precision = updatePrecision(precision, minutes, Precision.Minutes, false)

	seconds, err := extractDurationInfo(matches, dtRegexSeconds, false)
	if err != nil {
		return DT{}, ErrParseDT
	}
	precision = updatePrecision(precision, seconds, Precision.Seconds, false)

	nanos, err := extractDurationInfo(matches, dtRegexFractal, true)
	if err != nil {
		return DT{}, ErrParseDT
	} else if nanos.Present {
		precision = nanos.FractalPrecision
	}

	var offsetSpecified bool

	offsetHours, err := extractDurationInfo(matches, dtRegexOffsetHours, false)
	if err != nil {
		return DT{}, ErrParseDT
	}
	// If hours are not present, there is either no offset or the regex will fail,
	// so we only need to check this here.
	if offsetHours.Present {
		offsetSpecified = true
	}

	offsetMinutes, err := extractDurationInfo(matches, dtRegexOffsetMinutes, false)
	if err != nil {
		return DT{}, ErrParseDT
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

	return DT{
		Time:         parsed,
		Precision:    precision,
		IgnoreOffset: offsetSpecified,
	}, nil
}

// Converts DICOM TM (time) value to time.Time as UTC.
func ParseTM(tmString string) (TM, error) {
	matches := tmRegex.FindStringSubmatch(tmString)
	// If no full match is found, return an error
	if !hasMatches(matches, tmString) {
		return TM{}, ErrParseTM
	}

	var precision PrecisionLevel

	// Hours has to be present, if it is not, the regex will not pass.
	hours, err := extractDurationInfo(matches, tmRegexHours, false)
	if err != nil {
		return TM{}, ErrParseTM
	}
	precision = updatePrecision(precision, hours, Precision.Hours, false)

	minutes, err := extractDurationInfo(matches, tmRegexMinutes, false)
	if err != nil {
		return TM{}, ErrParseTM
	}
	precision = updatePrecision(precision, minutes, Precision.Minutes, false)

	seconds, err := extractDurationInfo(matches, tmRegexSeconds, false)
	if err != nil {
		return TM{}, ErrParseTM
	}
	precision = updatePrecision(precision, seconds, Precision.Seconds, false)

	nanos, err := extractDurationInfo(matches, tmRegexFractal, true)
	if err != nil {
		return TM{}, ErrParseTM
	} else if nanos.Present {
		precision = nanos.FractalPrecision
	}

	// Return a new time.Time with the given values and UTC encoding.
	parsed := time.Date(
		0,
		0,
		0,
		hours.Value,
		minutes.Value,
		seconds.Value,
		nanos.Value,
		time.UTC,
	)

	return TM{
		Time:      parsed,
		Precision: precision,
	}, nil
}
