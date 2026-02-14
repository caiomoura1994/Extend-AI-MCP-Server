package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
)

// GetParseRun retrieves the result of a parse operation
func (c *Client) GetParseRun(ctx context.Context, runID string) (*dto.ParserRun, error) {
	path := fmt.Sprintf("/parser_runs/%s", runID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.GetParserRunResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result.ParserRun, nil
}
