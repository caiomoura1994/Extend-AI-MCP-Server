package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
	"github.com/caiomoura/extend-mcp-server/internal/repositories"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// WorkflowHandlers contains all workflow-related MCP handlers
type WorkflowHandlers struct {
	client *repositories.Client
}

// NewWorkflowHandlers creates a new workflow handlers instance
func NewWorkflowHandlers(c *repositories.Client) *WorkflowHandlers {
	return &WorkflowHandlers{client: c}
}

// GetWorkflowRunInput defines the input schema for get_workflow_run tool
type GetWorkflowRunInput struct {
	RunID string `json:"run_id" jsonschema:"ID of the workflow run to retrieve"`
}

// HandleGetWorkflowRun handles the get_workflow_run tool call
func (h *WorkflowHandlers) HandleGetWorkflowRun(ctx context.Context, req *mcp.CallToolRequest, input GetWorkflowRunInput) (*mcp.CallToolResult, *dto.WorkflowRun, error) {
	if input.RunID == "" {
		return nil, nil, fmt.Errorf("run_id is required")
	}

	result, err := h.client.GetWorkflowRun(ctx, input.RunID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get workflow run: %w", err)
	}

	return nil, result, nil
}

// ListWorkflowRunsInput defines the input schema for list_workflow_runs tool
type ListWorkflowRunsInput struct {
	MaxPageSize      *int    `json:"max_page_size,omitempty" jsonschema:"Maximum number of items per page (default: 10). Ignored when date filters are used."`
	NextPageToken    *string `json:"next_page_token,omitempty" jsonschema:"Token for the next page of results. Pass the value from a previous response to get the next page. Ignored when date filters are used."`
	SortBy           *string `json:"sort_by,omitempty" jsonschema:"Sort by field: updatedAt or createdAt (default: updatedAt). Forced to createdAt when date filters are used."`
	SortDir          *string `json:"sort_dir,omitempty" jsonschema:"Sort direction: asc or desc (default: desc). Forced to desc when date filters are used."`
	Status           *string `json:"status,omitempty" jsonschema:"Filter by status: PENDING, PROCESSING, NEEDS_REVIEW, REJECTED, PROCESSED, FAILED, CANCELLED, or CANCELLING"`
	WorkflowID       *string `json:"workflow_id,omitempty" jsonschema:"Filter by workflow ID"`
	BatchID          *string `json:"batch_id,omitempty" jsonschema:"Filter by batch ID"`
	FileNameContains *string `json:"file_name_contains,omitempty" jsonschema:"Filter runs where the file name contains this string"`
	CreatedAfter     *string `json:"created_after,omitempty" jsonschema:"Return only runs created after this date (RFC 3339 format, e.g. 2026-02-17T00:00:00Z). Enables auto-pagination to collect all matching results."`
	CreatedBefore    *string `json:"created_before,omitempty" jsonschema:"Return only runs created before this date (RFC 3339 format, e.g. 2026-02-18T00:00:00Z). Enables auto-pagination to collect all matching results."`
}

// ListWorkflowRunsOutput wraps the workflow runs list with pagination
type ListWorkflowRunsOutput struct {
	WorkflowRuns  []dto.WorkflowRunSummary `json:"workflow_runs"`
	NextPageToken *string                  `json:"next_page_token,omitempty"`
}

// HandleListWorkflowRuns handles the list_workflow_runs tool call
func (h *WorkflowHandlers) HandleListWorkflowRuns(ctx context.Context, req *mcp.CallToolRequest, input ListWorkflowRunsInput) (*mcp.CallToolResult, *ListWorkflowRunsOutput, error) {
	filters := &dto.WorkflowRunFilters{
		Status:           input.Status,
		WorkflowID:       input.WorkflowID,
		BatchID:          input.BatchID,
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

		fetcher := func(ctx context.Context, token *string) ([]dto.WorkflowRunSummary, *string, error) {
			pagination := &dto.PaginationParams{
				MaxPageSize:   &pageSize,
				NextPageToken: token,
				SortBy:        &sortBy,
				SortDir:       &sortDir,
			}
			result, err := h.client.ListWorkflowRuns(ctx, pagination, filters)
			if err != nil {
				return nil, nil, err
			}
			if result.WorkflowRuns == nil {
				result.WorkflowRuns = []dto.WorkflowRunSummary{}
			}
			return result.WorkflowRuns, result.NextPageToken, nil
		}

		runs, err := FetchWithDateFilter(ctx, dateParams, fetcher, func(r dto.WorkflowRunSummary) time.Time {
			return r.CreatedAt
		})
		if err != nil {
			return nil, nil, fmt.Errorf("failed to list workflow runs with date filter: %w", err)
		}
		if runs == nil {
			runs = []dto.WorkflowRunSummary{}
		}
		return nil, &ListWorkflowRunsOutput{
			WorkflowRuns: runs,
		}, nil
	}

	// Normal single-page path (no date filter)
	pagination := &dto.PaginationParams{
		MaxPageSize:   input.MaxPageSize,
		NextPageToken: input.NextPageToken,
		SortBy:        input.SortBy,
		SortDir:       input.SortDir,
	}

	result, err := h.client.ListWorkflowRuns(ctx, pagination, filters)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list workflow runs: %w", err)
	}

	if result.WorkflowRuns == nil {
		result.WorkflowRuns = []dto.WorkflowRunSummary{}
	}
	return nil, &ListWorkflowRunsOutput{
		WorkflowRuns:  result.WorkflowRuns,
		NextPageToken: result.NextPageToken,
	}, nil
}

// ListWorkflowsInput defines the input schema for list_workflows tool
type ListWorkflowsInput struct {
	MaxPageSize   *int    `json:"max_page_size,omitempty" jsonschema:"Maximum number of items per page (default: 10)"`
	NextPageToken *string `json:"next_page_token,omitempty" jsonschema:"Token for the next page of results"`
	SortBy        *string `json:"sort_by,omitempty" jsonschema:"Sort by field: updatedAt or createdAt (default: updatedAt)"`
	SortDir       *string `json:"sort_dir,omitempty" jsonschema:"Sort direction: asc or desc (default: desc)"`
}

// ListWorkflowsOutput wraps the workflows list with pagination
type ListWorkflowsOutput struct {
	Workflows     []dto.Workflow `json:"workflows"`
	NextPageToken *string        `json:"next_page_token,omitempty"`
}

// HandleListWorkflows handles the list_workflows tool call
func (h *WorkflowHandlers) HandleListWorkflows(ctx context.Context, req *mcp.CallToolRequest, input ListWorkflowsInput) (*mcp.CallToolResult, *ListWorkflowsOutput, error) {
	pagination := &dto.PaginationParams{
		MaxPageSize:   input.MaxPageSize,
		NextPageToken: input.NextPageToken,
		SortBy:        input.SortBy,
		SortDir:       input.SortDir,
	}

	result, err := h.client.ListWorkflows(ctx, pagination)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list workflows: %w", err)
	}

	if result.Workflows == nil {
		result.Workflows = []dto.Workflow{}
	}
	return nil, &ListWorkflowsOutput{
		Workflows:     result.Workflows,
		NextPageToken: result.NextPageToken,
	}, nil
}
