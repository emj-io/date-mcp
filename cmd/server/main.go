package main

import (
	"fmt"
	"log"

	"github.com/ericjohnson/date-mcp/internal/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"date-mcp",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)

	s.AddTool(tools.GetCurrentDateTool(), tools.HandleGetCurrentDate)

	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v", err)
	}

	fmt.Println("Date MCP server started")
}