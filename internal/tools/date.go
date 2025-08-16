package tools

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetCurrentDateTool() mcp.Tool {
	return mcp.NewTool("get_current_date",
		mcp.WithDescription("Get the current date and time"),
		mcp.WithString("format",
			mcp.Description("Date format (optional). Defaults to RFC3339. Common formats: 'rfc3339', 'unix', 'date-only'"),
		),
	)
}

func HandleGetCurrentDate(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	format := mcp.ParseString(request, "format", "rfc3339")

	now := time.Now()
	var result string

	switch format {
	case "unix":
		result = fmt.Sprintf("%d", now.Unix())
	case "date-only":
		result = now.Format("2006-01-02")
	case "rfc3339":
		fallthrough
	default:
		result = now.Format(time.RFC3339)
	}

	return mcp.NewToolResultText(result), nil
}