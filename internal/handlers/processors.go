package handlers

import (
	"context"
	"fmt"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
	"github.com/caiomoura/extend-mcp-server/internal/repositories"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ProcessorHandlers contains all processor-related MCP handlers
type ProcessorHandlers struct {
	client *repositories.Client
}

// NewProcessorHandlers creates a new processor handlers instance
func NewProcessorHandlers(c *repositories.Client) *ProcessorHandlers {
	return &ProcessorHandlers{client: c}
}

// GetProcessorRunInput defines the input schema for get_processor_run tool
type GetProcessorRunInput struct {
	RunID string `json:"run_id" jsonschema:"ID of the processor run to retrieve"`
}

// HandleGetProcessorRun handles the get_processor_run tool call
func (h *ProcessorHandlers) HandleGetProcessorRun(ctx context.Context, req *mcp.CallToolRequest, input GetProcessorRunInput) (*mcp.CallToolResult, *dto.ProcessorRun, error) {
	// Validate
	if input.RunID == "" {
		return nil, nil, fmt.Errorf("run_id is required")
	}

	// Call API
	result, err := h.client.GetProcessorRun(ctx, input.RunID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get processor run: %w", err)
	}

	// Return result
	return nil, result, nil
}

// ListProcessorsOutput wraps the processors list
type ListProcessorsOutput struct {
	Processors []dto.Processor `json:"processors"`
}

// HandleListProcessors handles the list_processors tool call
func (h *ProcessorHandlers) HandleListProcessors(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, *ListProcessorsOutput, error) {
	// Call API
	result, err := h.client.ListProcessors(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list processors: %w", err)
	}

	// Ensure non-nil slice (MCP SDK requires array, not null)
	if result == nil {
		result = []dto.Processor{}
	}
	return nil, &ListProcessorsOutput{Processors: result}, nil
}
