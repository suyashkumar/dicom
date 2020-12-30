// Package tag enumerates element tags defined in the DICOM standard.
//
// ftp://medical.nema.org/medical/dicom/2011/11_06pu.pdf
package tag

//go:generate go run ./pkg/tag/codegen
//go:generate stringer -type VRKind

// Parses and populates the tag info dict from serialized spec information.
func init() {
	initTagDicts()
}
