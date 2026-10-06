package frame

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"strings"
	"sync"
	"testing"
)

func makeBaselineJPEG(t *testing.T) []byte {
	t.Helper()
	img := image.NewGray(image.Rect(0, 0, 8, 8))
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			img.SetGray(x, y, color.Gray{Y: uint8(x*10 + y*10)})
		}
	}
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 90}); err != nil {
		t.Fatalf("failed to encode baseline JPEG: %v", err)
	}
	return buf.Bytes()
}

func makeSOF3JPEG() []byte {
	// Minimal JPEG stream containing SOF3 (0xFFC3) lossless process marker
	return []byte{
		0xFF, 0xD8, // SOI
		0xFF, 0xC3, // SOF3 marker
		0x00, 0x0B, // Length = 11 bytes
		0x08,       // Precision = 8 bits
		0x00, 0x10, // Height = 16
		0x00, 0x10, // Width = 16
		0x01,             // 1 Component
		0x01, 0x11, 0x00, // Component 1: ID=1, sampling=0x11, table=0
		0xFF, 0xDA, // SOS
		0x00, 0x08, // Length = 8 bytes
		0x01,       // 1 Component
		0x01, 0x00, // Component 1, table=0
		0x00, 0x00, 0x00, // Ss, Se, Ah/Al
		0x12, 0x34, // Mock entropy payload
		0xFF, 0xD9, // EOI
	}
}

func makeSOF55JPEG() []byte {
	// Minimal JPEG stream containing SOF55 (0xFFF7) JPEG-LS marker
	return []byte{
		0xFF, 0xD8, // SOI
		0xFF, 0xF7, // SOF55 marker
		0x00, 0x0B, // Length = 11 bytes
		0x08,       // Precision = 8 bits
		0x00, 0x10, // Height = 16
		0x00, 0x10, // Width = 16
		0x01,             // 1 Component
		0x01, 0x11, 0x00, // Component 1
		0xFF, 0xD9, // EOI
	}
}

func makeJPEG2000Container() []byte {
	// JPEG 2000 JP2 container signature box (12 bytes)
	return []byte{
		0x00, 0x00, 0x00, 0x0C,
		0x6A, 0x50, 0x20, 0x20,
		0x0D, 0x0A, 0x87, 0x0A,
		0x00, 0x00, 0x00, 0x14, // ftyp box length
		0x66, 0x74, 0x79, 0x70, // "ftyp"
	}
}

func makeJPEG2000Codestream() []byte {
	// Raw JPEG 2000 codestream starting with SOC (0xFF4F) and SIZ (0xFF51)
	return []byte{
		0xFF, 0x4F, // SOC
		0xFF, 0x51, // SIZ
		0x00, 0x29, // Length
		0x00, 0x00, 0x00, 0x00,
	}
}

func TestEncapsulatedDecoderRegistrationAndDispatch(t *testing.T) {
	ResetEncapsulatedDecoders()
	t.Cleanup(ResetEncapsulatedDecoders)

	dummyImg := image.NewRGBA(image.Rect(0, 0, 4, 4))
	dummyImg.Set(0, 0, color.RGBA{R: 255, A: 255})

	called := false
	decoder := func(data []byte) (image.Image, error) {
		called = true
		if string(data) != "lossless-payload" {
			t.Errorf("unexpected decoder data payload: %s", string(data))
		}
		return dummyImg, nil
	}

	const tsUID = "1.2.840.10008.1.2.4.70"
	RegisterEncapsulatedDecoder(tsUID, decoder)

	gotDec, ok := GetEncapsulatedDecoder(tsUID)
	if !ok || gotDec == nil {
		t.Fatalf("expected registered decoder for %s, got ok=%v", tsUID, ok)
	}

	// Verify retrieval with formatting variations (e.g. leading =, whitespace)
	if _, found := GetEncapsulatedDecoder(" =" + tsUID + " "); !found {
		t.Errorf("expected normalized lookup to match %s", tsUID)
	}

	f := EncapsulatedFrame{
		Data:              []byte("lossless-payload"),
		TransferSyntaxUID: tsUID,
	}

	img, err := f.GetImage()
	if err != nil {
		t.Fatalf("unexpected error from GetImage: %v", err)
	}
	if !called {
		t.Errorf("custom registered decoder was not invoked")
	}
	if img != dummyImg {
		t.Errorf("GetImage returned unexpected image pointer")
	}

	// Verify invocation through Frame wrapper
	called = false
	frameWrapper := Frame{
		Encapsulated:     true,
		EncapsulatedData: f,
	}
	imgWrap, err := frameWrapper.GetImage()
	if err != nil {
		t.Fatalf("Frame.GetImage failed: %v", err)
	}
	if !called {
		t.Errorf("custom registered decoder not called via Frame wrapper")
	}
	if imgWrap != dummyImg {
		t.Errorf("Frame.GetImage returned unexpected image pointer")
	}
}

