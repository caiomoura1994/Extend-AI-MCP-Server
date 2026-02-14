# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a Model Context Protocol (MCP) server implementation that integrates with the Extend AI API. It exposes Extend AI's document processing capabilities (extraction, classification, parsing, workflows) as MCP tools that can be used by AI assistants like Claude Desktop.

## Tech Stack

- **Language**: Go 1.23+
- **Framework**: MCP Go SDK (`github.com/modelcontextprotocol/go-sdk`)
- **Transport**: stdio (standard input/output)
- **API**: Extend AI REST API (https://api.extend.ai)

## Architecture

### Three-Layer Architecture

The codebase follows a clean separation of concerns:

1. **Entry Point** (`cmd/server/main.go`)
   - Loads configuration from environment variables
   - Creates Extend AI HTTP client
   - Initializes MCP server and registers all tools (15 total)
   - Starts stdio transport

2. **Handlers Layer** (`internal/handlers/`)
   - `processors.go` - 5 processor-related tools
   - `workflows.go` - 6 workflow-related tools
   - `parse.go` - 4 parse-related tools
   - Each handler validates input, calls the client, and formats MCP responses

3. **Client Layer** (`internal/client/`)
   - `extend.go` - HTTP client with all Extend AI API methods
   - `types.go` - Request/response type definitions
   - Handles authentication (Bearer token), API versioning, error handling

### Tool Categories

**Processors** (5 tools): Run individual document processing operations
- `run_processor`, `get_processor_run`, `list_processors`, `create_processor`, `delete_processor_run`

**Workflows** (6 tools): Orchestrate multi-step document processing
- `run_workflow`, `get_workflow_run`, `list_workflows`, `batch_run_workflow`, `create_workflow`, `delete_workflow_run`

**Parse** (4 tools): Convert documents to structured text
- `parse_file` (sync), `parse_file_async`, `get_parse_run`, `delete_parse_run`

## Development Commands

### Build

```bash
# Download dependencies
go mod download

# Build the binary
go build -o extend-mcp-server cmd/server/main.go

# Install globally (optional)
go install ./cmd/server
```

### Run Locally

```bash
# Set required environment variable
export EXTEND_API_KEY=your_api_key_here

# Optional environment variables
export EXTEND_API_VERSION=2024-11-14              # Default: 2024-11-14
export EXTEND_BASE_URL=https://api.extend.ai      # Default: https://api.extend.ai

# Run the server (communicates via stdio)
go run cmd/server/main.go
```

### Testing

There are currently no automated tests. The server is tested by:
1. Building the binary
2. Configuring it in Claude Desktop's `claude_desktop_config.json`
3. Using the tools through Claude Desktop's interface

To add tests in the future:
- Use Go's standard `testing` package
- Create `*_test.go` files alongside source files
- Run with `go test ./...`

## Configuration

The server requires environment variables (see `internal/config/config.go`):

- **EXTEND_API_KEY** (required): Your Extend AI API key
- **EXTEND_API_VERSION** (optional): API version, defaults to "2024-11-14"
- **EXTEND_BASE_URL** (optional): Base URL, defaults to "https://api.extend.ai"

For Claude Desktop integration, add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS) or `%APPDATA%\Claude\claude_desktop_config.json` (Windows):

```json
{
  "mcpServers": {
    "extend-ai": {
      "command": "/absolute/path/to/extend-mcp-server",
      "env": {
        "EXTEND_API_KEY": "your_api_key_here"
      }
    }
  }
}
```

## Code Patterns

### Adding a New Tool

1. **Define input struct** in the appropriate handler file with JSON schema tags:
   ```go
   type MyToolInput struct {
       Field string `json:"field" jsonschema:"required,description=Field description"`
   }
   ```

2. **Implement handler method** that returns `(*mcp.CallToolResult, *SomeType, error)`:
   ```go
   func (h *Handlers) HandleMyTool(ctx context.Context, req *mcp.CallToolRequest, input MyToolInput) (*mcp.CallToolResult, *SomeType, error) {
       // Validate input
       // Call client method
       // Return result
   }
   ```

3. **Register in main.go**:
   ```go
   mcp.AddTool(server, &mcp.Tool{
       Name:        "my_tool",
       Description: "Description of what the tool does",
   }, handlers.HandleMyTool)
   ```

### Adding a New Client Method

1. **Define request/response types** in `internal/client/types.go`:
   ```go
   type MyRequest struct {
       Field string `json:"field"`
   }

   type MyResponse struct {
       Result string `json:"result"`
   }
   ```

2. **Implement client method** in `internal/client/extend.go`:
   ```go
   func (c *Client) MyMethod(ctx context.Context, req *MyRequest) (*MyResponse, error) {
       resp, err := c.doRequest(ctx, http.MethodPost, "/my-endpoint", req)
       if err != nil {
           return nil, err
       }
       defer resp.Body.Close()

       var result MyResponse
       if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
           return nil, fmt.Errorf("failed to decode response: %w", err)
       }
       return &result, nil
   }
   ```

### Error Handling

- Client methods return errors with context: `fmt.Errorf("failed to X: %w", err)`
- HTTP errors (4xx/5xx) are decoded from Extend AI's error response format
- Handlers return MCP error results with `IsError: true` for validation errors

## Key Technical Decisions

### Why stdio Transport?
MCP clients like Claude Desktop communicate with servers via standard input/output. This is simpler than HTTP for local integrations and enables bidirectional communication.

### Why No Middleware/Framework?
The MCP Go SDK handles request routing, JSON schema validation, and response formatting. Custom HTTP middleware isn't needed.

### Why Separate Handler Types?
Grouping related tools into `ProcessorHandlers`, `WorkflowHandlers`, and `ParseHandlers` improves code organization and makes it easy to find tool implementations.

### API Client Timeout
The HTTP client has a 30-second timeout (see `client.NewClient`). This may need adjustment for long-running workflows or large file uploads.

## Important Files

- `cmd/server/main.go` - Server entry point, tool registration
- `internal/client/extend.go` - All Extend AI API calls
- `internal/client/types.go` - API request/response types
- `internal/handlers/*.go` - MCP tool handlers
- `internal/config/config.go` - Environment variable configuration
- `.env.example` - Template for local environment setup
- `go.mod` - Go module dependencies

## External API Reference

The server wraps the Extend AI API. For details on processor types, workflow configurations, and response formats, see:
- [Extend AI Developer Docs](https://docs.extend.ai/developers/)
- [Extend AI API Reference](https://docs.extend.ai/api-reference/)
