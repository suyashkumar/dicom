package dcmtime_test

import (
	"fmt"
	"github.com/suyashkumar/dicom/pkg/dcmtime"
	"reflect"
	"testing"
	"time"
)

/*
This file is here to ensure that certain dcm types implement certain interfaces.
*/

type hasTime interface {
	GetTime() time.Time
}

type hasPrecision interface {
	GetPrecision() dcmtime.PrecisionLevel
}

type hasYear interface {
	Year() int
}

type hasMonth interface {
	Month() (time.Month, bool)
}

type hasDay interface {
	Day() (int, bool)
}

type hasHour interface {
	Hour() (int, bool)
}

type hasMinute interface {
	Minute() (int, bool)
}

type hasSecond interface {
	Second() (int, bool)
}

type hasNanosecond interface {
	Nanosecond() (int, bool)
}

type dcmVal interface {
	DCM() string
}

type isPrecisionChecker interface {
	HasPrecision(level dcmtime.PrecisionLevel) bool
}

// TestCommonInterfaces ensures that certain types continue to share certain common
// interfaces.
func TestCommonInterfaces(t *testing.T) {
	t.Run("hasTime", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Time{},
			dcmtime.Date{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(hasTime)
				if !ok {
					t.Errorf("%v does not implement hasTime", reflect.TypeOf(tc))
				}
			})
		}
	})

	t.Run("hasPrecision", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Time{},
			dcmtime.Date{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(hasPrecision)
				if !ok {
					t.Errorf("%v does not implement hasPrecision", reflect.TypeOf(tc))
				}
			})
		}
	})

	t.Run("hasYear", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Date{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(hasYear)
				if !ok {
					t.Errorf("%v does not implement hasYear", reflect.TypeOf(tc))
				}
			})
		}
	})

	t.Run("hasMonth", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Date{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(hasMonth)
				if !ok {
					t.Errorf("%v does not implement hasMonth", reflect.TypeOf(tc))
				}
			})
		}
	})

	t.Run("hasDay", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Date{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(hasDay)
				if !ok {
					t.Errorf("%v does not implement hasDay", reflect.TypeOf(tc))
				}
			})
		}
	})

	t.Run("hasHour", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Time{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(hasHour)
				if !ok {
					t.Errorf("%v does not implement hasHour", reflect.TypeOf(tc))
				}
			})
		}
	})

	t.Run("hasMinute", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Time{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(hasMinute)
				if !ok {
					t.Errorf("%v does not implement hasMinute", reflect.TypeOf(tc))
				}
			})
		}
	})

	t.Run("hasSecond", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Time{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(hasSecond)
				if !ok {
					t.Errorf("%v does not implement hasSecond", reflect.TypeOf(tc))
				}
			})
		}
	})

	t.Run("hasNanosecond", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Time{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(hasNanosecond)
				if !ok {
					t.Errorf("%v does not implement hasNanosecond", reflect.TypeOf(tc))
				}
			})
		}
	})

	t.Run("dcmVal", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Date{},
			dcmtime.Time{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(dcmVal)
				if !ok {
					t.Errorf("%v does not implement dcmVal", reflect.TypeOf(tc))
				}
			})
		}
	})

	t.Run("fmt.Stringer", func(t *testing.T) {
		testCases := []interface{}{
			dcmtime.Date{},
			dcmtime.Time{},
			dcmtime.Datetime{},
		}

		for _, tc := range testCases {
			t.Run(reflect.TypeOf(tc).String(), func(t *testing.T) {
				_, ok := tc.(fmt.Stringer)
				if !ok {
					t.Errorf("%v does not implement fmt.Stringer", reflect.TypeOf(tc))
				}
			})
		}
	})
}
