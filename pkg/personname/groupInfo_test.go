package personname

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
		expected.FamilyName,
		received.FamilyName,
		"Family Name, group %v",
		group,
	)

	assert.Equal(
		expected.GivenName,
		received.GivenName,
		"Given Name, group %v",
		group,
	)

	assert.Equal(
		expected.MiddleName,
		received.MiddleName,
		"Middle Name, group %v",
		group,
	)

	assert.Equal(
		expected.NamePrefix,
		received.NamePrefix,
		"Name Prefix, group %v",
		group,
	)

	assert.Equal(
		expected.NameSuffix,
		received.NameSuffix,
		"Name Suffix, group %v",
		group,
	)

	assert.Equal(
		expected.Raw,
		received.String(),
		"Formatted String, group %v",
		group,
	)
}

func TestNewPersonNameFromDicom(t *testing.T) {
	testCases := []struct {
		// The Raw string to parse from
		Raw string
		// The parsed information we expect.
		Expected GroupInfo
		// Whether removeTrailingSeps should be set to true when creating a new
		// GroupInfo to match Raw.
		RemoveTrailingSeps bool
		// Whether IsEmpty should return true after parsing Raw.
		IsEmpty bool
	}{
		// Full Name
		{
			Raw: "CROUCH^BARTEMIUS^'BARTY'^MR^JR",
			Expected: GroupInfo{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "'BARTY'",
				NamePrefix: "MR",
				NameSuffix: "JR",
			},
		},
		// No Middle Name
		{
			Raw: "CROUCH^BARTEMIUS^^MR^JR",
			Expected: GroupInfo{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "",
				NamePrefix: "MR",
				NameSuffix: "JR",
			},
		},
		// No Suffix
		{
			Raw: "CROUCH^BARTEMIUS^'BARTY'^MR^",
			Expected: GroupInfo{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "'BARTY'",
				NamePrefix: "MR",
				NameSuffix: "",
			},
		},
		// No Prefix
		{
			Raw: "CROUCH^BARTEMIUS^'BARTY'^^JR",
			Expected: GroupInfo{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "'BARTY'",
				NamePrefix: "",
				NameSuffix: "JR",
			},
		},
		// Only First and last
		{
			Raw: "CROUCH^BARTEMIUS^^^",
			Expected: GroupInfo{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "",
				NamePrefix: "",
				NameSuffix: "",
			},
		},
		// No first
		{
			Raw: "CROUCH^^'BARTY'^MR^JR",
			Expected: GroupInfo{
				FamilyName: "CROUCH",
				GivenName:  "",
				MiddleName: "'BARTY'",
				NamePrefix: "MR",
				NameSuffix: "JR",
			},
		},
		// Empty
		{
			Raw: "^^^^",
			Expected: GroupInfo{
				FamilyName: "",
				GivenName:  "",
				MiddleName: "",
				NamePrefix: "",
				NameSuffix: "",
			},
			IsEmpty: true,
		},
		// No suffix or trailing
		{
			Raw: "CROUCH^BARTEMIUS^'BARTY'^MR",
			Expected: GroupInfo{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "'BARTY'",
				NamePrefix: "MR",
				NameSuffix: "",
			},
			RemoveTrailingSeps: true,
		},
		// No prefix or trailing
		{
			Raw: "CROUCH^BARTEMIUS^'BARTY'",
			Expected: GroupInfo{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "'BARTY'",
				NamePrefix: "",
				NameSuffix: "",
			},
			RemoveTrailingSeps: true,
		},
		// No middle or trailing
		{
			Raw: "CROUCH^BARTEMIUS",
			Expected: GroupInfo{
				FamilyName: "CROUCH",
				GivenName:  "BARTEMIUS",
				MiddleName: "",
				NamePrefix: "",
				NameSuffix: "",
			},
			RemoveTrailingSeps: true,
		},
		// No given or trailing
		{
			Raw: "CROUCH",
			Expected: GroupInfo{
				FamilyName: "CROUCH",
				GivenName:  "",
				MiddleName: "",
				NamePrefix: "",
				NameSuffix: "",
			},
			RemoveTrailingSeps: true,
		},
		// No family or trailing
		{
			Raw: "",
			Expected: GroupInfo{
				FamilyName: "",
				GivenName:  "",
				MiddleName: "",
				NamePrefix: "",
				NameSuffix: "",
			},
			RemoveTrailingSeps: true,
			IsEmpty:            true,
		},
	}

	for _, tc := range testCases {
		tc.Expected.Raw = tc.Raw

		// Test creating a new GroupInfo object
		t.Run(tc.Raw+"_New", func(t *testing.T) {
			newGroup := NewGroupInfo(
				tc.Expected.FamilyName,
				tc.Expected.GivenName,
				tc.Expected.MiddleName,
				tc.Expected.NamePrefix,
				tc.Expected.NameSuffix,
				tc.RemoveTrailingSeps,
			)
			assert.Equal(
				t,
				tc.Raw,
				newGroup.String(),
				"convert to string",
			)
		})

		// Test parsing a group string.
		t.Run(tc.Raw+"_Parse", func(t *testing.T) {
			assert := assert.New(t)

			parsed, err := groupFromValueString(tc.Raw, "Alphabetic")
			if !assert.NoError(err, "parse string") {
				t.FailNow()
			}

			checkGroupInfo(t, tc.Expected, parsed, "")

			assert.Equal(tc.Raw, parsed.String(), "convert to string")
		})

		// Test .IsEmpty() method.
		t.Run(tc.Raw+"_IsEmpty", func(t *testing.T) {
			assert := assert.New(t)

			parsed, err := groupFromValueString(tc.Raw, "Alphabetic")
			if !assert.NoError(err, "parse string") {
				t.FailNow()
			}

			assert.Equal(tc.IsEmpty, parsed.IsEmpty(), "IsEmpty()")
		})
	}
}

func TestNewPersonNameFromDicom_Err(t *testing.T) {
	badName := "Malfoy^Draco^^^^"
	_, err := groupFromValueString(badName, "Alphabetic")

	// Check that we get a ErrParsePersonName
	assert.True(t, errors.Is(err, ErrParsePersonName))
	assert.EqualError(
		t,
		err,
		"string contains to many segments for PN value: PN group Alphabetic"+
			" contains 6 segments. No more than 5 segments with"+
			" '[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed",
	)
}
