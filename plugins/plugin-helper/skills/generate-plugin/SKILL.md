---
description: Generate custom plugin from analysis
---

Generate a complete custom Claude Code plugin based on the project analysis performed by `/analyze-project`.

## Prerequisites

You must have run `/analyze-project` first to understand the project patterns and generate the analysis report.

## What This Skill Creates

A complete plugin structure with:
1. Plugin metadata (plugin.json)
2. Custom skills (SKILL.md files)
3. Intelligent agents (agent.md files)
4. Validation hooks (bash scripts)
5. MCP server (if needed, in Go)
6. Documentation (README.md, ARCHITECTURE.md)

## Generation Process

### Step 1: Confirm Plugin Details

Ask the user to confirm/customize:
1. **Plugin name**: [suggested-name] or custom
2. **Plugin description**: Brief description
3. **Skills to include**: Which suggested skills?
4. **Agents to include**: Which suggested agents?
5. **Hooks to include**: Which validation hooks?
6. **MCP server**: Should one be created?

### Step 2: Create Plugin Structure

Generate the following directory structure:

```
plugins/[plugin-name]/
├── .claude-plugin/
│   └── plugin.json                    # Plugin metadata
├── skills/
│   ├── [skill-1]/
│   │   └── SKILL.md                   # Skill definition
│   ├── [skill-2]/
│   │   └── SKILL.md
│   └── ...
├── agents/
│   ├── [agent-1].md                   # Agent definition
│   ├── [agent-2].md
│   └── ...
├── scripts/
│   ├── validate-[pattern].sh          # Validation scripts
│   └── ...
├── servers/
│   └── [server-name]/                 # MCP server (if needed)
│       ├── [server-name].go
│       ├── go.mod
│       ├── go.sum
│       └── README.md
├── templates/
│   └── [template-files]               # Code templates
├── ARCHITECTURE.md                     # Plugin architecture docs
├── CHANGELOG.md                        # Version history
├── CODE_GENERATION_RULES.md           # Code gen guidelines
└── README.md                          # Plugin documentation
```

### Step 3: Generate plugin.json

Create the plugin metadata file with:

```json
{
  "name": "[plugin-name]",
  "description": "[description based on analysis]",
  "version": "1.0.0",
  "author": {
    "name": "[from git config or ask]",
    "email": "[from git config or ask]"
  },
  "hooks": {
    "PreToolUse": [],
    "PostToolUse": [
      {
        "matcher": "Write|Edit",
        "matchPath": "[file-pattern]",
        "hooks": [
          {
            "type": "command",
            "command": "${CLAUDE_PLUGIN_ROOT}/scripts/[validation-script].sh \"${TOOL_FILE_PATH}\"",
            "description": "[what it validates]"
          }
        ]
      }
    ]
  },
  "mcpServers": {
    "[server-name]": {
      "command": "${CLAUDE_PLUGIN_ROOT}/servers/[server-name]/[server-name]",
      "description": "[server purpose]"
    }
  }
}
```

### Step 4: Generate Skills

For each skill, create a SKILL.md file with:

```markdown
---
description: [Brief description]
---

[Detailed instructions for Claude when this skill is invoked]

## What to Create

1. Ask the user:
   - [Required parameters]
   - [Optional parameters]

2. Generate [what structures/files]:
   - [File 1] in [location]
   - [File 2] in [location]
   - [etc.]

## File Templates

### [File Type 1]
\`\`\`[language]
[Template content based on project patterns]
\`\`\`

### [File Type 2]
\`\`\`[language]
[Template content based on project patterns]
\`\`\`

## Best Practices

- [Practice 1 from project]
- [Practice 2 from project]
- [etc.]

## Example

\`\`\`bash
/[skill-name] [example-args]
\`\`\`
```

**Skill Content Guidelines:**
1. Use project's actual patterns and conventions
2. Include file templates based on existing code
3. Reference project's documentation style
4. Follow naming conventions found in codebase
5. Use project's preferred testing approach
6. Match the architectural patterns detected

### Step 5: Generate Agents

For each agent, create an agent.md file with:

