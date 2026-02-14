package handlers_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/caiomoura/extend-mcp-server/internal/client"
	"github.com/caiomoura/extend-mcp-server/internal/handlers"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

var testImpl = &mcp.Implementation{Name: "test-server", Version: "0.0.1"}

// setupServer creates a server+client pair with all list tools registered.
// Uses an invalid API key so every API call triggers the error path.
func setupServer(t *testing.T) *mcp.ClientSession {
	t.Helper()

	c := client.NewClient("https://api.extend.ai", "invalid_key", "2025-04-21")

	extractorH := handlers.NewExtractorHandlers(c)
	processorH := handlers.NewProcessorHandlers(c)
	workflowH := handlers.NewWorkflowHandlers(c)
	parseH := handlers.NewParseHandlers(c)

	server := mcp.NewServer(testImpl, nil)

	// Register list tools (these have slice output fields)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_extract_runs",
		Description: "List all extractor execution runs",
		InputSchema: json.RawMessage(`{"type": "object"}`),
	}, extractorH.HandleListExtractRuns)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_extractors",
		Description: "List all available extractors",
		InputSchema: json.RawMessage(`{"type": "object"}`),
	}, extractorH.HandleListExtractors)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_processors",
		Description: "List all available processors",
		InputSchema: json.RawMessage(`{"type": "object"}`),
	}, processorH.HandleListProcessors)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_workflows",
		Description: "List all available workflows",
		InputSchema: json.RawMessage(`{"type": "object"}`),
	}, workflowH.HandleListWorkflows)

	// Register get tools (these return pointer-to-struct)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_processor_run",
		Description: "Get the status and result of a processor execution",
	}, processorH.HandleGetProcessorRun)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_workflow_run",
		Description: "Get the status and result of a workflow execution",
	}, workflowH.HandleGetWorkflowRun)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_parse_run",
		Description: "Get the result of a parse operation",
	}, parseH.HandleGetParseRun)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_extractor",
		Description: "Get details of a specific extractor",
	}, extractorH.HandleGetExtractor)

	// Connect via in-memory transport
	ct, st := mcp.NewInMemoryTransports()
	ctx := context.Background()

	_, err := server.Connect(ctx, st, nil)
	if err != nil {
		t.Fatal(err)
	}

	clientImpl := &mcp.Implementation{Name: "test-client", Version: "0.0.1"}
	mcpClient := mcp.NewClient(clientImpl, nil)
	cs, err := mcpClient.Connect(ctx, ct, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { cs.Close() })

	return cs
}

// TestListToolsErrorPath verifies that all list tools handle API errors gracefully
// without triggering the MCP SDK output validation bug (nil slice -> JSON null -> schema error).
func TestListToolsErrorPath(t *testing.T) {
	cs := setupServer(t)

	tools := []string{
		"list_extract_runs",
		"list_extractors",
		"list_processors",
		"list_workflows",
	}

	for _, toolName := range tools {
		t.Run(toolName, func(t *testing.T) {
			ctx := context.Background()
			result, err := cs.CallTool(ctx, &mcp.CallToolParams{
				Name:      toolName,
				Arguments: map[string]any{},
			})

			// A protocol-level error means something went very wrong
			if err != nil {
				errStr := err.Error()
				if strings.Contains(errStr, "validating tool output") {
					t.Fatalf("BUG: SDK output validation error (nil slice issue): %v", err)
				}
				// Other protocol errors are unexpected too
				t.Fatalf("Unexpected protocol error: %v", err)
			}

			// The result should indicate a tool error (invalid API key)
			if result == nil {
				t.Fatal("Expected non-nil result")
			}
			if !result.IsError {
				t.Errorf("Expected IsError=true (API should fail with invalid key)")
			}

			t.Logf("OK - tool %s error path handled correctly: IsError=%v", toolName, result.IsError)
		})
	}
}

// TestGetToolsAPIErrorPath verifies that get tools handle API errors gracefully
// when called with valid input but an invalid API key.
func TestGetToolsAPIErrorPath(t *testing.T) {
	cs := setupServer(t)

	// Call get tools with a fake ID -- the API will reject with 401 (invalid key)
	tools := []struct {
		name string
		args map[string]any
	}{
		{"get_processor_run", map[string]any{"run_id": "fake-id"}},
		{"get_workflow_run", map[string]any{"run_id": "fake-id"}},
		{"get_parse_run", map[string]any{"run_id": "fake-id"}},
		{"get_extractor", map[string]any{"extractor_id": "fake-id"}},
	}

	for _, tc := range tools {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			result, err := cs.CallTool(ctx, &mcp.CallToolParams{
				Name:      tc.name,
				Arguments: tc.args,
			})

			if err != nil {
				errStr := err.Error()
				if strings.Contains(errStr, "validating tool output") {
					t.Fatalf("BUG: SDK output validation error: %v", err)
				}
				// Other protocol errors are unexpected
				t.Fatalf("Unexpected protocol error: %v", err)
			}

			if result == nil {
				t.Fatal("Expected non-nil result")
			}
			if !result.IsError {
				t.Errorf("Expected IsError=true (API should fail with invalid key)")
			}

			t.Logf("OK - tool %s error path handled correctly: IsError=%v", tc.name, result.IsError)
		})
	}
}
