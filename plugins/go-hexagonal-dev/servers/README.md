# GoArchTest MCP Server

Model Context Protocol (MCP) server for real-time Go architecture analysis using goarchtest.

## Features

The MCP server exposes the following tools to Claude:

### 1. `check_layer_dependencies`
Check if a specific layer has illegal dependencies on other layers.

**Parameters:**
- `layer`: domain, application, or infrastructure
- `domain`: The bounded context to check (e.g., "user", "order")

**Example:**
```json
{
  "layer": "domain",
  "domain": "user"
}
```

### 2. `check_domain_isolation`
Verify that domains are properly isolated from each other.

**Parameters:**
- `sourceDomain`: Domain to check
- `targetDomain`: Domain that should not be imported

**Example:**
```json
{
  "sourceDomain": "user",
  "targetDomain": "order"
}
```

### 3. `check_naming_conventions`
Validate naming conventions for repositories, use cases, and handlers.

**Parameters:**
- `pattern`: repository, usecase, or handler

### 4. `run_all_architecture_tests`
Execute all architecture tests in `test/architecture/`.

### 5. `generate_dependency_graph`
Generate a DOT format dependency graph visualization.

**Parameters:**
- `domain` (optional): Specific domain to visualize

## Installation

```bash
cd servers
npm install
```

## Usage

The MCP server is automatically configured in the plugin's `plugin.json`. It starts when the plugin is loaded.

Claude can invoke these tools during development to:
- Validate architecture in real-time
- Check dependencies before committing
- Generate visualizations of code structure
- Enforce architectural constraints proactively

## How It Works

The server uses `goarchtest` library to analyze Go code structure and dependencies. It generates temporary test files and executes them to validate architectural constraints.

## Development

To test the server manually:

```bash
node goarchtest-server --project-root /path/to/go/project
```

The server communicates via stdio following the MCP protocol.
