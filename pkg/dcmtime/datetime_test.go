package dcmtime_test

import (
	"errors"
	"github.com/suyashkumar/dicom/pkg/dcmtime"
	"testing"
	"time"
)

// dtPrecisionOmits describes the precision range not valid for datetime. Since datetime
// has no omits, the range is outside our enumerated values.
var dtPrecisionOmits = precisionRange{
	Min: dcmtime.PrecisionFull + 1,
	Max: dcmtime.PrecisionFull + 1,
}

func TestParseDatetime(t *testing.T) {
	testCases := []struct {
		// Name is the name of the sub-test
		Name string
		// DTValue is the RAW DICOM DT value we are going to parse.
		DTValue string
		// ExpectedTime is the time.Time value we expect to be returned by the parse.
		ExpectedTime time.Time
		// ExpectedPrecision is the PrecisionLevel we expect to be returned by the
		// parse.
		ExpectedPrecision dcmtime.PrecisionLevel
		// ExpectedString is the expected result from the value's String() method.
		ExpectedString string
		// ExpectedNoOffset is the ExpectedNoOffset value we expect to get from the parse
		ExpectedNoOffset bool
		// HasMonth is whether the parsed value's Month() method should return ok=true
		HasMonth bool
		// HasDay is whether the parsed value's Day() method should return ok=true
		HasDay bool
		// HasHour is whether the parsed value's Hour() method should return ok=true.
		HasHour bool
		// HasMinute is whether the parsed value's Minute() method should return ok=true.
		HasMinute bool
		// HasSecond is whether the parsed value's Second() method should return ok=true.
		HasSecond bool
		// HasNanosecond is whether the parsed value's HasNanosecond() method should
		// return ok=true.
		HasNanosecond bool
		// HasPrecisionRange is the range of Precision Values we expect the
		// HasPrecision() method to return true for.
		HasPrecisionRange precisionRange
	}{
		{
			Name:              "PrecisionFull-PositiveOffset",
			DTValue:           "10100203040506.456789+0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionFull,
			ExpectedString:    "1010-02-03 04:05:06.456789 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionFull,
			},
		},
		{
			Name:              "PrecisionMS5-PositiveOffset",
			DTValue:           "10100203040506.45678+0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456780000, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionMS5,
			ExpectedString:    "1010-02-03 04:05:06.45678 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS5,
			},
		},
		{
			Name:              "PrecisionMS4-PositiveOffset",
			DTValue:           "10100203040506.4567+0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456700000, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionMS4,
			ExpectedString:    "1010-02-03 04:05:06.4567 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS4,
			},
		},
		{
			Name:              "PrecisionMS3-PositiveOffset",
			DTValue:           "10100203040506.456+0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456000000, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionMS3,
			ExpectedString:    "1010-02-03 04:05:06.456 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS3,
			},
		},
		{
			Name:              "PrecisionMS2-PositiveOffset",
			DTValue:           "10100203040506.45+0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 450000000, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionMS2,
			ExpectedString:    "1010-02-03 04:05:06.45 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS2,
			},
		},
		{
			Name:              "PrecisionMS1-PositiveOffset",
			DTValue:           "10100203040506.4+0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 400000000, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionMS1,
			ExpectedString:    "1010-02-03 04:05:06.4 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS1,
			},
		},
		{
			Name:              "PrecisionSeconds-PositiveOffset",
			DTValue:           "10100203040506+0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 0, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionSeconds,
			ExpectedString:    "1010-02-03 04:05:06 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionSeconds,
			},
		},
		{
			Name:              "PrecisionMinutes-PositiveOffset",
			DTValue:           "101002030405+0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 0, 0, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionMinutes,
			ExpectedString:    "1010-02-03 04:05 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMinutes,
			},
		},
		{
			Name:              "PrecisionHours-PositiveOffset",
			DTValue:           "1010020304+0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 0, 0, 0, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionHours,
			ExpectedString:    "1010-02-03 04 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionHours,
			},
		},
		{
			Name:              "PrecisionDay-PositiveOffset",
			DTValue:           "10100203+0102",
			ExpectedTime:      time.Date(1010, 2, 3, 0, 0, 0, 0, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionDay,
			ExpectedString:    "1010-02-03 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           false,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionDay,
			},
		},
		{
			Name:              "PrecisionMonth-PositiveOffset",
			DTValue:           "101002+0102",
			ExpectedTime:      time.Date(1010, 2, 1, 0, 0, 0, 0, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionMonth,
			ExpectedString:    "1010-02 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            false,
			HasHour:           false,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMonth,
			},
		},
		{
			Name:              "PrecisionYear-PositiveOffset",
			DTValue:           "1010+0102",
			ExpectedTime:      time.Date(1010, 1, 1, 0, 0, 0, 0, time.FixedZone("", 3720)),
			ExpectedPrecision: dcmtime.PrecisionYear,
			ExpectedString:    "1010 +01:02",
			ExpectedNoOffset:  false,
			HasMonth:          false,
			HasDay:            false,
			HasHour:           false,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionYear,
			},
		},
		{
			Name:              "PrecisionFull-NegativeOffset",
			DTValue:           "10100203040506.456789-0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionFull,
			ExpectedString:    "1010-02-03 04:05:06.456789 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionFull,
			},
		},
		{
			Name:              "PrecisionMS5-NegativeOffset",
			DTValue:           "10100203040506.45678-0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456780000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionMS5,
			ExpectedString:    "1010-02-03 04:05:06.45678 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS5,
			},
		},
		{
			Name:              "PrecisionMS4-NegativeOffset",
			DTValue:           "10100203040506.4567-0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456700000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionMS4,
			ExpectedString:    "1010-02-03 04:05:06.4567 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS4,
			},
		},
		{
			Name:              "PrecisionMS3-NegativeOffset",
			DTValue:           "10100203040506.456-0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456000000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionMS3,
			ExpectedString:    "1010-02-03 04:05:06.456 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS3,
			},
		},
		{
			Name:              "PrecisionMS2-NegativeOffset",
			DTValue:           "10100203040506.45-0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 450000000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionMS2,
			ExpectedString:    "1010-02-03 04:05:06.45 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS2,
			},
		},
		{
			Name:              "PrecisionMS1-NegativeOffset",
			DTValue:           "10100203040506.4-0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 400000000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionMS1,
			ExpectedString:    "1010-02-03 04:05:06.4 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS1,
			},
		},
		{
			Name:              "PrecisionSeconds-NegativeOffset",
			DTValue:           "10100203040506-0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 000000000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionSeconds,
			ExpectedString:    "1010-02-03 04:05:06 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionSeconds,
			},
		},
		{
			Name:              "PrecisionMinutes-NegativeOffset",
			DTValue:           "101002030405-0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 0, 000000000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionMinutes,
			ExpectedString:    "1010-02-03 04:05 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMinutes,
			},
		},
		{
			Name:              "PrecisionHours-NegativeOffset",
			DTValue:           "1010020304-0102",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 0, 0, 000000000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionHours,
			ExpectedString:    "1010-02-03 04 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionHours,
			},
		},
		{
			Name:              "PrecisionDay-NegativeOffset",
			DTValue:           "10100203-0102",
			ExpectedTime:      time.Date(1010, 2, 3, 0, 0, 0, 000000000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionDay,
			ExpectedString:    "1010-02-03 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           false,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionDay,
			},
		},
		{
			Name:              "PrecisionMonth-NegativeOffset",
			DTValue:           "101002-0102",
			ExpectedTime:      time.Date(1010, 2, 1, 0, 0, 0, 000000000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionMonth,
			ExpectedString:    "1010-02 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          true,
			HasDay:            false,
			HasHour:           false,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMonth,
			},
		},
		{
			Name:              "PrecisionYear-NegativeOffset",
			DTValue:           "1010-0102",
			ExpectedTime:      time.Date(1010, 1, 1, 0, 0, 0, 000000000, time.FixedZone("", -3720)),
			ExpectedPrecision: dcmtime.PrecisionYear,
			ExpectedString:    "1010 -01:02",
			ExpectedNoOffset:  false,
			HasMonth:          false,
			HasDay:            false,
			HasHour:           false,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionYear,
			},
		},
		{
			Name:              "PrecisionFull-NoOffset",
			DTValue:           "10100203040506.456789",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionFull,
			ExpectedString:    "1010-02-03 04:05:06.456789",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionFull,
			},
		},
		{
			Name:              "PrecisionMS5-NoOffset",
			DTValue:           "10100203040506.45678",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456780000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS5,
			ExpectedString:    "1010-02-03 04:05:06.45678",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS5,
			},
		},
		{
			Name:              "PrecisionMS4-NoOffset",
			DTValue:           "10100203040506.4567",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456700000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS4,
			ExpectedString:    "1010-02-03 04:05:06.4567",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS4,
			},
		},
		{
			Name:              "PrecisionMS3-NoOffset",
			DTValue:           "10100203040506.456",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 456000000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS3,
			ExpectedString:    "1010-02-03 04:05:06.456",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS3,
			},
		},
		{
			Name:              "PrecisionMS2-NoOffset",
			DTValue:           "10100203040506.45",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 450000000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS2,
			ExpectedString:    "1010-02-03 04:05:06.45",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS2,
			},
		},
		{
			Name:              "PrecisionMS1-NoOffset",
			DTValue:           "10100203040506.4",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 400000000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS1,
			ExpectedString:    "1010-02-03 04:05:06.4",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMS1,
			},
		},
		{
			Name:              "PrecisionSeconds-NoOffset",
			DTValue:           "10100203040506",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 6, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionSeconds,
			ExpectedString:    "1010-02-03 04:05:06",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionSeconds,
			},
		},
		{
			Name:              "PrecisionMinutes-NoOffset",
			DTValue:           "101002030405",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 5, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMinutes,
			ExpectedString:    "1010-02-03 04:05",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         true,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMinutes,
			},
		},
		{
			Name:              "PrecisionHours-NoOffset",
			DTValue:           "1010020304",
			ExpectedTime:      time.Date(1010, 2, 3, 4, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionHours,
			ExpectedString:    "1010-02-03 04",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           true,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionHours,
			},
		},

		// Full value, no offset, no hours
		{
			Name:              "PrecisionDay-NoOffset",
			DTValue:           "10100203",
			ExpectedTime:      time.Date(1010, 2, 3, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionDay,
			ExpectedString:    "1010-02-03",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            true,
			HasHour:           false,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionDay,
			},
		},

		// Full value, no offset, no days
		{
			Name:              "PrecisionMonth-NoOffset",
			DTValue:           "101002",
			ExpectedTime:      time.Date(1010, 2, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMonth,
			ExpectedString:    "1010-02",
			ExpectedNoOffset:  true,
			HasMonth:          true,
			HasDay:            false,
			HasHour:           false,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMonth,
			},
		},

		// Full value, no offset, no month
		{
			Name:              "PrecisionYear-NoOffset",
			DTValue:           "1010",
			ExpectedTime:      time.Date(1010, 1, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionYear,
			ExpectedString:    "1010",
			ExpectedNoOffset:  true,
			HasMonth:          false,
			HasDay:            false,
			HasHour:           false,
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionYear,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			var parsed dcmtime.Datetime

			t.Run("ParseDatetime()", func(t *testing.T) {
				var err error

				parsed, err = dcmtime.ParseDatetime(tc.DTValue)
				if err != nil {
					t.Fatal("parse err:", err)
				}

				if !tc.ExpectedTime.Equal(parsed.Time) {
					t.Errorf(
						"Datetime.Time: expected %v, got %v",
						tc.ExpectedTime,
						parsed.Time,
					)

				}

				if parsed.Precision != tc.ExpectedPrecision {
					t.Errorf(
						"Datetime.Precision: expected %v, got %v",
						tc.ExpectedPrecision.String(),
						parsed.Precision.String(),
					)
				}

				if parsed.NoOffset != tc.ExpectedNoOffset {
					t.Errorf(
						"Datetime.NoOffset: expected %v, got %v",
						tc.ExpectedNoOffset,
						parsed.NoOffset,
					)
				}
			})

			t.Run("String()", func(t *testing.T) {
				stringVal := parsed.String()
				if stringVal != tc.ExpectedString {
					t.Fatalf(
						"got String() value '%v', expected '%v'",
						stringVal,
						tc.ExpectedString,
					)
				}
			})

			t.Run("DCM()", func(t *testing.T) {
				dcmVal := parsed.DCM()
				if dcmVal != tc.DTValue {
					t.Fatalf(
						"got DCM() value '%v', expected '%v'", dcmVal, tc.DTValue,
					)
				}
			})

			t.Run("Year()", func(t *testing.T) {
				year := parsed.Year()
				checkDateHelperOutput(t, "Year()", parsed.Time.Year(), year, true, true)
			})

			t.Run("Month()", func(t *testing.T) {
				month, ok := parsed.Month()
				checkDateHelperOutput(t, "Month()", int(parsed.Time.Month()), int(month), tc.HasMonth, ok)
			})

			t.Run("Day()", func(t *testing.T) {
				day, ok := parsed.Day()
				checkDateHelperOutput(t, "Day()", parsed.Time.Day(), day, tc.HasDay, ok)
			})

			t.Run("Hour()", func(t *testing.T) {
				hour, ok := parsed.Hour()
				checkDateHelperOutput(t, "Hour()", parsed.Time.Hour(), hour, tc.HasHour, ok)
			})

			t.Run("Minute()", func(t *testing.T) {
				minute, ok := parsed.Minute()
				checkDateHelperOutput(t, "Minute()", parsed.Time.Minute(), minute, tc.HasMinute, ok)
			})

			t.Run("Second()", func(t *testing.T) {
				minute, ok := parsed.Second()
				checkDateHelperOutput(t, "Second()", parsed.Time.Second(), minute, tc.HasSecond, ok)
			})

			t.Run("Nanosecond()", func(t *testing.T) {
				nanos, ok := parsed.Nanosecond()
				checkDateHelperOutput(t, "Nanosecond()", parsed.Time.Nanosecond(), nanos, tc.HasNanosecond, ok)
			})

			t.Run("HasPrecision()", func(t *testing.T) {
				checkHasPrecision(t, parsed, tc.HasPrecisionRange, dtPrecisionOmits)
			})

		})
	}
}

