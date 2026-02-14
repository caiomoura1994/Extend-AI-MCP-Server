package handlers

import (
	"context"
	"fmt"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
	"github.com/caiomoura/extend-mcp-server/internal/repositories"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// FileHandlers contains all file-related MCP handlers
type FileHandlers struct {
	client *repositories.Client
}

// NewFileHandlers creates a new file handlers instance
func NewFileHandlers(c *repositories.Client) *FileHandlers {
	return &FileHandlers{client: c}
}

// ListFilesInput defines the input schema for list_files tool
type ListFilesInput struct {
	MaxPageSize   *int    `json:"max_page_size,omitempty" jsonschema:"Maximum number of items per page (default: 10)"`
	NextPageToken *string `json:"next_page_token,omitempty" jsonschema:"Token for the next page of results"`
	SortBy        *string `json:"sort_by,omitempty" jsonschema:"Sort by field: updatedAt or createdAt (default: updatedAt)"`
	SortDir       *string `json:"sort_dir,omitempty" jsonschema:"Sort direction: asc or desc (default: desc)"`
}

// ListFilesOutput wraps the files list with pagination
type ListFilesOutput struct {
	Files         []dto.File `json:"files"`
	NextPageToken *string    `json:"next_page_token,omitempty"`
}

// HandleListFiles handles the list_files tool call
func (h *FileHandlers) HandleListFiles(ctx context.Context, req *mcp.CallToolRequest, input ListFilesInput) (*mcp.CallToolResult, *ListFilesOutput, error) {
	pagination := &dto.PaginationParams{
		MaxPageSize:   input.MaxPageSize,
		NextPageToken: input.NextPageToken,
		SortBy:        input.SortBy,
		SortDir:       input.SortDir,
	}

	result, err := h.client.ListFiles(ctx, pagination)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list files: %w", err)
	}

	if result.Files == nil {
		result.Files = []dto.File{}
	}
	return nil, &ListFilesOutput{
		Files:         result.Files,
		NextPageToken: result.NextPageToken,
	}, nil
}
