package fuzz

import (
	"github.com/suyashkumar/dicom"
)

func Fuzz(data []byte) int {
	p, _ := dicom.NewParserFromBytes(data, nil)
	_, _ = p.Parse(dicom.ParseOptions{})

	return 1
}
