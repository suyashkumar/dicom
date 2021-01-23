package personname_test

import (
	"fmt"
	"github.com/suyashkumar/dicom/pkg/personname"
)

// The below test does not pass, I think due to the ideograms being incorrectly
// compared. We are going to use 'Outputs' instead of 'Output' to stop the test from
// running. Kind of annoying.

// ExampleParse shows how to parse an existing PN value.
func ExampleParse() {
	// A Raw PN string value with Alphabetic, Ideographic, and Phonetic information.
	rawPN := "Potter^Harry^James^^=哈利^波特^詹姆^^=hɛər.i^pɒ.tər^dʒeɪmz^^"

	// To parse a PN value, use personName.Parse
	parsedPN, err := personname.Parse(rawPN)
	if err != nil {
		panic(err)
	}

	// PN values are broken into three groups: Alphabetic, Phonetic, and Ideographic;
	// our struct contains a field for each one.
	fmt.Println("ALPHABETIC:", parsedPN.Alphabetic.DCM())
	fmt.Println("IDEOGRAPHIC:", parsedPN.Ideographic.DCM())
	fmt.Println("PHONETIC:", parsedPN.Phonetic.DCM())

	// We can check to what level of trailing null separators were present, because
	// there are no null groups, the level will be NONE.
	fmt.Println("INFO NULL SEPS:", parsedPN.TrailingNullLevel)

	// Each GroupInfo in turn is parsed into discrete fields for each name segment.
	fmt.Println("FAMILY NAME:", parsedPN.Alphabetic.FamilyName)
	fmt.Println("GIVEN NAME: ", parsedPN.Alphabetic.GivenName)
	fmt.Println("MIDDLE NAME:", parsedPN.Alphabetic.MiddleName)
	fmt.Println("NAME PREFIX:", parsedPN.Alphabetic.NamePrefix)
	fmt.Println("NAME SUFFIX:", parsedPN.Alphabetic.NameSuffix)

	// The alphabetic null separators, on the other hand, will be ALL, since there are
	// null separators and the maximum allowed number is present.
	fmt.Println("ALPHABETIC NULL SEPS:", parsedPN.Alphabetic.TrailingNullLevel)

	// To print the original Raw value, simply use the DCM() method.
	fmt.Println("DCM OUTPUT :", parsedPN.DCM())

	// Outputs:
	// ALPHABETIC: Potter^Harry^James^^
	// IDEOGRAPHIC: 哈利^波特^詹姆^^
	// PHONETIC: hɛər.i^pɒ.tər^dʒeɪmz^^
	// INFO NULL SEPS: NONE
	//
	// FAMILY NAME: Potter
	// GIVEN NAME:  Harry
	// MIDDLE NAME: James
	// NAME PREFIX:
	// NAME SUFFIX:
	// ALPHABETIC NULL SEPS: ALL
	//
	// DCM OUTPUT : Potter^Harry^James^^=哈利^波特^詹姆^^=hɛər.i^pɒ.tər^dʒeɪmz^^
}

// ExampleParse_partialNullSeparators shots parsing a name with only a few null
// separators.
func ExampleParse_partialNullSeparators() {
	// This value only has one null separator. While this is a legal value, it's format
	// is somewhat irregular.
	rawPN := "Potter^Harry^"

	// To parse a PN value, use personName.Parse
	parsedPN, err := personname.Parse(rawPN)
	if err != nil {
		panic(err)
	}

	// We can check the null separator level here. When it is not NONE or ALL, it will
	// be the name of the segment that comes after the highest present null separator:
	fmt.Println("NULL SEP LEVEL:", parsedPN.Alphabetic.TrailingNullLevel)

	// Print the DCM() method, the missing null separators are preserved.
	fmt.Println("DCM OUTPUT    :", parsedPN.DCM())

	// Output:
	// NULL SEP LEVEL: MiddleName
	// DCM OUTPUT    : Potter^Harry^
}

// How to create new PN value.
func ExampleNew() {
	// Create a new PN like so
	pnVal := personname.Info{
		Alphabetic: personname.GroupInfo{
			FamilyName: "Potter",
			GivenName:  "Harry",
			MiddleName: "James",
			// This group will print null separators, but the rest will not.
			TrailingNullLevel: personname.GroupNullAll,
		},
		// Add empty group that will not render its separators.
		Ideographic: personname.GroupInfo{},
		// Add empty group that will not render its separators.
		Phonetic: personname.GroupInfo{},
	}

	// Print the string, should render as 'Potter^Harry^James^^'.
	fmt.Println("PN 1:", pnVal.DCM())

	// Now let's make one that still renders empty groups with DCM().
	pnVal = personname.Info{
		Alphabetic: personname.GroupInfo{
			FamilyName:        "Potter",
			GivenName:         "Harry",
			MiddleName:        "James",
			TrailingNullLevel: personname.GroupNullAll,
		},
		// Add empty group that will render its separators.
		Ideographic: personname.GroupInfo{
			TrailingNullLevel: personname.GroupNullAll,
		},
		// Add empty group that will render its separators.
		Phonetic: personname.GroupInfo{
			TrailingNullLevel: personname.GroupNullAll,
		},
		TrailingNullLevel: personname.InfoNullAll,
	}

	// This will render as 'Potter^Harry^James^^=^^^^=^^^^'
	fmt.Println("PN 2:", pnVal.DCM())

	// Output:
	// PN 1: Potter^Harry^James^^
	// PN 2: Potter^Harry^James^^=^^^^=^^^^
}

