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

	paginationNote := " PAGINATION: Results are paginated (default 10 per page). If next_page_token is present in the response, you MUST call this tool again with that token to get the next page. Repeat until next_page_token is absent to retrieve all results."

	// Processor tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_processor_run",
		Description: "Get the status and result of a processor execution.",
	}, processorHandlers.HandleGetProcessorRun)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_processors",
		Description: "List available processors." + paginationNote,
	}, processorHandlers.HandleListProcessors)

	// Workflow tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_workflow_run",
		Description: "Get the status and result of a workflow execution.",
	}, workflowHandlers.HandleGetWorkflowRun)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_workflows",
		Description: "List available workflows." + paginationNote,
	}, workflowHandlers.HandleListWorkflows)

	dateFilterNote := " DATE FILTERING: Use created_after and/or created_before (RFC 3339, e.g. '2026-02-17T00:00:00Z') to filter by date range. When date filters are used, auto-pagination is handled internally and all matching results are returned at once."

	mcp.AddTool(server, &mcp.Tool{
		Name: "list_workflow_runs",
		Description: "List workflow execution runs with status, timing, and metadata. " +
			"Filter by status (PENDING/PROCESSING/NEEDS_REVIEW/REJECTED/PROCESSED/FAILED/CANCELLED/CANCELLING), workflow_id, batch_id, or file_name_contains." +
			dateFilterNote +
			paginationNote,
	}, workflowHandlers.HandleListWorkflowRuns)

	// Parse tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_parse_run",
		Description: "Get the result of a parse operation.",
	}, parseHandlers.HandleGetParseRun)

	// Extractor tools
	mcp.AddTool(server, &mcp.Tool{
		Name: "list_extract_runs",
		Description: "List extractor execution runs. " +
			"Filter by status (PROCESSING/PROCESSED/FAILED/CANCELLED), extractor_id, source (API/WORKFLOW_RUN/STUDIO/ADMIN/BATCH_PROCESSOR_RUN/PLAYGROUND/WORKFLOW_CONFIGURATION), source_id, or file_name_contains." +
			dateFilterNote +
			paginationNote,
	}, extractorHandlers.HandleListExtractRuns)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_extractor",
		Description: "Get details of a specific extractor.",
	}, extractorHandlers.HandleGetExtractor)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_extractors",
		Description: "List available extractors." + paginationNote,
	}, extractorHandlers.HandleListExtractors)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_extractor_version",
		Description: "Get a specific version of an extractor.",
	}, extractorHandlers.HandleGetExtractorVersion)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_extractor_versions",
		Description: "List all versions of a specific extractor." + paginationNote,
	}, extractorHandlers.HandleListExtractorVersions)

	// File tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_files",
		Description: "List files uploaded to the account." + paginationNote,
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
