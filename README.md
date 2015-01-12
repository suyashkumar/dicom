# DICOM parser in Go

## Usage [![GoDoc](https://godoc.org/github.com/gillesdemey/go-dicom?status.svg)](https://godoc.org/github.com/gillesdemey/go-dicom)
```Go
package main

import (
	"fmt"
	"github.com/gillesdemey/go-dicom"
	"io/ioutil"
)

func main() {

	file, _ := ioutil.ReadFile("examples/IM-0001-0001.dcm")
	data, _ := dicom.Parse(file)

	if err != nil {
		fmt.Println(err)
	}

	for _, elem := range data.Elements {
		fmt.Printf("%+v\n", &elem)
	}

}
```

Will print something like:

```
&{Name:FileMetaInformationGroupLength Group:0002 Vr:UL Vl:4 Element:0000 Value:204}
&{Name:FileMetaInformationVersion Group:0002 Vr:OB Vl:2 Element:0001 Value:[0 1]}
&{Name:MediaStorageSOPClassUID Group:0002 Vr:UI Vl:26 Element:0002 Value:[1.2.840.10008.5.1.4.1.1.2]}
&{Name:MediaStorageSOPInstanceUID Group:0002 Vr:UI Vl:56 Element:0003 Value:[1.3.12.2.1107.5.1.4.54023.30000005032916373504600004748]}
&{Name:TransferSyntaxUID Group:0002 Vr:UI Vl:22 Element:0010 Value:[1.2.840.10008.1.2.4.91]}
&{Name:ImplementationClassUID Group:0002 Vr:UI Vl:22 Element:0012 Value:[1.3.6.1.4.1.19291.2.1]}
&{Name:ImplementationVersionName Group:0002 Vr:SH Vl:10 Element:0013 Value:[OSIRIX001]}
&{Name:SourceApplicationEntityTitle Group:0002 Vr:AE Vl:6 Element:0016 Value:[OsiriX]}
```
