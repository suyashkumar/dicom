package dcmtest

import (
	"embed"
)

// IncludedFS uses the embed package to read all of the test files in /pkg/dcmtest/data
// into an embed.FS for testing. This allows us to keep the files in memory without
// reading them from disk each time they are needed, or worrying about relative import
// paths.
//
// WARNING
//
// This var should only be used for testing purposes. It contains a number of full DICOM
// files that could significantly bloat a binary's size if included in production code.
//
//go:embed **/*.dcm
var IncludedFS embed.FS
