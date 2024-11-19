package tag

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name        string
		tag         Tag
		wantKeyword string
		wantVRs     []string
		wantRetired bool
	}{
		{
			name:        "basic",
			tag:         Tag{0x0010, 0x0010},
			wantKeyword: "PatientName",
			wantVRs:     []string{"PN"},
		},
		{
			name:        "multiple vrs allowed",
			tag:         Tag{0x7FE0, 0x0010},
			wantKeyword: "PixelData",
			wantVRs:     []string{"OW", "OB"},
		},
		{
			name:        "NA tag",
			tag:         Tag{0xfffe, 0xe000},
			wantKeyword: "Item",
			wantVRs:     []string{"NA"},
		},
		{
			name:        "retired tag",
			tag:         Tag{0x7f00, 0x0011},
			wantKeyword: "VariableNextDataGroup",
			wantVRs:     []string{"US"},
			wantRetired: true,
		},
		{
			name:        "file meta tag (group 0002)",
			tag:         Tag{0x0002, 0x0010},
			wantKeyword: "TransferSyntaxUID",
			wantVRs:     []string{"UI"},
		},
		{
			name:        "DICOMDIR tag (group 0004)",
			tag:         Tag{0x0004, 0x1130},
			wantKeyword: "FileSetID",
			wantVRs:     []string{"CS"},
		},
		{
			name:        "tag definition constant",
			tag:         CodeValue,
			wantKeyword: "CodeValue",
			wantVRs:     []string{"SH"},
		},
	}
	for _, test := range tests {
		t.Run(test.wantKeyword, func(t *testing.T) {
			elem, err := Find(test.tag)
			if err != nil {
				t.Fatalf("Find returned unexpected error: %v", err)
			}
			if elem.Keyword != test.wantKeyword {
				t.Errorf("Unexpected keyword, want %s, got %s", test.wantKeyword, elem.Keyword)
			}
			if !cmp.Equal(elem.VRs, test.wantVRs, cmpopts.SortSlices(func(x, y string) bool { return x < y })) {
				t.Errorf("Unexpected VRs list, want %v, got %v", test.wantVRs, elem.VRs)
			}
			if elem.Retired != test.wantRetired {
				t.Errorf("Unexpected Retired value, want %v, got %v", test.wantRetired, elem.Retired)
			}
		})
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
