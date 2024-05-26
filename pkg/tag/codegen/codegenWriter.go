package main

import (
	"fmt"
	"io"
)

const packageName = "package tag\n\n"
const codegenWarning = "// Code generated from './pkg/tag/codegen'. " +
	"DO NOT EDIT.\n\n"

// GeneratedFileLeader contains the preamble needed for a generated .go file in the
// tag package.
const GeneratedFileLeader = packageName + codegenWarning

// CodeWriter is an interface for writing to a dicom tag codegen file.
type CodeWriter interface {
	// Name to use in error messages related to this writer.
	Name() string

	// WriteLeading writes the opening part of a file. Called once before calls to
	// WriteTag() are made.
	WriteLeading() error

	// WriteTag writes codegen for a given tag. It may be called many times between
	// WriteLeading and WriteTrailing. WriteTag should return io.EOF when all tags
	// are successfully written.
	WriteTag(info TagInfo) error

	// WriteTrailing writes all codegen after WriteTag calls are complete.
	WriteTrailing() error

	// Close closes any underlying open resources.
	Close() error
}

// MasterCodeWriter takes in multiple CodeWriter instances and exposes them as a single
// CodeWriter.
type MasterCodeWriter struct {
	codegenWriters []CodeWriter
}

// Name implements CodeWriter, and returns "master codegen writer".
func (writer *MasterCodeWriter) Name() string {
	return "master codegen writer"
}

// WriteLeading implements CodeWriter, and calls WriteLeading on all child code writers.
func (writer *MasterCodeWriter) WriteLeading() error {
	for _, thisWriter := range writer.codegenWriters {
		err := thisWriter.WriteLeading()
		if err != nil {
			return fmt.Errorf("error writing for %v: %w", thisWriter.Name(), err)
		}
	}

	return nil
}

// WriteTag implements CodeWriter, and calls WriteTag with info on all child code
// writers.
func (writer *MasterCodeWriter) WriteTag(info TagInfo) error {
	for _, thisWriter := range writer.codegenWriters {
		err := thisWriter.WriteTag(info)
		if err != nil {
			return fmt.Errorf("error writing for %v: %w", thisWriter.Name(), err)
		}
	}

	return nil
}

// WriteTrailing implements CodeWriter, and calls WriteTrailing and all child readers.
func (writer *MasterCodeWriter) WriteTrailing() error {
	for _, thisWriter := range writer.codegenWriters {
		err := thisWriter.WriteTrailing()
		if err != nil {
			return fmt.Errorf("error writing for %v: %w", thisWriter.Name(), err)
		}
	}

	return nil
}

// Close implements CodeWriter, and closes all child readers.
func (writer *MasterCodeWriter) Close() (err error) {
	for _, thisWriter := range writer.codegenWriters {
		defer func(closer io.Closer) {
			thisErr := closer.Close()
			if thisErr != nil && err == nil {
				err = thisErr
			}
		}(thisWriter)
	}

	return err
}

// NewMasterCodeWriter creates a MasterCodeWriter from a slice of functions that
// create new child writers for the MasterCodeWriter to handle.
func NewMasterCodeWriter(
	codegenWriterCreators []func() (CodeWriter, error),
) (writer CodeWriter, err error) {
	masterWriter := &MasterCodeWriter{
		codegenWriters: make([]CodeWriter, 0, len(codegenWriterCreators)),
	}

	// Close master writer (and all sub-writers) on error.
	defer func() {
		if err != nil {
			masterWriter.Close()
		}
	}()

	var thisWriter CodeWriter
	for _, thisCreator := range codegenWriterCreators {
		thisWriter, err = thisCreator()
		if err != nil {
			return nil, err
		}

		masterWriter.codegenWriters = append(masterWriter.codegenWriters, thisWriter)
	}

	return masterWriter, nil
}
