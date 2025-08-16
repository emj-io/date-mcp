# date-mcp

A Model Context Protocol (MCP) server that provides date-related tools for AI applications.

## Overview

This MCP server exposes date and time functionality through the Model Context Protocol, allowing AI applications like Claude to get current date/time information in various formats.

## Features

- **get_current_date** tool with multiple format options:
  - `rfc3339` (default): RFC3339 timestamp format
  - `unix`: Unix timestamp
  - `date-only`: YYYY-MM-DD format

## Installation

```bash
git clone https://github.com/ericjohnson/date-mcp
cd date-mcp
go mod download
```

## Building

```bash
go build -o date-mcp ./cmd/server
```

## Running

The MCP server runs via stdio transport:

```bash
./date-mcp
```

## Integration

### Cursor IDE

1. **Build the server** (if not already built):
   ```bash
   go build -o date-mcp ./cmd/server
   ```

2. **Create MCP configuration file**:
   Create `.cursor/mcp.json` in your project root (or globally at `~/.cursor/mcp.json`):

   ```json
   {
     "mcpServers": {
       "date-mcp": {
         "command": "./date-mcp",
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
      "command": "./date-mcp",
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
go test ./...
```

Run specific test suites:
```bash
go test ./internal/tools/    # Tool unit tests
go test ./cmd/server/        # Server integration tests
```

### Project Structure

- `cmd/server/` - Main server entry point
- `internal/tools/` - Tool implementations
- Tests are co-located with source files

## License

MIT License - see LICENSE file for details.