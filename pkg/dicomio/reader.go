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
	ReadString() (string, error)
	Skip(n uint) error
}

type reader struct {
	in        io.Reader
	bo        binary.ByteOrder
	limit     int64
	bytesRead int64
}
