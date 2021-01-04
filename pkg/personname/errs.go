package personname

import (
	"errors"
	"fmt"
)

// ErrParsePersonName is returned when attempting to parse Info from a string.
var ErrParsePersonName = errors.New(
	"string contains to many segments for PN value",
)

func newErrTooManyGroups(groupsFound int) error {
	return fmt.Errorf(
		"%w: PN contains %v groups. No more than 3 groups with "+
			"'[Alphabetic]%v[Ideographic]%v[Phonetic]' format are allowed",
		ErrParsePersonName,
		groupsFound,
		groupSep,
		groupSep,
	)
}

func newErrTooManyGroupSegments(group string, segmentsFound int) error {
	return fmt.Errorf(
		"%w: PN group %v contains %v segments. No more than 5 segments with "+
			"'[Last]%v[First]%v[Middle]%v[Prefix]%v[Suffix]' format are allowed",
		ErrParsePersonName,
		group,
		segmentsFound,
		segmentSep,
		segmentSep,
		segmentSep,
		segmentSep,
	)
}
