package dicom_test

import (
	"testing"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomtag"
)

func TestParse0(t *testing.T) {
	ds, err := dicom.ReadDataSetFromFile("examples/I_000013.dcm", dicom.ReadOptions{})
	if err != nil {
		t.Fatal(err)
	}
	studyUID := "1.2.840.113857.1907.192833.1115.220048"
	match, elem, err := dicom.Query(ds, dicom.MustNewElement(dicomtag.StudyInstanceUID, studyUID))
	if !match || err != nil {
		t.Error(err)
	}
	if elem.MustGetString() != studyUID {
		t.Error(elem)
	}
}

func TestParseDate(t *testing.T) {
	goodDateRanges := []struct {
		str          string
		startISO8601 string
		endISO8601   string
	}{
		{"20170927-20170929", "2017-09-27", "2017-09-29"},
		{"-20170929", "0000-01-01", "2017-09-29"},
		{"20170927-", "2017-09-27", "9999-12-31"},
	}
	for _, r := range goodDateRanges {
		s, e, err := dicom.ParseDate(r.str)
		if err != nil {
			t.Error(err)
		}
		if s.String() != r.startISO8601 || e.String() != r.endISO8601 {
			t.Errorf("%+v: %+v %+v", r, s, e)
		}
	}
	goodDates := []struct {
		str     string
		iso8601 string
	}{
		{"20170101", "2017-01-01"},
		{"2017.02.03", "2017-02-03"},
	}
	for _, goodDate := range goodDates {
		s, e, err := dicom.ParseDate(goodDate.str)
		if err != nil {
			t.Errorf("%v: err %v", goodDate, err)
		}
		if s.String() != goodDate.iso8601 || e.Year != dicom.InvalidYear {
			t.Errorf("%v: wrong value %v", goodDate, s.String())
		}
	}
	badDates := []string{"2017.0101", "2017.01", "2017X01.02", "201X0405"}
	for _, badDate := range badDates {
		_, _, err := dicom.ParseDate(badDate)
		if err == nil {
			t.Errorf("%s: this should have failed", badDate)
		}
	}
}
