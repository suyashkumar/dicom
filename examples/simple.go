package main

import (
	"fmt"
	"github.com/gillesdemey/go-dicom"
	"io/ioutil"
	"path/filepath"
)

func main() {

	// parse all DICOM files
	files, _ := filepath.Glob("examples/*.dcm")

	for _, path := range files {

		fmt.Printf("Parsing file %s\n", path)

		file, err := ioutil.ReadFile(path)

		if err != nil {
			panic(err)
		}

		data, err := dicom.Parse(file)

		if err != nil {
			fmt.Println(err)
		}

		patient, _ := data.LookupElement("PatientName")
		name := patient.Value
		fmt.Printf("Patient name: %s\n", name)
	}

}
