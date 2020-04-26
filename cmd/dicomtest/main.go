package main

import (
	"flag"
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
		info, err := f.Stat()
		if err != nil {
			log.Println("err reading", err)
		}
		p, err := dicom.NewParser(f, info.Size())
		if err != nil {
			log.Println("err creating parser", err)
		}
		ds, err := p.Parse()
		if err != nil {
			log.Println("err parsing", err)
		}
		log.Println(len(ds.Elements))
		for _, elem := range ds.Elements {
			if elem.Tag != tag.PixelData {
				log.Println(elem.Tag)
				log.Println(elem.ValueLength)
				log.Println(elem.Value)
			}
		}
	}
}
