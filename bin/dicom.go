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
	file   = flag.String("file", "IM-0001-0001.dcm", "the DICOM file you want to parse")
	silent = flag.Bool("silent", false, "wether or not to print all Data Elements")
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
		// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
		table := tw.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)

		fmt.Fprintf(table, "Group\tElement\tName\tVR\tVL\tValue\n")

		for _, elem := range data.Elements {
			fmt.Fprintf(table, "%04X\t%04X\t%s\t%s\t%d\t%v\n", elem.Group, elem.Element, elem.Name, elem.Vr, elem.Vl, elem.Value)
		}

		table.Flush()
	}

}
