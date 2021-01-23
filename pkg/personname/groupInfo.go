package personname

import (
	"strings"
)

const segmentSep = "^"
const groupSegmentCount = 5

// GroupNullSepLevel represents how many null separators are present in the
// GroupInfo.DCM() return value.
type GroupNullSepLevel uint

const (
	// GroupNullSepNone will render no null seps.
	GroupNullSepNone GroupNullSepLevel = iota

	// NullSepGiven will render null separators up to the separator before the
	// GroupInfo.GivenName segment
	GroupNullSepGiven

	// NullSepGiven will render null separators up to the separator before the
	// GroupInfo.MiddleName segment
	GroupNullSepMiddle

	// NullSepGiven will render null separators up to the separator before the
	// GroupInfo.NamePrefix segment
	GroupNullSepPrefix

	// NullSepGiven will render null separators up to the separator before the
	// GroupInfo.NameSuffix segment, or ALL possible separators.
	GroupNullSepAll
)

func validateNullSepLevel(level GroupNullSepLevel) error {
	if level <= GroupNullSepAll {
		return nil
	}

	return newErrNullSepLevelInvalid(
		"group", uint(GroupNullSepAll), uint(level),
	)
}

// GroupInfo holds the parsed information for any one of these groups the person name
// groups specified in the DICOM spec:
//
//	- Alphabetic
//	- Ideographic
//	- Phonetic
type GroupInfo struct {
	// FamilyName is the person's family or last name.
	FamilyName string
	// GivenName if the person's given or first names.
	GivenName string
	// MiddleName is the person's middle name(s).
	MiddleName string
	// NamePrefix is the person's name prefix (ex: MR., MRS.).
	NamePrefix string
	// NameSuffix is the person's name suffix (ex: Jr, III).
	NameSuffix string

	// NullSepLevel is true when a PN value has trailing '=' separators around null
	// group values. If true, repeated separators around null groups will be removed
	// when calling DCM().
	NullSepLevel GroupNullSepLevel
}

// DCM Returns original, formatted string in
// '[FamilyName]^[GivenName]^[MiddleName]^[NamePrefix]^[NameSuffix]'.
func (group GroupInfo) DCM() string {
	// validate our NullSepLevel and panic if it is exceeded.
	err := validateNullSepLevel(group.NullSepLevel)
	if err != nil {
		panic(err)
	}

	// Put all the segments into an array.
	segments := [groupSegmentCount]string{
		group.FamilyName,
		group.GivenName,
		group.MiddleName,
		group.NamePrefix,
		group.NameSuffix,
	}

	// It's going to be easier to write the correct values if we iterate backwards over
	// them, since we need to look-ahead to know if there are interior null components.
	dcmString := ""
	nonZeroFound := false
	for i := groupSegmentCount - 1; i >= 0; i-- {
		segment := segments[i]
		dcmString = segment + dcmString

		if segment != "" {
			nonZeroFound = true
		}

		if i > 0 && (nonZeroFound || i <= int(group.NullSepLevel)) {
			dcmString = "^" + dcmString
		}
	}

	return dcmString
}

// IsEmpty returns true if all group segments are empty, even if Raw value was "^^^^".
func (group GroupInfo) IsEmpty() bool {
	return group.FamilyName == "" &&
		group.GivenName == "" &&
		group.MiddleName == "" &&
		group.NamePrefix == "" &&
		group.NameSuffix == ""
}

// groupFromValueString converts a string from a dicom element with a Value
// Representation of PN to a parsed Info struct.
func groupFromValueString(groupString string, group pnGroup) (GroupInfo, error) {
	segments := strings.Split(groupString, segmentSep)
	segmentCount := len(segments)

	if segmentCount > 5 {
		return GroupInfo{}, newErrTooManyGroupSegments(group, len(segments))
	}

	groupInfo := GroupInfo{}

	// Start off with our null segment level being None
	nullSegmentLevel := GroupNullSepNone
	for i, groupValue := range segments {
		// If this segment is empty, it means there is a null sep here. Our null sep
		// level needs to reflect this.
		if groupValue == "" {
			nullSegmentLevel = GroupNullSepLevel(i)
		} else {
			// Otherwise, if there is a non-zero string value, there is no null sep
			// after it.
			nullSegmentLevel = GroupNullSepNone
		}

		switch i {
		case 0:
			groupInfo.FamilyName = groupValue
		case 1:
			groupInfo.GivenName = groupValue
		case 2:
			groupInfo.MiddleName = groupValue
		case 3:
			groupInfo.NamePrefix = groupValue
		case 4:
			groupInfo.NameSuffix = groupValue
		}
	}

	// If the string is not empty, and any of our groups ARE empty, then we are using
	// null separators.
	if strings.HasSuffix(groupString, "^") {
		groupInfo.NullSepLevel = nullSegmentLevel
	}

	return groupInfo, nil
}
