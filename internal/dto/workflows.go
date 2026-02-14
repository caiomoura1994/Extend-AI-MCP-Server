package dto

import "time"

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
