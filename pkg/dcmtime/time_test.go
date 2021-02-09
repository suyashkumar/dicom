package dcmtime

import (
	"errors"
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
			_, err := ParseTime(tc.BadValue)
			if !errors.Is(err, ErrParseTM) {
				t.Errorf("got %v, expected error value of type ErrParseTM", err)
			}
		})
	}
}

func TestTime_Methods(t *testing.T) {
	testCases := []struct {
		Time           time.Time
		Precision      PrecisionLevel
		ExpectedDCM    string
		ExpectedString string
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
			Precision:      PrecisionFull,
			ExpectedDCM:    "010203.456789",
			ExpectedString: "01:02:03.456789",
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
			Precision:      PrecisionFull,
			ExpectedDCM:    "010203.000456",
			ExpectedString: "01:02:03.000456",
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
			Precision:      PrecisionFull,
			ExpectedDCM:    "010203.456789",
			ExpectedString: "01:02:03.456789",
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
			Precision:      PrecisionMS5,
			ExpectedDCM:    "010203.45678",
			ExpectedString: "01:02:03.45678",
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
			Precision:      PrecisionMS4,
			ExpectedDCM:    "010203.4567",
			ExpectedString: "01:02:03.4567",
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
			Precision:      PrecisionMS3,
			ExpectedDCM:    "010203.456",
			ExpectedString: "01:02:03.456",
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
			Precision:      PrecisionMS2,
			ExpectedDCM:    "010203.45",
			ExpectedString: "01:02:03.45",
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
			Precision:      PrecisionMS1,
			ExpectedDCM:    "010203.4",
			ExpectedString: "01:02:03.4",
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
			Precision:      PrecisionSeconds,
			ExpectedDCM:    "010203",
			ExpectedString: "01:02:03",
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
			Precision:      PrecisionMinutes,
			ExpectedDCM:    "0102",
			ExpectedString: "01:02",
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
			Precision:      PrecisionHours,
			ExpectedDCM:    "01",
			ExpectedString: "01",
		},
	}

	for _, tc := range testCases {
		tm := Time{
			Time:      tc.Time,
			Precision: tc.Precision,
		}

		t.Run(tc.ExpectedDCM+"_DCM", func(t *testing.T) {
			dcmVal := tm.DCM()
			if dcmVal != tc.ExpectedDCM {
				t.Errorf(
					"DCM(): expected '%v', got '%v'", tc.ExpectedDCM, dcmVal,
				)
			}
		})

		t.Run(tc.ExpectedDCM+"_String", func(t *testing.T) {
			strVal := tm.String()
			if strVal != tc.ExpectedString {
				t.Errorf(
					"String(): expected '%v', got '%v'",
					tc.ExpectedString,
					strVal,
				)
			}
		})
	}
}
