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
	"sync"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
)

var (
	filepath                = flag.String("path", "", "path")
	extractImagesStream     = flag.Bool("extract-images-stream", false, "Extract images using frame streaming capability")
	printJSON               = flag.Bool("json", false, "Print dataset as JSON")
	assumeNoHeaderAndOffset = flag.Bool("noheader", false,
		"Set to true if you want to read the dicom assuming no header or 128 byte offset")
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
			var p dicom.Parser
			var err error
			if *assumeNoHeaderAndOffset {
				p, err = dicom.NewParser(f, info.Size(), nil, dicom.AssumeNoHeaderAndOffset)
			} else {
				p, err = dicom.NewParser(f, info.Size(), nil)
			}
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
			log.Println("Printing DICOM dataset serialized as JSON to stdout")
			j, err := json.MarshalIndent(ds, "", "  ")
			if err != nil {
				panic(err)
			}

			fmt.Println(string(j))
		} else {
			log.Println("Printing DICOM dataset parsed elements to stdout:")
			fmt.Print(ds)

			// In non-streaming frame mode, we need to find all PixelData elements and generate images.
			for _, elem := range ds.Elements {
				if elem.Tag == tag.PixelData && !*extractImagesStream {
					writePixelDataElement(elem, "")
				}
				// TODO: remove image icon hack after implementing flat iterator
				if elem.Tag == tag.IconImageSequence {
					for _, item := range elem.Value.GetValue().([]*dicom.SequenceItemValue) {
						for _, subElem := range item.GetValue().([]*dicom.Element) {
							if subElem.Tag == tag.PixelData {
								writePixelDataElement(subElem, "_icon")
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
		go generateImage(frame, count, "", &wg)
	}
	wg.Wait()
	doneWG.Done()
}

func generateImage(fr *frame.Frame, frameIndex int, frameSuffix string, wg *sync.WaitGroup) {
	i, err := fr.GetImage()
	if err != nil {
		log.Fatal("Error while getting image")
	}

	name := fmt.Sprintf("image_%d%s.jpg", frameIndex, frameSuffix)
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

func writePixelDataElement(e *dicom.Element, suffix string) {
	imageInfo := e.Value.GetValue().(dicom.PixelDataInfo)
	for idx, f := range imageInfo.Frames {
		generateImage(&f, idx, suffix, nil)
	}
}