func TestParseDatetimeErr(t *testing.T) {
	testCases := []struct {
		Name     string
		BadValue string
	}{
		{
			Name:     "TotallyWrong",
			BadValue: "notaDT",
		},
		{
			Name:     "ContainsValidHead",
			BadValue: "10100203040506.456789+0102SomeText",
		},
		{
			Name:     "ContainsValidHead_LineBreak",
			BadValue: "10100203040506.456789+0102\nSomeText",
		},
		{
			Name:     "ContainsValidHead_WhiteSpace",
			BadValue: "10100203040506.456789+0102 SomeText",
		},
		{
			Name:     "ContainsValidTail",
			BadValue: "SomeText10100203040506.456789+0102",
		},
		{
			Name:     "ContainsValidTail_LineBreak",
			BadValue: "SomeText\n10100203040506.456789+0102",
		},
		{
			Name:     "ContainsValidTail_WhiteSpace",
			BadValue: "SomeText 10100203040506.456789+0102",
		},
		{
			Name:     "ExtraDigit_TZ",
			BadValue: "10100203040506.456789+01023",
		},
		{
			Name:     "MissingDigit_TZ",
			BadValue: "10100203040506.456789+010",
		},
		{
			Name:     "TZ_NoMinutes",
			BadValue: "10100203040506.456789+01",
		},
		{
			Name:     "TZ_SingleDigit",
			BadValue: "10100203040506.456789+1",
		},
		{
			Name:     "TZ_DoubleSign",
			BadValue: "10100203040506.456789++0102",
		},
		{
			Name:     "TZ_BadSign",
			BadValue: "10100203040506.456789&0102",
		},
		{
			Name:     "ExtraDigit_Milliseconds",
			BadValue: "10100203040506.4567891+0102",
		},
		{
			Name:     "ExtraDigit_Milliseconds_NoTZ",
			BadValue: "10100203040506.4567891",
		},
		{
			Name:     "ExtraDigit_Seconds",
			BadValue: "101002030405061.456789+0102",
		},
		{
			Name:     "ExtraDigit_Seconds_NoTZ",
			BadValue: "101002030405061.456789",
		},
		{
			Name:     "ExtraDigit_Seconds_NoMilliseconds",
			BadValue: "101002030405061",
		},
		{
			Name:     "MissingDigit_Seconds",
			BadValue: "1010020304056.456789+0102",
		},
		{
			Name:     "MissingDigit_Seconds_NoTZ",
			BadValue: "1010020304056.456789",
		},
		{
			Name:     "MissingDigit_Seconds_NoMilliseconds",
			BadValue: "1010020304056",
		},
		{
			Name:     "MissingDigit_Minutes",
			BadValue: "10100203045.456789+0102",
		},
		{
			Name:     "MissingDigit_Minutes_NoTZ",
			BadValue: "10100203045.456789",
		},
		{
			Name:     "MissingDigit_Minutes_NoMilliseconds",
			BadValue: "10100203045",
		},
		{
			Name:     "MissingDigit_Hours",
			BadValue: "101002034.456789+0102",
		},
		{
			Name:     "MissingDigit_Hours_NoTZ",
			BadValue: "101002034.456789",
		},
		{
			Name:     "MissingDigit_Hours_NoMilliseconds",
			BadValue: "101002034",
		},
		{
			Name:     "MissingDigit_Days",
			BadValue: "1010023.456789+0102",
		},
		{
			Name:     "MissingDigit_Days_NoTZ",
			BadValue: "1010023.456789",
		},
		{
			Name:     "MissingDigit_Days_NoMilliseconds",
			BadValue: "1010023",
		},
		{
			Name:     "MissingDigit_Months",
			BadValue: "10102.456789+0102",
		},
		{
			Name:     "MissingDigit_Months_NoTZ",
			BadValue: "10102.456789",
		},
		{
			Name:     "MissingDigit_Months_NoMilliseconds",
			BadValue: "10102.456789",
		},
		{
			Name:     "MissingDigit_Years",
			BadValue: "101.456789+0102",
		},
		{
			Name:     "MissingDigit_Years_NoTZ",
			BadValue: "101.456789",
		},
		{
			Name:     "MissingDigit_Years_NoMilliseconds",
			BadValue: "101.456789",
		},
		{
			Name:     "NoMillisecondsWithSeparator",
			BadValue: "10100203040506.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, err := dcmtime.ParseDatetime(tc.BadValue)
			if !errors.Is(err, dcmtime.ErrParseDT) {
				t.Errorf("expected ErrParseDT from ParseDatetime(), got %v", err)
			}
		})
	}
}