func TestEncapsulatedDecoderUnregister(t *testing.T) {
	ResetEncapsulatedDecoders()
	t.Cleanup(ResetEncapsulatedDecoders)

	const tsUID = "1.2.840.10008.1.2.4.70"
	dummyDecoder := func(data []byte) (image.Image, error) {
		return image.NewGray(image.Rect(0, 0, 1, 1)), nil
	}

	RegisterEncapsulatedDecoder(tsUID, dummyDecoder)
	if _, ok := GetEncapsulatedDecoder(tsUID); !ok {
		t.Fatalf("expected decoder to be registered")
	}

	UnregisterEncapsulatedDecoder(tsUID)
	if _, ok := GetEncapsulatedDecoder(tsUID); ok {
		t.Errorf("decoder still returned after UnregisterEncapsulatedDecoder")
	}

	// Registering nil unregisters as well
	RegisterEncapsulatedDecoder(tsUID, dummyDecoder)
	RegisterEncapsulatedDecoder(tsUID, nil)
	if _, ok := GetEncapsulatedDecoder(tsUID); ok {
		t.Errorf("decoder still returned after registering nil decoder")
	}

	// Test ResetEncapsulatedDecoders clearing multiple
	RegisterEncapsulatedDecoder("1.2.840.10008.1.2.4.70", dummyDecoder)
	RegisterEncapsulatedDecoder("1.2.840.10008.1.2.4.80", dummyDecoder)
	ResetEncapsulatedDecoders()
	if _, ok := GetEncapsulatedDecoder("1.2.840.10008.1.2.4.70"); ok {
		t.Errorf("decoder 1 still returned after ResetEncapsulatedDecoders")
	}
	if _, ok := GetEncapsulatedDecoder("1.2.840.10008.1.2.4.80"); ok {
		t.Errorf("decoder 2 still returned after ResetEncapsulatedDecoders")
	}
}

func TestFallbackBehaviorWhenNoDecoderRegistered(t *testing.T) {
	ResetEncapsulatedDecoders()
	t.Cleanup(ResetEncapsulatedDecoders)

	jpegBytes := makeBaselineJPEG(t)

	// Frame without transfer syntax UID set decodes using image/jpeg
	f1 := EncapsulatedFrame{
		Data: jpegBytes,
	}
	img1, err := f1.GetImage()
	if err != nil {
		t.Fatalf("GetImage on baseline JPEG without TS UID failed: %v", err)
	}
	if img1.Bounds().Dx() != 8 || img1.Bounds().Dy() != 8 {
		t.Errorf("unexpected image bounds: %v", img1.Bounds())
	}

	// Frame with standard baseline transfer syntax UID decodes cleanly
	f2 := EncapsulatedFrame{
		Data:              jpegBytes,
		TransferSyntaxUID: "1.2.840.10008.1.2.4.50",
	}
	img2, err := f2.GetImage()
	if err != nil {
		t.Fatalf("GetImage on baseline JPEG with TS UID 1.2.840.10008.1.2.4.50 failed: %v", err)
	}
	if img2.Bounds().Dx() != 8 || img2.Bounds().Dy() != 8 {
		t.Errorf("unexpected image bounds: %v", img2.Bounds())
	}
}

