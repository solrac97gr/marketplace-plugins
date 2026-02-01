# React-Dev Plugin

Comprehensive React development plugin for Claude Code with TypeScript, Vite, Tailwind CSS, and modern testing tools.

## Features

- **Feature-Based Architecture**: Organize code by features, not technical layers
- **TypeScript First**: Strict mode, comprehensive type safety
- **Modern Testing**: Vitest + React Testing Library + Playwright + Storybook
- **TDD Workflow**: Test-first development with automated coverage tracking
- **Accessibility**: WCAG 2.1 AA compliance built-in
- **Performance**: Automated optimization recommendations
- **Intelligent Agents**: 10 specialized code review agents
- **MCP Integration**: Playwright and Component Analyzer servers

## Quick Start

```bash
# Initialize a new React project
/start-project

# Create a new feature with TDD workflow
/new-feature

# Generate a component with tests and stories
/new-component

# Create a custom hook with tests
/new-hook

# Add context-based state management
/new-context
```

## Skills

### Project Initialization
- **`/start-project`** - Initialize React + Vite + TypeScript project with feature-based structure

### Code Generation
- **`/new-feature`** - Scaffold complete feature with TDD workflow
- **`/new-component`** - Create component with tests and Storybook story
- **`/new-hook`** - Create custom hook with tests
- **`/new-context`** - Create Context provider with useReducer

### Code Review & Optimization
- **`/review-code`** - Comprehensive React code quality review
- **`/review-a11y`** - Accessibility compliance review (WCAG 2.1 AA)
- **`/optimize-performance`** - Analyze and optimize React performance
- **`/update-tests`** - Update test coverage for existing code

### Advanced
- **`/from-diagram`** - Convert diagrams (Mermaid, PlantUML, GraphViz) to React components

## Intelligent Agents

The plugin includes 10 specialized agents that automatically review your code:

1. **Component Architecture Reviewer** - Reviews component structure and composition
2. **React Best Practices Enforcer** - Enforces React 18+ patterns and hooks rules
3. **Accessibility Reviewer** - WCAG 2.1 AA compliance validation
4. **Performance Reviewer** - React.memo, useMemo, useCallback optimization
5. **Security Reviewer** - XSS prevention, input sanitization, secure storage
6. **Test Coverage Guardian** - Testing best practices and coverage tracking
7. **State Management Reviewer** - Context patterns and prop drilling prevention
8. **Hook Usage Reviewer** - Custom hooks and rules of hooks compliance
9. **TypeScript Quality Reviewer** - Strict mode, typing patterns, no any
10. **Tailwind Best Practices Reviewer** - Utility classes, responsive design

## MCP Servers

### Playwright Server
Tools for E2E testing:
- `run_e2e_tests` - Execute Playwright tests
- `check_critical_flows` - Validate user flows
- `generate_test_report` - Create HTML report
- `check_browser_compatibility` - Cross-browser testing
- `capture_screenshots` - Visual regression baseline

### Component Analyzer
Tools for component analysis:
- `analyze_component_tree` - Generate component hierarchy
- `check_prop_drilling` - Identify excessive prop passing
- `find_unused_components` - Detect unused components
- `analyze_bundle_size` - Per-component bundle impact
- `check_circular_dependencies` - Detect circular imports

## Project Structure

```
project-name/
├── src/
│   ├── features/                          # Feature-based organization
│   │   ├── [feature-name]/
│   │   │   ├── components/                # Feature components
│   │   │   ├── hooks/                     # Feature hooks
│   │   │   ├── context/                   # Feature state
│   │   │   ├── services/                  # API & business logic
│   │   │   ├── types/                     # Feature types
│   │   │   └── utils/                     # Feature utilities
│   │   └── shared/                        # Shared code
│   ├── app/                               # App-level configuration
│   ├── assets/                            # Static assets
│   ├── styles/                            # Global styles
│   └── main.tsx                           # Entry point
├── tests/
│   ├── e2e/                               # Playwright E2E
│   └── integration/                       # Integration tests
├── .storybook/                            # Storybook config
└── public/                                # Public assets
```

## Code Standards

The plugin enforces comprehensive code standards:

- **React**: Functional components, React 18+ patterns, hooks rules
- **TypeScript**: Strict mode, no `any`, proper typing
- **Testing**: User-centric tests, 80%+ coverage, accessibility tests
- **Accessibility**: WCAG 2.1 AA, semantic HTML, keyboard navigation
- **Performance**: Memoization, code splitting, bundle optimization
- **Tailwind**: Utility-first, mobile-first responsive design

See [CODE_STANDARDS.md](./CODE_STANDARDS.md) for complete details.

## Architecture

Feature-based organization ensures:
- **Feature Isolation**: Each feature is self-contained
- **Clear Dependencies**: Features expose public APIs via index.ts
- **Testability**: Easy to test features in isolation
- **Scalability**: Add features without affecting others

See [ARCHITECTURE.md](./ARCHITECTURE.md) for complete architecture guide.

## Automation

### Pre-Commit Hooks
- Test coverage validation
- TypeScript quality checks

### Post-Edit Hooks
- Component architecture review
- React best practices enforcement
- Accessibility compliance
- Performance optimization suggestions
- Security vulnerability detection

### Validation Scripts
- Component purity validation
- Accessibility automated checks
- Bundle size analysis
- Unused dependency detection

## Installation

The plugin is automatically available when installed in the Claude Code marketplace.

### MCP Server Dependencies

To use MCP server tools, install dependencies:

```bash
cd plugins/react-dev/servers/playwright-server
npm install

cd ../component-analyzer
npm install
```

## Version History

See [CHANGELOG.md](./CHANGELOG.md) for version history.

## License

MIT
