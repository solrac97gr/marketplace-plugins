# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a personal Claude Code marketplace repository for custom plugins and skills. The repository follows the Claude Code plugin marketplace structure, with a `marketplace.json` manifest that defines available plugins.

**Available Plugins:**
1. **go-dev** - Go backend development with hexagonal architecture, DDD, and TDD/BDD
2. **react-dev** - React frontend development with modern best practices and testing
3. **plugin-helper** - Meta-plugin for analyzing projects and generating custom plugins

## Repository Structure

- **marketplace.json**: Root manifest defining the marketplace name, owner, and list of plugins
- **plugins/**: Directory containing individual plugin packages
  - Each plugin has its own `.claude-plugin/plugin.json` metadata file
  - Skills are defined in `skills/[skill-name]/SKILL.md` within each plugin

## Plugin Structure

Each plugin follows this structure:
```
plugins/[plugin-name]/
├── .claude-plugin/
│   └── plugin.json          # Plugin metadata (name, description, version)
└── skills/
    └── [skill-name]/
        └── SKILL.md         # Skill definition with frontmatter and instructions
```

### SKILL.md Format

Skill files use YAML frontmatter followed by the skill instructions:
```markdown
---
description: Brief description of what the skill does
disable-model-invocation: true  # Optional flag
---

[Skill instructions that Claude will follow when the skill is invoked]
```

## Adding New Plugins

1. Create a new directory under `plugins/[plugin-name]/`
2. Add `.claude-plugin/plugin.json` with metadata:
   ```json
   {
     "name": "plugin-name",
     "description": "What the plugin does",
     "version": "1.0.0"
   }
   ```
3. Create skills under `skills/[skill-name]/SKILL.md`
4. Register the plugin in root `marketplace.json` under the `plugins` array

## Adding New Skills to Existing Plugins

1. Create a new directory under `plugins/[plugin-name]/skills/[skill-name]/`
2. Add a `SKILL.md` file with frontmatter and instructions
3. Skills are automatically discovered from the plugin directory structure

## go-dev Plugin Architecture

The `go-dev` plugin creates Go projects following a specific microservices architecture with automated architecture testing.

### Structure Pattern
```
project-name/
├── cmd/
│   ├── main.go                     # Cobra CLI entry point
│   ├── root.go                     # Root command
│   └── serve/
│       ├── [domain].go             # Each domain microservice command
│       └── all.go                  # Run all microservices
├── api/proto/[domain]/             # gRPC definitions (if using gRPC)
├── internal/
│   ├── [domain]/                   # Each bounded context (user, order, etc.)
│   │   ├── domain/                 # Pure domain logic
│   │   ├── application/            # Use cases and DTOs
│   │   └── infrastructure/         # Adapters (DB, HTTP, gRPC)
│   └── shared/                     # Cross-domain shared code
└── test/
    ├── architecture/               # goarchtest tests
    ├── features/[domain]/          # Godog BDD tests per domain
    ├── integration/[domain]/
    └── unit/[domain]/
```

### Key Principles
- **Microservices Monorepo**: Single binary with Cobra commands for each microservice
- **Domain-First**: Organized by bounded context (`internal/[domain]/`)
- **Hexagonal Architecture**: Each domain has domain/application/infrastructure layers
- **Dependency Flow**: Infrastructure → Application → Domain (always inward)
- **Automated Architecture Testing**: Uses goarchtest (https://github.com/solrac97gr/goarchtest) to enforce constraints
- **TDD/BDD**: Features start with Godog feature files, then tests, then implementation
- **Protocol Agnostic**: Supports REST API, gRPC, or both
- **Independent Deployment**: Each domain can run as separate microservice from same binary

### Architecture Testing
The plugin integrates goarchtest to automatically validate:
- Layer dependencies (domain cannot depend on application/infrastructure)
- Domain isolation (domains cannot depend on each other directly)
- Naming conventions (repositories, use cases, handlers)
- Port & Adapter pattern enforcement

Tests run with `make test-arch` and are part of `/review-arch` workflow.

### Intelligent Agents

The plugin includes two specialized agents:

**Architecture Reviewer Agent** (`agents/architecture-reviewer.md`):
- Proactively reviews code changes for architectural violations
- Validates dependency direction and layer purity
- Checks DDD pattern usage
- Provides educational feedback

**DDD Consultant Agent** (`agents/ddd-consultant.md`):
- Strategic DDD guidance (bounded contexts, context mapping)
- Tactical pattern recommendations
- Entity vs Value Object decisions
- Aggregate design assistance

### Automated Hooks

**PostToolUse Hooks:**
1. **Domain Purity Validation** - Triggers on Write/Edit of `internal/*/domain/*.go`
   - Runs `scripts/validate-domain-purity.sh`
   - Validates only stdlib and shared imports
   - Prevents infrastructure dependencies

2. **Architecture Test Execution** - Triggers on Write/Edit of `test/architecture/*_test.go`
   - Runs `go test ./test/architecture/... -v`
   - Immediate feedback on architectural violations

### MCP Server

**GoArchTest Analyzer** (`servers/goarchtest-server`):
- Real-time architecture analysis via MCP protocol
- Tools: check_layer_dependencies, check_domain_isolation, check_naming_conventions
- Runs architecture tests on-demand
- Generates dependency graphs

To use MCP server tools, install dependencies:
```bash
cd plugins/go-dev/servers/goarchtest-server
go mod download
```

## react-dev Plugin Architecture

The `react-dev` plugin provides comprehensive React development tools with TypeScript, modern testing, and accessibility focus.

### Structure Pattern
```
project-name/
├── src/
│   ├── features/                    # Feature-based architecture
│   │   └── [feature-name]/
│   │       ├── components/          # Feature components
│   │       ├── hooks/               # Custom hooks
│   │       ├── services/            # API services
│   │       ├── types/               # TypeScript types
│   │       └── index.ts             # Public API
│   ├── components/                  # Shared components
│   ├── hooks/                       # Shared hooks
│   └── App.tsx
├── tests/                           # Test utilities
└── stories/                         # Storybook stories
```

### Key Principles
- **Feature-Based Architecture**: Organize by feature/domain, not by technical layer
- **Component Composition**: Small, focused components that compose well
- **TypeScript Strict Mode**: Full type safety throughout
- **Accessibility First**: WCAG 2.1 AA compliance built-in
- **Test-Driven Development**: Component → Test → Story workflow
- **Performance Optimized**: Memoization, code splitting, lazy loading
- **Modern React**: Hooks, Context API, Concurrent features

### Validation & Quality

**PostToolUse Hooks:**
1. **Component Purity Validation** - Triggers on Write/Edit of `src/**/*.{tsx,ts,jsx,js}`
   - Runs `scripts/validate-component-purity.sh`
   - Validates React best practices
   - Checks hook naming conventions
   - Ensures TypeScript strict types
   - Validates accessibility attributes

### Intelligent Agents

The plugin includes 10 specialized agents:
1. **Security Reviewer** - XSS prevention, input sanitization, CSP
2. **React Best Practices Enforcer** - Hooks rules, modern patterns
3. **State Management Reviewer** - Context API, state architecture
4. **Tailwind Best Practices Reviewer** - Utility-first patterns
5. **Performance Reviewer** - Memoization, code splitting
6. **Component Architecture Reviewer** - Composition, single responsibility
7. **TypeScript Quality Reviewer** - Strict mode compliance
8. **Hook Usage Reviewer** - Custom hook patterns, Rules of Hooks
9. **Accessibility Reviewer** - WCAG 2.1 AA enforcement
10. **Test Coverage Guardian** - User-centric testing approach

### MCP Server

**Component Analyzer** (`servers/component-analyzer`):
- Real-time component analysis via MCP protocol
- Tools: analyze_component_tree, detect_prop_drilling, check_hook_dependencies
- Complexity analysis and scoring
- Accessibility checking
- Unused props detection

To build the MCP server:
```bash
cd plugins/react-dev/servers/component-analyzer
go build -o component-analyzer component-analyzer.go
```

## plugin-helper Plugin Architecture

The `plugin-helper` plugin is a meta-plugin that analyzes existing codebases and generates custom Claude Code plugins.

### Purpose

Generate project-specific plugins by:
1. Analyzing codebase patterns and conventions
2. Identifying automation opportunities
3. Generating skills, agents, hooks, and MCP servers
4. Creating complete, ready-to-use plugins

### Skills

**`/analyze-project`**:
- Analyzes project structure, patterns, and workflows
- Identifies scaffolding opportunities
- Detects validation rules
- Suggests plugin capabilities
- Generates comprehensive analysis report

**`/generate-plugin`**:
- Creates complete plugin structure
- Generates project-specific skills with templates
- Creates intelligent agents with domain knowledge
- Generates validation hooks (bash scripts)
- Creates MCP servers in Go (optional)
- Produces full documentation

### Plugin Architect Agent

Expert at:
- Codebase pattern recognition
- Plugin architecture design
- Capability prioritization
- Best practices enforcement

Activates during plugin analysis and generation to provide guidance.

### Use Cases

1. **Team Onboarding**: Generate plugins that capture team conventions
2. **Project Consistency**: Automate scaffolding for consistent patterns
3. **Knowledge Capture**: Turn tribal knowledge into automated tools
4. **Multi-Project**: Generate custom plugins per project type
5. **Open Source**: Create contributor-friendly automation

### Generated Plugin Structure

```
plugins/[custom-plugin-name]/
├── .claude-plugin/
│   └── plugin.json              # With hooks and MCP servers
├── skills/
│   └── [skill-name]/
│       └── SKILL.md             # Project-specific skills
├── agents/
│   └── [agent-name].md          # Domain-aware agents
├── scripts/
│   └── validate-*.sh            # Validation hooks
├── servers/
│   └── [server-name]/           # Go MCP server (optional)
├── templates/                   # Code templates
└── documentation                # README, ARCHITECTURE, etc.
```

### Key Features

- **Project-Specific**: Uses actual project patterns, not generic templates
- **Template Extraction**: Extracts templates from existing code
- **Convention Matching**: Follows project naming and structure
- **Language Agnostic**: Works with any programming language
- **Complete Generation**: Generates all plugin components
- **Documentation**: Creates comprehensive docs automatically

## MCP Server Consistency

All plugins use **Go binaries** for MCP servers for consistency:
- **go-dev**: goarchtest-server (6.4 MB)
- **react-dev**: component-analyzer (6.2 MB)
- **plugin-helper**: No MCP server (generates them for other plugins)

Build any MCP server:
```bash
cd plugins/[plugin-name]/servers/[server-name]
go mod download
go build -o [server-name] [server-name].go
```
