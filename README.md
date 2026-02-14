# Extend AI MCP Server

> **Note:** This is an unofficial, community-maintained MCP server for Extend AI. It is not affiliated with or endorsed by Extend AI. Use at your own discretion.

A Model Context Protocol (MCP) server for integrating with the Extend AI API, enabling AI-powered document processing through Claude Desktop and other MCP clients.

## Overview

Extend AI is a platform for processing documents with AI, offering:
- **Document Extraction**: Extract structured data from documents
- **Document Splitting**: Split multi-page documents
- **Document Classification**: Classify document types
- **Document Parsing**: Convert documents to structured text
- **Workflows**: Orchestrate multiple processors

This MCP server exposes Extend AI functionality as tools that can be used by AI assistants like Claude.

## Features

### Processor Tools
- `run_processor` - Execute a processor to process documents
- `get_processor_run` - Get status and results of a processor execution
- `list_processors` - List all available processors
- `create_processor` - Create a new processor
- `delete_processor_run` - Delete a processor execution

### Workflow Tools
- `run_workflow` - Execute a workflow
- `get_workflow_run` - Get status and results of a workflow execution
- `list_workflows` - List all available workflows
- `batch_run_workflow` - Execute a workflow on multiple files
- `create_workflow` - Create a new workflow
- `delete_workflow_run` - Delete a workflow execution

### Parse Tools
- `parse_file` - Parse a file synchronously
- `parse_file_async` - Parse a file asynchronously
- `get_parse_run` - Get results of an async parse operation
- `delete_parse_run` - Delete a parse execution

### Extractor Tools (Read-Only)
- `list_extract_runs` - List all extractor execution runs
- `get_extractor` - Get details of a specific extractor
- `list_extractors` - List all available extractors
- `get_extractor_version` - Get a specific version of an extractor
- `list_extractor_versions` - List all versions of a specific extractor

### Unimplemented Tools (TODO)
The following endpoints are documented in the Extend AI API but not yet implemented in this MCP server:
- `extract_file` - POST /extractors/run - Execute an extractor on a document

## Installation

