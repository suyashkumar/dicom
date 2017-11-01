package dicom_test

import (
	"os"
	"testing"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomtag"
	"github.com/grailbio/go-dicom/dicomuid"
	"v.io/x/lib/vlog"
)

func mustReadFile(path string, options dicom.ReadOptions) *dicom.DataSet {
	data, err := dicom.ReadDataSetFromFile(path, options)
	if err != nil {
		vlog.Fatalf("%s: failed to read: %v", path, err)
	}
	return data
}

func TestAllFiles(t *testing.T) {
	dir, err := os.Open("examples")
	if err != nil {
		panic(err)
	}
	names, err := dir.Readdirnames(0)
	if err != nil {
		panic(err)
	}
	for _, name := range names {
		vlog.Infof("Reading %s", name)
		_ = mustReadFile("examples/"+name, dicom.ReadOptions{})
	}
}

func testWriteFile(t *testing.T, dcmPath, transferSyntaxUID string) {
	data := mustReadFile(dcmPath, dicom.ReadOptions{})
	dstPath := "/tmp/test.dcm"
	out, err := os.Create(dstPath)
	if err != nil {
		t.Fatal(err)
	}

	for i := range data.Elements {
		if data.Elements[i].Tag == dicomtag.TransferSyntaxUID {
			newElem := dicom.MustNewElement(dicomtag.TransferSyntaxUID, transferSyntaxUID)
			vlog.Infof("Setting transfer syntax UID from %v to %v",
				data.Elements[i].MustGetString(), newElem.MustGetString())
			data.Elements[i] = newElem
		}
	}
	err = dicom.WriteDataSet(out, data)
	if err != nil {
		t.Fatal(err)
	}
	data2 := mustReadFile(dstPath, dicom.ReadOptions{})

	if len(data.Elements) != len(data2.Elements) {
		t.Errorf("Wrong # of elements: %v %v", len(data.Elements), len(data2.Elements))
		for _, elem := range data.Elements {
			if _, err := data2.FindElementByTag(elem.Tag); err != nil {
				t.Errorf("Tag %v found in org, but not in new", dicomtag.DebugString(elem.Tag))
			}
		}
		for _, elem := range data2.Elements {
			if _, err := data.FindElementByTag(elem.Tag); err != nil {
				t.Errorf("Tag %v found in new, but not in org", dicomtag.DebugString(elem.Tag))
			}
		}
	}
	for _, elem := range data.Elements {
		elem2, err := data2.FindElementByTag(elem.Tag)
		if err != nil {
			t.Error(err)
			continue
		}
		if elem.Tag == dicomtag.FileMetaInformationGroupLength {
			// This element is expected to change when the file is transcoded.
			continue
		}
		if elem.String() != elem2.String() {
			t.Fatalf("Elem mismatch: %v %v", elem.String(), elem2.String())
		}
	}
	// TODO(saito) Fix below.
	//if !reflect.DeepEqual(data, data2) {
	//	t.Error("Files aren't equal")
	//}
}

func TestWriteFile(t *testing.T) {
	// path := "examples/IM-0001-0001.dcm"
	//testWriteFile(t, "examples/CT-MONO2-16-ort.dcm", dicomuid.ExplicitVRBigEndian)
	//testWriteFile(t, "examples/CT-MONO2-16-ort.dcm", dicomuid.ImplicitVRLittleEndian)
	testWriteFile(t, "examples/CT-MONO2-16-ort.dcm", dicomuid.ExplicitVRLittleEndian)
}

func TestReadDataSet(t *testing.T) {
	data := mustReadFile("examples/IM-0001-0001.dcm", dicom.ReadOptions{})
	elem, err := data.FindElementByName("PatientName")
	if err != nil {
		t.Error(err)
	}
	if elem.MustGetString() != "TOUTATIX" {
		t.Errorf("Incorrect patient name: %s", elem)
	}
	elem, err = data.FindElementByName("TransferSyntaxUID")
	if err != nil {
		t.Error(err)
	}
	if elem.MustGetString() != "1.2.840.10008.1.2.4.91" {
		t.Errorf("Incorrect TransferSyntaxUID: %s", elem)
	}
	if l := len(data.Elements); l != 98 {
		t.Errorf("Error parsing DICOM file, wrong number of elements: %d", l)
	}
	elem, err = data.FindElementByTag(dicomtag.PixelData)
	if err != nil {
		t.Error(err)
	}
}

// Test ReadOptions.DropPixelData.
func TestDropPixelData(t *testing.T) {
	data := mustReadFile("examples/IM-0001-0001.dcm", dicom.ReadOptions{DropPixelData: true})
	_, err := data.FindElementByTag(dicomtag.PatientName)
	if err != nil {
		t.Error(err)
	}
	_, err = data.FindElementByTag(dicomtag.PixelData)
	if err == nil {
		t.Errorf("PixelData should not be present")
	}
}

func BenchmarkParseSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = mustReadFile("examples/IM-0001-0001.dcm", dicom.ReadOptions{})
	}
}
