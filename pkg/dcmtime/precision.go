package dcmtime

// PrecisionLevel tells encoding methods how many segments of the value to render
// to the DICOM format string, as the spec allows eliding segments to communicate the
// precision.
//
// For instance: when rendering a DA (date) value, using a truncation of Precision.Year
// would render the entire date as 'YYYY'.
//
// Using Precision.Month would render 'YYYYMM'.
//
// Using Precision.Day or Precision.Full would render a full 'YYYYMMDD'.
type PrecisionLevel int

// String returns the name of the PrecisionLevel for debugging.
func (level PrecisionLevel) String() string {
	switch level {
	case PrecisionFull:
		return "FULL"
	case PrecisionYear:
		return "YEAR"
	case PrecisionMonth:
		return "MONTH"
	case PrecisionDay:
		return "DAY"
	case PrecisionHours:
		return "HOURS"
	case PrecisionMinutes:
		return "MINUTES"
	case PrecisionSeconds:
		return "SECONDS"
	case PrecisionMS1:
		return "MS1"
	case PrecisionMS2:
		return "MS2"
	case PrecisionMS3:
		return "MS3"
	case PrecisionMS4:
		return "MS4"
	case PrecisionMS5:
		return "MS5"
	default:
		return "!{PRECISION}"
	}
}

const (
	// PrecisionYear indicated that a given dcm time value is only precise to the year.
	PrecisionYear PrecisionLevel = iota
	// PrecisionMonth indicated that a given dcm time value is only precise to the
	// month.
	PrecisionMonth
	// PrecisionDay indicated that a given dcm time value is only precise to the day.
	PrecisionDay
	// PrecisionHours indicated that a given dcm time value is only precise to the hour.
	PrecisionHours
	// PrecisionMinutes indicated that a given dcm time value is only precise to the
	// minute.
	PrecisionMinutes
	// PrecisionSeconds indicated that a given dcm time value is only precise to the
	// second.
	PrecisionSeconds
	// PrecisionMS1 indicated that a given dcm time value is only precise to 1
	// millisecond place (1/10 of a second).
	PrecisionMS1
	// PrecisionMS2 indicated that a given dcm time value is only precise to 2
	// millisecond place (1/100 of a second).
	PrecisionMS2
	// PrecisionMS3 indicated that a given dcm time value is only precise to 3
	// millisecond place (1/1000 of a second).
	PrecisionMS3
	// PrecisionMS4 indicated that a given dcm time value is only precise to 4
	// millisecond place (1/10000 of a second).
	PrecisionMS4
	// PrecisionMS5 indicated that a given dcm time value is only precise to 4
	// millisecond place (1/100000 of a second).
	PrecisionMS5
	// PrecisionFull indicates that a given dcm time value is precise to the full extent
	// it is able to be.
	PrecisionFull
)
