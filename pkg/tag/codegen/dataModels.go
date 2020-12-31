package main

import (
	"strconv"
	"strings"
)

// Holds tag group and element.
type Tag struct {
	Group   uint16
	Element uint16
}

// Mirror of tag.Info
type TagInfo struct {
	Tag  Tag
	VR   string
	Name string
	VM   string
}

// Parse tag value from string.
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
