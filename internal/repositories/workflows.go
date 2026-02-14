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

// ListWorkflowRuns retrieves all workflow runs
func (c *Client) ListWorkflowRuns(ctx context.Context) ([]dto.WorkflowRunSummary, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/workflow_runs", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListWorkflowRunsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.WorkflowRuns, nil
}

// ListWorkflows retrieves all available workflows
func (c *Client) ListWorkflows(ctx context.Context) ([]dto.Workflow, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/workflows", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListWorkflowsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Workflows, nil
}
