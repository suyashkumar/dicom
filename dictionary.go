package dicom

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
)

const (
	dict_col_tag     = 0
	dict_col_vr      = 1
	dict_col_name    = 2
	dict_col_vm      = 3
	dict_col_version = 4
)

type dicomRow struct {
	tag     string
	vr      string
	name    string
	vm      string
	version string
}

var dictionary = make(map[string]*dicomRow)

func marshalFile() error {

	file, err := os.Open("dicom.dic")
	defer file.Close()

	if err != nil {
		return err
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
			return err
		}

		r := &dicomRow{
			tag:     row[0],
			vr:      row[1],
			name:    row[2],
			vm:      row[3],
			version: row[4],
		}

		dictionary[r.tag] = r
	}

	if err != nil {
		return err
	}

	return nil
}

// Read a tab-separated DICOM dictionary file and find the specified field
//
// Tag   VR  Name      VM  Version
func lookupTag(tag string) (*dicomRow, error) {

	if len(dictionary) == 0 {
		marshalFile()
	}

	tag = strings.ToUpper(tag)

	if dictionary[tag] != nil {
		return dictionary[tag], nil
	}

	return nil, ErrTagNotFound
}
