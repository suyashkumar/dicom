package personname

import (
	"errors"
	"fmt"
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

	if expected.TrailingNullLevel != received.TrailingNullLevel {
		t.Errorf(
			"TrailingNullLevel: expected %v, got %v",
			expected.TrailingNullLevel,
			received.TrailingNullLevel,
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
				FamilyName:        "CROUCH",
				GivenName:         "BARTEMIUS",
				MiddleName:        "'BARTY'",
				NamePrefix:        "MR",
				NameSuffix:        "",
				TrailingNullLevel: GroupNullAll,
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
				FamilyName:        "CROUCH",
				GivenName:         "BARTEMIUS",
				MiddleName:        "",
				NamePrefix:        "",
				NameSuffix:        "",
				TrailingNullLevel: GroupNullAll,
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
		// No family name.
		{
			Raw: "^BARTEMIUS^'BARTY'^MR^JR",
			Expected: GroupInfo{
				FamilyName: "",
				GivenName:  "BARTEMIUS",
				MiddleName: "'BARTY'",
				NamePrefix: "MR",
				NameSuffix: "JR",
			},
		},
		// No family name, no suffix
		{
			Raw: "^BARTEMIUS^'BARTY'^MR",
			Expected: GroupInfo{
				FamilyName: "",
				GivenName:  "BARTEMIUS",
				MiddleName: "'BARTY'",
				NamePrefix: "MR",
				NameSuffix: "",
			},
		},
		// Empty
		{
			Raw: "^^^^",
			Expected: GroupInfo{
				FamilyName:        "",
				GivenName:         "",
				MiddleName:        "",
				NamePrefix:        "",
				NameSuffix:        "",
				TrailingNullLevel: GroupNullAll,
			},
			IsEmpty: true,
		},
		// Empty, missing Prefix separator
		{
			Raw: "^^^",
			Expected: GroupInfo{
				FamilyName:        "",
				GivenName:         "",
				MiddleName:        "",
				NamePrefix:        "",
				NameSuffix:        "",
				TrailingNullLevel: GroupNullPrefix,
			},
			IsEmpty: true,
		},
		// Empty, missing Middle separator
		{
			Raw: "^^",
			Expected: GroupInfo{
				FamilyName:        "",
				GivenName:         "",
				MiddleName:        "",
				NamePrefix:        "",
				NameSuffix:        "",
				TrailingNullLevel: GroupNullMiddle,
			},
			IsEmpty: true,
		},
		// Empty, missing Given separator
		{
			Raw: "^",
			Expected: GroupInfo{
				FamilyName:        "",
				GivenName:         "",
				MiddleName:        "",
				NamePrefix:        "",
				NameSuffix:        "",
				TrailingNullLevel: GroupNullGiven,
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
		},
		// Middle Only with missing prefix separator
		{
			Raw: "^^'BARTY'^",
			Expected: GroupInfo{
				FamilyName:        "",
				GivenName:         "",
				MiddleName:        "'BARTY'",
				NamePrefix:        "",
				NameSuffix:        "",
				TrailingNullLevel: GroupNullPrefix,
			},
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
			IsEmpty: true,
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
				tc.Expected.TrailingNullLevel,
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

			fmt.Println("EXPECTED NULLS:", tc.Expected.TrailingNullLevel)
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

// TestGroupInfo_DCM_interiorNullsExceedTrailingLevel tests that if a group has interior
// null values that exceed the
func TestGroupInfo_DCM_interiorNullsExceedTrailingLevel(t *testing.T) {
	groupInfo := GroupInfo{
		FamilyName:        "CROUCH",
		GivenName:         "",
		MiddleName:        "",
		NamePrefix:        "",
		NameSuffix:        "JR",
		TrailingNullLevel: GroupNullMiddle,
	}

	if groupInfo.DCM() != "CROUCH^^^^JR" {
		t.Errorf("dcm returned uneexpected value: %v", groupInfo.DCM())
	}
}

func TestGroupInfo_DCM_panic(t *testing.T) {
	groupInfo := GroupInfo{
		TrailingNullLevel: 5,
	}

	var recovered interface{}

	func() {
		defer func() {
			recovered = recover()
		}()
		groupInfo.DCM()
	}()

	if recovered == nil {
		t.Error("did not recover panic")
	}

	err, ok := recovered.(error)
	if !ok {
		t.Error("recovered panic was not error")
	}

	if err.Error() != "TrailingNullLevel exceeded maximum: cannot be greater than 4, got 5" {
		t.Error("err had unexpected text:", err.Error())
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
