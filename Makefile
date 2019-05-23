BINARY = dicomutil

.PHONY: build
build:
	dep ensure
	$(MAKE) test
	go build -o build/${BINARY} ./cmd/dicomutil

.PHONY: build-fast
build-fast:
	go build -o build/${BINARY} ./cmd/dicomutil

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	make build
	./${BINARY}

.PHONY: release
release:
	dep ensure
	$(MAKE) test
	GOOS=linux GOARCH=amd64 go build -o build/${BINARY}-linux-amd64 ./cmd/dicomutil;
	GOOS=darwin GOARCH=amd64 go build -o build/${BINARY}-darwin-amd64 ./cmd/dicomutil;
	GOOS=windows GOARCH=amd64 go build -o build/${BINARY}-windows-amd64.exe ./cmd/dicomutil;
