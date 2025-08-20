# Makefile for FSM Modulo Three Project

.PHONY: help test build run clean verify

# Default target
help:
	@echo "Available targets:"
	@echo "  test     - Run all tests"
	@echo "  build    - Build the project"
	@echo "  run      - Run the interactive demo"
	@echo "  verify   - Run the verification script"
	@echo "  clean    - Clean build artifacts"
	@echo "  coverage - Run tests with coverage"

# Run all tests
test:
	go test -v ./...

# Build the project
build:
	go build -o bin/fsm-demo cmd/main.go

# Run the interactive demo
run:
	go run cmd/main.go

# Run the verification script
verify:
	go run verify.go

# Clean build artifacts
clean:
	go clean
	rm -rf bin/

# Run tests with coverage
coverage:
	go test -cover ./...

# Run tests with coverage report
coverage-report:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Format code
fmt:
	go fmt ./...

# Vet code
vet:
	go vet ./...

# Lint code (requires golint)
lint:
	golint ./...

# All checks
check: fmt vet test

