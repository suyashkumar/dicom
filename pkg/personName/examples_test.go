package personName_test

import (
	"fmt"
	"github.com/suyashkumar/dicom/pkg/personName"
)

// How to parse an existing PN value.
func ExampleParse() {
	// To parse a PN value, use personName.Parse
	pn, err := personName.Parse(
		"Potter^Harry^James^^===哈利^波特^詹姆^^===hɛər.i^pɒ.tər^dʒeɪmz^^",
	)
	if err != nil {
		panic(err)
	}

	// PN values are broken into three groups: Alphabetic, Phonetic, and Ideographic.
	// which can be accessed through getter methods.
	fmt.Println("ALPHABETIC", pn.Alphabetic().String())
	fmt.Println("IDEOGRAPHIC", pn.Ideographic().String())
	fmt.Println("PHONETIC", pn.Phonetic().String())

	// Each Group info in turn has Getter methods for retrieving individual data.
	fmt.Println("FAMILY NAME:", pn.Alphabetic().FamilyName())
	fmt.Println("GIVEN NAME: ", pn.Alphabetic().GivenName())
	fmt.Println("MIDDLE NAME:", pn.Alphabetic().MiddleName())
	fmt.Println("NAME PREFIX:", pn.Alphabetic().NamePrefix())
	fmt.Println("NAME SUFFIX:", pn.Alphabetic().NameSuffix())

	// To print the original raw value, simply use the string method.
	fmt.Println("ORIGINAL RAW:", pn.String())
}

// How to create new PN value.
func ExampleNew() {
	// Create a new PN like so
	pn := personName.New(
		personName.NewGroupInfo(
			"Potter",
			"Harry",
			"James",
			"",
			"",
			// This PN group will render trailing separators.
			false,
		),
		// Add empty group that will render it's separators.
		personName.NewGroupEmpty(false),
		// Add empty group that will render it's separators.
		personName.NewGroupEmpty(false),
		// Remove groups from string render that do not add information.
		true,
	)

	// Print the string, should render as 'Potter^Harry^James^^'. Even though those
	// groups were set to render trailing separators, because they are empty, they are
	// suppressed.
	fmt.Println("PN:", pn.String())

	// Now lets make one that still renders empty groups with String().
	pn = personName.New(
		personName.NewGroupInfo(
			"Potter",
			"Harry",
			"James",
			"",
			"",
			// This PN group will render trailing separators.
			false,
		),
		// Add empty group that will render it's separators.
		personName.NewGroupEmpty(false),
		// Add empty group that will render it's separators.
		personName.NewGroupEmpty(false),
		// Do not remove groups from string render that do not add information.
		false,
	)

	// This will render as 'Potter^Harry^James^^===^^^^===^^^^'
	fmt.Println("PN:", pn.String())
}
