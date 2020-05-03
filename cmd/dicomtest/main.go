// Really basic sanity check program
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

var (
	filepath = flag.String("path", "", "path")
)

func main() {
	log.Println("start")
	flag.Parse()
	log.Println(*filepath)
	if len(*filepath) > 0 {
		log.Println("start")
		f, err := os.Open(*filepath)
		defer f.Close()
		if err != nil {
			log.Println("err")
			return
		}
		info, err := f.Stat()
		if err != nil {
			log.Println("err reading", err)
			return
		}
		p, err := dicom.NewParser(f, info.Size())
		if err != nil {
			log.Println("err creating parser", err)
			return
		}
		ds, err := p.Parse()
		if err != nil {
			log.Println("err parsing", err)
			return
		}

		for _, elem := range ds.Elements {
			if elem.Tag != tag.PixelData {
				log.Println(elem.Tag)
				log.Println(elem.ValueLength)
				log.Println(elem.Value)
			} else {
				imageInfo := elem.Value.GetValue().(dicom.PixelDataInfo)
				for i, frame := range imageInfo.Frames {
					err := ioutil.WriteFile(fmt.Sprintf("image_%d.jpg", i), frame.EncapsulatedData.Data,
						0644)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}
