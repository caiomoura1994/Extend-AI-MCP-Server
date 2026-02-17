package handlers

import (
	"context"
	"fmt"
	"time"

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
	MaxPageSize      *int    `json:"max_page_size,omitempty" jsonschema:"Maximum number of items per page (default: 10). Ignored when date filters are used."`
	NextPageToken    *string `json:"next_page_token,omitempty" jsonschema:"Token for the next page of results. Pass the value from a previous response to get the next page. Ignored when date filters are used."`
	SortBy           *string `json:"sort_by,omitempty" jsonschema:"Sort by field: updatedAt or createdAt (default: updatedAt). Forced to createdAt when date filters are used."`
	SortDir          *string `json:"sort_dir,omitempty" jsonschema:"Sort direction: asc or desc (default: desc). Forced to desc when date filters are used."`
	Status           *string `json:"status,omitempty" jsonschema:"Filter by status: PROCESSING, PROCESSED, FAILED, or CANCELLED"`
	ExtractorID      *string `json:"extractor_id,omitempty" jsonschema:"Filter by extractor ID"`
	Source           *string `json:"source,omitempty" jsonschema:"Filter by source: API, WORKFLOW_RUN, STUDIO, ADMIN, BATCH_PROCESSOR_RUN, PLAYGROUND, or WORKFLOW_CONFIGURATION"`
	SourceID         *string `json:"source_id,omitempty" jsonschema:"Filter by source ID"`
	FileNameContains *string `json:"file_name_contains,omitempty" jsonschema:"Filter runs where the file name contains this string"`
	CreatedAfter     *string `json:"created_after,omitempty" jsonschema:"Return only runs created after this date (RFC 3339 format, e.g. 2026-02-17T00:00:00Z). Enables auto-pagination to collect all matching results."`
	CreatedBefore    *string `json:"created_before,omitempty" jsonschema:"Return only runs created before this date (RFC 3339 format, e.g. 2026-02-18T00:00:00Z). Enables auto-pagination to collect all matching results."`
}

// ListExtractRunsOutput wraps the extract runs list with pagination
type ListExtractRunsOutput struct {
	ExtractRuns   []dto.ExtractRun `json:"extract_runs"`
	NextPageToken *string          `json:"next_page_token,omitempty"`
}

// HandleListExtractRuns handles the list_extract_runs tool call
func (h *ExtractorHandlers) HandleListExtractRuns(ctx context.Context, req *mcp.CallToolRequest, input ListExtractRunsInput) (*mcp.CallToolResult, *ListExtractRunsOutput, error) {
	filters := &dto.ExtractRunFilters{
		Status:           input.Status,
		ExtractorID:      input.ExtractorID,
		Source:           input.Source,
		SourceID:         input.SourceID,
		FileNameContains: input.FileNameContains,
	}

	// Parse optional date filters
	createdAfter, err := ParseOptionalTime(input.CreatedAfter)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid created_after: %w", err)
	}
	createdBefore, err := ParseOptionalTime(input.CreatedBefore)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid created_before: %w", err)
	}

	dateParams := DateFilterParams{
		CreatedAfter:  createdAfter,
		CreatedBefore: createdBefore,
	}

	// When date filters are active, auto-paginate and collect all matching results
	if dateParams.HasDateFilter() {
		sortBy := "createdAt"
		sortDir := "desc"
		pageSize := autoPageSize

		fetcher := func(ctx context.Context, token *string) ([]dto.ExtractRun, *string, error) {
			pagination := &dto.PaginationParams{
				MaxPageSize:   &pageSize,
				NextPageToken: token,
				SortBy:        &sortBy,
				SortDir:       &sortDir,
			}
			result, err := h.client.ListExtractRuns(ctx, pagination, filters)
			if err != nil {
				return nil, nil, err
			}
			if result.Data == nil {
				result.Data = []dto.ExtractRun{}
			}
			return result.Data, result.NextPageToken, nil
		}

		runs, err := FetchWithDateFilter(ctx, dateParams, fetcher, func(r dto.ExtractRun) time.Time {
			return r.CreatedAt
		})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to list extract runs with date filter: %w", err)
		}
		if runs == nil {
			runs = []dto.ExtractRun{}
		}
		return nil, &ListExtractRunsOutput{
			ExtractRuns: runs,
		}, nil
	}

	// Normal single-page path (no date filter)
	pagination := &dto.PaginationParams{
		MaxPageSize:   input.MaxPageSize,
		NextPageToken: input.NextPageToken,
		SortBy:        input.SortBy,
		SortDir:       input.SortDir,
	}

	result, err := h.client.ListExtractRuns(ctx, pagination, filters)
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
