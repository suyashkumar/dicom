package main

import (
	"fmt"
	"github.com/gillesdemey/go-dicom"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	fh, err := os.Open("dicom.dic")
	defer fh.Close()

	dict := dicom.Dictionary(fh)

	parser, err := dicom.NewParser(dict)

	if err != nil {
		panic(err)
	}

	// parse all DICOM files
	files, _ := filepath.Glob("examples/*.dcm")

	for _, path := range files {

		file, err := ioutil.ReadFile(path)

		if err != nil {
			panic(err)
		}

		data, err := parser.Parse(file)

		if err != nil {
			fmt.Println(err)
		}

		elem, _ := data.LookupElement("PatientName")
		name := elem.Value
		fmt.Printf("Patient name: %s\n", name)

	}

}
