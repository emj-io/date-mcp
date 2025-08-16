.PHONY: deps build test run clean lint all help
.DEFAULT_GOAL := build

# Binary name and path
BINARY_NAME=date-mcp
BINARY_PATH=bin/$(BINARY_NAME)

# Build target
build:
	mkdir -p bin
	go build -o $(BINARY_PATH) ./cmd/server

# Download dependencies
deps:
	go mod download
	go mod tidy

# Run tests
test:
	go test -v ./...

# Run the application (depends on build)
run: build
	./$(BINARY_PATH)

# Clean build artifacts
clean:
	rm -rf bin/

# Run linter (prefer golangci-lint, fallback to basic tools)
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		echo "Running golangci-lint..."; \
		golangci-lint run; \
	elif [ -x "$$HOME/go/bin/golangci-lint" ]; then \
		echo "Running golangci-lint from GOPATH..."; \
		$$HOME/go/bin/golangci-lint run; \
	else \
		echo "golangci-lint not found, running basic Go tools..."; \
		echo "Running go fmt..."; \
		gofmt -d . | tee /tmp/gofmt.out; \
		if [ -s /tmp/gofmt.out ]; then \
			echo "❌ gofmt found formatting issues"; \
			exit 1; \
		fi; \
		echo "Running go vet..."; \
		go vet ./...; \
		echo "✅ Basic linting passed. Install golangci-lint for comprehensive checks:"; \
		echo "  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Full build pipeline
all: deps lint test build

# Show help
help:
	@echo "Available targets:"
	@echo "  deps     - Download and tidy dependencies"
	@echo "  build    - Build the MCP server binary"
	@echo "  test     - Run all tests"
	@echo "  run      - Build and run the server"
	@echo "  clean    - Remove build artifacts"
	@echo "  lint     - Run golangci-lint (if available)"
	@echo "  all      - Run deps, lint, test, and build"
	@echo "  help     - Show this help message"