package personname

import (
	"errors"
	"testing"
)

func TestInfo(t *testing.T) {
	testCases := []struct {
		// The Raw string to parse from
		Raw string
		// The parsed information we expect.
		Expected Info
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
					FamilyName:        "aFamily",
					GivenName:         "aGiven",
					MiddleName:        "aMiddle",
					NamePrefix:        "aPrefix",
					NameSuffix:        "aSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Ideographic: GroupInfo{
					FamilyName:        "iFamily",
					GivenName:         "iGiven",
					MiddleName:        "iMiddle",
					NamePrefix:        "iPrefix",
					NameSuffix:        "iSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Phonetic: GroupInfo{
					FamilyName:        "pFamily",
					GivenName:         "pGiven",
					MiddleName:        "pMiddle",
					NamePrefix:        "pPrefix",
					NameSuffix:        "pSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				TrailingNullLevel: InfoNullLevelNone,
			},
			IsEmpty: false,
		},
		// No Phonetic
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix=" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:        "aFamily",
					GivenName:         "aGiven",
					MiddleName:        "aMiddle",
					NamePrefix:        "aPrefix",
					NameSuffix:        "aSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Ideographic: GroupInfo{
					FamilyName:        "iFamily",
					GivenName:         "iGiven",
					MiddleName:        "iMiddle",
					NamePrefix:        "iPrefix",
					NameSuffix:        "iSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Phonetic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				TrailingNullLevel: InfoNullLevelAll,
			},
			IsEmpty: false,
		},
		// No Phonetic, no seps
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:        "aFamily",
					GivenName:         "aGiven",
					MiddleName:        "aMiddle",
					NamePrefix:        "aPrefix",
					NameSuffix:        "aSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Ideographic: GroupInfo{
					FamilyName:        "iFamily",
					GivenName:         "iGiven",
					MiddleName:        "iMiddle",
					NamePrefix:        "iPrefix",
					NameSuffix:        "iSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Phonetic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				TrailingNullLevel: InfoNullLevelNone,
			},
			IsEmpty: false,
		},
		// No Ideographic
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=" +
				"=" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:        "aFamily",
					GivenName:         "aGiven",
					MiddleName:        "aMiddle",
					NamePrefix:        "aPrefix",
					NameSuffix:        "aSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Ideographic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Phonetic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				TrailingNullLevel: InfoNullLevelAll,
			},
			IsEmpty: false,
		},
		// No Ideographic, no seps
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:        "aFamily",
					GivenName:         "aGiven",
					MiddleName:        "aMiddle",
					NamePrefix:        "aPrefix",
					NameSuffix:        "aSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Ideographic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Phonetic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				TrailingNullLevel: InfoNullLevelNone,
			},
			IsEmpty: false,
		},
		// No Alphabetic, with seps
		{
			Raw: "=" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix=" +
				"pFamily^pGiven^pMiddle^pPrefix^pSuffix",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Ideographic: GroupInfo{
					FamilyName:        "iFamily",
					GivenName:         "iGiven",
					MiddleName:        "iMiddle",
					NamePrefix:        "iPrefix",
					NameSuffix:        "iSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Phonetic: GroupInfo{
					FamilyName:        "pFamily",
					GivenName:         "pGiven",
					MiddleName:        "pMiddle",
					NamePrefix:        "pPrefix",
					NameSuffix:        "pSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
			},
			IsEmpty: false,
		},
		// Alphabetic, with only 1 sep
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix=",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:        "aFamily",
					GivenName:         "aGiven",
					MiddleName:        "aMiddle",
					NamePrefix:        "aPrefix",
					NameSuffix:        "aSuffix",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Ideographic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Phonetic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				TrailingNullLevel: InfoNullLevelIdeographic,
			},
			IsEmpty: false,
		},
		// Empty with seps
		{
			Raw: "=" +
				"=" +
				"",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Ideographic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Phonetic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				TrailingNullLevel: InfoNullLevelAll,
			},
			IsEmpty: true,
		},
		// Empty no seps
		{
			Raw: "",
			Expected: Info{
				Alphabetic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Ideographic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				Phonetic: GroupInfo{
					FamilyName:        "",
					GivenName:         "",
					MiddleName:        "",
					NamePrefix:        "",
					NameSuffix:        "",
					TrailingNullLevel: GroupNullLevelNone,
				},
				TrailingNullLevel: InfoNullLevelNone,
			},
			IsEmpty: true,
		},
	}

	for _, tc := range testCases {

		// Test creating a new Info value and getting it's DCM() value.
		t.Run(tc.Raw+"_String", func(t *testing.T) {
			newInfo := Info{
				Alphabetic:        tc.Expected.Alphabetic,
				Ideographic:       tc.Expected.Ideographic,
				Phonetic:          tc.Expected.Phonetic,
				TrailingNullLevel: tc.Expected.TrailingNullLevel,
			}

			dcm, err := newInfo.DCM()
			if err != nil {
				t.Fatal("DCM() returned error:", err)
			}

			if tc.Raw != dcm {
				t.Errorf(
					"formatted string: expected '%v', got '%v'",
					tc.Raw,
					dcm,
				)
			}
		})

		// Test parsing a full PN value.
		t.Run(tc.Raw+"_parse", func(t *testing.T) {

			parsed, err := Parse(tc.Raw)
			if err != nil {
				t.Fatal("error parsing value:", err)
			}

			checkGroupInfo(
				t,
				tc.Expected.Alphabetic,
				parsed.Alphabetic,
				tc.Expected.Alphabetic.MustDCM(),
				"Alphabetic",
			)
			checkGroupInfo(
				t,
				tc.Expected.Ideographic,
				parsed.Ideographic,
				tc.Expected.Ideographic.MustDCM(),
				"Ideographic",
			)
			checkGroupInfo(
				t,
				tc.Expected.Phonetic,
				parsed.Phonetic,
				tc.Expected.Phonetic.MustDCM(),
				"Phonetic",
			)
		})

		// Test the .IsEmpty() method.
		t.Run(tc.Raw+"_isEmpty", func(t *testing.T) {
			newInfo, err := Parse(tc.Raw)
			if err != nil {
				t.Error("error parsing value:", err)
				t.FailNow()
			}

			if tc.IsEmpty != newInfo.IsEmpty() {
				t.Errorf(
					".IsEmpty(): got '%v', expected '%v'",
					tc.IsEmpty,
					newInfo.IsEmpty(),
				)
			}
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
			ErrString: "error parsing PN value: no more than 3 groups with " +
				"'[Alphabetic]=[Ideographic]=[Phonetic]' format are allowed: value " +
				"contains 4 groups. see 'PN' entry in official dicom spec: " +
				"https://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_6.2",
		},
		{
			Raw: "^^^^^",
			ErrString: "error parsing PN value: no more than 5 segments with " +
				"'[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed: " +
				"value group Alphabetic contains 6 segments. see 'PN' entry in " +
				"official dicom spec: " +
				"https://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_6.2",
		},
		{
			Raw: "=^^^^^",
			ErrString: "error parsing PN value: no more than 5 segments with " +
				"'[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed: " +
				"value group Ideographic contains 6 segments. see 'PN' entry in " +
				"official dicom spec: " +
				"https://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_6.2",
		},
		{
			Raw: "==^^^^^",
			ErrString: "error parsing PN value: no more than 5 segments with " +
				"'[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed: " +
				"value group Phonetic contains 6 segments. see 'PN' entry in " +
				"official dicom spec: " +
				"https://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_6.2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Raw, func(t *testing.T) {
			_, err := Parse(tc.Raw)

			// Test errors.Is() with ErrParsePersonName
			if !errors.Is(err, ErrParsePersonName) {
				t.Errorf("err is not ErrParsePersonName")
			}

			if err.Error() != tc.ErrString {
				t.Error("error string unexpected:", err.Error())
			}
		})
	}
}

