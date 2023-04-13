package dcmtime_test

import (
	"errors"
	"testing"
	"time"

	"github.com/amitbet/dicom/pkg/dcmtime"
)

func TestParseTime(t *testing.T) {
	testCases := []struct {
		Name              string
		TMValue           string
		ExpectedTime      time.Time
		ExpectedPrecision dcmtime.PrecisionLevel
	}{
		// Full value, leading zeros
		{
			Name:              "PrecisionFull",
			TMValue:           "010203.456789",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 456789000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionFull,
		},

		// Remove one millisecond
		{
			Name:              "PrecisionMS5",
			TMValue:           "010203.45678",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 456780000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS5,
		},

		// Remove two millisecond
		{
			Name:              "PrecisionMS4",
			TMValue:           "010203.4567",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 456700000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS4,
		},

		// Remove three millisecond
		{
			Name:              "PrecisionMS3",
			TMValue:           "010203.456",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 456000000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS3,
		},

		// Remove four millisecond
		{
			Name:              "PrecisionMS2",
			TMValue:           "010203.45",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 450000000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS2,
		},

		// Remove five millisecond
		{
			Name:              "PrecisionMS1",
			TMValue:           "010203.4",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 400000000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS1,
		},

		// No milliseconds
		{
			Name:              "PrecisionSeconds",
			TMValue:           "010203",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionSeconds,
		},

		// No seconds
		{
			Name:              "PrecisionMinutes",
			TMValue:           "0102",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMinutes,
		},

		// No minutes
		{
			Name:              "PrecisionHours",
			TMValue:           "01",
			ExpectedTime:      time.Date(1, 1, 1, 1, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionHours,
		},

		// No leading zeroes
		{
			Name:              "PrecisionFullNoLeadingZeros",
			TMValue:           "102030.456789",
			ExpectedTime:      time.Date(1, 1, 1, 10, 20, 30, 456789000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionFull,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TMValue, func(t *testing.T) {

			parsed, err := dcmtime.ParseTime(tc.TMValue)
			if err != nil {
				t.Fatal("parse error:", err)
			}

			if !tc.ExpectedTime.Equal(parsed.Time) {
				t.Errorf(
					"parsed Time (%v) != expected (%v)", parsed, tc.ExpectedTime,
				)
			}

			if parsed.Precision != tc.ExpectedPrecision {
				t.Errorf(
					"Time.Precision: expected %v, got %v",
					tc.ExpectedPrecision.String(),
					parsed.Precision.String(),
				)
			}
		})
	}
}

func TestParseTimeErr(t *testing.T) {
	badValues := []struct {
		Name     string
		BadValue string
	}{
		{
			Name:     "Totally Wrong",
			BadValue: "NotADate",
		},
		{
			Name:     "ValidAtHead",
			BadValue: "120304.12345SomeText",
		},
		{
			Name:     "ValidAtHead_LineBreak",
			BadValue: "120304.12345\nSomeText",
		},
		{
			Name:     "ValidAtHead_WhiteSpace",
			BadValue: "120304.12345 SomeText",
		},
		{
			Name:     "ValidAtTail",
			BadValue: "SomeText120304.12345",
		},
		{
			Name:     "ValidAtTail_LineBreak",
			BadValue: "SomeText\n120304.12345",
		},
		{
			Name:     "ExtraDigit_Milliseconds",
			BadValue: "010304.1234567",
		},
		{
			Name:     "NoMillisecondsWithSeparator",
			BadValue: "010304.",
		},
		{
			Name:     "ExtraDigit_Seconds",
			BadValue: "01030405.345",
		},
		{
			Name:     "ExtraDigit_Seconds_NoMilliseconds",
			BadValue: "01030405",
		},
		{
			Name:     "MissingDigit_Seconds",
			BadValue: "01034.123456",
		},
		{
			Name:     "MissingDigit_Seconds_NoMilliseconds",
			BadValue: "01034",
		},
		{
			Name:     "MissingDigit_Minutes",
			BadValue: "013.123456",
		},
		{
			Name:     "MissingDigit_Minutes_NoMilliseconds",
			BadValue: "013",
		},
		{
			Name:     "MissingDigit_Hours",
			BadValue: "1.123456",
		},
		{
			Name:     "MissingDigit_Hours_NoMilliseconds",
			BadValue: "1",
		},
	}

	for _, tc := range badValues {
		t.Run(tc.Name, func(t *testing.T) {
			_, err := dcmtime.ParseTime(tc.BadValue)
			if !errors.Is(err, dcmtime.ErrParseTM) {
				t.Errorf("got %v, expected error value of type ErrParseTM", err)
			}
		})
	}
}

func TestTime_Methods(t *testing.T) {
	testCases := []struct {
		Name           string
		Time           time.Time
		Precision      dcmtime.PrecisionLevel
		ExpectedDCM    string
		ExpectedString string
	}{
		// PrecisionFull
		{
			Name:           "PrecisionFull",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789000, time.UTC),
			Precision:      dcmtime.PrecisionFull,
			ExpectedDCM:    "010203.456789",
			ExpectedString: "01:02:03.456789",
		},

		// PrecisionFull, leading zeros
		{
			Name:           "PrecisionFullMSLeadingZeros",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789, time.UTC),
			Precision:      dcmtime.PrecisionFull,
			ExpectedDCM:    "010203.000456",
			ExpectedString: "01:02:03.000456",
		},

		// PrecisionFull, tail truncated
		{
			Name:           "PrecisionFull",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789999, time.UTC),
			Precision:      dcmtime.PrecisionFull,
			ExpectedDCM:    "010203.456789",
			ExpectedString: "01:02:03.456789",
		},

		// PrecisionMS5
		{
			Name:           "PrecisionMS5",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789000, time.UTC),
			Precision:      dcmtime.PrecisionMS5,
			ExpectedDCM:    "010203.45678",
			ExpectedString: "01:02:03.45678",
		},

		// PrecisionMS4
		{
			Name:           "PrecisionMS4",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789000, time.UTC),
			Precision:      dcmtime.PrecisionMS4,
			ExpectedDCM:    "010203.4567",
			ExpectedString: "01:02:03.4567",
		},

		// PrecisionMS3
		{
			Name:           "PrecisionMS3",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789000, time.UTC),
			Precision:      dcmtime.PrecisionMS3,
			ExpectedDCM:    "010203.456",
			ExpectedString: "01:02:03.456",
		},

		// PrecisionMS2
		{
			Name:           "PrecisionMS2",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789000, time.UTC),
			Precision:      dcmtime.PrecisionMS2,
			ExpectedDCM:    "010203.45",
			ExpectedString: "01:02:03.45",
		},

		// PrecisionMS1
		{
			Name:           "PrecisionMS1",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789000, time.UTC),
			Precision:      dcmtime.PrecisionMS1,
			ExpectedDCM:    "010203.4",
			ExpectedString: "01:02:03.4",
		},

		// PrecisionSeconds
		{
			Name:           "PrecisionSeconds",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789000, time.UTC),
			Precision:      dcmtime.PrecisionSeconds,
			ExpectedDCM:    "010203",
			ExpectedString: "01:02:03",
		},

		// PrecisionMinutes
		{
			Name:           "PrecisionMinutes",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789000, time.UTC),
			Precision:      dcmtime.PrecisionMinutes,
			ExpectedDCM:    "0102",
			ExpectedString: "01:02",
		},

		// PrecisionHours
		{
			Name:           "PrecisionHours",
			Time:           time.Date(0, 0, 0, 1, 2, 3, 456789000, time.UTC),
			Precision:      dcmtime.PrecisionHours,
			ExpectedDCM:    "01",
			ExpectedString: "01",
		},
	}

	for _, tc := range testCases {
		tm := dcmtime.Time{
			Time:      tc.Time,
			Precision: tc.Precision,
		}

		// Run one test per case with broken out subtests for each method
		t.Run(tc.Name, func(t *testing.T) {

			t.Run("DCM", func(t *testing.T) {
				dcmVal := tm.DCM()
				if dcmVal != tc.ExpectedDCM {
					t.Errorf(
						"DCM(): expected '%v', got '%v'", tc.ExpectedDCM, dcmVal,
					)
				}
			})

			t.Run("String", func(t *testing.T) {
				strVal := tm.String()
				if strVal != tc.ExpectedString {
					t.Errorf(
						"String(): expected '%v', got '%v'",
						tc.ExpectedString,
						strVal,
					)
				}
			})
		})
	}
}

// TestTime_SaneDefaults tests that instantiating a new Time object with just the Time
// field specified yields a reasonable result.
func TestTime_SaneDefaults(t *testing.T) {
	newValue := dcmtime.Time{
		Time: time.Date(1, 1, 1, 12, 7, 56, 123456000, time.FixedZone("", 0)),
	}

	dcmVal := newValue.DCM()
	expected := "120756.123456"
	if dcmVal != expected {
		t.Errorf("DCM(): expected '%v', but got '%v'", expected, dcmVal)
	}
}
