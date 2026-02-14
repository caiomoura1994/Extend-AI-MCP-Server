# Contributing

Thanks for your interest in contributing to the Extend AI MCP Server!

## Development Setup

### Prerequisites

- [Go 1.23+](https://go.dev/dl/)
- An Extend AI API key for testing

### Getting Started

```bash
git clone https://github.com/caiomoura/extend-mcp-server.git
cd extend-mcp-server
go mod download
```

### Running Locally

```bash
export EXTEND_API_KEY=your_api_key_here
go run ./cmd/extend-mcp-server/
```

The server communicates via stdio (standard input/output), which is the MCP transport protocol.

### Testing with MCP Inspector

```bash
go build -o extend-mcp-server ./cmd/extend-mcp-server/
npx @modelcontextprotocol/inspector ./extend-mcp-server
```

This opens a web UI where you can list tools, call them, and inspect responses.

### Running Tests

```bash
go test -v ./...
```

Tests use in-memory MCP transports to validate handler behavior without hitting the real API.

## Architecture

The project follows a **Clean Architecture** pattern with clear separation of concerns:

```
extend-mcp-server/
├── cmd/
│   └── extend-mcp-server/
│       └── main.go                 # Entry point: loads config, creates client, starts server
├── internal/
│   ├── dto/                        # Data Transfer Objects (pure data types, no logic)
│   │   ├── processors.go
│   │   ├── workflows.go
│   │   ├── extractors.go
│   │   ├── parse.go
│   │   └── files.go
│   ├── repositories/               # Data access layer (HTTP client + API methods)
│   │   ├── client.go               # HTTP client: auth, headers, error handling
│   │   ├── processors.go
│   │   ├── workflows.go
│   │   ├── extractors.go
│   │   ├── parse.go
│   │   └── files.go
│   ├── handlers/                   # MCP tool handlers (validate input, call repos, format response)
│   │   ├── processors.go
│   │   ├── workflows.go
│   │   ├── extractors.go
│   │   ├── parse.go
│   │   ├── files.go
│   │   └── handlers_test.go
│   ├── registry/
│   │   └── registry.go             # Tool registration and server startup logging
│   └── config/
│       └── config.go               # Environment variable configuration
├── go.mod
├── go.sum
├── .env.example
├── CLAUDE.md
├── CONTRIBUTING.md
└── README.md
```

### Layer Dependencies

```
handlers  -->  repositories  -->  dto
   |                |
   +--- dto --------+
```

- **dto** - No dependencies. Pure structs that represent Extend AI API request/response shapes.
- **repositories** - Depends on `dto`. HTTP client and all API call methods.
- **handlers** - Depends on `dto` and `repositories`. Validates input, calls repositories, formats MCP responses.
- **registry** - Wires handlers to the MCP server.

## How to Contribute

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
4. **Make your changes** and ensure they build and pass tests:
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

## Adding a New Tool

### 1. Define the DTO

Add types in `internal/dto/<domain>.go` if new types are needed:

```go
type MyResource struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

type ListMyResourcesResponse struct {
    Success   bool         `json:"success"`
    Resources []MyResource `json:"resources"`
}
```

### 2. Add the repository method

Add the API call in `internal/repositories/<domain>.go`:

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

### 3. Create the handler

Add the MCP handler in `internal/handlers/<domain>.go`:

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

### 4. Register the tool

Add the tool in `internal/registry/registry.go`:

```go
mcp.AddTool(server, &mcp.Tool{
    Name:        "get_my_resource",
    Description: "Get details of a specific resource",
}, myHandlers.HandleGetMyResource)
```

Don't forget to update the tool count in the `return` statement.

### 5. Add tests

Add test cases in `internal/handlers/handlers_test.go`.

## Important Guidelines

- **Read-only only**: This server intentionally exposes only read operations (GET/LIST). Do not add tools that create, modify, or delete resources.
- **Error handling**: Always return `(nil, nil, fmt.Errorf(...))` for errors in handlers, never `(*mcp.CallToolResult{IsError: true}, nil, nil)`. This avoids an MCP SDK output validation bug with nil slices.
- **Non-nil slices**: For list handlers, always ensure slices are non-nil before returning (use `if result == nil { result = []dto.Type{} }`). The MCP SDK rejects JSON `null` where an array is expected.
- **Tests**: Run `go test ./...` before submitting. Add tests for new tools.

## Ideas for Contributions

- Add pagination support for list endpoints
- Add filtering/search parameters to list tools
- Improve error messages with more context
- Add more detailed tool descriptions for better AI understanding
- Add support for new Extend AI API endpoints as they are released
- Pre-built binaries via GitHub Releases (goreleaser)
