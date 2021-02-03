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
func truncateMilliseconds(
	nanoSeconds int, precision PrecisionLevel,
) (millisecondsStr string) {
	milliseconds := nanoSeconds / 1000
	millisecondsStr = fmt.Sprintf("%06d", milliseconds)
	millisecondsStr = millisecondsStr[:6-(PrecisionFull-precision)]

	return millisecondsStr
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

// extractDurationInfo extracts a piece of DA, TM, or DT info from a parsed regex.
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
		info.FractalPrecision = PrecisionFull - PrecisionLevel(missingPlaces-3)
	}

	info.Value, err = strconv.Atoi(valueStr)
	if err != nil {
		return info, fmt.Errorf("error parsing int: %w", err)
	}

	return info, nil
}

// updatePrecision takes in our current precision and a piece of parsed info, then
// updates the precision based on whether the info was present.
//
// levelIsFull should be set to true if the level we are checking is "Full" for the
// type of value we are parsing.
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

// hasMatches checks if a regex result has matches.
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

// writeTimezoneString writes the string representation of time to builder.
//
// if useSeps is true, hours and minutes will be separated by a ":",
func writeTimezoneString(builder *strings.Builder, time time.Time, useSeps bool) {
	// Get the seconds offset from the zone.
	_, offset := time.Zone()

	// Deduce the offset sign, then invert the offset to positive if it is negative.
	offsetSign := "+"
	if offset < 0 {
		offsetSign = "-"
		offset *= -1
	}

	// Divide seconds by 60 to get minutes
	offsetMinutes := offset / 60
	// Hours equal minutes divided by 60
	offsetHours := offsetMinutes / 60
	// Set minutes to remainder of minutes divided by 60
	offsetMinutes %= 60

	// Add the offset sign and hours.
	builder.WriteString(fmt.Sprintf("%v%02d", offsetSign, offsetHours))

	// If we are using seps, add a colon.
	if useSeps {
		builder.WriteRune(':')
	}

	// Add the Minutes.
	builder.WriteString(fmt.Sprintf("%02d", offsetMinutes))
}
