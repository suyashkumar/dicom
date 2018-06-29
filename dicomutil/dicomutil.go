package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gradienthealth/go-dicom"
	"github.com/gradienthealth/go-dicom/dicomtag"
	"github.com/gradienthealth/go-dicom/dicomlog"
	"math"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

var (
	printMetadata = flag.Bool("print-metadata", true, "Print image metadata")
	extractImages = flag.Bool("extract-images", false, "Extract images into separate files")
	verbose = flag.Bool("verbose", false, "Activate high verbosity log operation")
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Panic("dicomutil <dicomfile>")
	}
	if *verbose {
		dicomlog.SetLevel(math.MaxInt32)
	}
	path := flag.Arg(0)
	data, err := dicom.ReadDataSetFromFile(path, dicom.ReadOptions{DropPixelData: !*extractImages})
	if data == nil {
		log.Panic("Error reading %s: %v", path, err)
	}
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
				if data.Encapsulated {
					for _, frame := range data.EncapsulatedFrames {
						path := fmt.Sprintf("image.%d.jpg", n) // TODO: figure out the image format
						n++
						ioutil.WriteFile(path, frame, 0644)
						fmt.Printf("%s: %d bytes\n", path, len(frame))
					}
				} else {
					for _, frame := range data.NativeFrames {
						i := image.NewGray16(image.Rect(0, 0, 512, 512))
						for j := 0; j < len(frame); j++ {
							i.SetGray16(j % 512, j / 512, color.Gray16{Y: frame[j]})
						}
						f, _ := os.Create(fmt.Sprintf("image_%d.jpg", n))
						jpeg.Encode(f, i, &jpeg.Options{Quality: 100})
						n++
					}
				}
			}
		}
	}
	log.Println("Complete.")
}


