package frame

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"strings"
	"sync"

	"github.com/suyashkumar/dicom/pkg/uid"
)

// ErrUnsupportedEncapsulatedCodec is a sentinel error indicating that the encapsulated frame
// uses a compression format or JPEG marker unsupported by the standard Go image/jpeg decoder.
var ErrUnsupportedEncapsulatedCodec = errors.New("unsupported encapsulated codec")

// UnsupportedCodecError provides diagnostic details when an encapsulated frame cannot be decoded.
type UnsupportedCodecError struct {
	// TransferSyntaxUID is the DICOM transfer syntax UID if known or detected.
	TransferSyntaxUID string
	// Marker is the JPEG SOF marker code if detected (e.g. 0xFFC3 for SOF3).
	Marker uint16
	// Format is a human-readable name of the format (e.g. "JPEG Lossless (SOF3)").
	Format string
	// Err is the underlying decoder error from image/jpeg.
	Err error
}

func (e *UnsupportedCodecError) Error() string {
	var b strings.Builder
	if e.TransferSyntaxUID != "" {
		if e.Format != "" {
			fmt.Fprintf(&b, "unsupported transfer syntax %s (%s): ", e.TransferSyntaxUID, e.Format)
		} else {
			fmt.Fprintf(&b, "unsupported transfer syntax %s: ", e.TransferSyntaxUID)
		}
	} else if e.Format != "" {
		fmt.Fprintf(&b, "unsupported image format (%s): ", e.Format)
	} else if e.Marker != 0 {
		fmt.Fprintf(&b, "unsupported JPEG feature: marker 0x%04X: ", e.Marker)
	} else {
		b.WriteString("unsupported encapsulated image format: ")
	}

	b.WriteString("format is not supported by standard image/jpeg; ")
	if e.TransferSyntaxUID != "" {
		fmt.Fprintf(&b, "register a decoder via RegisterEncapsulatedDecoder for transfer syntax %s or access EncapsulatedFrame.Data directly", e.TransferSyntaxUID)
	} else {
		b.WriteString("register a decoder via RegisterEncapsulatedDecoder or access EncapsulatedFrame.Data directly")
	}
	return b.String()
}

// Unwrap returns the underlying decoder error.
func (e *UnsupportedCodecError) Unwrap() error {
	return e.Err
}

// Is reports whether this error matches target.
func (e *UnsupportedCodecError) Is(target error) bool {
	return target == ErrUnsupportedEncapsulatedCodec
}

// EncapsulatedDecoder defines a function type capable of decoding encapsulated frame bytes into an image.Image.
type EncapsulatedDecoder func(data []byte) (image.Image, error)

var (
	decoderMu sync.RWMutex
	decoders  = make(map[string]EncapsulatedDecoder)
)

func normalizeUID(uidStr string) string {
	cleaned := strings.Trim(uidStr, " \t\r\n\x00")
	cleaned = strings.TrimPrefix(cleaned, "=")
	if info, err := uid.Lookup(cleaned); err == nil && info.UID != "" {
		return info.UID
	}
	return cleaned
}

// RegisterEncapsulatedDecoder registers a custom decoder for the specified DICOM transfer syntax UID.
func RegisterEncapsulatedDecoder(transferSyntaxUID string, decoder EncapsulatedDecoder) {
	key := normalizeUID(transferSyntaxUID)
	decoderMu.Lock()
	defer decoderMu.Unlock()
	if decoder == nil {
		delete(decoders, key)
		return
	}
	decoders[key] = decoder
}

// GetEncapsulatedDecoder retrieves a registered decoder for the specified DICOM transfer syntax UID.
func GetEncapsulatedDecoder(transferSyntaxUID string) (EncapsulatedDecoder, bool) {
	key := normalizeUID(transferSyntaxUID)
	decoderMu.RLock()
	defer decoderMu.RUnlock()
	decoder, ok := decoders[key]
	return decoder, ok
}

// UnregisterEncapsulatedDecoder removes any registered decoder for the specified DICOM transfer syntax UID.
func UnregisterEncapsulatedDecoder(transferSyntaxUID string) {
	key := normalizeUID(transferSyntaxUID)
	decoderMu.Lock()
	defer decoderMu.Unlock()
	delete(decoders, key)
}

// ResetEncapsulatedDecoders removes all registered encapsulated decoders.
func ResetEncapsulatedDecoders() {
	decoderMu.Lock()
	defer decoderMu.Unlock()
	decoders = make(map[string]EncapsulatedDecoder)
}

