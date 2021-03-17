package dcmtime_test

import (
	"errors"
	"github.com/suyashkumar/dicom/pkg/dcmtime"
	"testing"
	"time"
)

// daPrecisionOmits is the range of precision values not relevant to Date.
var daPrecisionOmits = precisionRange{
	Min: dcmtime.PrecisionHours,
	Max: dcmtime.PrecisionMS5,
}

func TestDate(t *testing.T) {
	testCases := []struct {
		// Name is the name of the test case.
		Name string
		// DAValue is the DICOM string value we are going to parse.
		DAValue string
		// ExpectedString is the expected value of the String() method.
		ExpectedString string
		// ExpectedTime is the expected time.Time value of the parsed value.
		ExpectedTime time.Time
		// ExpectedPrecision is the expected precision value of the parsed value.
		ExpectedPrecision dcmtime.PrecisionLevel
		// HasMonth is whether the parsed value's Month() method should return ok=true
		HasMonth bool
		// HasDay is whether the parsed value's Day() method should return ok=true
		HasDay bool
		// HasPrecisionRange is the range of Precision Values we expect the
		// HasPrecision() method to return true for.
		HasPrecisionRange precisionRange
	}{
		{
			Name:              "PrecisionFull",
			DAValue:           "20200304",
			ExpectedString:    "2020-03-04",
			ExpectedTime:      time.Date(2020, 3, 4, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionFull,
			HasMonth:          true,
			HasDay:            true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionFull,
			},
		},
		{
			Name:              "PrecisionMonth",
			DAValue:           "202003",
			ExpectedString:    "2020-03",
			ExpectedTime:      time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMonth,
			HasMonth:          true,
			HasDay:            false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMonth,
			},
		},
		{
			Name:              "PrecisionYear",
			DAValue:           "2020",
			ExpectedString:    "2020",
			ExpectedTime:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionYear,
			HasMonth:          false,
			HasDay:            false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionYear,
			},
		},
		{
			Name:              "PrecisionFullNEMA",
			DAValue:           "2020.03.04",
			ExpectedString:    "2020-03-04",
			ExpectedTime:      time.Date(2020, 3, 4, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionFull,
			HasMonth:          true,
			HasDay:            true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionFull,
			},
		},
		{
			Name:              "PrecisionMonthNEMA",
			DAValue:           "2020.03",
			ExpectedString:    "2020-03",
			ExpectedTime:      time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMonth,
			HasMonth:          true,
			HasDay:            false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionMonth,
			},
		},
		{
			Name:              "PrecisionYearNEMA",
			DAValue:           "2020",
			ExpectedString:    "2020",
			ExpectedTime:      time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionYear,
			HasMonth:          false,
			HasDay:            false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionYear,
				Max: dcmtime.PrecisionYear,
			},
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

				if !tc.ExpectedTime.Equal(parsed.Time) {
					t.Errorf(
						"parsed time (%v) != expected (%v) from source DA '%v'",
						parsed.Time,
						tc.ExpectedTime,
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

			t.Run("HasPrecision()", func(t *testing.T) {
				checkHasPrecision(t, parsed, tc.HasPrecisionRange, daPrecisionOmits)
			})
		})

	}
}

// checkDateHelperOutput check the output of a helper value getter like Date.Month()
func checkDateHelperOutput(t *testing.T, methodName string, expectedValue int, value int, expectedOK bool, ok bool) {
	if expectedValue != value {
		t.Errorf("got %v int value of '%v', expected '%v'", methodName, value, expectedValue)
	}

	if expectedOK != ok {
		t.Errorf("got %v ok value of '%v', expected '%v'", methodName, ok, expectedOK)
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
