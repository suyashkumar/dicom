package dcmtime

import (
	"fmt"
	"time"
)

// Parse a date string.
func ExampleParseDate() {
	// This is a DA value like we would expect
	daString := "20201210"

	// We are parsing the date string without allowing nema
	da, err := ParseDate(daString)
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
	da, err := ParseDate(daString)
	if err != nil {
		panic(err)
	}

	// The resulting da value contains a native time.Time value.
	fmt.Println("TIME MONTH :", da.Time.Month())
	// It also reports the precision, of the value. This value is Precision.Month,
	// so we know that even though da.Time.Day() will equal 1, we should disregard it.

	fmt.Println("PRECISION  :", da.Precision)

	// This date is not a NEMA-300 date.
	fmt.Println("IS NEMA    :", da.IsNEMA)

	// Our Date value has some methods similar to time.Time's methods, but also
	// returns presence information since not all DICOM dates contain all date
	// components.
	//
	// Try to get the Month value. Our value included a month, so 'ok' will be true.
	if month, ok := da.Month(); ok {
		fmt.Println("MONTH      :", month)
	}

	// Try to get the Day value. Because minutes are not included, 'ok' will be false
	// and this will not print.
	if minute, ok := da.Day(); ok {
		fmt.Println("DAY:", minute)
	}

	// We can also easily check if the value contains a certain precision:
	hasMonth := da.HasPrecision(PrecisionMonth)
	fmt.Println("HAS MONTH  :", hasMonth)

	// Output:
	// TIME MONTH : December
	// PRECISION  : MONTH
	// IS NEMA    : false
	// MONTH      : December
	// HAS MONTH  : true
}

// Parse a NEMA date string.
func ExampleParseDate_nema() {
	// This is a DA value using the legacy NEMA-format.
	daString := "2020.12.10"

	da, err := ParseDate(daString)
	if err != nil {
		panic(err)
	}

	fmt.Println("TIME VALUE:", da.Time)
	fmt.Println("PRECISION :", da.Precision)
	// This is a NEMA-300 date.
	fmt.Println("IS NEMA   :", da.IsNEMA)

	// Output:
	// TIME VALUE: 2020-12-10 00:00:00 +0000 +0000
	// PRECISION : FULL
	// IS NEMA   : true
}

func ExampleDate_create() {
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

func ExampleDate_createNEMA300() {
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

func ExampleParseTime_precisionMinute() {
	// This is a TM value like we would expect for 12:35
	tmString := "1235"

	tm, err := ParseTime(tmString)
	if err != nil {
		panic(err)
	}

	fmt.Println("TIME VALUE   :", tm.Time)
	fmt.Println("PRECISION    :", tm.Precision)

	// Our Time value has some methods similar to time.Time's methods, but also
	// returns presence information since not all DICOM times contain all time
	// components.
	//
	// Try to get the Minute value. Our value included a day, so 'ok' will be true.
	if day, ok := tm.Minute(); ok {
		fmt.Println("MINUTE VALUE :", day)
	}

	// Try to get the Second value. Because minutes are not included, 'ok' will be false
	// and this will not print.
	if minute, ok := tm.Second(); ok {
		fmt.Println("SECOND VALUE :", minute)
	}

	// We can also easily check if the value contains a certain precision:
	hasSeconds := tm.HasPrecision(PrecisionSeconds)
	fmt.Println("HAS SECONDS  :", hasSeconds)

	// Output:
	// TIME VALUE   : 0001-01-01 12:35:00 +0000 +0000
	// PRECISION    : MINUTES
	// MINUTE VALUE : 35
	// HAS SECONDS  : false
}

func ExampleTime_create() {
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
	fmt.Println("NO OFFSET :", dt.NoOffset)

	// Output:
	// TIME VALUE: 2020-12-10 12:30:01.000431 +0100 +0100
	// PRECISION : FULL
	// NO OFFSET : false
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
	fmt.Println("NO OFFSET :", dt.NoOffset)

	// Output:
	// TIME VALUE: 2020-12-10 12:30:01.000431 +0000 +0000
	// PRECISION : FULL
	// NO OFFSET : true
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

	fmt.Println("TIME VALUE  :", dt.Time)
	fmt.Println("PRECISION   :", dt.Precision)
	fmt.Println("NO OFFSET   :", dt.NoOffset)

	// Our Datetime value has some methods similar to time.Time's methods, but also
	// returns presence information since not all DICOM datetimes contain all datetime
	// components.
	//
	// Try to get the Day value. Our value included a day, so 'ok' will be true
	if day, ok := dt.Day(); ok {
		fmt.Println("DAY VALUE   :", day)
	}

	// Try to get the Minute value. Because minutes are not included. 'ok' will be false
	// and this will not print.
	if minute, ok := dt.Minute(); ok {
		fmt.Println("MINUTE VALUE :", minute)
	}

	// We can also easily check if the value contains a certain precision:
	hasMinutes := dt.HasPrecision(PrecisionMinutes)
	fmt.Println("HAS MINUTES :", hasMinutes)

	// Output:
	// TIME VALUE  : 2020-12-10 12:00:00 +0000 +0000
	// PRECISION   : HOURS
	// NO OFFSET   : true
	// DAY VALUE   : 10
	// HAS MINUTES : false
}

func ExampleDatetime_create() {
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

func ExampleDatetime_createNoOffset() {
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
