package dicomio

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"testing"
)

func TestReadInt16(t *testing.T) {
	cases := []struct {
		val int16
		bo  binary.ByteOrder
	}{
		{0, binary.LittleEndian},
		{12345, binary.LittleEndian},
		{-1, binary.LittleEndian},
		{-12345, binary.LittleEndian},
		{0, binary.BigEndian},
		{12345, binary.BigEndian},
		{-1, binary.BigEndian},
		{-12345, binary.BigEndian},
	}

	for _, tc := range cases {
		buf := &bytes.Buffer{}
		if err := binary.Write(buf, tc.bo, tc.val); err != nil {
			t.Fatalf("setup error: %v", err)
		}

		reader := NewReader(bufio.NewReader(buf), tc.bo, int64(buf.Len()))
		got, err := reader.ReadInt16()
		if err != nil {
			t.Errorf("ReadInt16 error for val %d: %v", tc.val, err)
		}
		if got != tc.val {
			t.Errorf("ReadInt16 for val %d under %v: got %d, want %d", tc.val, tc.bo, got, tc.val)
		}
	}
}

func TestReadInt32(t *testing.T) {
	cases := []struct {
		val int32
		bo  binary.ByteOrder
	}{
		{0, binary.LittleEndian},
		{123456789, binary.LittleEndian},
		{-1, binary.LittleEndian},
		{-123456789, binary.LittleEndian},
		{0, binary.BigEndian},
		{123456789, binary.BigEndian},
		{-1, binary.BigEndian},
		{-123456789, binary.BigEndian},
	}

	for _, tc := range cases {
		buf := &bytes.Buffer{}
		if err := binary.Write(buf, tc.bo, tc.val); err != nil {
			t.Fatalf("setup error: %v", err)
		}

		reader := NewReader(bufio.NewReader(buf), tc.bo, int64(buf.Len()))
		got, err := reader.ReadInt32()
		if err != nil {
			t.Errorf("ReadInt32 error for val %d: %v", tc.val, err)
		}
		if got != tc.val {
			t.Errorf("ReadInt32 for val %d under %v: got %d, want %d", tc.val, tc.bo, got, tc.val)
		}
	}
}
