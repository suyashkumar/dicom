/*
The personName package provides methods and data types for inspecting Person Name (PN)
Value Representations, as defined here:

http://dicom.nema.org/dicom/2013/output/chtml/part05/sect_6.2.html
*/
package personName

import (
	"strings"
)

const groupSep = "="

// Expected Info value information from an element with a "PN" VR. See the "PN"
// entry at: http://dicom.nema.org/dicom/2013/output/chtml/part05/sect_6.2.html
//
// PN values are split into three groups which represent three different ways to
// represent a name:
//
//	- Alphabetic: How a name is formally spelled using a phonetic alphabet.
//  - Ideographic: How a name is represented using ideograms / ideographs.
//  - Phonetic: How a name is pronounced.
//
// Each of these groups can be inspected to access their individual segments (family
// name, Given name, etc.)
type Info struct {
	// The original, raw PN value.
	raw string
	// Expected information about the alphabetic group.
	alphabetic GroupInfo
	// Expected information about the ideographic group.
	ideographic GroupInfo
	// Expected information about the phonetic group.
	phonetic GroupInfo
}

// Returns PN value group 1: Alphabetic representation of the person's name.
func (info Info) Alphabetic() GroupInfo {
	return info.alphabetic
}

// Returns PN value group 2: Ideographic representation of the person's name.
func (info Info) Ideographic() GroupInfo {
	return info.ideographic
}

// Returns PN value group 3: Phonetic representation of the person's name.
func (info Info) Phonetic() GroupInfo {
	return info.phonetic
}

// Original raw representation of the PN value, in
// '[alphabetic]===[ideographic]===[phonetic]' format.
func (info Info) String() string {
	return info.raw
}

// IsEmpty returns whether the PN value contains any actual information. This method
// ignores separator, so both '' and '^^^^=^^^^=^^^^' would return true.
func (info Info) IsEmpty() bool {
	return info.alphabetic.IsEmpty() &&
		info.ideographic.IsEmpty() &&
		info.phonetic.IsEmpty()
}

// Creates a new Info object detailing the individual data.
//
// If removeTrailingEmpty is set to true, null trailing groups and their separators
// will be removed for the raw field if they contain no information, so
// "Potter^Harry^James^^=^^^^=^^^^" will be rendered as "Potter^Harry^James^^"
// instead.
func New(
	alphabetic GroupInfo,
	ideographic GroupInfo,
	phonetic GroupInfo,
	removeTrailingEmpty bool,
) Info {
	info := Info{
		raw:         "",
		alphabetic:  alphabetic,
		ideographic: ideographic,
		phonetic:    phonetic,
	}

	groups := []GroupInfo{info.Alphabetic(), info.Ideographic(), info.Phonetic()}

	// If we are removing trailing emtpy, we are going to check if the last group is
	// empty, then trim it off the end if it is, until we hit a non-empty slice.
	if removeTrailingEmpty {
		out := 3
		for i := 2; i >= 0; i-- {
			if groups[i].IsEmpty() {
				out = i
			} else {
				break
			}
		}

		// Trim the groups by getting a slice of our slice.
		groups = groups[0:out]
	}

	rawGroups := make([]string, len(groups))

	for i := 0; i < len(groups); i++ {
		rawGroups[i] = groups[i].String()
	}

	info.raw = strings.Join(rawGroups, groupSep)

	return info
}

// Parse PN dicom value into informational value.
func Parse(valueString string) (Info, error) {
	groups := strings.Split(valueString, groupSep)

	if len(groups) > 3 {
		return Info{}, newErrParsePersonNameTooManyGroups(len(groups))
	}

	info := Info{raw: valueString}

	// Range over the groups and assign them based on index.
	for i, groupString := range groups {
		switch i {
		case 0:
			groupInfo, err := groupFromValueString(groupString, "alphabetic")
			if err != nil {
				return Info{}, err
			}
			info.alphabetic = groupInfo
		case 1:
			groupInfo, err := groupFromValueString(groupString, "ideographic")
			if err != nil {
				return Info{}, err
			}
			info.ideographic = groupInfo
		case 2:
			groupInfo, err := groupFromValueString(groupString, "phonetic")
			if err != nil {
				return Info{}, err
			}
			info.phonetic = groupInfo
		}
	}

	return info, nil
}
