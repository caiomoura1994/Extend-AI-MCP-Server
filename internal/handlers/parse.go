package handlers

import (
	"context"
	"fmt"

	"github.com/caiomoura/extend-mcp-server/internal/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ParseHandlers contains all parse-related MCP handlers
type ParseHandlers struct {
	client *client.Client
}

// NewParseHandlers creates a new parse handlers instance
func NewParseHandlers(c *client.Client) *ParseHandlers {
	return &ParseHandlers{client: c}
}

// GetParseRunInput defines the input schema for get_parse_run tool
type GetParseRunInput struct {
	RunID string `json:"run_id" jsonschema:"ID of the parse run to retrieve"`
}

// HandleGetParseRun handles the get_parse_run tool call
func (h *ParseHandlers) HandleGetParseRun(ctx context.Context, req *mcp.CallToolRequest, input GetParseRunInput) (*mcp.CallToolResult, *client.ParserRun, error) {
	// Validate
	if input.RunID == "" {
		return nil, nil, fmt.Errorf("run_id is required")
	}

	// Call API
	result, err := h.client.GetParseRun(ctx, input.RunID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get parse run: %w", err)
	}

	// Return result
	return nil, result, nil
}
