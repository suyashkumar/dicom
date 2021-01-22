package personname_test

import (
	"fmt"
	"github.com/suyashkumar/dicom/pkg/personname"
)

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

	// Each GroupInfo in turn is parsed into discrete fields for each name segment.
	fmt.Println("FAMILY NAME:", parsedPN.Alphabetic.FamilyName)
	fmt.Println("GIVEN NAME: ", parsedPN.Alphabetic.GivenName)
	fmt.Println("MIDDLE NAME:", parsedPN.Alphabetic.MiddleName)
	fmt.Println("NAME PREFIX:", parsedPN.Alphabetic.NamePrefix)
	fmt.Println("NAME SUFFIX:", parsedPN.Alphabetic.NameSuffix)

	// To print the original Raw value, simply use the DCM() method.
	fmt.Println("DCM OUTPUT :", parsedPN.DCM())

	// Output (gotest comparison here is busted, maybe due to ideographs?):
	// ALPHABETIC: Potter^Harry^James^^
	// IDEOGRAPHIC: 哈利^波特^詹姆^^
	// PHONETIC: hɛər.i^pɒ.tər^dʒeɪmz^^
	// FAMILY NAME: Potter
	// GIVEN NAME:  Harry
	// MIDDLE NAME: James
	// NAME PREFIX:
	// NAME SUFFIX:
	// DCM OUTPUT : Potter^Harry^James^^=哈利^波特^詹姆^^=hɛər.i^pɒ.tər^dʒeɪmz^^
}

func ExampleParse_partialNullSeparators() {
	// This value only has one null separator. While this is a legal value, it's format
	// is somewhat irregular.
	rawPN := "Potter^Harry^"

	// To parse a PN value, use personName.Parse
	parsedPN, err := personname.Parse(rawPN)
	if err != nil {
		panic(err)
	}

	// Print the DCM() method, the missing null separators have been added.
	fmt.Println("DCM OUTPUT:", parsedPN.DCM())

	// Output:
	// DCM OUTPUT: Potter^Harry^^^
}

// How to create new PN value.
func ExampleNew() {
	// Create a new PN like so
	pnVal := personname.Info{
		Alphabetic: personname.GroupInfo{
			FamilyName:        "Potter",
			GivenName:         "Harry",
			MiddleName:        "James",
			HasNullSeparators: true,
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
			HasNullSeparators: true,
		},
		// Add empty group that will render its separators.
		Ideographic: personname.GroupInfo{
			HasNullSeparators: true,
		},
		// Add empty group that will render its separators.
		Phonetic: personname.GroupInfo{
			HasNullSeparators: true,
		},
		HasNullSeparators: true,
	}

	// This will render as 'Potter^Harry^James^^=^^^^=^^^^'
	fmt.Println("PN 2:", pnVal.DCM())

	// Output:
	// PN 1: Potter^Harry^James^^
	// PN 2: Potter^Harry^James^^=^^^^=^^^^
}

// ExampleInfo_WithNullSeparators shows reformatting a PN value without trailing
// separators.
func ExampleInfo_WithNullSeparators() {
	rawVal := "Potter^Harry"

	parsedPN, err := personname.Parse(rawVal)
	if err != nil {
		panic(err)
	}

	fmt.Println("ORIGINAL   :", parsedPN.DCM())

	reformatted := parsedPN.WithNullSeparators()
	fmt.Println("REFORMATTED:", reformatted.DCM())

	// Output:
	// ORIGINAL   : Potter^Harry
	// REFORMATTED: Potter^Harry^^^=^^^^=^^^^
}

// ExampleInfo_WithoutNullSeparators shows reformatting a PN values without trailing
// separators.
func ExampleInfo_WithoutNullSeparators() {
	rawVal := "Potter^Harry^^^=^^^^=^^^^"

	parsedPN, err := personname.Parse(rawVal)
	if err != nil {
		panic(err)
	}

	fmt.Println("ORIGINAL   :", parsedPN.DCM())

	reformatted := parsedPN.WithoutNullSeparators()
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
		false, // useGroupNullSeparators
		// We will keep empty separators for the Alphabetic group.
		true,  // useAlphabeticNullSeparators
		false, // useIdeographicNullSeparators
		false, // usePhoneticNullSeparators
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
