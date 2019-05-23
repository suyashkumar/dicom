package dicom_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/dicomtag"
	"github.com/suyashkumar/dicom/element"
)

func TestParse0(t *testing.T) {
	p, err := dicom.NewParserFromFile("examples/I_000013.dcm", nil)
	if err != nil {
		t.Fatal(err)
	}
	ds, err := p.Parse(dicom.ParseOptions{})
	if err != nil {
		t.Fatal(err)
	}
	studyUID := "1.2.840.113857.1907.192833.1115.220048"
	match, elem, err := dicom.Query(ds, element.MustNewElement(dicomtag.StudyInstanceUID, studyUID))
	assert.True(t, match)
	assert.NoError(t, err)
	assert.Equal(t, elem.MustGetString(), studyUID)
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
		assert.NoError(t, err)
		assert.Equal(t, s.String(), r.startISO8601)
		assert.Equal(t, e.String(), r.endISO8601)
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
		assert.NoError(t, err, "Date:", goodDate)
		assert.Equal(t, s.String(), goodDate.iso8601)
		assert.Equal(t, e.Year, dicom.InvalidYear)
	}
	badDates := []string{"2017.0101", "2017.01", "2017X01.02", "201X0405"}
	for _, badDate := range badDates {
		_, _, err := dicom.ParseDate(badDate)
		assert.Error(t, err, "Date:", badDate)
	}
}
