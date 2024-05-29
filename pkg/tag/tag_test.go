package tag

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name            string
		tag             Tag
		expectedKeyword string
		expectedVRs     []string
		expectedRetired bool
	}{
		{
			name:            "basic",
			tag:             Tag{0x0010, 0x0010},
			expectedKeyword: "PatientName",
			expectedVRs:     []string{"PN"},
		},
		{
			name:            "multiple vrs allowed",
			tag:             Tag{0x7FE0, 0x0010},
			expectedKeyword: "PixelData",
			expectedVRs:     []string{"OW", "OB"},
		},
		{
			name:            "NA tag",
			tag:             Tag{0xfffe, 0xe000},
			expectedKeyword: "Item",
			expectedVRs:     []string{"NA"},
		},
		{
			name:            "retired tag",
			tag:             Tag{0x7f00, 0x0011},
			expectedKeyword: "VariableNextDataGroup",
			expectedVRs:     []string{"US"},
			expectedRetired: true,
		},
		{
			name:            "file meta tag (group 0002)",
			tag:             Tag{0x0002, 0x0010},
			expectedKeyword: "TransferSyntaxUID",
			expectedVRs:     []string{"UI"},
		},
		{
			name:            "DICOMDIR tag (group 0004)",
			tag:             Tag{0x0004, 0x1130},
			expectedKeyword: "FileSetID",
			expectedVRs:     []string{"CS"},
		},
		{
			name:            "tag definition constant",
			tag:             CodeValue,
			expectedKeyword: "CodeValue",
			expectedVRs:     []string{"SH"},
		},
	}
	for _, test := range tests {
		t.Run(test.expectedKeyword, func(t *testing.T) {
			elem, err := Find(test.tag)
			if err != nil {
				t.Fatalf("Find returned unecpected error: %v", err)
			}
			if elem.Keyword != test.expectedKeyword {
				t.Errorf("Unexpected keyword, want %s, got %s", test.expectedKeyword, elem.Keyword)
			}
			if !cmp.Equal(elem.VRs, test.expectedVRs, cmpopts.SortSlices(func(x, y string) bool { return x < y })) {
				t.Errorf("Unexpected VRs list, want %v, got %v", test.expectedVRs, elem.VRs)
			}
			if elem.Retired != test.expectedRetired {
				t.Errorf("Unexpected Retired value, want %v, got %v", test.expectedRetired, elem.Retired)
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
