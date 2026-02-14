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

You only need **Node.js 18+** (which you probably already have) and an [Extend AI API key](https://extend.ai).

### Claude Desktop

Edit your config file:

- **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "extend-ai": {
      "command": "npx",
      "args": ["-y", "extend-ai-mcp-server"],
      "env": {
        "EXTEND_API_KEY": "your_api_key_here"
      }
    }
  }
}
```

Restart Claude Desktop after saving.

### Cursor

Add to your project's `.cursor/mcp.json` or global Cursor MCP settings:

```json
{
  "mcpServers": {
    "extend-ai": {
      "command": "npx",
      "args": ["-y", "extend-ai-mcp-server"],
      "env": {
        "EXTEND_API_KEY": "your_api_key_here"
      }
    }
  }
}
```

### Other MCP Clients

Any MCP client that supports stdio transport can use this server:

```json
{
  "command": "npx",
  "args": ["-y", "extend-ai-mcp-server"],
  "env": { "EXTEND_API_KEY": "your_api_key_here" }
}
```

> **How it works:** `npx` automatically downloads a thin wrapper that fetches the correct pre-built binary for your OS on first run. The binary is cached at `~/.extend-mcp-server/`, so subsequent starts are instant.

## Alternative Installation Methods

<details>
<summary><strong>Using <code>go run</code> (requires Go 1.23+)</strong></summary>

If you have Go installed, you can skip npm entirely:

```json
{
  "mcpServers": {
    "extend-ai": {
      "command": "go",
      "args": ["run", "github.com/caiomoura/extend-mcp-server/cmd/extend-mcp-server@latest"],
      "env": {
        "EXTEND_API_KEY": "your_api_key_here"
      }
    }
  }
}
```

Go downloads and compiles the server automatically. The build is cached, so subsequent starts are fast.

</details>

<details>
<summary><strong>Pre-built binary (no Node.js or Go required)</strong></summary>

Download the latest binary for your platform from [GitHub Releases](https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases):

**macOS (Apple Silicon):**

```bash
curl -Lo extend-mcp-server.tar.gz https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases/latest/download/extend-mcp-server_darwin_arm64.tar.gz
tar xzf extend-mcp-server.tar.gz
sudo mv extend-mcp-server /usr/local/bin/
rm extend-mcp-server.tar.gz
```

**macOS (Intel):**

```bash
curl -Lo extend-mcp-server.tar.gz https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases/latest/download/extend-mcp-server_darwin_amd64.tar.gz
tar xzf extend-mcp-server.tar.gz
sudo mv extend-mcp-server /usr/local/bin/
rm extend-mcp-server.tar.gz
```

**Linux (x86_64):**

```bash
curl -Lo extend-mcp-server.tar.gz https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases/latest/download/extend-mcp-server_linux_amd64.tar.gz
tar xzf extend-mcp-server.tar.gz
sudo mv extend-mcp-server /usr/local/bin/
rm extend-mcp-server.tar.gz
```

**Linux (ARM64):**

```bash
curl -Lo extend-mcp-server.tar.gz https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases/latest/download/extend-mcp-server_linux_arm64.tar.gz
tar xzf extend-mcp-server.tar.gz
sudo mv extend-mcp-server /usr/local/bin/
rm extend-mcp-server.tar.gz
```

**Windows (x86_64):**

1. Download [extend-mcp-server_windows_amd64.zip](https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases/latest/download/extend-mcp-server_windows_amd64.zip)
2. Extract `extend-mcp-server.exe`
3. Move it to a directory in your PATH (e.g. `C:\Program Files\extend-mcp-server\`)

Then use the binary directly in your MCP config:

```json
{
  "mcpServers": {
    "extend-ai": {
      "command": "/usr/local/bin/extend-mcp-server",
      "env": {
        "EXTEND_API_KEY": "your_api_key_here"
      }
    }
  }
}
```

> On Windows use `"C:\\Program Files\\extend-mcp-server\\extend-mcp-server.exe"` as the command.

</details>

<details>
<summary><strong>Install with <code>go install</code></strong></summary>

```bash
go install github.com/caiomoura/extend-mcp-server/cmd/extend-mcp-server@latest
```

The binary is placed in `$(go env GOPATH)/bin/`. Then use `"command": "extend-mcp-server"` in your MCP config (if the Go bin directory is in your PATH).

</details>

<details>
<summary><strong>Build from source</strong></summary>

```bash
git clone https://github.com/caiomoura1994/Extend-AI-MCP-Server.git
cd Extend-AI-MCP-Server
go build -o extend-mcp-server ./cmd/extend-mcp-server/
```

</details>

## Configuration

All configuration is done via environment variables passed through your MCP client config:

| Variable | Required | Default | Description |
|---|---|---|---|
| `EXTEND_API_KEY` | Yes | - | Your Extend AI API key |
| `EXTEND_API_VERSION` | No | `2026-02-09` | Extend AI API version |
| `EXTEND_BASE_URL` | No | `https://api.extend.ai` | Extend AI API base URL |

### Regional Endpoints

If your Extend AI account is on a regional instance (e.g. US2), set the base URL:

```json
{
  "mcpServers": {
    "extend-ai": {
      "command": "extend-mcp-server",
      "env": {
        "EXTEND_API_KEY": "your_api_key_here",
        "EXTEND_BASE_URL": "https://api.us2.extend.app"
      }
    }
  }
}
```

## Troubleshooting

| Problem | Solution |
|---|---|
| `command not found` | Use the full path to the binary, or move it to a directory in your PATH |
| Server won't start | Check that `EXTEND_API_KEY` is set in your MCP client config |
| Tools not appearing | Restart your MCP client after config changes |
| 404 API errors | Check `EXTEND_BASE_URL` matches your regional instance |
| 401 API errors | Verify your API key is valid and active |
| Endpoint not found | Check `EXTEND_API_VERSION` matches available endpoints |

## Contributing

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for development setup, architecture overview, and guidelines.

## Links

- [Model Context Protocol](https://modelcontextprotocol.io/)
- [MCP Go SDK](https://github.com/modelcontextprotocol/go-sdk)
- [Extend AI Docs](https://docs.extend.ai/developers/)
- [Extend AI API Reference](https://docs.extend.ai/api-reference/)

## License

MIT License - See [LICENSE](LICENSE) file for details.
