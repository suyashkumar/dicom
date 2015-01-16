package dicom

import (
	"fmt"
	"testing"
)

var parser *Parser

func init() {
	parser, _ = NewParser()
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