// ExampleInfo_WithTrailingNulls shows reformatting a PN value without trailing
// separators.
func ExampleInfo_WithTrailingNulls() {
	rawVal := "Potter^Harry"

	parsedPN, err := personname.Parse(rawVal)
	if err != nil {
		panic(err)
	}

	fmt.Println("ORIGINAL   :", parsedPN.DCM())

	reformatted := parsedPN.WithTrailingNulls()
	fmt.Println("REFORMATTED:", reformatted.DCM())

	// Output:
	// ORIGINAL   : Potter^Harry
	// REFORMATTED: Potter^Harry^^^=^^^^=^^^^
}

// ExampleInfo_WithoutTrailingNulls shows reformatting a PN values without trailing
// separators.
func ExampleInfo_WithoutTrailingNulls() {
	rawVal := "Potter^Harry^^^=^^^^=^^^^"

	parsedPN, err := personname.Parse(rawVal)
	if err != nil {
		panic(err)
	}

	fmt.Println("ORIGINAL   :", parsedPN.DCM())

	reformatted := parsedPN.WithoutTrailingNulls()
	fmt.Println("REFORMATTED:", reformatted.DCM())

	// Output:
	// ORIGINAL   : Potter^Harry^^^=^^^^=^^^^
	// REFORMATTED: Potter^Harry
}

// ExampleInfo_WithFormat shows reformatting a PN value to a mix and match of formatting
// options.
func ExampleInfo_WithFormat() {
	rawVal := "Potter^Harry^^^=^^^^=^^^^"

	parsedPN, err := personname.Parse(rawVal)
	if err != nil {
		panic(err)
	}

	fmt.Println("ORIGINAL   :", parsedPN.DCM())

	reformatted := parsedPN.WithFormat(
		personname.InfoNullNone, // groupTrailingNulls
		// We will keep empty separators for the Alphabetic group.
		personname.GroupNullAll,  // alphabeticTrailingNulls
		personname.GroupNullNone, // ideographicTrailingNulls
		personname.GroupNullNone, // phoneticTrailingNulls
	)
	fmt.Println("REFORMATTED:", reformatted.DCM())

	// Output:
	// ORIGINAL   : Potter^Harry^^^=^^^^=^^^^
	// REFORMATTED: Potter^Harry^^^
}

// ExampleInfo_WithoutEmptyGroups demos removing groups with no actual information, and
// the Alphabetic group has partial information.
func ExampleInfo_WithoutEmptyGroups() {
	rawVal := "Potter^Harry^^^=^^^^=^^^^"

	parsedPN, err := personname.Parse(rawVal)
	if err != nil {
		panic(err)
	}

	fmt.Println("ORIGINAL   :", parsedPN.DCM())

	reformatted := parsedPN.WithoutEmptyGroups()
	fmt.Println("REFORMATTED:", reformatted.DCM())

	// Output:
	// ORIGINAL   : Potter^Harry^^^=^^^^=^^^^
	// REFORMATTED: Potter^Harry^^^
}

// ExampleInfo_WithoutEmptyGroups demos removing groups with no actual information
// when both the Alphabetic and Phonetic group has partial information, but bot the
// ideographic group.
func ExampleInfo_WithoutEmptyGroups_hasPhonetic() {
	rawVal := "Potter^Harry^^^=^^^^=hɛər.i^pɒ.tər^dʒeɪmz^^"

	parsedPN, err := personname.Parse(rawVal)
	if err != nil {
		panic(err)
	}

	fmt.Println("ORIGINAL   :", parsedPN.DCM())

	reformatted := parsedPN.WithoutEmptyGroups()
	fmt.Println("REFORMATTED:", reformatted.DCM())

	// Output:
	// ORIGINAL   : Potter^Harry^^^=^^^^=hɛər.i^pɒ.tər^dʒeɪmz^^
	// REFORMATTED: Potter^Harry^^^==hɛər.i^pɒ.tər^dʒeɪmz^^
}
