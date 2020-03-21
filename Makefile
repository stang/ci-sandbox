VERSION ?= $(shell git describe --match "v*.*" --abbrev=7 --tags --dirty)
PKGS := $(wildcard ./internal/*)
BINARY := hello
BUILD_ARGS := -ldflags "-X main.version=${VERSION}"

all: test build

.PHONY: $(PKGS)
$(PKGS):
	go test -count=1 "./$@"

test: $(PKGS)

build:
	mkdir -p bin/
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${BUILD_ARGS} -o bin/${BINARY}-darwin-amd64 cmd/${BINARY}/*.go

version:
	@echo "$(VERSION)"
