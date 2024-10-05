# Makefile for Slack Bot Server

# Variables
BINARY_NAME=slack-bot-server
GO=go

# Main targets
.PHONY: all deps build build-linux run clean

all: deps build build-linux

deps:
	@echo "Installing dependencies..."
	$(GO) get github.com/slack-go/slack

build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -o $(BINARY_NAME) .

build-linux:
	@echo "Cross-compiling $(BINARY_NAME) for Linux AMD64..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -o $(BINARY_NAME)-linux-amd64 .

run: build
	@if [ -z "$(TOKEN)" ] || [ -z "$(CHANNEL)" ]; then \
		echo "Error: TOKEN and CHANNEL must be set"; \
		exit 1; \
	fi
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME) -token=$(TOKEN) -channel=$(CHANNEL)

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME) $(BINARY_NAME)-linux-amd64
	$(GO) clean

# Helper targets
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  deps         - Install dependencies"
	@echo "  build        - Build the bot for the current system"
	@echo "  build-linux  - Cross-compile the bot for Linux AMD64"
	@echo "  run          - Run the bot (requires TOKEN and CHANNEL to be set)"
	@echo "  clean        - Remove built binaries"
	@echo "  help         - Show this help message"
