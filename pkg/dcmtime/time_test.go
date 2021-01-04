package dcmtime

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	testCases := []struct {
		TMValue           string
		ExpectedTime      time.Time
		ExpectedPrecision PrecisionLevel
	}{
		// Full value, leading zeros
		{
			TMValue: "010203.456789",
			ExpectedTime: time.Date(
				1,
				1,
				1,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionFull,
		},
		// Remove one millisecond
		{
			TMValue: "010203.45678",
			ExpectedTime: time.Date(
				1,
				1,
				1,
				1,
				2,
				3,
				456780000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMS5,
		},
		// Remove one millisecond
		{
			TMValue: "010203.4567",
			ExpectedTime: time.Date(
				1,
				1,
				1,
				1,
				2,
				3,
				456700000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMS4,
		},
		// Remove one millisecond
		{
			TMValue: "010203.456",
			ExpectedTime: time.Date(
				1,
				1,
				1,
				1,
				2,
				3,
				456000000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMS3,
		},
		// Remove one millisecond
		{
			TMValue: "010203.45",
			ExpectedTime: time.Date(
				1,
				1,
				1,
				1,
				2,
				3,
				450000000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMS2,
		},
		// Remove one millisecond
		{
			TMValue: "010203.4",
			ExpectedTime: time.Date(
				1,
				1,
				1,
				1,
				2,
				3,
				400000000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMS1,
		},
		// No milliseconds
		{
			TMValue: "010203",
			ExpectedTime: time.Date(
				1,
				1,
				1,
				1,
				2,
				3,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionSeconds,
		},
		// No seconds
		{
			TMValue: "0102",
			ExpectedTime: time.Date(
				1,
				1,
				1,
				1,
				2,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMinutes,
		},
		// No minutes
		{
			TMValue: "01",
			ExpectedTime: time.Date(
				1,
				1,
				1,
				1,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionHours,
		},
		// No leading zeroes
		{
			TMValue: "102030.456789",
			ExpectedTime: time.Date(
				1,
				1,
				1,
				10,
				20,
				30,
				456789000,
				time.UTC,
			),
			ExpectedPrecision: PrecisionFull,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TMValue, func(t *testing.T) {

			parsed, err := ParseTime(tc.TMValue)
			if err != nil {
				t.Fatal("parse error:", err)
			}

			if !tc.ExpectedTime.Equal(parsed.Time) {
				t.Errorf(
					"parsed time (%v) != expected (%v)", parsed, tc.ExpectedTime,
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

func TestTime_DCM(t *testing.T) {
	testCases := []struct {
		Time      time.Time
		Precision PrecisionLevel
		Expected  string
	}{
		// Precision.Full
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
			Precision: PrecisionFull,
			Expected:  "010203.456789",
		},
		// Precision.Full, leading zeros
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789,
				time.UTC,
			),
			Precision: PrecisionFull,
			Expected:  "010203.000456",
		},
		// Precision.Full, tail truncated
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789999,
				time.UTC,
			),
			Precision: PrecisionFull,
			Expected:  "010203.456789",
		},
		// Precision.MS5
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
			Precision: PrecisionMS5,
			Expected:  "010203.45678",
		},
		// Precision.MS4
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
			Precision: PrecisionMS4,
			Expected:  "010203.4567",
		},
		// Precision.MS3
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
			Precision: PrecisionMS3,
			Expected:  "010203.456",
		},
		// Precision.MS2
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
			Precision: PrecisionMS2,
			Expected:  "010203.45",
		},
		// Precision.MS1
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
			Precision: PrecisionMS1,
			Expected:  "010203.4",
		},
		// Precision.Seconds
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
			Precision: PrecisionSeconds,
			Expected:  "010203",
		},
		// Precision.Minutes
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
			Precision: PrecisionMinutes,
			Expected:  "0102",
		},
		// Precision.Hours
		{
			Time: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
			Precision: PrecisionHours,
			Expected:  "01",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Expected, func(t *testing.T) {
			tm := Time{
				Time:      tc.Time,
				Precision: tc.Precision,
			}

			if tm.DCM() != tc.Expected {
				t.Errorf("DCM(): expected '%v', got '%v'", tc.Expected, tm.DCM())
			}
		})
	}
}
