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

		file, err := ioutil.ReadFile(path)

		if err != nil {
			panic(err)
		}

		data, err := dicom.Parse(file)

		if err != nil {
			fmt.Println(err)
		}

		for _, elem := range data.Elements {
			fmt.Printf("%+v\n", &elem)
		}

	}

}
