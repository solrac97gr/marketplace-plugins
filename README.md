# Marketplace Plugins

A personal collection of custom plugins and skills for [Claude Code](https://claude.ai/code).

## What is this?

This repository contains custom plugins that extend Claude Code's functionality through user-defined skills. Each plugin can provide one or more skills that can be invoked during Claude Code sessions using the `/skill-name` command.

## Available Plugins

### 1. go-dev

Complete toolkit for Go backend development following hexagonal architecture, Domain-Driven Design (DDD), and Test-Driven Development with Godog BDD framework.

**Features:**
- 7 production-ready skills for project scaffolding and development
- 9 intelligent AI agents for architecture review, DDD guidance, and quality assurance
- Real-time MCP server (Go) for architecture testing with goarchtest
- Automated validation hooks for domain purity and architecture tests

**Skills:** `/start-project`, `/new-feature`, `/new-entity`, `/new-usecase`, `/from-diagram`, `/review-arch`, `/update-arch-tests`

**[📖 Full Documentation](./plugins/go-dev/README.md)**

---

### 2. react-dev

Comprehensive React development plugin with TypeScript, Vite, Tailwind CSS, and modern testing tools. Provides TDD workflows, accessibility compliance, performance optimization, and intelligent code review agents.

**Features:**
- 10 production-ready skills for React development workflows
- 10 specialized AI agents covering security, performance, accessibility, and best practices
- Real-time MCP server (Go) for component analysis and complexity metrics
- Automated validation hooks for component purity and best practices
- Playwright integration for E2E testing

**Skills:** `/start-project`, `/new-component`, `/new-hook`, `/new-context`, `/new-feature`, `/from-diagram`, `/review-code`, `/review-a11y`, `/update-tests`, `/optimize-performance`

**[📖 Full Documentation](./plugins/react-dev/README.md)**

---

### 3. plugin-helper

Meta-plugin for analyzing projects and generating custom Claude Code plugins tailored to your codebase patterns and workflows.

**Features:**
- Analyzes any codebase to identify patterns and automation opportunities
- Generates complete custom plugins with skills, agents, hooks, and MCP servers
- Extracts templates from actual project code for consistency
- Creates project-specific documentation automatically
- Language-agnostic pattern recognition
- Plugin Architect agent for intelligent plugin design

**Skills:** `/analyze-project`, `/generate-plugin`

**Use Cases:**
- Generate team-specific plugins for onboarding and consistency
- Capture tribal knowledge as automated tools
- Create project-specific scaffolding and validation
- Automate repetitive development workflows

**[📖 Full Documentation](./plugins/plugin-helper/README.md)**

---

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
│       ├── skills/           # Skill definitions
│       ├── agents/           # AI agent definitions
│       ├── servers/          # MCP servers
│       └── README.md         # Plugin documentation
└── CLAUDE.md                 # Development guide
```

## Quick Start

### Using Existing Plugins

1. Clone this repository
2. Configure Claude Code to use this marketplace
3. Start using skills:
   ```bash
   # Go backend development
   /start-project my-backend-api
   /new-feature user-management

   # React frontend development
   /start-project my-react-app
   /new-component Button

   # Generate custom plugin for your project
   /analyze-project
   /generate-plugin
   ```

### Creating Custom Plugins

**Easy Way (Recommended):**
Use the `plugin-helper` plugin to generate a custom plugin for your project:

```bash
# Analyze your project
/analyze-project

# Review the analysis report
# Then generate the plugin
/generate-plugin
```

**Manual Way:**

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

## Plugin Statistics

| Plugin | Skills | Agents | MCP Servers | Hooks | LOC |
|--------|--------|--------|-------------|-------|-----|
| go-dev | 7 | 9 | 1 (Go) | 2 | ~12,000 |
| react-dev | 10 | 10 | 1 (Go) | 1 | ~13,000 |
| plugin-helper | 2 | 1 | 0 | 0 | ~5,000 |
| **Total** | **19** | **20** | **2** | **3** | **~30,000** |

## Technology Stack

### MCP Servers
All MCP servers are written in **Go** for consistency and performance:
- Built with `github.com/mark3labs/mcp-go v0.43.2`
- Compiled to single binaries (~6 MB each)
- No runtime dependencies
- Fast startup and execution

### Validation Hooks
All validation hooks are **bash scripts**:
- Fast execution (< 1 second)
- Universal compatibility
- Easy to debug and maintain
- Clear error messages

### Skills & Agents
Written in **Markdown** with YAML frontmatter:
- Easy to read and edit
- Version control friendly
- No compilation needed
- Human and machine readable

## Examples

### Generate a Backend API with go-dev

```bash
# Start a new Go project with microservices architecture
/start-project ecommerce-api

# You'll be guided through:
# - Project type selection
# - Domain identification (user, product, order, payment)
# - Protocol choice (REST, gRPC, or both)
# - Database selection

# Add a new feature with TDD/BDD
/new-feature user-authentication

# Result: Complete feature with:
# - Domain entities and logic
# - Use cases and DTOs
# - Repository interfaces
# - HTTP/gRPC handlers
# - Godog BDD tests
# - Architecture tests
```

### Build a React App with react-dev

```bash
# Start a new React project with modern setup
/start-project my-dashboard

# Creates: Vite + React 18 + TypeScript + Tailwind + Vitest + Storybook

# Create a new feature
/new-feature user-profile

# Result: Complete feature with:
# - Feature directory structure
# - Component files with tests
# - Custom hooks
# - API service layer
# - TypeScript types
# - Storybook stories
```

### Generate Custom Plugin with plugin-helper

```bash
# Analyze your existing project
/analyze-project

# Input:
# - Project path: /Users/you/my-python-fastapi-app
# - Project type: Backend API
# - Technologies: Python, FastAPI, SQLAlchemy, Alembic
# - Workflows: Creating endpoints, models, migrations
# - Pain points: Repetitive CRUD setup

# Output: Analysis report with suggestions

# Generate the plugin
/generate-plugin

# Result: Custom plugin my-python-fastapi-app-dev with:
# - /new-endpoint skill
# - /new-model skill
# - /new-migration skill
# - API Best Practices agent
# - SQLAlchemy Pattern Reviewer agent
# - Validation hooks for Python style
```

## Use Cases

### Team Onboarding
- New developers get productive immediately with project-specific skills
- Capture team conventions and best practices in agents
- Automate repetitive setup tasks
- Ensure consistency across the team

### Open Source Projects
- Make it easy for contributors to add features correctly
- Reduce PR review time with automated validation
- Document patterns through working code generation
- Lower the barrier to contribution

### Enterprise Development
- Enforce architectural standards automatically
- Ensure security and compliance requirements
- Scale best practices across multiple teams
- Reduce technical debt through consistency

### Personal Projects
- Speed up development workflows
- Maintain consistency even when context-switching
- Experiment with new patterns quickly
- Build faster without sacrificing quality

## Architecture Patterns Supported

### Backend
- **Hexagonal Architecture** (Ports & Adapters) - go-dev
- **Clean Architecture** - go-dev
- **Domain-Driven Design** - go-dev
- **Microservices** - go-dev
- **REST APIs** - go-dev
- **gRPC Services** - go-dev

### Frontend
- **Feature-Based Architecture** - react-dev
- **Component Composition** - react-dev
- **Atomic Design** - react-dev (customizable)
- **JAMstack** - react-dev
- **Server Components** - react-dev (Next.js compatible)

### Custom
- **Any Pattern** - plugin-helper analyzes and generates

## Contributing

This is a personal marketplace, but you're welcome to:
- Fork and create your own custom plugins
- Share generated plugins with the community
- Report issues or suggest improvements
- Use plugin-helper to generate plugins for your projects

## Troubleshooting

### Plugin not showing up
- Check `.claude-plugin/marketplace.json` includes the plugin
- Verify `plugin.json` exists in the plugin directory
- Restart Claude Code

### MCP Server not working
- Ensure the binary is built: `cd servers/[name] && go build`
- Check binary is executable: `chmod +x [name]`
- Verify go.mod has correct dependencies

### Hooks not triggering
- Check file path regex in plugin.json
- Ensure script is executable: `chmod +x script.sh`
- Test script manually with a file path

### Skills not generating code correctly
- Review the SKILL.md for proper instructions
- Check if templates match your project structure
- Use plugin-helper to regenerate with updated analysis

## Roadmap

### Short Term
- [ ] Add more example plugins
- [ ] Improve plugin-helper analysis algorithms
- [ ] Add plugin testing framework
- [ ] Create plugin marketplace submission tool

### Long Term
- [ ] Visual plugin builder
- [ ] Plugin composition (extend/inherit plugins)
- [ ] Plugin analytics and usage tracking
- [ ] Community plugin templates
- [ ] Multi-language monorepo support

## Acknowledgments

Built with:
- [Claude Code](https://claude.ai/code) by Anthropic
- [mcp-go](https://github.com/mark3labs/mcp-go) for MCP servers
- [goarchtest](https://github.com/solrac97gr/goarchtest) for architecture testing

## License

MIT

---

**Total Capabilities:**
- 19 skills across all plugins
- 20 intelligent agents
- 2 Go MCP servers
- 3 validation hooks
- ~30,000 lines of code
- Supports 10+ project types
- Language-agnostic plugin generation

Ready to supercharge your development workflow! 🚀
