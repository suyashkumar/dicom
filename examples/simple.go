package main

import (
	"fmt"
	"github.com/gillesdemey/go-dicom"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	dict, err := os.Open("dicom.dic")
	defer dict.Close()
	if err != nil {
		panic(err)
	}

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

		for _, elem := range data.Elements {
			fmt.Printf("%+v\n", &elem)
		}

	}

}
