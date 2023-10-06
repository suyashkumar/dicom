// Copyright (c) 2023 Segmed Inc.
package codec

import (
	"bufio"
	"bytes"
	"image/jpeg"
	"os"
	"strings"
	"testing"

	"github.com/suyashkumar/dicom/pkg/frame"
)

type testify struct {
	in                   string
	out                  string
	transferSyntax       string
	w, h, prec, numcomps int
}

func TestMonoDecompresing(t *testing.T) {
	for _, tc := range []testify{
		{
			in:             "testdata/a1_mono.j2c",
			out:            "testdata/a1_mono.ppm",
			w:              303,
			h:              179,
			prec:           8,
			numcomps:       1,
			transferSyntax: "1.2.840.10008.1.2.4.90", // JPEG2000 Image Compression
		},
		{
			in:             "testdata/3_jpeg.dcm.1.raw",
			out:            "testdata/3.dcm.0.raw",
			w:              176,
			h:              176,
			prec:           16,
			numcomps:       1,
			transferSyntax: "1.2.840.10008.1.2.4.50", // JPEGProcess1TransferSyntax
		},
		{
			in:             "testdata/1_jpls.dcm.3.raw",
			out:            "testdata/1.dcm.1.raw",
			w:              64,
			h:              64,
			prec:           8,
			numcomps:       1,
			transferSyntax: "1.2.840.10008.1.2.4.80", // JPEG-LS Lossless
		},
	} {
		t.Run("decompress_"+tc.in, func(t *testing.T) {
			assertJPEGDecompressing(t, tc)
		})
	}
}

func readPPMPixelData(t testing.TB, fn string) []byte {
	file, err := os.Open(fn)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if file.Close() != nil {
			t.Fatal(err)
		}
	})
	scanner := bufio.NewScanner(file)
	var (
		expectedData []byte
		i            int
	)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		i++
		if i < 4 {
			continue
		}
		expectedData = scanner.Bytes()
		break
	}
	return expectedData
}

func assertJPEGDecompressing(t *testing.T, tc testify) {
	input, err := os.ReadFile(tc.in)
	if err != nil {
		t.Fatal(err)
	}
	fr := frame.EncapsulatedFrame{
		Data:            input,
		Rows:            tc.h,
		Cols:            tc.w,
		BitsAllocated:   tc.prec,
		SamplesPerPixel: tc.numcomps,
		TransferSyntax:  tc.transferSyntax,
	}
	uncompressedBytes, err := GetNativePixelData(fr)
	if err != nil {
		t.Fatal(err)
	}
	if len(uncompressedBytes) == 0 {
		t.Fatal("failed to decode image")
	}

	fr.Data = uncompressedBytes
	img, err := GetStdImage(fr)
	if err != nil {
		t.Fatal(err)
	}

	// convert to jpeg using go std lib
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		t.Fatal(err)
	}
	var expectedData []byte
	if strings.HasSuffix(tc.out, ".ppm") {
		expectedData = readPPMPixelData(t, tc.out)
	}
	if strings.HasSuffix(tc.out, ".raw") {
		expectedData, err = os.ReadFile(tc.out)
		if err != nil {
			t.Fatal(err)
		}
	}

	n := 10
	if !bytes.Equal(expectedData[:n], uncompressedBytes[:n]) {
		for i := 0; i < n; i++ {
			t.Logf("read byte %d in binary, expected %08b, got %08b", i, expectedData[i], uncompressedBytes[i])
		}
		t.Errorf("read the first %d bytes of %s: expected %v got %v", n, tc.out, expectedData[:n], uncompressedBytes[:n])
	}
	if len(uncompressedBytes) != len(expectedData) {
		t.Fatalf("read %s: expected length=%d, got %d", tc.out, len(expectedData), len(uncompressedBytes))
	}
	if !bytes.Equal(expectedData, uncompressedBytes) {
		t.Fatalf("read %s: expected %v got %v", tc.out, expectedData, uncompressedBytes)
	}
}
