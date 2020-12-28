package personName

import "strings"

const segmentSep = "^"

// The dicom spec splits each PN value into 3 representation groups:
//
//	- Alphabetic
//	- Ideographic
//	- Phonetic
//
// GroupInfo holds the parsed information for any one of these groups.
type GroupInfo struct {
	// The original Raw string which created this value.
	Raw string
	// The person's family or last name.
	FamilyName string
	// The person's given or first names.
	GivenName string
	// The person's middle names.
	MiddleName string
	// The person's name prefix.
	NamePrefix string
	// The person's name suffix.
	NameSuffix string
}

// Returns original, formatted string in
// '[FamilyName]^[GivenName]^[MiddleName]^[NamePrefix]^[NameSuffix]'.
func (group GroupInfo) String() string {
	return group.Raw
}

// Returns true if all group segments are empty, even if Raw value was "^^^^".
func (group GroupInfo) IsEmpty() bool {
	return group.FamilyName == "" &&
		group.GivenName == "" &&
		group.MiddleName == "" &&
		group.NamePrefix == "" &&
		group.NameSuffix == ""
}

// NewGroupInfo returns a group info object with the given segments. If
// removeTrailingSeparators is set to true, null trailing segments and their separators
// will be removed for the Raw field, so "Potter^Harry^James^^" will be rendered as
// "Potter^Harry^James" instead.
func NewGroupInfo(
	familyName string,
	givenName string,
	middleName string,
	namePrefix string,
	nameSuffix string,
	removeTrailingSeparators bool,
) GroupInfo {
	info := GroupInfo{
		Raw:        "",
		FamilyName: familyName,
		GivenName:  givenName,
		MiddleName: middleName,
		NamePrefix: namePrefix,
		NameSuffix: nameSuffix,
	}

	rawString := strings.Join(
		[]string{familyName, givenName, middleName, namePrefix, nameSuffix},
		segmentSep,
	)

	if removeTrailingSeparators {
		rawString = strings.TrimRight(rawString, segmentSep)
	}

	info.Raw = rawString
	return info
}

// Returns an empty GroupInfo value, if noSeps is false, returned String() value
// will be "^^^^", otherwise it will be "".
func NewGroupEmpty(noSeps bool) GroupInfo {
	return NewGroupInfo(
		"",
		"",
		"",
		"",
		"",
		noSeps,
	)
}

// groupFromValueString converts a string from a dicom element with a Value
// Representation of PN to a parsed Info struct.
func groupFromValueString(groupString string, group string) (GroupInfo, error) {
	segments := strings.Split(groupString, segmentSep)

	if len(segments) > 5 {
		return GroupInfo{}, newErrParsePersonNameTooGroupSegments(group, len(segments))
	}

	groupInfo := GroupInfo{Raw: groupString}

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

	return groupInfo, nil
}
