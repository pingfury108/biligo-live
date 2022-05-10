
VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

go-build:
	mkdir -p $(GOBIN)
	go build  -o $(GOBIN)/$(PROJECTNAME) ./cmd/biligo-live

clean:
	rm $(GOBIN)/$(PROJECTNAME)
