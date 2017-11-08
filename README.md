[![GoDoc](https://godoc.org/github.com/grailbio/go-dicom?status.svg)](https://godoc.org/github.com/grailbio/go-dicom) [![Build Status](https://travis-ci.org/grailbio/go-dicom.svg?branch=master)](https://travis-ci.org/grailbio/go-dicom.svg?branch=master)

# DICOM parser in Go



<<<<<<< HEAD
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
=======
This is a fork of github.com/gillesdemey/go-dicom. Changes are:

- Many bug fixes, especially around handling of sequences.
- Handle non-ASCII characters more properly.
- Simplify the API. All the functions are synchronous.
- Better library supports around tags & uids.
- Rudimentary support for writing DICOM files. This part is not complete yet.
- Adds fuzz tests and tests that ensure compatibility with pydicom.

TODO:
- Implement mixed-coding-system files more properly. We currently botch
  patient-name (PN) elements that mixes coding systems.

- A multi-image file. Functionality is almost there, but I haven't had time to complete it.

- Native pixeldata format. It'll be parsed as just []byte.


See doc.go for usage. dicomutil contains a sample program that dumps DICOM
elements in a file.

>>>>>>> 3bad039... Initial commit

### Acknowledgements

I'd like to thank my friend [Seppe Stas](https://github.com/Bitbored/) for helping me get through the horrific DICOM image specification and some of the harder parts of the parser.

Some more inspiration and helpful resource that brought this library to life (in no particular order):

DWV by ivmartel https://github.com/ivmartel/dwv/ <br>
dicomParser by Chris Hafey https://github.com/chafey/dicomParser <br>
http://www.dicomlibrary.com <br>
http://dicom.nema.org/medical/dicom/current/output/pdf/part05.pdf <br>
