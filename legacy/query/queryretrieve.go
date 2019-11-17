package query

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gobwas/glob"
	"github.com/suyashkumar/dicom/dicomtag"
	"github.com/suyashkumar/dicom/element"
)

func querySequence(elem *element.Element, f *element.Element) (match bool, err error) {
	// TODO(saito) Implement!
	return true, nil
}

func matchString(pattern string, value string) (bool, error) {
	g, err := glob.Compile(pattern)
	if err != nil {
		return false, err
	}
	return g.Match(value), nil
}

func queryElement(elem *element.Element, f *element.Element) (match bool, err error) {
	if isEmptyQuery(f) {
		// Universal match
		return true, nil
	}
	if f.VR == "SQ" {
		return querySequence(f, elem)
	}
	if elem == nil {
		// TODO(saito) This is probably wrong. We shouldn't distinguish between
		// nonexistent element and an empty element.
		return false, err
	}
	if f.VR != elem.VR {
		// This shouldn't happen, but be paranoid
		return false, fmt.Errorf("VR mismatch: filter %v, value %v", f, elem)
	}
	if f.VR == "UI" {
		// See if elem contains at last one uid listed in the filter.
		for _, expected := range f.Value {
			e := expected.(string)
			for _, value := range elem.Value {
				if value.(string) == e {
					return true, nil
				}
			}
		}
		return false, nil
	}
	// TODO: handle date-range matches
	switch v := f.Value[0].(type) {
	case int32:
		for _, value := range elem.Value {
			if v == value.(int32) {
				return true, nil
			}
		}
	case int16:
		for _, value := range elem.Value {
			if v == value.(int16) {
				return true, nil
			}
		}
	case uint32:
		for _, value := range elem.Value {
			if v == value.(uint32) {
				return true, nil
			}
		}
	case uint16:
		for _, value := range elem.Value {
			if v == value.(uint16) {
				return true, nil
			}
		}
	case float32:
		for _, value := range elem.Value {
			if v == value.(float32) {
				return true, nil
			}
		}
	case float64:
		for _, value := range elem.Value {
			if v == value.(float64) {
				return true, nil
			}
		}

	case string:
		for _, value := range elem.Value {
			ok, err := matchString(v, value.(string))
			if err != nil {
				return false, err
			}
			return ok, nil
		}
	default:
		panic(fmt.Sprintf("Unknown data: %v", f))
	}
	return false, nil
}

func isEmptyQuery(f *element.Element) bool {
	// Check if the glob pattern is a sequence of '*'s.
	// "*" is the same as an empty query. P3.4, C2.2.2.4.
	isUniversalGlob := func(s string) bool {
		for i := 0; i < len(s); i++ {
			if s[i] != '*' {
				return false
			}
		}
		return true
	}

	if len(f.Value) == 0 {
		return true
	}
	switch dicomtag.GetVRKind(f.Tag, f.VR) {
	case dicomtag.VRBytes:
		if len(f.Value[0].([]byte)) == 0 {
			return true
		}
	case dicomtag.VRString, dicomtag.VRDate:
		pattern := f.Value[0].(string)
		if len(pattern) == 0 {
			return true
		}
		if isUniversalGlob(pattern) {
			return true
		}
	case dicomtag.VRStringList:
		pattern := f.Value[0].(string)
		if isUniversalGlob(pattern) {
			return true
		}
	}
	return false
}

// Query checks if the dataset matches a QR condition "f". If so, it returns the
// <true, matched element, nil>. If "f" asks for a universal match (i.e., empty
// query value), and the element for f.Tag doesn't exist, the function returns
// <true, nil, nil>. If "f" is malformed, the function returns <false, nil,
// error reason>.
func Query(ds *element.DataSet, f *element.Element) (match bool, matchedElem *element.Element, err error) {
	if len(f.Value) > 1 {
		// A filter can't contain multiple values. P3.4, C.2.2.2.1
		return false, nil, fmt.Errorf("Multiple values found in filter '%v'", f)
	}
	if f.Tag == dicomtag.QueryRetrieveLevel || f.Tag == dicomtag.SpecificCharacterSet {
		return true, nil, nil
	}
	elem, err := ds.FindElementByTag(f.Tag)
	if err != nil {
		elem = nil
	}
	match, err = queryElement(elem, f)
	if match {
		return true, elem, nil
	}
	return false, nil, err
}

const (
	InvalidYear  = -1
	InvalidMonth = -1
	InvalidDay   = -1
)

// DateInfo is a result of parsing a date string.
type DateInfo struct {
	// Input string.
	Str string
	// Results of parsing Str
	Year  int // Year (CE), in range [0,9999]. E.g., 2015
	Month int // Month of year, in range [1,12].
	Day   int // Day of month, in range [1,31].
}

func (d DateInfo) String() string {
	// Convert to ISO-8601
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

var (
	dateRangeRE = regexp.MustCompile("^([\\d.]*)-([\\d.]*)$")

	// A date string can be 8-byte Date string of form 19930822 or 10-byte
	// ACR-NEMA300 string of form "1993.08.22". The latter is not compliant
	// according to P3.5 6.2, but it still happens in real life.
	newDateRE = regexp.MustCompile("^(\\d\\d\\d\\d)(\\d\\d)(\\d\\d)$")
	oldDateRE = regexp.MustCompile("^(\\d\\d\\d\\d)\\.(\\d\\d)\\.(\\d\\d)$")
)

// ParseDate parses a date string or date-range string as defined for the VR
// type "DA". See
// https://www.medicalconnections.co.uk/kb/Dicom_Query_DateTime_range.
//
// If "s" is for a point in time, startDate will show that point, and endDate
// will be {Year:-1,Month:-1,Day:-1}. Is "s" is for a range of dates, [startDate,
// endDate] stores the range (note: both ends are closed, even though the DICOM
// spec isn't clear about it.).
func ParseDate(s string) (startDate, endDate DateInfo, err error) {
	m := dateRangeRE.FindStringSubmatch(s)
	if m != nil { // Date range.
		if m[1] == "" {
			startDate = DateInfo{"", 0, 1, 1}
		} else {
			startDate, _, err = ParseDate(m[1])
			if err != nil {
				return startDate, endDate, err
			}
		}
		if m[2] == "" {
			endDate = DateInfo{"", 9999, 12, 31}
		} else {
			endDate, _, err = ParseDate(m[2])
			if err != nil {
				return startDate, endDate, err
			}
		}
		return startDate, endDate, nil
	}

	parseInt := func(v string) int {
		i, e := strconv.Atoi(v)
		if e != nil {
			if err == nil {
				err = e
			}
			return -1
		}
		return i
	}
	m = newDateRE.FindStringSubmatch(s)
	if m == nil {
		m = oldDateRE.FindStringSubmatch(s)
		if m == nil {
			return startDate, endDate, fmt.Errorf("%s: Failed to parse as date", s)
		}
	}
	startDate = DateInfo{s, parseInt(m[1]), parseInt(m[2]), parseInt(m[3])}
	endDate = DateInfo{"", InvalidYear, InvalidMonth, InvalidDay}
	return startDate, endDate, err
}
