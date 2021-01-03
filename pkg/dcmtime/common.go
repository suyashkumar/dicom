package dcmtime

import (
	"fmt"
	"time"
)

// Returns whether `check` is included in `limit`.
//
// Example: to test whether seconds should be included, you would:
//	isIncluded(Precision.Seconds, [caller-passed-limit])
func isIncluded(check PrecisionLevel, precision PrecisionLevel) bool {
	return check <= precision
}

// Truncate nanosecond time.Time value to arbitrary precision.
func truncateMilliseconds(
	nanoSeconds int, precision PrecisionLevel,
) (millisecondsStr string) {
	milliseconds := nanoSeconds / 1000
	millisecondsStr = fmt.Sprintf("%06d", milliseconds)
	millisecondsStr = millisecondsStr[:6-(Precision.Full-precision)]

	return millisecondsStr
}

// Const for time.Time timezone. We don't want a timezone because we don't REALLY know
// that a TM or DA value is UTC, so we don't want UTC to be made explicit.
var zeroTimezone = time.FixedZone("", 0)
