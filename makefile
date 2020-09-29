SHELL := /bin/bash

GOCMD=go
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install

BINARY_NAME=cb-store
BINARY_LINUX=$(BINARY_NAME)
TAGS=$(BINARY_NAME)

.PHONY: all clean build build-linux
all: clean test build

clean:
	@echo "==> Cleaning project"
	rm -f $(BINARY_NAME)

build: clean
	@echo "==> Build project"
	$(GOBUILD) -tags "$(TAGS)" -o $(BINARY_NAME) -v $(BUILD_PATH)

build-linux: clean
	@echo "==> Build (Linux-64) project"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -tags "$(TAGS)" -o $(BINARY_LINUX) -v $(BUILD_PATH)

install:
	@echo "==> Install project"
	$(GOINSTALL) -tags "$(TAGS)"