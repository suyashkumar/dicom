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

	// NoNullSeparators will remove repeated separators around null groups when
	// calling DCM() if set to true.
	NoNullSeparators bool
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

	if group.NoNullSeparators {
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

	if len(segments) > 5 {
		return GroupInfo{}, newErrTooManyGroupSegments(group, len(segments))
	}

	groupInfo := GroupInfo{}

	// Range over the groups and assign them based on index.
	for i, groupValue := range segments {
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

	// If there are less than 5 segments, that means trailing separators were not
	// included, and when we call GroupInfo.DCM(), they should not be rendered.
	if len(segments) < 5 {
		groupInfo.NoNullSeparators = true
	}

	return groupInfo, nil
}
