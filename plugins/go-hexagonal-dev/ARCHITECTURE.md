# Go Hexagonal Dev Plugin - Architecture

This document describes the internal architecture of the plugin itself and how all components work together.

## Plugin Structure

```
go-hexagonal-dev/
├── .claude-plugin/
│   └── plugin.json              # Plugin manifest with metadata, hooks, MCP servers
├── skills/                      # User-invocable skills
│   ├── start-project/
│   ├── new-feature/
│   ├── new-entity/
│   ├── new-usecase/
│   ├── review-arch/
│   └── update-arch-tests/
├── agents/                      # Intelligent agents
│   ├── architecture-reviewer.md
│   └── ddd-consultant.md
├── scripts/                     # Automation scripts
│   └── validate-domain-purity.sh
├── servers/                     # MCP servers
│   ├── goarchtest-server
│   ├── package.json
│   └── README.md
├── README.md                    # User documentation
├── CHANGELOG.md                 # Version history
└── ARCHITECTURE.md              # This file
```

## Component Interaction

```
┌─────────────────────────────────────────────────────────────┐
│                         User Input                          │
│                   (Invokes /start-project)                  │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                    Skill Execution                          │
│  1. Ask project type                                        │
│  2. Suggest domains (using domain knowledge)                │
│  3. User selects domains (multi-select)                     │
│  4. Ask technical preferences (DB, protocol, framework)     │
│  5. Generate project structure                              │
│  6. Create architecture tests                               │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                   File Creation                             │
│  (Write tool creates domain files)                          │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                PostToolUse Hook Triggers                    │
│  If file matches: internal/*/domain/*.go                    │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│             Validation Script Execution                     │
│  scripts/validate-domain-purity.sh                          │
│  - Checks imports                                           │
│  - Validates only stdlib + shared                           │
│  - Reports violations                                       │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                Agent Activation (Optional)                  │
│  Architecture Reviewer Agent                                │
│  - Reviews the change                                       │
│  - Provides feedback                                        │
└────────────────────────┬────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│              User continues development                     │
└─────────────────────────────────────────────────────────────┘
```

## Data Flow

### 1. Skill Invocation Flow

```
User → /skill-name → SKILL.md instructions → Claude execution → Generated code
                                                                      ↓
                                                              PostToolUse Hook
                                                                      ↓
                                                              Validation Script
                                                                      ↓
                                                                 User feedback
```

### 2. Architecture Review Flow

```
User → /review-arch → Execute goarchtest → Parse results
                            ↓                      ↓
                    AI code analysis ←────────────┘
                            ↓
                    Combined report → User
```

### 3. MCP Server Integration

```
Claude needs architecture info
         ↓
    MCP Protocol
         ↓
goarchtest-server (Node.js)
         ↓
Execute goarchtest programmatically
         ↓
Return results to Claude
         ↓
Claude uses info in response
```

## Key Design Decisions

### 1. Domain-First Organization

**Decision:** Organize by domain/bounded context first, then hexagonal layers within each domain.

**Rationale:**
- Aligns with DDD bounded context principle
- Each domain can evolve independently
- Clear microservice boundaries
- Easier to understand business capabilities

**Alternative considered:** Layer-first organization (all domains in one `domain/` folder)
**Why rejected:** Harder to maintain domain isolation, less clear microservice boundaries

### 2. Cobra CLI for Microservices

**Decision:** Single binary with Cobra commands for each microservice.

**Rationale:**
- Developer experience: one build, multiple run options
- Deployment flexibility: same artifact, different commands
- Reduced maintenance: single codebase
- Shared code easily accessible

**Alternative considered:** Separate binaries per microservice
**Why rejected:** More build complexity, harder to share code

### 3. Automated Architecture Testing

**Decision:** Use goarchtest library for automated constraint validation.

**Rationale:**
- Catch violations in CI/CD before code review
- Objective, automated enforcement
- No manual checking needed
- Fast feedback loop

**Alternative considered:** Manual code review only
**Why rejected:** Human error, inconsistent enforcement, slower feedback

### 4. Hook-Based Validation

**Decision:** Use PostToolUse hooks for real-time validation.

**Rationale:**
- Immediate feedback to developers
- Prevent violations at creation time
- Educational (explains why violations are bad)
- No waiting for CI/CD

**Alternative considered:** Only CI/CD validation
**Why rejected:** Slower feedback, more fix cycles

### 5. Intelligent Agents vs Simple Scripts

**Decision:** Use dedicated agent files with rich context.