func TestInfo_WithNullSeparators(t *testing.T) {
	pnVal := "Potter^Harry"

	parsed, err := Parse(pnVal)
	if err != nil {
		t.Fatalf("error parsing pnVal: %v", err)
	}

	altered := parsed.WithTrailingNulls()

	expected := "Potter^Harry^^^=^^^^=^^^^"

	dcm, err := altered.DCM()
	if err != nil {
		t.Fatal("DCM() returned error:", err)
	}

	if dcm != expected {
		t.Errorf("expected '%v', got '%v'", expected, dcm)
	}
}

func TestInfo_WithoutNullSeparators(t *testing.T) {
	pnVal := "Potter^Harry^^^=^^^^=^^^^"

	parsed, err := Parse(pnVal)
	if err != nil {
		t.Fatalf("error parsing pnVal: %v", err)
	}

	altered := parsed.WithoutTrailingNulls()

	expected := "Potter^Harry"

	dcm, err := altered.DCM()
	if err != nil {
		t.Fatal("DCM() returned error", err)
	}

	if dcm != expected {
		t.Errorf("expected '%v', got '%v'", expected, dcm)
	}
}

func TestInfo_DCM_panic(t *testing.T) {
	groupInfo := Info{
		TrailingNullLevel: 3,
	}

	var recovered any

	func() {
		defer func() {
			recovered = recover()
		}()
		groupInfo.MustDCM()
	}()

	if recovered == nil {
		t.Fatal("did not recover panic")
	}

	err, ok := recovered.(error)
	if !ok {
		t.Fatal("recovered panic was not error")
	}

	if err.Error() != "TrailingNullLevel exceeded maximum: cannot be greater than 2, got 3" {
		t.Error("err had unexpected text:", err.Error())
	}
}