```markdown
---
name: [Agent Name]
description: [What this agent does]
skills:
  - [related-skill-1]
  - [related-skill-2]
---

# [Agent Name]

You are an expert in [domain based on project].

## Your Role

[Describe the agent's expertise and when it activates]

## When to Activate

Automatically activate when:
- [Trigger condition 1]
- [Trigger condition 2]
- The user runs `/[related-skill]` command
- [Architecture/quality violations detected]

## What to Review

### 1. [Check Category 1]
- ✅ [What's good]
- ❌ [What's bad]
- [Specific rules from project]

### 2. [Check Category 2]
- ✅ [What's good]
- ❌ [What's bad]
- [Specific rules from project]

### 3. [Check Category 3]
[Based on project's architecture and patterns]

## Common Anti-Patterns

❌ **[Anti-pattern 1]**: [Description based on project issues]
❌ **[Anti-pattern 2]**: [Description based on project issues]
❌ **[Anti-pattern 3]**: [Description based on project issues]

## Review Process

1. Analyze the file/changes
2. Check against project conventions
3. Identify violations or improvements
4. Provide specific, actionable feedback
5. Reference project documentation when relevant

## Feedback Format

Use this format for feedback:

\`\`\`
✅ Good: [What follows project patterns]
❌ Issue: [What violates project rules]
💡 Suggestion: [How to fix it]
📚 Reference: [Link to project docs/examples]
\`\`\`

## Example Violations

[Provide examples based on actual project code]
```

**Agent Content Guidelines:**
1. Base checks on project's actual architecture
2. Reference project-specific patterns
3. Use project's terminology and ubiquitous language
4. Include examples from the actual codebase
5. Link to project documentation
6. Enforce project-specific rules, not generic ones

### Step 6: Generate Validation Scripts

For each hook, create a bash script:

```bash
#!/bin/bash

# [Script Purpose]
# This script validates [what based on project rules]

set -e

FILE_PATH="$1"

# Check if file matches pattern
if [[ ! "$FILE_PATH" =~ [regex-pattern] ]]; then
    exit 0  # Not a relevant file, skip
fi

echo "🔍 Validating [what]: $FILE_PATH"

VIOLATIONS=0

# Check 1: [Project-specific rule]
if [condition based on project]; then
    echo "❌ VIOLATION: [Message]"
    VIOLATIONS=$((VIOLATIONS + 1))
fi

# Check 2: [Project-specific rule]
if [condition based on project]; then
    echo "❌ VIOLATION: [Message]"
    VIOLATIONS=$((VIOLATIONS + 1))
fi

# [More checks based on project]

if [ $VIOLATIONS -gt 0 ]; then
    echo ""
    echo "⚠️  [Type] validation failed with $VIOLATIONS violation(s)"
    echo ""
    echo "[Project name] Best Practices:"
    echo "  - [Rule 1 from project]"
    echo "  - [Rule 2 from project]"
    echo "  - [Rule 3 from project]"
    echo ""
    exit 1
fi

echo "✅ [Type] validated"
exit 0
```

**Script Content Guidelines:**
1. Check project-specific rules, not generic ones
2. Use project's file patterns and conventions
3. Reference project documentation in error messages
4. Make scripts executable: `chmod +x`
5. Test scripts on actual project files

### Step 7: Generate MCP Server (Optional)

If the analysis suggests an MCP server is needed, create a Go-based server:

**When to create MCP servers:**
- Project has custom tooling that could be exposed
- Complex analysis needed (architecture, dependencies, etc.)
- Project-specific linting or validation
- Integration with project-specific tools

