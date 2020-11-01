package dicomio

import (
	"encoding/binary"
	"io"
)

// Writer is a lower level encoder that manages writing out entities to an
// io.Reader.
type Writer struct {
	out      io.Writer
	bo       binary.ByteOrder
	implicit bool
}

// NewWriter initializes and returns a Writer.
func NewWriter(out io.Writer, bo binary.ByteOrder, implicit bool) Writer {
	return Writer{
		out:      out,
		bo:       bo,
		implicit: implicit,
	}
}

func (w *Writer) SetTransferSyntax(bo binary.ByteOrder, implicit bool) {
	w.bo = bo
	w.implicit = implicit
}

func (w *Writer) GetTransferSyntax() (binary.ByteOrder, bool) {
	return w.bo, w.implicit
}

func (w *Writer) WriteZeros(len int) error {
	zeros := make([]byte, len)
	_, err := w.out.Write(zeros)
	return err
}

func (w *Writer) WriteString(v string) error {
	_, err := w.out.Write([]byte(v))
	return err
}

func (w *Writer) WriteByte(v byte) error {
	return binary.Write(w.out, w.bo, &v)
}

func (w *Writer) WriteBytes(v []byte) error {
	_, err := w.out.Write(v)
	return err
}

func (w *Writer) WriteUInt16(v uint16) error {
	return binary.Write(w.out, w.bo, &v)
}

func (w *Writer) WriteUInt32(v uint32) error {
	return binary.Write(w.out, w.bo, &v)
}

func (w *Writer) WriteFloat32(v float32) error {
	return binary.Write(w.out, w.bo, &v)
}

func (w *Writer) WriteFloat64(v float64) error {
	return binary.Write(w.out, w.bo, &v)
}
