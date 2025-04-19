# GoQuick CLI Makefile

BINARY_NAME=quick
VERSION=1.0.0
BUILD_DIR=bin
MAIN_PACKAGE=main.go
LDFLAGS=-ldflags "-w -s -X main.VERSION=$(VERSION)"

# Detect the operating system and architecture
OS=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)

.PHONY: all build clean run test install release-all release-linux release-darwin release-windows

## clean and build
all: clean build

# ========================
# 🛠️  Build Targets
# ========================

## Build the CLI binary
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

## Run the CLI directly (use: make run CMD="addc user")
run:
	@echo "🚀 Running: $(BINARY_NAME) $(CMD)"
	@go run $(MAIN_PACKAGE) $(CMD)

## Clean Build
clean:
	@rm -rf $(BUILD_DIR)
	@go clean

## Run all tests
test:
	@echo "🧪 Running tests..."
	@go test ./...

## Run tests with coverage
cover:
	@go test -v -covermode=count -coverprofile=coverage.out -tags exclude_test ./... 

## Show HTML coverage report
cover-html: cover ## 🌐 Open HTML coverage
	@go tool cover -html=coverage.out

## Install CLI
install:
	@go install $(LDFLAGS) $(MAIN_PACKAGE)

# ========================
# 📦 Cross Compilation
# ========================

## Release builds for all platforms
release-all: release-linux release-darwin release-windows

## Build for Linux
release-linux:
	@mkdir -p $(BUILD_DIR)/linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/linux/$(BINARY_NAME) $(MAIN_PACKAGE)

## Build for MacOS
release-darwin:
	@mkdir -p $(BUILD_DIR)/darwin
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/darwin/$(BINARY_NAME) $(MAIN_PACKAGE)

## Build for Windows
release-windows:
	@mkdir -p $(BUILD_DIR)/windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/windows/$(BINARY_NAME).exe $(MAIN_PACKAGE)

# ========================
# 🆘 Help & Default
# ========================

## Default target
help:
	@echo "🧰 GoQuick CLI Makefile"
	@echo ""
	@echo "📦 Usage:"
	@echo "  make build         - Build the binary for current platform"
	@echo "  make run CMD=...   - Run the CLI with arguments"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make install       - Install binary to \$GOBIN"
	@echo ""
	@echo "🧪 Testing:"
	@echo "  make test          - Run tests"
	@echo "  make cover         - Run tests with coverage"
	@echo "  make cover-html    - Open HTML coverage report"
	@echo ""
	@echo "🌍 Cross Compilation:"
	@echo "  make release-linux     - Build Linux binary"
	@echo "  make release-darwin    - Build macOS binary"
	@echo "  make release-windows   - Build Windows binary"
	@echo "  make release-all       - Build for all platforms"
	@echo ""
	@echo "Current settings:"
	@echo "  OS: $(OS)"
	@echo "  ARCH: $(ARCH)"
	@echo "  VERSION: $(VERSION)"

.DEFAULT_GOAL := help
