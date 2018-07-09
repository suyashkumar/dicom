FROM golang:alpine
WORKDIR /go/src/github.com/gradienthealth/go-dicom
COPY . .
RUN apk add --no-cache make git
RUN go get github.com/golang/dep/cmd/dep
RUN make