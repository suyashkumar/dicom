package pn

import (
	"errors"
	"regexp"
)

// Parses PersonName value from an element with a "PN" VR. See the "PN" entry at:
// http://dicom.nema.org/dicom/2013/output/chtml/part05/sect_6.2.html, and the
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

var pnRegex = regexp.MustCompile(
	`(?P<Fist>.*)\^(?P<Last>.*)\^(?P<Middle>.*)\^(?P<Prefix>.*)\^(?P<Suffix>.*)`,
)

var ErrParsePersonName = errors.New("person name value does not match Dicom Spec" +
	" '[Last]^[First]^[Middle]^[Prefix]^[Suffix]'")

// Converts a Person Name string to parsed PersonName struct.
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