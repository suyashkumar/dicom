package personname

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInfo(t *testing.T) {
	testCases := []struct {
		// The Raw string to parse from
		Raw string
		// The parsed information we expect.
		Expected Info
		// Whether RemoveTrailingEmpty should be set to true when creating a new
		// GroupInfo to match Raw.
		RemoveTrailingEmpty bool
		// Whether IsEmpty should return true after parsing Raw.
		IsEmpty bool
	}{
		// All groups
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix=" +
				"pFamily^pGiven^pMiddle^pPrefix^pSuffix",
			Expected: Info{
				Alphabetic: GroupInfo{
					Raw:        "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
					FamilyName: "aFamily",
					GivenName:  "aGiven",
					MiddleName: "aMiddle",
					NamePrefix: "aPrefix",
					NameSuffix: "aSuffix",
				},
				Ideographic: GroupInfo{
					Raw:        "iFamily^iGiven^iMiddle^iPrefix^iSuffix",
					FamilyName: "iFamily",
					GivenName:  "iGiven",
					MiddleName: "iMiddle",
					NamePrefix: "iPrefix",
					NameSuffix: "iSuffix",
				},
				Phonetic: GroupInfo{
					Raw:        "pFamily^pGiven^pMiddle^pPrefix^pSuffix",
					FamilyName: "pFamily",
					GivenName:  "pGiven",
					MiddleName: "pMiddle",
					NamePrefix: "pPrefix",
					NameSuffix: "pSuffix",
				},
			},
			RemoveTrailingEmpty: false,
			IsEmpty:             false,
		},
		// No Phonetic
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix=" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					Raw:        "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
					FamilyName: "aFamily",
					GivenName:  "aGiven",
					MiddleName: "aMiddle",
					NamePrefix: "aPrefix",
					NameSuffix: "aSuffix",
				},
				Ideographic: GroupInfo{
					Raw:        "iFamily^iGiven^iMiddle^iPrefix^iSuffix",
					FamilyName: "iFamily",
					GivenName:  "iGiven",
					MiddleName: "iMiddle",
					NamePrefix: "iPrefix",
					NameSuffix: "iSuffix",
				},
				Phonetic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
			},
			RemoveTrailingEmpty: false,
			IsEmpty:             false,
		},
		// No Phonetic, no seps
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					Raw:        "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
					FamilyName: "aFamily",
					GivenName:  "aGiven",
					MiddleName: "aMiddle",
					NamePrefix: "aPrefix",
					NameSuffix: "aSuffix",
				},
				Ideographic: GroupInfo{
					Raw:        "iFamily^iGiven^iMiddle^iPrefix^iSuffix",
					FamilyName: "iFamily",
					GivenName:  "iGiven",
					MiddleName: "iMiddle",
					NamePrefix: "iPrefix",
					NameSuffix: "iSuffix",
				},
				Phonetic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
			},
			RemoveTrailingEmpty: true,
			IsEmpty:             false,
		},
		// No Ideographic
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=" +
				"=" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					Raw:        "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
					FamilyName: "aFamily",
					GivenName:  "aGiven",
					MiddleName: "aMiddle",
					NamePrefix: "aPrefix",
					NameSuffix: "aSuffix",
				},
				Ideographic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
				Phonetic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
			},
			RemoveTrailingEmpty: false,
			IsEmpty:             false,
		},
		// No Ideographic, no seps
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
			Expected: Info{
				Alphabetic: GroupInfo{
					Raw:        "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
					FamilyName: "aFamily",
					GivenName:  "aGiven",
					MiddleName: "aMiddle",
					NamePrefix: "aPrefix",
					NameSuffix: "aSuffix",
				},
				Ideographic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
				Phonetic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
			},
			RemoveTrailingEmpty: true,
			IsEmpty:             false,
		},
		// Empty with seps
		{
			Raw: "=" +
				"=" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
				Ideographic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
				Phonetic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
			},
			RemoveTrailingEmpty: false,
			IsEmpty:             true,
		},
		// Empty no seps
		{
			Raw: "",
			Expected: Info{
				Alphabetic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
				Ideographic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
				Phonetic: GroupInfo{
					Raw:        "",
					FamilyName: "",
					GivenName:  "",
					MiddleName: "",
					NamePrefix: "",
					NameSuffix: "",
				},
			},
			RemoveTrailingEmpty: true,
			IsEmpty:             true,
		},
	}

	for _, tc := range testCases {
		tc.Expected.Raw = tc.Raw

		// Test creating a new Info object.
		t.Run(tc.Raw+"_new", func(t *testing.T) {
			newInfo := New(
				tc.Expected.Alphabetic,
				tc.Expected.Ideographic,
				tc.Expected.Phonetic,
				tc.RemoveTrailingEmpty,
			)

			assert.Equal(t, tc.Raw, newInfo.String())
		})

		// Test pasring a full PN value.
		t.Run(tc.Raw+"_parse", func(t *testing.T) {
			assert := assert.New(t)

			parsed, err := Parse(tc.Raw)
			if !assert.NoError(err, "parse Raw") {
				t.FailNow()
			}

			checkGroupInfo(
				t,
				tc.Expected.Alphabetic,
				parsed.Alphabetic,
				"Alphabetic",
			)
			checkGroupInfo(
				t,
				tc.Expected.Ideographic,
				parsed.Ideographic,
				"Ideographic",
			)
			checkGroupInfo(
				t,
				tc.Expected.Phonetic,
				parsed.Phonetic,
				"Phonetic",
			)
		})

		// Test the .IsEmpty() method.
		t.Run(tc.Raw+"_isEmpty", func(t *testing.T) {
			assert := assert.New(t)

			newInfo, err := Parse(tc.Raw)
			if !assert.NoError(err, "parse Raw") {
				t.FailNow()
			}

			assert.Equal(tc.IsEmpty, newInfo.IsEmpty())
		})
	}
}

func TestParse_Err(t *testing.T) {
	testCases := []struct {
		Raw       string
		ErrString string
	}{
		{
			Raw: "===",
			ErrString: "string contains to many segments for PN value:" +
				" PN contains 4 groups. No more than 3 groups with" +
				" '[Alphabetic]=[Ideographic]=[Phonetic]' format are allowed",
		},
		{
			Raw: "^^^^^",
			ErrString: "string contains to many segments for PN value: PN group" +
				" Alphabetic contains 6 segments. No more than 5 segments with" +
				" '[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed",
		},
		{
			Raw: "=^^^^^",
			ErrString: "string contains to many segments for PN value: PN group" +
				" Ideographic contains 6 segments. No more than 5 segments with" +
				" '[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed",
		},
		{
			Raw: "==^^^^^",
			ErrString: "string contains to many segments for PN value: PN group" +
				" Phonetic contains 6 segments. No more than 5 segments with" +
				" '[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Raw, func(t *testing.T) {
			assert := assert.New(t)

			_, err := Parse(tc.Raw)

			// Test errors.Is() with ErrParsePersonName
			assert.True(
				errors.Is(err, ErrParsePersonName),
				"error unwraps to ErrParsePersonName",
			)

			// Test full error string is correct.
			assert.EqualError(
				err,
				tc.ErrString,
			)
		})
	}
}
