package personName

import (
	"errors"
	"fmt"
)

// ErrParsePersonName is returned when attempting to parse Info from a string.
var ErrParsePersonName = errors.New(
	"string contains to many segments for PN value",
)

func newErrParsePersonNameTooManyGroups(groupsFound int) error {
	return fmt.Errorf(
		"%w: PN contains %v groups. No more than 3 groups with "+
			"'[alphabetic]===[ideographic]===[phonetic]' format are allowed",
		ErrParsePersonName,
		groupsFound,
	)
}

func newErrParsePersonNameTooGroupSegments(group string, segmentsFound int) error {
	return fmt.Errorf(
		"%w: PN group %v contains %v segments. No more than 5 segments with "+
			"'[Last]^[First]^[Middle]^[Prefix]^[Suffix]' format are allowed",
		ErrParsePersonName,
		group,
		segmentsFound,
	)
}
