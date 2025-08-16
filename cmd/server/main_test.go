package main

import (
	"context"
	"testing"

	"github.com/ericjohnson/date-mcp/internal/tools"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func TestServerCreation(t *testing.T) {
	s := server.NewMCPServer(
		"date-mcp",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)
	
	if s == nil {
		t.Fatal("Expected server to be created")
	}
}

func TestToolIntegration(t *testing.T) {
	s := server.NewMCPServer(
		"date-mcp",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)
	
	s.AddTool(tools.GetCurrentDateTool(), tools.HandleGetCurrentDate)
	
	// Test that the tool handler works in the context of the server
	ctx := context.Background()
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Name: "get_current_date",
			Arguments: map[string]interface{}{
				"format": "date-only",
			},
		},
	}
	
	result, err := tools.HandleGetCurrentDate(ctx, request)
	if err != nil {
		t.Fatalf("Expected no error from tool handler, got %v", err)
	}
	
	if len(result.Content) != 1 {
		t.Fatalf("Expected 1 content item, got %d", len(result.Content))
	}
	
	textContent, ok := mcp.AsTextContent(result.Content[0])
	if !ok {
		t.Fatal("Expected content to be TextContent")
	}
	
	if textContent.Text == "" {
		t.Error("Expected non-empty text result")
	}
}