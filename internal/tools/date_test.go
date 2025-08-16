package tools

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestGetCurrentDateTool(t *testing.T) {
	tool := GetCurrentDateTool()

	if tool.Name != "get_current_date" {
		t.Errorf("Expected tool name 'get_current_date', got '%s'", tool.Name)
	}

	if tool.Description == "" {
		t.Error("Tool description should not be empty")
	}
}

func TestHandleGetCurrentDate_DefaultFormat(t *testing.T) {
	ctx := context.Background()
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{},
		},
	}

	result, err := HandleGetCurrentDate(ctx, request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(result.Content) != 1 {
		t.Fatalf("Expected 1 content item, got %d", len(result.Content))
	}

	textContent, ok := mcp.AsTextContent(result.Content[0])
	if !ok {
		t.Fatal("Expected content to be TextContent")
	}

	// Try to parse as RFC3339 (default format)
	_, err = time.Parse(time.RFC3339, textContent.Text)
	if err != nil {
		t.Errorf("Expected valid RFC3339 timestamp, got '%s': %v", textContent.Text, err)
	}
}

func TestHandleGetCurrentDate_UnixFormat(t *testing.T) {
	ctx := context.Background()
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"format": "unix",
			},
		},
	}

	result, err := HandleGetCurrentDate(ctx, request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	textContent, ok := mcp.AsTextContent(result.Content[0])
	if !ok {
		t.Fatal("Expected content to be TextContent")
	}

	// Try to parse as unix timestamp
	_, err = strconv.ParseInt(textContent.Text, 10, 64)
	if err != nil {
		t.Errorf("Expected valid unix timestamp, got '%s': %v", textContent.Text, err)
	}
}

func TestHandleGetCurrentDate_DateOnlyFormat(t *testing.T) {
	ctx := context.Background()
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"format": "date-only",
			},
		},
	}

	result, err := HandleGetCurrentDate(ctx, request)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	textContent, ok := mcp.AsTextContent(result.Content[0])
	if !ok {
		t.Fatal("Expected content to be TextContent")
	}

	// Try to parse as date-only format (YYYY-MM-DD)
	_, err = time.Parse("2006-01-02", textContent.Text)
	if err != nil {
		t.Errorf("Expected valid date-only format, got '%s': %v", textContent.Text, err)
	}
}
