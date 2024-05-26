package main

import (
	"fmt"
	"os"
)

const tagDictPath = "./dicts.go"

// TagDictCodeWriter writes codegen to dicts.go via CodeWriter interface.
type TagDictCodeWriter struct {
	fileWriter *os.File
}

// Name implements CodeWriter, and returns "tag info dict writer".
func (writer *TagDictCodeWriter) Name() string {
	return "tag info dict writer"
}

// WriteLeading implements CodeWriter, and writes the generic codegen preamble, as well
// as the var declarations for our dicts, and opens declaration of the
// function we will use to init them.
func (writer *TagDictCodeWriter) WriteLeading() error {
	// Set up the two dicts and open the function we will use to init them.
	declareDict := "var tagDict = make(map[Tag]Info)\n"
	declareDict += "var tagDictByKeyword = make(map[string]Info)\n\n"
	initDictOpen := "func initTagDicts() {\n\tvar thisInfo Info\n\n"

	leader := GeneratedFileLeader + declareDict + initDictOpen

	_, err := writer.fileWriter.WriteString(leader)
	return err
}

// Each tag will generate the below code, loading the info into an existing variable
// and then inserting it into both maps.
const tagDeclarationTemplate = `	thisInfo = Info{
		Tag: Tag{0x%04x,0x%04x},
		Name: "%v",
		VR: "%v",
		VM: "%v",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

`

// WriteTag implements CodeWriter, and adds each tag passed to it to each of our dicts.
func (writer *TagDictCodeWriter) WriteTag(info TagInfo) error {
	// Insert our parsed values into the code template.
	addToDicts := fmt.Sprintf(
		tagDeclarationTemplate,
		info.Tag.Group,
		info.Tag.Element,
		info.Name,
		info.VR,
		info.VM,
	)

	_, err := writer.fileWriter.WriteString(addToDicts)
	return err
}

// WriteTrailing implements CodeWriter, and closes our dict init function.
func (writer *TagDictCodeWriter) WriteTrailing() error {
	// Close initDicts() function
	_, err := writer.fileWriter.WriteString("\n}\n")
	return err
}

// Close implements CodeWriter, and closes the dicts.go file.
func (writer *TagDictCodeWriter) Close() error {
	return writer.fileWriter.Close()
}

// NewTagDictCodeWriter creates a new TagDictCodeWriter for writing the dicts.go file.
func NewTagDictCodeWriter() (CodeWriter, error) {
	fileWriter, err := os.Create(tagDictPath)
	if err != nil {
		return nil, fmt.Errorf("error opening '%v': %w", tagDictPath, err)
	}

	return &TagDictCodeWriter{
		fileWriter: fileWriter,
	}, nil
}
