package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
)

// GetProcessorRun retrieves the status and result of a processor run
func (c *Client) GetProcessorRun(ctx context.Context, runID string) (*dto.ProcessorRun, error) {
	path := fmt.Sprintf("/processor_runs/%s", runID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.GetProcessorRunResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result.ProcessorRun, nil
}

// ListProcessors retrieves all available processors
func (c *Client) ListProcessors(ctx context.Context) ([]dto.Processor, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/processors", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListProcessorsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Processors, nil
}
