package dto

import "time"

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
	Object             string                   `json:"object"`
	ID                 string                   `json:"id"`
	ProcessorID        string                   `json:"processorId"`
	ProcessorVersionID string                   `json:"processorVersionId,omitempty"`
	ProcessorName      string                   `json:"processorName,omitempty"`
	Status             string                   `json:"status"` // PENDING, PROCESSING, PROCESSED, FAILED, CANCELLED
	Type               string                   `json:"type,omitempty"`
	Output             map[string]interface{}   `json:"output,omitempty"`
	InitialOutput      map[string]interface{}   `json:"initialOutput,omitempty"`
	ReviewedOutput     map[string]interface{}   `json:"reviewedOutput,omitempty"`
	FailureReason      string                   `json:"failureReason,omitempty"`
	FailureMessage     string                   `json:"failureMessage,omitempty"`
	Metadata           map[string]interface{}   `json:"metadata,omitempty"`
	Reviewed           bool                     `json:"reviewed"`
	Edited             bool                     `json:"edited"`
	Config             map[string]interface{}   `json:"config,omitempty"`
	Files              []map[string]interface{} `json:"files,omitempty"`
	URL                string                   `json:"url,omitempty"`
	Usage              map[string]interface{}   `json:"usage,omitempty"`
}

// GetProcessorRunResponse wraps the get processor run API response
type GetProcessorRunResponse struct {
	Success      bool         `json:"success"`
	ProcessorRun ProcessorRun `json:"processorRun"`
}
