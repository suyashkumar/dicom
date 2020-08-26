package dicomio

import (
  "io"
  "binary"
)

// Writer is  what Encoder object used to be
type Writer interface {
  SetTransferSynax(bo binary.ByteOrder)
}

type writer struct {
  out io.Writer
  bo binary.ByteOrder
  implicit bool
}

// See https://github.com/suyashkumar/dicom/blob/6ffe547e2a080b3dcc1ce01946a7c7350d531bdc/dicomio/buffer.go#L80
func NewWriter(out io.Writer, bo binary.ByteOrder, implicit bool) *Writer {
  return &Writer{
    out: out,
    bo: bo,
    implicit: implicit,
  }
}

func (w *Writer) SetTransferSynax(bo binary.ByteOrder, implicit bool) {
  w.bo = bo
  w.implicit = implicit
}

// Retrieve the a []byte representation of what's contained in writer.out
func (w *Writer) Bytes() []byte {return ni}


// Low-level functions to write to writer.out
func (w *Writer) WriteZeros() {}
func (w *Writer) WriteZeros() {}
