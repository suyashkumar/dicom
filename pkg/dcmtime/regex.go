package dcmtime

import "regexp"

// Parses dicom DA (date) of YYYYMMDD
var daRegex = regexp.MustCompile(
	`^(?P<YEAR>[0-9]{4})(?P<MONTH>[0-9]{2})?(?P<DAY>[0-9]{2})?$`,
)

// Parses old pre-dicom NEMA spec date of YYYY.MM.DD.
var daRegexNema = regexp.MustCompile(
	`^(?P<YEAR>[0-9]{4})(?:\.(?P<MONTH>[0-9]{2}))?(?:\.(?P<DAY>[0-9]{2}))?$`,
)

// Sub-match group indexes for DA
const daRegexGroupYear = 1
const daRegexGroupMonth = 2
const daRegexGroupDay = 3

// Parses dicom DT (datetime) value of YYYYMMDDHHMMSS.FFFFFF&ZZXX
var dtRegex = regexp.MustCompile(
	`^(?P<YEAR>[0-9]{4})` +
		`(?P<MONTH>[0-9]{2})?` +
		`(?P<DAY>[0-9]{2})?` +
		`(?P<HOURS>[0-9]{2})?` +
		`(?P<MINUTES>[0-9]{2})?` +
		`(?P<SECONDS>[0-9]{2})?` +
		`(:?\.(?P<FRACTAL>[0-9]{1,6}))?(:?(?P<OFFSET_SIGN>[-+])` +
		`(?P<OFFSET_HOURS>[0-9]{2})(?P<OFFSET_SECONDS>[0-9]{2}))?$`,
)

// Sub-match group indexes for DT
const dtRegexGroupYear = 1
const dtRegexGroupMonth = 2
const dtRegexGroupDay = 3
const dtRegexGroupHours = 4
const dtRegexGroupMinutes = 5
const dtRegexGroupSeconds = 6
const dtRegexGroupFractal = 8
const dtRegexGroupOffsetSign = 10
const dtRegexGroupOffsetHours = 11
const dtRegexGroupOffsetMinutes = 12

// Parses dicom TM (time) value of HHMMSS.FFFFFF.
var tmRegex = regexp.MustCompile(
	`^(?P<HOURS>[0-9]{2})?(?P<MINUTES>[0-9]{2})?(?P<SECONDS>[0-9]{2})?(?:\.(?P<FRACTAL>[0-9]{1,6}))?$`,
)

// Sub-match group indexes for TM
const tmRegexGroupHours = 1
const tmRegexGroupMinutes = 2
const tmRegexGroupSeconds = 3
const tmRegexGroupFractal = 4
