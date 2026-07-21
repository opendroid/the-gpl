.PHONY: build install test clean fmt vet lint check

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

fmt:
	gofmt -l .

vet:
	go vet ./...

lint:
	golangci-lint run

test:
	go test ./...

check: fmt vet lint test

clean:
	go clean
	rm -f $(GOBIN)/$(BINARY_NAME)
