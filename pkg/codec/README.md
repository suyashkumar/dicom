# Codec
DICOM codec for decoding JPEG/JPEG-LS/JPEG2000 images, with a focus on providing support for [suyashkumar/dicom](https://github.com/suyashkumar/dicom).

### Prerequisites
To be able to decompress compressed DICOM pixel data, you first have to install the corresponding libraries, that are able to handle the format the data is encoded in:

- LINUX
```bash
apt-get install pkg-config swig libcharls-dev libdcmtk-dev libopenjp2-7-dev
```
- MACOS
```bash
brew install swig pkg-config
brew install dcmtk
brew install openjpeg
brew install team-charls/tap/charls
```

Generate the necessary C++ wrapper code to allow Go to bind to it. You need to do this each time the dependent *.h or *.cxx files change.
```bash
swig -go -cgo -c++ -intgosize 64 codec.i
```

## Usage
```go
import (
	"github.com/suyashkumar/dicom"
	_ "github.com/suyashkumar/dicom/pkg/codec"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func main() {
    dataset, err := dicom.ParseFile("testdata/1.dcm", nil)
    pixelDataElem, err := dataset.FindElementByTag(tag.PixelData)
    pixelData := dicom.MustGetPixelDataInfo(pixelDataElem.Value)
    encapsulatedFrame, err := pixelData.Frames[0].GetEncapsulatedFrame()
    image, err = encapsulatedFrame.GetImage()
    ...
}
```
