package personName

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInfo(t *testing.T) {
	type TestCase struct {
		// The raw string to parse from
		Raw string
		// The parsed information we expect.
		Expected Info
		// Whether RemoveTrailingEmpty should be set to true when creating a new
		// GroupInfo to match Raw.
		RemoveTrailingEmpty bool
		// Whether IsEmpty should return true after parsing Raw.
		IsEmpty bool
	}

	var thisCase TestCase

	runNewTest := func(t *testing.T) {
		newInfo := New(
			thisCase.Expected.alphabetic,
			thisCase.Expected.ideographic,
			thisCase.Expected.phonetic,
			thisCase.RemoveTrailingEmpty,
		)

		assert.Equal(t, thisCase.Raw, newInfo.String())
	}

	runParseTest := func(t *testing.T) {
		assert := assert.New(t)

		parsed, err := Parse(thisCase.Raw)
		if !assert.NoError(err, "parse raw") {
			t.FailNow()
		}

		checkGroupInfo(
			t,
			thisCase.Expected.alphabetic,
			parsed.alphabetic,
			"alphabetic",
		)
		checkGroupInfo(
			t,
			thisCase.Expected.ideographic,
			parsed.ideographic,
			"ideographic",
		)
		checkGroupInfo(
			t,
			thisCase.Expected.phonetic,
			parsed.phonetic,
			"phonetic",
		)
	}

	runIsEmptyTest := func(t *testing.T) {
		assert := assert.New(t)

		newInfo, err := Parse(thisCase.Raw)
		if !assert.NoError(err, "parse raw") {
			t.FailNow()
		}

		assert.Equal(thisCase.IsEmpty, newInfo.IsEmpty())
	}

	testCases := []TestCase{
		// All groups
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix===" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix===" +
				"pFamily^pGiven^pMiddle^pPrefix^pSuffix",
			Expected: Info{
				alphabetic: GroupInfo{
					raw:        "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
					familyName: "aFamily",
					givenName:  "aGiven",
					middleName: "aMiddle",
					namePrefix: "aPrefix",
					nameSuffix: "aSuffix",
				},
				ideographic: GroupInfo{
					raw:        "iFamily^iGiven^iMiddle^iPrefix^iSuffix",
					familyName: "iFamily",
					givenName:  "iGiven",
					middleName: "iMiddle",
					namePrefix: "iPrefix",
					nameSuffix: "iSuffix",
				},
				phonetic: GroupInfo{
					raw:        "pFamily^pGiven^pMiddle^pPrefix^pSuffix",
					familyName: "pFamily",
					givenName:  "pGiven",
					middleName: "pMiddle",
					namePrefix: "pPrefix",
					nameSuffix: "pSuffix",
				},
			},
			RemoveTrailingEmpty: false,
			IsEmpty:             false,
		},
		// No phonetic
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix===" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix===" +
				"",
			Expected: Info{
				alphabetic: GroupInfo{
					raw:        "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
					familyName: "aFamily",
					givenName:  "aGiven",
					middleName: "aMiddle",
					namePrefix: "aPrefix",
					nameSuffix: "aSuffix",
				},
				ideographic: GroupInfo{
					raw:        "iFamily^iGiven^iMiddle^iPrefix^iSuffix",
					familyName: "iFamily",
					givenName:  "iGiven",
					middleName: "iMiddle",
					namePrefix: "iPrefix",
					nameSuffix: "iSuffix",
				},
				phonetic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
			},
			RemoveTrailingEmpty: false,
			IsEmpty:             false,
		},
		// No phonetic, no seps
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix===" +
				"iFamily^iGiven^iMiddle^iPrefix^iSuffix" +
				"",
			Expected: Info{
				alphabetic: GroupInfo{
					raw:        "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
					familyName: "aFamily",
					givenName:  "aGiven",
					middleName: "aMiddle",
					namePrefix: "aPrefix",
					nameSuffix: "aSuffix",
				},
				ideographic: GroupInfo{
					raw:        "iFamily^iGiven^iMiddle^iPrefix^iSuffix",
					familyName: "iFamily",
					givenName:  "iGiven",
					middleName: "iMiddle",
					namePrefix: "iPrefix",
					nameSuffix: "iSuffix",
				},
				phonetic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
			},
			RemoveTrailingEmpty: true,
			IsEmpty:             false,
		},
		// No ideographic
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix===" +
				"===" +
				"",
			Expected: Info{
				alphabetic: GroupInfo{
					raw:        "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
					familyName: "aFamily",
					givenName:  "aGiven",
					middleName: "aMiddle",
					namePrefix: "aPrefix",
					nameSuffix: "aSuffix",
				},
				ideographic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
				phonetic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
			},
			RemoveTrailingEmpty: false,
			IsEmpty:             false,
		},
		// No ideographic, no seps
		{
			Raw: "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
			Expected: Info{
				alphabetic: GroupInfo{
					raw:        "aFamily^aGiven^aMiddle^aPrefix^aSuffix",
					familyName: "aFamily",
					givenName:  "aGiven",
					middleName: "aMiddle",
					namePrefix: "aPrefix",
					nameSuffix: "aSuffix",
				},
				ideographic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
				phonetic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
			},
			RemoveTrailingEmpty: true,
			IsEmpty:             false,
		},
		// Empty with seps
		{
			Raw: "===" +
				"===" +
				"",
			Expected: Info{
				alphabetic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
				ideographic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
				phonetic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
			},
			RemoveTrailingEmpty: false,
			IsEmpty:             true,
		},
		// Empty no seps
		{
			Raw: "",
			Expected: Info{
				alphabetic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
				ideographic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
				phonetic: GroupInfo{
					raw:        "",
					familyName: "",
					givenName:  "",
					middleName: "",
					namePrefix: "",
					nameSuffix: "",
				},
			},
			RemoveTrailingEmpty: true,
			IsEmpty:             true,
		},
	}

	for _, thisCase = range testCases {
		thisCase.Expected.raw = thisCase.Raw

		t.Run(thisCase.Raw+"_new", runNewTest)
		t.Run(thisCase.Raw+"_parse", runParseTest)
		t.Run(thisCase.Raw+"_isEmpty", runIsEmptyTest)
	}
}

func TestParse_Err(t *testing.T) {
	type TestCase struct {
		Raw       string
		ErrString string
	}

	var thisCase TestCase

	runTest := func(t *testing.T) {
		assert := assert.New(t)

		_, err := Parse(thisCase.Raw)
		assert.True(
			errors.Is(err, ErrParsePersonName),
			"error unwraps to ErrParsePersonName",
		)
		assert.EqualError(
			err,
			thisCase.ErrString,
		)
	}

	testCases := []TestCase{
		{
			Raw: "=========",
			ErrString: "string contains to many segments for PN value:" +
				" PN contains 4 groups. No more than 3 groups with" +
				" '[alphabetic]===[ideographic]===[phonetic]' format are allowed",
		},
		{
			Raw: "^^^^^",
			ErrString: "string contains to many segments for PN value: PN group" +
				" alphabetic contains 6 segments. No more than 5 segments with" +
				" '[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed",
		},
		{
			Raw: "===^^^^^",
			ErrString: "string contains to many segments for PN value: PN group" +
				" ideographic contains 6 segments. No more than 5 segments with" +
				" '[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed",
		},
		{
			Raw: "======^^^^^",
			ErrString: "string contains to many segments for PN value: PN group" +
				" phonetic contains 6 segments. No more than 5 segments with" +
				" '[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed",
		},
	}

	for _, thisCase = range testCases {
		t.Run(thisCase.Raw, runTest)
	}
}
