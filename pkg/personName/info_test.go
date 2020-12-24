package personName

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPersonNameFromDicom(t *testing.T) {
	type TestCase struct {
		Raw    string
		Parsed Info
	}
	
	testCases := []TestCase{
		// Full Name
		{
			Raw:    "CROUCH^BARTEMIUS^'BARTY'^MR^JR",
			Parsed: Info{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "'BARTY'",
				NamePrefix: "MR",
				NameSuffix: "JR",
			},
		},
		// No Middle Name
		{
			Raw:    "CROUCH^BARTEMIUS^^MR^JR",
			Parsed: Info{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "",
				NamePrefix: "MR",
				NameSuffix: "JR",
			},
		},
		// No Suffix
		{
			Raw:    "CROUCH^BARTEMIUS^'BARTY'^MR^",
			Parsed: Info{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "'BARTY'",
				NamePrefix: "MR",
				NameSuffix: "",
			},
		},
		// No Prefix
		{
			Raw:    "CROUCH^BARTEMIUS^'BARTY'^^JR",
			Parsed: Info{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "'BARTY'",
				NamePrefix: "",
				NameSuffix: "JR",
			},
		},
		// Only First and last
		{
			Raw:    "CROUCH^BARTEMIUS^^^",
			Parsed: Info{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "",
				NamePrefix: "",
				NameSuffix: "",
			},
		},
		// No first
		{
			Raw:    "CROUCH^^'BARTY'^MR^JR",
			Parsed: Info{
				FamilyName: "CROUCH",
				GivenName:  "",
				MiddleName: "'BARTY'",
				NamePrefix: "MR",
				NameSuffix: "JR",
			},
		},
		// Empty
		{
			Raw:    "^^^^",
			Parsed: Info{
				FamilyName: "",
				GivenName:  "",
				MiddleName: "",
				NamePrefix: "",
				NameSuffix: "",
			},
		},
	}

	var thisCase TestCase

	runParseTest := func(t *testing.T) {
		assert := assert.New(t)

		parsed, err := FromDicomValueString(thisCase.Raw)
		if !assert.NoError(err, "parse string") {
			t.FailNow()
		}

		expected := thisCase.Parsed
		assert.Equal(
			expected.FamilyName, parsed.FamilyName, "Family Name",
		)

		assert.Equal(
			expected.GivenName, parsed.GivenName, "Given Name",
		)

		assert.Equal(
			expected.MiddleName, parsed.MiddleName, "Middle Name",
		)

		assert.Equal(
			expected.NamePrefix, parsed.NamePrefix, "Name Prefix",
		)

		assert.Equal(
			expected.NameSuffix, parsed.NameSuffix, "Name Suffix",
		)
	}

	runStringTest := func(t *testing.T) {
		personStringed := thisCase.Parsed.String()
		assert.Equal(t, thisCase.Raw, personStringed, "convert to string")
	}

	for _, thisCase = range testCases {
		t.Run(thisCase.Raw + "_Parse", runParseTest)
		t.Run(thisCase.Raw + "_String", runStringTest)
	}
}

func TestNewPersonNameFromDicom_Err(t *testing.T) {
	badName := "Malfoy^Draco"
	_, err := FromDicomValueString(badName)

	// Check that we get a ErrParsePersonName
	assert.True(t, errors.Is(err, ErrParsePersonName))
}