**Rationale:**
- Proactive guidance, not reactive fixes
- Educational feedback
- Context-aware recommendations
- Consistent voice and approach

**Alternative considered:** Simple linting scripts
**Why rejected:** Less helpful, no explanations, reactive only

### 6. MCP Server for Architecture Analysis

**Decision:** Build Node.js MCP server wrapping goarchtest.

**Rationale:**
- Real-time analysis during development
- Claude can check architecture before suggesting changes
- Programmatic access to goarchtest
- Interactive dependency graphs

**Alternative considered:** Only CLI-based goarchtest
**Why rejected:** Not integrated into Claude's workflow

## Extension Points

The plugin is designed to be extensible:

### 1. New Skills

Add new skills by creating `skills/[skill-name]/SKILL.md`:
```markdown
---
description: Brief description
---

Instructions for Claude to follow...
```

Register in `plugin.json`:
```json
{
  "skills": [
    "./skills/new-skill/"
  ]
}
```

### 2. New Hooks

Add hooks in `plugin.json`:
```json
{
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write|Edit",
        "matchPath": "pattern",
        "hooks": [
          {
            "type": "command",
            "command": "script.sh"
          }
        ]
      }
    ]
  }
}
```

### 3. New MCP Tools

Extend `servers/goarchtest-server`:
```javascript
{
  name: "new_tool",
  description: "...",
  inputSchema: { ... }
}
```

### 4. New Agents

Create `agents/new-agent.md` with agent instructions and register in `plugin.json`.

## Performance Considerations

### 1. Hook Execution Time

**Concern:** Hooks run after every file edit/write.

**Mitigation:**
- Hooks only match specific paths (regex patterns)
- Scripts exit early if not applicable
- Validation is fast (imports parsing only)

### 2. Architecture Test Execution

**Concern:** Full test suite can be slow on large projects.

**Mitigation:**
- Tests are opt-in (not auto-triggered except on test file edits)
- MCP server can run targeted tests (specific layer/domain)
- Tests are cached by Go test framework

### 3. MCP Server Startup

**Concern:** Node.js server needs to start.

**Mitigation:**
- Server starts once per Claude session
- Stays running for duration of session
- Lightweight implementation (minimal dependencies)

## Security Considerations

### 1. Script Execution

**Risk:** PostToolUse hooks execute shell scripts.

**Mitigation:**
- Scripts are part of the plugin (trusted source)
- No user input in script execution
- Scripts are read-only (no file modifications)
- Limited to validation only

### 2. MCP Server

**Risk:** MCP server executes commands.

**Mitigation:**
- Server only runs goarchtest (read-only analysis)
- No file system modifications
- Sandboxed to project directory
- Stdio communication only (no network)

### 3. Generated Code

**Risk:** Plugin generates Go code in user's project.

**Mitigation:**
- Code follows best practices
- No hardcoded secrets or credentials
- User reviews all generated code
- Architecture tests validate generated code

## Testing Strategy

The plugin itself should be tested:

1. **Skill Testing**: Verify each skill generates correct structure
2. **Hook Testing**: Ensure hooks trigger correctly and validate properly
3. **Agent Testing**: Validate agent recommendations are accurate
4. **MCP Server Testing**: Test each MCP tool endpoint
5. **Integration Testing**: End-to-end workflow testing

## Future Enhancements

### Short Term
- [ ] More domain suggestions (20+ project types)
- [ ] GraphQL protocol support
- [ ] Event-driven architecture patterns
- [ ] CQRS/Event Sourcing templates

### Medium Term
- [ ] Visual architecture diagrams (auto-generated)
- [ ] Migration guides (from existing projects)
- [ ] Performance testing templates
- [ ] API documentation generation

### Long Term
- [ ] AI-powered domain discovery (analyze existing code)
- [ ] Automatic refactoring suggestions
- [ ] Compliance checking (GDPR, SOC2, etc.)
- [ ] Multi-language support (TypeScript, Rust)

## Contributing

To contribute to this plugin:

1. Understand the architecture (this document)
2. Follow the structure conventions
3. Add tests for new features
4. Update documentation
5. Submit PR with clear description

See individual component README files for specific contribution guidelines.

## Resources

- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)
- [Domain-Driven Design](https://www.domainlanguage.com/ddd/)
- [goarchtest](https://github.com/solrac97gr/goarchtest)
- [MCP Protocol](https://modelcontextprotocol.io/)
- [Claude Code Plugins](https://docs.anthropic.com/claude/docs/claude-code)
