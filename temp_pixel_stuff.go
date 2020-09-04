package dicom

import (
   "math/rand"

   "github.com/suyashkumar/dicom/pkg/tag"
   "github.com/suyashkumar/dicom/pkg/frame"
)

func generateRandImage(cols int, rows int) []*Element {
      nf := frame.NativeFrame{
              Data:          randPixelData(cols, rows),
              Rows:          rows,
              Cols:          cols,
              BitsPerSample: 8,
      }
      f := frame.Frame{
              Encapsulated: false,
              NativeData:   nf,
      }
      frames := []frame.Frame{f}

      pd := PixelDataInfo{
              IsEncapsulated: false,
              Frames:         frames,
      }
      elems := []*Element{
              mustNewElement(tag.Tag{0x7FE0, 0x0010}, pd),            // PixelData
              mustNewElement(tag.Tag{0x0028, 0x0002}, []int{1}),     //SamplePerPixel
              mustNewElement(tag.Tag{0x0028, 0x0034}, []string{"01"}),          //PixelAspectRatio
              mustNewElement(tag.Tag{0x0028, 0x0103}, []int{0}),     //PixelRepresentation
              mustNewElement(tag.Tag{0x0028, 0x0010}, []int{rows}),  //Rows
              mustNewElement(tag.Tag{0x0028, 0x0011}, []int{cols}),  //Cols
              mustNewElement(tag.Tag{0x0028, 0x0004}, []string{"MONOCHROME1"}), //PhotometricInterpretation
              mustNewElement(tag.Tag{0x0028, 0x0100}, []int{16}),    //BitsAllocated
              mustNewElement(tag.Tag{0x0028, 0x0101}, []int{10}),    //BitsStored
              mustNewElement(tag.Tag{0x0028, 0x0102}, []int{9}),     //HighBit
              mustNewElement(tag.Tag{0x0028, 0x1050}, []string{"550"}),         //WindowCenter
              mustNewElement(tag.Tag{0x0028, 0x1051}, []string{"1024"}),        //WindowWidth
      }

      return elems
}

func randPixelData(cols int, rows int) [][]int {
        data := make([][]int, cols)

        for col := 0; col < cols; col++ {
                data[col] = make([]int, rows)
                for row := 0; row < rows; row++ {
                        data[col][row] = rand.Int()
                }
        }
        return data
}
