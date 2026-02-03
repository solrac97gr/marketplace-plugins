---
description: Analyze project to generate custom plugin
---

Analyze an existing project's codebase, structure, and documentation to understand its patterns and suggest a custom Claude Code plugin tailored to the project.

## Objective

Create a comprehensive analysis of a project to generate:
1. Custom skills that match the project's workflows
2. Intelligent agents for code review and guidance
3. Validation hooks for code quality
4. MCP servers for project-specific tooling (if needed)

## Analysis Process

### Step 1: Project Discovery

Ask the user:
1. **Project path**: Where is the project located?
2. **Project type**: What kind of project is it?
   - Web application (React, Vue, Angular, etc.)
   - Backend API (Go, Node.js, Python, etc.)
   - Mobile app (React Native, Flutter, etc.)
   - Library/SDK
   - Microservices
   - Other (specify)
3. **Main technologies**: What are the core technologies used?
4. **Development workflows**: What are the common tasks developers perform?
   - Creating new features/components
   - Testing (unit, integration, e2e)
   - Deployment
   - Code review
   - Documentation
5. **Pain points**: What repetitive tasks would benefit from automation?

### Step 2: Codebase Analysis

**IMPORTANT**: Use the Task tool with subagent_type=Explore to perform thorough codebase exploration.

Analyze the following aspects:

#### Project Structure
- Directory layout and organization patterns
- File naming conventions
- Module/package structure
- Configuration files present

#### Code Patterns
- Architectural patterns (MVC, hexagonal, clean architecture, etc.)
- Design patterns frequently used
- Common abstractions (repositories, services, controllers, etc.)
- Testing patterns and frameworks
- Error handling patterns

#### Technology Stack
- Programming languages
- Frameworks and libraries
- Build tools and package managers
- Testing frameworks
- CI/CD tools
- Database systems

#### Documentation
- README.md patterns
- API documentation
- Code comments style
- Architecture decision records (ADRs)
- Contributing guidelines

#### Existing Automation
- Scripts in package.json, Makefile, or similar
- Git hooks
- CI/CD pipelines
- Code generation tools

### Step 3: Pattern Recognition

Identify common patterns that can be automated:

1. **Scaffolding Patterns**
   - New feature structure
   - New component/module structure
   - Test file patterns
   - Configuration file patterns

2. **Code Quality Patterns**
   - Import rules (what can import what)
   - Naming conventions
   - Required file presence (e.g., tests for every feature)
   - Architectural boundaries

3. **Workflow Patterns**
   - Feature development workflow (TDD, BDD, etc.)
   - Code review checklist items
   - Deployment prerequisites
   - Documentation requirements

4. **Project-Specific Tools**
   - Custom CLI commands
   - Build tools
   - Code generators
   - Analysis tools

### Step 4: Generate Analysis Report

Create a comprehensive report with:

```markdown
# Project Analysis Report

## Project Overview
- **Name**: [project-name]
- **Type**: [project-type]
- **Main Language**: [language]
- **Framework**: [framework]

## Structure Analysis
[Describe the directory structure and organization]

## Code Patterns Detected
[List architectural and design patterns]

## Technology Stack
- **Languages**: [list]
- **Frameworks**: [list]
- **Testing**: [frameworks]
- **Build Tools**: [tools]

## Suggested Plugin Capabilities

### Skills (Actions users can invoke)
1. **[skill-name]**: [description]
   - **Purpose**: [what it does]
   - **Template**: [what it generates]
   - **Example**: `/[skill-name] [args]`

2. [more skills...]

### Agents (Intelligent reviewers)
1. **[agent-name]**: [description]
   - **Expertise**: [domain]
   - **Triggers**: [when it activates]
   - **Checks**: [what it validates]

2. [more agents...]

### Hooks (Automatic validations)
1. **[hook-type]**: [description]
   - **Trigger**: [file pattern]
   - **Validates**: [what to check]
   - **Script**: [validation logic]

2. [more hooks...]

### MCP Server (Optional)
- **Needed**: Yes/No
- **Purpose**: [what analysis/tools it provides]
- **Tools**: [list of tools]

## Recommended Plugin Name
**[suggested-plugin-name]**

## Next Steps
To generate the plugin, run:
`/generate-plugin`
```

## Important Guidelines

1. **Be thorough**: Analyze at least 50+ files to understand patterns
2. **Look for repetition**: What structures appear multiple times?
3. **Check documentation**: README, CONTRIBUTING, docs/ folder
4. **Analyze tests**: Test patterns reveal development workflows
5. **Check configs**: package.json, Makefile, etc. show common commands
6. **Consider scale**: How many developers? How big is the codebase?

## Example Analysis

For a **React + TypeScript** project with **feature-based architecture**:

```
Detected Patterns:
- Features in src/features/[feature-name]/
- Each feature has: components/, hooks/, services/, types/
- Tests colocated with source files
- Storybook stories for components
- React Query for data fetching

Suggested Skills:
1. new-feature: Scaffold complete feature with all subdirectories
2. new-component: Create component + test + story
3. new-hook: Create custom hook with tests
4. new-service: Create API service with types

Suggested Agents:
1. React Best Practices Enforcer
2. TypeScript Type Safety Reviewer
3. Test Coverage Guardian

Suggested Hooks:
1. Validate component has test file
2. Check for proper TypeScript types (no 'any')
3. Ensure React hooks follow Rules of Hooks
```

## Output Format

Present the analysis in a clear, structured format. Ask the user if they want to proceed with generating the plugin based on this analysis.

If the user approves, they can run `/generate-plugin` to create the actual plugin files.
