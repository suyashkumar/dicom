package dcmtime

import (
	"testing"
	"time"
)

func TestDate_DCM(t *testing.T) {
	testCases := []struct {
		Time      time.Time
		Precision PrecisionLevel
		Expected  string
	}{
		{
			Time: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			Precision: PrecisionFull,
			Expected:  "10100203",
		},
		{
			Time: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			Precision: PrecisionDay,
			Expected:  "10100203",
		},
		{
			Time: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			Precision: PrecisionMonth,
			Expected:  "101002",
		},
		{
			Time: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			Precision: PrecisionYear,
			Expected:  "1010",
		},
	}

	for _, tc := range testCases {
		name := tc.Expected + "_" + tc.Precision.String()
		t.Run(name, func(t *testing.T) {
			da := Date{
				Time:      tc.Time,
				Precision: tc.Precision,
			}

			if da.DCM() != tc.Expected {
				t.Errorf("DCM(): expected '%v', got '%v'", tc.Expected, da.DCM())
			}
		})
	}
}
