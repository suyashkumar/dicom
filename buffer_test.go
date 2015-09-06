package dicom

import (
	"testing"
)

func TestReadString(t *testing.T) {

	const expect string = "OsiriX"

	// OsiriX
	b := newDicomBuffer([]byte{0x4f, 0x73, 0x69, 0x72, 0x69, 0x58})
	l := uint32(b.Len())

	if s := b.readString(l); s != expect {
		t.Errorf("Incorrect string comparison %#x. Should be %#x.", s, expect)
	}

}

func TestReadStringWithNullBytes(t *testing.T) {

	const expect string = "OsiriX"

	// null byte (0x00)

	// OsiriX
	b := newDicomBuffer([]byte{0x00, 0x4f, 0x73, 0x69, 0x72, 0x69, 0x58, 0x00})
	l := uint32(b.Len())

	if s := b.readString(l); s != expect {
		t.Errorf("Incorrect string comparison %#x. Should be %#x.", s, expect)
	}

}

func TestReadStringWithZeroWidthCharacter(t *testing.T) {

	const expect string = "OsiriX"

	// zero-width character (0xE2, 0x80, 0x8B)

	// OsiriX
	b := newDicomBuffer([]byte{0x4f, 0x73, 0x69, 0x72, 0x69, 0x58, 0xE2, 0x80, 0x8B})
	l := uint32(b.Len())

	if s := b.readString(l); s != expect {
		t.Errorf("Incorrect string comparison %#x. Should be %#x.", s, expect)
	}

}
