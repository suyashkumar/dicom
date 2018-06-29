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
	"sync"
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
		for _, elem := range data.Elements {
			if elem.Tag == dicomtag.PixelData {
				data := elem.Value[0].(dicom.PixelDataInfo)
				if data.Encapsulated {
					var wg sync.WaitGroup
					for frameIndex, frame := range data.EncapsulatedFrames {
						wg.Add(1)
						go generateEncapsulatedImage(frame, frameIndex, &wg)
					}
					wg.Wait()
				} else {
					var wg sync.WaitGroup
					for frameIndex, frame := range data.NativeFrames {
						wg.Add(1)
						go generateNativeImage(frame, frameIndex, &wg)
					}
					wg.Wait()
				}
			}
		}
	}
	log.Println("Complete.")
}

func generateEncapsulatedImage(frame []byte, frameIndex int, wg *sync.WaitGroup) {
	defer wg.Done()
	path := fmt.Sprintf("image_%d.jpg", frameIndex) // TODO: figure out the image format
	ioutil.WriteFile(path, frame, 0644)
	fmt.Printf("%s: %d bytes\n", path, len(frame))
}

func generateNativeImage(frame []uint16, frameIndex int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := image.NewGray16(image.Rect(0, 0, 512, 512)) //TODO(suyash): dont assume 512 x 512
	for j := 0; j < len(frame); j++ {
		i.SetGray16(j % 512, j / 512, color.Gray16{Y: frame[j]})
	}
	name := fmt.Sprintf("image_%d.jpg", frameIndex)
	f, _ := os.Create(name)
	jpeg.Encode(f, i, &jpeg.Options{Quality: 100})
	fmt.Printf("%s written \n", name)
}
