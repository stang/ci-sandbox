VERSION ?= $(shell git describe --match "v*.*" --abbrev=7 --tags --dirty)
BINARY_NAME := hello
BUILD_ARGS := -ldflags "-X main.version=${VERSION}"

all: test build

.PHONY: test
test:
	go test -count=1 -race -coverprofile=coverage.txt -covermode=atomic ./...

build:
	mkdir -p bin/
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${BUILD_ARGS} -o bin/${BINARY_NAME}-darwin-amd64 cmd/${BINARY_NAME}/*.go

version:
	@echo "$(VERSION)"
