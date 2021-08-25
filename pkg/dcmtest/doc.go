// Package dcmtest contains helper methods and functions for testing against dicom
// files.
//
// WARNING
//
// This module should only be used by test files or test helper packages to be run by
// `gotest`, as it included an embedded dataset which will lead to significant binary
// size bloat when imported into production code.
package dcmtest
