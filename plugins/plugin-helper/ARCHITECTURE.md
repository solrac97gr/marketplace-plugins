# Plugin Helper Architecture

## Overview

Plugin Helper is a meta-plugin that analyzes existing codebases and generates custom Claude Code plugins tailored to the project's specific patterns and workflows.

## Core Concept

Instead of manually creating plugins, Plugin Helper:
1. **Analyzes** your project to understand its patterns
2. **Identifies** automation opportunities
3. **Generates** a complete custom plugin
4. **Adapts** to your project's unique needs

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                     Plugin Helper                            │
├─────────────────────────────────────────────────────────────┤
│                                                              │
│  ┌───────────────────┐         ┌──────────────────────┐    │
│  │ analyze-project   │────────▶│  Plugin Architect    │    │
│  │     Skill         │         │      Agent           │    │
│  └───────────────────┘         └──────────────────────┘    │
│           │                              │                  │
│           │                              │                  │
│           ▼                              ▼                  │
│  ┌─────────────────────────────────────────────────┐       │
│  │         Project Analysis Engine                  │       │
│  ├─────────────────────────────────────────────────┤       │
│  │ • Structure Analysis                             │       │
│  │ • Pattern Recognition                            │       │
│  │ • Convention Detection                           │       │
│  │ • Workflow Identification                        │       │
│  └─────────────────────────────────────────────────┘       │
│           │                                                  │
│           ▼                                                  │
│  ┌─────────────────────────────────────────────────┐       │
│  │         Analysis Report Generator                │       │
│  └─────────────────────────────────────────────────┘       │
│           │                                                  │
│           │                                                  │
│  ┌───────────────────┐                                      │
│  │ generate-plugin   │                                      │
│  │     Skill         │                                      │
│  └───────────────────┘                                      │
│           │                                                  │
│           ▼                                                  │
│  ┌─────────────────────────────────────────────────┐       │
│  │         Plugin Generator Engine                  │       │
│  ├─────────────────────────────────────────────────┤       │
│  │ • Skill Generator                                │       │
│  │ • Agent Generator                                │       │
│  │ • Hook Generator                                 │       │
│  │ • MCP Server Generator (Go)                      │       │
│  │ • Documentation Generator                        │       │
│  └─────────────────────────────────────────────────┘       │
│           │                                                  │
│           ▼                                                  │
│  ┌─────────────────────────────────────────────────┐       │
│  │         Generated Custom Plugin                  │       │
│  ├─────────────────────────────────────────────────┤       │
│  │ plugins/[custom-plugin-name]/                    │       │
│  │ ├── .claude-plugin/plugin.json                   │       │
│  │ ├── skills/                                      │       │
│  │ ├── agents/                                      │       │
│  │ ├── scripts/                                     │       │
│  │ ├── servers/ (optional)                          │       │
│  │ └── documentation                                │       │
│  └─────────────────────────────────────────────────┘       │
│                                                              │
└─────────────────────────────────────────────────────────────┘
```

## Components

### 1. Skills

#### analyze-project
**Purpose**: Analyze a codebase to understand its patterns and workflows

**Process**:
1. **Discovery**: Ask user about project type and context
2. **Exploration**: Use Task tool with Explore agent to scan codebase
3. **Pattern Recognition**: Identify repeated structures and conventions
4. **Opportunity Identification**: Find automation opportunities
5. **Report Generation**: Create comprehensive analysis report

**Output**: Detailed analysis report with plugin recommendations

#### generate-plugin
**Purpose**: Generate a complete custom plugin from analysis

**Process**:
1. **Configuration**: Confirm plugin details with user
2. **Structure Creation**: Create plugin directory structure
3. **Metadata Generation**: Create plugin.json with hooks and servers
4. **Skill Generation**: Create SKILL.md files with project-specific templates
5. **Agent Generation**: Create agent.md files with project context
6. **Hook Generation**: Create validation scripts in bash
7. **Server Generation**: Create Go MCP server (if needed)
8. **Documentation**: Create README, ARCHITECTURE, etc.

**Output**: Complete, ready-to-use custom plugin

### 2. Agents

#### Plugin Architect
**Expertise**: Codebase analysis and plugin design

**Responsibilities**:
- Guide project analysis process
- Identify patterns and opportunities
- Design plugin capabilities
- Ensure generated plugins follow best practices
- Provide implementation recommendations

**Activation**:
- Automatically during `/analyze-project`
- Automatically during `/generate-plugin`
- When user asks about plugin creation

### 3. Templates Directory

Stores reusable templates for generating plugins:
- Skill template structure
- Agent template structure
- Hook script templates
- MCP server templates
- Documentation templates

## Analysis Process

### Phase 1: Static Analysis

**Directory Structure**:
```
Analyze:
- Root directory layout
- Nested directory patterns
- File organization principles
- Configuration file locations
```

**File Patterns**:
```
Detect:
- Naming conventions (camelCase, snake_case, etc.)
- File extensions and types
- Required file presence patterns
- File relationships (test colocated, etc.)
```

**Code Patterns**:
```
Identify:
- Import/dependency patterns
- Architectural boundaries
- Abstraction layers
- Design patterns in use
- Testing approaches
```

### Phase 2: Convention Detection

**Naming Conventions**:
- Class/function naming style
- Variable naming patterns
- File naming rules
- Module/package naming

**Structural Conventions**:
- Feature/module organization
- Layer separation (if any)
- Testing structure
- Documentation location

**Quality Standards**:
- Code formatting style
- Comment patterns
- Error handling approach
- Logging conventions

### Phase 3: Workflow Analysis

**Development Tasks**:
- Feature creation process
- Component scaffolding needs
- Test creation workflow
- Documentation requirements

**Automation Opportunities**:
- Repetitive file creation
- Boilerplate code patterns
- Validation that could be automated
- Common multi-step processes

**Integration Points**:
- Build tools
- CI/CD pipelines
- Testing frameworks
- External tools

## Generation Process

### Phase 1: Design

**Capability Selection**:
- Which skills to generate
- Which agents to create
- Which hooks to implement
- Whether MCP server is needed

**Priority Assignment**:
- High-value capabilities first
- Quick wins vs. complex features
- Must-have vs. nice-to-have

### Phase 2: Implementation

**Skill Generation**:
```markdown
For each skill:
1. Extract template from actual project files
2. Create parameterized version
3. Add project-specific instructions
4. Include examples from project
5. Document usage
```

**Agent Generation**:
```markdown
For each agent:
1. Define domain expertise (from project)
2. List activation triggers
3. Specify validation rules (from project conventions)
4. Provide feedback templates
5. Reference project documentation
```

**Hook Generation**:
```bash
For each hook:
1. Define trigger pattern (file regex)
2. Implement validation logic
3. Use project-specific rules
4. Provide clear error messages
5. Make executable
```

**MCP Server Generation** (if needed):
```go
1. Define server purpose
2. Create tool interfaces
3. Implement analysis logic
4. Handle errors gracefully
5. Document all tools
6. Build binary
```

### Phase 3: Documentation

**README.md**:
- Plugin overview
- Available skills and usage
- Agent descriptions
- Installation instructions
- Examples from the project

**ARCHITECTURE.md**:
- Plugin structure explanation
- Design decisions
- Component descriptions
- Integration points

## Design Principles

### 1. Project-Specific Over Generic

❌ Bad (Generic):
```markdown
Create a component with these files:
- Component.tsx
- Component.test.tsx
- Component.css
```

✅ Good (Project-Specific):
```markdown
Create a component following [ProjectName]'s conventions:
- src/features/[feature]/components/[Name]/[Name].tsx
- src/features/[feature]/components/[Name]/[Name].test.tsx
- src/features/[feature]/components/[Name]/[Name].stories.tsx
- Uses Tailwind CSS (no .css files)
- Exports [Name]Props interface
- Includes accessibility attributes
```

### 2. Extract, Don't Invent

Always base templates on actual project files:
```
1. Find exemplar files in project
2. Extract pattern
3. Parameterize
4. Validate against other examples
5. Use in skill template
```

### 3. Educational, Not Just Enforcement

Agents should teach, not just check:
```markdown
❌ "Naming convention violated"
✅ "Function names should use camelCase (e.g., 'getUserData')
   as specified in [project]/CONTRIBUTING.md section 3.2"
