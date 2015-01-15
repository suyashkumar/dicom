package dicom

import (
  "fmt"
  "testing"
)

func init() {

}

func BenchmarkFindMetaGroupLengthTag(b *testing.B) {
  for i := 0; i < b.N; i++ {

    _, err := lookupTag("(0002,0000)", "Name")

    if err != nil {
      fmt.Println(err)
    }

  }
}

func BenchmarkFindPixelDataTag(b *testing.B) {
  for i := 0; i < b.N; i++ {

    _, err := lookupTag("(7FE0,0010)", "Name")

    if err != nil {
      fmt.Println(err)
    }

  }
}