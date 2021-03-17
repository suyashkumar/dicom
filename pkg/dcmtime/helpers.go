package dcmtime

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// hasPrecision returns whether `check` is included in `limit`.
//
// Example: to test whether seconds should be included, you would:
// hasPrecision(PrecisionSeconds, [caller-passed-limit])
func hasPrecision(check PrecisionLevel, precision PrecisionLevel) bool {
	return check >= precision
}

// hasPrecisionOmits is the underlying call made on [type].HasPrecision() call.
//
// check is the value passed in by the caller to check.
//
// valuePrecision is the precision of the value we are checking about.
//
// omits are a set of Precision levels the value cannot have. For instance. Date can
// have a precision of PrecisionYear, but not PrecisionSeconds
func hasPrecisionOmits(check PrecisionLevel, valuePrecision PrecisionLevel, omits precisionRange) bool {
	if check < valuePrecision {
		return false
	}

	// If this value falls within the omit range, it is false.
	if omits.Contains(check) {
		return false
	}

	return true
}

// truncateMilliseconds truncate nanosecond time.Time value to arbitrary precision.
func truncateMilliseconds(nanoSeconds int, precision PrecisionLevel) (millis string) {
	milliseconds := nanoSeconds / 1000
	millis = fmt.Sprintf("%06d", milliseconds)
	millis = millis[:6-(PrecisionFull+precision)]

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
		info.FractalPrecision = PrecisionFull + PrecisionLevel(missingPlaces-3)
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

// precisionRange defines the (inclusive) minimum and maximum precision for omits.
type precisionRange struct {
	// Min is the the minimum precision in this range (inclusive).
	Min PrecisionLevel
	// Max is the maximum precision in this range (inclusive).
	Max PrecisionLevel
}

// Contains returns true if a value falls within the given range (inclusive).
func (pRange precisionRange) Contains(val PrecisionLevel) bool {
	// Because the precision iota is reversed to make PrecisionFull the default, we
	// need to invert the normal range logic here.
	if val > pRange.Min {
		return false
	}

	if val < pRange.Max {
		return false
	}

	// If this value falls within the omit range, it is false.
	return true
}

// combineDAAndTM is the backing function for both DA.Combine and TM.Combine
//
// If no location is given time.FixedZone("", 0) will be used, and NoOffset will be
// set to 'true'.
func combineDateAndTime(da Date, tm Time, location *time.Location) (Datetime, error) {
	// User-created values might use PrecisionDay in place of PrecisionFull for
	// full-precision dates.
	//
	// We need a full precision date, because Datetime values can only elide fom the
	// end. Since combining a Date and Time value implies the Time has at least an
	// Hour value, no Date values can be missing.
	if !da.HasPrecision(PrecisionFull) && !da.HasPrecision(PrecisionDay) {
		return Datetime{}, fmt.Errorf(
			"DA value must have full precision, got '%v'", da.Precision,
		)
	}

	noOffset := false
	if location == nil {
		location = time.FixedZone("", 0)
		noOffset = true
	}

	return Datetime{
		Time: time.Date(
			// Date values.
			da.Time.Year(),
			da.Time.Month(),
			da.Time.Day(),
			// Time values.
			tm.Time.Hour(),
			tm.Time.Minute(),
			tm.Time.Second(),
			tm.Time.Nanosecond(),
			location,
		),
		// We'll inherit our Precision from the Time value.
		Precision: tm.Precision,
		NoOffset:  noOffset,
	}, nil
}
