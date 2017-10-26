/*
 DICOM file parser and creator.

 Reading a DICOM file

   package main

   import (
 	"fmt"
 	"github.com/yasushi-saito/go-dicom"
   )

   func main() {
     data, err := dicom.ReadDataSetFromFile("myfile.dcm", st.Size())
     if err != nil {
         panic(err)
     }
     for _, elem := range data.Elements {
         fmt.Printf("%+v\n", elem)
     }
     elem, err := data.FindElementByTag(dicom.TagPatientName)
     if err != nil {
         panic(err)
     }
     log.Printf("Patient name is: ", elem.MustGetString())
  }

Writing a DICOM file

   package main

   import (
 	"fmt"
 	"github.com/yasushi-saito/go-dicom"
   )

   func main() {
     // The first three elements are mandatory. The remaining elements are optional.
     elems := []*dicom.Element{
           dicom.MustNewElement(dicom.TagTransferSyntax, dicom.ExplicitVRLittleEndian),
           dicom.MustNewElement(dicom.TagMediaStorageSOPClassUID, "1.2.840.10008.5.1.4.1.1.1.2"),
           dicom.MustNewElement(dicom.TagMediaStorageSOPInstanceUID, "1.2.3.4.5.6"),
           dicom.MustNewElement(dicom.TagPatientName, "Alice Doe"),
     }
     ds := dicom.DataSet{Elements: elems}
     elem, err := dicom.WriteDataSetToFile("test.dcm", ds)
     if err != nil {
         panic(err)
     }
  }

*/

package dicom
