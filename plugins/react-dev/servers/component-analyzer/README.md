# Component Analyzer MCP Server

A Go-based MCP (Model Context Protocol) server for analyzing React components, hooks, and component architecture.

## Features

This server provides the following analysis tools:

### 1. analyze_component_tree
Analyzes component hierarchy and structure metrics:
- Nesting depth
- Number of child components
- Custom hooks usage
- Conditional rendering patterns

**Parameters:**
- `componentPath` (required): Path to the component file

### 2. detect_prop_drilling
Detects prop drilling anti-patterns in your component tree.

**Parameters:**
- `featurePath` (required): Path to feature directory
- `depth` (optional): Minimum prop passing depth to flag (default: 3)

### 3. check_hook_dependencies
Analyzes React hook dependency arrays for common issues:
- Empty dependency arrays
- Missing dependencies
- useEffect without deps array

**Parameters:**
- `filePath` (required): Path to file with hooks

### 4. analyze_component_complexity
Calculates component complexity metrics:
- Lines of code
- Number of hooks
- Conditional branches
- State variables count
- Overall complexity score

**Parameters:**
- `componentPath` (required): Path to component to analyze

### 5. find_unused_props
Finds props defined in the interface but not used in the component.

**Parameters:**
- `componentPath` (required): Path to component file

### 6. check_accessibility
Performs basic accessibility checks:
- Images without alt attributes
- Buttons without labels
- Inputs without labels
- Interactive handlers on non-interactive elements

**Parameters:**
- `componentPath` (required): Path to component file

## Building

```bash
go mod download
go build -o component-analyzer component-analyzer.go
```

## Usage

The server is automatically invoked by Claude Code when the react-dev plugin is active.

You can also run it manually:

```bash
./component-analyzer [project-root-path]
```

## Integration

This server is configured in the react-dev plugin's `plugin.json`:

```json
{
  "mcpServers": {
    "component-analyzer": {
      "command": "${CLAUDE_PLUGIN_ROOT}/servers/component-analyzer/component-analyzer",
      "description": "React component analysis tools"
    }
  }
}
```

## Architecture

The server uses the [mcp-go](https://github.com/mark3labs/mcp-go) library to implement the Model Context Protocol, providing analysis capabilities through stdio communication.
