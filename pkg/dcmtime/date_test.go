package dcmtime_test

import (
	"errors"
	"github.com/suyashkumar/dicom/pkg/dcmtime"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	testCases := []struct {
		Name              string
		DAValue           string
		ExpectedString    string
		Expected          time.Time
		ExpectedPrecision dcmtime.PrecisionLevel
	}{
		{
			Name:              "PrecisionFull",
			DAValue:           "20200304",
			ExpectedString:    "2020-03-04",
			Expected:          time.Date(2020, 3, 4, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionFull,
		},
		{
			Name:              "PrecisionMonth",
			DAValue:           "202003",
			ExpectedString:    "2020-03",
			Expected:          time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMonth,
		},
		{
			Name:              "PrecisionYear",
			DAValue:           "2020",
			ExpectedString:    "2020",
			Expected:          time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionYear,
		},
		{
			Name:              "PrecisionFullNEMA",
			DAValue:           "2020.03.04",
			ExpectedString:    "2020-03-04",
			Expected:          time.Date(2020, 3, 4, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionFull,
		},
		{
			Name:              "PrecisionMonthNEMA",
			DAValue:           "2020.03",
			ExpectedString:    "2020-03",
			Expected:          time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMonth,
		},
		{
			Name:              "PrecisionYearNEMA",
			DAValue:           "2020",
			ExpectedString:    "2020",
			Expected:          time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionYear,
		},
	}

	for _, tc := range testCases {
		// Run a master test on the test case
		t.Run(tc.Name, func(t *testing.T) {

			// We'll store the parsed object here for subsequent subtests
			var parsed dcmtime.Date

			t.Run("ParseDate", func(t *testing.T) {
				var err error
				parsed, err = dcmtime.ParseDate(tc.DAValue)
				if err != nil {
					t.Fatal("parse err:", err)
				}

				if !tc.Expected.Equal(parsed.Time) {
					t.Errorf(
						"parsed time (%v) != expected (%v) from source DA '%v'",
						parsed.Time,
						tc.Expected,
						tc.DAValue,
					)

				}

				if parsed.Precision != tc.ExpectedPrecision {
					t.Errorf(
						"precision: expected %v, got %v from source DA '%v'",
						tc.ExpectedPrecision.String(),
						parsed.Precision.String(),
						tc.DAValue,
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
				if dcmVal != tc.DAValue {
					t.Fatalf(
						"got DCM() value '%v', expected '%v'", dcmVal, tc.DAValue,
					)
				}
			})
		})

	}
}

func TestParseDateErr(t *testing.T) {
	testCases := []struct {
		Name     string
		BadValue string
	}{
		{
			Name:     "TooManyDigits",
			BadValue: "101002034",
		},
		{
			Name:     "MissingDigit_Days",
			BadValue: "1010023",
		},
		{
			Name:     "MissingDigit_Months",
			BadValue: "10102",
		},
		{
			Name:     "MissingDigit_Year",
			BadValue: "101",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			_, err := dcmtime.ParseDate(tc.BadValue)
			if !errors.Is(err, dcmtime.ErrParseDA) {
				t.Errorf("expected ErrParseDA, got %v", err)
			}
		})
	}
}

// TestDate_DCMTrimming tests that dates are trimmed to the correct precision on DCM
// out
func TestDate_DCMTrimming(t *testing.T) {
	testCases := []struct {
		Name      string
		Time      time.Time
		Precision dcmtime.PrecisionLevel
		Expected  string
	}{
		{
			Name:      "PrecisionFull",
			Time:      time.Date(1010, 2, 3, 5, 6, 7, 8, time.UTC),
			Precision: dcmtime.PrecisionFull,
			Expected:  "10100203",
		},
		{
			Name:      "PrecisionDay",
			Time:      time.Date(1010, 2, 3, 5, 6, 7, 8, time.UTC),
			Precision: dcmtime.PrecisionDay,
			Expected:  "10100203",
		},
		{
			Name:      "PrecisionMonth",
			Time:      time.Date(1010, 2, 3, 5, 6, 7, 8, time.UTC),
			Precision: dcmtime.PrecisionMonth,
			Expected:  "101002",
		},
		{
			Name:      "PrecisionYear",
			Time:      time.Date(1010, 2, 3, 5, 6, 7, 8, time.UTC),
			Precision: dcmtime.PrecisionYear,
			Expected:  "1010",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			da := dcmtime.Date{
				Time:      tc.Time,
				Precision: tc.Precision,
			}

			if da.DCM() != tc.Expected {
				t.Errorf(
					"Date.DCM(): expected '%v', got '%v'", tc.Expected, da.DCM(),
				)
			}
		})
	}
}

// TestDate_SaneDefaults tests that instantiating a new Date object with just the Time
// field specified yields a reasonable result.
func TestDate_SaneDefaults(t *testing.T) {
	newValue := dcmtime.Date{
		Time: time.Date(2021, 03, 16, 0, 0, 0, 0, time.FixedZone("", 0)),
	}

	dcmVal := newValue.DCM()
	expected := "20210316"
	if dcmVal != expected {
		t.Errorf("DCM(): expected '%v', but got '%v'", expected, dcmVal)
	}
}