func TestActionableError_SOF3JPEGLossless(t *testing.T) {
	ResetEncapsulatedDecoders()
	t.Cleanup(ResetEncapsulatedDecoders)

	sof3Bytes := makeSOF3JPEG()

	// Scenario A: No TransferSyntaxUID provided, raw SOF3 stream
	fA := EncapsulatedFrame{
		Data: sof3Bytes,
	}
	_, errA := fA.GetImage()
	if errA == nil {
		t.Fatalf("expected error on SOF3 data without registered decoder, got nil")
	}

	if !errors.Is(errA, ErrUnsupportedEncapsulatedCodec) {
		t.Errorf("expected error matching ErrUnsupportedEncapsulatedCodec, got: %v", errA)
	}

	var uErrA *UnsupportedCodecError
	if !errors.As(errA, &uErrA) {
		t.Fatalf("expected error convertible to *UnsupportedCodecError, got: %T", errA)
	}
	if uErrA.Marker != 0xFFC3 {
		t.Errorf("expected marker 0xFFC3, got 0x%04X", uErrA.Marker)
	}

	errMsgA := errA.Error()
	if !strings.Contains(errMsgA, "SOF3 marker 0xFFC3") {
		t.Errorf("error message missing marker information: %s", errMsgA)
	}
	if !strings.Contains(errMsgA, "RegisterEncapsulatedDecoder") {
		t.Errorf("error message missing RegisterEncapsulatedDecoder guidance: %s", errMsgA)
	}
	if !strings.Contains(errMsgA, "EncapsulatedFrame.Data") {
		t.Errorf("error message missing EncapsulatedFrame.Data guidance: %s", errMsgA)
	}

	// Scenario B: TransferSyntaxUID set to 1.2.840.10008.1.2.4.70
	const tsUID = "1.2.840.10008.1.2.4.70"
	fB := EncapsulatedFrame{
		Data:              sof3Bytes,
		TransferSyntaxUID: tsUID,
	}
	_, errB := fB.GetImage()
	if errB == nil {
		t.Fatalf("expected error on SOF3 data without registered decoder, got nil")
	}
	if !strings.Contains(errB.Error(), tsUID) {
		t.Errorf("error message should mention transfer syntax %s, got: %s", tsUID, errB.Error())
	}

	// Scenario C: Register decoder for 1.2.840.10008.1.2.4.70, verify successful decode
	decodedImg := image.NewGray(image.Rect(0, 0, 16, 16))
	RegisterEncapsulatedDecoder(tsUID, func(data []byte) (image.Image, error) {
		return decodedImg, nil
	})

	imgC, errC := fB.GetImage()
	if errC != nil {
		t.Fatalf("unexpected error after registering decoder: %v", errC)
	}
	if imgC != decodedImg {
		t.Errorf("returned unexpected image from registered decoder")
	}

	// Also verify that fA (with empty TransferSyntaxUID) now decodes via detected SOF3 transfer syntax
	imgA2, errA2 := fA.GetImage()
	if errA2 != nil {
		t.Fatalf("unexpected error when decoding fA with detected transfer syntax decoder: %v", errA2)
	}
	if imgA2 != decodedImg {
		t.Errorf("returned unexpected image for fA from detected decoder")
	}
}

func TestActionableError_SOF55JPEGLS(t *testing.T) {
	ResetEncapsulatedDecoders()
	t.Cleanup(ResetEncapsulatedDecoders)

	sof55Bytes := makeSOF55JPEG()
	f := EncapsulatedFrame{
		Data: sof55Bytes,
	}
	_, err := f.GetImage()
	if err == nil {
		t.Fatalf("expected error on SOF55 data without registered decoder, got nil")
	}

	if !errors.Is(err, ErrUnsupportedEncapsulatedCodec) {
		t.Errorf("expected ErrUnsupportedEncapsulatedCodec, got: %v", err)
	}

	var uErr *UnsupportedCodecError
	if !errors.As(err, &uErr) {
		t.Fatalf("expected *UnsupportedCodecError, got: %T", err)
	}
	if uErr.Marker != 0xFFF7 {
		t.Errorf("expected marker 0xFFF7, got 0x%04X", uErr.Marker)
	}

	msg := err.Error()
	if !strings.Contains(msg, "SOF55 marker 0xFFF7") || !strings.Contains(msg, "JPEG-LS") {
		t.Errorf("error message missing JPEG-LS / SOF55 details: %s", msg)
	}

	// Register decoder and verify dispatch
	customImg := image.NewRGBA(image.Rect(0, 0, 16, 16))
	RegisterEncapsulatedDecoder("1.2.840.10008.1.2.4.80", func(data []byte) (image.Image, error) {
		return customImg, nil
	})

	gotImg, err := f.GetImage()
	if err != nil {
		t.Fatalf("unexpected error after registering JPEG-LS decoder: %v", err)
	}
	if gotImg != customImg {
		t.Errorf("expected custom image from JPEG-LS decoder")
	}
}

