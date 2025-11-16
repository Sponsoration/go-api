.PHONY: test test-verbose test-coverage test-race bench clean fmt vet lint run-test-email help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

# Binary names
BINARY_NAME=sponsoration-api
TEST_EMAIL_BINARY=test-email

# Colors for output
GREEN=\033[0;32m
YELLOW=\033[0;33m
RED=\033[0;31m
NC=\033[0m # No Color

## help: Show this help message
help:
	@echo "$(GREEN)Sponsoration Go API - Available Commands$(NC)"
	@echo ""
	@echo "$(YELLOW)Testing:$(NC)"
	@echo "  make test              - Run all tests"
	@echo "  make test-verbose      - Run tests with verbose output"
	@echo "  make test-coverage     - Run tests with coverage report"
	@echo "  make test-race         - Run tests with race detector"
	@echo "  make bench             - Run benchmark tests"
	@echo ""
	@echo "$(YELLOW)Code Quality:$(NC)"
	@echo "  make fmt               - Format code"
	@echo "  make vet               - Run go vet"
	@echo "  make lint              - Run linter (requires golangci-lint)"
	@echo ""
	@echo "$(YELLOW)Build:$(NC)"
	@echo "  make build             - Build the API server"
	@echo "  make run-test-email    - Run email test program"
	@echo ""
	@echo "$(YELLOW)Maintenance:$(NC)"
	@echo "  make clean             - Clean build artifacts"
	@echo "  make deps              - Download dependencies"
	@echo "  make deps-update       - Update dependencies"
	@echo "  make tidy              - Tidy dependencies"
	@echo ""

## test: Run all tests
test:
	@echo "$(GREEN)Running tests...$(NC)"
	$(GOTEST) -v ./...

## test-verbose: Run tests with verbose output
test-verbose:
	@echo "$(GREEN)Running tests (verbose)...$(NC)"
	$(GOTEST) -v -count=1 ./...

## test-coverage: Run tests with coverage report
test-coverage:
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	$(GOTEST) -v -coverprofile=coverage.out -covermode=atomic ./...
	@echo ""
	@echo "$(GREEN)Coverage Summary:$(NC)"
	$(GOCMD) tool cover -func=coverage.out | tail -1
	@echo ""
	@echo "$(YELLOW)To view HTML coverage report, run:$(NC)"
	@echo "  go tool cover -html=coverage.out"

## test-race: Run tests with race detector
test-race:
	@echo "$(GREEN)Running tests with race detector...$(NC)"
	$(GOTEST) -v -race ./...

## bench: Run benchmark tests
bench:
	@echo "$(GREEN)Running benchmarks...$(NC)"
	$(GOTEST) -bench=. -benchmem ./...

## fmt: Format code
fmt:
	@echo "$(GREEN)Formatting code...$(NC)"
	$(GOFMT) ./...

## vet: Run go vet
vet:
	@echo "$(GREEN)Running go vet...$(NC)"
	$(GOVET) ./...

## lint: Run golangci-lint (requires installation)
lint:
	@echo "$(GREEN)Running linter...$(NC)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "$(YELLOW)golangci-lint not installed. Install with:$(NC)"; \
		echo "  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin"; \
	fi

## build: Build the API server
build:
	@echo "$(GREEN)Building $(BINARY_NAME)...$(NC)"
	$(GOBUILD) -o bin/$(BINARY_NAME) -v ./cmd/api

## run-test-email: Run email test program
run-test-email:
	@echo "$(GREEN)Running email test program...$(NC)"
	@if [ -z "$(EMAIL)" ]; then \
		echo "$(YELLOW)Usage: make run-test-email EMAIL=your-email@example.com$(NC)"; \
		exit 1; \
	fi
	$(GOCMD) run cmd/test-email/main.go $(EMAIL)

## clean: Clean build artifacts
clean:
	@echo "$(GREEN)Cleaning...$(NC)"
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -f bin/$(TEST_EMAIL_BINARY)
	rm -f coverage.out
	rm -f coverage.html

## deps: Download dependencies
deps:
	@echo "$(GREEN)Downloading dependencies...$(NC)"
	$(GOMOD) download

## deps-update: Update dependencies
deps-update:
	@echo "$(GREEN)Updating dependencies...$(NC)"
	$(GOGET) -u ./...
	$(GOMOD) tidy

## tidy: Tidy dependencies
tidy:
	@echo "$(GREEN)Tidying dependencies...$(NC)"
	$(GOMOD) tidy

# Default target
.DEFAULT_GOAL := help
