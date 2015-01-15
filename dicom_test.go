package dicom

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var files [][]byte

func init() {
	f, err := filepath.Glob("examples/*.dcm")
	if err != nil {
		fmt.Println("failed to glob files")
		panic(err)
	}

	for _, path := range f {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("failed to read file")
			panic(err)
		}
		files = append(files, file)

	}
}

func BenchmarkParseSingle(b *testing.B) {
	dict, err := os.Open("dicom.dic")
	defer dict.Close()
	if err != nil {
		panic(err)
	}
	parser, err := NewParser(dict)
	for i := 0; i < b.N; i++ {
		_, err := parser.Parse(files[0])
		if err != nil {
			fmt.Println("failed to parse dicom file")
			panic(err)
		}
	}
}

func BenchmarkParseMultiple(b *testing.B) {
	dict, err := os.Open("dicom.dic")
	defer dict.Close()
	if err != nil {
		panic(err)
	}
	parser, err := NewParser(dict)
	for i := 0; i < b.N; i++ {
		for _, file := range files {
			_, err := parser.Parse(file)
			if err != nil {
				fmt.Println("failed to parse dicom file")
				panic(err)
			}
		}
	}
}
