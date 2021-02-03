package dcmtime

import (
	"errors"
	"fmt"
	"time"
)

// Parse a date string.
func ExampleParseDate() {
	// This is a DA value like we would expect
	daString := "20201210"

	// We are parsing the date string without allowing nema
	da, err := ParseDate(daString, false)
	if err != nil {
		panic(err)
	}

	fmt.Println("TIME VALUE:", da.Time)
	fmt.Println("PRECISION :", da.Precision)

	// Output:
	// TIME VALUE: 2020-12-10 00:00:00 +0000 +0000
	// PRECISION : FULL
}

// Parse a date string missing dat values.
func ExampleParseDate_lessPrecision() {
	// This is a DA value like we would expect, but it is missing the day value.
	daString := "202012"

	// We are parsing the date string without allowing NEMA-300 formatted dates.
	da, err := ParseDate(daString, false)
	if err != nil {
		panic(err)
	}

	// The resulting da value contains a native time.Time value.
	fmt.Println("TIME MONTH:", da.Time.Month())
	// It also reports the precision, of the value. This value is Precision.Month,
	// so we know that even though da.Time.Day() will equal 1, we should disregard it.

	fmt.Println("PRECISION :", da.Precision)

	// This date is not a NEMA-300 date.
	fmt.Println("IS NEMA   :", da.IsNEMA)

	// Output:
	// TIME MONTH: December
	// PRECISION : MONTH
	// IS NEMA   : false
}

// Parse a NEMA date string.
func ExampleParseDate_nema() {
	// This is a DA value like we would expect, but it is missing the day value.
	daString := "2020.12.10"

	// We are parsing the date string without allowing nema. This will result in an
	// error.
	_, err := ParseDate(daString, false)
	if err == nil {
		panic("should have error")
	}
	fmt.Println("ERROR     :", err)

	// This error is a ErrParseDA with some extra context, which we can test like so:
	if !errors.Is(err, ErrParseDA) {
		panic("bad error test")
	}

	// But if we allow NEMA-formatted values...
	da, err := ParseDate(daString, true)
	if err != nil {
		panic(err)
	}

	fmt.Println("TIME VALUE:", da.Time)
	fmt.Println("PRECISION :", da.Precision)
	// This is a NEMA-300 date.
	fmt.Println("IS NEMA   :", da.IsNEMA)

	// Output:
	// ERROR     : error parsing dicom DA (date) value -- expected format is 'YYYYMMDD'. for more details on proper DA value formatting, see here: http://dicom.nema.org/medical/dicom/current/output/html/part05.html#table_6.2-1
	// TIME VALUE: 2020-12-10 00:00:00 +0000 +0000
	// PRECISION : FULL
	// IS NEMA   : true
}

func ExampleDate() {
	// We'll use the reference date as our date
	date, err := time.Parse(
		"Mon Jan 2, 2006",
		"Mon Jan 2, 2006",
	)
	if err != nil {
		panic(err)
	}

	// Create a nw DA object like so:
	da := Date{
		Time:      date,
		Precision: PrecisionFull,
		IsNEMA:    false,
	}

	// Get the DICOM string value
	fmt.Println("DCM   :", da.DCM())

	// Our String() method will yield a more readable non-DICOM-compliant value.
	fmt.Println("STRING:", da.String())

	// Output:
	// DCM   : 20060102
	// STRING: 2006-01-02
}

func ExampleDate_nema300() {
	// We'll use the reference date as our date
	date, err := time.Parse(
		"Mon Jan 2, 2006",
		"Mon Jan 2, 2006",
	)
	if err != nil {
		panic(err)
	}

	// Create a nw DA object like so:
	da := Date{
		Time:      date,
		Precision: PrecisionFull,
		IsNEMA:    true,
	}

	// Get the DICOM string value, this will have period separators ala NEMA-300.
	fmt.Println("DCM   :", da.DCM())

	// Our String() method will yield a more readable non-DICOM-compliant value.
	fmt.Println("STRING:", da.String())

	// Output:
	// DCM   : 2006.01.02
	// STRING: 2006-01-02
}

