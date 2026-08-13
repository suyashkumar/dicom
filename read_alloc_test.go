package dicom

import (
	"bytes"
	"encoding/binary"
	"runtime"
	"testing"
)

// A defined-length element must never allocate more than the reader can
// supply. Before the bound, readBytes/ReadString ran make([]byte, vl) from the
// declared length before io.ReadFull saw the bytes were absent, so a 20-byte
// file could allocate gigabytes (#349).
func TestDefinedLengthDoesNotOverAllocate(t *testing.T) {
	var b bytes.Buffer
	b.Write(make([]byte, 128))
	b.WriteString("DICM")
	// (0002,0000) UL group length 4, value 0 — minimal valid meta start is not
	// required; a crafted OB with a 4 GiB length in a tiny file is the point.
	binary.Write(&b, binary.LittleEndian, uint16(0x0009))
	binary.Write(&b, binary.LittleEndian, uint16(0x0001))
	b.WriteString("OB")
	binary.Write(&b, binary.LittleEndian, uint16(0)) // reserved
	binary.Write(&b, binary.LittleEndian, uint32(0xFFFFFFF0))
	raw := b.Bytes()

	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	p, err := NewParser(bytes.NewReader(raw), int64(len(raw)), nil,
		SkipPixelData(), AllowMissingMetaElementGroupLength())
	if err == nil {
		for {
			if _, e := p.Next(); e != nil {
				break
			}
		}
	}
	runtime.ReadMemStats(&after)

	if mb := float64(after.TotalAlloc-before.TotalAlloc) / (1 << 20); mb > 64 {
		t.Errorf("a %d-byte file allocated %.0f MB from a declared length", len(raw), mb)
	}
}
