package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
)

// ListExtractRuns retrieves all extract runs
func (c *Client) ListExtractRuns(ctx context.Context) ([]dto.ExtractRun, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/extract_runs", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListExtractRunsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Data, nil
}

// GetExtractor retrieves a specific extractor by ID
func (c *Client) GetExtractor(ctx context.Context, extractorID string) (*dto.Extractor, error) {
	path := fmt.Sprintf("/extractors/%s", extractorID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.Extractor
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// ListExtractors retrieves all extractors
func (c *Client) ListExtractors(ctx context.Context) ([]dto.Extractor, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/extractors", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListExtractorsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Data, nil
}

// GetExtractorVersion retrieves a specific version of an extractor
func (c *Client) GetExtractorVersion(ctx context.Context, extractorID, versionID string) (*dto.ExtractorVersion, error) {
	path := fmt.Sprintf("/extractors/%s/versions/%s", extractorID, versionID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ExtractorVersion
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// ListExtractorVersions retrieves all versions of a specific extractor
func (c *Client) ListExtractorVersions(ctx context.Context, extractorID string) ([]dto.ExtractorVersion, error) {
	path := fmt.Sprintf("/extractors/%s/versions", extractorID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListExtractorVersionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Data, nil
}