func ExampleDate_precisionYear() {
	// We'll use the reference date as our date
	date, err := time.Parse("2006-01", "2006-01")
	if err != nil {
		panic(err)
	}

	// Create a nw DA object that only represent the year like so:
	da := Date{
		Time:      date,
		Precision: PrecisionMonth,
		IsNEMA:    false,
	}

	// Get the DICOM string value
	fmt.Println("DCM   :", da.DCM())

	// Our String() method will yield a more readable non-DICOM-compliant value.
	fmt.Println("STRING:", da.String())

	// Output:
	// DCM   : 200601
	// STRING: 2006-01
}

func ExampleParseTime() {
	// This is a TM value like we would expect for 12:30:01 and 400 microseconds
	tmString := "123001.000431"

	tm, err := ParseTime(tmString)
	if err != nil {
		panic(err)
	}

	fmt.Println("TIME VALUE:", tm.Time)
	fmt.Println("PRECISION :", tm.Precision)

	// Output:
	// TIME VALUE: 0001-01-01 12:30:01.000431 +0000 +0000
	// PRECISION : FULL
}

func ExampleParseTime_precisionMS() {
	// This is a TM value like we would expect for 12:30:01 and 400 microseconds
	tmString := "123001.431"

	tm, err := ParseTime(tmString)
	if err != nil {
		panic(err)
	}

	fmt.Println("TIME VALUE:", tm.Time)
	fmt.Println("PRECISION :", tm.Precision)

	// Output:
	// TIME VALUE: 0001-01-01 12:30:01.431 +0000 +0000
	// PRECISION : MS3
}

func ExampleParseTime_precisionHour() {
	// This is a TM value like we would expect for 12:30:01 and 400 microseconds
	tmString := "12"

	tm, err := ParseTime(tmString)
	if err != nil {
		panic(err)
	}

	fmt.Println("TIME VALUE:", tm.Time)
	fmt.Println("PRECISION :", tm.Precision)

	// Output:
	// TIME VALUE: 0001-01-01 12:00:00 +0000 +0000
	// PRECISION : HOURS
}

func ExampleTime() {
	// We'll use the reference date as our date
	timeVal, err := time.Parse(
		"15:04:05.000",
		"15:04:05.431",
	)
	if err != nil {
		panic(err)
	}

	// Create a nw TM object like so:
	tm := Time{
		Time:      timeVal,
		Precision: PrecisionFull,
	}

	// Get the DICOM string value
	fmt.Println("DCM   :", tm.DCM())

	// Our String() method will yield a more readable non-DICOM-compliant value.
	fmt.Println("STRING:", tm.String())

	// Output:
	// DCM   : 150405.431000
	// STRING: 15:04:05.431000
}

func ExampleTime_precision3MS() {
	// We'll use the reference date as our date
	timeVal, err := time.Parse(
		"15:04:05.000",
		"15:04:05.431",
	)
	if err != nil {
		panic(err)
	}

	// Create a nw TM object like so:
	tm := Time{
		Time:      timeVal,
		Precision: PrecisionMS3,
	}

	// Get the DICOM string value
	fmt.Println("DCM   :", tm.DCM())

	// Our String() method will yield a more readable non-DICOM-compliant value.
	fmt.Println("STRING:", tm.String())

	// Output:
	// DCM   : 150405.431
	// STRING: 15:04:05.431
}

func ExampleTime_precisionMinutes() {
	// We'll use the reference date as our date
	timeVal, err := time.Parse(
		"15:04",
		"15:04",
	)
	if err != nil {
		panic(err)
	}

	// Create a nw TM object like so:
	tm := Time{
		Time:      timeVal,
		Precision: PrecisionMinutes,
	}

	// Get the DICOM string value
	fmt.Println("DCM   :", tm.DCM())

	// Our String() method will yield a more readable non-DICOM-compliant value.
	fmt.Println("STRING:", tm.String())

	// Output:
	// DCM   : 1504
	// STRING: 15:04
}

