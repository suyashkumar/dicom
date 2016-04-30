package dicom

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
)

type DicomMessage struct {
	msg  *DicomElement
	wait chan bool
}

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

func fileName(folderName, uid, ext string, i, n int) string {
	posf := ""
	if n > 1 {
		posf = "_" + strconv.Itoa(i)
	}
	return folderName + "\\" + uid + posf + "." + ext
}

// Writes pixel data to files
func (di *DicomFile) WriteImagesToFile(in <-chan DicomMessage, done *sync.WaitGroup, folderName string) <-chan DicomMessage {

	out := make(chan DicomMessage)
	waitMsg := make(chan bool)

	done.Add(1)
	go func() {

		var inImg bool = false
		var idx int
		var n int
		var txUID, instanceUID, fext string

		//first := true

		for dcmMsg := range in {

			switch dcmMsg.msg.Name {

			case "TransferSyntaxUID":
				txUID = dcmMsg.msg.Value[0].(string)

			case "SOPInstanceUID":
				instanceUID = dcmMsg.msg.Value[0].(string)

			case "PixelData":
				inImg = true

				switch txUID {
				//JPEG baseline 1
				case "1.2.840.10008.1.2.4.50":
					fext = "jpg"
				//JPEG 2000 Part 1
				case "1.2.840.10008.1.2.4.91":
					fext = "jp2"
				default:
					//panic("Non imlpemented Transfer Syntax: \"" + txUID + "\"")

				}

			case "Item":
				if inImg == true {

					// skip
					if idx == 0 {
						n = int(dcmMsg.msg.Vl/4) - 1
					}

					if idx > 0 {
						pb := dcmMsg.msg.Value[0].([]byte)
						err := ioutil.WriteFile(fileName(folderName, instanceUID, fext, idx, n), pb, 0644)
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
func (di *DicomFile) PrintToFile(in <-chan DicomMessage, done *sync.WaitGroup, folderName string) <-chan DicomMessage {

	out := make(chan DicomMessage)
	waitMsg := make(chan bool)
	done.Add(1)

	go func() {
		var f *os.File
		var instanceUID string

		fn1 := folderName + "/_tmp.txt"
		if _, err := os.Stat(fn1); os.IsExist(err) {
			err := os.Remove(fn1)
			if err != nil {
				panic(err)
			}
		}

		f1, err := os.Create(fn1)
		f = f1
		if err != nil {
			panic(err)
		}

		for dcmMsg := range in {
			if dcmMsg.msg.Name == "SOPInstanceUID" {
				instanceUID = dcmMsg.msg.Value[0].(string)
			}

			_, err := f.WriteString(fmt.Sprintln(dcmMsg.msg))
			if err != nil {
				panic(err)
			}

			out <- DicomMessage{dcmMsg.msg, waitMsg}
			<-waitMsg
			dcmMsg.wait <- true

		}

		f.Close()

		fn2 := folderName + "/" + instanceUID + ".txt"
		os.Rename(fn1, fn2)

		close(out)
		done.Done()
	}()
	return out

}

// Logs dicom elements
func (di *DicomFile) Log(in <-chan DicomMessage, done *sync.WaitGroup) <-chan DicomMessage {

	out := make(chan DicomMessage)
	waitMsg := make(chan bool)

	done.Add(1)
	go func() {
		for dcmMsg := range in {
			log.Println("->", dcmMsg.msg)
			out <- DicomMessage{dcmMsg.msg, waitMsg}
			<-waitMsg
			dcmMsg.wait <- true
		}
		close(out)
		done.Done()
	}()
	return out

}
