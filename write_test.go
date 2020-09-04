package dicom

import (
	"testing"
	"os"
	"encoding/binary"

  "github.com/stretchr/testify/assert"
  "github.com/suyashkumar/dicom/pkg/dicomio"
  "github.com/suyashkumar/dicom/pkg/tag"
	"github.com/suyashkumar/dicom/pkg/uid"
)

func TestWrite(t *testing.T) {

}

func TestWriteFileHeader(t *testing.T) {
	// Create the file
	location := "fileheader.dcm"
	file, err := os.Create(location)
	assert.Nil(t, err)
	defer file.Close()

	// create a writer on the file
  w := dicomio.NewWriter(file, binary.LittleEndian, false)

	// Create the meta elements
	transferSyntax, err := newElement(tag.TransferSyntaxUID, []string{uid.ImplicitVRLittleEndian})
	assert.Nil(t, err)
	mediaStorage, err := newElement(tag.MediaStorageSOPClassUID, []string{"1.2.840.10008.5.1.4.1.1.1.2"})
	assert.Nil(t, err)
	metaElems := []*Element{
							 transferSyntax,
							 mediaStorage,
	}
	ds := &Dataset{Elements: metaElems}

	// write the file header
	err = writeFileHeader(w, ds, metaElems, func (set *writeOptSet){})
	assert.Nil(t, err)

	// // read the file
	// info, err := file.Stat()
	// assert.Nil(t, err)
	// p, err := NewParser(file, info.Size(), nil)
	//
	// parsedData, err := p.Parse()
	// assert.Nil(t, err)

	// Verify the the corrrect things were written to the file header


}

func TestEncodeElementHeader(t *testing.T) {}

func TestWriteValue(t *testing.T) {}

func TestWriteTag(t *testing.T) {}

func TestWriteVRVL(t *testing.T) {}

func TestVerifyVR(t *testing.T) {
	tg := tag.Tag{ // FileMetaInformationGroupLength tag
		Group: 0x0002,
		Element: 0x0000,
	}

	// give the wrong VR
	vr, err := verifyVR(tg, "OB")
	assert.Equal(t, "", vr)
	assert.NotNil(t, err)

	// No vr given
	vr, err = verifyVR(tg, "")
	assert.Nil(t, err)
	assert.Equal(t, "UL", vr)

	// made up tag
	tg = tag.Tag{
		Group: 0x9999,
		Element: 0x9999,
	}
	vr, err = verifyVR(tg, "")
	assert.Nil(t, err)
	assert.Equal(t, "UN", vr)
}

func TestVerifyValueType(t *testing.T) {
	tg := tag.Tag{ // FileMetaInformationGroupLength tag
		Group: 0x0002,
		Element: 0x0000,
	}

	// VALID
	value, err := NewValue([]int{128})
	assert.Nil(t, err)
	err = verifyValueType(tg, value, Ints, "UL")
	assert.Nil(t, err)

	// INVALID VR
	err = verifyValueType(tg, value, Ints, "NA") // incorrect vr
	assert.NotNil(t, err)

	// WRONG VALUE TYPE
	err = verifyValueType(tg, value, Strings, "UL")
	assert.NotNil(t, err)
}
