package dcmtime

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// isIncluded returns whether `check` is included in `limit`.
//
// Example: to test whether seconds should be included, you would:
// isIncluded(PrecisionSeconds, [caller-passed-limit])
func isIncluded(check PrecisionLevel, precision PrecisionLevel) bool {
	return check <= precision
}

// truncateMilliseconds truncate nanosecond time.Time value to arbitrary precision.
func truncateMilliseconds(nanoSeconds int, precision PrecisionLevel) (millis string) {
	milliseconds := nanoSeconds / 1000
	millis = fmt.Sprintf("%06d", milliseconds)
	millis = millis[:6-(PrecisionFull-precision)]

	return millis
}

// Const for time.Time timezone. We don't want a timezone because we don't REALLY know
// that a TM or DA value is UTC, so we don't want UTC to be made explicit.
var zeroTimezone = time.FixedZone("", 0)

// durationInfo holds parsed piece of DA, TM, or DT info.
type durationInfo struct {
	// Value is the parsed int value.
	Value int
	// PresentInSource is whether it was present in the string.
	PresentInSource bool
	// FractalPrecision is only used when parsing Fractal seconds: the level of
	// precision used.
	FractalPrecision PrecisionLevel
}

// extractDurationInfo extracts a piece of DA, TM, or DT info from a parsed regex,
// handling all validation checks, emtpy values, etc.
func extractDurationInfo(subMatches []string, index int, isFractal bool) (durationInfo, error) {
	info := durationInfo{}

	// If the submatch array does not contain the group index we are looking for, then
	// we need to return a parsing error.
	if len(subMatches) <= index {
		return info, errors.New("not enough sub-matches")
	}

	// Get the value of our capture group.
	valueStr := subMatches[index]

	// If there was no match for the specific subgroup, this value is 0, and we leave
	// info.PresentInSource as false.
	if valueStr != "" {
		// Otherwise set the info.PresentInSource to true
		info.PresentInSource = true
	}

	// If this is a isFractal seconds value, we need to pad it out to nanoseconds for
	// go.
	if info.PresentInSource && isFractal {
		// The isFractal value can be any length, with truncation implying a loss of
		// precision, we need to add trailing zeros to the hundred-millionth place to
		// get our nano-seconds.
		missingPlaces := 9 - len(valueStr)
		valueStr = valueStr + strings.Repeat("0", missingPlaces)
		info.FractalPrecision = PrecisionFull - PrecisionLevel(missingPlaces-3)
	}

	// If our info is present, parse the value into an int.
	if info.PresentInSource {
		var err error
		info.Value, err = strconv.Atoi(valueStr)
		if err != nil {
			return info, fmt.Errorf("error parsing int: %w", err)
		}
	}

	return info, nil
}

// updatePrecision takes in our current precision and a piece of parsed info, then
// updates the precision based on whether the info was present.
//
// levelIsFull should be set to true if the level we are checking is "Full" for the
// type of value we are parsing. For instance, PrecisionHours would be the full
// precision of a TM value, and we'll use PrecisionFull instead.
func updatePrecision(info durationInfo, current, infoLevel PrecisionLevel, levelIsFull bool) PrecisionLevel {
	// If the info was not present, return our current level.
	if !info.PresentInSource {
		return current
	}

	if levelIsFull {
		return PrecisionFull
	}
	return infoLevel
}
