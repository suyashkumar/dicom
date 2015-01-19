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

			group, element, err := splitTag(row[0])

			if err != nil {
				// return err
				continue
			}

			if cap(dictionary[group]) == 0 {
				dictionary[group] = make([]*dictEntry, 0xffff+1)
			}

			dictionary[group][element] = &dictEntry{row[1], row[2], row[3], row[4]}
		}

		p.dictionary = dictionary
		return nil
	}

}

func (p *Parser) getDictEntry(group, element int) (*dictEntry, error) {
	entry := p.dictionary[group][element]
	if entry == nil {
		return nil, ErrTagNotFound
	}

	return entry, nil
}

// Split a tag into a group and element, represented as a hex value
// TODO: support group ranges (6000-60FF,0803)
func splitTag(tag string) (int64, int64, error) {

	parts := strings.Split(strings.Trim(tag, "()"), ",")

	group, err := strconv.ParseInt(parts[0], 16, 0)
	if err != nil {
		return 0, 0, err
	}
	elem, err := strconv.ParseInt(parts[1], 16, 0)
	if err != nil {
		return 0, 0, err
	}

	return group, elem, nil
}
