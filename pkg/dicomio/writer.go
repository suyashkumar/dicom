package dicomio

import (
  "io"
  "binary"
)

// Writer is  what Encoder object used to be
type Writer interface {
}

type writer struct {
  out io.Writer
  bo binary.ByteOrder
}

// See https://github.com/suyashkumar/dicom/blob/6ffe547e2a080b3dcc1ce01946a7c7350d531bdc/dicomio/buffer.go#L80
func NewWriter(out io.Writer, bo binary.ByteOrder, ...) *Writer {
  return &Writer{
    out: out,
    bo: bo,
  }
}