### Prerequisites
- Go 1.23 or higher
- Extend AI API key ([Get one here](https://extend.ai))
- Claude Desktop or another MCP client

### Build from Source

1. Clone the repository:
```bash
git clone https://github.com/caiomoura/extend-mcp-server.git
cd extend-mcp-server
```

2. Install dependencies:
```bash
go mod download
```

3. Build the server:
```bash
go build -o extend-mcp-server cmd/server/main.go
```

4. (Optional) Install globally:
```bash
go install ./cmd/server
```

## Configuration

### Environment Variables

Create a `.env` file or set the following environment variables:

```bash
# Required
EXTEND_API_KEY=your_api_key_here

# Optional - API Config
EXTEND_API_VERSION=2024-11-14           # Default: 2024-11-14
EXTEND_BASE_URL=https://api.extend.ai   # Default: https://api.extend.ai

# Optional - Tool Enablement (default: all true)
EXTEND_ENABLE_PROCESSORS=true          # Enable processor tools
EXTEND_ENABLE_WORKFLOWS=true           # Enable workflow tools
EXTEND_ENABLE_PARSE=true                # Enable parse tools
EXTEND_ENABLE_EXTRACTORS=true          # Enable extractor tools

# Optional - Operation Control (default: false)
EXTEND_READ_ONLY_MODE=false             # When true, only GET/LIST operations are enabled
```

### Configuration Examples

**Example 1: Read-Only Mode (Security-Focused)**
```bash
EXTEND_API_KEY=your_api_key_here
EXTEND_READ_ONLY_MODE=true              # Only allow read operations (GET, LIST)
```
This configuration blocks all POST (run_*, create_*) and DELETE operations, allowing only safe read operations like `list_processors`, `get_workflow_run`, etc.

**Example 2: Extractors Only**
```bash
EXTEND_API_KEY=your_api_key_here
EXTEND_ENABLE_PROCESSORS=false
EXTEND_ENABLE_WORKFLOWS=false
EXTEND_ENABLE_PARSE=false
EXTEND_ENABLE_EXTRACTORS=true           # Only enable extractor tools
```
This configuration registers only the 5 extractor tools, useful when you only need metadata and version management.

**Example 3: Processors and Workflows Only**
```bash
EXTEND_API_KEY=your_api_key_here
EXTEND_ENABLE_PROCESSORS=true
EXTEND_ENABLE_WORKFLOWS=true
EXTEND_ENABLE_PARSE=false               # Disable parse tools
EXTEND_ENABLE_EXTRACTORS=false          # Disable extractor tools
```
This configuration focuses on document processing operations, excluding parsing and extractor management.

### Claude Desktop Setup

Edit your Claude Desktop configuration file:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

Add the following configuration:

```json
{
  "mcpServers": {
    "extend-ai": {
      "command": "/path/to/extend-mcp-server",
      "env": {
        "EXTEND_API_KEY": "your_api_key_here",
        "EXTEND_API_VERSION": "2024-11-14"
      }
    }
  }
}
```

Replace `/path/to/extend-mcp-server` with the actual path to the built binary.

After updating the configuration, restart Claude Desktop.

## Usage Examples

### Example 1: Parse a PDF Document

```
User: Parse this PDF and extract the text: https://example.com/document.pdf

Claude will use: parse_file tool
Input: {
  "file_url": "https://example.com/document.pdf",
  "format": "text"
}
```

### Example 2: Run a Document Extraction Workflow

```
User: Extract invoice data from this document using workflow wf_12345

Claude will use: run_workflow tool
Input: {
  "workflow_id": "wf_12345",
  "file_url": "https://example.com/invoice.pdf"
}

Then: get_workflow_run tool to check status
Input: {
  "run_id": "run_67890"
}
```

### Example 3: List Available Processors

```
User: What processors are available in Extend AI?

Claude will use: list_processors tool
Returns: List of all processors with their configurations
```

### Example 4: Batch Process Documents

```
User: Process these 5 invoices using the invoice extraction workflow

Claude will use: batch_run_workflow tool
Input: {
  "workflow_id": "wf_invoice",
  "files": [
    {"file_url": "https://example.com/invoice1.pdf"},
    {"file_url": "https://example.com/invoice2.pdf"},
    ...
  ]
}
```

## Tool Reference

### Processor Tools

#### `run_processor`
Execute a processor on a document.

**Input:**
```json
{
  "processor_id": "proc_12345",        // Required: Processor ID
  "file_url": "https://...",           // Either file_url or file_content required
  "file_content": "base64...",         // Base64-encoded file
  "parameters": {"key": "value"}       // Optional: Additional parameters
}
```

**Output:**
```json
{
  "id": "run_67890",
  "processor_id": "proc_12345",
  "status": "pending",
  "created_at": "2024-11-14T10:00:00Z"
}
```

#### `get_processor_run`
Get the status and result of a processor execution.

**Input:**
```json
{
  "run_id": "run_67890"
}
```

**Output:**
```json
{
  "id": "run_67890",
  "status": "completed",
  "progress": 100,
  "result": {
    "extracted_data": {...}
  }
}
```

### Workflow Tools

#### `run_workflow`
Execute a workflow on a document.

**Input:**
```json
{
  "workflow_id": "wf_12345",
  "file_url": "https://...",
  "parameters": {"key": "value"}
}
```

**Output:**
```json
{
  "id": "run_67890",
  "workflow_id": "wf_12345",
  "status": "running",
  "steps": [...]
}
```

#### `batch_run_workflow`
Execute a workflow on multiple files.

**Input:**
```json
{
  "workflow_id": "wf_12345",
  "files": [
    {"file_url": "https://example.com/doc1.pdf"},
    {"file_url": "https://example.com/doc2.pdf"}
  ],
  "parameters": {"key": "value"}
}
```

**Output:**
```json
[
  {"id": "run_1", "status": "pending"},
  {"id": "run_2", "status": "pending"}
]
```

### Parse Tools

#### `parse_file`
Parse a file synchronously.

**Input:**
```json
{
  "file_url": "https://...",
  "format": "text"              // Optional: text, markdown, json
}
```

**Output:**
```json
{
  "text": "Extracted text content...",
  "format": "text",
  "pages": 5
}
```

#### `parse_file_async`
Parse a file asynchronously (for large files).

**Input:**
```json
{
  "file_url": "https://...",
  "format": "markdown"
}
```

**Output:**
```json
{
  "id": "parse_12345",
  "status": "pending"
}
```

### Extractor Tools

#### `list_extract_runs`
List all extractor execution runs.

**Input:**
```json
{}
```

**Output:**
```json
{
  "extract_runs": [
    {
      "id": "extract_run_12345",
      "extractor_id": "ext_67890",
      "status": "completed",
      "progress": 1.0,
      "result": {"field": "value"},
      "created_at": "2024-11-14T10:00:00Z",
      "updated_at": "2024-11-14T10:05:00Z"
    }
  ]
}
```

#### `get_extractor`
Get details of a specific extractor.

**Input:**
```json
{
  "extractor_id": "ext_67890"
}
```

**Output:**
```json
{
  "id": "ext_67890",
  "name": "Invoice Extractor",
  "description": "Extract structured data from invoices",
  "type": "structured",
  "schema": {"fields": ["invoice_number", "total", "date"]},
  "created_at": "2024-11-14T10:00:00Z",
  "updated_at": "2024-11-14T10:00:00Z"
}
```

#### `list_extractors`
List all available extractors.

**Input:**
```json
{}
```

**Output:**
```json
{
  "extractors": [
    {
      "id": "ext_67890",
      "name": "Invoice Extractor",
      "description": "Extract structured data from invoices",
      "type": "structured",
      "schema": {"fields": ["invoice_number", "total", "date"]},
      "created_at": "2024-11-14T10:00:00Z",
      "updated_at": "2024-11-14T10:00:00Z"
    }
  ]
}
```

#### `get_extractor_version`
Get a specific version of an extractor.

**Input:**
```json
{
  "extractor_id": "ext_67890",
  "version_id": "v_123"
}
```

**Output:**
```json
{
  "id": "v_123",
  "extractor_id": "ext_67890",
  "version": "1.0.0",
  "schema": {"fields": ["invoice_number", "total"]},
  "status": "published",
  "created_at": "2024-11-14T10:00:00Z",
  "updated_at": "2024-11-14T10:00:00Z"
}
```

#### `list_extractor_versions`
List all versions of a specific extractor.

**Input:**
```json
{
  "extractor_id": "ext_67890"
}
```

**Output:**
```json
{
  "versions": [
    {
      "id": "v_123",
      "extractor_id": "ext_67890",
      "version": "1.0.0",
      "schema": {"fields": ["invoice_number", "total"]},
      "status": "published",
      "created_at": "2024-11-14T10:00:00Z",
      "updated_at": "2024-11-14T10:00:00Z"
    }
  ]
}
```

## Development

### Project Structure

```
extend-mcp-server/
├── cmd/
│   └── server/
│       └── main.go           # Main server entry point (minimal, delegates to registry)
├── internal/
│   ├── client/
│   │   ├── extend.go         # Extend AI HTTP client
│   │   └── types.go          # Request/response types
│   ├── handlers/
│   │   ├── processors.go     # Processor tool handlers
│   │   ├── workflows.go      # Workflow tool handlers
│   │   ├── parse.go          # Parse tool handlers
│   │   └── extractors.go     # Extractor tool handlers
│   ├── registry/
│   │   └── registry.go       # Tool registration logic and configuration
│   └── config/
│       └── config.go         # Configuration management
├── go.mod
├── go.sum
├── README.md
└── .env.example
```

### Running in Development

```bash
# Set environment variables
export EXTEND_API_KEY=your_api_key

# Run the server
go run cmd/server/main.go
```

### Testing

The server communicates via stdio, so you can test it manually:

```bash
# Build the server
go build -o extend-mcp-server cmd/server/main.go

# Test with a simple request (optional, requires MCP client)
./extend-mcp-server
```

## Troubleshooting

### Server Not Starting
- Check that `EXTEND_API_KEY` is set
- Verify the binary has execute permissions
- Check Claude Desktop logs for errors

### Tools Not Appearing in Claude
- Ensure `claude_desktop_config.json` is valid JSON
- Restart Claude Desktop after config changes
- Check the command path is correct

### API Errors
- Verify your API key is valid
- Check API version compatibility
- Review Extend AI documentation for endpoint changes

## API Documentation

For full Extend AI API documentation, visit:
- [Extend AI Docs](https://docs.extend.ai/developers/)
- [API Reference](https://docs.extend.ai/api-reference/)

## Contributing

Contributions are welcome! Please:
1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

MIT License - See LICENSE file for details

## Support

For issues or questions:
- GitHub Issues: [github.com/caiomoura/extend-mcp-server/issues](https://github.com/caiomoura/extend-mcp-server/issues)
- Extend AI Support: [extend.ai/support](https://extend.ai/support)

## Credits

Built with:
- [Model Context Protocol](https://modelcontextprotocol.io/)
- [MCP Go SDK](https://github.com/modelcontextprotocol/go-sdk)
- [Extend AI](https://extend.ai/)
