.PHONY: version clean test build clean
# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all:	test
test:
	$(GOTEST) -coverprofile=coverage.out -cover -v ./...
