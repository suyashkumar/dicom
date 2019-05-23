package frame

// NativeFrame represents a native image frame
type NativeFrame struct {
	// Data is a slice of pixels, where each pixel can have multiple values
	Data          [][]int
	Rows          int
	Cols          int
	BitsPerSample int
}

// EncapsulatedFrame represents an encapsulated image frame
type EncapsulatedFrame struct {
	// Data is a collection of bytes representing a JPEG encoded image frame
	Data []byte
}

// Frame wraps a single encapsulated or native image frame
type Frame struct {
	IsEncapsulated   bool
	EncapsulatedData EncapsulatedFrame
	NativeData       NativeFrame
}
