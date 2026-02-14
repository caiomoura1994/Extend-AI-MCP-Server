package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents the Extend AI API client
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
	apiVersion string
}

// NewClient creates a new Extend AI API client
func NewClient(baseURL, apiKey, apiVersion string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL:    baseURL,
		apiKey:     apiKey,
		apiVersion: apiVersion,
	}
}

// doRequest performs an HTTP request with proper headers
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	url := c.baseURL + path
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("x-extend-api-version", c.apiVersion)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return nil, fmt.Errorf("API error (status %d): failed to decode error response", resp.StatusCode)
		}
		return nil, fmt.Errorf("API error (status %d): %s - %s", resp.StatusCode, errResp.Error, errResp.Message)
	}

	return resp, nil
}

// --- Processor Methods ---

// GetProcessorRun retrieves the status and result of a processor run
func (c *Client) GetProcessorRun(ctx context.Context, runID string) (*ProcessorRun, error) {
	path := fmt.Sprintf("/processor_runs/%s", runID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GetProcessorRunResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result.ProcessorRun, nil
}

// ListProcessors retrieves all available processors
func (c *Client) ListProcessors(ctx context.Context) ([]Processor, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/processors", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ListProcessorsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Processors, nil
}

// --- Workflow Methods ---

// GetWorkflowRun retrieves the status and result of a workflow run
func (c *Client) GetWorkflowRun(ctx context.Context, runID string) (*WorkflowRun, error) {
	path := fmt.Sprintf("/workflows/runs/%s", runID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result WorkflowRun
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// ListWorkflowRuns retrieves all workflow runs
func (c *Client) ListWorkflowRuns(ctx context.Context) ([]WorkflowRunSummary, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/workflow_runs", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ListWorkflowRunsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.WorkflowRuns, nil
}

// ListWorkflows retrieves all available workflows
func (c *Client) ListWorkflows(ctx context.Context) ([]Workflow, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/workflows", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ListWorkflowsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Workflows, nil
}

// --- File Methods ---

// ListFiles retrieves all files
func (c *Client) ListFiles(ctx context.Context) ([]File, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/files", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ListFilesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Files, nil
}

// --- Parse Methods ---

// GetParseRun retrieves the result of a parse operation
func (c *Client) GetParseRun(ctx context.Context, runID string) (*ParserRun, error) {
	path := fmt.Sprintf("/parser_runs/%s", runID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result GetParserRunResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result.ParserRun, nil
}

// --- Extractor Methods ---

// ListExtractRuns retrieves all extract runs
func (c *Client) ListExtractRuns(ctx context.Context) ([]ExtractRun, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/extract_runs", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ListExtractRunsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Data, nil
}

// GetExtractor retrieves a specific extractor by ID
func (c *Client) GetExtractor(ctx context.Context, extractorID string) (*Extractor, error) {
	path := fmt.Sprintf("/extractors/%s", extractorID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Extractor
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// ListExtractors retrieves all extractors
func (c *Client) ListExtractors(ctx context.Context) ([]Extractor, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/extractors", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ListExtractorsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Data, nil
}

// GetExtractorVersion retrieves a specific version of an extractor
func (c *Client) GetExtractorVersion(ctx context.Context, extractorID, versionID string) (*ExtractorVersion, error) {
	path := fmt.Sprintf("/extractors/%s/versions/%s", extractorID, versionID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ExtractorVersion
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// ListExtractorVersions retrieves all versions of a specific extractor
func (c *Client) ListExtractorVersions(ctx context.Context, extractorID string) ([]ExtractorVersion, error) {
	path := fmt.Sprintf("/extractors/%s/versions", extractorID)
	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ListExtractorVersionsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Data, nil
}
