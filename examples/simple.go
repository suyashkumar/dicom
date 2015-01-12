package main

import (
	"fmt"
	"github.com/gillesdemey/go-dicom"
	"io/ioutil"
)

func main() {

	file, err := ioutil.ReadFile("examples/IM-0001-0001.dcm")

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
