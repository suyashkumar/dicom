package dcmtime

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	testCases := []struct {
		DAValue           string
		Expected          time.Time
		ExpectedPrecision PrecisionLevel
		AllowNema         bool
	}{
		// DICOM full date
		{
			DAValue: "20200304",
			Expected: time.Date(
				2020,
				3,
				4,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionFull,
			AllowNema:         false,
		},
		// DICOM no day
		{
			DAValue: "202003",
			Expected: time.Date(
				2020,
				3,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMonth,
			AllowNema:         false,
		},
		// DICOM no month
		{
			DAValue: "2020",
			Expected: time.Date(
				2020,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionYear,
			AllowNema:         false,
		},
		// NEMA full date
		{
			DAValue: "2020.03.04",
			Expected: time.Date(
				2020,
				3,
				4,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionFull,
			AllowNema:         true,
		},
		// NEMA no day
		{
			DAValue: "2020.03",
			Expected: time.Date(
				2020,
				3,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionMonth,
			AllowNema:         true,
		},
		// NEMA no month
		{
			DAValue: "2020",
			Expected: time.Date(
				2020,
				1,
				1,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			ExpectedPrecision: PrecisionYear,
			AllowNema:         true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.DAValue, func(t *testing.T) {
			parsed, err := ParseDate(tc.DAValue, tc.AllowNema)
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
