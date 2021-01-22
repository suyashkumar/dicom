package personname

import (
	"strings"
)

const segmentSep = "^"

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

	// HasNullSeparators will remove repeated separators around null groups when
	// calling DCM() if set to true.
	HasNullSeparators bool
}

// DCM Returns original, formatted string in
// '[FamilyName]^[GivenName]^[MiddleName]^[NamePrefix]^[NameSuffix]'.
func (group GroupInfo) DCM() string {
	dcmString := strings.Join(
		[]string{
			group.FamilyName,
			group.GivenName,
			group.MiddleName,
			group.NamePrefix,
			group.NameSuffix,
		},
		segmentSep,
	)

	if !group.HasNullSeparators {
		dcmString = strings.TrimRight(dcmString, segmentSep)
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

	// Range over the groups and assign them based on index.
	hasEmpty := false
	for i, groupValue := range segments {
		// If this group value is empty, set hasEmpty to true
		if groupValue == "" {
			hasEmpty = true
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
	if len(groupString) > 0 && hasEmpty {
		groupInfo.HasNullSeparators = true
	}

	return groupInfo, nil
}
