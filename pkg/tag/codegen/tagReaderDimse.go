package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

const dimseCSVPath = "./codegen/dimse.csv"

const dimseTagRecordIndex = 0
const dimseVRRecordIndex = 1
const dimseKeywordRecordIndex = 2
const dimseVMRecordIndex = 3

// DimseReader reads group 0x0000 tags from the dimse.csv file.
type DimseReader struct {
	fileCloser  io.Closer
	csvReader   *csv.Reader
	recordIndex uint
}

// Name implements TagReader and returns "dimse reader".
func (reader *DimseReader) Name() string {
	return "dimse reader"
}

// Next implements TagReader and returns information about the next tag in the CSV file.
func (reader *DimseReader) Next() (TagInfo, error) {
	defer func() {
		reader.recordIndex++
	}()

	// Read next record.
	record, err := reader.csvReader.Read()
	if err != nil {
		return TagInfo{}, fmt.Errorf(
			"error decoding record %v: %w", reader.recordIndex, err,
		)
	}

	// Parse the info.
	tagRaw := record[dimseTagRecordIndex]
	thisTag, err := ParseTag(tagRaw)
	if err != nil {
		return TagInfo{}, fmt.Errorf(
			"error decoding tag %v for record %v: %w",
			tagRaw,
			reader.recordIndex,
			err,
		)
	}

	return TagInfo{
		Tag:  thisTag,
		VR:   record[dimseVRRecordIndex],
		Name: record[dimseKeywordRecordIndex],
		VM:   record[dimseVMRecordIndex],
	}, nil
}

// Close implements TagReader and closes our dimse.csv.
func (reader *DimseReader) Close() error {
	return reader.fileCloser.Close()
}

// NewDimseReader is a factory method for creating our DimseReader.
func NewDimseReader() (TagReader, error) {
	fileReader, err := os.Open(dimseCSVPath)
	if err != nil {
		return nil, fmt.Errorf(
			"error opening dimse csv at path '%v': %w", dimseCSVPath, err,
		)
	}

	csvReader := csv.NewReader(fileReader)
	csvReader.Comma = '\t'

	return &DimseReader{
		fileCloser: fileReader,
		csvReader:  csvReader,
	}, nil
}
