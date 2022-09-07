// Really basic sanity check program
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"sync"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/frame"
	"github.com/suyashkumar/dicom/pkg/tag"
)

// GitVersion is the current version of dicomutil, will be replaced in release step with current git commit hash or tag.
var GitVersion = "unknown"

var (
	version                  = flag.Bool("version", false, "print current version and exit")
	filepath                 = flag.String("path", "", "path")
	extractImagesStream      = flag.Bool("extract-images-stream", false, "Extract images using frame streaming capability")
	printJSON                = flag.Bool("json", false, "Print dataset as JSON")
	allowPixelDataVLMismatch = flag.Bool("allow-pixel-data-mismatch", false, "Allows the pixel data mismatch")
)

// FrameBufferSize represents the size of the *Frame buffered channel for streaming calls
const FrameBufferSize = 100

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("dicomutil: %s\n", GitVersion)
		os.Exit(0)
	}

	if len(*filepath) > 0 {

		f, err := os.Open(*filepath)
		if err != nil {
			log.Printf("Error opening file: %s\n", *filepath)
			return
		}
		defer f.Close()

		info, err := f.Stat()
		if err != nil {
			log.Println("err reading", err)
			return
		}

		var ds *dicom.Dataset
		if *extractImagesStream {
			ds = parseWithStreaming(f, info.Size())
		} else {
			var opts []dicom.ParseOption
			if *allowPixelDataVLMismatch {
				opts = append(opts, dicom.AllowMismatchPixelDataLength())
			}
			data, err := dicom.Parse(f, info.Size(), nil, opts...)
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

	// Go routine to process frames as they are sent to frameChannel
	var wg sync.WaitGroup
	wg.Add(1)
	go writeStreamingFrames(fc, &wg)

	ds, err := dicom.Parse(in, size, fc)
	if err != nil {
		log.Fatalf("error parsing: %v", err)
	}
	wg.Wait()

	return &ds

}

func writeStreamingFrames(frameChan chan *frame.Frame, doneWG *sync.WaitGroup) {
	count := 0 // may not correspond to frame number
	var wg sync.WaitGroup
	for fr := range frameChan {
		count++
		wg.Add(1)
		go generateImage(fr, count, "", &wg)
	}
	wg.Wait()
	doneWG.Done()
}

func generateImage(fr *frame.Frame, frameIndex int, frameSuffix string, wg *sync.WaitGroup) {
	i, err := fr.GetImage()
	if err != nil {
		fmt.Printf("Error while getting image: %v\n", err)
		return
	}

	ext := ".jpg"
	if !fr.IsEncapsulated() {
		ext = ".png"
	}

	name := fmt.Sprintf("image_%d%s%s", frameIndex, frameSuffix, ext)
	f, err := os.Create(name)
	if err != nil {
		fmt.Printf("Error while creating file: %s", err.Error())
		return
	}

	if !fr.IsEncapsulated() {
		// Native (non-encapsulated) frames are written as PNGs to exactly
		// preserve the pixel values.
		err := png.Encode(f, i)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		err = jpeg.Encode(f, i, &jpeg.Options{Quality: 100})
		if err != nil {
			log.Println(err)
			return
		}
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
