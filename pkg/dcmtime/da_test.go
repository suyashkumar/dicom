package dcmtime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDA_String(t *testing.T) {
	testCases := []struct {
		Time      time.Time
		Precision PrecisionLevel
		Expected  string
	}{
		{
			Time: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			Precision: Precision.Full,
			Expected:  "10100203",
		},
		{
			Time: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			Precision: Precision.Day,
			Expected:  "10100203",
		},
		{
			Time: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			Precision: Precision.Month,
			Expected:  "101002",
		},
		{
			Time: time.Date(
				1010,
				2,
				3,
				0,
				0,
				0,
				0,
				time.UTC,
			),
			Precision: Precision.Year,
			Expected:  "1010",
		},
	}

	for _, tc := range testCases {
		name := tc.Expected + "_" + tc.Precision.String()
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			da := NewDA(tc.Time, tc.Precision)

			assert.Equal(da.DCM(), tc.Expected)
		})
	}
}
