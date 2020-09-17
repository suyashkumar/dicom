package dicom

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/suyashkumar/dicom/pkg/dicomio"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
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

func TestReadFloat_float64(t *testing.T) {
	cases := []struct{
		name string
		floats []float64
		VR string
		want Value
		expectedErr error
	}{
		{
			name: "float64",
			floats:[]float64{20.1, 32.22},
			VR: "FD",
			want: &floatsValue{value: []float64{20.1, 32.22}},
			expectedErr: nil,
		},
		{
			name: "float64 with wrong VR",
			floats:[]float64{20.1, 32.22},
			VR: "XX",
			want: nil,
			expectedErr: errorUnableToParseFloat,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := bytes.Buffer{}
			for _, fl := range tc.floats {
				if err := binary.Write(&data, binary.LittleEndian, fl); err != nil {
					t.Errorf("TestReadFloat: Unable to setup test buffer")
				}
			}

			r, err := dicomio.NewReader(&data, binary.LittleEndian, int64(data.Len()))
			if err != nil {
				t.Errorf("TestReadFloat: unable to create new dicomio.Reader")
			}

			got, err := readFloat(r, tag.Tag{}, tc.VR, uint32(data.Len()))
			if err != tc.expectedErr {
				t.Fatalf("readFloat(r, tg, %s, %d) got unexpected error: got: %v, want: %v", tc.VR, data.Len(), err, tc.expectedErr)
			}
			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(floatsValue{})); diff != "" {
				t.Errorf("readFloat(r, tg, %s, %d) unexpected diff: %s", tc.VR, data.Len(), diff)
			}
		})
	}
}

func TestReadFloat_float32(t *testing.T) {
	cases := []struct{
		name string
		floats []float32
		VR string
		want Value
		expectedErr error
	}{
		{
			name: "float32",
			floats:[]float32{20.1, 32.22},
			VR: "FL",
			// TODO(suyashkumar): look into extra significant digits added when going from float32 -> float64. One
			// option is having both a float32 and float64 ValueType.
			want: &floatsValue{value: []float64{float64(float32(20.1)), float64(float32(32.22))}},
			expectedErr: nil,
		},
		{
			name: "float32 with wrong VR",
			floats:[]float32{20.1, 32.22},
			VR: "XX",
			want: nil,
			expectedErr: errorUnableToParseFloat,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			data := bytes.Buffer{}
			for _, fl := range tc.floats {
				if err := binary.Write(&data, binary.LittleEndian, fl); err != nil {
					t.Errorf("TestReadFloat: Unable to setup test buffer")
				}
			}

			r, err := dicomio.NewReader(&data, binary.LittleEndian, int64(data.Len()))
			if err != nil {
				t.Errorf("TestReadFloat: unable to create new dicomio.Reader")
			}

			got, err := readFloat(r, tag.Tag{}, tc.VR, uint32(data.Len()))
			if err != tc.expectedErr {
				t.Fatalf("readFloat(r, tg, %s, %d) got unexpected error: got: %v, want: %v", tc.VR, data.Len(), err, tc.expectedErr)
			}
			if diff := cmp.Diff(got, tc.want, cmp.AllowUnexported(floatsValue{})); diff != "" {
				t.Errorf("readFloat(r, tg, %s, %d) unexpected diff: %s", tc.VR, data.Len(), diff)
			}
		})
	}
}
