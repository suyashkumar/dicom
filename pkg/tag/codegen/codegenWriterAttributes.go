package main

import (
	"fmt"
	"os"
)

const attributesFilePath = "./attributes.go"

// Writes codegen to attributes.go via CodegenWriter interface.
type AttributeCodeWriter struct {
	attributesFileWriter *os.File
}

func (writer *AttributeCodeWriter) Name() string {
	return "attribute var writer"
}

func (writer *AttributeCodeWriter) WriteLeading() error {
	// We just need to write the basics here, package name and generated comment.
	_, err := writer.attributesFileWriter.WriteString(GeneratedFileLeader)
	return err
}

func (writer *AttributeCodeWriter) WriteTag(info TagInfo) error {
	// Each tag is a var declaration on a new line.
	entry := fmt.Sprintf(
		"var %v = Tag{0x%04x,0x%04x}\n",
		info.Name,
		info.Tag.Group,
		info.Tag.Element,
	)
	_, err := writer.attributesFileWriter.WriteString(entry)
	return err
}

func (writer *AttributeCodeWriter) WriteTrailing() error {
	return nil
}

func (writer *AttributeCodeWriter) Close() (err error) {
	return writer.attributesFileWriter.Close()
}

// Create a new CodeWriter for writing the attributes.go file.
func NewAttributesCodeWriter() (CodeWriter, error) {
	fileWriter, err := os.Create(attributesFilePath)
	if err != nil {
		return nil, fmt.Errorf(
			"error file '%v' for write: %w", attributesFilePath, err,
		)
	}

	return &AttributeCodeWriter{
		attributesFileWriter: fileWriter,
	}, nil
}
