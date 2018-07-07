package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
	"sync"

	"github.com/gradienthealth/go-dicom"
	"github.com/gradienthealth/go-dicom/dicomlog"
	"github.com/gradienthealth/go-dicom/dicomtag"
)

var (
	printMetadata = flag.Bool("print-metadata", true, "Print image metadata")
	extractImages = flag.Bool("extract-images", false, "Extract images into separate files")
	verbose       = flag.Bool("verbose", false, "Activate high verbosity log operation")
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
	p, err := dicom.NewParserFromFile(path, nil)
	parsedData, err := p.Parse(dicom.ParseOptions{DropPixelData: !*extractImages})
	if parsedData == nil {
		log.Panic("Error reading %s: %v", path, err)
	}
	if *printMetadata {
		for _, elem := range parsedData.Elements {
			fmt.Printf("%v\n", elem.String())
		}
	}
	if *extractImages {
		for _, elem := range parsedData.Elements {
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
					x, y, err := getDimensions(parsedData)
					if err != nil {

					}
					var wg sync.WaitGroup
					for frameIndex, frame := range data.NativeFrames {
						wg.Add(1)
						go generateNativeImage(frame, frameIndex, x, y, &wg) // parse image as gray scale
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

func generateNativeImage(frame [][]int, frameIndex, x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := image.NewGray16(image.Rect(0, 0, x, y))
	for j := 0; j < len(frame); j++ {
		i.SetGray16(j%x, j/y, color.Gray16{Y: uint16(frame[j][0])}) // for now, assume we're not overflowing uint16, assume gray image
	}
	name := fmt.Sprintf("image_%d.jpg", frameIndex)
	f, err := os.Create(name)
	if err != nil {
		fmt.Printf("Error while creating file: %s", err.Error())
	}
	jpeg.Encode(f, i, &jpeg.Options{Quality: 100})
	fmt.Printf("%s written \n", name)
}

func getDimensions(d *dicom.DataSet) (x, y int, err error) {
	rows, err := d.FindElementByTag(dicomtag.Rows)
	if err != nil {
		return 0, 0, err
	}

	cols, err := d.FindElementByTag(dicomtag.Columns)
	if err != nil {
		return 0, 0, err
	}

	return int(cols.MustGetUInt16()), int(rows.MustGetUInt16()), nil
}
