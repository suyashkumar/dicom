package dcmtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseTM(t *testing.T) {
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
			ExpectedPrecision: Precision.Full,
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
			ExpectedPrecision: Precision.MS5,
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
			ExpectedPrecision: Precision.MS4,
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
			ExpectedPrecision: Precision.MS3,
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
			ExpectedPrecision: Precision.MS2,
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
			ExpectedPrecision: Precision.MS1,
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
			ExpectedPrecision: Precision.Seconds,
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
			ExpectedPrecision: Precision.Minutes,
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
			ExpectedPrecision: Precision.Hours,
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
			ExpectedPrecision: Precision.Full,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.TMValue, func(t *testing.T) {
			assert := assert.New(t)

			parsed, err := ParseTM(tc.TMValue)
			if !assert.NoError(err, "parse TM value") {
				t.FailNow()
			}

			assert.Truef(
				tc.ExpectedTime.Equal(parsed.Time),
				"parsed time (%v) equals expected (%v)",
				parsed,
				tc.ExpectedTime,
			)

			assert.Equalf(
				tc.ExpectedPrecision,
				parsed.Precision,
				"precision. expected %v, got %v",
				tc.ExpectedPrecision.String(),
				parsed.Precision.String(),
			)
		})
	}
}

func TestParseDA(t *testing.T) {
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
			ExpectedPrecision: Precision.Full,
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
			ExpectedPrecision: Precision.Month,
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
			ExpectedPrecision: Precision.Year,
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
			ExpectedPrecision: Precision.Full,
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
			ExpectedPrecision: Precision.Month,
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
			ExpectedPrecision: Precision.Year,
			AllowNema:         true,
		},
	}

	for _, thisCase := range testCases {
		t.Run(thisCase.DAValue, func(t *testing.T) {
			assert := assert.New(t)

			parsed, err := ParseDA(thisCase.DAValue, thisCase.AllowNema)
			if !assert.NoError(err, "parse DA value") {
				t.FailNow()
			}

			assert.Truef(
				thisCase.Expected.Equal(parsed.Time),
				"parsed time (%v) equals expected (%v)",
				parsed.Time,
				thisCase.Expected,
			)

			assert.Equalf(
				thisCase.ExpectedPrecision,
				parsed.Precision,
				"precision. expected %v, got %v",
				thisCase.ExpectedPrecision.String(),
				parsed.Precision.String(),
			)
		})
	}
}

func TestParseDT(t *testing.T) {
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
			ExpectedPrecision: Precision.Full,
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
			ExpectedPrecision: Precision.MS5,
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
			ExpectedPrecision: Precision.MS4,
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
			ExpectedPrecision: Precision.MS3,
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
			ExpectedPrecision: Precision.MS2,
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
			ExpectedPrecision: Precision.MS1,
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
			ExpectedPrecision: Precision.Seconds,
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
			ExpectedPrecision: Precision.Minutes,
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
			ExpectedPrecision: Precision.Hours,
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
			ExpectedPrecision: Precision.Day,
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
			ExpectedPrecision: Precision.Month,
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
			ExpectedPrecision: Precision.Year,
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
			ExpectedPrecision: Precision.Full,
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
			ExpectedPrecision: Precision.MS5,
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
			ExpectedPrecision: Precision.MS4,
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
			ExpectedPrecision: Precision.MS3,
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
			ExpectedPrecision: Precision.MS2,
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
			ExpectedPrecision: Precision.MS1,
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
			ExpectedPrecision: Precision.Seconds,
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
			ExpectedPrecision: Precision.Minutes,
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
			ExpectedPrecision: Precision.Hours,
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
			ExpectedPrecision: Precision.Day,
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
			ExpectedPrecision: Precision.Month,
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
			ExpectedPrecision: Precision.Year,
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
			ExpectedPrecision: Precision.Full,
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
			ExpectedPrecision: Precision.MS5,
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
			ExpectedPrecision: Precision.MS4,
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
			ExpectedPrecision: Precision.MS3,
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
			ExpectedPrecision: Precision.MS2,
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
			ExpectedPrecision: Precision.MS1,
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
			ExpectedPrecision: Precision.Seconds,
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
			ExpectedPrecision: Precision.Minutes,
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
			ExpectedPrecision: Precision.Hours,
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
			ExpectedPrecision: Precision.Day,
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
			ExpectedPrecision: Precision.Month,
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
			ExpectedPrecision: Precision.Year,
			HasOffset:         false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.DTValue, func(t *testing.T) {
			assert := assert.New(t)

			parsed, err := ParseDT(tc.DTValue)
			if !assert.NoError(err, "parse DT value") {
				t.FailNow()
			}

			assert.Truef(
				tc.Expected.Equal(parsed.Time),
				"parsed time (%v) equals expected (%v)",
				parsed,
				tc.Expected,
			)

			assert.Equal(
				tc.ExpectedPrecision,
				parsed.Precision,
				"precision. expected %v, got %v",
				tc.ExpectedPrecision,
				parsed.Precision,
			)

			assert.Equal(
				tc.HasOffset,
				parsed.NoOffset,
				"offset specified",
			)
		})
	}
}
