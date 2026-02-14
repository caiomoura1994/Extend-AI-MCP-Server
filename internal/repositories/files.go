package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
)

// ListFiles retrieves all files
func (c *Client) ListFiles(ctx context.Context) ([]dto.File, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/files", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListFilesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Files, nil
}
