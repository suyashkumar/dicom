package dicomio

import (
  "io"
  "encoding/binary"
  "bytes"
)

type Writer interface {
  SetTransferSynax(bo binary.ByteOrder, implicit bool)
  Bytes() []byte
  WriteZeros(len int)
  WriteString(v string)
  WriteByte(v byte)
  WriteBytes(v []byte)
  WriteUInt16(v uint16)
  WriteUInt32(v uint32)
  GetTransferSyntax() (binary.ByteOrder, bool)
}

type writer struct {
  out io.Writer
  bo binary.ByteOrder
  implicit bool
}

// See https://github.com/suyashkumar/dicom/blob/6ffe547e2a080b3dcc1ce01946a7c7350d531bdc/dicomio/buffer.go#L80
func NewWriter(out io.Writer, bo binary.ByteOrder, implicit bool) Writer {
  return &writer{
    out: out,
    bo: bo,
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

// Retrieve the a []byte representation of what's contained in writer.out
// ONLY works when out is &bytes.Buffer{}
func (w *writer) Bytes() []byte {
  return w.out.(*bytes.Buffer).Bytes()
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
