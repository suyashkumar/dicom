// Really basic sanity check program
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image/jpeg"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
)

var (
	filepath            = flag.String("path", "", "path")
	extractImagesStream = flag.Bool("extract-images-stream", false, "Extract images using frame streaming capability")
	printJSON           = flag.Bool("json", false, "Print dataset as JSON")
)

// FrameBufferSize represents the size of the *Frame buffered channel for streaming calls
const FrameBufferSize = 100

func main() {
	flag.Parse()
	if len(*filepath) > 0 {

		f, err := os.Open(*filepath)
		defer f.Close()
		if err != nil {
			log.Printf("Error opening file: %s\n", *filepath)
			return
		}

		info, err := f.Stat()
		if err != nil {
			log.Println("err reading", err)
			return
		}

		var ds *dicom.Dataset
		if *extractImagesStream {
			ds = parseWithStreaming(f, info.Size())
		} else {
			p, err := dicom.NewParser(f, info.Size(), nil)
			if err != nil {
				log.Fatalf("error creating parser: %v", err)
			}
			data, err := p.Parse()
			if err != nil {
				log.Fatalf("error parsing data: %v", err)
			}
			ds = &data
		}

		if *printJSON {
			j, err := json.MarshalIndent(ds, "", "  ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(j))
		} else {

			for z, elem := range ds.Elements {
				if elem.Tag == tag.PixelData && !*extractImagesStream {
					writePixelDataElement(elem, strconv.Itoa(z))
				}
				log.Println(elem.Tag)
				log.Println(elem.ValueLength)
				log.Println(elem.Value)
				// TODO: remove image icon hack after implementing flat iterator
				if elem.Tag == tag.IconImageSequence {
					for _, item := range elem.Value.GetValue().([]*dicom.SequenceItemValue) {
						for _, subElem := range item.GetValue().([]*dicom.Element) {
							if subElem.Tag == tag.PixelData {
								writePixelDataElement(subElem, "icon")
							}
						}
					}
				}

			}
		}
	}
}

func parseWithStreaming(in io.Reader, size int64) *dicom.Dataset {
	fc := make(chan *frame.Frame, FrameBufferSize)
	p, err := dicom.NewParser(in, size, fc)
	if err != nil {
		log.Fatalf("error creating parser: %v", err)
	}

	// Go routine to process frames as they are sent to frameChannel
	var wg sync.WaitGroup
	wg.Add(1)
	go writeStreamingFrames(fc, &wg)

	ds, err := p.Parse()
	if err != nil {
		log.Fatalf("error parsing: %v", err)
	}
	wg.Wait()

	return &ds

}

func writeStreamingFrames(frameChan chan *frame.Frame, doneWG *sync.WaitGroup) {
	count := 0 // may not correspond to frame number
	var wg sync.WaitGroup
	for frame := range frameChan {
		count++
		wg.Add(1)
		go generateImage(frame, count, &wg)
	}
	wg.Wait()
	doneWG.Done()
}

func generateImage(fr *frame.Frame, frameIndex int, wg *sync.WaitGroup) {
	i, err := fr.GetImage()
	if err != nil {
		log.Fatal("Error while getting image")
	}

	name := fmt.Sprintf("image_%d.jpg", frameIndex)
	f, err := os.Create(name)
	if err != nil {
		fmt.Printf("Error while creating file: %s", err.Error())
	}
	err = jpeg.Encode(f, i, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Println(err)
	}
	if err = f.Close(); err != nil {
		log.Println("ERROR: unable to properly close file: ", f.Name())
	}
	log.Printf("Image %s written\n", name)
	if wg != nil {
		wg.Done()
	}
}

func writePixelDataElement(e *dicom.Element, id string) {
	imageInfo := e.Value.GetValue().(dicom.PixelDataInfo)
	for idx, f := range imageInfo.Frames {
		generateImage(&f, idx, nil)
	}
}
