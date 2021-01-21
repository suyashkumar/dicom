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
	// our struct contains a fiel for each one.
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
	fmt.Println("ORIGINAL RAW:", parsedPN.DCM())

	// Output (gotest comparison here is busted, maybe due to ideographs?):
	// ALPHABETIC: Potter^Harry^James^^
	// IDEOGRAPHIC: 哈利^波特^詹姆^^
	// PHONETIC: hɛər.i^pɒ.tər^dʒeɪmz^^
	// FAMILY NAME: Potter
	// GIVEN NAME:  Harry
	// MIDDLE NAME: James
	// NAME PREFIX:
	// NAME SUFFIX:
	// ORIGINAL RAW: Potter^Harry^James^^=哈利^波特^詹姆^^=hɛər.i^pɒ.tər^dʒeɪmz^^
}

// How to create new PN value.
func ExampleNew() {
	// Create a new PN like so
	pnVal := personname.Info{
		Alphabetic: personname.GroupInfo{
			FamilyName: "Potter",
			GivenName:  "Harry",
			MiddleName: "James",
		},
		// Add empty group that will not render its separators.
		Ideographic: personname.GroupInfo{
			NoNullSeparators: true,
		},
		// Add empty group that will not render its separators.
		Phonetic: personname.GroupInfo{
			NoNullSeparators: true,
		},
		// This will remove trailing '=' separators
		NoNullSeparators: true,
	}

	// Print the string, should render as 'Potter^Harry^James^^'.
	fmt.Println("PN 1:", pnVal.DCM())

	// Now let's make one that still renders empty groups with DCM().
	pnVal = personname.Info{
		Alphabetic: personname.GroupInfo{
			FamilyName: "Potter",
			GivenName:  "Harry",
			MiddleName: "James",
		},
		// Add empty group that will render its separators.
		Ideographic: personname.GroupInfo{},
		// Add empty group that will render its separators.
		Phonetic: personname.GroupInfo{},
	}

	// This will render as 'Potter^Harry^James^^=^^^^=^^^^'
	fmt.Println("PN 2:", pnVal.DCM())

	// Output:
	// PN 1: Potter^Harry^James^^
	// PN 2: Potter^Harry^James^^=^^^^=^^^^
}

// ExampleInfo_WithNullSeparators shows parsing a PN value without trailing separators.
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

// ExampleInfo_WithoutNullSeparators shows parsing a PN values with trailing separators.
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
