package personname

import (
	"strings"
)

const segmentSep = "^"

// GroupTrailingNullLevel represents how many null '^' separators are present in the
// GroupInfo.DCM() return value.
type GroupTrailingNullLevel uint

// String implements fmt.Stringer, giving human-readable names to the trailing null
// level.
//
// Returns "NONE" if no null separators were present.
//
// Returns "ALL" if the highest possible null separator was present.
//
// Otherwise, returns the name of the section that comes after the highest present null
// separator.
//
// String will panic if called on a value that exceeds GroupNullLevelAll.
func (level GroupTrailingNullLevel) String() string {
	switch level {
	case GroupNullLevelNone:
		return "NONE"
	case GroupNullLevelGiven:
		return "GivenName"
	case GroupNullLevelMiddle:
		return "MiddleName"
	case GroupNullLevelPrefix:
		return "NamePrefix"
	case GroupNullLevelAll:
		return "ALL"
	default:
		return "[INVALID]"
	}
}

const (
	// GroupNullLevelNone will render no null seps.
	GroupNullLevelNone GroupTrailingNullLevel = iota

	// NullSepGiven will render null separators up to the separator before the
	// GroupInfo.GivenName segment
	GroupNullLevelGiven

	// NullSepGiven will render null separators up to the separator before the
	// GroupInfo.MiddleName segment
	GroupNullLevelMiddle

	// NullSepGiven will render null separators up to the separator before the
	// GroupInfo.NamePrefix segment
	GroupNullLevelPrefix

	// NullSepGiven will render null separators up to the separator before the
	// GroupInfo.NameSuffix segment (ALL possible separators).
	GroupNullLevelAll
)

func validateGroupNullSepLevel(level GroupTrailingNullLevel) error {
	if level <= GroupNullLevelAll {
		return nil
	}

	return newErrNullSepLevelInvalid(uint(GroupNullLevelAll), uint(level))
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

	// TrailingNullLevel contains the highest present null '^' separator in the DCM()
	// value. For most use cases GroupNullLevelAll or GroupNullLevelNone should be used when
	// creating new PN values. Use other levels only if you know what you are doing!
	TrailingNullLevel GroupTrailingNullLevel
}

// DCM Returns original, formatted string in
// '[FamilyName]^[GivenName]^[MiddleName]^[NamePrefix]^[NameSuffix]'.
func (group GroupInfo) DCM() (string, error) {
	// validate our TrailingNullLevel and panic if it is exceeded.
	if err := validateGroupNullSepLevel(group.TrailingNullLevel); err != nil {
		return "", err
	}

	// Put all the segments into an array.
	segments := []string{
		group.FamilyName,
		group.GivenName,
		group.MiddleName,
		group.NamePrefix,
		group.NameSuffix,
	}

	// Render our segments with the correct number of null-separators.
	return renderWithSeps(segments, segmentSep, uint(group.TrailingNullLevel)), nil
}

// MustDCM is as DCM, but panics on error.
//
// MustDCM will only panic if TrailingNullLevel exceeds GroupNullLevelAll.
func (group GroupInfo) MustDCM() string {
	dcm, err := group.DCM()
	if err != nil {
		panic(err)
	}
	return dcm
}

// IsEmpty returns true if all group segments are empty, even if Raw value was "^^^^".
func (group GroupInfo) IsEmpty() bool {
	return group.FamilyName == "" &&
		group.GivenName == "" &&
		group.MiddleName == "" &&
		group.NamePrefix == "" &&
		group.NameSuffix == ""
}

// Veterinary returns a read / write data access object that points to the calling
// GroupInfo value with helper methods specific veterinary contexts.
//
// This helper objects should not be passed around itself, but used in chained commands:
//
//  petName := groupInfo.Veterinary().PatientName()
//  groupInfo.Veterinary().SetPatientName(petName)
func (group *GroupInfo) Veterinary() VeterinaryInfo {
	return VeterinaryInfo{groupInfo: group}
}

// VeterinaryInfo acts as a namespace for veterinary-specific methods on GroupInfo.
//
// VeterinaryInfo should not be extracted and used on it's own. It is included as a more
// explicit API for veterinary settings and should be invoked via chained calls:
//
//  petName := groupInfo.Veterinary().PatientName()
//  groupInfo.Veterinary().SetPatientName(petName)
type VeterinaryInfo struct {
	groupInfo *GroupInfo
}

// ResponsibleParty returns the Responsible Party Family Name / Responsible
// organization name for the patient.
//
// ResponsibleParty simply returns GroupInfo.FamilyName, and is included a more explicit
// way of communicating you are getting this information in a Veterinary context.
func (vet VeterinaryInfo) ResponsibleParty() string {
	return vet.groupInfo.FamilyName
}

// SetResponsibleParty sets the Responsible Party Family Name / Responsible
// organization name for the patient.
//
// SetResponsibleParty simply sets GroupInfo.FamilyName, and is included a more explicit
// way of communicating you are getting this information in a Veterinary context.
func (vet VeterinaryInfo) SetResponsibleParty(name string) {
	vet.groupInfo.FamilyName = name
}

// PatientName is the full name of the Pet, Animal, etc.
//
// PatientName simply returns GroupInfo.GivenName, and is included a more explicit
// way of communicating you are getting this information in a Veterinary context.
func (vet VeterinaryInfo) PatientName() string {
	return vet.groupInfo.GivenName
}

// SetPatientName sets the full name of the Pet, Animal, etc.
//
// SetPatientName simply sets GroupInfo.GivenName, and is included a more explicit
// way of communicating you are getting this information in a Veterinary context.
func (vet VeterinaryInfo) SetPatientName(name string) {
	vet.groupInfo.GivenName = name
}

// NewVeterinaryGroupInfo is a helper function for creating new GroupInfo values in a
// veterinary context.
//
// Returns a GroupInfo value with GroupInfo.FamilyName set to responsibleParty,
// GroupInfo.GivenName set to patientName, and GroupInfo.TrailingNullLevel set to
// GroupNullLevelNone.
//
// From the dicom spec:
//
//  For veterinary use, the first two of the five components in their order of
//  occurrence are: responsible party family name or responsible organization name,
//  patient name. The remaining components are not used and shall not be present.
func NewVeterinaryGroupInfo(responsibleParty, patientName string) GroupInfo {
	return GroupInfo{
		FamilyName:        responsibleParty,
		GivenName:         patientName,
		TrailingNullLevel: GroupNullLevelNone,
	}
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
	nullSepLevel := GroupNullLevelNone
	for i, groupValue := range segments {
		// If this segment is empty, it means there is a null sep here. Our null sep
		// level needs to reflect this.
		if groupValue == "" {
			nullSepLevel = GroupTrailingNullLevel(i)
		} else {
			// Otherwise, if there is a non-zero string value, there is no null sep
			// after it.
			nullSepLevel = GroupNullLevelNone
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
		groupInfo.TrailingNullLevel = nullSepLevel
	}

	return groupInfo, nil
}
