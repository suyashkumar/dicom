package dicom

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"
)

// Errors
var (
	ErrTagNotFound = errors.New("Could not find tag in dicom dictionary")
)

// Read a tab-separated DICOM dictionary file and find the specified field
//
// Tag   VR  Name      VM  Version
func lookupTag(tag string, field string) (string, error) {

	tag = strings.ToUpper(tag)

	column := 0

	switch field {
	case "Tag":
		column = 0
	case "VR":
		column = 1
	case "Name":
		column = 2
	case "VM":
		column = 3
	case "Version":
		column = 4
	}

	file, err := os.Open("dicom.dic")
	defer file.Close()

	if err != nil {
		return "", err
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'  // tab separated file
	reader.Comment = '#' // comments start with #

	// read until we find the row, or EOF
	for {

		row, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}

		if row[0] == tag {
			return row[column], nil
		}

	}

	if err != nil {
		return "", err
	}

	return "", ErrTagNotFound
}
