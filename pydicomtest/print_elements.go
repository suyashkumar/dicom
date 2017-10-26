package main

// Print the contents of a file in a format used by pydicom.  Used to ensure
// that pydicom and go-dicom parses files in the same way.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"github.com/grailbio/go-dicom"
	"github.com/grailbio/go-dicom/dicomuid"
)

var (
	printMetadata = flag.Bool("print-metadata", true, "Print image metadata")
)

// Sorter
type elemSorter struct {
	elems []*dicom.Element
}

func (e *elemSorter) Len() int {
	return len(e.elems)
}

func (e *elemSorter) Swap(i, j int) {
	tmp := e.elems[i]
	e.elems[i] = e.elems[j]
	e.elems[j] = tmp
}

func (e *elemSorter) Less(i, j int) bool {
	elemi := e.elems[i]
	elemj := e.elems[j]
	if elemi.Tag.Group < elemj.Tag.Group {
		return true
	}
	if elemi.Tag.Group > elemj.Tag.Group {
		return false
	}
	return elemi.Tag.Element < elemj.Tag.Element
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Panic("print_elements_test <dicomfile>")
	}
	path := flag.Arg(0)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data, err := dicom.ReadDataSetInBytes(bytes, dicom.ReadOptions{})
	if err != nil {
		panic(err)
	}

	printElements(data.Elements, 0)
}

func printScalar(vr string, i interface{}, indent int) string {
	var s string
	switch v := i.(type) {
	case float32:
		s = fmt.Sprintf("%.4f", v)
	case float64:
		s = fmt.Sprintf("%.4f", v)
	case string:
		if vr == "UI" {
			// Resolve UID
			e, err := dicomuid.Lookup(v)
			if err == nil {
				v = e.Name
			}
		}
		s = fmt.Sprintf("%s", v)
	case dicom.Tag:
		return v.String()

	default:
		s = fmt.Sprintf("%v", i)
	}
	return s
}

func printTag(tag dicom.Tag) string {
	return fmt.Sprintf("(%04x, %04x)", tag.Group, tag.Element)
}

func printElement(elem *dicom.Element, indent int) {
	if elem.Tag.Group == 2 {
		// Don't print the meta elements found in the beginning of the
		// file. Pydicom doesn't for some reason.
		return
	}

	fmt.Printf("%s%s %s:", strings.Repeat(" ", indent*2), printTag(elem.Tag), elem.VR)

	if elem.Tag == dicom.TagPixelData {
		// PixelData encoding differs between godicom and pydicom.  Skip
		// for now.  TODO(saito) fix.
		fmt.Print(" [omitted]\n")
	} else if elem.VR == "OW" || elem.VR == "OB" || elem.VR == "OD" || elem.VR == "OF" {
		if len(elem.Value) != 1 {
			fmt.Printf(" [%d values]\n", len(elem.Value))
		} else if v, ok := elem.Value[0].([]byte); ok {
			fmt.Printf(" %dB\n", len(v))
		} else {
			v := elem.Value[0].(string)
			fmt.Printf(" %dB\n", len(v))
		}
	} else if elem.VR == "LT" {
		// pydicom trims one (but not more) trailing space from the
		// string.  Whereas go-dicom doesn't trim anything. The spec
		// says an impl "*may* trim trailing space*s* from the string".
		// So the behavior of pydicom is a bit strange, but follows the
		// wording.
		v := elem.Value[0].(string)
		n := len(v)
		if strings.HasSuffix(v, " ") {
			n--
		}
		fmt.Printf(" %dB\n", n)
	} else if elem.VR == "SQ" {
		var childElems []*dicom.Element
		if len(elem.Value) == 1 {
			// If SQ contains one Item, unwrap the item.
			items := elem.Value[0].(*dicom.Element)
			if items.Tag != dicom.TagItem {
				log.Panicf("A SQ item must be of type Item, but found %v", items)
			}
			for _, item := range items.Value {
				childElems = append(childElems, item.(*dicom.Element))
			}
		} else {
			for _, v := range elem.Value {
				child := v.(*dicom.Element)
				if child.Tag != dicom.TagItem {
					log.Panicf("A SQ item must be of type Item, but found %v", child)
				}
				childElems = append(childElems, child)
			}
		}
		fmt.Print("\n")
		printElements(childElems, indent+1)
	} else if elem.VR == "NA" {
		var childElems []*dicom.Element
		if len(elem.Value) == 1 {
			items := elem.Value[0].(*dicom.Element)
			for _, item := range items.Value {
				childElems = append(childElems, item.(*dicom.Element))
			}
		} else {
			for _, v := range elem.Value {
				child := v.(*dicom.Element)
				childElems = append(childElems, child)
			}
		}
		fmt.Print("\n")
		printElements(childElems, indent+1)
	} else { // Other scalar value
		if len(elem.Value) == 0 {
			fmt.Print("\n")
		} else if len(elem.Value) == 1 {
			fmt.Printf(" %s\n", printScalar(elem.VR, elem.Value[0], 1))
		} else {
			fmt.Print(" [")
			for i, value := range elem.Value {
				if i > 0 {
					fmt.Print(", ")
				}
				// Follow the pydicom's printing format.  It
				// encloses the value in quotes only at the
				// toplevel.
				fmt.Print(printScalar(elem.VR, value, indent))
			}
			fmt.Print("]\n")
		}
	}
}

func printElements(elems []*dicom.Element, indent int) {
	sort.Sort(&elemSorter{elems: elems})
	for _, elem := range elems {
		printElement(elem, indent)
	}
}
