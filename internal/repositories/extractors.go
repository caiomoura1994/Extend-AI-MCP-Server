package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caiomoura/extend-mcp-server/internal/dto"
)

// ListExtractRuns retrieves extract runs with optional pagination
func (c *Client) ListExtractRuns(ctx context.Context, pagination *dto.PaginationParams) (*dto.ListExtractRunsResponse, error) {
	path := "/extract_runs" + pagination.QueryString()
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListExtractRunsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
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

// ListExtractors retrieves extractors with optional pagination
func (c *Client) ListExtractors(ctx context.Context, pagination *dto.PaginationParams) (*dto.ListExtractorsResponse, error) {
	path := "/extractors" + pagination.QueryString()
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListExtractorsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
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

// ListExtractorVersions retrieves versions of a specific extractor with optional pagination
func (c *Client) ListExtractorVersions(ctx context.Context, extractorID string, pagination *dto.PaginationParams) (*dto.ListExtractorVersionsResponse, error) {
	path := fmt.Sprintf("/extractors/%s/versions", extractorID) + pagination.QueryString()
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result dto.ListExtractorVersionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
