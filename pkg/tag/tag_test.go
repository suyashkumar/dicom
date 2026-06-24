package tag

import (
	"fmt"
	"strings"
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
		{
			name:        "command group tag",
			tag:         CommandField,
			wantKeyword: "CommandField",
			wantVRs:     []string{"US"},
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
		testInfo := Info{
			Tag:  Tag{Group: 0x0063, Element: 0x0020},
			VRs:  []string{"UT"},
			Name: "TestTag",
			VM:   "1",
		}
		// Given a certain tag does not exist
		_, err := FindByName("TestTag")
		if err == nil {
			t.Fatalf("expected TestTag to not exist")
		}

		// When a new tag is registered
		err = Add(testInfo, false)
		if err != nil {
			t.Fatalf("Add(TestInfo, false) = %v, want: nil", err)
		}

		// Then the tag is now part of the tag collection
		_, err = FindByName("TestTag")
		if err != nil {
			t.Errorf("error when finding newly added tag (FindByName\"TestTag\"), got: %v, want: nil", err)
		}
		info, err := Find(testInfo.Tag)
		if err != nil {
			t.Fatalf("unexpected error in Find(\"TestTag\"). got err: %v, want err: nil", err)
		}
		if diff := cmp.Diff(testInfo, info); diff != "" {
			t.Errorf("unexpected diff in retrieved tag info after Add: %v: \n", diff)
		}
	})
	t.Run("force overwrite existing tag", func(t *testing.T) {
		// setup a test tag
		testInfo := Info{
			Tag:  Tag{Group: 0x0073, Element: 0x0021},
			VRs:  []string{"UT"},
			Name: "TestTag",
			VM:   "1",
		}
		err := Add(testInfo, false)
		if err != nil {
			t.Fatalf("Add(testInfo, false) = %v, want: nil for unknown tag", err)
		}

		// now change the info
		testInfo.VRs = []string{"LO"}
		// and try to overwrite it with force
		err = Add(testInfo, true)
		if err != nil {
			t.Fatalf("Add(testInfo, true) = %v, want: nil. Add with force = true should never return an error", err)
		}

		// validate that the tag has been overwritten
		info, err := Find(testInfo.Tag)
		if err != nil {
			t.Fatalf("unexpected err on Find, got: %v, want: nil", err)
		}
		if info.VRs[0] != "LO" {
			t.Errorf("unexpected VR on updated tag, got: %v, want: LO", info.VRs[0])
		}
	})
	t.Run("fail to overwrite existing tag", func(t *testing.T) {
		// setup a test tag
		testInfo := Info{
			Tag:  Tag{Group: 0x0083, Element: 0x0031},
			VRs:  []string{"UT"},
			Name: "TestTag",
			VM:   "1",
		}
		err := Add(testInfo, false)
		if err != nil {
			t.Fatalf("Add(testInfo, false) = error(%v), expected nil for unknown tag", err)
		}

		// now change the info
		testInfo.VRs = []string{"LO"}
		// and try to overwrite it with force
		err = Add(testInfo, false)
		if err == nil || !strings.Contains(err.Error(), "tag already exists") {
			t.Fatalf("Add(testInfo, true) = %v, want: error(\"tag already exists\"). Add with force = false should return an error when trying to overwrite an existing tag", err)
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
