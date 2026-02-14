package registry

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caiomoura/extend-mcp-server/internal/config"
	"github.com/caiomoura/extend-mcp-server/internal/handlers"
	"github.com/caiomoura/extend-mcp-server/internal/repositories"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterAllTools registers all read-only MCP tools
func RegisterAllTools(server *mcp.Server, extendClient *repositories.Client) int {
	// Create handlers
	processorHandlers := handlers.NewProcessorHandlers(extendClient)
	workflowHandlers := handlers.NewWorkflowHandlers(extendClient)
	parseHandlers := handlers.NewParseHandlers(extendClient)
	extractorHandlers := handlers.NewExtractorHandlers(extendClient)
	fileHandlers := handlers.NewFileHandlers(extendClient)

	// Processor tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_processor_run",
		Description: "Get the status and result of a processor execution",
	}, processorHandlers.HandleGetProcessorRun)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_processors",
		Description: "List all available processors",
		InputSchema: json.RawMessage(`{"type": "object"}`),
	}, processorHandlers.HandleListProcessors)

	// Workflow tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_workflow_run",
		Description: "Get the status and result of a workflow execution",
	}, workflowHandlers.HandleGetWorkflowRun)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_workflows",
		Description: "List all available workflows",
		InputSchema: json.RawMessage(`{"type": "object"}`),
	}, workflowHandlers.HandleListWorkflows)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_workflow_runs",
		Description: "List all workflow execution runs with status, timing and metadata",
		InputSchema: json.RawMessage(`{"type": "object"}`),
	}, workflowHandlers.HandleListWorkflowRuns)

	// Parse tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_parse_run",
		Description: "Get the result of a parse operation",
	}, parseHandlers.HandleGetParseRun)

	// Extractor tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_extract_runs",
		Description: "List all extractor execution runs",
		InputSchema: json.RawMessage(`{"type": "object"}`),
	}, extractorHandlers.HandleListExtractRuns)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_extractor",
		Description: "Get details of a specific extractor",
	}, extractorHandlers.HandleGetExtractor)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_extractors",
		Description: "List all available extractors",
		InputSchema: json.RawMessage(`{"type": "object"}`),
	}, extractorHandlers.HandleListExtractors)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_extractor_version",
		Description: "Get a specific version of an extractor",
	}, extractorHandlers.HandleGetExtractorVersion)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_extractor_versions",
		Description: "List all versions of a specific extractor",
	}, extractorHandlers.HandleListExtractorVersions)

	// File tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_files",
		Description: "List all files uploaded to the account",
		InputSchema: json.RawMessage(`{"type": "object"}`),
	}, fileHandlers.HandleListFiles)

	return 12
}

// LogServerInfo logs server startup information
func LogServerInfo(cfg *config.Config, toolCount int) {
	fmt.Fprintf(os.Stderr, "Extend AI MCP Server starting...\n")
	fmt.Fprintf(os.Stderr, "API Version: %s\n", cfg.APIVersion)
	fmt.Fprintf(os.Stderr, "Base URL: %s\n", cfg.BaseURL)
	fmt.Fprintf(os.Stderr, "Registered %d tools (read-only)\n", toolCount)
}
