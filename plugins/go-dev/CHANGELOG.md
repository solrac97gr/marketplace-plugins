# Changelog

All notable changes to the Go Hexagonal Architecture Development Plugin will be documented in this file.

## [1.0.0] - 2026-01-30

### Added

#### Core Skills
- **`/start-project`** - Initialize Go project with hexagonal architecture
  - Intelligent domain suggestions based on project type
  - Multi-select domain creation
  - Support for REST API, gRPC, or both
  - Multiple database support (PostgreSQL, MongoDB, MySQL, SQLite)
  - Cobra CLI with microservice commands
  - Automated architecture test generation
  - Godog BDD setup following v0.12.0+ best practices
  - Complete Makefile with all commands
  - Docker Compose configuration

- **`/new-feature`** - Complete feature scaffolding with TDD/BDD
  - Starts with Godog feature file creation
  - Uses modern Godog patterns (context.Context, TestSuite)
  - Interactive clarification questions
  - Domain entity generation
  - Use case creation
  - Infrastructure adapters (HTTP/gRPC)
  - Comprehensive test suite with proper step definitions

- **`/new-entity`** - DDD entity creation
  - Domain entities with business logic
  - Repository interfaces (ports)
  - Value objects
  - Unit tests

- **`/new-usecase`** - Application use case creation
  - Use case with workflow orchestration
  - Input/Output DTOs
  - Unit tests with mocks
  - TDD approach

- **`/review-arch`** - Architectural compliance review
  - Automated architecture testing with goarchtest
  - AI-powered code analysis
  - Combined reporting
  - Dependency violation detection
  - DDD pattern validation
  - Anti-pattern identification

- **`/update-arch-tests`** - Architecture test management
  - Automatic test generation based on project structure
  - Layer dependency tests
  - Domain isolation tests
  - Naming convention validation
  - Custom rule support

#### Intelligent Agents
- **Architecture Reviewer Agent**
  - Proactive code review
  - Dependency direction validation
  - Layer purity checks
  - DDD pattern enforcement
  - Educational feedback

- **DDD Consultant Agent**
  - Strategic DDD guidance
  - Bounded context identification
  - Tactical pattern recommendations
  - Entity vs Value Object decisions
  - Aggregate design assistance
  - Ubiquitous language enforcement

- **Domain Documentation Writer Agent**
  - Standardized domain documentation generation
  - Mermaid architecture diagrams
  - Entity and use case documentation
  - API endpoint listings
  - Database schema documentation

- **Test Coverage Guardian Agent**
  - TDD/BDD practice enforcement
  - Test coverage monitoring (90%+ domain, 85%+ application, 70%+ infrastructure)
  - Godog feature validation with modern patterns (context.Context, TestSuite)
  - Step definition best practices enforcement
  - Test quality review
  - Red-Green-Refactor workflow guidance
  - go test integration validation

- **API Contract Validator Agent**
  - REST/gRPC standards enforcement
  - DTO validation and consistency
  - Breaking change detection
  - Error response standardization
  - Pagination and versioning validation

- **Security Reviewer Agent**
  - Vulnerability detection (SQL injection, XSS, command injection)
  - Authentication/authorization review
  - Secrets management validation
  - CORS and security headers
  - Rate limiting and input validation

- **Error Handling Consultant Agent**
  - Error pattern consistency
  - Error wrapping and context
  - Domain error type design
  - HTTP error mapping
  - Structured logging guidance

- **Database Migration Assistant Agent**
  - Migration safety review
  - Backward compatibility validation
  - Safe index creation strategies
  - Data loss prevention
  - Transaction and rollback guidance

#### Automation Features
- **PostToolUse Hooks**
  - Domain purity validation on file edits
  - Automatic architecture test execution
  - Real-time feedback

- **MCP Server - GoArchTest Analyzer**
  - `check_layer_dependencies` - Layer dependency validation
  - `check_domain_isolation` - Domain boundary verification
  - `check_naming_conventions` - Naming standard enforcement
  - `run_all_architecture_tests` - Full test suite execution
  - `generate_dependency_graph` - Dependency visualization

#### Scripts
- `validate-domain-purity.sh` - Domain layer validation script
  - Ensures only stdlib and shared imports
  - Prevents infrastructure leakage
  - Clear violation reporting

#### Documentation
- Comprehensive README with usage examples
- MCP server documentation
- Architecture principles guide
- DDD tactical patterns reference
- Agent role descriptions
- **GODOG_BEST_PRACTICES.md** - Detailed Godog implementation guide
  - Modern patterns (v0.12.0+: context.Context, TestSuite, go test)
  - Step definition best practices
  - Concurrency and thread-safety
  - Feature file quality standards
  - Common pitfalls and anti-patterns

### Architecture Principles

- **Hexagonal Architecture**: Ports and Adapters pattern
- **Domain-Driven Design**: Tactical and strategic patterns
- **Test-Driven Development**: Red-Green-Refactor cycle
- **Behavior-Driven Development**: Godog with Gherkin syntax
- **Microservices**: Monorepo with independent services
- **Protocol Agnostic**: REST, gRPC, or both
- **Automated Testing**: goarchtest integration

### Dependencies

- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - Configuration management
- `github.com/cucumber/godog` - BDD testing
- `github.com/stretchr/testify` - Test assertions
- `github.com/solrac97gr/goarchtest` - Architecture testing
- `@modelcontextprotocol/sdk` - MCP server (Node.js)

### Keywords

golang, hexagonal-architecture, clean-architecture, ddd, domain-driven-design, tdd, bdd, godog, microservices, architecture-testing, goarchtest

---

## Future Enhancements (Planned)

- [ ] GraphQL support
- [ ] Event-driven architecture patterns
- [ ] CQRS support
- [ ] Event sourcing templates
- [ ] Kubernetes deployment manifests
- [ ] CI/CD pipeline templates
- [ ] API documentation generation
- [ ] Performance testing templates
- [ ] Chaos engineering support
- [ ] Multi-tenancy patterns

---

**Legend:**
- Added: New features
- Changed: Changes in existing functionality
- Deprecated: Soon-to-be removed features
- Removed: Removed features
- Fixed: Bug fixes
- Security: Security improvements
