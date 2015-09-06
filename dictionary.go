package dicom

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type dictEntry struct {
	tag     string
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
				continue // we don't support groups yet
			}

			if cap(dictionary[group]) == 0 {
				dictionary[group] = make([]*dictEntry, 0xffff+1)
			}

			dictionary[group][element] = &dictEntry{
				row[0],
				strings.ToUpper(row[1]),
				row[2],
				row[3],
				row[4],
			}
		}

		p.dictionary = dictionary
		return nil
	}

}

func (p *Parser) getDictEntry(group, element uint16) (*dictEntry, error) {

	var entry *dictEntry

	tag := fmt.Sprintf("(%s,%s)", group, element)

	// does the entry exist?
	exists := p.dictionary[group] != nil && p.dictionary[group][element] != nil

	if exists {
		entry = p.dictionary[group][element]
	}

	if !exists {
		// (0000-u-ffff,0000)	UL	GenericGroupLength	1	GENERIC
		if group%2 == 0 && element == 0x0000 {
			entry = &dictEntry{tag, "UL", "GenericGroupLength", "1", "GENERIC"}
		}
	}

	// nope, still nothing
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
