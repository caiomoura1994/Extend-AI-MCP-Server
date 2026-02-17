package dto

import (
	"net/url"
	"time"
)

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

// ExtractRunFilters represents optional filter parameters for listing extract runs
type ExtractRunFilters struct {
	Status           *string `json:"status,omitempty"`
	ExtractorID      *string `json:"extractor_id,omitempty"`
	Source           *string `json:"source,omitempty"`
	SourceID         *string `json:"source_id,omitempty"`
	FileNameContains *string `json:"file_name_contains,omitempty"`
}

// QueryValues returns url.Values for the extract run filters
func (f *ExtractRunFilters) QueryValues() url.Values {
	params := url.Values{}
	if f == nil {
		return params
	}
	if f.Status != nil && *f.Status != "" {
		params.Set("status", *f.Status)
	}
	if f.ExtractorID != nil && *f.ExtractorID != "" {
		params.Set("extractorId", *f.ExtractorID)
	}
	if f.Source != nil && *f.Source != "" {
		params.Set("source", *f.Source)
	}
	if f.SourceID != nil && *f.SourceID != "" {
		params.Set("sourceId", *f.SourceID)
	}
	if f.FileNameContains != nil && *f.FileNameContains != "" {
		params.Set("fileNameContains", *f.FileNameContains)
	}
	return params
}

// ExtractRun represents an extractor execution run
type ExtractRun struct {
	Object           string                 `json:"object"`
	ID               string                 `json:"id"`
	Status           string                 `json:"status"` // PROCESSING, PROCESSED, FAILED, CANCELLED
	Extractor        *Extractor             `json:"extractor,omitempty"`
	ExtractorVersion *ExtractorVersion      `json:"extractorVersion,omitempty"`
	Reviewed         bool                   `json:"reviewed"`
	Edited           bool                   `json:"edited"`
	FailureReason    *string                `json:"failureReason,omitempty"`
	FailureMessage   *string                `json:"failureMessage,omitempty"`
	File             map[string]interface{} `json:"file,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	ParseRunID       *string                `json:"parseRunId,omitempty"`
	DashboardURL     string                 `json:"dashboardUrl,omitempty"`
	Usage            map[string]interface{} `json:"usage,omitempty"`
	CreatedAt        time.Time              `json:"createdAt"`
	UpdatedAt        time.Time              `json:"updatedAt"`
}

// ListExtractRunsResponse wraps the list extract runs API response
type ListExtractRunsResponse struct {
	Object        string       `json:"object"`
	Data          []ExtractRun `json:"data"`
	NextPageToken *string      `json:"nextPageToken"`
}
