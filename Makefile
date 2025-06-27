# ==============================
# Build configuration
# ==============================

# Target platform (override via `make TARGETOS=linux`)
TARGETOS     ?= linux
TARGETARCH   ?= arm64
CGO_ENABLED  ?= 0

# Default component name (subgraph, indexer)
TAG_SUFFIX   := "-arm64"
COMPONENT := subgraph
BIN_DIR   := bin
BIN_PATH  := $(BIN_DIR)/$(COMPONENT)

# Application name and registry
APP       := indexer.starknet.stablecoin.$(COMPONENT)
REGISTRY  := ghcr.io/tsisar
VERSION   := $(shell git describe --tags --abbrev=0 2>/dev/null || echo dev)-$(shell git rev-parse --short HEAD)

# ==============================
# Phony targets
# ==============================
.PHONY: help format lint test get build image push clean dev release install-lint generate-filters generate-gqlgen generate-ent generate-indexer generate-mapper generate

# ==============================
# Help
# ==============================
help:
	@echo "Available targets:"
	@echo "  help     - Show this help message"
	@echo "  format   - Format Go code"
	@echo "  lint     - Run golangci-lint"
	@echo "  test     - Run tests"
	@echo "  get      - Get dependencies"
	@echo "  build    - Build the application"
	@echo "  image    - Build Docker image"
	@echo "  push     - Push Docker image to registry"
	@echo "  clean    - Clean build artifacts"
	@echo "  dev      - Development build (with debug info)"
	@echo "  release  - Production release build and push"
	@echo ""
	@echo "Configuration:"
	@echo "  TARGETOS     = $(TARGETOS) (linux, windows, darwin)"
	@echo "  TARGETARCH   = $(TARGETARCH) (amd64, arm64)"
	@echo "  CGO_ENABLED  = $(CGO_ENABLED) (0 for no cgo, 1 for cgo enabled)"

# ==============================
# Go Tools
# ==============================
format:
	@echo "Formatting Go code..."
	@gofmt -s -w ./

install-lint:
	@which golangci-lint >/dev/null || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)

lint: install-lint
	@echo "Running linter..."
	@golangci-lint run ./...

test:
	@echo "Running tests..."
	@go test -v -cover ./...

generate-indexer:
	@echo "Generating core code..."
	@go run github.com/tsisar/starknet-indexer/cmd/generator-indexer

generate-filters:
	@echo "Generating filter code..."
	@go run github.com/tsisar/starknet-indexer/cmd/generator-graphql

generate-ent:
	@echo "Generating ent code..."
	@go run github.com/tsisar/starknet-indexer/cmd/generator-ent
	@ent generate ./generated/ent/schema

generate-gqlgen:
	@echo "Generating GraphQL schema..."
	@go run github.com/99designs/gqlgen generate

generate-mapper:
	@echo "Generating mapper code..."
	@go run github.com/tsisar/starknet-indexer/cmd/generator-mapper

generate: generate-indexer generate-mapper generate-filters generate-ent generate-gqlgen
	@echo "All generators executed successfully"

get:
	@echo "Getting dependencies..."
	@go mod tidy
	@go mod download

# ==============================
# Build targets
# ==============================
dev: format get
	@echo "Building development binary..."
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) \
		go build -v -o $(COMPONENT) ./cmd/$(COMPONENT)

build: format get
	@echo "Building production binary..."
	@CGO_ENABLED=$(CGO_ENABLED) GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) \
		go build -v -o $(BIN_PATH) -ldflags="-s -w" ./cmd/$(COMPONENT)

# ==============================
# Docker
# ==============================
image:
	@echo "Building Docker image for $(TARGETOS)/$(TARGETARCH)..."
	@docker buildx build \
		--platform $(TARGETOS)/$(TARGETARCH) \
		--build-arg BIN_PATH=$(BIN_PATH) \
		--tag $(REGISTRY)/$(APP):$(VERSION)$(TAG_SUFFIX) \
		--load .

push:
	@echo "Pushing image to registry..."
	@docker push $(REGISTRY)/$(APP):$(VERSION)$(TAG_SUFFIX)
	@docker tag $(REGISTRY)/$(APP):$(VERSION)$(TAG_SUFFIX) $(REGISTRY)/$(APP):latest$(TAG_SUFFIX)
	@docker push $(REGISTRY)/$(APP):latest$(TAG_SUFFIX)

# ==============================
# Utilities
# ==============================
clean:
	@echo "Cleaning artifacts..."
	@rm -f indexer
	@docker rmi $(REGISTRY)/$(APP):$(VERSION)$(TAG_SUFFIX) || true

release: clean test build image push
	@echo "Release $(VERSION) completed successfully"