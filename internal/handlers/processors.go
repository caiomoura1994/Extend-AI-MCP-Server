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
	if input.RunID == "" {
		return nil, nil, fmt.Errorf("run_id is required")
	}

	result, err := h.client.GetProcessorRun(ctx, input.RunID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get processor run: %w", err)
	}

	return nil, result, nil
}

// ListProcessorsInput defines the input schema for list_processors tool
type ListProcessorsInput struct {
	MaxPageSize   *int    `json:"max_page_size,omitempty" jsonschema:"Maximum number of items per page (default: 10)"`
	NextPageToken *string `json:"next_page_token,omitempty" jsonschema:"Token for the next page of results"`
	SortBy        *string `json:"sort_by,omitempty" jsonschema:"Sort by field: updatedAt or createdAt (default: updatedAt)"`
	SortDir       *string `json:"sort_dir,omitempty" jsonschema:"Sort direction: asc or desc (default: desc)"`
}

// ListProcessorsOutput wraps the processors list with pagination
type ListProcessorsOutput struct {
	Processors    []dto.Processor `json:"processors"`
	NextPageToken *string         `json:"next_page_token,omitempty"`
}

// HandleListProcessors handles the list_processors tool call
func (h *ProcessorHandlers) HandleListProcessors(ctx context.Context, req *mcp.CallToolRequest, input ListProcessorsInput) (*mcp.CallToolResult, *ListProcessorsOutput, error) {
	pagination := &dto.PaginationParams{
		MaxPageSize:   input.MaxPageSize,
		NextPageToken: input.NextPageToken,
		SortBy:        input.SortBy,
		SortDir:       input.SortDir,
	}

	result, err := h.client.ListProcessors(ctx, pagination)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list processors: %w", err)
	}

	if result.Processors == nil {
		result.Processors = []dto.Processor{}
	}
	return nil, &ListProcessorsOutput{
		Processors:    result.Processors,
		NextPageToken: result.NextPageToken,
	}, nil
}
