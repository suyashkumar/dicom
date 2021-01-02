package dcmtime

import "errors"

var ErrParseDate = errors.New(
	"error parsing dicom DA (date) value -- expected format is 'YYYYMMDD'. " +
		"for more details on proper DA value formatting, see here: " +
		"http://dicom.nema.org/medical/dicom/current/output/html/part05.html#table_6.2-1",
)

var ErrParseDatetime = errors.New(
	"error parsing dicom DT (datetime) value -- expected format is" +
		" 'YYYYMMDDHHMMSS.FFFFFF&ZZXX'. " +
		"for more details on proper DT value formatting, see here: " +
		"http://dicom.nema.org/medical/dicom/current/output/html/part05.html#table_6.2-1",
)

var ErrParseTime = errors.New(
	"error parsing dicom TM (time) value, but expected format is 'HHMMSS.FFFFFF'. " +
		"for more details on proper TM value formatting, see here: " +
		"http://dicom.nema.org/medical/dicom/current/output/html/part05.html#table_6.2-1",
)
