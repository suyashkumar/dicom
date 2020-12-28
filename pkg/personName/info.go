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
//	- Alphabetic: How a name is formally spelled using a Phonetic alphabet.
//  - Ideographic: How a name is represented using ideograms / ideographs.
//  - Phonetic: How a name is pronounced.
//
// Each of these groups can be inspected to access their individual segments (family
// name, Given name, etc.)
type Info struct {
	// The original, Raw PN value.
	Raw string
	// Expected information about the Alphabetic group.
	Alphabetic GroupInfo
	// Expected information about the Ideographic group.
	Ideographic GroupInfo
	// Expected information about the Phonetic group.
	Phonetic GroupInfo
}

// Original Raw representation of the PN value, in
// '[Alphabetic]===[Ideographic]===[Phonetic]' format.
func (info Info) String() string {
	return info.Raw
}

// IsEmpty returns whether the PN value contains any actual information. This method
// ignores separator, so both '' and '^^^^=^^^^=^^^^' would return true.
func (info Info) IsEmpty() bool {
	return info.Alphabetic.IsEmpty() &&
		info.Ideographic.IsEmpty() &&
		info.Phonetic.IsEmpty()
}

// Creates a new Info object detailing the individual data.
//
// If removeTrailingEmpty is set to true, null trailing groups and their separators
// will be removed for the Raw field if they contain no information, so
// "Potter^Harry^James^^=^^^^=^^^^" will be rendered as "Potter^Harry^James^^"
// instead.
func New(
	alphabetic GroupInfo,
	ideographic GroupInfo,
	phonetic GroupInfo,
	removeTrailingEmpty bool,
) Info {
	info := Info{
		Raw:         "",
		Alphabetic:  alphabetic,
		Ideographic: ideographic,
		Phonetic:    phonetic,
	}

	groups := []GroupInfo{info.Alphabetic, info.Ideographic, info.Phonetic}

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

	info.Raw = strings.Join(rawGroups, groupSep)

	return info
}

// Parse PN dicom value into informational value.
func Parse(valueString string) (Info, error) {
	groups := strings.Split(valueString, groupSep)

	if len(groups) > 3 {
		return Info{}, newErrParsePersonNameTooManyGroups(len(groups))
	}

	info := Info{Raw: valueString}

	// Range over the groups and assign them based on index.
	for i, groupString := range groups {
		switch i {
		case 0:
			groupInfo, err := groupFromValueString(groupString, "Alphabetic")
			if err != nil {
				return Info{}, err
			}
			info.Alphabetic = groupInfo
		case 1:
			groupInfo, err := groupFromValueString(groupString, "Ideographic")
			if err != nil {
				return Info{}, err
			}
			info.Ideographic = groupInfo
		case 2:
			groupInfo, err := groupFromValueString(groupString, "Phonetic")
			if err != nil {
				return Info{}, err
			}
			info.Phonetic = groupInfo
		}
	}

	return info, nil
}
