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

	bytes, err := ioutil.ReadFile("myfile.dcm")
	
	parser, err := dicom.NewParser()
	data, err := parser.Parse(bytes)

	for _, elem := range data.Elements {
		fmt.Printf("%+v\n", &elem)
	}

}
```

## Commandline Interface

`dicom -file=myfile.dcm`

Will print something like:

```
Group	Element	Name				VR	VL	Value
0002	0000	FileMetaInformationGroupLength	UL	4	204
0002	0001	FileMetaInformationVersion	OB	2	[0 1]
0002	0002	MediaStorageSOPClassUID		UI	26	[1.2.840.10008.5.1.4.1.1.2]
0002	0003	MediaStorageSOPInstanceUID	UI	56	[1.3.12.2.1107.5.1.4.54023.30000005032916373504600004748]
0002	0010	TransferSyntaxUID		UI	22	[1.2.840.10008.1.2.4.91]
0002	0012	ImplementationClassUID		UI	22	[1.3.6.1.4.1.19291.2.1]
0002	0013	ImplementationVersionName	SH	10	[OSIRIX001]
0002	0016	SourceApplicationEntityTitle	AE	6	[OsiriX]
0008	0008	ImageType			CS	34	[ORIGINAL PRIMARY AXIAL CT_SOM5 SPI]
0008	0016	SOPClassUID			UI	26	[1.2.840.10008.5.1.4.1.1.2]
0008	0018	SOPInstanceUID			UI	56	[1.3.12.2.1107.5.1.4.54023.30000005032916373504600004748]
0008	0020	StudyDate			DA	8	[20050329]
0008	0021	SeriesDate			DA	8	[20050329]
0008	0022	AcquisitionDate			DA	8	[20050329]
0008	0023	ContentDate			DA	8	[20050329]
0008	0030	StudyTime			TM	14	[142530.125000 ]
0008	0031	SeriesTime			TM	14	[144801.203000 ]
0008	0032	AcquisitionTime			TM	14	[143840.611848 ]
0008	0033	ContentTime			TM	14	[143840.611848 ]
0008	0050	AccessionNumber			SH	8	[2386679]
0008	0060	Modality			CS	2	[CT]
0008	0070	Manufacturer			LO	8	[SIEMENS]
0008	0080	InstitutionName			LO	20	[UCLA  Medical Center]
0008	0081	InstitutionAddress		ST	52	[UCLA Medical PlazaLos Angeles/2782F4/Los AngelesUSA]
0008	0090	ReferringPhysicianName		PN	16	[MIYAMOTO^MICHAEL]
0008	1010	StationName			SH	8	[CT54023]
0008	1030	StudyDescription		LO	48	[Cardiac^1CTA_CORONARY_ARTERIES_TESTBOLUS (Adult)]
```

### Acknowledgements

I'd like to thank my friend [Seppe Stas](https://github.com/Bitbored/) for helping me get through the horrific DICOM image specification and some of the harder parts of the parser.

Some more inspiration and helpful resource that brought this library to life (in no particular order):

DWV by ivmartel https://github.com/ivmartel/dwv/ <br>
dicomParser by Chris Hafey https://github.com/chafey/dicomParser <br>
http://www.dicomlibrary.com <br>
http://dicom.nema.org/medical/dicom/current/output/pdf/part05.pdf <br>
