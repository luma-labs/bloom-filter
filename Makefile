# Define project variables
BINARY_NAME=bloomfilter
PKG=github.com/luma-labs/bloom-filter

# Default target: Build the binary
build:
	@echo "Building the binary..."
	go build -o bin/$(BINARY_NAME) ./cmd/bloomfilter/

# Run the main application
run:
	@echo "Running the application..."
	go run ./cmd/bloomfilter/main.go

# Run all tests
test:
	@echo "Running tests..."
	go test ./pkg/... ./tests/...

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	go test -bench=. ./benchmarks/...

# Clean the binary and other build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf bin/

# Lint the project (you'll need to install golangci-lint)
lint:
	@echo "Linting the project..."
	golangci-lint run ./...

# Install necessary tools (e.g., linters, formatters)
tools:
	@echo "Installing tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Generate test coverage report
cover:
	@echo "Running test coverage..."
	go test ./pkg/... ./tests/... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
