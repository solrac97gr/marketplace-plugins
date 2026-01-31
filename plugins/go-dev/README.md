# Go Hexagonal Architecture Development Plugin

Complete toolkit for building Go backend applications with hexagonal architecture (ports and adapters), Domain-Driven Design (DDD), and Test-Driven Development (TDD/BDD).

## Features

### 🏗️ Project Scaffolding
- **Domain-first organization** with intelligent domain suggestions
- **Microservices monorepo** architecture with Cobra CLI
- **Multi-protocol support**: REST API, gRPC, or both
- **Multi-database support**: PostgreSQL, MongoDB, MySQL, SQLite
- **Automated architecture tests** using [goarchtest](https://github.com/solrac97gr/goarchtest)

### 🎯 Skills

#### `/start-project`
Initialize a new Go project with complete hexagonal architecture setup.

**Process:**
1. Asks about your project type (e-commerce, banking, etc.)
2. Suggests relevant domains/bounded contexts
3. User selects which domains to create (multi-select)
4. Configures database and protocol preferences
5. Generates complete project structure with architecture tests

**Creates:**
- Cobra CLI with commands for each microservice
- Domain/Application/Infrastructure layers per domain
- Automated architecture tests
- BDD test setup with Godog
- Makefile with common commands
- Docker Compose for local development

#### `/new-feature`
Scaffold a complete feature following TDD/BDD workflow.

**Process:**
1. Asks clarifying questions about the feature
2. Creates Godog `.feature` file with Gherkin syntax (human-readable)
3. Generates domain entities with business logic
4. Creates use cases in application layer
5. Implements adapters in infrastructure layer
6. Writes unit, integration, and BDD tests

**Follows:** Feature file → Domain → Application → Infrastructure → Tests

#### `/new-entity`
Create a domain entity following DDD principles.

**Creates:**
- Domain entity with business logic
- Repository interface (port)
- Value objects if needed
- Unit tests

#### `/new-usecase`
Create an application use case.

**Creates:**
- Use case with business workflow
- Input/Output DTOs
- Unit tests with mocks

#### `/review-arch`
Comprehensive architecture review combining automated tests and AI analysis.

**Process:**
1. **Runs automated tests** using goarchtest
2. **AI code analysis** for patterns and violations
3. **Combined report** with prioritized recommendations

**Checks:**
- Layer dependency violations
- Domain isolation
- DDD pattern compliance
- Test coverage
- Naming conventions
- Anti-patterns

#### `/update-arch-tests`
Create or update architecture tests using goarchtest.

**Generates tests for:**
- Layer dependencies (domain → application → infrastructure)
- Domain isolation (no cross-domain dependencies)
- Naming conventions (repositories, use cases, handlers)
- Custom architectural rules

### 🤖 Intelligent Agents

#### Architecture Reviewer Agent
Proactively reviews code changes for architectural compliance.

**Reviews:**
- Dependency direction
- Layer purity
- DDD pattern usage
- Common anti-patterns

#### DDD Consultant Agent
Expert guidance on Domain-Driven Design modeling.

**Helps with:**
- Bounded context identification
- Entity vs Value Object decisions
- Aggregate design
- Repository pattern
- Domain events
- Ubiquitous language

#### Domain Documentation Writer Agent
Automatically generates and maintains standardized domain documentation.

**Creates:**
- Architecture diagrams using Mermaid
- Domain entity documentation
- Use case descriptions
- API endpoint listings
- Database schema documentation

#### Test Coverage Guardian Agent
Enforces TDD/BDD practices and monitors test coverage.

**Ensures:**
- Tests written before implementation
- Coverage targets (Domain 90%+, Application 85%+, Infrastructure 70%+)
- Godog feature completeness
- Test quality standards

#### API Contract Validator Agent
Maintains API consistency and prevents breaking changes.

**Validates:**
- REST/gRPC standards compliance
- Request/response DTOs
- Breaking change detection
- Error response consistency
- Pagination standards

#### Security Reviewer Agent
Identifies security vulnerabilities and enforces security best practices.

**Checks:**
- SQL injection, XSS, command injection
- Authentication/authorization patterns
- Password hashing and secrets management
- CORS and security headers
- Rate limiting and input validation

#### Error Handling Consultant Agent
Ensures consistent error handling patterns.

**Reviews:**
- Error creation and wrapping
- Domain error types
- HTTP error mapping
- Logging practices
- Error message quality

#### Database Migration Assistant Agent
Reviews migration scripts for safety and compatibility.

**Validates:**
- Backward compatibility
- Safe index creation
- Data loss prevention
- Transaction usage
- Rollback strategies

#### Code Quality Reviewer Agent
Enforces Uber Go Style Guide best practices on all Go code.

**Dynamic Approach:**
- Fetches latest guide from https://github.com/uber-go/guide/blob/master/style.md
- Always up-to-date with language evolution
- No stale documentation

**Reviews:**
- Interface usage patterns
- Concurrency safety (mutexes, atomics)
- Error handling idioms
- Initialization patterns
- Performance optimizations
- Code style and idioms

### 🔧 Automated Hooks

#### Domain Purity Validation
Automatically runs when domain layer files are modified.

**Validates:**
- No external dependencies (only stdlib and shared)
- No infrastructure imports
- No framework dependencies

#### Architecture Test Execution
Automatically runs architecture tests when test files are updated.

### 🔌 MCP Server

**GoArchTest Analyzer** - Real-time architecture analysis

**Tools:**
- `check_layer_dependencies` - Validate layer dependencies
- `check_domain_isolation` - Verify domain boundaries
- `check_naming_conventions` - Enforce naming standards
- `run_all_architecture_tests` - Execute full test suite
- `generate_dependency_graph` - Visualize dependencies

## Project Structure

```
your-project/
├── cmd/
│   ├── main.go                    # Cobra CLI entry
│   ├── root.go                    # Root command
│   └── serve/
│       ├── user.go                # User microservice command
│       ├── order.go               # Order microservice command
│       └── all.go                 # Run all microservices
├── api/proto/                     # gRPC definitions (if using gRPC)
│   └── [domain]/
├── internal/
│   ├── user/                      # User bounded context
│   │   ├── domain/                # Pure domain logic
│   │   │   ├── entity.go
│   │   │   ├── repository.go      # Interfaces (ports)
│   │   │   └── errors.go
│   │   ├── application/           # Use cases & DTOs
│   │   │   ├── usecase/
│   │   │   └── dto/
│   │   └── infrastructure/        # Adapters
│   │       ├── persistence/
│   │       ├── http/              # REST handlers
│   │       └── grpc/              # gRPC handlers
│   ├── order/                     # Order bounded context
│   └── shared/                    # Shared kernel
├── test/
│   ├── architecture/              # goarchtest tests
│   ├── features/[domain]/         # Godog BDD tests
│   ├── integration/[domain]/
│   └── unit/[domain]/
└── Makefile
```

## Architecture Principles

### Hexagonal Architecture (Ports & Adapters)
- **Domain** (center): Pure business logic, no dependencies
- **Application**: Use cases, orchestrates domain
- **Infrastructure** (outer): Database, HTTP, gRPC, external services

### Dependency Direction
```
Infrastructure → Application → Domain
```

### DDD Tactical Patterns
- **Entities**: Identity + behavior
- **Value Objects**: Immutable, self-validating
- **Aggregates**: Consistency boundaries
- **Repositories**: Persistence abstraction
- **Domain Services**: Cross-entity behavior
- **Domain Events**: Business occurrences

### TDD/BDD Workflow
1. Write Godog feature file (Given/When/Then)
2. Write failing tests (Red)
3. Implement minimum code (Green)
4. Refactor
5. Repeat

## Usage Example

```bash
# Start a new project
/start-project
# Answers: "E-commerce platform"
# Selects: user, product, order, payment domains
# Chooses: PostgreSQL, REST + gRPC, Gin framework

# Add a new feature
/new-feature
# Answers questions about the feature
# Creates feature file, domain, application, infrastructure

# Review architecture
/review-arch
# Runs goarchtest + AI analysis
# Shows violations and recommendations

# Run microservices
make serve-user     # User service on :8080
make serve-order    # Order service on :8081
make serve-all      # All services together

# Run tests
make test           # All tests
make test-arch      # Architecture tests only
make test-bdd       # BDD tests only
```

## Installation

This plugin is part of the solrac97gr marketplace. Configure your Claude Code to use this marketplace.

## Dependencies

The plugin will install these Go packages:
- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - Configuration
- `github.com/cucumber/godog` - BDD testing
- `github.com/stretchr/testify` - Test assertions
- `github.com/solrac97gr/goarchtest` - Architecture testing
- Database drivers (based on selection)
- Web frameworks (based on selection)

## Benefits

✅ **Enforced Architecture**: Automated tests prevent violations
✅ **Fast Scaffolding**: Generate complete features in minutes
✅ **DDD Best Practices**: Built-in expert guidance
✅ **Microservices Ready**: Single binary, multiple services
✅ **Protocol Agnostic**: REST, gRPC, or both
✅ **Test-Driven**: BDD with Godog from the start
✅ **Maintainable**: Clean separation of concerns
✅ **Scalable**: Independent domains, clear boundaries

## Learn More

- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)
- [Domain-Driven Design](https://martinfowler.com/bliki/DomainDrivenDesign.html)
- [goarchtest](https://github.com/solrac97gr/goarchtest)
- [Godog BDD](https://github.com/cucumber/godog)

## License

MIT
