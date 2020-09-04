package dicom

import (
	"testing"
	// "os"

  "github.com/stretchr/testify/assert"
  // "github.com/suyashkumar/dicom/pkg/dicomio"
  "github.com/suyashkumar/dicom/pkg/tag"
)

func TestWriteFileHeader(t *testing.T) {
	// // Create the file
	// location := "fileheader.dcm"
	// file, err := os.Open(location)
	// assert.Nil(t, err_)
	// defer file.Close()
	//
	// // create a writer on the file
  // w := dicomio.NewWriter(file, binary.LittleEndian, false)
	//
	// // Create the meta elements
	//
	// // write the file header
	// err := writeFileHeader(w, ds, metaElems, nil)
	// assert.Nil(t, err)
	//
	// // read the file
	// info, err := file.Stat()
	// assert.Nil(t, err)
	// p, err := dicom.NewParser(file, info.Size(), nil)
	//
	// parsedData, err := p.Parse()
	// assert.Nil(t, err)
	//
	// // Verify the the corrrect things were written to the file header


}

func TestEncodeElementHeader(t *testing.T) {}

func TestWriteValue(t *testing.T) {}

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

func TestWrite(t *testing.T) {}
