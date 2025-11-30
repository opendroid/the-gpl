.PHONY: build install test clean

# Default GOBIN if not set
GOBIN ?= $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN := $(shell go env GOPATH)/bin
endif

BINARY_NAME=the-gpl

build:
	go build -o $(GOBIN)/$(BINARY_NAME) main.go

install:
	go install .

test:
	go test ./...

clean:
	go clean
	rm -f $(GOBIN)/$(BINARY_NAME)
