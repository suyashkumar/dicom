package dicom

import (
	"fmt"
	"os"
	"testing"
)

var parser *Parser

func init() {
	dict, err := os.Open("dicom.dic")
	defer dict.Close()
	if err != nil {
		panic(err)
	}
	parser, err = NewParser(dict)
	if err != nil {
		panic(err)
	}
}

func BenchmarkFindMetaGroupLengthTag(b *testing.B) {
	for i := 0; i < b.N; i++ {

		_, err := parser.getDictEntry(2, 0)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func BenchmarkFindPixelDataTag(b *testing.B) {
	for i := 0; i < b.N; i++ {

		_, err := parser.getDictEntry(32736, 16)
		if err != nil {
			fmt.Println(err)
		}

	}
}
