# Makefile for the TRUMP Programming Language
# Author: Andrew Donelson

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=trumpc
BINARY_UNIX=$(BINARY_NAME)_unix

# Build directory
BUILD_DIR=build

# Source directories
CMD_DIR=cmd/trumpc
INTERNAL_DIRS=internal/lexer internal/parser internal/interpreter internal/errors

# All packages
ALL_PACKAGES=./$(CMD_DIR) ./internal/...

.PHONY: all build clean test install uninstall fmt lint run help examples docs

all: test build

build:
	@echo "Building TRUMP compiler..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./$(CMD_DIR)
	@echo "Build complete! Binary at $(BUILD_DIR)/$(BINARY_NAME)"

clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	@echo "Clean complete!"

test:
	@echo "Running tests..."
	$(GOTEST) -v $(ALL_PACKAGES)
	@echo "Tests complete!"

install: build
	@echo "Installing TRUMP compiler..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)
	@echo "Installation complete! You can now use '$(BINARY_NAME)' from anywhere."

uninstall:
	@echo "Uninstalling TRUMP compiler..."
	rm -f $(GOPATH)/bin/$(BINARY_NAME)
	@echo "Uninstalled!"

fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt $(ALL_PACKAGES)
	@echo "Formatting complete!"

lint:
	@echo "Linting code..."
	golint $(ALL_PACKAGES)
	@echo "Linting complete!"

run: build
	@echo "Running TRUMP compiler..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Build examples
examples: build
	@echo "Building examples..."
	mkdir -p examples/hello
	echo '// A tremendous hello world program\nTWEET "Hello, World! It is going to be TREMENDOUS!";' > examples/hello/main.trump
	mkdir -p examples/fibonacci
	echo '// Calculate Fibonacci numbers\n\nYUGE FUNCTION fibonacci(n) {\n  BUILD WALL IF (n <= 1) {\n    RETURN n;\n  } ELSE {\n    RETURN fibonacci(n-1) + fibonacci(n-2);\n  }\n}\n\nMAKE AMERICA GREAT AGAIN FOR (YUGE i = 0; i < 10; i = i + 1) {\n  TWEET "Fibonacci #" + i + ": " + fibonacci(i);\n}' > examples/fibonacci/main.trump
	@echo "Examples created in examples/ directory!"

# Build documentation
docs:
	@echo "Building documentation..."
	mkdir -p docs
	cp docs.md.md docs/index.md
	@echo "Documentation built in docs/ directory!"

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) -v ./$(CMD_DIR)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME).exe -v ./$(CMD_DIR)

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)_mac -v ./$(CMD_DIR)

# Build all platforms
release: build-linux build-windows build-mac

# Show help
help:
	@echo "TRUMP Programming Language - Make Programming Great Again!"
	@echo ""
	@echo "Available commands:"
	@echo "  make all         - Run tests and build binary"
	@echo "  make build       - Build the TRUMP compiler"
	@echo "  make clean       - Remove build artifacts"
	@echo "  make test        - Run tests"
	@echo "  make install     - Install TRUMP compiler to GOPATH/bin"
	@echo "  make uninstall   - Remove TRUMP compiler from GOPATH/bin"
	@echo "  make fmt         - Format the code"
	@echo "  make lint        - Run linter"
	@echo "  make run         - Run the TRUMP compiler"
	@echo "  make examples    - Build example programs"
	@echo "  make docs        - Build documentation"
	@echo "  make release     - Build for multiple platforms"
	@echo ""
	@echo "Usage after install: trumpc run yourprogram.trump"