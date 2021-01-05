package main

import (
	"strconv"
	"strings"
)

// Tag holds tag group and element.
type Tag struct {
	// Group is the metadata group piece of the tag.
	Group uint16
	// Element is the metadata element identifier half of the tag.
	Element uint16
}

// TagInfo is a mirror of tag.Info that we will parse our spec data into.
type TagInfo struct {
	// Tag is the (XXXX,YYYY) metadata element identifier.
	Tag Tag
	// VR is the DICOM Value Representation string that indicates what type of data
	// this tag holds.
	VR string
	// Name is the dicom machine readable Keyword associated with the tag.
	Name string
	// VM is the value multiplicity of the tag (how many VR values does / can the
	// element store)
	VM string
}

// ParseTag parses a Tag from a (XXXX,YYYY) string.
func ParseTag(tag string) (Tag, error) {
	parts := strings.Split(strings.Trim(tag, "()"), ",")
	group, err := strconv.ParseInt(parts[0], 16, 0)
	if err != nil {
		return Tag{}, err
	}
	elem, err := strconv.ParseInt(parts[1], 16, 0)
	if err != nil {
		return Tag{}, err
	}
	return Tag{Group: uint16(group), Element: uint16(elem)}, nil
}
