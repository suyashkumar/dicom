package dicom

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

type dictEntry struct {
	vr      string
	name    string
	vm      string
	version string
}

// Sets the dictionary for the Parser
func Dictionary(r io.Reader) func(*Parser) error {

	return func(p *Parser) error {

		reader := csv.NewReader(r)
		reader.Comma = '\t'  // tab separated file
		reader.Comment = '#' // comments start with #

		dictionary := make([][]*dictEntry, 0xffff+1)

		for {

			row, err := reader.Read()

			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}

			tag := strings.Split(strings.Trim(row[0], "()"), ",")

			group, err := strconv.ParseInt(tag[0], 16, 0)
			if err != nil {
				// TODO: Get this to work for ranges
				// return err
				continue
			}
			element, err := strconv.ParseInt(tag[1], 16, 0)
			if err != nil {
				// return err
				continue
			}

			if cap(dictionary[group]) == 0 {
				dictionary[group] = make([]*dictEntry, 0xffff+1)
			}

			dictionary[group][element] = &dictEntry{row[1], row[2], row[3], row[4]}
		}

		p.Dictionary = dictionary
		return nil
	}

}

func (p *Parser) getDictEntry(group, element int) (*dictEntry, error) {
	entry := p.Dictionary[group][element]
	if entry == nil {
		return nil, ErrTagNotFound
	}

	return entry, nil
}
