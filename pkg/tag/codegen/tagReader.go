package main

import (
	"errors"
	"fmt"
	"io"
)

// Interface to implement for reading tags from a source.
type TagReader interface {
	// Name of reader for error messages.
	Name() string
	// Yield next tag or io.EOF if done.
	Next() (TagInfo, error)
	// Implements io.Closer() for closing underlying reader.
	Close() error
}

// Reads tags from multiple TagReader implementations until all readers are exhausted.
type MasterTagReader struct {
	tagReaders []TagReader
}

func (reader *MasterTagReader) Name() string {
	return "master tag reader"
}

func (reader *MasterTagReader) Next() (TagInfo, error) {
	// If there are no readers left to process, we are done.
	if len(reader.tagReaders) == 0 {
		return TagInfo{}, io.EOF
	}

	// The current reader will be the first in the slice.
	currentReader := reader.tagReaders[0]

	// Try to read the next tag from the current reader.
	info, err := currentReader.Next()

	// If this reader is complete.
	if errors.Is(err, io.EOF) {
		// close reader and remove from list.
		_ = currentReader.Close()
		reader.tagReaders = reader.tagReaders[1:]

		// Get the next tag.
		return reader.Next()
	} else if err != nil {
		// If this is some other error, return it.
		return TagInfo{}, fmt.Errorf(
			"error getting tag from reader %v: %w", reader, err,
		)
	}

	return info, nil
}

func (reader *MasterTagReader) Close() (err error) {
	// Defer all closes so if we panic, files still get closed.
	for _, thisReader := range reader.tagReaders {
		defer func(closer io.Closer) {
			thisErr := closer.Close()
			if thisErr != nil && err == nil {
				err = thisErr
			}
		}(thisReader)
	}

	return err
}

// Creates master TagReader we can use to read from all of our child readers.
func NewMasterTagReader(
	tagReaderCreators []func() (TagReader, error),
) (reader TagReader, err error) {
	masterReader := &MasterTagReader{
		tagReaders: make([]TagReader, 0, len(tagReaderCreators)),
	}

	// Close master reader and all sub-readers on error
	defer func() {
		if err != nil {
			masterReader.Close()
		}
	}()

	var thisReader TagReader
	for _, thisCreator := range tagReaderCreators {
		thisReader, err = thisCreator()
		if err != nil {
			return nil, err
		}

		masterReader.tagReaders = append(masterReader.tagReaders, thisReader)
	}

	return masterReader, nil
}