func TestActionableError_JPEG2000(t *testing.T) {
	ResetEncapsulatedDecoders()
	t.Cleanup(ResetEncapsulatedDecoders)

	// Test JP2 container format
	jp2Bytes := makeJPEG2000Container()
	f1 := EncapsulatedFrame{Data: jp2Bytes}
	_, err1 := f1.GetImage()
	if err1 == nil {
		t.Fatalf("expected error on JP2 container data, got nil")
	}
	if !errors.Is(err1, ErrUnsupportedEncapsulatedCodec) {
		t.Errorf("expected ErrUnsupportedEncapsulatedCodec, got: %v", err1)
	}
	if !strings.Contains(err1.Error(), "JPEG 2000") {
		t.Errorf("error message missing JPEG 2000 description: %s", err1.Error())
	}

	// Test raw codestream format
	j2kBytes := makeJPEG2000Codestream()
	f2 := EncapsulatedFrame{Data: j2kBytes}
	_, err2 := f2.GetImage()
	if err2 == nil {
		t.Fatalf("expected error on J2K raw codestream, got nil")
	}
	if !errors.Is(err2, ErrUnsupportedEncapsulatedCodec) {
		t.Errorf("expected ErrUnsupportedEncapsulatedCodec for raw codestream, got: %v", err2)
	}
}

func TestActionableError_DeclaredTransferSyntax(t *testing.T) {
	ResetEncapsulatedDecoders()
	t.Cleanup(ResetEncapsulatedDecoders)

	// Truncated payload with explicit RLE transfer syntax
	f := EncapsulatedFrame{
		Data:              []byte{0x01, 0x02, 0x03, 0x04},
		TransferSyntaxUID: "1.2.840.10008.1.2.5",
	}
	_, err := f.GetImage()
	if err == nil {
		t.Fatalf("expected error on RLE transfer syntax without decoder, got nil")
	}
	if !errors.Is(err, ErrUnsupportedEncapsulatedCodec) {
		t.Errorf("expected ErrUnsupportedEncapsulatedCodec for declared RLE transfer syntax, got: %v", err)
	}
	if !strings.Contains(err.Error(), "1.2.840.10008.1.2.5") {
		t.Errorf("error message missing transfer syntax: %s", err.Error())
	}
}

func TestMalformedData_ReturnsStandardError(t *testing.T) {
	ResetEncapsulatedDecoders()
	t.Cleanup(ResetEncapsulatedDecoders)

	// Random corrupt bytes (no SOI, no known format)
	f1 := EncapsulatedFrame{
		Data: []byte{0xDE, 0xAD, 0xBE, 0xEF},
	}
	_, err1 := f1.GetImage()
	if err1 == nil {
		t.Fatalf("expected error on invalid bytes, got nil")
	}
	if errors.Is(err1, ErrUnsupportedEncapsulatedCodec) {
		t.Errorf("malformed data should NOT return ErrUnsupportedEncapsulatedCodec, got: %v", err1)
	}

	// Truncated SOI without markers
	f2 := EncapsulatedFrame{
		Data: []byte{0xFF, 0xD8},
	}
	_, err2 := f2.GetImage()
	if err2 == nil {
		t.Fatalf("expected error on truncated SOI bytes, got nil")
	}
	if errors.Is(err2, ErrUnsupportedEncapsulatedCodec) {
		t.Errorf("truncated data should NOT return ErrUnsupportedEncapsulatedCodec, got: %v", err2)
	}
}

func TestEmptyFrameData(t *testing.T) {
	f := EncapsulatedFrame{Data: nil}
	_, err := f.GetImage()
	if err == nil {
		t.Fatalf("expected error on empty frame data, got nil")
	}
}

