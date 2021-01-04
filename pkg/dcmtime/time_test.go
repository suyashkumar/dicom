package dcmtime

import (
	"github.com/stretchr/testify/assert"
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
			Precision: Precision.Full,
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
			Precision: Precision.Full,
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
			Precision: Precision.Full,
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
			Precision: Precision.MS5,
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
			Precision: Precision.MS4,
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
			Precision: Precision.MS3,
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
			Precision: Precision.MS2,
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
			Precision: Precision.MS1,
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
			Precision: Precision.Seconds,
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
			Precision: Precision.Minutes,
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
			Precision: Precision.Hours,
			Expected:  "01",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Expected, func(t *testing.T) {
			assert := assert.New(t)

			tm := NewTM(tc.Time, tc.Precision)

			assert.Equal(tc.Expected, tm.DCM(), ".DCM()")
		})
	}
}
