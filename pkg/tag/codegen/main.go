package main

import (
	"errors"
	"io"
	"log"
)

var tagReaderCreators = []func() (TagReader, error){
	NewDimseReader,
	NewInnoliticsReader,
}

var codegenWriterCreators = []func() (CodeWriter, error){
	NewAttributesCodeWriter,
	NewTagDictCodeWriter,
}

func main() {
	log.Println("creating tag readers")
	// Creates reader object which can read tag info from multiple sources in sequence.
	tagReader, err := NewMasterTagReader(tagReaderCreators)
	if err != nil {
		log.Fatal("error creating tag readers:", err)
	}
	defer tagReader.Close()

	log.Println("creating code writers")
	// Creates writer object which can write codegen to multiple targets.
	codeWriter, err := NewMasterCodegenWriter(codegenWriterCreators)
	if err != nil {
		log.Fatalln("error creating code writers:", err)
	}
	defer codeWriter.Close()

	log.Println("starting codegen")
	// Write any leading codegen that needs to come before tag information
	err = codeWriter.WriteLeading()
	if err != nil {
		log.Fatalln("error writing leading code:", err)
	}

	// Iterate over tag info reader and write codegen for each tag info object read
	var thisInfo TagInfo

	var i uint
	for i = 0; true; i++ {
		thisInfo, err = tagReader.Next()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Fatalln("error reading tag:", err)
		}

		err = codeWriter.WriteTag(thisInfo)
		if err != nil {
			log.Fatalf("error writing tag %+v: %v\n", thisInfo.Tag, err)
		}
	}
	log.Println("codegen complete. tags processed:", i)

	// Write any codegen that needs to come after tag info.
	err = codeWriter.WriteTrailing()
	if err != nil {
		log.Fatalln("error writing trailing code:", err)
	}

	// Done
}
