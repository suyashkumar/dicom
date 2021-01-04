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
		// Whether NoNullSeparators should be set to true when creating a new
		// Info to match Raw.
		NoNullSeparators bool
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
					FamilyName:       "aFamily",
					GivenName:        "aGiven",
					MiddleName:       "aMiddle",
					NamePrefix:       "aPrefix",
					NameSuffix:       "aSuffix",
					NoNullSeparators: true,
				},
				Ideographic: GroupInfo{
					FamilyName:       "iFamily",
					GivenName:        "iGiven",
					MiddleName:       "iMiddle",
					NamePrefix:       "iPrefix",
					NameSuffix:       "iSuffix",
					NoNullSeparators: true,
				},
				Phonetic: GroupInfo{
					FamilyName:       "pFamily",
					GivenName:        "pGiven",
					MiddleName:       "pMiddle",
					NamePrefix:       "pPrefix",
					NameSuffix:       "pSuffix",
					NoNullSeparators: true,
				},
			},
			NoNullSeparators: false,
			IsEmpty:          false,
		},
		// No Phonetic
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix=" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:       "aFamily",
					GivenName:        "aGiven",
					MiddleName:       "aMiddle",
					NamePrefix:       "aPrefix",
					NameSuffix:       "aSuffix",
					NoNullSeparators: true,
				},
				Ideographic: GroupInfo{
					FamilyName:       "iFamily",
					GivenName:        "iGiven",
					MiddleName:       "iMiddle",
					NamePrefix:       "iPrefix",
					NameSuffix:       "iSuffix",
					NoNullSeparators: true,
				},
				Phonetic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
			},
			NoNullSeparators: false,
			IsEmpty:          false,
		},
		// No Phonetic, no seps
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:       "aFamily",
					GivenName:        "aGiven",
					MiddleName:       "aMiddle",
					NamePrefix:       "aPrefix",
					NameSuffix:       "aSuffix",
					NoNullSeparators: true,
				},
				Ideographic: GroupInfo{
					FamilyName:       "iFamily",
					GivenName:        "iGiven",
					MiddleName:       "iMiddle",
					NamePrefix:       "iPrefix",
					NameSuffix:       "iSuffix",
					NoNullSeparators: true,
				},
				Phonetic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
			},
			NoNullSeparators: true,
			IsEmpty:          false,
		},
		// No Ideographic
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=" +
				"=" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:       "aFamily",
					GivenName:        "aGiven",
					MiddleName:       "aMiddle",
					NamePrefix:       "aPrefix",
					NameSuffix:       "aSuffix",
					NoNullSeparators: true,
				},
				Ideographic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
				Phonetic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
			},
			NoNullSeparators: false,
			IsEmpty:          false,
		},
		// No Ideographic, no seps
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:       "aFamily",
					GivenName:        "aGiven",
					MiddleName:       "aMiddle",
					NamePrefix:       "aPrefix",
					NameSuffix:       "aSuffix",
					NoNullSeparators: true,
				},
				Ideographic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
				Phonetic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
			},
			NoNullSeparators: true,
			IsEmpty:          false,
		},
		// Empty with seps
		{
			Raw: "=" +
				"=" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
				Ideographic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
				Phonetic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
			},
			NoNullSeparators: false,
			IsEmpty:          true,
		},
		// Empty no seps
		{
			Raw: "",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
				Ideographic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
				Phonetic: GroupInfo{
					FamilyName:       "",
					GivenName:        "",
					MiddleName:       "",
					NamePrefix:       "",
					NameSuffix:       "",
					NoNullSeparators: true,
				},
			},
			NoNullSeparators: true,
			IsEmpty:          true,
		},
	}

	for _, tc := range testCases {

		// Test creating a new Info object.
		t.Run(tc.Raw+"_new", func(t *testing.T) {
			newInfo := Info{
				Alphabetic:       tc.Expected.Alphabetic,
				Ideographic:      tc.Expected.Ideographic,
				Phonetic:         tc.Expected.Phonetic,
				NoNullSeparators: tc.NoNullSeparators,
			}

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
				tc.Expected.Alphabetic.String(),
				parsed.Alphabetic,
				"Alphabetic",
			)
			checkGroupInfo(
				t,
				tc.Expected.Ideographic,
				tc.Expected.Ideographic.String(),
				parsed.Ideographic,
				"Ideographic",
			)
			checkGroupInfo(
				t,
				tc.Expected.Phonetic,
				tc.Expected.Phonetic.String(),
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

func TestInfo_WithNullSeparators(t *testing.T) {
	pnVal := "Potter^Harry"

	parsed, err := Parse(pnVal)
	if err != nil {
		t.Errorf("error parsing pnVal: %v", err)
		t.FailNow()
	}

	altered := parsed.WithNullSeparators()

	expected := "Potter^Harry^^^=^^^^=^^^^"
	if altered.String() != expected {
		t.Errorf("expected '%v', got '%v'", expected, altered.String())
	}
}

func TestInfo_WithoutNullSeparators(t *testing.T) {
	pnVal := "Potter^Harry^^^=^^^^=^^^^"

	parsed, err := Parse(pnVal)
	if err != nil {
		t.Errorf("error parsing pnVal: %v", err)
		t.FailNow()
	}

	altered := parsed.WithoutNullSeparators()

	expected := "Potter^Harry"
	if altered.String() != expected {
		t.Errorf("expected '%v', got '%v'", expected, altered.String())
	}
}
