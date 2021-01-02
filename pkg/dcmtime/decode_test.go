package dcmtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTMToTime(t *testing.T) {
	testCases := []struct{
		TMValue  string
		Expected time.Time
	}{
		// Full value, leading zeros
		{
			TMValue: "010203.456789",
			Expected: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456789000,
				time.UTC,
			),
		},
		// Remove one millisecond
		{
			TMValue: "010203.45678",
			Expected: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456780000,
				time.UTC,
			),
		},
		// Remove one millisecond
		{
			TMValue: "010203.4567",
			Expected: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456700000,
				time.UTC,
			),
		},
		// Remove one millisecond
		{
			TMValue: "010203.456",
			Expected: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				456000000,
				time.UTC,
			),
		},
		// Remove one millisecond
		{
			TMValue: "010203.45",
			Expected: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				450000000,
				time.UTC,
			),
		},
		// Remove one millisecond
		{
			TMValue: "010203.4",
			Expected: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				400000000,
				time.UTC,
			),
		},
		// No milliseconds
		{
			TMValue: "010203",
			Expected: time.Date(
				0,
				0,
				0,
				1,
				2,
				3,
				0,
				time.UTC,
			),
		},
		// No seconds
		{
			TMValue: "0102",
			Expected: time.Date(
				0,
				0,
				0,
				1,
				2,
				0,
				0,
				time.UTC,
			),
		},
		// No minutes
		{
			TMValue: "01",
			Expected: time.Date(
				0,
				0,
				0,
				1,
				0,
				0,
				0,
				time.UTC,
			),
		},
		// No leading zeroes
		{
			TMValue: "102030.456789",
			Expected: time.Date(
				0,
				0,
				0,
				10,
				20,
				30,
				456789000,
				time.UTC,
			),
		},
	}

	for _, thisCase := range testCases {
		t.Run(thisCase.TMValue, func(t *testing.T) {
			assert := assert.New(t)

			parsed, err := TMToTime(thisCase.TMValue)
			if !assert.NoError(err, "parse TM value") {
				t.FailNow()
			}

			assert.Truef(
				thisCase.Expected.Equal(parsed),
				"parsed (%v) equals expected (%v)",
				parsed,
				thisCase.Expected,
			)
		})
	}
}

func TestDAToTime(t *testing.T) {
	testCases := []struct{
		DAValue  string
		Expected time.Time
		AllowNema bool
	}{
		// DICOM full date
		{
			DAValue:   "20200304",
			Expected:  time.Date(
				2020,
				3,
				4,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			AllowNema: false,
		},
		// DICOM no day
		{
			DAValue:   "202003",
			Expected:  time.Date(
				2020,
				3,
				0,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			AllowNema: false,
		},
		// DICOM no month
		{
			DAValue:   "2020",
			Expected:  time.Date(
				2020,
				0,
				0,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			AllowNema: false,
		},
		// NEMA full date
		{
			DAValue:   "2020.03.04",
			Expected:  time.Date(
				2020,
				3,
				4,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			AllowNema: true,
		},
		// NEMA no day
		{
			DAValue:   "2020.03",
			Expected:  time.Date(
				2020,
				3,
				0,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			AllowNema: true,
		},
		// NEMA no month
		{
			DAValue:   "2020",
			Expected:  time.Date(
				2020,
				0,
				0,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			AllowNema: true,
		},
	}

	for _, thisCase := range testCases {
		t.Run(thisCase.DAValue, func(t *testing.T) {
			assert := assert.New(t)

			parsed, err := DAToTime(thisCase.DAValue, thisCase.AllowNema)
			if !assert.NoError(err, "parse DA value") {
				t.FailNow()
			}

			assert.Truef(
				thisCase.Expected.Equal(parsed),
				"parsed (%v) equals expected (%v)",
				parsed,
				thisCase.Expected,
			)
		})
	}
}

func TestDDTToTime(t *testing.T) {
	testCases := []struct{
		DTValue  string
		Expected time.Time
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
		},
		// Full value, no hours, positive offset
		{
			DTValue: "101002+0102",
			Expected: time.Date(
				1010,
				2,
				0,
				0,
				0,
				0,
				0,
				time.FixedZone("", 3720),
			),
		},
		// Full value, no hours, positive offset
		{
			DTValue: "1010+0102",
			Expected: time.Date(
				1010,
				0,
				0,
				0,
				0,
				0,
				0,
				time.FixedZone("", 3720),
			),
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
		},
		// Full value, negative offset, no days
		{
			DTValue: "101002-0102",
			Expected: time.Date(
				1010,
				2,
				0,
				0,
				0,
				0,
				000000000,
				time.FixedZone("", -3720),
			),
		},
		// Full value, negative offset, no hours
		{
			DTValue: "101002-0102",
			Expected: time.Date(
				1010,
				2,
				0,
				0,
				0,
				0,
				000000000,
				time.FixedZone("", -3720),
			),
		},
		// Full value, negative offset, no hours
		{
			DTValue: "1010-0102",
			Expected: time.Date(
				1010,
				0,
				0,
				0,
				0,
				0,
				000000000,
				time.FixedZone("", -3720),
			),
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
		},
		// Full value, no offset, no days
		{
			DTValue: "101002",
			Expected: time.Date(
				1010,
				2,
				0,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
		// Full value, no offset, no month
		{
			DTValue: "1010",
			Expected: time.Date(
				1010,
				0,
				0,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
	}

	for _, thisCase := range testCases {
		t.Run(thisCase.DTValue, func(t *testing.T) {
			assert := assert.New(t)

			parsed, err := DTToTime(thisCase.DTValue)
			if !assert.NoError(err, "parse DT value") {
				t.FailNow()
			}

			assert.Truef(
				thisCase.Expected.Equal(parsed),
				"parsed (%v) equals expected (%v)",
				parsed,
				thisCase.Expected,
			)
		})
	}
}
