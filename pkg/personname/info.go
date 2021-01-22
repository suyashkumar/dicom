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
	case pnGroupAlphabetic:
		return "Alphabetic"
	case pnGroupIdeographic:
		return "Ideographic"
	case pnGroupPhonetic:
		return "Phonetic"
	default:
		panic(fmt.Errorf("bad pnGroup value: %v", int(group)))
	}
}

// Enum values for pnGroup
const (
	pnGroupAlphabetic  = 0
	pnGroupIdeographic = 1
	pnGroupPhonetic    = 2
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

	// HasNullSeparators will remove repeated separators around null groups when
	// calling DCM() if set to true.
	HasNullSeparators bool
}

// WithFormat returns a new Info object with null separator settings applied to the
// relevant Info / GroupInfo objects.
//
// useGroupNullSeparators will add trailing "=" characters, even when one or more
// trailing groups are emtpy strings.
//
// The remaining options will apply the passed value to their groups respective
// HasNullSeparators value.
//
// This method does not mutate its receiver value, instead returning a new value
// to the caller with the passed settings.
func (info Info) WithFormat(
	useGroupNullSeparators,
	useAlphabeticNullSeparators,
	useIdeographicNullSeparators,
	usePhoneticNullSeparators bool,
) Info {
	info.HasNullSeparators = useGroupNullSeparators
	info.Alphabetic.HasNullSeparators = useAlphabeticNullSeparators
	info.Ideographic.HasNullSeparators = useIdeographicNullSeparators
	info.Phonetic.HasNullSeparators = usePhoneticNullSeparators
	return info
}

// WithNullSeparators returns a new Info object that will keep trailing separators
// that surround both null groups AND group segments: (ex: 'Potter^Harry^^^==').
//
// WithNullSeparators is equivalent to calling WithFormat() with all options set to
// true.
//
// WithNullSeparators does not mutate its receiver value, instead returning a new value
// to the caller with the passed settings.
func (info Info) WithNullSeparators() Info {
	info.HasNullSeparators = true
	info.Alphabetic.HasNullSeparators = true
	info.Ideographic.HasNullSeparators = true
	info.Phonetic.HasNullSeparators = true

	return info
}

// WithoutNullSeparators returns a new Info object that will remove trailing
// separators that surround both null groups AND group segments:
// (ex: 'Potter^Harry').
//
// WithoutNullSeparators is equivalent to calling WithFormat() with all options set to
// false.
//
// WithoutNullSeparators does not mutate its receiver value, instead returning a new
// value to the caller with the passed settings.
func (info Info) WithoutNullSeparators() Info {
	info.HasNullSeparators = false
	info.Alphabetic.HasNullSeparators = false
	info.Ideographic.HasNullSeparators = false
	info.Phonetic.HasNullSeparators = false

	return info
}

// WithoutEmptyGroups sets Info.HasNullSeparators to false, then checks eac
// group, and if it contains no actual information, sets that group's HasNullSeparators
// to false.
//
// Groups with Partial information will retain their null separators.
func (info Info) WithoutEmptyGroups() Info {
	info.HasNullSeparators = false

	// Iterate over references to our group values (we aren't mutating our receiver
	// here since it's passed by value and already a deep copy).
	for _, group := range []*GroupInfo{&info.Alphabetic, &info.Ideographic, &info.Phonetic} {
		if group.IsEmpty() {
			group.HasNullSeparators = false
		}
	}

	return info
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

// dcmRemoveNullStrings removes null group strings from the list of groups to be
// rendered.
func (info Info) dcmRemoveNullStrings(groupStrings []string) []string {
	// If we are not removing null separators, return immediately.
	if info.HasNullSeparators {
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

	// Start with the groupCount point containing all three groups
	groupCount := len(groupStrings)
	// Iterate backwards over the strings
	for i := groupCount - 1; i >= 0; i-- {
		// If the string is blank, shift the slice out back 1 place.
		if groupStrings[i] == "" {
			groupCount = i
		} else {
			// Otherwise, if it is not blank, we have to render all remaining
			// separators, even if their values are null, so bail.
			break
		}
	}

	// Trim the groups by getting a slice of our slice using the groupCount value.
	groupStrings = groupStrings[0:groupCount]
	return groupStrings
}

// IsEmpty returns whether the PN value contains any actual information. This method
// ignores separator, so both '' and '^^^^=^^^^=^^^^' would return true.
func (info Info) IsEmpty() bool {
	return info.Alphabetic.IsEmpty() &&
		info.Ideographic.IsEmpty() &&
		info.Phonetic.IsEmpty()
}

// Parse PN dicom value into a personname.Info value.
//
// NOTE ON PARSING:
//
// The personname.Info and personname.GroupInfo values only track whether any null
// separators were used, not how many. This means if a PN value has some null
// separators, but not the full amount, round-tripping the value will result in adding
// the missing separators. See examples below. If you wish to make sure that NO
// alterations are made to the original value after inspecting re-serializing, the
// original value should be used directly instead.
func Parse(valueString string) (Info, error) {
	groups := strings.Split(valueString, groupSep)

	// If there are more than three groups, then the value does not conform to the DICOM
	// spec.
	if len(groups) > 3 {
		return Info{}, newErrTooManyGroups(len(groups))
	}

	// Set up a new info value.
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
		info.HasNullSeparators = false

		// If we were missing the last group, we know there was no ^^^^ either, so we
		// need to reflect that in the Phonetic group.
		info.Phonetic.HasNullSeparators = false

		// Same idea if we are missing the second-to-last group (ideographic).
		if len(groups) < 2 {
			info.Ideographic.HasNullSeparators = false
		}

		// Split will always result in at least one value, even on an emtpy slice, so
		// we don't need to worry about null separators on the Alphabetic group.
	}

	return info, nil
}
