// Package vrraw offers aliases to all VR abbreviations as defined here:
//
// http://dicom.nema.org/medical/dicom/current/output/html/part05.html#table_6.2-1
//
// Deprecated VRs from older editions are also included
package vrraw

const (
	ApplicationEntity   = "AE"
	AgeString           = "AS"
	AttributeTag        = "AT"
	CodeString          = "CS"
	Date                = "DA"
	DecimalString       = "DS"
	DateTime            = "DT"
	FloatingPointSingle = "FL"
	FloatingPointDouble = "FD"
	IntegerString       = "IS"
	LongString          = "LO"
	LongText            = "LT"
	OtherByte           = "OB"
	OtherDouble         = "OD"
	OtherFloat          = "OF"
	OtherLong           = "OL"
	OtherVeryLong       = "OV"
	OtherWord           = "OW"
	PersonName          = "PN"
	ShortString         = "SH"
	SignedLong          = "SL"
	Sequence            = "SQ"
	SignedShort         = "SS"
	ShortText           = "ST"
	SignedVeryLong      = "SV"
	Time                = "TM"
	UnlimitedCharacters = "UC"

	// NOTE: UniqueIdentifier is also referred to as UID.
	UniqueIdentifier = "UI"

	// NOTE: UniversalResourceIdentifier is also referred to as
	// UniversalResourceLocator, URI and  URL.
	UniversalResourceIdentifier = "UR"

	UnsignedLong     = "UL"
	Unknown          = "UN"
	UnsignedShort    = "US"
	UnlimitedText    = "UT"
	UnsignedVeryLong = "UV"
)