```

### 4. Fail Fast, Explain Clearly

Hooks should provide actionable feedback:
```bash
❌ "Validation failed"
✅ "Domain layer imports 'database/sql' package.
   Domain layer should only import stdlib and internal/shared.
   Move database logic to infrastructure layer.
   See: docs/architecture.md"
```

## Integration with Other Plugins

Plugin Helper complements existing plugins:

```
go-dev → analyze Go project → generate go-[project]-dev
react-dev → analyze React project → generate react-[project]-dev
plugin-helper → analyze any project → generate custom plugin
```

Generated plugins can coexist with generic plugins:
- Use generic plugin for standard tasks
- Use custom plugin for project-specific tasks

## Maintenance

### Updating Generated Plugins

When project evolves:
1. Re-run `/analyze-project` to detect new patterns
2. Compare with existing plugin
3. Add new capabilities
4. Update hooks if rules changed
5. Regenerate or manually update

### Versioning

Generated plugins should be versioned:
```json
{
  "version": "1.0.0",  // Initial generation
  "version": "1.1.0",  // Added new skill
  "version": "1.2.0",  // Added new agent
  "version": "2.0.0"   // Major restructure
}
```

### Testing

Test generated plugins:
1. **Smoke test**: Can skills run without errors?
2. **Template test**: Do generated files match project style?
3. **Hook test**: Do validations catch actual violations?
4. **Agent test**: Do agents provide useful feedback?
5. **Integration test**: Does plugin work with real project?

## Future Enhancements

Potential improvements:
- [ ] Plugin update detection (suggest updates when project changes)
- [ ] Plugin marketplace submission helper
- [ ] Multi-language project support (monorepos)
- [ ] Visual plugin builder (UI instead of CLI)
- [ ] Plugin testing framework
- [ ] Plugin analytics (which skills are most used?)
- [ ] Community plugin templates
- [ ] Plugin inheritance (extend base plugins)

## Technical Decisions

### Why Skills Over Single Monolithic Command?

✅ Skills are granular and composable
✅ Users can invoke specific capabilities
✅ Easier to maintain and extend
✅ Better discoverability

### Why Go for MCP Servers?

✅ Consistency with go-dev and react-dev
✅ Single binary distribution
✅ Fast performance
✅ Strong typing
✅ Excellent standard library

### Why Bash for Hooks?

✅ Universal availability
✅ Simple for file validation
✅ Fast execution
✅ Easy to debug
✅ No dependencies

### Why Extract Templates?

✅ Ensures generated code matches project
✅ Captures real patterns, not idealized ones
✅ Includes project-specific details
✅ Validates against actual examples

## Security Considerations

### Analysis Phase
- Only analyze files in project directory
- Don't execute arbitrary code
- Don't expose sensitive data in reports

### Generation Phase
- Don't include secrets in generated files
- Validate all user inputs
- Use project's .gitignore patterns

### Execution Phase
- Hooks should fail safe (not block critical work)
- MCP servers should sanitize inputs
- Agents should not modify files without permission

## Performance

### Analysis
- Use parallel file reading where possible
- Cache analysis results
- Limit directory depth for large projects
- Use streaming for large files

### Generation
- Generate files incrementally
- Show progress for long operations
- Allow cancellation

### Runtime (Generated Plugin)
- Hooks should complete in < 1 second
- Skills should show progress for long tasks
- MCP servers should have reasonable timeouts

## Conclusion

Plugin Helper enables rapid creation of high-quality, project-specific plugins by:
1. Understanding your project deeply
2. Identifying automation opportunities
3. Generating tailored capabilities
4. Maintaining consistency with your patterns

The result: plugins that feel native to your project because they're built from your project's DNA.
