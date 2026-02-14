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

## Installation

Choose one of the three options below. After installing, jump to [Setup](#setup) to configure your MCP client.

### Option 1: npx (recommended)

Requires [Node.js 18+](https://nodejs.org/). No install step -- runs directly:

```
npx -y extend-ai-mcp-server
```

> On first run, it downloads the correct binary for your OS and caches it at `~/.extend-mcp-server/`. Subsequent runs start instantly.

### Option 2: Go

Requires [Go 1.23+](https://go.dev/dl/). Also runs directly, no install step:

```
go run github.com/caiomoura/extend-mcp-server/cmd/extend-mcp-server@latest
```

> Go downloads and compiles the server automatically. The build is cached for fast subsequent starts.

### Option 3: Pre-built binary

No runtime dependencies. Download from [GitHub Releases](https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases):

<details>
<summary>macOS (Apple Silicon)</summary>

```bash
curl -Lo extend-mcp-server.tar.gz https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases/latest/download/extend-mcp-server_darwin_arm64.tar.gz
tar xzf extend-mcp-server.tar.gz && rm extend-mcp-server.tar.gz
sudo mv extend-mcp-server /usr/local/bin/
```

</details>

<details>
<summary>macOS (Intel)</summary>

```bash
curl -Lo extend-mcp-server.tar.gz https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases/latest/download/extend-mcp-server_darwin_amd64.tar.gz
tar xzf extend-mcp-server.tar.gz && rm extend-mcp-server.tar.gz
sudo mv extend-mcp-server /usr/local/bin/
```

</details>

<details>
<summary>Linux (x86_64)</summary>

```bash
curl -Lo extend-mcp-server.tar.gz https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases/latest/download/extend-mcp-server_linux_amd64.tar.gz
tar xzf extend-mcp-server.tar.gz && rm extend-mcp-server.tar.gz
sudo mv extend-mcp-server /usr/local/bin/
```

</details>

<details>
<summary>Linux (ARM64)</summary>

```bash
curl -Lo extend-mcp-server.tar.gz https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases/latest/download/extend-mcp-server_linux_arm64.tar.gz
tar xzf extend-mcp-server.tar.gz && rm extend-mcp-server.tar.gz
sudo mv extend-mcp-server /usr/local/bin/
```

</details>

<details>
<summary>Windows (x86_64)</summary>

1. Download [extend-mcp-server_windows_amd64.zip](https://github.com/caiomoura1994/Extend-AI-MCP-Server/releases/latest/download/extend-mcp-server_windows_amd64.zip)
2. Extract `extend-mcp-server.exe`
3. Move it to a directory in your PATH (e.g. `C:\Program Files\extend-mcp-server\`)

</details>

## Setup

Add the following to your MCP client config. Replace the `command` and `args` depending on which installation option you chose:

| Installation | `command` | `args` |
|---|---|---|
| **npx** | `"npx"` | `["-y", "extend-ai-mcp-server"]` |
| **Go** | `"go"` | `["run", "github.com/caiomoura/extend-mcp-server/cmd/extend-mcp-server@latest"]` |
| **Binary** | `"/usr/local/bin/extend-mcp-server"` | not needed |

Example using npx:

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

Where to put this config:

| Client | Config file |
|---|---|
| **Claude Desktop** (macOS) | `~/Library/Application Support/Claude/claude_desktop_config.json` |
| **Claude Desktop** (Windows) | `%APPDATA%\Claude\claude_desktop_config.json` |
| **Cursor** | `.cursor/mcp.json` in your project, or global Cursor MCP settings |

Restart your MCP client after saving.

## Configuration

All configuration is done via environment variables in the `env` block of your MCP config:

| Variable | Required | Default | Description |
|---|---|---|---|
| `EXTEND_API_KEY` | Yes | - | Your Extend AI API key |
| `EXTEND_API_VERSION` | No | `2026-02-09` | Extend AI API version |
| `EXTEND_BASE_URL` | No | `https://api.extend.ai` | Extend AI API base URL |

### Regional Endpoints

If your Extend AI account is on a regional instance (e.g. US2), add `EXTEND_BASE_URL` to the `env` block:

```json
"env": {
  "EXTEND_API_KEY": "your_api_key_here",
  "EXTEND_BASE_URL": "https://api.us2.extend.app"
}
```

## Troubleshooting

| Problem | Solution |
|---|---|
| `command not found` | Use the full path to the binary, or ensure Node.js/Go is in your PATH |
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
