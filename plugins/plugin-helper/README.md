# Plugin Helper

A meta-plugin for analyzing projects and generating custom Claude Code plugins tailored to your codebase patterns and workflows.

## Overview

Plugin Helper enables you to automatically create custom Claude Code plugins based on your existing project's code, structure, and documentation. Instead of manually crafting skills and agents, let Plugin Helper analyze your codebase and generate a plugin that perfectly matches your development patterns.

## Use Cases

- **New Team Members**: Generate a plugin that captures your team's workflows and best practices
- **Consistency**: Automate code generation to ensure consistency across the codebase
- **Knowledge Capture**: Turn tribal knowledge into automated skills and validations
- **Productivity**: Eliminate repetitive scaffolding and boilerplate tasks
- **Quality**: Enforce architectural rules and code standards automatically

## Features

### 1. Project Analysis
Deeply analyze your codebase to identify:
- Directory structure and organization patterns
- File naming conventions
- Architectural patterns (hexagonal, clean architecture, etc.)
- Testing approaches and patterns
- Common workflows and tasks
- Code quality rules and standards

### 2. Intelligent Plugin Generation
Automatically generate:
- **Skills**: User-invocable commands for common tasks
- **Agents**: Intelligent code reviewers for quality enforcement
- **Hooks**: Automatic validations on file changes
- **MCP Servers**: Advanced tooling (optional, in Go)
- **Documentation**: Complete README and architecture docs

### 3. Project-Specific Templates
All generated code uses your project's actual patterns:
- Extracted from existing code
- Matches your naming conventions
- Follows your architectural rules
- Uses your preferred testing approach
- References your documentation

## Skills

### `/analyze-project`

Analyzes your project to understand its patterns and suggest plugin capabilities.

**What it does:**
1. Asks about your project (type, technologies, workflows)
2. Explores the codebase using intelligent agents
3. Identifies patterns and opportunities for automation
4. Generates a comprehensive analysis report
5. Suggests skills, agents, hooks, and MCP servers

**Example:**
```bash
/analyze-project

# You'll be asked:
# - Project path
# - Project type (web app, API, etc.)
# - Main technologies
# - Common workflows
# - Pain points to automate
```

**Output:**
- Detailed analysis report
- List of suggested capabilities
- Implementation recommendations
- Estimated time savings

### `/generate-plugin`

Generates a complete custom plugin based on the analysis.

**Prerequisites:** Must run `/analyze-project` first

**What it creates:**
- Complete plugin structure
- plugin.json with metadata and hooks
- SKILL.md files for each capability
- Agent.md files for intelligent reviewers
- Validation scripts (bash)
- MCP server (if needed, in Go)
- Comprehensive documentation

**Example:**
```bash
/generate-plugin

# You'll be asked to confirm:
# - Plugin name
# - Which skills to include
# - Which agents to include
# - Which hooks to include
# - MCP server creation
```

**Output:**
- `plugins/[your-plugin-name]/` with complete structure
- Updated marketplace.json
- Usage instructions
- Testing guidelines

## Agents

### Plugin Architect

Expert agent that helps design high-quality plugins.

**Expertise:**
- Codebase pattern recognition
- Plugin architecture design
- Capability prioritization
- Best practices enforcement

**When it activates:**
- During `/analyze-project` execution
- During `/generate-plugin` execution
- When you ask about creating plugins
- When you mention automation needs

**What it provides:**
- Pattern analysis and insights
- Capability suggestions
- Implementation guidance
- Quality assessment

## How It Works

```
1. Run /analyze-project
   ↓
2. Plugin Architect analyzes codebase
   ↓
3. Identifies patterns and opportunities
   ↓
4. Generates analysis report
   ↓
5. User reviews and approves
   ↓
6. Run /generate-plugin
   ↓
7. Generates complete plugin structure
   ↓
8. Plugin ready to use!
```

## Example Workflow

