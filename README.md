# DICOM parser in Go [![GoDoc](https://godoc.org/github.com/gillesdemey/go-dicom?status.svg)](https://godoc.org/github.com/gillesdemey/go-dicom) [![wercker status](https://app.wercker.com/status/c250d72bc82a5d8f267c7ee0b9e839bc/s/master "wercker status")](https://app.wercker.com/project/bykey/c250d72bc82a5d8f267c7ee0b9e839bc)

## Usage
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
...
&{Name:RequestedProcedureID Group:0040 Element:1001 Vr:SH Vl:8 Value:[2386679]}
&{Name:PixelData Group:7FE0 Element:0010 Vr:OW Vl:0 Value:[]}
&{Name:Item Group:FFFE Element:E000 Vr:na Vl:4 Value:[0 0 0 0]}
```

### Acknowledgements

I'd like to thank my friend [Seppe Stas](https://github.com/Bitbored/) for helping me get through the horrific DICOM image specification and some of the harder parts of the parser.

Some more inspiration and helpful resource that brought this library to life (in no particular order):

DWV by ivmartel https://github.com/ivmartel/dwv/ <br>
dicomParser by Chris Hafey https://github.com/chafey/dicomParser <br>
http://www.dicomlibrary.com <br>
http://dicom.nema.org/medical/dicom/current/output/pdf/part05.pdf <br>
