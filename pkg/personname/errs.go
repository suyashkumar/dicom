package personname

import (
	"errors"
	"fmt"
)

// ErrParsePersonName is returned when attempting to parse Info from a string.
var ErrParsePersonName = errors.New(
	"error parsing PN value",
)

// ErrParseGroupCount is returned when a value has more than 3 groups, violating the
// dicom spec. ErrParseGroupCount wraps ErrParsePersonName.
var ErrParseGroupCount = fmt.Errorf(
	"%w: no more than 3 groups with '[Alphabetic]%v[Ideographic]%v[Phonetic]' "+
		"format are allowed",
	ErrParsePersonName,
	groupSep,
	groupSep,
)

// ErrParseGroupSegmentCount is returned when a group contains more than 5 segments,
// violating the dicom spec. ErrParseGroupSegmentCount wraps ErrParsePersonName.
var ErrParseGroupSegmentCount = fmt.Errorf(
	"%w: no more than 5 segments with"+
		" '[Last]%v[First]%v[Middle]%v[Prefix]%v[Suffix]' format are allowed",
	ErrParsePersonName,
	segmentSep,
	segmentSep,
	segmentSep,
	segmentSep,
)

// ErrNullSepLevelInvalid is a sentinel error returned when Info.TrailingNullLevel or
// GroupInfo.TrailingNullLevel exceeds the maximum allowed value.
var ErrNullSepLevelInvalid = fmt.Errorf("TrailingNullLevel exceeded maximum")

// newErrTooManyGroups returns a new ErrParseGroupCount wrapped with some context.
func newErrTooManyGroups(groupsFound int) error {
	return fmt.Errorf(
		"%w: value contains %v groups. see 'PN' entry in official dicom spec: "+
			"http://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_6.2",
		ErrParseGroupCount,
		groupsFound,
	)
}

// newErrTooManyGroupSegments returns a new ErrParseGroupSegmentCount wrapped with some
// context.
func newErrTooManyGroupSegments(group pnGroup, segmentsFound int) error {
	return fmt.Errorf(
		"%w: value group %v contains %v segments. see 'PN' entry in official "+
			"dicom spec: http://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_6.2",
		ErrParseGroupSegmentCount,
		group,
		segmentsFound,
	)
}

// newErrNullSepLevelInvalid returns a new ErrNullSepLevelInvalid wrapped with some
// context.
func newErrNullSepLevelInvalid(maxAllowed, found uint) error {
	return fmt.Errorf(
		"%w: cannot be greater than %v, got %v",
		ErrNullSepLevelInvalid,
		maxAllowed,
		found,
	)
}
