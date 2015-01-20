package dicom

import (
	"fmt"
	"testing"
)

var parser *Parser

func init() {
	parser, _ = NewParser()
}

func TestDefaultDictionary(t *testing.T) {
	if _, err := NewParser(); err != nil {
		t.Error(err)
	}
}

func TestGetDictEntry(t *testing.T) {
	elem, err := parser.getDictEntry(32736, 16)
	if err != nil {
		t.Error(err)
	}

	if elem.name != "PixelData" {
		t.Errorf("Wrong element name: %s", elem.name)
	}

	if elem.vr != "OX" {
		t.Errorf("Wrong element VR: %s", elem.vr)
	}

}

// TODO: add a test for correctly splitting ranges
func TestSplitTag(t *testing.T) {

	group, element, err := splitTag("(7FE0,0010)")
	if err != nil {
		t.Error(err)
	}

	if group != 0x7FE0 {
		t.Errorf("Error splitting tag. Wrong group: %#x", group)
	}

	if element != 0x0010 {
		t.Errorf("Error splitting tag. Wrong element: %#x", element)
	}

}

func BenchmarkFindMetaGroupLengthTag(b *testing.B) {
	for i := 0; i < b.N; i++ {

		if _, err := parser.getDictEntry(2, 0); err != nil {
			fmt.Println(err)
		}

	}
}

func BenchmarkFindPixelDataTag(b *testing.B) {
	for i := 0; i < b.N; i++ {

		if _, err := parser.getDictEntry(32736, 16); err != nil {
			fmt.Println(err)
		}

	}
}
