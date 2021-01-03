package dcmtime

import "regexp"

// Parses dicom DA (date) of YYYYMMDD
var daRegex = regexp.MustCompile(
	`(?P<YEAR>[0-9]{4})(?P<MONTH>[0-9]{2})?(?P<DAY>[0-9]{2})?`,
)

// Parses old pre-dicom NEMA spec date of YYYY.MM.DD.
var daRegexNema = regexp.MustCompile(
	`(?P<YEAR>[0-9]{4})(?:\.(?P<MONTH>[0-9]{2}))?(?:\.(?P<DAY>[0-9]{2}))?`,
)

// Sub-match group indexes for DA
const daRegexYear = 1
const daRegexMonth = 2
const daRegexDay = 3

// Parses dicom DT (datetime) value of YYYYMMDDHHMMSS.FFFFFF&ZZXX
var dtRegex = regexp.MustCompile(
	`(?P<YEAR>[0-9]{4})` +
		`(?P<MONTH>[0-9]{2})?` +
		`(?P<DAY>[0-9]{2})?` +
		`(?P<HOURS>[0-9]{2})?` +
		`(?P<MINUTES>[0-9]{2})?` +
		`(?P<SECONDS>[0-9]{2})?` +
		`(:?\.(?P<FRACTAL>[0-9]+))?(:?(?P<OFFSET_SIGN>[-+])` +
		`(?P<OFFSET_HOURS>[0-9]{2})(?P<OFFSET_SECONDS>[0-9]{2}))?`,
)

// Sub-match group indexes for DT
const dtRegexYear = 1
const dtRegexMonth = 2
const dtRegexDay = 3
const dtRegexHours = 4
const dtRegexMinutes = 5
const dtRegexSeconds = 6
const dtRegexFractal = 8
const dtRegexOffsetSign = 10
const dtRegexOffsetHours = 11
const dtRegexOffsetMinutes = 12

// Parses dicom TM (time) value of HHMMSS.FFFFFF.
var tmRegex = regexp.MustCompile(
	`(?P<HOURS>[0-9]{2})?(?P<MINUTES>[0-9]{2})?(?P<SECONDS>[0-9]{2})?(?:\.(?P<FRACTAL>[0-9]+))?`,
)

// Sub-match group indexes for TM
const tmRegexHours = 1
const tmRegexMinutes = 2
const tmRegexSeconds = 3
const tmRegexFractal = 4
