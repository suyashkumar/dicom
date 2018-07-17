# DICOM parser in Go
[![Build Status](https://travis-ci.org/gradienthealth/dicom.svg?branch=master)](https://travis-ci.org/gradienthealth/dicom)
[![GoDoc Reference](https://godoc.org/github.com/gradienthealth/dicom?status.svg)](https://godoc.org/github.com/gradienthealth/dicom)

This is a (hard-ish) fork of [go-dicom](https://github.com/gillesdemey/go-dicom)--a golang DICOM image parsing library. A command line tool to parse imagery and data out of DICOM files is also included (`dicomutil`). We have been working on this package with the goal of building a full-featured and high-performance dicom parser with new features and improvements. So far improvements include: 
* parsing and extracting multi-frame DICOM imagery (both encapsulated and native pixel data)
* exposing a `Parser` golang interface to make mock-based testing easier for clients

Upcoming features:
* Channel-based streaming of frames to a client _as they are parsed_ out of the dicom

We're open to suggestions and comments -- open an issue if you have any. 

## Usage
To use this in your golang project, simply import our pacakge `github.com/gradienthealth/go-dicom` and then you can use our `Parser` for your parsing needs:
```go 
p, err := dicom.NewParserFromFile("myfile.dcm", nil)
opts := dicom.ParseOptions{DropPixelData: true}

element := p.ParseNext(opts) // parse and return the next dicom element
// or
dataset, err := p.Parse(opts) // parse whole dicom
```
More details about the package can be found in the [godoc](https://godoc.org/github.com/gradienthealth/go-dicom)

## CLI Tool
A CLI tool that uses this package to parse imagery and metadata out of DICOMs is provided in the `dicomutil` package. 
### Usage
```
dicomutil --extract-images myfile.dcm
```
Note: for some dicoms (with native pixel data) no automatic intensity scaling is applied yet (this is coming). You can apply this in your image viewer if needed. 
### Docker build
To build the tool for all platforms (Mac, Windows, Linux) from source using docker, execute the following in the cloned repo:
```bash
docker build . -t godicom
docker run -it -v $PWD/build:/go/src/github.com/gradienthealth/go-dicom/build godicom make release
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
* Innolitics [DICOM browser](https://dicom.innolitics.com/ciods)
* [DICOM Specification](http://dicom.nema.org/medical/dicom/current/output/pdf/part05.pdf)
