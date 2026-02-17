package dto

import (
	"net/url"
	"time"
)

// Workflow represents a workflow configuration
type Workflow struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	Description string                   `json:"description,omitempty"`
	Steps       []map[string]interface{} `json:"steps,omitempty"`
	CreatedAt   time.Time                `json:"createdAt"`
	UpdatedAt   time.Time                `json:"updatedAt"`
}

// ListWorkflowsResponse wraps the list workflows API response
type ListWorkflowsResponse struct {
	Success       bool       `json:"success"`
	Workflows     []Workflow `json:"workflows"`
	NextPageToken *string    `json:"nextPageToken"`
}

// WorkflowRun represents a workflow execution
type WorkflowRun struct {
	ID         string                 `json:"id"`
	WorkflowID string                 `json:"workflow_id"`
	Status     string                 `json:"status"`
	Progress   float64                `json:"progress,omitempty"`
	Steps      []WorkflowStep         `json:"steps,omitempty"`
	Result     map[string]interface{} `json:"result,omitempty"`
	Error      string                 `json:"error,omitempty"`
	CreatedAt  time.Time              `json:"createdAt"`
	UpdatedAt  time.Time              `json:"updatedAt"`
}

// WorkflowStep represents a step in a workflow execution
type WorkflowStep struct {
	Name      string                 `json:"name"`
	Status    string                 `json:"status"`
	Result    map[string]interface{} `json:"result,omitempty"`
	Error     string                 `json:"error,omitempty"`
	StartedAt time.Time              `json:"startedAt,omitempty"`
	EndedAt   time.Time              `json:"endedAt,omitempty"`
}

// WorkflowRunFilters represents optional filter parameters for listing workflow runs
type WorkflowRunFilters struct {
	Status           *string `json:"status,omitempty"`
	WorkflowID       *string `json:"workflow_id,omitempty"`
	BatchID          *string `json:"batch_id,omitempty"`
	FileNameContains *string `json:"file_name_contains,omitempty"`
}

// QueryValues returns url.Values for the workflow run filters
func (f *WorkflowRunFilters) QueryValues() url.Values {
	params := url.Values{}
	if f == nil {
		return params
	}
	if f.Status != nil && *f.Status != "" {
		params.Set("status", *f.Status)
	}
	if f.WorkflowID != nil && *f.WorkflowID != "" {
		params.Set("workflowId", *f.WorkflowID)
	}
	if f.BatchID != nil && *f.BatchID != "" {
		params.Set("batchId", *f.BatchID)
	}
	if f.FileNameContains != nil && *f.FileNameContains != "" {
		params.Set("fileNameContains", *f.FileNameContains)
	}
	return params
}

// WorkflowRunSummary represents a workflow run summary (GET /workflow_runs)
type WorkflowRunSummary struct {
	ID                string                 `json:"id"`
	Status            string                 `json:"status"`
	InitialRunAt      *time.Time             `json:"initialRunAt,omitempty"`
	ReviewedByUser    string                 `json:"reviewedByUser,omitempty"`
	ReviewedAt        *time.Time             `json:"reviewedAt,omitempty"`
	StartTime         *time.Time             `json:"startTime,omitempty"`
	EndTime           *time.Time             `json:"endTime,omitempty"`
	WorkflowID        string                 `json:"workflowId"`
	WorkflowName      string                 `json:"workflowName"`
	WorkflowVersionID string                 `json:"workflowVersionId"`
	BatchID           string                 `json:"batchId,omitempty"`
	RejectionNote     string                 `json:"rejectionNote,omitempty"`
	DashboardURL      string                 `json:"dashboardUrl,omitempty"`
	CreatedAt         time.Time              `json:"createdAt"`
	UpdatedAt         time.Time              `json:"updatedAt"`
	Usage             map[string]interface{} `json:"usage,omitempty"`
}

// ListWorkflowRunsResponse wraps the list workflow runs API response
type ListWorkflowRunsResponse struct {
	Success       bool                 `json:"success"`
	WorkflowRuns  []WorkflowRunSummary `json:"workflowRuns"`
	NextPageToken *string              `json:"nextPageToken"`
}