### Step 1: Analyze Your React Project

```bash
/analyze-project
```

**Input:**
- Project path: `/Users/you/my-react-app`
- Project type: Web application
- Technologies: React, TypeScript, Vite, Tailwind
- Workflows: Feature development, component creation, testing
- Pain points: Repetitive component scaffolding

**Analysis Result:**
```
Detected Patterns:
- Features in src/features/[feature-name]/
- Each feature has: components/, hooks/, services/, types/
- Tests colocated with source files
- Storybook stories for all components

Suggested Plugin: my-react-app-dev

Skills:
1. new-feature: Scaffold complete feature
2. new-component: Component + test + story
3. new-hook: Custom hook with tests
4. new-service: API service with types

Agents:
1. React Best Practices Enforcer
2. TypeScript Type Safety Reviewer
3. Component Architecture Reviewer

Hooks:
1. Validate component has test
2. Check for TypeScript strict types
3. Ensure components have stories

Time Saved: ~15-20 minutes per feature
```

### Step 2: Generate the Plugin

```bash
/generate-plugin
```

**Result:**
```
✅ Plugin Generated: my-react-app-dev

Location: plugins/my-react-app-dev/

Created:
- plugin.json
- 4 skills (new-feature, new-component, new-hook, new-service)
- 3 agents (React Best Practices, TypeScript Quality, Architecture)
- 3 validation hooks
- Complete documentation

Next Steps:
1. Test with: /new-feature user-profile
2. Review generated files
3. Customize as needed
```

### Step 3: Use Your Custom Plugin

```bash
/new-feature shopping-cart

# Creates:
# src/features/shopping-cart/
# ├── components/
# │   ├── ShoppingCart.tsx
# │   ├── ShoppingCart.test.tsx
# │   └── ShoppingCart.stories.tsx
# ├── hooks/
# │   └── useShoppingCart.ts
# ├── services/
# │   └── shoppingCartService.ts
# ├── types/
# │   └── shoppingCart.types.ts
# └── index.ts
```

## Generated Plugin Examples

### Example 1: Go Microservices

**Project Type:** Backend API with hexagonal architecture

**Generated Skills:**
- `/new-service [name]` - Scaffold new microservice
- `/new-endpoint [service] [route]` - Add REST endpoint
- `/new-domain-entity [service] [entity]` - Add domain entity

**Generated Agents:**
- Hexagonal Architecture Reviewer
- API Contract Validator
- DDD Consultant

**Generated Hooks:**
- Domain layer purity validation
- Architecture test runner

**MCP Server:**
- Architecture testing tools
- Dependency graph generator

### Example 2: Python Data Pipeline

**Project Type:** Airflow data processing

**Generated Skills:**
- `/new-dag [name]` - Create Airflow DAG
- `/new-operator [name]` - Custom operator
- `/new-transformation [dag] [step]` - Add data transformation

**Generated Agents:**
- DAG Complexity Reviewer
- Data Quality Validator
- Pipeline Best Practices

**Generated Hooks:**
- DAG validation (no circular deps)
- SQL query validation

**MCP Server:**
- DAG dependency analyzer
- Data lineage tracker

### Example 3: Mobile App (React Native)

**Project Type:** Cross-platform mobile app

**Generated Skills:**
- `/new-screen [name]` - Create screen with navigation
- `/new-component [name]` - Mobile component
- `/new-api-hook [endpoint]` - API integration hook

**Generated Agents:**
- Mobile UX Reviewer
- Performance Optimizer
- Accessibility Checker (mobile)

**Generated Hooks:**
- Platform-specific code validation
- Navigation structure validation

## Benefits

### Time Savings
- **Scaffolding**: 10-20 minutes per component/feature → automated
- **Validation**: Catch issues immediately, not in code review
- **Consistency**: No more copy-paste errors
- **Onboarding**: New developers productive faster

