package main

import (
	"flag"
	"fmt"
	"github.com/gillesdemey/go-dicom"
	"io/ioutil"
	"os"
	tw "text/tabwriter"
)

var (
	file   = flag.String("file", "", "the DICOM file you want to parse")
	silent = flag.Bool("silent", false, "wether or not to print all Data Elements")
	out    = flag.String("out", "", "where to write the program's output")
)

func init() {
	flag.Parse()
}

func main() {

	parser, err := dicom.NewParser()
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(*file)

	if err != nil {
		panic(err)
	}

	data, err := parser.Parse(bytes)

	if err != nil {
		fmt.Println(err)
	}

	if *silent == false {

		var writer *os.File

		if *out != "" {
			writer, _ = os.Create(*out)
		} else {
			writer = os.Stdout
		}

		// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
		table := tw.NewWriter(writer, 0, 8, 0, '\t', 0)

		fmt.Fprintf(table, "Tag\tVR\tValue\tVL\tName\n")

		for _, elem := range data.Elements {
			fmt.Fprintf(table, "(%04X,%04X)\t%s\t%v\t%d\t%s\n", elem.Group, elem.Element, elem.Vr, elem.Value, elem.Vl, elem.Name)
		}

		table.Flush()
	}

}
