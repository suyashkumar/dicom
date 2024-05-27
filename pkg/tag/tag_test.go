package tag

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	elem, err := Find(Tag{32736, 16})
	if err != nil {
		t.Error(err)
	}
	if elem.Keyword != "PixelData" || elem.VR != "OW" {
		t.Errorf("Wrong element name: %s", elem.Keyword)
	}
	elem, err = Find(Tag{0, 0x1002})
	if err != nil {
		t.Error(err)
	}
	if elem.Keyword != "EventTypeID" || elem.VR != "US" {
		t.Errorf("Wrong element name: %s", elem.Keyword)
	}

	elem, err = FindByName("TransferSyntaxUID")
	if err != nil {
		t.Error(err)
	}
	if (elem.Tag != Tag{2, 0x10}) {
		t.Errorf("Wrong element: %v", elem)
	}
}

// TODO: add a test for correctly splitting ranges
func TestSplitTag(t *testing.T) {
	tag, err := parseTag("(7FE0,0010)")
	if err != nil {
		t.Error(err)
	}
	if tag.Group != 0x7FE0 {
		t.Errorf("Error splitting tag. Wrong group: %#x", tag.Group)
	}
	if tag.Element != 0x0010 {
		t.Errorf("Error splitting tag. Wrong element: %#x", tag.Element)
	}
}

func TestUint32Conversion(t *testing.T) {
	var got, want uint32 = Tag{Group: 0x7FE0, Element: 0x0010}.Uint32(), 0x7FE00010
	if got != want {
		t.Errorf("Uint32() got unexpected value %x; want %x", got, want)
	}
}

func BenchmarkFindMetaGroupLengthTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := Find(Tag{2, 0}); err != nil {
			fmt.Println(err)
		}

	}
}

func BenchmarkFindPixelDataTag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := Find(Tag{32736, 16}); err != nil {
			fmt.Println(err)
		}

	}
}
