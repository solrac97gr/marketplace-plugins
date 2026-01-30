# Marketplace Plugins

A personal collection of custom plugins and skills for [Claude Code](https://claude.ai/code).

## What is this?

This repository contains custom plugins that extend Claude Code's functionality through user-defined skills. Each plugin can provide one or more skills that can be invoked during Claude Code sessions using the `/skill-name` command.

## Available Plugins

### go-hexagonal-dev

Complete toolkit for Go backend development following hexagonal architecture, DDD, and TDD/BDD with Godog. Features automated architecture testing, intelligent agents, and real-time validation hooks.

**Version:** 1.0.0 | **Category:** Development | **License:** MIT

**Keywords:** golang, hexagonal-architecture, ddd, tdd, bdd, godog, microservices, goarchtest

**Skills included:**

- **`/start-project`** - Initialize a new Go project with complete hexagonal architecture structure
  - Discovers project type and suggests relevant domains/bounded contexts
  - User selects which domains to create (multi-select)
  - Supports REST API, gRPC, or both
  - Supports multiple databases (PostgreSQL, MongoDB, MySQL, SQLite)
  - Creates Cobra CLI with commands for each microservice
  - Sets up automated architecture tests using goarchtest
  - Configures Godog for BDD testing
  - Single binary, multiple microservices architecture

- **`/new-feature`** - Scaffold a complete feature following TDD/BDD workflow
  - Starts by creating Godog feature file (Gherkin syntax)
  - Asks clarifying questions until feature is clear
  - Generates domain entities, use cases, repositories, handlers, and tests
  - Follows strict TDD: feature file → tests → implementation

- **`/new-entity`** - Create a domain entity with repository interface
  - Generates entity with business logic
  - Creates repository interface (port)
  - Includes unit tests
  - Follows DDD principles

- **`/new-usecase`** - Create an application use case with tests
  - Generates use case in application layer
  - Creates DTOs
  - Writes unit tests with mocks
  - Follows TDD approach

- **`/review-arch`** - Review code for architecture compliance
  - **Executes automated architecture tests** using goarchtest
  - Validates hexagonal architecture dependency rules
  - Checks DDD principles adherence
  - Reviews test coverage
  - Identifies anti-patterns
  - Combines automated testing with AI analysis
  - Provides actionable recommendations

- **`/update-arch-tests`** - Create or update architecture tests
  - Analyzes project structure and domains
  - Generates comprehensive goarchtest test suites
  - Tests layer dependencies (domain → application → infrastructure)
  - Tests domain isolation (no cross-domain dependencies)
  - Tests naming conventions
  - Custom rules for project-specific patterns

**Intelligent Agents:**

- **Architecture Reviewer Agent** - Proactive code review
  - Automatically reviews files in `internal/` directory
  - Validates dependency direction (Infrastructure → Application → Domain)
  - Checks layer purity (domain has no external dependencies)
  - Identifies DDD pattern violations
  - Educational feedback with explanations

- **DDD Consultant Agent** - Domain modeling guidance
  - Strategic DDD advice (bounded contexts, context mapping)
  - Tactical pattern recommendations (entities, value objects, aggregates)
  - Entity vs Value Object decision support
  - Aggregate design assistance
  - Ubiquitous language enforcement

**Automated Hooks:**

- **Domain Purity Validation** - Runs on domain file edits
  - Validates no external dependencies in domain layer
  - Only allows stdlib and shared kernel imports
  - Prevents infrastructure leakage

- **Architecture Test Execution** - Runs on test file updates
  - Automatically executes architecture tests
  - Immediate feedback on violations

**MCP Server - GoArchTest Analyzer:**

Real-time architecture analysis tools:
- `check_layer_dependencies` - Validate layer boundaries
- `check_domain_isolation` - Verify domain separation
- `check_naming_conventions` - Enforce naming standards
- `run_all_architecture_tests` - Execute full test suite
- `generate_dependency_graph` - Visualize dependencies

## Installation

To use these plugins with Claude Code:

1. Clone this repository
2. Configure your Claude Code marketplace settings to point to this repository
3. The plugins will be available in your Claude Code sessions

## Repository Structure

```
marketplace-plugins/
├── .claude-plugin/
│   └── marketplace.json      # Marketplace manifest
├── plugins/
│   └── [plugin-name]/
│       ├── .claude-plugin/
│       │   └── plugin.json   # Plugin metadata
│       └── skills/
│           └── [skill-name]/
│               └── SKILL.md  # Skill definition
└── CLAUDE.md                 # Development guide
```

## Creating a New Plugin

1. Create a new directory under `plugins/`:
   ```bash
   mkdir -p plugins/my-plugin/.claude-plugin
   mkdir -p plugins/my-plugin/skills/my-skill
   ```

2. Add plugin metadata in `.claude-plugin/plugin.json`:
   ```json
   {
     "name": "my-plugin",
     "description": "Description of what your plugin does",
     "version": "1.0.0"
   }
   ```

3. Create a skill in `skills/my-skill/SKILL.md`:
   ```markdown
   ---
   description: Brief description
   disable-model-invocation: true
   ---

   Instructions for Claude when this skill is invoked.
   ```

4. Register the plugin in `.claude-plugin/marketplace.json`:
   ```json
   {
     "plugins": [
       {
         "name": "my-plugin",
         "source": "./plugins/my-plugin",
         "description": "Description of what your plugin does"
       }
     ]
   }
   ```

## Contributing

This is a personal marketplace, but feel free to fork and create your own custom plugins!

## License

MIT
