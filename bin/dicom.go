package main

import (
	"flag"
	"fmt"
	"github.com/gillesdemey/go-dicom"
	"io/ioutil"
)

var (
	file = flag.String("file", "IM-0001-0001.dcm", "the DICOM file you want to parse")
)

func init() {
	flag.Parse()
}

func main() {

	parser, err := dicom.NewParser()
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(*file)

	if err != nil {
		panic(err)
	}

	data, err := parser.Parse(bytes)

	if err != nil {
		fmt.Println(err)
	}

	elem, _ := data.LookupElement("PatientName")
	name := elem.Value
	fmt.Printf("Patient name: %s\n", name)

}
