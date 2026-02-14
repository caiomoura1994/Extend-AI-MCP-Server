package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
)

// ListFiles retrieves files with optional pagination
func (c *Client) ListFiles(ctx context.Context, pagination *dto.PaginationParams) (*dto.ListFilesResponse, error) {
	path := "/files" + pagination.QueryString()
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListFilesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
