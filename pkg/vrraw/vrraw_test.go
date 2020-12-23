package vrraw_test

import (
	"github.com/suyashkumar/dicom/pkg/vrraw"
	"testing"
)

// Tests that the VR consts are what we expect
func TestVRRawValues(t *testing.T) {
	type TestCase struct {
		VR string
		Expected string
	}

	testCases := []TestCase{
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
			VR:       vrraw.OtherByteString,
			Expected: "OB",
		},
		{
			VR:       vrraw.OtherDoubleString,
			Expected: "OD",
		},
		{
			VR:       vrraw.OtherFloatString,
			Expected: "OF",
		},
		{
			VR:       vrraw.OtherWordString,
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
			VR:       vrraw.Time,
			Expected: "TM",
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
	}

	var thisCase TestCase

	runTest := func(t *testing.T) {
		if thisCase.Expected !=  thisCase.VR {
			t.Errorf("expected %v, got %v", thisCase.Expected, thisCase.VR)
		}
	}

	for _, thisCase = range testCases {
		t.Run(thisCase.Expected, runTest)
	}
}
