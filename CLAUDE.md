# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a personal Claude Code marketplace repository for custom plugins and skills. The repository follows the Claude Code plugin marketplace structure, with a `marketplace.json` manifest that defines available plugins.

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
cd plugins/go-dev/servers
npm install
```
