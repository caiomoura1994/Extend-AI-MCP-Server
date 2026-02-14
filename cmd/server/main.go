package main

import (
	"context"
	"log"

	"github.com/caiomoura/extend-mcp-server/internal/repositories"
	"github.com/caiomoura/extend-mcp-server/internal/config"
	"github.com/caiomoura/extend-mcp-server/internal/registry"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create Extend API client
	extendClient := repositories.NewClient(cfg.BaseURL, cfg.APIKey, cfg.APIVersion)

	// Create MCP server
	impl := &mcp.Implementation{
		Name:    "extend-mcp-server",
		Title:   "Extend AI MCP Server",
		Version: "1.0.0",
	}
	server := mcp.NewServer(impl, nil)

	// Register all tools based on configuration
	toolCount := registry.RegisterAllTools(server, cfg, extendClient)

	// Log server startup information
	registry.LogServerInfo(cfg, toolCount)

	// Start server via stdio transport
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
