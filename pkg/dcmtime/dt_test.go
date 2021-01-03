package dcmtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDT_String(t *testing.T) {
	testCases := []struct {
		TimeVal      time.Time
		Expected     string
		Precision    PrecisionLevel
		IgnoreOffset bool
	}{
		// Full Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.456789-0102",
			Precision:    Precision.Full,
			IgnoreOffset: false,
		},
		// Full Precision, no offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.456789",
			Precision:    Precision.Full,
			IgnoreOffset: true,
		},
		// -1 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.45678-0102",
			Precision:    Precision.MS5,
			IgnoreOffset: false,
		},
		// -1 Precision, no offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.45678",
			Precision:    Precision.MS5,
			IgnoreOffset: true,
		},
		// -2 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.4567-0102",
			Precision:    Precision.MS4,
			IgnoreOffset: false,
		},
		// -2 Precision, no offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.4567",
			Precision:    Precision.MS4,
			IgnoreOffset: true,
		},
		// -3 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.456-0102",
			Precision:    Precision.MS3,
			IgnoreOffset: false,
		},
		// -3 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.456",
			Precision:    Precision.MS3,
			IgnoreOffset: true,
		},
		// -4 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.45-0102",
			Precision:    Precision.MS2,
			IgnoreOffset: false,
		},
		// -4 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.45",
			Precision:    Precision.MS2,
			IgnoreOffset: true,
		},
		// -5 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.4-0102",
			Precision:    Precision.MS1,
			IgnoreOffset: false,
		},
		// -5 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506.4",
			Precision:    Precision.MS1,
			IgnoreOffset: true,
		},
		// -6 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506-0102",
			Precision:    Precision.Seconds,
			IgnoreOffset: false,
		},
		// -6 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203040506",
			Precision:    Precision.Seconds,
			IgnoreOffset: true,
		},
		// -7 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "101002030405-0102",
			Precision:    Precision.Minutes,
			IgnoreOffset: false,
		},
		// -7 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "101002030405",
			Precision:    Precision.Minutes,
			IgnoreOffset: true,
		},
		// -8 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "1010020304-0102",
			Precision:    Precision.Hours,
			IgnoreOffset: false,
		},
		// -8 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "1010020304",
			Precision:    Precision.Hours,
			IgnoreOffset: true,
		},
		// -9 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203-0102",
			Precision:    Precision.Day,
			IgnoreOffset: false,
		},
		// -9 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "10100203",
			Precision:    Precision.Day,
			IgnoreOffset: true,
		},
		// -10 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "101002-0102",
			Precision:    Precision.Month,
			IgnoreOffset: false,
		},
		// -10 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "101002",
			Precision:    Precision.Month,
			IgnoreOffset: true,
		},
		// -11 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "1010-0102",
			Precision:    Precision.Year,
			IgnoreOffset: false,
		},
		// -11 Precision, with offset
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:     "1010",
			Precision:    Precision.Year,
			IgnoreOffset: true,
		},
	}

	for _, tc := range testCases {
		name := tc.Expected
		if tc.IgnoreOffset {
			name += "_NoOffset"
		}
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			daVal := NewDT(tc.TimeVal, tc.Precision, tc.IgnoreOffset)
			assert.Equal(
				tc.Expected, daVal.String(), "output equals expected",
			)
		})
	}
}
