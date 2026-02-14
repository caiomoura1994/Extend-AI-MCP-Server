package dto

import "time"

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
