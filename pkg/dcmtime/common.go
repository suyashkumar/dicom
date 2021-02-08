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
	// The parsed int value
	Value int
	// Whether it was present in the string
	Present bool
	// Only used when parsing Fractal seconds: the level of precision used.
	FractalPrecision PrecisionLevel
}

// extractDurationInfo extracts a piece of DA, TM, or DT info from a parsed regex,
// handling all validation checks, emtpy values, etc.
func extractDurationInfo(subMatches []string, index int, fractal bool) (durationInfo, error) {
	info := durationInfo{}

	// If the submatch array does not contain the group index we are looking for, then
	// we need to return a parsing error.
	if len(subMatches) <= index {
		return info, errors.New("not enough sub-matches")
	}

	// Get the value of our capture group.
	valueStr := subMatches[index]

	// If there was no match for the specific subgroup, this value is 0, and we leave
	// info.Present as false.
	if valueStr != "" {
		// Otherwise set the info.Present to true
		info.Present = true
	}

	// If this is a fractal seconds value, we need to pad it out to nanoseconds for go.
	if fractal {
		// The fractal value can be any length, with truncation implying a loss of
		// precision, we need to add trailing zeros to the hundred-millionth place to
		// get our nano-seconds.
		missingPlaces := 9 - len(valueStr)
		valueStr = valueStr + strings.Repeat("0", missingPlaces)
		info.FractalPrecision = PrecisionFull - PrecisionLevel(missingPlaces-3)
	}

	// If our info is present, parse the value into an int.
	if info.Present {
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
func updatePrecision(
	current PrecisionLevel,
	info durationInfo,
	infoLevel PrecisionLevel,
	levelIsFull bool,
) PrecisionLevel {
	// If the info was not present, return our current level.
	if !info.Present {
		return current
	}

	if levelIsFull {
		return PrecisionFull
	}
	return infoLevel
}

// validateRegexpMatchResult validates a regexp.Match result from a match on the
// rawValue string.
func validateRegexpMatchResult(matchResult []string, rawValue string) bool {
	// There must be at least one full match. If the result is empty, then there was no
	// match, and we can return false immediately.
	if len(matchResult) == 0 {
		return false
	}

	// If the full match is not the entire rawValue string, then it is not a valid
	// value.
	//
	// For example, when parsing a TM value, "123004.4567SomeText" would
	// return a valid result with the first entry being the full match of "123004.4567",
	// even though the input is obviously an illegal value.
	//
	// We need to validate that the first entry (full expression match) matchResult the
	// ENTIRE rawValue string. Otherwise we are dealing with a non-valid value that
	// HAPPENS to contain a valid value inside of it.
	if matchResult[0] != rawValue {
		return false
	}

	// Otherwise it's good.
	return true
}

// writeTimezoneString writes the string representation of time to builder.
//
// if useSeps is true, hours and minutes will be separated by a ":". This is used for
// the human-readable string.
func writeTimezoneString(builder *strings.Builder, time time.Time, useSeps bool) {
	layout := "-0700"
	if useSeps {
		layout = "-07:00"
	}
	tzString := time.Format(layout)
	builder.WriteString(tzString)
}
