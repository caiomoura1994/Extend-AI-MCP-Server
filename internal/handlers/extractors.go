package handlers

import (
	"context"
	"fmt"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
	"github.com/caiomoura/extend-mcp-server/internal/repositories"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ExtractorHandlers contains all extractor-related MCP handlers
type ExtractorHandlers struct {
	client *repositories.Client
}

// NewExtractorHandlers creates a new extractor handlers instance
func NewExtractorHandlers(c *repositories.Client) *ExtractorHandlers {
	return &ExtractorHandlers{client: c}
}

// ListExtractRunsOutput wraps the extract runs list
type ListExtractRunsOutput struct {
	ExtractRuns []dto.ExtractRun `json:"extract_runs"`
}

// HandleListExtractRuns handles the list_extract_runs tool call
func (h *ExtractorHandlers) HandleListExtractRuns(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, *ListExtractRunsOutput, error) {
	// Call API
	result, err := h.client.ListExtractRuns(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list extract runs: %w", err)
	}

	// Ensure non-nil slice (MCP SDK requires array, not null)
	if result == nil {
		result = []dto.ExtractRun{}
	}
	return nil, &ListExtractRunsOutput{ExtractRuns: result}, nil
}

// GetExtractorInput defines the input schema for get_extractor tool
type GetExtractorInput struct {
	ExtractorID string `json:"extractor_id" jsonschema:"ID of the extractor to retrieve"`
}

// HandleGetExtractor handles the get_extractor tool call
func (h *ExtractorHandlers) HandleGetExtractor(ctx context.Context, req *mcp.CallToolRequest, input GetExtractorInput) (*mcp.CallToolResult, *dto.Extractor, error) {
	// Validate
	if input.ExtractorID == "" {
		return nil, nil, fmt.Errorf("extractor_id is required")
	}

	// Call API
	result, err := h.client.GetExtractor(ctx, input.ExtractorID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get extractor: %w", err)
	}

	// Return result
	return nil, result, nil
}

// ListExtractorsOutput wraps the extractors list
type ListExtractorsOutput struct {
	Extractors []dto.Extractor `json:"extractors"`
}

// HandleListExtractors handles the list_extractors tool call
func (h *ExtractorHandlers) HandleListExtractors(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, *ListExtractorsOutput, error) {
	// Call API
	result, err := h.client.ListExtractors(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list extractors: %w", err)
	}

	// Ensure non-nil slice (MCP SDK requires array, not null)
	if result == nil {
		result = []dto.Extractor{}
	}
	return nil, &ListExtractorsOutput{Extractors: result}, nil
}

// GetExtractorVersionInput defines the input schema for get_extractor_version tool
type GetExtractorVersionInput struct {
	ExtractorID string `json:"extractor_id" jsonschema:"ID of the extractor"`
	VersionID   string `json:"version_id" jsonschema:"ID of the version to retrieve"`
}

// HandleGetExtractorVersion handles the get_extractor_version tool call
func (h *ExtractorHandlers) HandleGetExtractorVersion(ctx context.Context, req *mcp.CallToolRequest, input GetExtractorVersionInput) (*mcp.CallToolResult, *dto.ExtractorVersion, error) {
	// Validate
	if input.ExtractorID == "" {
		return nil, nil, fmt.Errorf("extractor_id is required")
	}
	if input.VersionID == "" {
		return nil, nil, fmt.Errorf("version_id is required")
	}

	// Call API
	result, err := h.client.GetExtractorVersion(ctx, input.ExtractorID, input.VersionID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get extractor version: %w", err)
	}

	// Return result
	return nil, result, nil
}

// ListExtractorVersionsInput defines the input schema for list_extractor_versions tool
type ListExtractorVersionsInput struct {
	ExtractorID string `json:"extractor_id" jsonschema:"ID of the extractor"`
}

// ListExtractorVersionsOutput wraps the extractor versions list
type ListExtractorVersionsOutput struct {
	Versions []dto.ExtractorVersion `json:"versions"`
}

// HandleListExtractorVersions handles the list_extractor_versions tool call
func (h *ExtractorHandlers) HandleListExtractorVersions(ctx context.Context, req *mcp.CallToolRequest, input ListExtractorVersionsInput) (*mcp.CallToolResult, *ListExtractorVersionsOutput, error) {
	// Validate
	if input.ExtractorID == "" {
		return nil, nil, fmt.Errorf("extractor_id is required")
	}

	// Call API
	result, err := h.client.ListExtractorVersions(ctx, input.ExtractorID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list extractor versions: %w", err)
	}

	// Ensure non-nil slice (MCP SDK requires array, not null)
	if result == nil {
		result = []dto.ExtractorVersion{}
	}
	return nil, &ListExtractorVersionsOutput{Versions: result}, nil
}
