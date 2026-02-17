package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
)

// GetWorkflowRun retrieves the status and result of a workflow run
func (c *Client) GetWorkflowRun(ctx context.Context, runID string) (*dto.WorkflowRun, error) {
	path := fmt.Sprintf("/workflows/runs/%s", runID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.WorkflowRun
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// ListWorkflowRuns retrieves workflow runs with optional pagination and filters
func (c *Client) ListWorkflowRuns(ctx context.Context, pagination *dto.PaginationParams, filters *dto.WorkflowRunFilters) (*dto.ListWorkflowRunsResponse, error) {
	path := "/workflow_runs" + dto.BuildQueryString(pagination.QueryValues(), filters.QueryValues())
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListWorkflowRunsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// ListWorkflows retrieves available workflows with optional pagination
func (c *Client) ListWorkflows(ctx context.Context, pagination *dto.PaginationParams) (*dto.ListWorkflowsResponse, error) {
	path := "/workflows" + pagination.QueryString()
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListWorkflowsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
