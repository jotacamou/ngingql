VERSION?=    0.0.1
GOARCH?=     amd64
GOOS?=       linux

.PHONY: dep test build clean

all: dep test build

dep:
	dep ensure -v

test:
	go test ./...

build: dep
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -x -v

clean:
	go clean -x -v
	rm -rf vendor/

