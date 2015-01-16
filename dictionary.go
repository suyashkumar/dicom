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

func (p *Parser) loadDictionary(dictionary io.Reader) error {
	reader := csv.NewReader(dictionary)
	reader.Comma = '\t'  // tab separated file
	reader.Comment = '#' // comments start with #

	p.dict = make([][]*dictEntry, 0xffff+1)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		group, element, err := splitTag(row[0])
		if err != nil {
			// TODO: fix group ranges (6000-60FF,0803) causing error
			continue
		}

		if cap(p.dict[group]) == 0 {

			p.dict[group] = make([]*dictEntry, 0xffff+1)
		}

		p.dict[group][element] = &dictEntry{row[1], row[2], row[3], row[4]}
	}
	return nil
}

func (p *Parser) getDictEntry(group, element int) (*dictEntry, error) {
	entry := p.dict[group][element]
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