func TestDatetime_PrecisionTrimming(t *testing.T) {
	testCases := []struct {
		Name           string
		TimeVal        time.Time
		ExpectedDCM    string
		ExpectedString string
		Precision      dcmtime.PrecisionLevel
		NoOffset       bool
	}{
		{
			Name:           "PrecisionFull-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.456789-0102",
			ExpectedString: "1010-02-03 04:05:06.456789 -01:02",
			Precision:      dcmtime.PrecisionFull,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionFull-WithOffset-MSLeadingZeroes",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.000456-0102",
			ExpectedString: "1010-02-03 04:05:06.000456 -01:02",
			Precision:      dcmtime.PrecisionFull,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionFull-WithOffset-TruncateNanos",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789999, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.456789-0102",
			ExpectedString: "1010-02-03 04:05:06.456789 -01:02",
			Precision:      dcmtime.PrecisionFull,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionFull-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.456789",
			ExpectedString: "1010-02-03 04:05:06.456789",
			Precision:      dcmtime.PrecisionFull,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionMS5-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.45678-0102",
			ExpectedString: "1010-02-03 04:05:06.45678 -01:02",
			Precision:      dcmtime.PrecisionMS5,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionMS5-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.45678",
			ExpectedString: "1010-02-03 04:05:06.45678",
			Precision:      dcmtime.PrecisionMS5,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionMS4-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.4567-0102",
			ExpectedString: "1010-02-03 04:05:06.4567 -01:02",
			Precision:      dcmtime.PrecisionMS4,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionMS4-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.4567",
			ExpectedString: "1010-02-03 04:05:06.4567",
			Precision:      dcmtime.PrecisionMS4,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionMS3-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.456-0102",
			ExpectedString: "1010-02-03 04:05:06.456 -01:02",
			Precision:      dcmtime.PrecisionMS3,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionMS3-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.456",
			ExpectedString: "1010-02-03 04:05:06.456",
			Precision:      dcmtime.PrecisionMS3,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionMS2-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.45-0102",
			ExpectedString: "1010-02-03 04:05:06.45 -01:02",
			Precision:      dcmtime.PrecisionMS2,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionMS2-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.45",
			ExpectedString: "1010-02-03 04:05:06.45",
			Precision:      dcmtime.PrecisionMS2,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionMS1-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.4-0102",
			ExpectedString: "1010-02-03 04:05:06.4 -01:02",
			Precision:      dcmtime.PrecisionMS1,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionMS1-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506.4",
			ExpectedString: "1010-02-03 04:05:06.4",
			Precision:      dcmtime.PrecisionMS1,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionSeconds-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506-0102",
			ExpectedString: "1010-02-03 04:05:06 -01:02",
			Precision:      dcmtime.PrecisionSeconds,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionSeconds-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203040506",
			ExpectedString: "1010-02-03 04:05:06",
			Precision:      dcmtime.PrecisionSeconds,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionMinutes-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "101002030405-0102",
			ExpectedString: "1010-02-03 04:05 -01:02",
			Precision:      dcmtime.PrecisionMinutes,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionMinutes-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "101002030405",
			ExpectedString: "1010-02-03 04:05",
			Precision:      dcmtime.PrecisionMinutes,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionHours-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "1010020304-0102",
			ExpectedString: "1010-02-03 04 -01:02",
			Precision:      dcmtime.PrecisionHours,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionHours-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "1010020304",
			ExpectedString: "1010-02-03 04",
			Precision:      dcmtime.PrecisionHours,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionDay-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203-0102",
			ExpectedString: "1010-02-03 -01:02",
			Precision:      dcmtime.PrecisionDay,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionDay-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "10100203",
			ExpectedString: "1010-02-03",
			Precision:      dcmtime.PrecisionDay,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionMonth-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "101002-0102",
			ExpectedString: "1010-02 -01:02",
			Precision:      dcmtime.PrecisionMonth,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionMonth-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "101002",
			ExpectedString: "1010-02",
			Precision:      dcmtime.PrecisionMonth,
			NoOffset:       true,
		},
		{
			Name:           "PrecisionYear-WithOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "1010-0102",
			ExpectedString: "1010 -01:02",
			Precision:      dcmtime.PrecisionYear,
			NoOffset:       false,
		},
		{
			Name:           "PrecisionYear-NoOffset",
			TimeVal:        time.Date(1010, 2, 3, 4, 5, 6, 456789000, time.FixedZone("", -3720)),
			ExpectedDCM:    "1010",
			ExpectedString: "1010",
			Precision:      dcmtime.PrecisionYear,
			NoOffset:       true,
		},
	}

	for _, tc := range testCases {
		dt := dcmtime.Datetime{
			Time:      tc.TimeVal,
			Precision: tc.Precision,
			NoOffset:  tc.NoOffset,
		}

		// Run one master sub-test for each case
		t.Run(tc.Name, func(t *testing.T) {

			// Break out methods into sub-sub tests
			t.Run("DCM()", func(t *testing.T) {
				dcmVal := dt.DCM()
				if dcmVal != tc.ExpectedDCM {
					t.Errorf("DCM(): expected '%v', got '%v'", tc.ExpectedDCM, dcmVal)
				}
			})

			t.Run("String()", func(t *testing.T) {
				strVal := dt.String()
				if strVal != tc.ExpectedString {
					t.Errorf(
						"String(): expected '%v', got '%v'", tc.ExpectedString, strVal,
					)
				}
			})
		})
	}
}
