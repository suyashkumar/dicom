package dcmtime

import (
	"testing"
	"time"
)

func TestDatetime_DCM(t *testing.T) {
	testCases := []struct {
		TimeVal   time.Time
		Expected  string
		Precision PrecisionLevel
		NoOffset  bool
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
			Expected:  "10100203040506.456789-0102",
			Precision: PrecisionFull,
			NoOffset:  false,
		},
		// Full Precision, with offset, fractal leading zero
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:  "10100203040506.000456-0102",
			Precision: PrecisionFull,
			NoOffset:  false,
		},
		// Full Precision, truncate trailing nanos
		{
			TimeVal: time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789999,
				time.FixedZone(
					"",
					-3720,
				),
			),
			Expected:  "10100203040506.456789-0102",
			Precision: PrecisionFull,
			NoOffset:  false,
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
			Expected:  "10100203040506.456789",
			Precision: PrecisionFull,
			NoOffset:  true,
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
			Expected:  "10100203040506.45678-0102",
			Precision: PrecisionMS5,
			NoOffset:  false,
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
			Expected:  "10100203040506.45678",
			Precision: PrecisionMS5,
			NoOffset:  true,
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
			Expected:  "10100203040506.4567-0102",
			Precision: PrecisionMS4,
			NoOffset:  false,
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
			Expected:  "10100203040506.4567",
			Precision: PrecisionMS4,
			NoOffset:  true,
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
			Expected:  "10100203040506.456-0102",
			Precision: PrecisionMS3,
			NoOffset:  false,
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
			Expected:  "10100203040506.456",
			Precision: PrecisionMS3,
			NoOffset:  true,
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
			Expected:  "10100203040506.45-0102",
			Precision: PrecisionMS2,
			NoOffset:  false,
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
			Expected:  "10100203040506.45",
			Precision: PrecisionMS2,
			NoOffset:  true,
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
			Expected:  "10100203040506.4-0102",
			Precision: PrecisionMS1,
			NoOffset:  false,
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
			Expected:  "10100203040506.4",
			Precision: PrecisionMS1,
			NoOffset:  true,
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
			Expected:  "10100203040506-0102",
			Precision: PrecisionSeconds,
			NoOffset:  false,
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
			Expected:  "10100203040506",
			Precision: PrecisionSeconds,
			NoOffset:  true,
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
			Expected:  "101002030405-0102",
			Precision: PrecisionMinutes,
			NoOffset:  false,
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
			Expected:  "101002030405",
			Precision: PrecisionMinutes,
			NoOffset:  true,
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
			Expected:  "1010020304-0102",
			Precision: PrecisionHours,
			NoOffset:  false,
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
			Expected:  "1010020304",
			Precision: PrecisionHours,
			NoOffset:  true,
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
			Expected:  "10100203-0102",
			Precision: PrecisionDay,
			NoOffset:  false,
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
			Expected:  "10100203",
			Precision: PrecisionDay,
			NoOffset:  true,
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
			Expected:  "101002-0102",
			Precision: PrecisionMonth,
			NoOffset:  false,
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
			Expected:  "101002",
			Precision: PrecisionMonth,
			NoOffset:  true,
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
			Expected:  "1010-0102",
			Precision: PrecisionYear,
			NoOffset:  false,
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
			Expected:  "1010",
			Precision: PrecisionYear,
			NoOffset:  true,
		},
	}

	for _, tc := range testCases {
		name := tc.Expected
		if tc.NoOffset {
			name += "_NoOffset"
		}
		t.Run(name, func(t *testing.T) {
			dt := Datetime{
				Time:      tc.TimeVal,
				Precision: tc.Precision,
				NoOffset:  tc.NoOffset,
			}

			if dt.DCM() != tc.Expected {
				t.Errorf("DCM(): expected '%v', got '%v'", tc.Expected, dt.DCM())
			}
		})
	}
}