### Quality Improvements
- **Enforce standards**: Automatically check architectural rules
- **Catch mistakes**: Validate before commit
- **Capture knowledge**: Team patterns become automated
- **Prevent drift**: Keep codebase consistent

### Developer Experience
- **Less tedious work**: Focus on business logic
- **Faster feedback**: Issues caught immediately
- **Clear guidance**: Agents provide educational feedback
- **Confidence**: Know your code follows standards

## Best Practices

### When Analyzing Projects
1. **Be thorough**: Let the agent explore extensively
2. **Provide context**: Explain your workflows and pain points
3. **Share examples**: Point to exemplar files/features
4. **Be specific**: "We always do X" is more useful than "Sometimes we do X"

### When Generating Plugins
1. **Start focused**: Don't try to automate everything at once
2. **Test immediately**: Verify generated code works
3. **Iterate**: Start with 2-3 skills, add more later
4. **Document**: Keep examples and usage notes updated
5. **Maintain**: Update plugin as project evolves

### Plugin Naming
- Use project/team name: `my-app-dev`, `team-backend-tools`
- Be specific: `ecommerce-platform-dev` not `web-dev`
- Avoid generic names: `helpers`, `utils`, `common`

## Limitations

### What It Can Do
- ✅ Analyze most programming languages
- ✅ Detect common architectural patterns
- ✅ Generate project-specific templates
- ✅ Create validation rules
- ✅ Build MCP servers in Go

### What It Cannot Do
- ❌ Understand business logic (you provide this context)
- ❌ Read your mind (be explicit about patterns)
- ❌ Modify existing plugins (generates new ones)
- ❌ Generate perfect code (review and refine needed)

## Troubleshooting

### Analysis seems shallow
- Provide more context about your workflows
- Point to specific exemplar files
- Explain your architecture explicitly
- Share internal documentation

### Generated code doesn't match style
- Review the analysis report
- Check if patterns were correctly detected
- Provide specific examples of preferred style
- Regenerate with more guidance

### Too many/too few capabilities
- Prioritize by impact: which saves most time?
- Start small: 2-3 high-value skills
- Add more later as needed
- Focus on repetitive tasks first

## Advanced Usage

### Custom Templates
After generation, you can add custom templates to:
```
plugins/[your-plugin]/templates/
```

Reference these in your SKILL.md files.

### Multiple Plugins per Project
Create specialized plugins:
- `[project]-frontend-dev` - Frontend skills
- `[project]-backend-dev` - Backend skills
- `[project]-testing` - Testing utilities
- `[project]-deployment` - Deployment automation

### Evolution
Update your plugin as project evolves:
1. Re-run `/analyze-project` on updated codebase
2. Compare new analysis with existing plugin
3. Add new capabilities
4. Update hooks and agents
5. Version your plugin (update version in plugin.json)

## FAQ

**Q: How long does analysis take?**
A: 2-5 minutes for small projects, 5-15 minutes for large codebases.

**Q: Will it work with my language/framework?**
A: Yes, Plugin Helper is language-agnostic. It analyzes patterns, not just syntax.

**Q: Can I edit generated plugins?**
A: Absolutely! Generated plugins are fully editable. Customize as needed.

**Q: Do I need to know how plugins work?**
A: No! Plugin Helper handles the technical details. Just describe your workflows.

**Q: Can I share my generated plugin?**
A: Yes! Generated plugins can be shared with your team or published.

**Q: What if my project is unique?**
A: That's the point! Plugin Helper creates plugins specific to your unique patterns.

## Examples in This Marketplace

Both `go-dev` and `react-dev` plugins in this marketplace were created with similar analysis processes:
- Analyzed common Go and React patterns
- Identified scaffolding opportunities
- Created comprehensive agents
- Built MCP servers for advanced tooling

You can use these as references for plugin structure and quality.

## Contributing

Found a bug or have a suggestion? Please open an issue on GitHub.

## License

MIT
