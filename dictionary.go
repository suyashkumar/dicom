package dicom

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

const (
	dict_col_tag = 0
	dict_col_vr = 1
	dict_col_name = 2
	dict_col_vm = 3
	dict_col_version = 4
)

// Read a tab-separated DICOM dictionary file and find the specified field
//
// Tag   VR  Name      VM  Version
func lookupTag(tag string, field string) (string, error) {

	tag = strings.ToUpper(tag)

	var column int

	switch field {
	case "Tag":
		column = dict_col_tag
	case "VR":
		column = dict_col_vr
	case "Name":
		column = dict_col_name
	case "VM":
		column = dict_col_vm
	case "Version":
		column = dict_col_version
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
