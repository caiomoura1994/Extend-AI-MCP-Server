package registry

import (
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
		Description: "List available processors. Supports pagination via max_page_size (default: 10), next_page_token, sort_by, and sort_dir",
	}, processorHandlers.HandleListProcessors)

	// Workflow tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_workflow_run",
		Description: "Get the status and result of a workflow execution",
	}, workflowHandlers.HandleGetWorkflowRun)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_workflows",
		Description: "List available workflows. Supports pagination via max_page_size (default: 10), next_page_token, sort_by, and sort_dir",
	}, workflowHandlers.HandleListWorkflows)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_workflow_runs",
		Description: "List workflow execution runs with status, timing and metadata. Supports pagination via max_page_size (default: 10), next_page_token, sort_by, and sort_dir",
	}, workflowHandlers.HandleListWorkflowRuns)

	// Parse tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_parse_run",
		Description: "Get the result of a parse operation",
	}, parseHandlers.HandleGetParseRun)

	// Extractor tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_extract_runs",
		Description: "List extractor execution runs. Supports pagination via max_page_size (default: 10), next_page_token, sort_by, and sort_dir",
	}, extractorHandlers.HandleListExtractRuns)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_extractor",
		Description: "Get details of a specific extractor",
	}, extractorHandlers.HandleGetExtractor)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_extractors",
		Description: "List available extractors. Supports pagination via max_page_size (default: 10), next_page_token, sort_by, and sort_dir",
	}, extractorHandlers.HandleListExtractors)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_extractor_version",
		Description: "Get a specific version of an extractor",
	}, extractorHandlers.HandleGetExtractorVersion)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_extractor_versions",
		Description: "List all versions of a specific extractor. Supports pagination via max_page_size (default: 10), next_page_token, sort_by, and sort_dir",
	}, extractorHandlers.HandleListExtractorVersions)

	// File tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_files",
		Description: "List files uploaded to the account. Supports pagination via max_page_size (default: 10), next_page_token, sort_by, and sort_dir",
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
