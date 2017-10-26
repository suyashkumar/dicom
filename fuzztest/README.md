This directory contains fuzz tests. It uses go-fuzz. Visit
https://github.com/dvyukov/go-fuzz and install two packages, go-fuzz-build and go-fuzz.
Then run:

```
go-fuzz-build github.com/yasushi-saito/go-dicom/fuzztest
mkdir -p /tmp/fuzz/corpus
cp -r ../examples/* /tmp/fuzz/corpus/
go-fuzz -bin fuzz-fuzz.zip -workdir /tmp/fuzz
```
