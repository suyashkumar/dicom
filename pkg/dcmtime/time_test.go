package dcmtime_test

import (
	"errors"
	"github.com/suyashkumar/dicom/pkg/dcmtime"
	"testing"
	"time"
)

// daPrecisionOmits is the range of precision values not relevant to Date.
var tmPrecisionOmits = precisionRange{
	Min: dcmtime.PrecisionYear,
	Max: dcmtime.PrecisionDay,
}

func TestTime(t *testing.T) {
	testCases := []struct {
		// Name is the name for the sub-test.
		Name string
		// TMValue is the raw DICOM TM string we are parsing.
		TMValue string
		// ExpectedString is the expected value of the String() method.
		ExpectedTime time.Time
		// ExpectedPrecision is the expected precision value of the parsed value.
		ExpectedPrecision dcmtime.PrecisionLevel
		// ExpectedString is the expected result of the String() value.
		ExpectedString string
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
		// Full value, leading zeros
		{
			Name:              "PrecisionFull",
			TMValue:           "010203.456789",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 456789000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionFull,
			ExpectedString:    "01:02:03.456789",
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionHours,
				Max: dcmtime.PrecisionFull,
			},
		},

		// Remove one millisecond
		{
			Name:              "PrecisionMS5",
			TMValue:           "010203.45678",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 456780000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS5,
			ExpectedString:    "01:02:03.45678",
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionHours,
				Max: dcmtime.PrecisionMS5,
			},
		},

		// Remove two millisecond
		{
			Name:              "PrecisionMS4",
			TMValue:           "010203.4567",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 456700000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS4,
			ExpectedString:    "01:02:03.4567",
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionHours,
				Max: dcmtime.PrecisionMS4,
			},
		},

		// Remove three millisecond
		{
			Name:              "PrecisionMS3",
			TMValue:           "010203.456",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 456000000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS3,
			ExpectedString:    "01:02:03.456",
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionHours,
				Max: dcmtime.PrecisionMS3,
			},
		},

		// Remove four millisecond
		{
			Name:              "PrecisionMS2",
			TMValue:           "010203.45",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 450000000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS2,
			ExpectedString:    "01:02:03.45",
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionHours,
				Max: dcmtime.PrecisionMS2,
			},
		},

		// Remove five millisecond
		{
			Name:              "PrecisionMS1",
			TMValue:           "010203.4",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 400000000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMS1,
			ExpectedString:    "01:02:03.4",
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionHours,
				Max: dcmtime.PrecisionMS1,
			},
		},

		// No milliseconds
		{
			Name:              "PrecisionSeconds",
			TMValue:           "010203",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 3, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionSeconds,
			ExpectedString:    "01:02:03",
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionHours,
				Max: dcmtime.PrecisionSeconds,
			},
		},

		// No seconds
		{
			Name:              "PrecisionMinutes",
			TMValue:           "0102",
			ExpectedTime:      time.Date(1, 1, 1, 1, 2, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionMinutes,
			ExpectedString:    "01:02",
			HasMinute:         true,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionHours,
				Max: dcmtime.PrecisionMinutes,
			},
		},

		// No minutes
		{
			Name:              "PrecisionHours",
			TMValue:           "01",
			ExpectedTime:      time.Date(1, 1, 1, 1, 0, 0, 0, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionHours,
			ExpectedString:    "01",
			HasMinute:         false,
			HasSecond:         false,
			HasNanosecond:     false,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionHours,
				Max: dcmtime.PrecisionHours,
			},
		},

		// No leading zeroes
		{
			Name:              "PrecisionFullNoLeadingZeros",
			TMValue:           "102030.456789",
			ExpectedTime:      time.Date(1, 1, 1, 10, 20, 30, 456789000, time.UTC),
			ExpectedPrecision: dcmtime.PrecisionFull,
			ExpectedString:    "10:20:30.456789",
			HasMinute:         true,
			HasSecond:         true,
			HasNanosecond:     true,
			HasPrecisionRange: precisionRange{
				Min: dcmtime.PrecisionHours,
				Max: dcmtime.PrecisionFull,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			// We'll store the parsed object here for subsequent subtests
			var parsed dcmtime.Time

			t.Run("ParseTime()", func(t *testing.T) {
				var err error
				parsed, err = dcmtime.ParseTime(tc.TMValue)
				if err != nil {
					t.Fatal("ParseTime() error:", err)
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

			t.Run("GetTime()", func(t *testing.T) {
				if !tc.ExpectedTime.Equal(parsed.GetTime()) {
					t.Errorf(
						"Datetime.GetTime(): expected %v, got %v",
						tc.ExpectedTime,
						parsed.Time,
					)
				}
			})

			t.Run("GetPrecision()", func(t *testing.T) {
				if parsed.GetPrecision() != tc.ExpectedPrecision {
					t.Errorf(
						"Datetime.GetPrecision(): expected %v, got %v",
						tc.ExpectedPrecision.String(),
						parsed.Precision.String(),
					)
				}
			})

			t.Run("DCM()", func(t *testing.T) {
				dcmVal := parsed.DCM()
				if dcmVal != tc.TMValue {
					t.Errorf(
						"DCM(): expected '%v', got '%v'", tc.TMValue, dcmVal,
					)
				}
			})

			t.Run("String()", func(t *testing.T) {
				strVal := parsed.String()
				if strVal != tc.ExpectedString {
					t.Errorf(
						"String(): expected '%v', got '%v'",
						tc.ExpectedString,
						strVal,
					)
				}
			})

			t.Run("Hour()", func(t *testing.T) {
				hour, ok := parsed.Hour()
				checkDateHelperOutput(t, "Hour()", parsed.Time.Hour(), hour, true, ok)
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
				checkHasPrecision(t, parsed, tc.HasPrecisionRange, tmPrecisionOmits)
			})
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

func TestTime_PrecisionTrimming(t *testing.T) {
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
