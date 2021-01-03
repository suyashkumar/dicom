package dcmtime

import (
	"fmt"
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
