package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomtag"
)

var (
	printMetadata = flag.Bool("print-metadata", true, "Print image metadata")
	extractImages = flag.Bool("extract-images", false, "Extract images into separate files")
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Panic("dicomutil <dicomfile>")
	}
	path := flag.Arg(0)
	data, err := dicom.ReadDataSetFromFile(path, dicom.ReadOptions{DropPixelData: !*extractImages})
	if data == nil {
		log.Panic("Error reading %s: %v", path, err)
	}
	log.Printf("Error reading %s: %v", path, err)
	if *printMetadata {
		for _, elem := range data.Elements {
			fmt.Printf("%v\n", elem.String())
		}
	}
	if *extractImages {
		n := 0
		for _, elem := range data.Elements {
			if elem.Tag == dicomtag.PixelData {
				data := elem.Value[0].(dicom.PixelDataInfo)
				for _, frame := range data.Frames {
					path := fmt.Sprintf("image.%d.jpg", n) // TODO: figure out the image format
					n++
					ioutil.WriteFile(path, frame, 0644)
					fmt.Printf("%s: %d bytes\n", path, len(frame))
				}
			}
		}
	}
}
