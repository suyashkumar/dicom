package dicomio

import (
  "io"
  "encoding/binary"
)

// Writer is  what Encoder object used to be
type Writer interface {
  SetTransferSynax(bo binary.ByteOrder, implicit bool)
  Bytes() []byte
  WriteZeros(len int)
  WriteString(str string)
  WriteBytes(bytes []byte)
  WriteUInt16(v uint16)
  WriteUInt32(v uint32)
  GetTransferSyntax() (binary.ByteOrder, bool)
  // TODO fill in other functions that Writer has
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
func (w *writer) Bytes() []byte {return nil}


// Low-level functions to write to writer.out
func (w *writer) WriteZeros(len int) {
  zeros := make([]byte, len)
  w.out.Write(zeros)
}

func (w *writer) WriteString(str string) {
  w.out.Write([]byte(str))
}

func (w *writer) WriteBytes(bytes []byte) {}

func (w *writer) WriteUInt16(v uint16) {
  binary.Write(w.out, w.bo, &v)
}

func (w *writer) WriteUInt32(v uint32) {
  binary.Write(w.out, w.bo, &v)
}
