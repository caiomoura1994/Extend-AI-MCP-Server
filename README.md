# Extend AI MCP Server

> **Disclaimer:** This is an **unofficial**, community-maintained MCP server for [Extend AI](https://extend.ai/). It is **not** affiliated with, endorsed by, or officially supported by Extend AI. This project was built independently to integrate Extend AI's document processing capabilities with MCP-compatible clients. Use at your own discretion.

A read-only [Model Context Protocol](https://modelcontextprotocol.io/) (MCP) server that exposes [Extend AI](https://extend.ai/) document processing data as tools for AI assistants like Claude Desktop, Cursor, and other MCP clients.

## What It Does

This server gives your AI assistant read-only access to your Extend AI account data:

- **Processors** - List processors and inspect processor run results
- **Workflows** - List workflows and inspect workflow run results
- **Extractors** - List extractors, versions, and inspect extraction runs
- **Parse** - Inspect parse run results
- **Files** - List uploaded files

All operations are **read-only** -- this server cannot create, modify, or delete any resources in your Extend AI account.

## Available Tools

| Category | Tool | Description |
|---|---|---|
| **Processors** | `list_processors` | List all available processors |
| | `get_processor_run` | Get the status and result of a processor execution |
| **Workflows** | `list_workflows` | List all available workflows |
| | `list_workflow_runs` | List all workflow execution runs with status and metadata |
| | `get_workflow_run` | Get the status and result of a workflow execution |
| **Extractors** | `list_extractors` | List all available extractors |
| | `list_extract_runs` | List all extractor execution runs |
| | `get_extractor` | Get details of a specific extractor |
| | `list_extractor_versions` | List all versions of a specific extractor |
| | `get_extractor_version` | Get a specific version of an extractor |
| **Parse** | `get_parse_run` | Get the result of a parse operation |
| **Files** | `list_files` | List all files uploaded to the account |

## Quick Start

### Prerequisites

- [Go 1.23+](https://go.dev/dl/)
- An Extend AI API key ([get one here](https://extend.ai))
- An MCP-compatible client (Claude Desktop, Cursor, etc.)

### Build

```bash
git clone https://github.com/caiomoura/extend-mcp-server.git
cd extend-mcp-server
go build -o extend-mcp-server cmd/server/main.go
```

### Configure Claude Desktop

Edit your Claude Desktop config:

- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

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

Restart Claude Desktop after saving.

### Configure Cursor

Add to your Cursor MCP settings (`.cursor/mcp.json`):

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

## Configuration

All configuration is done via environment variables:

| Variable | Required | Default | Description |
|---|---|---|---|
| `EXTEND_API_KEY` | Yes | - | Your Extend AI API key |
| `EXTEND_API_VERSION` | No | `2026-02-09` | Extend AI API version |
| `EXTEND_BASE_URL` | No | `https://api.extend.ai` | Extend AI API base URL |

### Regional Endpoints

If your Extend AI account is on a regional instance (e.g. US2), set the base URL accordingly:

```bash
EXTEND_BASE_URL=https://api.us2.extend.app
```

## Architecture

The project follows a **Clean Architecture** (Layered Architecture) pattern with clear separation of concerns:

```
extend-mcp-server/
├── cmd/
│   └── server/
│       └── main.go                 # Entry point: loads config, creates client, starts server
├── internal/
│   ├── dto/                        # Data Transfer Objects (pure data types, no logic)
│   │   ├── processors.go           # Processor, ProcessorRun, response wrappers
│   │   ├── workflows.go            # Workflow, WorkflowRun, WorkflowRunSummary
│   │   ├── extractors.go           # Extractor, ExtractorVersion, ExtractRun
│   │   ├── parse.go                # ParserRun, ParserRunChunk, metrics
│   │   └── files.go                # File, FileMetadata
│   ├── repositories/               # Data access layer (HTTP client + API methods)
│   │   ├── client.go               # HTTP client: auth, headers, error handling
│   │   ├── processors.go           # GetProcessorRun, ListProcessors
│   │   ├── workflows.go            # GetWorkflowRun, ListWorkflowRuns, ListWorkflows
│   │   ├── extractors.go           # ListExtractRuns, GetExtractor, ListExtractors, etc.
│   │   ├── parse.go                # GetParseRun
│   │   └── files.go                # ListFiles
│   ├── handlers/                   # MCP tool handlers (orchestrate repos + format responses)
│   │   ├── processors.go           # Processor tool handlers
│   │   ├── workflows.go            # Workflow tool handlers
│   │   ├── extractors.go           # Extractor tool handlers
│   │   ├── parse.go                # Parse tool handlers
│   │   ├── files.go                # File tool handlers
│   │   └── handlers_test.go        # Integration tests
│   ├── registry/
│   │   └── registry.go             # Tool registration and server startup logging
│   └── config/
│       └── config.go               # Environment variable configuration
├── go.mod
├── go.sum
├── .env.example
├── CLAUDE.md
└── README.md
```

### Layer Dependencies

```
handlers  -->  repositories  -->  dto
   |                |
   +--- dto --------+
```

- **dto** - No dependencies. Pure structs that represent Extend AI API shapes.
- **repositories** - Depends on `dto`. Contains the HTTP client and all API call methods.
- **handlers** - Depends on `dto` and `repositories`. Validates input, calls repositories, formats MCP responses.
- **registry** - Wires handlers to the MCP server based on configuration.

## Development

### Running Locally

```bash
export EXTEND_API_KEY=your_api_key_here
go run cmd/server/main.go
```

The server communicates via stdio (standard input/output), which is the MCP transport protocol.

### Testing with MCP Inspector

```bash
# Build the binary
go build -o extend-mcp-server cmd/server/main.go

# Run the MCP Inspector UI
npx @modelcontextprotocol/inspector ./extend-mcp-server
```

This opens a web UI where you can list tools, call them, and inspect responses.

### Running Tests

```bash
go test -v ./...
```

Tests use in-memory MCP transports to validate handler behavior without hitting the real API.

## Contributing

Contributions are welcome! Here's how you can help:

### Getting Started

1. **Fork** the repository
2. **Clone** your fork:
   ```bash
   git clone https://github.com/your-username/extend-mcp-server.git
   cd extend-mcp-server
   ```
3. **Create a branch** for your feature:
   ```bash
   git checkout -b feat/my-feature
   ```
4. **Make your changes** and ensure they build:
   ```bash
   go build ./...
   go test ./...
   ```
5. **Commit** with a conventional commit message:
   ```bash
   git commit -m "feat: add support for document uploads"
   ```
6. **Push** and open a **Pull Request**

### Commit Convention

We use [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` - New feature (e.g. new tool, new endpoint)
- `fix:` - Bug fix
- `refactor:` - Code restructuring without behavior change
- `docs:` - Documentation changes
- `test:` - Adding or updating tests
- `chore:` - Maintenance tasks (deps, CI, etc.)

### Adding a New Tool

1. **Define the DTO** in `internal/dto/<domain>.go` if new types are needed:
   ```go
   type MyResource struct {
       ID   string `json:"id"`
       Name string `json:"name"`
   }
   ```

2. **Add the repository method** in `internal/repositories/<domain>.go`:
   ```go
   func (c *Client) GetMyResource(ctx context.Context, id string) (*dto.MyResource, error) {
       path := fmt.Sprintf("/my_resources/%s", id)
       resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
       if err != nil {
           return nil, err
       }
       defer resp.Body.Close()

       var result dto.MyResource
       if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
           return nil, fmt.Errorf("failed to decode response: %w", err)
       }
       return &result, nil
   }
   ```

3. **Create the handler** in `internal/handlers/<domain>.go`:
   ```go
   func (h *MyHandlers) HandleGetMyResource(ctx context.Context, req *mcp.CallToolRequest, input GetMyResourceInput) (*mcp.CallToolResult, *dto.MyResource, error) {
       if input.ID == "" {
           return nil, nil, fmt.Errorf("id is required")
       }
       result, err := h.client.GetMyResource(ctx, input.ID)
       if err != nil {
           return nil, nil, fmt.Errorf("failed to get resource: %w", err)
       }
       return nil, result, nil
   }
   ```

4. **Register the tool** in `internal/registry/registry.go`:
   ```go
   mcp.AddTool(server, &mcp.Tool{
       Name:        "get_my_resource",
       Description: "Get details of a specific resource",
   }, myHandlers.HandleGetMyResource)
   ```

5. **Add tests** in `internal/handlers/handlers_test.go`

### Important Guidelines

- **Read-only only**: This server intentionally exposes only read operations (GET/LIST). Do not add tools that create, modify, or delete resources.
- **Error handling**: Always return `(nil, nil, fmt.Errorf(...))` for errors in handlers, never `(*mcp.CallToolResult{IsError: true}, nil, nil)`. This avoids an MCP SDK output validation bug with nil slices.
- **Non-nil slices**: For list handlers, always ensure slices are non-nil before returning (use `if result == nil { result = []dto.Type{} }`). The MCP SDK rejects JSON `null` where an array is expected.
- **Tests**: Run `go test ./...` before submitting. Add tests for new tools.

### Ideas for Contributions

- Add pagination support for list endpoints
- Add filtering/search parameters to list tools
- Improve error messages with more context
- Add more detailed tool descriptions for better AI understanding
- Add support for new Extend AI API endpoints as they are released

## Troubleshooting

| Problem | Solution |
|---|---|
| Server won't start | Check that `EXTEND_API_KEY` is set |
| Tools not appearing | Restart your MCP client after config changes |
| 404 API errors | Check `EXTEND_BASE_URL` matches your regional instance |
| 401 API errors | Verify your API key is valid and active |
| Endpoint not found | Check `EXTEND_API_VERSION` matches available endpoints |

## Links

- [Model Context Protocol](https://modelcontextprotocol.io/)
- [MCP Go SDK](https://github.com/modelcontextprotocol/go-sdk)
- [Extend AI Docs](https://docs.extend.ai/developers/)
- [Extend AI API Reference](https://docs.extend.ai/api-reference/)

## License

MIT License - See [LICENSE](LICENSE) file for details.
