package registry

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/caiomoura/extend-mcp-server/internal/client"
	"github.com/caiomoura/extend-mcp-server/internal/config"
	"github.com/caiomoura/extend-mcp-server/internal/handlers"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ToolCategory represents the category of an MCP tool
type ToolCategory string

const (
	CategoryProcessors ToolCategory = "processors"
	CategoryWorkflows  ToolCategory = "workflows"
	CategoryParse      ToolCategory = "parse"
	CategoryExtractors ToolCategory = "extractors"
	CategoryFiles      ToolCategory = "files"
)

// isCategoryEnabled checks if a tool category is enabled in the configuration
func isCategoryEnabled(cfg *config.Config, category ToolCategory) bool {
	switch category {
	case CategoryProcessors:
		return cfg.EnableProcessors
	case CategoryWorkflows:
		return cfg.EnableWorkflows
	case CategoryParse:
		return cfg.EnableParse
	case CategoryExtractors:
		return cfg.EnableExtractors
	case CategoryFiles:
		return cfg.EnableFiles
	}
	return false
}

// RegisterAllTools registers all read-only MCP tools based on configuration
func RegisterAllTools(server *mcp.Server, cfg *config.Config, extendClient *client.Client) int {
	// Create handlers
	processorHandlers := handlers.NewProcessorHandlers(extendClient)
	workflowHandlers := handlers.NewWorkflowHandlers(extendClient)
	parseHandlers := handlers.NewParseHandlers(extendClient)
	extractorHandlers := handlers.NewExtractorHandlers(extendClient)
	fileHandlers := handlers.NewFileHandlers(extendClient)

	// Tool registration counter
	toolCount := 0

	// Register Processor tools (read-only)
	if isCategoryEnabled(cfg, CategoryProcessors) {
		mcp.AddTool(server, &mcp.Tool{
			Name:        "get_processor_run",
			Description: "Get the status and result of a processor execution",
		}, processorHandlers.HandleGetProcessorRun)
		toolCount++

		mcp.AddTool(server, &mcp.Tool{
			Name:        "list_processors",
			Description: "List all available processors",
			InputSchema: json.RawMessage(`{"type": "object"}`),
		}, processorHandlers.HandleListProcessors)
		toolCount++
	}

	// Register Workflow tools (read-only)
	if isCategoryEnabled(cfg, CategoryWorkflows) {
		mcp.AddTool(server, &mcp.Tool{
			Name:        "get_workflow_run",
			Description: "Get the status and result of a workflow execution",
		}, workflowHandlers.HandleGetWorkflowRun)
		toolCount++

		mcp.AddTool(server, &mcp.Tool{
			Name:        "list_workflows",
			Description: "List all available workflows",
			InputSchema: json.RawMessage(`{"type": "object"}`),
		}, workflowHandlers.HandleListWorkflows)
		toolCount++

		mcp.AddTool(server, &mcp.Tool{
			Name:        "list_workflow_runs",
			Description: "List all workflow execution runs with status, timing and metadata",
			InputSchema: json.RawMessage(`{"type": "object"}`),
		}, workflowHandlers.HandleListWorkflowRuns)
		toolCount++
	}

	// Register Parse tools (read-only)
	if isCategoryEnabled(cfg, CategoryParse) {
		mcp.AddTool(server, &mcp.Tool{
			Name:        "get_parse_run",
			Description: "Get the result of a parse operation",
		}, parseHandlers.HandleGetParseRun)
		toolCount++
	}

	// Register Extractor tools (read-only)
	if isCategoryEnabled(cfg, CategoryExtractors) {
		mcp.AddTool(server, &mcp.Tool{
			Name:        "list_extract_runs",
			Description: "List all extractor execution runs",
			InputSchema: json.RawMessage(`{"type": "object"}`),
		}, extractorHandlers.HandleListExtractRuns)
		toolCount++

		mcp.AddTool(server, &mcp.Tool{
			Name:        "get_extractor",
			Description: "Get details of a specific extractor",
		}, extractorHandlers.HandleGetExtractor)
		toolCount++

		mcp.AddTool(server, &mcp.Tool{
			Name:        "list_extractors",
			Description: "List all available extractors",
			InputSchema: json.RawMessage(`{"type": "object"}`),
		}, extractorHandlers.HandleListExtractors)
		toolCount++

		mcp.AddTool(server, &mcp.Tool{
			Name:        "get_extractor_version",
			Description: "Get a specific version of an extractor",
		}, extractorHandlers.HandleGetExtractorVersion)
		toolCount++

		mcp.AddTool(server, &mcp.Tool{
			Name:        "list_extractor_versions",
			Description: "List all versions of a specific extractor",
		}, extractorHandlers.HandleListExtractorVersions)
		toolCount++
	}

	return toolCount
}

// LogServerInfo logs server startup information
func LogServerInfo(cfg *config.Config, toolCount int) {
	enabledCategories := []string{}
	if cfg.EnableProcessors {
		enabledCategories = append(enabledCategories, "processors")
	}
	if cfg.EnableWorkflows {
		enabledCategories = append(enabledCategories, "workflows")
	}
	if cfg.EnableParse {
		enabledCategories = append(enabledCategories, "parse")
	}
	if cfg.EnableExtractors {
		enabledCategories = append(enabledCategories, "extractors")
	}

	fmt.Fprintf(os.Stderr, "Extend AI MCP Server starting...\n")
	fmt.Fprintf(os.Stderr, "API Version: %s\n", cfg.APIVersion)
	fmt.Fprintf(os.Stderr, "Base URL: %s\n", cfg.BaseURL)
	fmt.Fprintf(os.Stderr, "Enabled Categories: %v\n", enabledCategories)
	fmt.Fprintf(os.Stderr, "Registered %d tools (read-only)\n", toolCount)
}
