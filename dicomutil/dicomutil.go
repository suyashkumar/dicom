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
	printMetadata       = flag.Bool("print-metadata", true, "Print image metadata")
	extractImages       = flag.Bool("extract-images", false, "Extract images into separate files")
	extractImagesStream = flag.Bool("extract-images-stream", false, "Extract images using frame streaming capability")
	verbose             = flag.Bool("verbose", false, "Activate high verbosity log operation")
)

// FrameBufferSize represents the size of the *Frame buffered channel for streaming calls
const FrameBufferSize = 100

func main() {
	// Update usage docs
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n%s <dicom file> [flags]\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if *verbose {
		dicomlog.SetLevel(math.MaxInt32)
	}
	path := flag.Arg(0)

	var parsedData *dicom.DataSet

	if *extractImagesStream {
		// Stream process frames as they become available:
		frameChannel := make(chan *dicom.Frame, FrameBufferSize)
		p, err := dicom.NewParserFromFile(path, frameChannel)
		if err != nil {
			log.Panic("error creating parser", err)
		}

		// Go process frames published to frameChannel
		var wg sync.WaitGroup
		wg.Add(1)
		go writeStreamingFrames(frameChannel, &wg)

		// Begin parsing
		parsedData, err = p.Parse(dicom.ParseOptions{})
		if err != nil {
			log.Panic("error parsing", err)
		}

		// Wait for all frames to be streamed and processed
		wg.Wait()

	} else {
		// Non-streaming parsing:
		p, err := dicom.NewParserFromFile(path, nil)
		if err != nil {
			log.Panic("error creating new parser", err)
		}
		parsedData, err = p.Parse(dicom.ParseOptions{DropPixelData: !*extractImages})
		if parsedData == nil || err != nil {
			log.Panicf("Error reading %s: %v", path, err)
		}
		if *extractImages {
			for _, elem := range parsedData.Elements {
				if elem.Tag == dicomtag.PixelData {
					data := elem.Value[0].(dicom.PixelDataInfo)

					var wg sync.WaitGroup
					for frameIndex, frame := range data.Frames {
						wg.Add(1)
						go generateImage(&frame, frameIndex, &wg)
					}
					wg.Wait()

				}
			}
		}
	}

	// Print Metadata from parsedData if needed
	if *printMetadata {
		log.Println(parsedData)
		for _, elem := range parsedData.Elements {
			fmt.Printf("%v\n", elem.String())
		}
	}

	log.Println("Complete.")
}

func writeStreamingFrames(frameChan chan *dicom.Frame, doneWG *sync.WaitGroup) {
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

func generateImage(frame *dicom.Frame, frameIndex int, wg *sync.WaitGroup) {
	if frame.IsEncapsulated {
		go generateEncapsulatedImage(frame.EncapsulatedData, frameIndex, wg)
	} else {
		go generateNativeImage(frame.NativeData, frameIndex, wg)
	}
}

func generateEncapsulatedImage(frame dicom.EncapsulatedFrame, frameIndex int, wg *sync.WaitGroup) {
	defer wg.Done()
	path := fmt.Sprintf("image_%d.jpg", frameIndex) // TODO: figure out the image format
	ioutil.WriteFile(path, frame.Data, 0644)
	log.Printf("%s: %d bytes\n", path, len(frame.Data))
}

func generateNativeImage(frame dicom.NativeFrame, frameIndex int, wg *sync.WaitGroup) {
	defer wg.Done()
	i := image.NewGray16(image.Rect(0, 0, frame.Cols, frame.Rows))
	for j := 0; j < len(frame.Data); j++ {
		i.SetGray16(j%frame.Cols, j/frame.Rows, color.Gray16{Y: uint16(frame.Data[j][0])}) // for now, assume we're not overflowing uint16, assume gray image
	}
	name := fmt.Sprintf("image_%d.jpg", frameIndex)
	f, err := os.Create(name)
	if err != nil {
		fmt.Printf("Error while creating file: %s", err.Error())
	}
	jpeg.Encode(f, i, &jpeg.Options{Quality: 100})
	log.Printf("%s written \n", name)
}
