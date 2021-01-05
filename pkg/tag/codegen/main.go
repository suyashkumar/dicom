package main

import (
	"errors"
	"io"
	"log"
)

// tagReaderFactories are the factory functions that generate our tag readers.
var tagReaderFactories = []func() (TagReader, error){
	NewDimseReader,
	NewInnoliticsReader,
}

// codegenWriterFactories are the factory functions that generate our tag writers.
var codegenWriterFactories = []func() (CodeWriter, error){
	NewAttributesCodeWriter,
	NewTagDictCodeWriter,
}

// Our main program.
func main() {
	// Create reader object which can read tag info from multiple sources in sequence.
	log.Println("creating tag readers")
	tagReader, err := NewMasterTagReader(tagReaderFactories)
	if err != nil {
		log.Fatal("error creating tag readers:", err)
	}

	// Close the reader on exit.
	defer tagReader.Close()

	// Create writer object which can write codegen to multiple targets.
	log.Println("creating code writers")
	codeWriter, err := NewMasterCodeWriter(codegenWriterFactories)
	if err != nil {
		log.Fatalln("error creating code writers:", err)
	}

	// Close the writer on exit.
	defer codeWriter.Close()

	// Write any leading codegen that needs to come before tag information.
	log.Println("starting codegen")
	err = codeWriter.WriteLeading()
	if err != nil {
		log.Fatalln("error writing leading code:", err)
	}

	// Iterate over tag info reader and write codegen for each tag info object read
	var thisInfo TagInfo

	// We are going to report the number of tags generated at the end.
	var tagCount uint
	for tagCount = 0; true; tagCount++ {
		// Read the next tag from the reader.
		thisInfo, err = tagReader.Next()
		if errors.Is(err, io.EOF) {
			// On EOF error, we are done.
			break
		} else if err != nil {
			// On any other errors, report it and exit as fatal.
			log.Fatalln("error reading tag:", err)
		}

		// Write the tag.
		err = codeWriter.WriteTag(thisInfo)
		if err != nil {
			// If there are any errors on writing the tag, exit as fatal.
			log.Fatalf("error writing tag %+v: %v\n", thisInfo.Tag, err)
		}
	}

	// Write any codegen that needs to come after tag info.
	err = codeWriter.WriteTrailing()
	if err != nil {
		log.Fatalln("error writing trailing code:", err)
	}

	// Report the number of written tags.
	log.Println("codegen complete. tags processed:", tagCount)

	// Done
}
