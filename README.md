# date-mcp

[![CI](https://github.com/emj-io/date-mcp/actions/workflows/ci.yml/badge.svg)](https://github.com/emj-io/date-mcp/actions/workflows/ci.yml)
[![Security](https://github.com/emj-io/date-mcp/actions/workflows/security.yml/badge.svg)](https://github.com/emj-io/date-mcp/actions/workflows/security.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/emj-io/date-mcp)](https://goreportcard.com/report/github.com/emj-io/date-mcp)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Model Context Protocol (MCP) server that provides date-related tools for AI applications.

## Overview

This MCP server exposes date and time functionality through the Model Context Protocol, allowing AI applications like Claude to get current date/time information in various formats.

## Features

- **get_current_date** tool with multiple format options:
  - `rfc3339` (default): RFC3339 timestamp format
  - `unix`: Unix timestamp
  - `date-only`: YYYY-MM-DD format

## Quick Start

```bash
git clone https://github.com/ericjohnson/date-mcp
cd date-mcp
make all    # Download deps, run tests, and build
make run    # Run the server
```

## Development Commands

This project uses a Makefile for common development tasks:

```bash
make deps     # Download and tidy dependencies
make build    # Build the MCP server binary
make test     # Run all tests
make run      # Build and run the server
make clean    # Remove build artifacts
make lint     # Run golangci-lint (if available)
make all      # Full build pipeline (deps, test, build)
make help     # Show all available commands
```

### Alternative: Direct Go Commands

If you prefer to use Go commands directly:

```bash
go mod download                    # Download dependencies
go build -o bin/date-mcp ./cmd/server  # Build
go test ./...                     # Test
./bin/date-mcp                    # Run
```

## Integration

### Cursor IDE

1. **Build the server** (if not already built):
   ```bash
   make build
   ```

2. **Create MCP configuration file**:
   Create `.cursor/mcp.json` in your project root (or globally at `~/.cursor/mcp.json`):

   ```json
   {
     "mcpServers": {
       "date-mcp": {
         "command": "./bin/date-mcp",
         "args": []
       }
     }
   }
   ```

3. **Configure Cursor**:
   - Open Cursor
   - Go to Settings > MCP
   - Verify the server shows as "Active" with a green status
   - Switch to Agent Mode (not Ask Mode) to access MCP tools

4. **Test the integration**:
   Ask the Cursor Agent: "What's the current date?" or "Get the current date in unix format"
   
   The agent will automatically use the `get_current_date` tool when relevant.

### Claude Desktop

Add to your Claude Desktop MCP configuration:

```json
{
  "mcpServers": {
    "date-mcp": {
      "command": "./bin/date-mcp",
      "args": []
    }
  }
}
```

## Usage Examples

Once integrated with an MCP client, you can use the date tool:

- Get current date in default format: `get_current_date`
- Get unix timestamp: `get_current_date` with format="unix"
- Get date only: `get_current_date` with format="date-only"

## Development

### Testing

Run all tests:
```bash
make test
```

Or use Go directly:
```bash
go test ./...               # All tests
go test ./internal/tools/   # Tool unit tests  
go test ./cmd/server/       # Server integration tests
```

### Project Structure

- `cmd/server/` - Main server entry point
- `internal/tools/` - Tool implementations
- Tests are co-located with source files

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Workflow

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`make test`)
5. Run linting (`make lint`)
6. Commit your changes (`git commit -m 'Add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

### Security

Please report security vulnerabilities by emailing the maintainers directly rather than opening public issues.

## License

MIT License - see LICENSE file for details.