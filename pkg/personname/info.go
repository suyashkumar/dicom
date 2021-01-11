/*
The personname package provides methods and data types for inspecting Person Name (PN)
Value Representations, as defined here:

http://dicom.nema.org/dicom/2013/output/chtml/part05/sect_6.2.html
*/
package personname

import (
	"fmt"
	"strings"
)

const groupSep = "="

// pnGroup is an enum value for the PN group types (Alphabetic, Ideographic & Phonetic).
type pnGroup int

// String representation -- mostly for formatting error messages.
func (group pnGroup) String() string {
	switch group {
	case 0:
		return "Alphabetic"
	case 1:
		return "Ideographic"
	case 2:
		return "Phonetic"
	default:
		panic(fmt.Errorf("bad pnGroup value: %v", int(group)))
	}
}

// Enum values for pnGroup
const (
	pnGroupAlphabetic pnGroup = iota
	pnGroupIdeographic
	pnGroupPhonetic
)

// Info holds information from an element with a "PN" VR. See the "PN"
// entry at: http://dicom.nema.org/dicom/2013/output/chtml/part05/sect_6.2.html
//
// PN values are split into three groups which represent three different ways to
// represent a name:
//
// - Alphabetic: How a name is formally spelled using a Phonetic alphabet.
// - Ideographic: How a name is represented using ideograms / ideographs.
// - Phonetic: How a name is pronounced.
//
// Each of these groups can be inspected to access their individual segments (family
// name, Given name, etc.)
type Info struct {
	// Alphabetic group information about the Alphabetic group.
	Alphabetic GroupInfo
	// Ideographic group information about the Ideographic group.
	Ideographic GroupInfo
	// Phonetic group information about the Phonetic group.
	Phonetic GroupInfo

	// NoNullSeparators will remove repeated separators around null groups when
	// calling DCM() if set to true.
	NoNullSeparators bool
}

// WithNullSeparators returns a new Info object that will keep trailing separators
// that surround both null groups AND group segments: (ex: 'Potter^Harry^^^==').
func (info Info) WithNullSeparators() Info {
	info.NoNullSeparators = false
	info.Alphabetic.NoNullSeparators = false
	info.Ideographic.NoNullSeparators = false
	info.Phonetic.NoNullSeparators = false

	return info
}

// WithoutNullSeparators returns a new Info object that will remove trailing
// separators that surround both null groups AND group segments:
// (ex: 'Potter^Harry').
func (info Info) WithoutNullSeparators() Info {
	info.NoNullSeparators = true
	info.Alphabetic.NoNullSeparators = true
	info.Ideographic.NoNullSeparators = true
	info.Phonetic.NoNullSeparators = true

	return info
}

// dcmRemoveNullStrings removes null group strings from the list of groups to be
// rendered.
func (info Info) dcmRemoveNullStrings(groupStrings []string) []string {
	// If we are not removing null separators, return immediately.
	if !info.NoNullSeparators {
		return groupStrings
	}

	// If we are removing trailing emtpy, we are going to check if the last group is
	// empty, and if it is, trim it off the end until we hit a non-empty slice.
	//
	// If we were to move front-to-back instead of back-to-front, we would get a bad
	// result if the middle group is blank, but the third group is not, as we would
	// bail on the middle group, and not end up rendering the last.
	//
	// 'Harry^Potter==Phonetic^Spelling is a valid PN value, which we need to account
	// for.

	// Start with the sliceOut point containing all three groups
	sliceOut := 3
	// Iterate backwards over the strings
	for i := 2; i >= 0; i-- {
		// If the string is blank, shift the slice out back 1 place.
		if groupStrings[i] == "" {
			sliceOut = i
		} else {
			// Otherwise, if it is not blank, we have to render all remaining
			// separators, even if their values are null, so bail.
			break
		}
	}

	// Trim the groups by getting a slice of our slice using the sliceOut value.
	groupStrings = groupStrings[0:sliceOut]
	return groupStrings
}

// DCM returns the serialized DICOM representation of the PN value, in
// '[Alphabetic]=[Ideographic]=[Phonetic]' format.
func (info Info) DCM() string {
	// Convert the groups into their formatted string representations.=
	groupStrings := []string{
		info.Alphabetic.DCM(), info.Ideographic.DCM(), info.Phonetic.DCM(),
	}

	// Remove groups based on the trailing null setting and the formatted group strings.
	groupStrings = info.dcmRemoveNullStrings(groupStrings)

	// Join the remaining groups with '='
	return strings.Join(groupStrings, groupSep)
}

// IsEmpty returns whether the PN value contains any actual information. This method
// ignores separator, so both '' and '^^^^=^^^^=^^^^' would return true.
func (info Info) IsEmpty() bool {
	return info.Alphabetic.IsEmpty() &&
		info.Ideographic.IsEmpty() &&
		info.Phonetic.IsEmpty()
}

// Parse PN dicom value into a personname.Info value.
func Parse(valueString string) (Info, error) {
	groups := strings.Split(valueString, groupSep)

	// If there are more than three groups, then the value does not conform to the DICOM
	// spec.
	if len(groups) > 3 {
		return Info{}, newErrTooManyGroups(len(groups))
	}

	// Set up out new info value.
	info := Info{}

	// Range over the groups and assign them based on index.
	for i, groupString := range groups {
		// Convert the index into a pnGroup enum.
		group := pnGroup(i)

		// Parse the group info.
		groupInfo, err := groupFromValueString(groupString, group)
		if err != nil {
			return Info{}, err
		}

		// Apply the group info.
		switch group {
		case pnGroupAlphabetic:
			info.Alphabetic = groupInfo
		case pnGroupIdeographic:
			info.Ideographic = groupInfo
		case pnGroupPhonetic:
			info.Phonetic = groupInfo
		}
	}

	// If there are less than three groups, that means this value was removing null
	// separators, and this behavior should be replicated when calling Info.DCM().
	if len(groups) < 3 {
		info.NoNullSeparators = true

		// If we were missing the last group, we know there was no ^^^^ either, so we
		// need to reflect that in the Phonetic group.
		info.Phonetic.NoNullSeparators = true

		// Same idea if we are missing the second-to-last group (ideographic).
		if len(groups) < 2 {
			info.Ideographic.NoNullSeparators = true
		}

		// Split will always result in at least one value, even on an emtpy slice, so
		// we don't need to worry about null separators on the Alphabetic group.
	}

	return info, nil
}
