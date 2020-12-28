package personName

import "strings"

const segmentSep = "^"

// The dicom spec splits each PN value into 3 representation groups:
//
//	- alphabetic
//	- ideographic
//	- phonetic
//
// GroupInfo holds the parsed information for any one of these groups.
type GroupInfo struct {
	// The original raw string which created this value.
	raw string
	// The person's family or last name.
	familyName string
	// The person's given or first names.
	givenName string
	// The person's middle names.
	middleName string
	// The person's name prefix.
	namePrefix string
	// The person's name suffix.
	nameSuffix string
}

// FamilyName returns segment 1 of the PN: the person's family or last name.
func (group GroupInfo) FamilyName() string {
	return group.familyName
}

// GivenName returns segment 2 of the PN: the person's given or first names.
func (group GroupInfo) GivenName() string {
	return group.givenName
}

// MiddleName returns segment 3 of the PN: the person's middle names.
func (group GroupInfo) MiddleName() string {
	return group.middleName
}

// NamePrefix returns segment 4 of the PN: the person's name prefix: e.g., 'Mr' and
// 'Mrs'.
func (group GroupInfo) NamePrefix() string {
	return group.namePrefix
}

// NameSuffix returns segment 5 of the PN: the person's name suffix, e.g. 'Jr' and
// 'III'.
func (group GroupInfo) NameSuffix() string {
	return group.nameSuffix
}

// Returns original raw dicom format PN group string:
// '[Last]^[First]^[Middle]^[Prefix]^[Suffix]'.
func (group GroupInfo) String() string {
	return group.raw
}

// Returns true if all group segments are empty, even if raw value was "^^^^".
func (group GroupInfo) IsEmpty() bool {
	return group.familyName == "" &&
		group.givenName == "" &&
		group.middleName == "" &&
		group.namePrefix == "" &&
		group.nameSuffix == ""
}

// NewGroupInfo returns a group info object with the given segments. If
// removeTrailingSeparators is set to true, null trailing segments and their separators
// will be removed for the raw field, so "Potter^Harry^James^^" will be rendered as
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
		raw:        "",
		familyName: familyName,
		givenName:  givenName,
		middleName: middleName,
		namePrefix: namePrefix,
		nameSuffix: nameSuffix,
	}

	rawString := strings.Join(
		[]string{familyName, givenName, middleName, namePrefix, nameSuffix},
		segmentSep,
	)

	if removeTrailingSeparators {
		rawString = strings.TrimRight(rawString, segmentSep)
	}

	info.raw = rawString
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

	groupInfo := GroupInfo{raw: groupString}

	// Range over the groups and assign them based on index.
	for i, groupValue := range segments {
		switch i {
		case 0:
			groupInfo.familyName = groupValue
		case 1:
			groupInfo.givenName = groupValue
		case 2:
			groupInfo.middleName = groupValue
		case 3:
			groupInfo.namePrefix = groupValue
		case 4:
			groupInfo.nameSuffix = groupValue
		}
	}

	return groupInfo, nil
}
