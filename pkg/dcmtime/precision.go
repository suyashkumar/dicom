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
	case Precision.Full:
		return "FULL"
	case Precision.Year:
		return "YEAR"
	case Precision.Month:
		return "MONTH"
	case Precision.Day:
		return "DAY"
	case Precision.Hours:
		return "HOURS"
	case Precision.Minutes:
		return "MINUTES"
	case Precision.Seconds:
		return "SECONDS"
	case Precision.MS1:
		return "MS1"
	case Precision.MS2:
		return "MS2"
	case Precision.MS3:
		return "MS3"
	case Precision.MS4:
		return "MS4"
	case Precision.MS5:
		return "MS5"
	default:
		return "!{PRECISION}"
	}
}

// Precision is a namespace for PrecisionLevel values.
var Precision = struct {
	Year,
	Month,
	Day,
	Hours,
	Minutes,
	Seconds,
	MS1,
	MS2,
	MS3,
	MS4,
	MS5,
	Full PrecisionLevel
}{
	Year:    0,
	Month:   1,
	Day:     2,
	Hours:   3,
	Minutes: 4,
	Seconds: 5,
	MS1:     6,
	MS2:     7,
	MS3:     8,
	MS4:     9,
	MS5:     10,
	Full:    11,
}
