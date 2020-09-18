package dicomio

import (
	"encoding/binary"
	"io"
)

// Writer is a lower level encoder that takes abstracted input and writes it at the byte-level
type Writer interface {
	SetTransferSynax(bo binary.ByteOrder, implicit bool)
	WriteZeros(len int)
	WriteString(v string)
	WriteByte(v byte)
	WriteBytes(v []byte)
	WriteUInt16(v uint16)
	WriteUInt32(v uint32)
	WriteFloat32(v float32) error
	WriteFloat64(v float64) error
	GetTransferSyntax() (binary.ByteOrder, bool)
}

type writer struct {
	out      io.Writer
	bo       binary.ByteOrder
	implicit bool
}

// NewWriter creates and returns a Writer struct
func NewWriter(out io.Writer, bo binary.ByteOrder, implicit bool) Writer {
	return &writer{
		out:      out,
		bo:       bo,
		implicit: implicit,
	}
}

func (w *writer) SetTransferSynax(bo binary.ByteOrder, implicit bool) {
	w.bo = bo
	w.implicit = implicit
}

func (w *writer) GetTransferSyntax() (binary.ByteOrder, bool) {
	return w.bo, w.implicit
}

func (w *writer) WriteZeros(len int) {
	zeros := make([]byte, len)
	w.out.Write(zeros)
}

func (w *writer) WriteString(v string) {
	w.out.Write([]byte(v))
}

func (w *writer) WriteByte(v byte) {
	binary.Write(w.out, w.bo, &v)
}

func (w *writer) WriteBytes(v []byte) {
	w.out.Write(v)
}

func (w *writer) WriteUInt16(v uint16) {
	binary.Write(w.out, w.bo, &v)
}

func (w *writer) WriteUInt32(v uint32) {
	binary.Write(w.out, w.bo, &v)
}

func (w *writer) WriteFloat32(v float32) error {
	return binary.Write(w.out, w.bo, &v)
}

func (w *writer) WriteFloat64(v float64) error {
	return binary.Write(w.out, w.bo, &v)
}
