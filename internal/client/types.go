package client

import "time"

// --- Processor Types ---

// ProcessorVersion represents a version of a processor
type ProcessorVersion struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

// Processor represents a processor configuration
type Processor struct {
	Object    string             `json:"object"`
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Type      string             `json:"type"` // EXTRACT, CLASSIFY, SPLITTER
	Versions  []ProcessorVersion `json:"versions,omitempty"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

// ListProcessorsResponse wraps the list processors API response
type ListProcessorsResponse struct {
	Success       bool        `json:"success"`
	Warning       *string     `json:"warning"`
	Processors    []Processor `json:"processors"`
	NextPageToken *string     `json:"nextPageToken"`
}

// ProcessorRun represents a processor execution (GET /processor_runs/{id})
type ProcessorRun struct {
	Object             string                 `json:"object"`
	ID                 string                 `json:"id"`
	ProcessorID        string                 `json:"processorId"`
	ProcessorVersionID string                 `json:"processorVersionId,omitempty"`
	ProcessorName      string                 `json:"processorName,omitempty"`
	Status             string                 `json:"status"` // PENDING, PROCESSING, PROCESSED, FAILED, CANCELLED
	Type               string                 `json:"type,omitempty"`
	Output             map[string]interface{} `json:"output,omitempty"`
	InitialOutput      map[string]interface{} `json:"initialOutput,omitempty"`
	ReviewedOutput     map[string]interface{} `json:"reviewedOutput,omitempty"`
	FailureReason      string                 `json:"failureReason,omitempty"`
	FailureMessage     string                 `json:"failureMessage,omitempty"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
	Reviewed           bool                   `json:"reviewed"`
	Edited             bool                   `json:"edited"`
	Config             map[string]interface{} `json:"config,omitempty"`
	Files              []map[string]interface{} `json:"files,omitempty"`
	URL                string                 `json:"url,omitempty"`
	Usage              map[string]interface{} `json:"usage,omitempty"`
}

// GetProcessorRunResponse wraps the get processor run API response
type GetProcessorRunResponse struct {
	Success      bool         `json:"success"`
	ProcessorRun ProcessorRun `json:"processorRun"`
}

// --- Workflow Types ---

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

// --- File Types ---

// FileMetadata contains metadata about a file
type FileMetadata struct {
	PageCount   *float64               `json:"pageCount,omitempty"`
	ParentSplit map[string]interface{} `json:"parentSplit,omitempty"`
}

// File represents a file in the Extend AI system
type File struct {
	Object       string                 `json:"object"`
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Type         string                 `json:"type,omitempty"`
	PresignedURL string                 `json:"presignedUrl,omitempty"`
	ParentFileID string                 `json:"parentFileId,omitempty"`
	Metadata     FileMetadata           `json:"metadata"`
	CreatedAt    time.Time              `json:"createdAt"`
	UpdatedAt    time.Time              `json:"updatedAt"`
	Usage        map[string]interface{} `json:"usage,omitempty"`
}

// ListFilesResponse wraps the list files API response
type ListFilesResponse struct {
	Success       bool    `json:"success"`
	Files         []File  `json:"files"`
	NextPageToken *string `json:"nextPageToken"`
}

// --- Parse Types ---

// ParserRunMetrics contains metrics about the parsing process
type ParserRunMetrics struct {
	ProcessingTimeMs float64 `json:"processingTimeMs"`
	PageCount        float64 `json:"pageCount"`
}

// ParserRunChunkMetadataPageRange represents the page range of a chunk
type ParserRunChunkMetadataPageRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// ParserRunChunkMetadata contains metadata about a chunk
type ParserRunChunkMetadata struct {
	PageRange ParserRunChunkMetadataPageRange `json:"pageRange"`
}

// ParserRunChunk represents a parsed chunk from a document
type ParserRunChunk struct {
	Object   string                 `json:"object"`
	Type     string                 `json:"type"` // page, document, section
	Content  string                 `json:"content"`
	Metadata ParserRunChunkMetadata `json:"metadata"`
	Blocks   []map[string]interface{} `json:"blocks,omitempty"`
}

// ParserRun represents a parser run result (GET /parser_runs/{id})
type ParserRun struct {
	Object        string                 `json:"object"` // parser_run or parser_run_status
	ID            string                 `json:"id"`
	FileID        string                 `json:"fileId,omitempty"`
	Status        string                 `json:"status"` // PROCESSING, PROCESSED, FAILED
	FailureReason string                 `json:"failureReason,omitempty"`
	Chunks        []ParserRunChunk       `json:"chunks,omitempty"`
	Metrics       *ParserRunMetrics      `json:"metrics,omitempty"`
	Config        map[string]interface{} `json:"config,omitempty"`
	Usage         map[string]interface{} `json:"usage,omitempty"`
}

// GetParserRunResponse wraps the get parser run API response
type GetParserRunResponse struct {
	Success   bool      `json:"success"`
	Warning   string    `json:"warning,omitempty"`
	ParserRun ParserRun `json:"parserRun"`
}

// --- Extractor Types ---

// Extractor represents an extractor configuration
type Extractor struct {
	Object    string    `json:"object"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ListExtractorsResponse wraps the list extractors API response
type ListExtractorsResponse struct {
	Object        string      `json:"object"`
	Data          []Extractor `json:"data"`
	NextPageToken *string     `json:"nextPageToken"`
}

// ExtractorVersion represents a version of an extractor
type ExtractorVersion struct {
	Object      string    `json:"object"`
	ID          string    `json:"id"`
	Description *string   `json:"description"`
	Version     string    `json:"version"`
	ExtractorID string    `json:"extractorId"`
	CreatedAt   time.Time `json:"createdAt"`
}

// ListExtractorVersionsResponse wraps the list extractor versions API response
type ListExtractorVersionsResponse struct {
	Object        string             `json:"object"`
	Data          []ExtractorVersion `json:"data"`
	NextPageToken *string            `json:"nextPageToken"`
}

// ExtractRun represents an extractor execution run
type ExtractRun struct {
	Object    string                 `json:"object"`
	ID        string                 `json:"id"`
	Status    string                 `json:"status"` // PROCESSING, PROCESSED, FAILED, CANCELLED
	Extractor *Extractor             `json:"extractor,omitempty"`
	Reviewed  bool                   `json:"reviewed"`
	Edited    bool                   `json:"edited"`
	File      map[string]interface{} `json:"file,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Usage     map[string]interface{} `json:"usage,omitempty"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

// ListExtractRunsResponse wraps the list extract runs API response
type ListExtractRunsResponse struct {
	Object        string       `json:"object"`
	Data          []ExtractRun `json:"data"`
	NextPageToken *string      `json:"nextPageToken"`
}

// --- Error Types ---

// ErrorResponse represents an API error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}
