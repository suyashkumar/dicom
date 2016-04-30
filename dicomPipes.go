package dicom

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	fp "path/filepath"
	"sync"
)

type DicomMessage struct {
	msg  *DicomElement
	wait chan bool
}

const (
	JPEG_2000       = "1.2.840.10008.1.2.4.91"
	JPEG_BASELINE_1 = "1.2.840.10008.1.2.4.50"
)

// generator
func (di *DicomFile) Parse(buff []byte) <-chan DicomMessage {

	parser, err := NewParser()
	if err != nil {
		panic(err)
	}

	_, c := parser.Parse(buff)
	if err != nil {
		panic(err)
	}
	return c
}

// Discard messages
func (di *DicomFile) Discard(in <-chan DicomMessage, done *sync.WaitGroup) {
	done.Add(1)
	go func() {
		for dcmMsg := range in {
			dcmMsg.wait <- true
		}
		done.Done()
	}()

}

func fileName(folder string, i int, ext string) string {
	basename := fp.Base(folder)
	filename := basename + "_" + fmt.Sprintf("%03d\n", i)
	return fp.Join(folder, filename) + "." + ext
}

// Writes pixel data to folder
func (di *DicomFile) WriteImagesToFolder(in <-chan DicomMessage, done *sync.WaitGroup, folder string) <-chan DicomMessage {

	out := make(chan DicomMessage)
	waitMsg := make(chan bool)

	done.Add(1)
	go func() {

		var inImg bool = false
		var idx int
		var txUID, fext string

		for dcmMsg := range in {

			switch dcmMsg.msg.Name {

			case "TransferSyntaxUID":
				txUID = dcmMsg.msg.Value[0].(string)

			case "PixelData":
				inImg = true

				switch txUID {

				// JPEG baseline 1
				case JPEG_BASELINE_1:
					fext = "jpg"

				// JPEG 2000 Part 1
				case JPEG_2000:
					fext = "jp2"

				// not implemented
				default:
					//panic("Non implemented Transfer Syntax: \"" + txUID + "\"")

				}

			case "Item":
				if inImg == true {

					if idx > 0 {
						pb := dcmMsg.msg.Value[0].([]byte)
						err := ioutil.WriteFile(fileName(folder, idx, fext), pb, 0644)
						if err != nil {
							panic(err)
						}
					}

					idx++
				}
			}

			out <- DicomMessage{dcmMsg.msg, waitMsg}
			<-waitMsg
			dcmMsg.wait <- true
		}
		close(out)
		done.Done()
	}()
	return out

}

// Writes dicom elements to file
func (di *DicomFile) WriteToFile(in <-chan DicomMessage, done *sync.WaitGroup, file *os.File) <-chan DicomMessage {

	out := make(chan DicomMessage)
	waitMsg := make(chan bool)
	done.Add(1)

	go func() {
		for dcmMsg := range in {
			_, err := file.WriteString(fmt.Sprintln(dcmMsg.msg))
			if err != nil {
				panic(err)
			}

			out <- DicomMessage{dcmMsg.msg, waitMsg}
			<-waitMsg
			dcmMsg.wait <- true
		}
		file.Close()

		close(out)
		done.Done()
	}()

	return out
}

// Logs dicom elements
func (di *DicomFile) Log(in <-chan DicomMessage, done *sync.WaitGroup) <-chan DicomMessage {
	logger := log.New(os.Stdout, "", 0)
	out := make(chan DicomMessage)
	waitMsg := make(chan bool)

	done.Add(1)
	go func() {
		for dcmMsg := range in {
			logger.Println(dcmMsg.msg)
			out <- DicomMessage{dcmMsg.msg, waitMsg}
			<-waitMsg
			dcmMsg.wait <- true
		}
		close(out)
		done.Done()
	}()

	return out
}
