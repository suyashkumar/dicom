package dicomio

import (
	"encoding/binary"
	"io"
)

type Reader interface {
	io.Reader
	ReadN(n uint32) ([]byte, error)
	ReadUInt16() (uint16, error)
	ReadUInt32() (uint32, error)
	ReadInt16() (int16, error)
	ReadInt32() (int32, error)
	ReadString(n uint32) (string, error)
	Skip(n uint) error
	PushLimit(n int)
	PopLimit()
	IsLimitExhausted() bool
}

type reader struct {
	in         io.Reader
	bo         binary.ByteOrder
	limit      int64
	bytesRead  int64
	limitStack []int64
}

func NewReader(in io.Reader, bo binary.ByteOrder, limit int64) (Reader, error) {
	return &reader{
		in:        in,
		bo:        bo,
		limit:     limit,
		bytesRead: 0,
	}, nil
}

// TODO: Implement the rest of these interface functions
func (r *reader) ReadN(n uint32) ([]byte, error) {
	return nil, nil
}

func (r *reader) ReadUInt16() (uint16, error) {
	var out uint16
	err := binary.Read(r, r.bo, &out)
	return out, err
}

func (r *reader) ReadUInt32() (uint32, error)         { return 0, nil }
func (r *reader) ReadInt16() (int16, error)           { return 0, nil }
func (r *reader) ReadInt32() (int32, error)           { return 0, nil }
func (r *reader) ReadString(n uint32) (string, error) { return "", nil }
func (r *reader) Skip(n uint) error                   { return nil }
func (r *reader) Read(p []byte) (n int, err error)    { return r.in.Read(p) }
func (r *reader) PushLimit(n int)                     {}
func (r *reader) PopLimit()                           {}
func (r *reader) IsLimitExhausted() bool              { return false }
