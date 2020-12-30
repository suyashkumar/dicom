package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const innoliticsJSONPath = "./dicom-standard/standard/attributes.json"

// Data model to unmarshal innolitics json.
type innoliticsTagInfo struct {
	Tag                 string `json:"tag"`
	Name                string `json:"name"`
	Keyword             string `json:"keyword"`
	ValueRepresentation string `json:"valueRepresentation"`
	ValueMultiplicity   string `json:"valueMultiplicity"`
	Retired             string `json:"retired"`
	ID                  string `json:"id"`
}

func infoFromInnolitics(tagSpec innoliticsTagInfo) (TagInfo, error) {
	// If this tag contains an "X" it is information about a general range of tags,
	// and should be skipped.
	//
	// Retired tags should also be skipped.
	if strings.ContainsAny(tagSpec.Tag, "X") || tagSpec.Retired == "Y" {
		return TagInfo{}, errSkipTag
	}

	// Parse the tag
	tag, err := ParseTag(tagSpec.Tag)
	if err != nil {
		return TagInfo{}, fmt.Errorf(
			"error parsing tag '%v': %w", tagSpec.Tag, err,
		)
	}

	// some VRs have multiple allowed VR values such as "OB or OW" or "US or SS". For
	// now we should use only one of them, but it may be worth in the future including
	// both in a slice and having a helper function for IsVR(vr string) for testing
	// whether a tag is a VR.
	//
	// The CSV from our previous method of generation seemed to favor the second value,
	// so that's what we will do here as well..
	vrList := strings.Split(tagSpec.ValueRepresentation, " or ")
	vr := vrList[len(vrList)-1]

	// Some vrs have values like that are ill-defined, like "See Note 2". This affects
	// tags like ItemDelimitationItem. The previous CSV we used defined these VR's as
	// 'na', so we will do the same here.
	if vr == "See Note 2" {
		vr = "na"
	}

	// By default, we will set the name equal to the Attribute Keyword
	name := tagSpec.Keyword
	// If there is no keyword, use the name, subbing out spaces for underscores
	if name == "" {
		name = strings.Replace(tagSpec.Name, " ", "_", -1)
	}

	// If there is still no name, skip the tag
	if name == "" {
		return TagInfo{}, errSkipTag
	}

	info := TagInfo{
		Tag:  tag,
		VR:   vr,
		Name: tagSpec.Keyword,
		VM:   tagSpec.ValueMultiplicity,
	}

	return info, nil
}

// Sentinel error to signal to main parsing loop that this tag should be skipped.
var errSkipTag = errors.New("skip tag")

type InnoliticsTagReader struct {
	// Closes underlying file reader
	fileCloser io.Closer

	// Json decoder we are fetching from.
	jsonDecoder *json.Decoder

	// Element index for error messages.
	elementIndex uint
}

func (reader *InnoliticsTagReader) Name() string {
	return "innolitics json reader"
}

func (reader *InnoliticsTagReader) Next() (TagInfo, error) {
	defer func() {
		reader.elementIndex++
	}()
	// If there is no more, return an io.EOF
	if !reader.jsonDecoder.More() {
		return TagInfo{}, io.EOF
	}

	// Get the next attribute.
	thisDecoded := innoliticsTagInfo{}
	err := reader.jsonDecoder.Decode(&thisDecoded)
	if err != nil {
		return TagInfo{}, fmt.Errorf(
			"error decoding attribute %v: %w", reader.elementIndex, err,
		)
	}

	// Convert the raw unmarshalled info to a TagInfo object.
	info, err := infoFromInnolitics(thisDecoded)
	if errors.Is(err, errSkipTag) {
		// If we get a skip tag error, then read the next tag.
		return reader.Next()
	} else if err != nil {
		return TagInfo{}, fmt.Errorf(
			"error parsing attribute %v: %w", reader.elementIndex, err,
		)
	}

	return info, nil
}

func (reader *InnoliticsTagReader) Close() error {
	return reader.fileCloser.Close()
}

// Creates a new tag reader that parses tags from the innolitics json
func NewInnoliticsReader() (TagReader, error) {
	// open json file.
	fileReader, err := os.Open(innoliticsJSONPath)
	if err != nil {
		return nil, fmt.Errorf(
			"error opening innolitics json at path '%v': %w. make sure git"+
				" submodules are fetched and initialized using command: "+
				" git submodule update --init --recursive",
			innoliticsJSONPath,
			err,
		)
	}

	// Get json decoder.
	decoder := json.NewDecoder(fileReader)

	// Get the first token '[' of our open list
	_, err = decoder.Token()
	if err != nil {
		// On error, close the file reader.
		defer fileReader.Close()
		return nil, fmt.Errorf(
			"error starting innolitics attribute list decode: %w", err,
		)
	}

	// return reader
	return &InnoliticsTagReader{
		fileCloser:  fileReader,
		jsonDecoder: decoder,
	}, nil
}
