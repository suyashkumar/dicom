<p align="center">
  <img src="https://suyashkumar.com/assets/img/magnetic-resonance.png" width="125px"/>
  <h3 align="center">dicom</h3>
  <p align="center">High Performance Golang DICOM Medical Image Parser<p>
  <p align="center"> <a href="https://travis-ci.org/suyashkumar/dicom"><img src="https://travis-ci.org/suyashkumar/dicom.svg?branch=master" /></a> <a href="https://godoc.org/github.com/suyashkumar/dicom"><img src="https://godoc.org/github.com/suyashkumar/dicom?status.svg" alt=""></a> 
  </p>
</p>

This is a (hard-ish) fork of work I did at [gradienthealth](https://github.com/gradienthealth/dicom) which built on top of [go-dicom](https://github.com/gillesdemey/go-dicom)--a golang DICOM image parsing library and command line tool. We have been working on this package with the goal of building a full-featured and high-performance dicom parser in Golang with new features and improvements. I will continue to add some (potentially API breaking) improvements on my repository fork here.

So far, improvements that have made on top of [go-dicom](https://github.com/gillesdemey/go-dicom) include: 
- [x] parsing and extracting multi-frame DICOM imagery (both encapsulated and native pixel data)
- [x] exposing a `Parser` golang interface to make mock-based testing easier for clients
- [x] Channel-based streaming of `Frame`s to a client _as they are parsed_ out of the dicom
- [x] Parsing performance improvements 
- [x] General refactors to the [go-dicom](https://github.com/gillesdemey/go-dicom) code (though there's more work to be done here) for maintainability an readability. 

## Usage
To use this in your golang project, import `github.com/suyashkumar/dicom` and then you can use `dicom.Parser` for your parsing needs:
```go 
p, err := dicom.NewParserFromFile("myfile.dcm", nil)
opts := dicom.ParseOptions{DropPixelData: true}

element := p.ParseNext(opts) // parse and return the next dicom element
// or
dataset, err := p.Parse(opts) // parse whole dicom
```
More details about the package can be found in the [godoc](https://godoc.org/github.com/suyashkumar/dicom)

## CLI Tool
A CLI tool that uses this package to parse imagery and metadata out of DICOMs is provided in the `dicomutil` package. All dicom tags present are printed to STDOUT by default. 

### Installation
You can download the prebuilt binaries from the [releases tab](https://github.com/suyashkumar/dicom/releases), or use the following to download the binary at the command line using my [getbin tool](https://github.com/suyashkumar/getbin):

```sh
wget -qO- "https://getbin.io/suyashkumar/dicom" | tar xvz
```
(This attempts to infer your OS and 301 redirects `wget` to the latest github release asset for your system. Downloads come from GitHub releases).
### Usage
```
dicomutil --extract-images-stream myfile.dcm
```
Note: for some dicoms (with native pixel data) no automatic intensity scaling is applied yet (this is coming). You can apply this in your image viewer if needed (in Preview on mac, go to Tools->Adjust Color). 
### Docker build
To build the tool for all platforms (Mac, Windows, Linux) from source using docker, execute the following in the cloned repo:
```bash
docker build . -t godicom
docker run -it -v $PWD/build:/go/src/github.com/suyashkumar/dicom/build godicom make release
```
You can then use the binaries that will show up in the `build` folder in your current working directory
### Build manually
To build manually, ensure you have `make`, golang, and [dep](https://github.com/golang/dep) installed on your machine. Clone (or `go get`) this repo into your gopath then:
```
make
```

## Acknowledgements

* Original [go-dicom](https://github.com/gillesdemey/go-dicom)
* Grailbio [go-dicom](https://github.com/grailbio/go-dicom) -- commits from their fork were applied to ours
* GradientHealth for supporting work I did on this while there [gradienthealth/dicom](https://github.com/gradienthealth/dicom)
* Innolitics [DICOM browser](https://dicom.innolitics.com/ciods)
* [DICOM Specification](http://dicom.nema.org/medical/dicom/current/output/pdf/part05.pdf)
* <div>Icons made by <a href="https://www.freepik.com/?__hstc=57440181.48e262e7f01bcb2b41259e2e5a8103b3.1557697512782.1557697512782.1557697512782.1&__hssc=57440181.4.1557697512783&__hsfp=2768524783" title="Freepik">Freepik</a> from <a href="https://www.flaticon.com/" 			    title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" 			    title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a></div>
