package dicom

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"testing"
)

func readFile() []byte {
	file, err := ioutil.ReadFile("examples/IM-0001-0001.dcm")
	if err != nil {
		fmt.Println("failed to read file")
		panic(err)
	}

	return file
}

func TestParseFile(t *testing.T) {

	file := readFile()

	parser, err := NewParser()
	if err != nil {
		t.Error(err)
	}

	data, err := parser.Parse(file)
	if err != nil {
		t.Errorf("failed to parse dicom file: %s", err)
	}

	elem, err := data.LookupElement("PatientName")
	if err != nil {
		t.Error(err)
	}

	pn := elem.Value.([]string)

	if pn[0] != "TOUTATIX" {
		t.Errorf("Incorrect patient name: %s", pn)
	}

	if l := len(pn); l != 1 {
		t.Errorf("Incorrect patient name length: %i", l)
	}

	elem, err = data.LookupElement("TransferSyntaxUID")
	if err != nil {
		t.Error(err)
	}

	ts := elem.Value.([]string)

	if ts[0] != "1.2.840.10008.1.2.4.91" {
		t.Errorf("Incorrect TransferSyntaxUID: %s", ts)
	}

	if l := len(data.Elements); l != 99 {
		t.Errorf("Error parsing DICOM file, wrong number of elements: %i", l)
	}

}

func TestGetTransferSyntaxImplicitLittleEndian(t *testing.T) {

	file := &DicomFile{}
	file.appendDataElement(&DicomElement{0002, 0010, "TransferSyntaxUID", "UI", 0, []string{"1.2.840.10008.1.2"}})

	bo, implicit, err := file.getTransferSyntax()
	if err != nil {
		t.Errorf("Could not get TransferSyntaxUID. %s", err)
	}

	if bo != binary.LittleEndian {
		t.Errorf("Incorrect ByteOrder %v. Should be LittleEndian.", bo)
	}

	if implicit != true {
		t.Errorf("Incorrect implicitness %v. Should be true.", implicit)
	}

}

func BenchmarkParseSingle(b *testing.B) {

	parser, _ := NewParser()

	for i := 0; i < b.N; i++ {

		file := readFile()

		_, err := parser.Parse(file)
		if err != nil {
			fmt.Println("failed to parse dicom file")
			panic(err)
		}
	}
}
