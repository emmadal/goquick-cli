# GoQuick CLI Makefile

BINARY_NAME=quick
VERSION=1.0.0
BUILD_DIR=build
MAIN_PACKAGE=./cmd/quick
LDFLAGS=-ldflags "-w -s -X main.VERSION=$(VERSION)"

# Detect the operating system and architecture
OS=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)

.PHONY: all build clean run test install release-all release-linux release-darwin release-windows

all: clean build

build:
	mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

run:
	go run $(MAIN_PACKAGE)

clean:
	rm -rf $(BUILD_DIR)
	go clean

test:
	go test ./...

install:
	go install $(LDFLAGS) $(MAIN_PACKAGE)

# Release builds for all platforms
release-all: release-linux release-darwin release-windows

# Build for Linux
release-linux:
	mkdir -p $(BUILD_DIR)/linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/linux/$(BINARY_NAME) $(MAIN_PACKAGE)

# Build for MacOS
release-darwin:
	mkdir -p $(BUILD_DIR)/darwin
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/darwin/$(BINARY_NAME) $(MAIN_PACKAGE)

# Build for Windows
release-windows:
	mkdir -p $(BUILD_DIR)/windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/windows/$(BINARY_NAME).exe $(MAIN_PACKAGE)

# Default target
help:
	@echo "GoQuick CLI Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build       - Build the binary for current platform"
	@echo "  make run         - Run the application"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make test        - Run tests"
	@echo "  make install     - Install binary to GOPATH/bin"
	@echo "  make release-all - Build binaries for all platforms"
	@echo ""
	@echo "Current settings:"
	@echo "  OS: $(OS)"
	@echo "  ARCH: $(ARCH)"
	@echo "  VERSION: $(VERSION)"

.DEFAULT_GOAL := help
