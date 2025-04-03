package dicomio

import (
	"bufio"
	"encoding/binary"
	"strings"
	"testing"
)

func TestReader_PeekAtMost(t *testing.T) {
	buf := bufio.NewReaderSize(strings.NewReader("abcdefghijklmnop"), 16)
	r := NewReader(buf, binary.LittleEndian, 0)

	if s, err := r.PeekAtMost(3); string(s) != "abc" || err != nil {
		t.Fatalf("want %q got %q, err=%v", "abc", string(s), err)
	}

	if s, err := r.PeekAtMost(0); len(s) != 0 || err != nil {
		t.Fatalf("want %d got %d, err=%v", 0, len(s), err)
	}

	if _, err := r.PeekAtMost(-1); err != bufio.ErrNegativeCount {
		t.Fatalf("want ErrNegativeCount got %v", err)
	}

	if s, err := r.PeekAtMost(32); string(s) != "abcdefghijklmnop" || err != nil {
		t.Fatalf("want %q,  without error, got %q, err=%v", "abcdefghijklmnop", string(s), err)
	}

	n, err := r.in.Read(make([]byte, 16))
	if n != 16 || err != nil {
		t.Fatalf("want 16 bytes, without error, got %d bytes, err=%v", n, err)
	}

	if s, err := r.PeekAtMost(1); len(s) == 0 && err == nil {
		t.Fatalf("want %d got %d, err=%v", 1, len(s), err)
	}

}
