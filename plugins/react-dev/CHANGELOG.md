# Changelog

All notable changes to the react-dev plugin will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-02-01

### Added

#### Skills
- `/start-project` - Initialize React + Vite + TypeScript projects with feature-based architecture
- `/new-feature` - Scaffold features with TDD workflow
- `/new-component` - Create components with tests and Storybook stories
- `/new-hook` - Create custom hooks with tests
- `/new-context` - Create Context providers with useReducer
- `/from-diagram` - Convert diagrams (Mermaid, PlantUML, GraphViz) to React components
- `/review-code` - Comprehensive React code quality review
- `/review-a11y` - Accessibility compliance review (WCAG 2.1 AA)
- `/optimize-performance` - React performance analysis and optimization
- `/update-tests` - Update test coverage for existing code

#### Intelligent Agents
- Component Architecture Reviewer - Reviews component structure and composition
- React Best Practices Enforcer - Enforces React 18+ patterns
- Accessibility Reviewer - WCAG 2.1 AA compliance
- Performance Reviewer - React.memo, useMemo, useCallback optimization
- Security Reviewer - XSS prevention, input sanitization
- Test Coverage Guardian - Testing best practices and coverage tracking
- State Management Reviewer - Context patterns and prop drilling prevention
- Hook Usage Reviewer - Custom hooks and rules of hooks compliance
- TypeScript Quality Reviewer - Strict mode enforcement
- Tailwind Best Practices Reviewer - Utility-first CSS patterns

#### MCP Servers
- Playwright Server - E2E testing tools (run tests, check flows, generate reports)
- Component Analyzer - Component analysis tools (tree analysis, prop drilling, bundle size)

#### Automation
- PreToolUse hooks for commit validation
- PostToolUse hooks for automatic code review on file changes
- Validation scripts for component purity, accessibility, bundle size
- Template system for consistent code generation

#### Documentation
- Comprehensive CODE_STANDARDS.md
- ARCHITECTURE.md with feature-based patterns
- README.md with quick start guide
- Complete TypeScript, React, and Tailwind standards
- Testing strategy and best practices

### Features

- **Feature-Based Architecture**: Organize by features, not technical layers
- **TypeScript First**: Strict mode, comprehensive type safety
- **Modern Testing Stack**: Vitest + React Testing Library + Playwright + Storybook
- **TDD Workflow**: Test-first development with automated coverage tracking
- **Accessibility Built-In**: WCAG 2.1 AA compliance validation
- **Performance Monitoring**: Automated optimization recommendations
- **Tailwind CSS Integration**: Utility-first, mobile-first responsive design
- **Dark Mode Support**: Consistent theming patterns
- **Code Quality**: 10 specialized review agents
- **Automated Validation**: Pre-commit and post-edit hooks

### Dependencies

- React 18.3+
- TypeScript 5.5+
- Vite 5.3+
- Vitest 2.0+
- React Testing Library 16.0+
- Playwright 1.45+
- Tailwind CSS 3.4+
- Storybook 8.1+

[1.0.0]: https://github.com/solrac97gr/marketplace-plugins/releases/tag/react-dev-v1.0.0
