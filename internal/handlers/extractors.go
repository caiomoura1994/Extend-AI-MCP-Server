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

// ListExtractRunsInput defines the input schema for list_extract_runs tool
type ListExtractRunsInput struct {
	MaxPageSize   *int    `json:"max_page_size,omitempty" jsonschema:"Maximum number of items per page (default: 10)"`
	NextPageToken *string `json:"next_page_token,omitempty" jsonschema:"Token for the next page of results"`
	SortBy        *string `json:"sort_by,omitempty" jsonschema:"Sort by field: updatedAt or createdAt (default: updatedAt)"`
	SortDir       *string `json:"sort_dir,omitempty" jsonschema:"Sort direction: asc or desc (default: desc)"`
}

// ListExtractRunsOutput wraps the extract runs list with pagination
type ListExtractRunsOutput struct {
	ExtractRuns   []dto.ExtractRun `json:"extract_runs"`
	NextPageToken *string          `json:"next_page_token,omitempty"`
}

// HandleListExtractRuns handles the list_extract_runs tool call
func (h *ExtractorHandlers) HandleListExtractRuns(ctx context.Context, req *mcp.CallToolRequest, input ListExtractRunsInput) (*mcp.CallToolResult, *ListExtractRunsOutput, error) {
	pagination := &dto.PaginationParams{
		MaxPageSize:   input.MaxPageSize,
		NextPageToken: input.NextPageToken,
		SortBy:        input.SortBy,
		SortDir:       input.SortDir,
	}

	result, err := h.client.ListExtractRuns(ctx, pagination)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list extract runs: %w", err)
	}

	if result.Data == nil {
		result.Data = []dto.ExtractRun{}
	}
	return nil, &ListExtractRunsOutput{
		ExtractRuns:   result.Data,
		NextPageToken: result.NextPageToken,
	}, nil
}

// GetExtractorInput defines the input schema for get_extractor tool
type GetExtractorInput struct {
	ExtractorID string `json:"extractor_id" jsonschema:"ID of the extractor to retrieve"`
}

// HandleGetExtractor handles the get_extractor tool call
func (h *ExtractorHandlers) HandleGetExtractor(ctx context.Context, req *mcp.CallToolRequest, input GetExtractorInput) (*mcp.CallToolResult, *dto.Extractor, error) {
	if input.ExtractorID == "" {
		return nil, nil, fmt.Errorf("extractor_id is required")
	}

	result, err := h.client.GetExtractor(ctx, input.ExtractorID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get extractor: %w", err)
	}

	return nil, result, nil
}

// ListExtractorsInput defines the input schema for list_extractors tool
type ListExtractorsInput struct {
	MaxPageSize   *int    `json:"max_page_size,omitempty" jsonschema:"Maximum number of items per page (default: 10)"`
	NextPageToken *string `json:"next_page_token,omitempty" jsonschema:"Token for the next page of results"`
	SortBy        *string `json:"sort_by,omitempty" jsonschema:"Sort by field: updatedAt or createdAt (default: updatedAt)"`
	SortDir       *string `json:"sort_dir,omitempty" jsonschema:"Sort direction: asc or desc (default: desc)"`
}

// ListExtractorsOutput wraps the extractors list with pagination
type ListExtractorsOutput struct {
	Extractors    []dto.Extractor `json:"extractors"`
	NextPageToken *string         `json:"next_page_token,omitempty"`
}

// HandleListExtractors handles the list_extractors tool call
func (h *ExtractorHandlers) HandleListExtractors(ctx context.Context, req *mcp.CallToolRequest, input ListExtractorsInput) (*mcp.CallToolResult, *ListExtractorsOutput, error) {
	pagination := &dto.PaginationParams{
		MaxPageSize:   input.MaxPageSize,
		NextPageToken: input.NextPageToken,
		SortBy:        input.SortBy,
		SortDir:       input.SortDir,
	}

	result, err := h.client.ListExtractors(ctx, pagination)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list extractors: %w", err)
	}

	if result.Data == nil {
		result.Data = []dto.Extractor{}
	}
	return nil, &ListExtractorsOutput{
		Extractors:    result.Data,
		NextPageToken: result.NextPageToken,
	}, nil
}

// GetExtractorVersionInput defines the input schema for get_extractor_version tool
type GetExtractorVersionInput struct {
	ExtractorID string `json:"extractor_id" jsonschema:"ID of the extractor"`
	VersionID   string `json:"version_id" jsonschema:"ID of the version to retrieve"`
}

// HandleGetExtractorVersion handles the get_extractor_version tool call
func (h *ExtractorHandlers) HandleGetExtractorVersion(ctx context.Context, req *mcp.CallToolRequest, input GetExtractorVersionInput) (*mcp.CallToolResult, *dto.ExtractorVersion, error) {
	if input.ExtractorID == "" {
		return nil, nil, fmt.Errorf("extractor_id is required")
	}
	if input.VersionID == "" {
		return nil, nil, fmt.Errorf("version_id is required")
	}

	result, err := h.client.GetExtractorVersion(ctx, input.ExtractorID, input.VersionID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get extractor version: %w", err)
	}

	return nil, result, nil
}

// ListExtractorVersionsInput defines the input schema for list_extractor_versions tool
type ListExtractorVersionsInput struct {
	ExtractorID   string  `json:"extractor_id" jsonschema:"ID of the extractor"`
	MaxPageSize   *int    `json:"max_page_size,omitempty" jsonschema:"Maximum number of items per page (default: 10)"`
	NextPageToken *string `json:"next_page_token,omitempty" jsonschema:"Token for the next page of results"`
	SortBy        *string `json:"sort_by,omitempty" jsonschema:"Sort by field: updatedAt or createdAt (default: updatedAt)"`
	SortDir       *string `json:"sort_dir,omitempty" jsonschema:"Sort direction: asc or desc (default: desc)"`
}

// ListExtractorVersionsOutput wraps the extractor versions list with pagination
type ListExtractorVersionsOutput struct {
	Versions      []dto.ExtractorVersion `json:"versions"`
	NextPageToken *string                `json:"next_page_token,omitempty"`
}

// HandleListExtractorVersions handles the list_extractor_versions tool call
func (h *ExtractorHandlers) HandleListExtractorVersions(ctx context.Context, req *mcp.CallToolRequest, input ListExtractorVersionsInput) (*mcp.CallToolResult, *ListExtractorVersionsOutput, error) {
	if input.ExtractorID == "" {
		return nil, nil, fmt.Errorf("extractor_id is required")
	}

	pagination := &dto.PaginationParams{
		MaxPageSize:   input.MaxPageSize,
		NextPageToken: input.NextPageToken,
		SortBy:        input.SortBy,
		SortDir:       input.SortDir,
	}

	result, err := h.client.ListExtractorVersions(ctx, input.ExtractorID, pagination)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list extractor versions: %w", err)
	}

	if result.Data == nil {
		result.Data = []dto.ExtractorVersion{}
	}
	return nil, &ListExtractorVersionsOutput{
		Versions:      result.Data,
		NextPageToken: result.NextPageToken,
	}, nil
}
