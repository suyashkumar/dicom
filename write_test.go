package dicom

import (
	"testing"
	"os"

  "github.com/stretchr/testify/assert"
  "github.com/suyashkumar/dicom/pkg/dicomio"
  "github.com/suyashkumar/dicom/pkg/tag"
)

func TestWriteFileHeader(t *testing.T) {
	// Create the file
	location := "fileheader.dcm"
	file, err := os.Open(location)
	assert.Nil(t, err_)
	defer file.Close()

	// create a writer on the file
  w := dicomio.NewWriter(file, binary.LittleEndian, false)

	// Create the meta elements

	// write the file header
	err := writeFileHeader(w, ds, metaElems, nil)
	assert.Nil(t, err)

	// read the file
	info, err := file.Stat()
	assert.Nil(t, err)
	p, err := dicom.NewParser(file, info.Size(), nil)

	parsedData, err := p.Parse()
	assert.Nil(t, err)

	// Verify the the corrrect things were written to the file header


}
