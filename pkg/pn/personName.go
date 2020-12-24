/*
The pn package provides methods and data types for inspecting Person Name (PN)
Value Representations.
 */
package pn

import (
	"errors"
	"regexp"
)

// Parsed PersonName value information from an element with a "PN" VR. See the "PN"
// entry at: http://dicom.nema.org/dicom/2013/output/chtml/part05/sect_6.2.html
type PersonName struct {
	// The person's family or last name.
	FamilyName string
	// The person's given or first names
	GivenName string
	// The person's middle names
	MiddleName string
	// The person's name prefix
	NamePrefix string
	// The person's name suffix
	NameSuffix string
}

// Returns dicom format PN string: '[Last]^[First]^[Middle]^[Prefix]^[Suffix]'
func (pn PersonName) String() string {
	return pn.FamilyName +
		"^" + pn.GivenName +
		"^" + pn.MiddleName +
		"^" + pn.NamePrefix +
		"^" + pn.NameSuffix
}

// Regex for parsing PN values.
var pnRegex = regexp.MustCompile(
	`(?P<Fist>.*)\^(?P<Last>.*)\^(?P<Middle>.*)\^(?P<Prefix>.*)\^(?P<Suffix>.*)`,
)

// ErrParsePersonName is returned when attempting to parse PersonName from a string.
var ErrParsePersonName = errors.New("person name value does not match Dicom Spec" +
	" '[Last]^[First]^[Middle]^[Prefix]^[Suffix]'")

// Converts a string from a dicom element with a Value Representation of PN to a
// parsed PersonName struct.
func NewPersonNameFromDicom(pnString string) (name PersonName, err error) {
	// Run the regex against the name
	matches := pnRegex.FindStringSubmatch(pnString)
	if len(matches) == 0 || matches[0] == "" {
		return PersonName{}, ErrParsePersonName
	}
	name.FamilyName = matches[1]
	name.GivenName = matches[2]
	name.MiddleName = matches[3]
	name.NamePrefix = matches[4]
	name.NameSuffix = matches[5]

	return name, err
}