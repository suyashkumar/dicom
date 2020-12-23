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

// Collects all unique raw VM values and checks that they were parsed correctly.
func TestVMInfo(t *testing.T) {
	type TestCase struct {
		RawVM  string
		Parsed VMInfo
	}

	var thisCase TestCase

	runTest := func(t *testing.T) {
		var expected VMInfo

		// We're going to hard code these values here so it's easier to check
		// correctness. If future Dicom standards significantly expand the set of
		// VMs in the spec, we may need to rethink this approach.
		switch thisCase.RawVM {
		case "1":
			expected = VMInfo{Minimum: 1, Maximum: 1, Step: 1}
		case "2":
			expected = VMInfo{Minimum: 2, Maximum: 2, Step: 1}
		case "3":
			expected = VMInfo{Minimum: 3, Maximum: 3, Step: 1}
		case "4":
			expected = VMInfo{Minimum: 4, Maximum: 4, Step: 1}
		case "6":
			expected = VMInfo{Minimum: 6, Maximum: 6, Step: 1}
		case "9":
			expected = VMInfo{Minimum: 9, Maximum: 9, Step: 1}
		case "16":
			expected = VMInfo{Minimum: 16, Maximum: 16, Step: 1}
		case "1-2":
			expected = VMInfo{Minimum: 1, Maximum: 2, Step: 1}
		case "1-3":
			expected = VMInfo{Minimum: 1, Maximum: 3, Step: 1}
		case "1-8":
			expected = VMInfo{Minimum: 1, Maximum: 8, Step: 1}
		case "1-32":
			expected = VMInfo{Minimum: 1, Maximum: 32, Step: 1}
		case "1-99":
			expected = VMInfo{Minimum: 1, Maximum: 99, Step: 1}
		case "1-n":
			expected = VMInfo{Minimum: 1, Maximum: -1, Step: 1}
		case "2-n":
			expected = VMInfo{Minimum: 2, Maximum: -1, Step: 1}
		case "3-n":
			expected = VMInfo{Minimum: 3, Maximum: -1, Step: 1}
		case "6-n":
			expected = VMInfo{Minimum: 6, Maximum: -1, Step: 1}
		case "2-2n":
			expected = VMInfo{Minimum: 2, Maximum: -1, Step: 2}
		case "3-3n":
			expected = VMInfo{Minimum: 3, Maximum: -1, Step: 3}
		default:
			t.Errorf("do not have matching cae for %v", thisCase.RawVM)
			t.FailNow()
		}

		if thisCase.Parsed.Maximum != expected.Maximum {
			t.Errorf(
				"expected max '%v' for vm '%v', got '%v'",
				expected.Maximum,
				thisCase.RawVM,
				thisCase.Parsed.Maximum,
			)
		}

		if thisCase.Parsed.Minimum != expected.Minimum {
			t.Errorf(
				"expected min '%v' for vm '%v', got '%v'",
				expected.Minimum,
				thisCase.RawVM,
				thisCase.Parsed.Minimum,
			)
		}

		if thisCase.Parsed.Step != expected.Step {
			t.Errorf(
				"expected step '%v' for vm '%v', got '%v'",
				expected.Step,
				thisCase.RawVM,
				thisCase.Parsed.Step,
			)
		}
	}

	// Collect only the unique VM values. We can use dict keys like a set.
	uniqueVMs := make(map[string]VMInfo)
	for _, thisInfo := range tagDict {
		uniqueVMs[thisInfo.VM] = thisInfo.VMInfo
	}

	// Range over the dict. Create and run a test case for each entry.
	for vmRaw, vmInfo := range uniqueVMs {
		thisCase = TestCase{
			RawVM:  vmRaw,
			Parsed: vmInfo,
		}

		t.Run(thisCase.RawVM, runTest)
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
