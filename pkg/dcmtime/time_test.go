package dcmtime

import (
	"testing"
	"time"
)

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
