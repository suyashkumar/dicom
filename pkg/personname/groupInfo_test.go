package personname

import (
	"errors"
	"testing"
)

func checkGroupInfo(t *testing.T, expected, received GroupInfo, expectedString, group string) {

	if expected.FamilyName != received.FamilyName {
		t.Errorf(
			"FamilyName: expected '%v', got '%v'. Group '%v'",
			expected.FamilyName,
			received.FamilyName,
			group,
		)
	}

	if expected.GivenName != received.GivenName {
		t.Errorf(
			"GivenName: expected '%v', got '%v'. Group '%v'",
			expected.GivenName,
			received.GivenName,
			group,
		)
	}

	if expected.MiddleName != received.MiddleName {
		t.Errorf(
			"MiddleName: expected '%v', got '%v'. Group '%v'",
			expected.MiddleName,
			received.MiddleName,
			group,
		)
	}

	if expected.NamePrefix != received.NamePrefix {
		t.Errorf(
			"MiddleName: expected '%v', got '%v'. Group '%v'",
			expected.NamePrefix,
			received.NamePrefix,
			group,
		)
	}

	if expected.NameSuffix != received.NameSuffix {
		t.Errorf(
			"NameSuffix: expected '%v', got '%v'. Group '%v'",
			expected.NameSuffix,
			received.NameSuffix,
			group,
		)
	}

	if expectedString != received.DCM() {
		t.Errorf(
			"formatted .DCM(): expected '%v', got '%v'. Group '%v'",
			expectedString,
			received.DCM(),
			group,
		)
	}

}

func TestNewPersonNameFromDicom(t *testing.T) {
	testCases := []struct {
		// The Raw string to parse from
		Raw string
		// The parsed information we expect.
		Expected GroupInfo
		// Whether NoNullSeparators should be set to true when creating a new
		// GroupInfo to match Raw.
		NoNullSeps bool
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
			NoNullSeps: true,
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
			NoNullSeps: true,
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
			NoNullSeps: true,
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
			NoNullSeps: true,
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
			NoNullSeps: true,
			IsEmpty:    true,
		},
	}

	for _, tc := range testCases {
		// Test creating a new GroupInfo object
		t.Run(tc.Raw+"_New", func(t *testing.T) {
			newGroup := GroupInfo{
				tc.Expected.FamilyName,
				tc.Expected.GivenName,
				tc.Expected.MiddleName,
				tc.Expected.NamePrefix,
				tc.Expected.NameSuffix,
				tc.NoNullSeps,
			}
			if tc.Raw != newGroup.DCM() {
				t.Errorf(
					"formatted .DCM() does not match input: "+
						"expected '%v', got '%v'",
					tc.Raw,
					newGroup.DCM(),
				)
			}
		})

		// Test parsing a group string.
		t.Run(tc.Raw+"_Parse", func(t *testing.T) {

			parsed, err := groupFromValueString(tc.Raw, pnGroupAlphabetic)
			if err != nil {
				t.Fatal("error parsing value:", err)
			}

			checkGroupInfo(t, tc.Expected, parsed, tc.Raw, "")

			if tc.Raw != parsed.DCM() {
				t.Errorf(
					"formatted .DCM() does not match input: "+
						"expected '%v', got '%v'",
					tc.Raw,
					parsed.DCM(),
				)
			}
		})

		// Test .IsEmpty() method.
		t.Run(tc.Raw+"_IsEmpty", func(t *testing.T) {
			parsed, err := groupFromValueString(tc.Raw, pnGroupAlphabetic)
			if err != nil {
				t.Fatalf("error parsing value '%v': %v", tc.Raw, err)
			}

			if tc.IsEmpty != parsed.IsEmpty() {
				t.Errorf(
					".IsEmpty() returned %v, extected %v",
					parsed.IsEmpty(),
					tc.IsEmpty,
				)
			}
		})
	}
}

func TestNewPersonNameFromDicom_Err(t *testing.T) {
	badName := "Malfoy^Draco^^^^"
	_, err := groupFromValueString(badName, pnGroupAlphabetic)

	// Check that we get a ErrParsePersonName
	if !errors.Is(err, ErrParsePersonName) {
		t.Errorf("returned err is not ErrParsePersonName: %v", err)
	}

	// Check that we get a ErrParsePersonName
	if !errors.Is(err, ErrParseGroupSegmentCount) {
		t.Errorf("returned err is not ErrParseGroupSegmentCount: %v", err)
	}

	expectedString := "error parsing PN value: no more than 5 segments with " +
		"'[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed: value group " +
		"Alphabetic contains 6 segments. see 'PN' entry in official dicom spec: " +
		"http://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_6.2"

	if err.Error() != expectedString {
		t.Errorf("unexpected error text: %v", err.Error())
	}
}