// EncapsulatedFrame represents an encapsulated image frame
type EncapsulatedFrame struct {
	// Data is a collection of bytes representing an encapsulated image frame
	Data []byte
	// TransferSyntaxUID represents the DICOM transfer syntax UID for this frame, if known.
	TransferSyntaxUID string
}

// IsEncapsulated indicates if the frame is encapsulated or not.
func (e *EncapsulatedFrame) IsEncapsulated() bool { return true }

// GetEncapsulatedFrame returns an EncapsulatedFrame from this frame.
func (e *EncapsulatedFrame) GetEncapsulatedFrame() (*EncapsulatedFrame, error) {
	return e, nil
}

// GetNativeFrame returns ErrorFrameTypeNotPresent, because this struct does
// not hold a NativeFrame.
func (e *EncapsulatedFrame) GetNativeFrame() (INativeFrame, error) {
	return nil, ErrorFrameTypeNotPresent
}

// SetTransferSyntaxUID sets the transfer syntax UID on this encapsulated frame.
func (e *EncapsulatedFrame) SetTransferSyntaxUID(uid string) {
	e.TransferSyntaxUID = uid
}

// GetImage returns a Go image.Image from the underlying frame.
func (e *EncapsulatedFrame) GetImage() (image.Image, error) {
	if len(e.Data) == 0 {
		return nil, errors.New("cannot decode empty encapsulated frame data")
	}

	// 1. Look up registered decoder by explicit TransferSyntaxUID if present.
	if e.TransferSyntaxUID != "" {
		if decoder, ok := GetEncapsulatedDecoder(e.TransferSyntaxUID); ok {
			return decoder(e.Data)
		}
	}

	// 2. Inspect stream to identify format or markers.
	detectedTS, marker, isFormatDetected := inspectStream(e.Data)

	// If detected transfer syntax has a registered decoder, invoke it.
	if detectedTS != "" && detectedTS != e.TransferSyntaxUID {
		if decoder, ok := GetEncapsulatedDecoder(detectedTS); ok {
			return decoder(e.Data)
		}
	}

	// For SOF3, check alternative lossless transfer syntax 1.2.840.10008.1.2.4.57
	if marker == 0xFFC3 {
		if decoder, ok := GetEncapsulatedDecoder("1.2.840.10008.1.2.4.57"); ok {
			return decoder(e.Data)
		}
	}

	// 3. Fall back to standard jpeg.Decode.
	img, err := jpeg.Decode(bytes.NewReader(e.Data))
	if err == nil {
		return img, nil
	}

	// 4. On failure, check if format is unsupported or TransferSyntaxUID is unsupported.
	if isFormatDetected {
		return nil, newUnsupportedCodecError(e.TransferSyntaxUID, detectedTS, marker, err)
	}

	if isKnownUnsupportedTransferSyntax(e.TransferSyntaxUID) {
		return nil, newUnsupportedCodecError(e.TransferSyntaxUID, e.TransferSyntaxUID, marker, err)
	}

	// Return standard decoder error for malformed or corrupted data.
	return nil, err
}

// Equals returns true if this frame equals the provided target frame, otherwise
// false.
func (e *EncapsulatedFrame) Equals(target *EncapsulatedFrame) bool {
	if target == nil || e == nil {
		return e == target
	}
	if e.TransferSyntaxUID != target.TransferSyntaxUID {
		return false
	}
	return bytes.Equal(e.Data, target.Data)
}

func scanJPEGMarkers(data []byte) (uint16, bool) {
	if len(data) < 4 || data[0] != 0xFF || data[1] != 0xD8 {
		return 0, false
	}
	i := 2
	for i < len(data) {
		if data[i] != 0xFF {
			i++
			continue
		}
		for i < len(data) && data[i] == 0xFF {
			i++
		}
		if i >= len(data) {
			break
		}
		markerByte := data[i]
		i++

		if markerByte == 0x00 {
			continue
		}

		if (markerByte >= 0xD0 && markerByte <= 0xD7) || markerByte == 0xD8 || markerByte == 0x01 {
			continue
		}

		if markerByte == 0xD9 {
			break
		}

		fullMarker := 0xFF00 | uint16(markerByte)

		switch markerByte {
		case 0xC0, 0xC1, 0xC2, 0xC3, 0xC5, 0xC6, 0xC7,
			0xC9, 0xCA, 0xCB, 0xCD, 0xCE, 0xCF, 0xF7:
			return fullMarker, true
		}

		if markerByte == 0xDA {
			break
		}

		if i+2 > len(data) {
			break
		}
		length := int(binary.BigEndian.Uint16(data[i : i+2]))
		if length < 2 {
			break
		}
		i += length
	}
	return 0, false
}

