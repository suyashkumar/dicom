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
	if elem.Name != "PixelData" || elem.VR != "OW" {
		t.Errorf("Wrong element name: %s", elem.Name)
	}
	elem, err = Find(Tag{0, 0x1002})
	if err != nil {
		t.Error(err)
	}
	if elem.Name != "EventTypeID" || elem.VR != "US" {
		t.Errorf("Wrong element name: %s", elem.Name)
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

func TestRegisterCustom(t *testing.T) {

	t.Run("Add new tag", func(t *testing.T) {
		// Given a certain tag does not exist
		_, err := FindByName("TestTag")
		if err == nil {
			t.Fatalf("expected TestTag to not exist")
		}

		// When a new tag is registered
		TagTestTag := Tag{Group: 0x0063, Element: 0x0020}
		RegisterCustom(Info{
			Tag:  TagTestTag,
			VR:   "UT",
			Name: "TestTag",
			VM:   "1",
		})

		// Then the tag is now part of the tag collection
		_, err = FindByName("TestTag")
		if err != nil {
			t.Fatalf("expected TestTag to be accessible with FindByName")
		}
		info, err := Find(TagTestTag)
		if err != nil {
			t.Fatalf("expected TestTag to be accessible with Find")
		}
		if info.VR != "UT" ||
			info.Name != "TestTag" ||
			info.VM != "1" {
			t.Fatal("info of new registered tag is wrong")
		}
	})
	t.Run("override existing tag", func(t *testing.T) {
		// Given a tag already exists
		TagTestTag := Tag{Group: 0x0010, Element: 0x0010} // this is the PatientName tag
		info, err := Find(TagTestTag)
		if err != nil {
			t.Fatalf("expected TestTag to be accessible with Find")
		}
		if info.VR != "PN" {
			t.Fatal("expected PatientName VR is originally PN")
		}

		// When the tag is registered with different content
		RegisterCustom(Info{
			Tag:  TagTestTag,
			VR:   "LO", // originally this is PN
			Name: "PatientName",
			VM:   "1",
		})

		// Then the tag information is overridden
		info, err = Find(TagTestTag)
		if err != nil {
			t.Fatalf("expected TestTag to be accessible with Find")
		}
		if info.VR != "LO" {
			t.Fatal("expected the VR to have changed to LO")
		}
	})

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
