# Makefile for Slack Bot Server

# Variables
BINARY_NAME=slack-bot-server
GO=go

# Main targets
.PHONY: all deps build run clean

all: deps build

deps:
	@echo "Installing dependencies..."
	$(GO) get github.com/slack-go/slack

build:
	@echo "Building $(BINARY_NAME)..."
	$(GO) build -o $(BINARY_NAME) .

run: build
	@if [ -z "$(TOKEN)" ] || [ -z "$(CHANNEL)" ]; then \
		echo "Error: TOKEN and CHANNEL must be set"; \
		exit 1; \
	fi
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME) -token=$(TOKEN) -channel=$(CHANNEL)

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	$(GO) clean

# Helper targets
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  deps   - Install dependencies"
	@echo "  build  - Build the bot"
	@echo "  run    - Run the bot (requires TOKEN and CHANNEL to be set)"
	@echo "  clean  - Remove built binary"
	@echo "  help   - Show this help message"
