package vrraw_test

import (
	"github.com/suyashkumar/dicom/pkg/vrraw"
	"testing"
)

// Tests that the VR consts are what we expect
func TestVRRawValues(t *testing.T) {
	testCases := []struct {
		VR       string
		Expected string
	}{
		{
			VR:       vrraw.ApplicationEntity,
			Expected: "AE",
		},
		{
			VR:       vrraw.AgeString,
			Expected: "AS",
		},
		{
			VR:       vrraw.AttributeTag,
			Expected: "AT",
		},
		{
			VR:       vrraw.CodeString,
			Expected: "CS",
		},
		{
			VR:       vrraw.Date,
			Expected: "DA",
		},
		{
			VR:       vrraw.DecimalString,
			Expected: "DS",
		},
		{
			VR:       vrraw.DateTime,
			Expected: "DT",
		},
		{
			VR:       vrraw.FloatingPointSingle,
			Expected: "FL",
		},
		{
			VR:       vrraw.FloatingPointDouble,
			Expected: "FD",
		},
		{
			VR:       vrraw.IntegerString,
			Expected: "IS",
		},
		{
			VR:       vrraw.LongString,
			Expected: "LO",
		},
		{
			VR:       vrraw.LongText,
			Expected: "LT",
		},
		{
			VR:       vrraw.OtherByte,
			Expected: "OB",
		},
		{
			VR:       vrraw.OtherDouble,
			Expected: "OD",
		},
		{
			VR:       vrraw.OtherFloat,
			Expected: "OF",
		},
		{
			VR:       vrraw.OtherLong,
			Expected: "OL",
		},
		{
			VR:       vrraw.OtherVeryLong,
			Expected: "OV",
		},
		{
			VR:       vrraw.OtherWord,
			Expected: "OW",
		},
		{
			VR:       vrraw.PersonName,
			Expected: "PN",
		},
		{
			VR:       vrraw.ShortString,
			Expected: "SH",
		},
		{
			VR:       vrraw.SignedLong,
			Expected: "SL",
		},
		{
			VR:       vrraw.Sequence,
			Expected: "SQ",
		},
		{
			VR:       vrraw.SignedShort,
			Expected: "SS",
		},
		{
			VR:       vrraw.ShortText,
			Expected: "ST",
		},
		{
			VR:       vrraw.SignedVeryLong,
			Expected: "SV",
		},
		{
			VR:       vrraw.Time,
			Expected: "TM",
		},
		{
			VR:       vrraw.UnlimitedCharacters,
			Expected: "UC",
		},
		{
			VR:       vrraw.UniqueIdentifier,
			Expected: "UI",
		},
		{
			VR:       vrraw.UID,
			Expected: "UI",
		},
		{
			VR:       vrraw.UniversalResourceIdentifier,
			Expected: "UR",
		},
		{
			VR:       vrraw.UniversalResourceLocator,
			Expected: "UR",
		},
		{
			VR:       vrraw.URL,
			Expected: "UR",
		},
		{
			VR:       vrraw.URI,
			Expected: "UR",
		},
		{
			VR:       vrraw.UnsignedLong,
			Expected: "UL",
		},
		{
			VR:       vrraw.Unknown,
			Expected: "UN",
		},
		{
			VR:       vrraw.UnsignedShort,
			Expected: "US",
		},
		{
			VR:       vrraw.UnlimitedText,
			Expected: "UT",
		},
		{
			VR:       vrraw.UnsignedVeryLong,
			Expected: "UV",
		},
	}

	for _, thisCase := range testCases {
		t.Run(thisCase.Expected, func(t *testing.T) {
			if thisCase.Expected != thisCase.VR {
				t.Errorf("expected %v, got %v", thisCase.Expected, thisCase.VR)
			}
		})
	}
}
