package dcmtime

import (
	"testing"
	"time"
)

func TestParseDatetime(t *testing.T) {
	testCases := []struct {
		DTValue           string
		Expected          time.Time
		ExpectedPrecision PrecisionLevel
		HasOffset         bool
	}{
		// Full value, positive offset
		{
			DTValue: "10100203040506.456789+0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionFull,
			HasOffset:         true,
		},
		// Full value, minus 1 millisecond places, positive offset
		{
			DTValue: "10100203040506.45678+0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456780000,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionMS5,
			HasOffset:         true,
		},
		// Full value, minus 2 millisecond places, positive offset
		{
			DTValue: "10100203040506.4567+0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456700000,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionMS4,
			HasOffset:         true,
		},
		// Full value, minus 3 millisecond places, positive offset
		{
			DTValue: "10100203040506.456+0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456000000,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionMS3,
			HasOffset:         true,
		},
		// Full value, minus 4 millisecond places, positive offset
		{
			DTValue: "10100203040506.45+0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				450000000,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionMS2,
			HasOffset:         true,
		},
		// Full value, minus 5 millisecond places, positive offset
		{
			DTValue: "10100203040506.4+0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				400000000,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionMS1,
			HasOffset:         true,
		},
		// Full value, no millisecond, positive offset
		{
			DTValue: "10100203040506+0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				0,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionSeconds,
			HasOffset:         true,
		},
		// Full value, no seconds, positive offset
		{
			DTValue: "101002030405+0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				0,
				0,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionMinutes,
			HasOffset:         true,
		},
		// Full value, no minutes, positive offset
		{
			DTValue: "1010020304+0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				0,
				0,
				0,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionHours,
			HasOffset:         true,
		},
		// Full value, no hours, positive offset
		{
			DTValue: "10100203+0102",
			Expected: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				0,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionDay,
			HasOffset:         true,
		},
		// Full value, no hours, positive offset
		{
			DTValue: "101002+0102",
			Expected: time.Date(
				1010,
				2,
				1,
				0,
				0,
				0,
				0,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionMonth,
			HasOffset:         true,
		},
		// Full value, no hours, positive offset
		{
			DTValue: "1010+0102",
			Expected: time.Date(
				1010,
				1,
				1,
				0,
				0,
				0,
				0,
				time.FixedZone("", 3720),
			),
			ExpectedPrecision: PrecisionYear,
			HasOffset:         true,
		},
		// Full value, negative offset
		{
			DTValue: "10100203040506.456789-0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionFull,
			HasOffset:         true,
		},
		// Full value, negative offset, minus 1 millisecond places
		{
			DTValue: "10100203040506.45678-0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456780000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionMS5,
			HasOffset:         true,
		},
		// Full value, negative offset, minus 2 millisecond places
		{
			DTValue: "10100203040506.4567-0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456700000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionMS4,
			HasOffset:         true,
		},
		// Full value, negative offset, minus 3 millisecond places
		{
			DTValue: "10100203040506.456-0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456000000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionMS3,
			HasOffset:         true,
		},
		// Full value, negative offset, minus 4 millisecond places
		{
			DTValue: "10100203040506.45-0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				450000000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionMS2,
			HasOffset:         true,
		},
		// Full value, negative offset, minus 5 millisecond places
		{
			DTValue: "10100203040506.4-0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				400000000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionMS1,
			HasOffset:         true,
		},
		// Full value, negative offset, no milliseconds
		{
			DTValue: "10100203040506-0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				000000000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionSeconds,
			HasOffset:         true,
		},
		// Full value, negative offset, no seconds
		{
			DTValue: "101002030405-0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				0,
				000000000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionMinutes,
			HasOffset:         true,
		},
		// Full value, negative offset, no minutes
		{
			DTValue: "1010020304-0102",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				0,
				0,
				000000000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionHours,
			HasOffset:         true,
		},
		// Full value, negative offset, no hours
		{
			DTValue: "10100203-0102",
			Expected: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				000000000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionDay,
			HasOffset:         true,
		},
		// Full value, negative offset, no days
		{
			DTValue: "101002-0102",
			Expected: time.Date(
				1010,
				2,
				1,
				0,
				0,
				0,
				000000000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionMonth,
			HasOffset:         true,
		},
		// Full value, negative offset, no month
		{
			DTValue: "1010-0102",
			Expected: time.Date(
				1010,
				1,
				1,
				0,
				0,
				0,
				000000000,
				time.FixedZone("", -3720),
			),
			ExpectedPrecision: PrecisionYear,
			HasOffset:         true,
		},
		// Full value, no offset
		{
			DTValue: "10100203040506.456789",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionFull,
			HasOffset:         false,
		},
		// Full value, no offset, minus 1 millisecond place
		{
			DTValue: "10100203040506.45678",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456780000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMS5,
			HasOffset:         false,
		},
		// Full value, no offset, minus 2 millisecond place
		{
			DTValue: "10100203040506.4567",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456700000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMS4,
			HasOffset:         false,
		},
		// Full value, no offset, minus 3 millisecond place
		{
			DTValue: "10100203040506.456",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456000000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMS3,
			HasOffset:         false,
		},
		// Full value, no offset, minus 4 millisecond place
		{
			DTValue: "10100203040506.45",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				450000000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMS2,
			HasOffset:         false,
		},
		// Full value, no offset, minus 5 millisecond place
		{
			DTValue: "10100203040506.4",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				400000000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMS1,
			HasOffset:         false,
		},
		// Full value, no offset, no milliseconds
		{
			DTValue: "10100203040506",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionSeconds,
			HasOffset:         false,
		},
		// Full value, no offset, no seconds
		{
			DTValue: "101002030405",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				5,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMinutes,
			HasOffset:         false,
		},
		// Full value, no offset, no minutes
		{
			DTValue: "1010020304",
			Expected: time.Date(
				1010,
				2,
				3,
				4,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionHours,
			HasOffset:         false,
		},
		// Full value, no offset, no seconds
		{
			DTValue: "10100203",
			Expected: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionDay,
			HasOffset:         false,
		},
		// Full value, no offset, no days
		{
			DTValue: "101002",
			Expected: time.Date(
				1010,
				2,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMonth,
			HasOffset:         false,
		},
		// Full value, no offset, no month
		{
			DTValue: "1010",
			Expected: time.Date(
				1010,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionYear,
			HasOffset:         false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.DTValue, func(t *testing.T) {
			parsed, err := ParseDatetime(tc.DTValue)
			if err != nil {
				t.Fatal("parse err:", err)
			}

			if !tc.Expected.Equal(parsed.Time) {
				t.Errorf(
					"parsed time (%v) != expected (%v)",
					parsed.Time,
					tc.Expected,
				)

			}

			if parsed.Precision != tc.ExpectedPrecision {
				t.Errorf(
					"precision: expected %v, got %v",
					tc.ExpectedPrecision.String(),
					parsed.Precision.String(),
				)
			}
		})
	}
}

func TestDatetime_Methods(t *testing.T) {
	testCases := []struct {
		TimeVal        time.Time
		ExpectedDCM    string
		ExpectedString string
		Precision      PrecisionLevel
		NoOffset       bool
	}{
		// Full Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.456789-0102",
			ExpectedString: "1010-02-03 04:05:06.456789 -01:02",
			Precision:      PrecisionFull,
			NoOffset:       false,
		},
		// Full Precision, with offset, fractal leading zero
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.000456-0102",
			ExpectedString: "1010-02-03 04:05:06.000456 -01:02",
			Precision:      PrecisionFull,
			NoOffset:       false,
		},
		// Full Precision, truncate trailing nanos
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789999,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.456789-0102",
			ExpectedString: "1010-02-03 04:05:06.456789 -01:02",
			Precision:      PrecisionFull,
			NoOffset:       false,
		},
		// Full Precision, no offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.456789",
			ExpectedString: "1010-02-03 04:05:06.456789",
			Precision:      PrecisionFull,
			NoOffset:       true,
		},
		// -1 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.45678-0102",
			ExpectedString: "1010-02-03 04:05:06.45678 -01:02",
			Precision:      PrecisionMS5,
			NoOffset:       false,
		},
		// -1 Precision, no offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.45678",
			ExpectedString: "1010-02-03 04:05:06.45678",

			Precision: PrecisionMS5,
			NoOffset:  true,
		},
		// -2 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.4567-0102",
			ExpectedString: "1010-02-03 04:05:06.4567 -01:02",
			Precision:      PrecisionMS4,
			NoOffset:       false,
		},
		// -2 Precision, no offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.4567",
			ExpectedString: "1010-02-03 04:05:06.4567",
			Precision:      PrecisionMS4,
			NoOffset:       true,
		},
		// -3 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.456-0102",
			ExpectedString: "1010-02-03 04:05:06.456 -01:02",
			Precision:      PrecisionMS3,
			NoOffset:       false,
		},
		// -3 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.456",
			ExpectedString: "1010-02-03 04:05:06.456",
			Precision:      PrecisionMS3,
			NoOffset:       true,
		},
		// -4 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.45-0102",
			ExpectedString: "1010-02-03 04:05:06.45 -01:02",
			Precision:      PrecisionMS2,
			NoOffset:       false,
		},
		// -4 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.45",
			ExpectedString: "1010-02-03 04:05:06.45",
			Precision:      PrecisionMS2,
			NoOffset:       true,
		},
		// -5 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.4-0102",
			ExpectedString: "1010-02-03 04:05:06.4 -01:02",
			Precision:      PrecisionMS1,
			NoOffset:       false,
		},
		// -5 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506.4",
			ExpectedString: "1010-02-03 04:05:06.4",
			Precision:      PrecisionMS1,
			NoOffset:       true,
		},
		// -6 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506-0102",
			ExpectedString: "1010-02-03 04:05:06 -01:02",
			Precision:      PrecisionSeconds,
			NoOffset:       false,
		},
		// -6 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203040506",
			ExpectedString: "1010-02-03 04:05:06",
			Precision:      PrecisionSeconds,
			NoOffset:       true,
		},
		// -7 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "101002030405-0102",
			ExpectedString: "1010-02-03 04:05 -01:02",
			Precision:      PrecisionMinutes,
			NoOffset:       false,
		},
		// -7 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "101002030405",
			ExpectedString: "1010-02-03 04:05",
			Precision:      PrecisionMinutes,
			NoOffset:       true,
		},
		// -8 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "1010020304-0102",
			ExpectedString: "1010-02-03 04 -01:02",
			Precision:      PrecisionHours,
			NoOffset:       false,
		},
		// -8 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "1010020304",
			ExpectedString: "1010-02-03 04",
			Precision:      PrecisionHours,
			NoOffset:       true,
		},
		// -9 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203-0102",
			ExpectedString: "1010-02-03 -01:02",
			Precision:      PrecisionDay,
			NoOffset:       false,
		},
		// -9 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "10100203",
			ExpectedString: "1010-02-03",
			Precision:      PrecisionDay,
			NoOffset:       true,
		},
		// -10 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "101002-0102",
			ExpectedString: "1010-02 -01:02",
			Precision:      PrecisionMonth,
			NoOffset:       false,
		},
		// -10 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "101002",
			ExpectedString: "1010-02",
			Precision:      PrecisionMonth,
			NoOffset:       true,
		},
		// -11 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "1010-0102",
			ExpectedString: "1010 -01:02",
			Precision:      PrecisionYear,
			NoOffset:       false,
		},
		// -11 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			ExpectedDCM:    "1010",
			ExpectedString: "1010",
			Precision:      PrecisionYear,
			NoOffset:       true,
		},
	}

	for _, tc := range testCases {
		name := tc.ExpectedDCM
		if tc.NoOffset {
			name += "_NoOffset"
		}
		dt := Datetime{
			Time:      tc.TimeVal,
			Precision: tc.Precision,
			NoOffset:  tc.NoOffset,
		}

		t.Run(name+"_DCM", func(t *testing.T) {
			dcmVal := dt.DCM()
			if dcmVal != tc.ExpectedDCM {
				t.Errorf("DCM(): expected '%v', got '%v'", tc.ExpectedDCM, dcmVal)
			}
		})

		t.Run(name+"_String", func(t *testing.T) {
			strVal := dt.String()
			if strVal != tc.ExpectedString {
				t.Errorf(
					"String(): expected '%v', got '%v'", tc.ExpectedString, strVal,
				)
			}
		})
	}
}
