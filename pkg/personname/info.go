/*
The personName package provides methods and data types for inspecting Person Name (PN)
Value Representations, as defined here:

http://dicom.nema.org/dicom/2013/output/chtml/part05/sect_6.2.html
*/
package personname

import (
	"strings"
)

const groupSep = "="

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
	// calling String() if set to true.
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

// String returns the original Raw representation of the PN value, in
// '[Alphabetic]=[Ideographic]=[Phonetic]' format.
func (info Info) String() string {
	groups := []GroupInfo{info.Alphabetic, info.Ideographic, info.Phonetic}

	// If we are removing trailing emtpy, we are going to check if the last group is
	// empty, then trim it off the end if it is, until we hit a non-empty slice.
	if info.NoNullSeparators {
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

	return strings.Join(rawGroups, groupSep)
}

// IsEmpty returns whether the PN value contains any actual information. This method
// ignores separator, so both '' and '^^^^=^^^^=^^^^' would return true.
func (info Info) IsEmpty() bool {
	return info.Alphabetic.IsEmpty() &&
		info.Ideographic.IsEmpty() &&
		info.Phonetic.IsEmpty()
}

// Parse PN dicom value into informational value.
func Parse(valueString string) (Info, error) {
	groups := strings.Split(valueString, groupSep)

	if len(groups) > 3 {
		return Info{}, newErrTooManyGroups(len(groups))
	}

	info := Info{}

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

	// If there are less than three groups, that means this value was removing null
	// separators, and this behavior should be replicated when calling Info.String().
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
