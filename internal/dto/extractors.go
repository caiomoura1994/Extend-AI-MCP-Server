package dto

import "time"

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