**Server structure:**
```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type [ProjectName]Server struct {
	projectRoot string
	mcpServer   *server.MCPServer
}

func New[ProjectName]Server(projectRoot string) *[ProjectName]Server {
	if projectRoot == "" {
		projectRoot, _ = os.Getwd()
	}

	s := &[ProjectName]Server{
		projectRoot: projectRoot,
	}

	mcpServer := server.NewMCPServer(
		"[server-name]",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	s.mcpServer = mcpServer
	s.setupHandlers()

	return s
}

func (s *[ProjectName]Server) setupHandlers() {
	// Register tools based on project needs
	s.mcpServer.AddTool(
		mcp.NewTool("[tool-name]",
			mcp.WithDescription("[what it does for project]"),
			mcp.WithString("[param]",
				mcp.Required(),
				mcp.Description("[param description]"),
			),
		),
		s.[toolHandler],
	)
}

func (s *[ProjectName]Server) [toolHandler](ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Implementation based on project needs
	return mcp.NewToolResultText("Result"), nil
}

func main() {
	projectRoot := ""
	if len(os.Args) > 1 {
		projectRoot = os.Args[1]
	}

	s := New[ProjectName]Server(projectRoot)

	fmt.Fprintln(os.Stderr, "[ProjectName] MCP server running")

	if err := server.ServeStdio(s.mcpServer); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
```

Include:
- `go.mod` with `github.com/mark3labs/mcp-go v0.43.2`
- Build instructions in README.md
- Tool documentation

### Step 8: Generate Documentation

#### README.md
```markdown
# [Plugin Name]

[Description based on project]

## Features

- [Feature 1]
- [Feature 2]
- [Feature 3]

## Skills

### /[skill-1]
[Description]

**Usage:**
\`\`\`bash
/[skill-1] [args]
\`\`\`

### /[skill-2]
[Description]

[More skills...]

## Agents

### [Agent 1]
[Description and when it activates]

### [Agent 2]
[Description]

[More agents...]

## Installation

1. Clone this marketplace repository
2. The plugin will be available in Claude Code sessions
3. (If MCP server exists) Build the server:
   \`\`\`bash
   cd servers/[server-name]
   go build -o [server-name] [server-name].go
   \`\`\`

## Usage Examples

[Examples specific to the project]

## Architecture

See [ARCHITECTURE.md](./ARCHITECTURE.md) for details on how this plugin is structured.

## Contributing

[Project-specific contribution guidelines]
```

#### ARCHITECTURE.md
```markdown
# Plugin Architecture

## Overview

This plugin is designed to support [project type] development following [project's patterns].

## Structure

[Explain the plugin organization]

## Skills

[Describe each skill's purpose and implementation]

## Agents

[Describe each agent's role and triggers]

## Hooks

[Explain validation hooks and what they check]

## MCP Server (if applicable)

[Explain the MCP server's tools and purpose]

## Design Decisions

[Document key decisions made during plugin design]
```

### Step 9: Final Steps

1. **Test the plugin**:
   - Verify all SKILL.md files are valid
   - Check agent.md files are properly formatted
   - Test validation scripts on sample files
   - Build MCP server if created

2. **Add to marketplace**:
   - Update `marketplace.json` to include the new plugin
   - Commit the plugin to git

3. **Documentation**:
   - Create usage examples
   - Document any prerequisites
   - Add troubleshooting guide

4. **Provide summary** to user:
   ```markdown
   ## ✅ Plugin Generated Successfully!

   **Plugin**: [plugin-name]
   **Location**: plugins/[plugin-name]/

   **Contents:**
   - [X] plugin.json
   - [X] [N] skills
   - [X] [N] agents
   - [X] [N] validation hooks
   - [X] MCP server (if applicable)
   - [X] Documentation

   **Next Steps:**
   1. Review the generated files
   2. Customize as needed
   3. Test with: [test command]
   4. Use skills: /[skill-name]

   **Available Skills:**
   - /[skill-1]: [description]
   - /[skill-2]: [description]
   - [...]
   ```

## Important Guidelines

1. **Use actual project patterns**: Don't generate generic templates, use the project's real patterns
2. **Extract from existing code**: Find existing files as examples and base templates on them
3. **Match project style**: Naming, formatting, structure should match the project
4. **Include project context**: Reference project docs, architecture, and conventions
5. **Test on real files**: Validate that generated code works with actual project files
6. **Make it useful**: Focus on common, repetitive tasks developers actually do
7. **Document thoroughly**: Explain why each component exists and how to use it

## Example Output

For a React + TypeScript project:
- Generated `react-my-project` plugin
- 5 skills for common tasks
- 3 agents for code review
- 2 validation hooks
- Component analysis MCP server
- Complete documentation

All based on the actual project's patterns and conventions.
