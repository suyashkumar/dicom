# Test files Details

This document contains citations and further details for the test DICOMs used here.


## Generated Files

The sub-bullets mention potentially interesting characteristics of the test file.
Some of the interesting characteristics may apply to more than one file, but may only
be mentioned in one of them for brevity.

* [1.dcm.1.raw](1.dcm.1.raw) is generated from https://github.com/suyashkumar/dicom/testdata/1.dcm by running the following commands:
    ```sh
    dcmdump --write-pixel ../pkg/codec/testdata/ 1.dcm
    ```
* [1_jpls.dcm.1.raw](1_jpls.dcm.1.raw) is generated from https://github.com/suyashkumar/dicom/testdata/1.dcm by running the following commands:
    ```sh
    dcmcjpls 1.dcm 1_jpls.dcm
    dcmdump --write-pixel ../pkg/codec/testdata/ 1_jpls.dcm
    ```
* [3.dcm.0.raw](3.dcm.0.raw) is generated from https://github.com/suyashkumar/dicom/testdata/3.dcm by running the following commands:
    ```sh
    dcmdump --write-pixel ../pkg/codec/testdata/ 3.dcm
    ```
* [3_jpeg.dcm.1.raw](3_jpeg.dcm.1.raw) is generated from https://github.com/suyashkumar/dicom/testdata/3.dcm by running the following commands:
    ```sh
    dcmcjpeg 3.dcm 3_jpeg.dcm
    dcmdump --write-pixel ../pkg/codec/testdata/ 3_jpeg.dcm
    ```

#### Files a1_mono.j2c and a1_mono.ppm
These files were sourced from [openjpeg-data](https://github.com/uclouvain/openjpeg-data)