func TestEncapsulatedFrame_Equals(t *testing.T) {
	f1 := &EncapsulatedFrame{Data: []byte{1, 2, 3}, TransferSyntaxUID: "1.2.3"}
	f2 := &EncapsulatedFrame{Data: []byte{1, 2, 3}, TransferSyntaxUID: "1.2.3"}
	fDiffData := &EncapsulatedFrame{Data: []byte{1, 2, 4}, TransferSyntaxUID: "1.2.3"}
	fDiffTS := &EncapsulatedFrame{Data: []byte{1, 2, 3}, TransferSyntaxUID: "1.2.4"}

	if !f1.Equals(f2) {
		t.Errorf("f1.Equals(f2) returned false, want true")
	}
	if f1.Equals(fDiffData) {
		t.Errorf("f1.Equals(fDiffData) returned true, want false")
	}
	if f1.Equals(fDiffTS) {
		t.Errorf("f1.Equals(fDiffTS) returned true, want false")
	}
	if f1.Equals(nil) {
		t.Errorf("f1.Equals(nil) returned true, want false")
	}

	var nilFrame *EncapsulatedFrame
	if !nilFrame.Equals(nil) {
		t.Errorf("nilFrame.Equals(nil) returned false, want true")
	}
	if nilFrame.Equals(f1) {
		t.Errorf("nilFrame.Equals(f1) returned true, want false")
	}
}

func TestFrame_SetTransferSyntaxUID(t *testing.T) {
	enc := EncapsulatedFrame{Data: []byte{1, 2}}
	enc.SetTransferSyntaxUID("1.2.840.10008.1.2.4.70")
	if enc.TransferSyntaxUID != "1.2.840.10008.1.2.4.70" {
		t.Errorf("unexpected TransferSyntaxUID: %s", enc.TransferSyntaxUID)
	}

	frame := Frame{
		Encapsulated:     true,
		EncapsulatedData: EncapsulatedFrame{Data: []byte{1, 2}},
	}
	frame.SetTransferSyntaxUID("1.2.840.10008.1.2.4.80")
	if frame.EncapsulatedData.TransferSyntaxUID != "1.2.840.10008.1.2.4.80" {
		t.Errorf("unexpected TransferSyntaxUID on Frame: %s", frame.EncapsulatedData.TransferSyntaxUID)
	}

	if !frame.IsEncapsulated() {
		t.Errorf("frame.IsEncapsulated() returned false")
	}
	ef, err := frame.GetEncapsulatedFrame()
	if err != nil || ef == nil {
		t.Errorf("frame.GetEncapsulatedFrame() failed: %v", err)
	}
	if _, err := frame.GetNativeFrame(); !errors.Is(err, ErrorFrameTypeNotPresent) {
		t.Errorf("expected ErrorFrameTypeNotPresent, got: %v", err)
	}
}

func TestEncapsulatedFrame_InterfaceMethods(t *testing.T) {
	ef := &EncapsulatedFrame{Data: []byte{1, 2}}
	if !ef.IsEncapsulated() {
		t.Errorf("ef.IsEncapsulated() returned false")
	}
	retEf, err := ef.GetEncapsulatedFrame()
	if err != nil || retEf != ef {
		t.Errorf("ef.GetEncapsulatedFrame() failed: %v", err)
	}
	if _, err := ef.GetNativeFrame(); !errors.Is(err, ErrorFrameTypeNotPresent) {
		t.Errorf("expected ErrorFrameTypeNotPresent, got: %v", err)
	}
}

func TestConcurrency_DecoderRegistry(t *testing.T) {
	ResetEncapsulatedDecoders()
	t.Cleanup(ResetEncapsulatedDecoders)

	const goroutines = 40
	const iterations = 50

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(id int) {
			defer wg.Done()
			ts := fmt.Sprintf("1.2.840.10008.1.2.4.%d", 50+(id%10))
			dummy := func(data []byte) (image.Image, error) {
				return image.NewGray(image.Rect(0, 0, 1, 1)), nil
			}

			for i := 0; i < iterations; i++ {
				switch i % 4 {
				case 0:
					RegisterEncapsulatedDecoder(ts, dummy)
				case 1:
					GetEncapsulatedDecoder(ts)
				case 2:
					f := EncapsulatedFrame{
						Data:              []byte{1, 2, 3},
						TransferSyntaxUID: ts,
					}
					// Will either succeed via registered decoder or fail via image/jpeg
					_, _ = f.GetImage()
				case 3:
					UnregisterEncapsulatedDecoder(ts)
				}
			}
		}(g)
	}

	wg.Wait()
}
