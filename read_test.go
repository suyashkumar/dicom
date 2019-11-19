package dicom

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_dicomio "github.com/suyashkumar/dicom/mocks/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func TestReadTag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	assert := assert.New(t)

	m := mock_dicomio.NewMockReader(ctrl)

	first := m.EXPECT().ReadUInt16().Return(uint16(0x7FE0), nil)
	m.EXPECT().ReadUInt16().Return(uint16(0x0010), nil).After(first)

	retTag, err := readTag(m)

	assert.NoError(err)
	fmt.Println(err)
	assert.Equal(tag.PixelData, *retTag)

}