// Parse a datetime string.
func ExampleParseDatetime() {
	// This is a DT value like we would expect
	daString := "20201210123001.000431+0100"

	// We are parsing the date string without allowing nema
	dt, err := ParseDatetime(daString)
	if err != nil {
		panic(err)
	}

	fmt.Println("TIME VALUE:", dt.Time)
	fmt.Println("PRECISION :", dt.Precision)
	fmt.Println("HAS OFFSET:", dt.NoOffset)

	// Output:
	// TIME VALUE: 2020-12-10 12:30:01.000431 +0100 +0100
	// PRECISION : FULL
	// HAS OFFSET: true
}

// Parse a datetime string with no timezone.
func ExampleParseDatetime_noTimezone() {
	// This is a DT value like we would expect
	daString := "20201210123001.000431"

	// We are parsing the date string without allowing nema
	dt, err := ParseDatetime(daString)
	if err != nil {
		panic(err)
	}

	fmt.Println("TIME VALUE:", dt.Time)
	fmt.Println("PRECISION :", dt.Precision)
	fmt.Println("HAS OFFSET:", dt.NoOffset)

	// Output:
	// TIME VALUE: 2020-12-10 12:30:01.000431 +0000 +0000
	// PRECISION : FULL
	// HAS OFFSET: false
}

// Parse a datetime string with no timezone.
func ExampleParseDatetime_precisionHour() {
	// This is a DT value like we would expect
	daString := "2020121012"

	// We are parsing the date string without allowing nema
	dt, err := ParseDatetime(daString)
	if err != nil {
		panic(err)
	}

	fmt.Println("TIME VALUE:", dt.Time)
	fmt.Println("PRECISION :", dt.Precision)
	fmt.Println("HAS OFFSET:", dt.NoOffset)

	// Output:
	// TIME VALUE: 2020-12-10 12:00:00 +0000 +0000
	// PRECISION : HOURS
	// HAS OFFSET: false
}

func ExampleDatetime() {
	// We'll use the reference date as our date
	timeVal, err := time.Parse(
		"2006-01-02T15:04:05.000000-07:00",
		"2006-01-02T15:04:05.123456+01:00",
	)
	if err != nil {
		panic(err)
	}

	// Create a nw TM object like so:
	dt := Datetime{
		Time:      timeVal,
		Precision: PrecisionFull,
		NoOffset:  false,
	}

	// Get the DICOM string value
	fmt.Println("DCM   :", dt.DCM())

	// Our String() method will yield a more readable non-DICOM-compliant value.
	fmt.Println("STRING:", dt.String())

	// Output:
	// DCM   : 20060102150405.123456+0100
	// STRING: 2006-01-02 15:04:05.123456 +01:00
}

func ExampleDatetime_noOffset() {
	// We'll use the reference date as our date
	timeVal, err := time.Parse(
		"2006-01-02T15:04:05.000000",
		"2006-01-02T15:04:05.123456",
	)
	if err != nil {
		panic(err)
	}

	// Create a nw TM object like so:
	dt := Datetime{
		Time:      timeVal,
		Precision: PrecisionFull,
		NoOffset:  true,
	}

	// Get the DICOM string value
	fmt.Println("DT:", dt.DCM())

	// Output:
	// DT: 20060102150405.123456
}

func ExampleDatetime_precisionMinute() {
	// We'll use the reference date as our date
	timeVal, err := time.Parse(
		"2006-01-02T15:04",
		"2006-01-02T15:04",
	)
	if err != nil {
		panic(err)
	}

	// Create a nw TM object like so:
	dt := Datetime{
		Time:      timeVal,
		Precision: PrecisionMinutes,
		NoOffset:  true,
	}

	// Get the DICOM string value
	fmt.Println("DT:", dt.DCM())

	// Output:
	// DT: 200601021504
}

func ExampleDatetime_precisionMinuteWithOffset() {
	// We'll use the reference date as our date
	timeVal, err := time.Parse(
		"2006-01-02T15:04-07:00",
		"2006-01-02T15:04+01:00",
	)
	if err != nil {
		panic(err)
	}

	// Create a nw TM object like so:
	dt := Datetime{
		Time:      timeVal,
		Precision: PrecisionMinutes,
		NoOffset:  false,
	}

	// Get the DICOM string value
	fmt.Println("DT:", dt.DCM())

	// Output:
	// DT: 200601021504+0100
}
