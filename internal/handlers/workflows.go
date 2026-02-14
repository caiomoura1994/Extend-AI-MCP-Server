package handlers

import (
	"context"
	"fmt"

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
	// Validate
	if input.RunID == "" {
		return nil, nil, fmt.Errorf("run_id is required")
	}

	// Call API
	result, err := h.client.GetWorkflowRun(ctx, input.RunID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get workflow run: %w", err)
	}

	// Return result
	return nil, result, nil
}

// ListWorkflowRunsOutput wraps the workflow runs list
type ListWorkflowRunsOutput struct {
	WorkflowRuns []dto.WorkflowRunSummary `json:"workflow_runs"`
}

// HandleListWorkflowRuns handles the list_workflow_runs tool call
func (h *WorkflowHandlers) HandleListWorkflowRuns(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, *ListWorkflowRunsOutput, error) {
	// Call API
	result, err := h.client.ListWorkflowRuns(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list workflow runs: %w", err)
	}

	// Ensure non-nil slice
	if result == nil {
		result = []dto.WorkflowRunSummary{}
	}
	return nil, &ListWorkflowRunsOutput{WorkflowRuns: result}, nil
}

// ListWorkflowsOutput wraps the workflows list
type ListWorkflowsOutput struct {
	Workflows []dto.Workflow `json:"workflows"`
}

// HandleListWorkflows handles the list_workflows tool call
func (h *WorkflowHandlers) HandleListWorkflows(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, *ListWorkflowsOutput, error) {
	// Call API
	result, err := h.client.ListWorkflows(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list workflows: %w", err)
	}

	// Ensure non-nil slice (MCP SDK requires array, not null)
	if result == nil {
		result = []dto.Workflow{}
	}
	return nil, &ListWorkflowsOutput{Workflows: result}, nil
}