func isJPEG2000(data []byte) bool {
	if len(data) >= 8 && bytes.Equal(data[:8], []byte{0x00, 0x00, 0x00, 0x0C, 0x6A, 0x50, 0x20, 0x20}) {
		return true
	}
	if len(data) >= 4 && data[0] == 0xFF && data[1] == 0x4F && data[2] == 0xFF {
		return true
	}
	return false
}

func inspectStream(data []byte) (detectedSyntax string, marker uint16, isFormatDetected bool) {
	if isJPEG2000(data) {
		return "1.2.840.10008.1.2.4.90", 0, true
	}
	m, found := scanJPEGMarkers(data)
	if !found {
		return "", 0, false
	}
	switch m {
	case 0xFFC0, 0xFFC1, 0xFFC2:
		return "1.2.840.10008.1.2.4.50", m, false
	case 0xFFC3:
		return "1.2.840.10008.1.2.4.70", m, true
	case 0xFFF7:
		return "1.2.840.10008.1.2.4.80", m, true
	default:
		return "", m, true
	}
}

func isKnownUnsupportedTransferSyntax(tsUID string) bool {
	switch normalizeUID(tsUID) {
	case "1.2.840.10008.1.2.4.70",
		"1.2.840.10008.1.2.4.57",
		"1.2.840.10008.1.2.4.65",
		"1.2.840.10008.1.2.4.66",
		"1.2.840.10008.1.2.4.80",
		"1.2.840.10008.1.2.4.81",
		"1.2.840.10008.1.2.4.90",
		"1.2.840.10008.1.2.4.91",
		"1.2.840.10008.1.2.4.92",
		"1.2.840.10008.1.2.4.93",
		"1.2.840.10008.1.2.5":
		return true
	default:
		return false
	}
}

func formatNameForMarker(marker uint16) string {
	switch marker {
	case 0xFFC3:
		return "JPEG Lossless (SOF3 marker 0xFFC3)"
	case 0xFFF7:
		return "JPEG-LS (SOF55 marker 0xFFF7)"
	case 0xFFC5:
		return "JPEG Differential Sequential DCT (SOF5 marker 0xFFC5)"
	case 0xFFC6:
		return "JPEG Differential Progressive DCT (SOF6 marker 0xFFC6)"
	case 0xFFC7:
		return "JPEG Differential Lossless (SOF7 marker 0xFFC7)"
	case 0xFFC9:
		return "JPEG Extended Sequential DCT Arithmetic Coding (SOF9 marker 0xFFC9)"
	case 0xFFCA:
		return "JPEG Progressive DCT Arithmetic Coding (SOF10 marker 0xFFCA)"
	case 0xFFCB:
		return "JPEG Lossless Arithmetic Coding (SOF11 marker 0xFFCB)"
	case 0xFFCD:
		return "JPEG Differential Sequential DCT Arithmetic Coding (SOF13 marker 0xFFCD)"
	case 0xFFCE:
		return "JPEG Differential Progressive DCT Arithmetic Coding (SOF14 marker 0xFFCE)"
	case 0xFFCF:
		return "JPEG Differential Lossless Arithmetic Coding (SOF15 marker 0xFFCF)"
	default:
		if marker != 0 {
			return fmt.Sprintf("JPEG marker 0x%04X", marker)
		}
		return ""
	}
}

func formatNameForSyntax(tsUID string) string {
	switch normalizeUID(tsUID) {
	case "1.2.840.10008.1.2.4.70":
		return "JPEG Lossless, First-Order Prediction"
	case "1.2.840.10008.1.2.4.57":
		return "JPEG Lossless, Non-Hierarchical (Process 14)"
	case "1.2.840.10008.1.2.4.80":
		return "JPEG-LS Lossless"
	case "1.2.840.10008.1.2.4.81":
		return "JPEG-LS Lossy"
	case "1.2.840.10008.1.2.4.90":
		return "JPEG 2000 (Lossless Only)"
	case "1.2.840.10008.1.2.4.91":
		return "JPEG 2000"
	case "1.2.840.10008.1.2.5":
		return "RLE Lossless"
	default:
		return ""
	}
}

func newUnsupportedCodecError(frameTS, detectedTS string, marker uint16, err error) *UnsupportedCodecError {
	tsUID := frameTS
	if tsUID == "" {
		tsUID = detectedTS
	}

	formatName := formatNameForMarker(marker)
	if formatName == "" && tsUID != "" {
		formatName = formatNameForSyntax(tsUID)
	}
	if formatName == "" {
		formatName = "unsupported format"
	}

	return &UnsupportedCodecError{
		TransferSyntaxUID: tsUID,
		Marker:            marker,
		Format:            formatName,
		Err:               err,
	}
}
