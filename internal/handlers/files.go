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

// ListFilesOutput wraps the files list
type ListFilesOutput struct {
	Files []dto.File `json:"files"`
}

// HandleListFiles handles the list_files tool call
func (h *FileHandlers) HandleListFiles(ctx context.Context, req *mcp.CallToolRequest, input struct{}) (*mcp.CallToolResult, *ListFilesOutput, error) {
	// Call API
	result, err := h.client.ListFiles(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to list files: %w", err)
	}

	// Ensure non-nil slice
	if result == nil {
		result = []dto.File{}
	}
	return nil, &ListFilesOutput{Files: result}, nil
}
