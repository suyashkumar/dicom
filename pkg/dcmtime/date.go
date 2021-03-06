package dcmtime

import (
	"fmt"
	"strings"
	"time"
)

// Date holds data for a parsed DICOM date (DA value).
type Date struct {
	// Time is the underlying time.Time value.
	Time time.Time

	// Precision with which the raw DA value was stored. For instance, a Date value
	// with a precision of PrecisionYear ONLY stored the year.
	Precision PrecisionLevel

	// IsNEMA will be true if this is a legacy NEMA-300 formatted value (1999.01.03).
	IsNEMA bool
}

// DCM converts time.Time value to dicom DA string. Values are truncated to the
// Date.Precision value.
//
// NOTE: Time zones are ignored in this operation, as DA does not support encoding them.
// Make sure values are converted to UTC before passing if that is the desired output.
func (da Date) DCM() string {
	year, month, day := da.Time.Date()
	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("%04d", year))
	if !isIncluded(PrecisionMonth, da.Precision) {
		return builder.String()
	}

	// If this is a NEMA-formatted time we need to separate the values with a ".".
	if da.IsNEMA {
		builder.WriteString(".")
	}

	builder.WriteString(fmt.Sprintf("%02d", month))
	if !isIncluded(PrecisionDay, da.Precision) {
		return builder.String()
	}

	if da.IsNEMA {
		builder.WriteString(".")
	}

	builder.WriteString(fmt.Sprintf("%02d", day))
	return builder.String()
}

// String implements fmt.Stringer.
func (da Date) String() string {
	builder := strings.Builder{}
	_, _ = builder.WriteString(fmt.Sprintf("%04d", da.Time.Year()))
	if !isIncluded(PrecisionMonth, da.Precision) {
		return builder.String()
	}

	_, _ = builder.WriteString(fmt.Sprintf("-%02d", da.Time.Month()))
	if !isIncluded(PrecisionDay, da.Precision) {
		return builder.String()
	}

	_, _ = builder.WriteString(fmt.Sprintf("-%02d", da.Time.Day()))
	return builder.String()
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
	year,
	month,
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
	precisionOut = updatePrecision(year, precisionIn, PrecisionYear, false)

	month, err = extractDurationInfo(matches, regexGroups.Month, false)
	if err != nil {
		return year, month, day, precisionOut, err
	}
	// If there is no month, we need to set it to 1 instead of zero, otherwise the year
	// will roll back 1.
	if !month.PresentInSource {
		month.Value = 1
	}
	precisionOut = updatePrecision(month, precisionOut, PrecisionMonth, false)

	day, err = extractDurationInfo(matches, regexGroups.Day, false)
	if err != nil {
		return year, month, day, precisionOut, err
	}
	// If there is no day, we need to set it to 1 instead of zero, otherwise the month
	// will roll back 1, ie December -> November.
	if !day.PresentInSource {
		day.Value = 1
	}
	precisionOut = updatePrecision(day, precisionOut, PrecisionDay, isDA)

	return year, month, day, precisionOut, err
}

// ParseDate converts DICOM DA (date) value to time.Time as UTC with PrecisionLevel.
func ParseDate(daString string) (Date, error) {
	// isNema will store whether this value is a NEMA-300 formatted value.
	isNEMA := false

	// Try to match against the normal DA value regex
	matches := daRegex.FindStringSubmatch(daString)
	// If that resulted in no matches and the caller wants to allow for NEMA-300 legacy
	// parsing, try that next.
	if len(matches) == 0 {
		matches = daRegexNEMA.FindStringSubmatch(daString)
		isNEMA = true
	}

	// If we did not find any matches, this is not a valid DA value, exit with an error.
	if len(matches) == 0 {
		return Date{}, ErrParseDA
	}

	year, month, day, precision, err := extractDate(matches, PrecisionFull, true)
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
		IsNEMA:    isNEMA,
		Precision: precision,
	}, nil
}
