package dicomio

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"strings"
	"testing"
)

func TestReader_PeekAtMost(t *testing.T) {
	tests := []struct {
		name      string
		data      string
		peekCount int
		wantData  string
		wantErr   error
	}{
		{
			name:      "peek_valid_count_within_buffer",
			data:      "abcdefghijklmnop",
			peekCount: 3,
			wantData:  "abc",
			wantErr:   nil,
		},
		{
			name:      "peek_zero_bytes",
			data:      "abcdefghijklmnop",
			peekCount: 0,
			wantData:  "",
			wantErr:   nil,
		},
		{
			name:      "peek_negative_count_returns_error",
			data:      "abcdefghijklmnop",
			peekCount: -1,
			wantData:  "",
			wantErr:   bufio.ErrNegativeCount,
		},
		{
			name:      "peek_count_larger_than_available_data",
			data:      "abcdefghijklmnop",
			peekCount: 32,
			wantData:  "abcdefghijklmnop",
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bufio.NewReaderSize(strings.NewReader(tt.data), 16)
			r := NewReader(buf, binary.LittleEndian, 0)

			data, err := r.PeekAtMost(tt.peekCount)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got error %v, want %v", err, tt.wantErr)
			}

			if string(data) != tt.wantData {
				t.Errorf("got data %q, want %q", string(data), tt.wantData)
			}
		})
	}
}

func TestReader_PeekAtMost_AfterConsumingAllData(t *testing.T) {
	buf := bufio.NewReaderSize(strings.NewReader("abcdefghijklmnop"), 16)
	r := NewReader(buf, binary.LittleEndian, 0)

	// Consume all data
	n, err := r.in.Read(make([]byte, 16))
	if n != 16 || err != nil {
		t.Fatalf("failed to consume all data: got %d bytes, err=%v", n, err)
	}
	// Now try to peek - should return empty slice with EOF error
	data, err := r.PeekAtMost(1)
	if len(data) != 0 || !errors.Is(err, io.EOF) {
		t.Errorf("got len=%d, err=%v, want len=0, err=EOF", len(data), err)
	}
}
