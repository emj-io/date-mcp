# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based MCP (Model Context Protocol) server that provides date-related functionality. It exposes a `get_current_date` tool that returns the current date/time in various formats through the MCP protocol.

## Development Setup

This project uses Go and the mark3labs/mcp-go library for MCP implementation.

## Project Structure

```
.
├── cmd/server/          # Main server entry point
│   ├── main.go         # Server startup and tool registration
│   └── main_test.go    # Integration tests
├── internal/tools/     # Tool implementations
│   ├── date.go        # get_current_date tool implementation
│   └── date_test.go   # Unit tests for date tool
├── go.mod             # Go module definition
├── go.sum             # Dependency checksums
├── CLAUDE.md          # This file
├── LICENSE            # MIT license
└── README.md          # Project documentation
```

## Architecture

The MCP server exposes tools through the Model Context Protocol:
- **Server**: Built using mark3labs/mcp-go, serves over stdio transport
- **Tools**: Date-related functionality in internal/tools package
- **get_current_date tool**: Returns current date/time with format options:
  - `rfc3339` (default): RFC3339 timestamp format
  - `unix`: Unix timestamp
  - `date-only`: YYYY-MM-DD format

## Commands

- `go build ./cmd/server` - Build the MCP server
- `go test ./...` - Run all tests
- `go test ./internal/tools/` - Run tool unit tests
- `go test ./cmd/server/` - Run server integration tests
- `go mod tidy` - Clean up dependencies
- `./server` - Run the MCP server (after building)

## Tool Implementation

Tools are implemented in the `internal/tools` package following this pattern:
1. Tool definition function (e.g., `GetCurrentDateTool()`) using mcp.NewTool
2. Handler function (e.g., `HandleGetCurrentDate()`) that processes requests
3. Registration in main.go using `s.AddTool()`

## Testing

Tests use the standard Go testing framework:
- Unit tests verify tool logic and output formats
- Integration tests verify server setup and tool registration
- Type assertions use mcp.AsTextContent() for accessing text content