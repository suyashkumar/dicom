package vrraw

// Value Representation abbreviations as defined here:
// http://dicom.nema.org/dicom/2013/output/chtml/part05/sect_6.2.html
const (
	ApplicationEntity = "AE"
	AgeString = "AS"
	AttributeTag = "AT"
	CodeString = "CS"
	Date = "DA"
	DecimalString = "DS"
	DateTime = "DT"
	FloatingPointSingle = "FL"
	FloatingPointDouble = "FD"
	IntegerString = "IS"
	LongString = "LO"
	LongText = "LT"
	OtherByteString = "OB"
	OtherDoubleString = "OD"
	OtherFloatString = "OF"
	OtherWordString = "OW"
	PersonName = "PN"
	ShortString = "SH"
	SignedLong = "SL"
	Sequence = "SQ"
	SignedShort = "SS"
	ShortText = "ST"
	Time = "TM"

	// The spec uses both UniqueIdentifier AND UID in the VR Name column for UI, so
	// we will include both here.
	UniqueIdentifier = "UI"
	UID = UniqueIdentifier

	UnsignedLong = "UL"
	Unknown = "UN"
	UnsignedShort = "US"
	UnlimitedText = "UT"
)
