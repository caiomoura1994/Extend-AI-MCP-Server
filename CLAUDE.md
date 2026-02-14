# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is an unofficial, community-maintained Model Context Protocol (MCP) server that integrates with the Extend AI API. It exposes Extend AI's document processing capabilities as read-only MCP tools for AI assistants like Claude Desktop and Cursor.

## Tech Stack

- **Language**: Go 1.23+
- **Framework**: MCP Go SDK (`github.com/modelcontextprotocol/go-sdk`)
- **Transport**: stdio (standard input/output)
- **API**: Extend AI REST API

## Architecture

The codebase follows Clean Architecture with four layers:

1. **Entry Point** (`cmd/extend-mcp-server/main.go`)
   - Loads configuration from environment variables
   - Creates Extend AI HTTP client
   - Initializes MCP server and registers all 12 tools
   - Starts stdio transport

2. **DTO Layer** (`internal/dto/`)
   - Pure data types with no dependencies
   - `processors.go`, `workflows.go`, `extractors.go`, `parse.go`, `files.go`
   - Represents Extend AI API request/response shapes

3. **Repository Layer** (`internal/repositories/`)
   - `client.go` - HTTP client with auth, headers, error handling
   - `processors.go`, `workflows.go`, `extractors.go`, `parse.go`, `files.go`
   - Each file contains API methods for its domain

4. **Handler Layer** (`internal/handlers/`)
   - `processors.go`, `workflows.go`, `extractors.go`, `parse.go`, `files.go`
   - Each handler validates input, calls a repository method, and formats MCP responses
   - `handlers_test.go` - Integration tests using in-memory MCP transports

5. **Registry** (`internal/registry/registry.go`)
   - Wires all handlers to the MCP server
   - Logs server startup information

### Tool Categories (12 tools, all read-only)

**Processors** (2): `list_processors`, `get_processor_run`
**Workflows** (3): `list_workflows`, `list_workflow_runs`, `get_workflow_run`
**Extractors** (5): `list_extractors`, `list_extract_runs`, `get_extractor`, `list_extractor_versions`, `get_extractor_version`
**Parse** (1): `get_parse_run`
**Files** (1): `list_files`

## Development Commands

### Build

```bash
go mod download
go build -o extend-mcp-server ./cmd/extend-mcp-server/
```

### Run Locally

```bash
export EXTEND_API_KEY=your_api_key_here
go run ./cmd/extend-mcp-server/
```

### Testing

```bash
go test -v ./...
```

Tests use in-memory MCP transports (`mcp.NewInMemoryTransports()`) to validate handler behavior without hitting the real API.

### MCP Inspector

```bash
go build -o extend-mcp-server ./cmd/extend-mcp-server/
npx @modelcontextprotocol/inspector ./extend-mcp-server
```

## Configuration

Environment variables (see `internal/config/config.go`):

- **EXTEND_API_KEY** (required): Your Extend AI API key
- **EXTEND_API_VERSION** (optional): API version, defaults to "2026-02-09"
- **EXTEND_BASE_URL** (optional): Base URL, defaults to "https://api.extend.ai"

## Code Patterns

### Adding a New Tool

1. **Define DTO** in `internal/dto/<domain>.go`
2. **Add repository method** in `internal/repositories/<domain>.go` (imports `dto`)
3. **Create handler** in `internal/handlers/<domain>.go` (imports `dto` + `repositories`)
4. **Register tool** in `internal/registry/registry.go`
5. **Add tests** in `internal/handlers/handlers_test.go`

### Error Handling

- Handlers return `(nil, nil, fmt.Errorf(...))` for errors -- never `(*mcp.CallToolResult{IsError: true}, nil, nil)`. This avoids an MCP SDK output validation bug with nil slices.
- For list handlers, ensure slices are non-nil: `if result == nil { result = []dto.Type{} }`
- Repository methods return errors with context: `fmt.Errorf("failed to X: %w", err)`

### Key Technical Decisions

- **Read-only only**: No create/update/delete operations. The MCP client decides which tools to use.
- **stdio transport**: MCP clients communicate via stdin/stdout, simpler than HTTP for local integrations.
- **No enable flags**: All 12 tools are always registered. The MCP client selects which tools to invoke.
- **30s HTTP timeout**: Set in `repositories.NewClient`. May need adjustment for large responses.

## Important Files

- `cmd/extend-mcp-server/main.go` - Server entry point
- `internal/dto/*.go` - API data types
- `internal/repositories/client.go` - HTTP client core
- `internal/repositories/*.go` - API call methods
- `internal/handlers/*.go` - MCP tool handlers
- `internal/handlers/handlers_test.go` - Integration tests
- `internal/registry/registry.go` - Tool registration
- `internal/config/config.go` - Environment variable configuration
- `.env.example` - Template for local environment setup

## External API Reference

- [Extend AI Developer Docs](https://docs.extend.ai/developers/)
- [Extend AI API Reference](https://docs.extend.ai/api-reference/)
