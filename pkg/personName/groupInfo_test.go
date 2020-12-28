package personName

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkGroupInfo(
	t *testing.T,
	expected GroupInfo,
	received GroupInfo,
	group string,
) {
	assert := assert.New(t)

	assert.Equal(
		expected.FamilyName(),
		received.familyName,
		"Family Name, group %v",
		group,
	)

	assert.Equal(
		expected.GivenName(),
		received.givenName,
		"Given Name, group %v",
		group,
	)

	assert.Equal(
		expected.MiddleName(),
		received.middleName,
		"Middle Name, group %v",
		group,
	)

	assert.Equal(
		expected.NamePrefix(),
		received.namePrefix,
		"Name Prefix, group %v",
		group,
	)

	assert.Equal(
		expected.NameSuffix(),
		received.nameSuffix,
		"Name Suffix, group %v",
		group,
	)

	assert.Equal(
		expected.raw,
		received.String(),
		"Formatted String, group %v",
		group,
	)
}

func TestNewPersonNameFromDicom(t *testing.T) {
	type TestCase struct {
		// The raw string to parse from
		Raw string
		// The parsed information we expect.
		Expected GroupInfo
		// Whether removeTrailingSeps should be set to true when creating a new
		// GroupInfo to match Raw.
		RemoveTrailingSeps bool
		// Whether IsEmpty should return true after parsing Raw.
		IsEmpty bool
	}

	testCases := []TestCase{
		// Full Name
		{
			Raw: "CROUCH^BARTEMIUS^'BARTY'^MR^JR",
			Expected: GroupInfo{
				familyName: "CROUCH",
				givenName:  "BARTEMIUS",
				middleName: "'BARTY'",
				namePrefix: "MR",
				nameSuffix: "JR",
			},
		},
		// No Middle Name
		{
			Raw: "CROUCH^BARTEMIUS^^MR^JR",
			Expected: GroupInfo{
				familyName: "CROUCH",
				givenName:  "BARTEMIUS",
				middleName: "",
				namePrefix: "MR",
				nameSuffix: "JR",
			},
		},
		// No Suffix
		{
			Raw: "CROUCH^BARTEMIUS^'BARTY'^MR^",
			Expected: GroupInfo{
				familyName: "CROUCH",
				givenName:  "BARTEMIUS",
				middleName: "'BARTY'",
				namePrefix: "MR",
				nameSuffix: "",
			},
		},
		// No Prefix
		{
			Raw: "CROUCH^BARTEMIUS^'BARTY'^^JR",
			Expected: GroupInfo{
				familyName: "CROUCH",
				givenName:  "BARTEMIUS",
				middleName: "'BARTY'",
				namePrefix: "",
				nameSuffix: "JR",
			},
		},
		// Only First and last
		{
			Raw: "CROUCH^BARTEMIUS^^^",
			Expected: GroupInfo{
				familyName: "CROUCH",
				givenName:  "BARTEMIUS",
				middleName: "",
				namePrefix: "",
				nameSuffix: "",
			},
		},
		// No first
		{
			Raw: "CROUCH^^'BARTY'^MR^JR",
			Expected: GroupInfo{
				familyName: "CROUCH",
				givenName:  "",
				middleName: "'BARTY'",
				namePrefix: "MR",
				nameSuffix: "JR",
			},
		},
		// Empty
		{
			Raw: "^^^^",
			Expected: GroupInfo{
				familyName: "",
				givenName:  "",
				middleName: "",
				namePrefix: "",
				nameSuffix: "",
			},
			IsEmpty: true,
		},
		// No suffix or trailing
		{
			Raw: "CROUCH^BARTEMIUS^'BARTY'^MR",
			Expected: GroupInfo{
				familyName: "CROUCH",
				givenName:  "BARTEMIUS",
				middleName: "'BARTY'",
				namePrefix: "MR",
				nameSuffix: "",
			},
			RemoveTrailingSeps: true,
		},
		// No prefix or trailing
		{
			Raw: "CROUCH^BARTEMIUS^'BARTY'",
			Expected: GroupInfo{
				familyName: "CROUCH",
				givenName:  "BARTEMIUS",
				middleName: "'BARTY'",
				namePrefix: "",
				nameSuffix: "",
			},
			RemoveTrailingSeps: true,
		},
		// No middle or trailing
		{
			Raw: "CROUCH^BARTEMIUS",
			Expected: GroupInfo{
				familyName: "CROUCH",
				givenName:  "BARTEMIUS",
				middleName: "",
				namePrefix: "",
				nameSuffix: "",
			},
			RemoveTrailingSeps: true,
		},
		// No given or trailing
		{
			Raw: "CROUCH",
			Expected: GroupInfo{
				familyName: "CROUCH",
				givenName:  "",
				middleName: "",
				namePrefix: "",
				nameSuffix: "",
			},
			RemoveTrailingSeps: true,
		},
		// No family or trailing
		{
			Raw: "",
			Expected: GroupInfo{
				familyName: "",
				givenName:  "",
				middleName: "",
				namePrefix: "",
				nameSuffix: "",
			},
			RemoveTrailingSeps: true,
			IsEmpty:            true,
		},
	}

	var thisCase TestCase

	runNewTest := func(t *testing.T) {
		newGroup := NewGroupInfo(
			thisCase.Expected.familyName,
			thisCase.Expected.givenName,
			thisCase.Expected.middleName,
			thisCase.Expected.namePrefix,
			thisCase.Expected.nameSuffix,
			thisCase.RemoveTrailingSeps,
		)
		assert.Equal(
			t,
			thisCase.Raw,
			newGroup.String(),
			"convert to string",
		)
	}

	runParseTest := func(t *testing.T) {
		assert := assert.New(t)

		parsed, err := groupFromValueString(thisCase.Raw, "alphabetic")
		if !assert.NoError(err, "parse string") {
			t.FailNow()
		}

		checkGroupInfo(t, thisCase.Expected, parsed, "")

		assert.Equal(thisCase.Raw, parsed.String(), "convert to string")
	}

	runIsEmptyTest := func(t *testing.T) {
		assert := assert.New(t)

		parsed, err := groupFromValueString(thisCase.Raw, "alphabetic")
		if !assert.NoError(err, "parse string") {
			t.FailNow()
		}

		assert.Equal(thisCase.IsEmpty, parsed.IsEmpty(), "IsEmpty()")
	}

	for _, thisCase = range testCases {
		thisCase.Expected.raw = thisCase.Raw
		t.Run(thisCase.Raw+"_New", runNewTest)
		t.Run(thisCase.Raw+"_Parse", runParseTest)
		t.Run(thisCase.Raw+"_IsEmpty", runIsEmptyTest)
	}
}

func TestNewPersonNameFromDicom_Err(t *testing.T) {
	badName := "Malfoy^Draco^^^^"
	_, err := groupFromValueString(badName, "alphabetic")

	// Check that we get a ErrParsePersonName
	assert.True(t, errors.Is(err, ErrParsePersonName))
	assert.EqualError(
		t,
		err,
		"string contains to many segments for PN value: PN group alphabetic"+
			" contains 6 segments. No more than 5 segments with"+
			" '[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed",
	)
}
