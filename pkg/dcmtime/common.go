package dcmtime

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Returns whether `check` is included in `limit`.
//
// Example: to test whether seconds should be included, you would:
// isIncluded(PrecisionSeconds, [caller-passed-limit])
func isIncluded(check PrecisionLevel, precision PrecisionLevel) bool {
	return check <= precision
}

// Truncate nanosecond time.Time value to arbitrary precision.
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
		info.FractalPrecision = PrecisionFull - PrecisionLevel(missingPlaces-3)
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
			return PrecisionFull
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
