---
description: Scaffold new feature following TDD/BDD with Godog
---

Create a complete feature following TDD/BDD principles, starting with the Godog feature file.

**Workflow (CRITICAL - Follow this order):**

## Phase 1: Feature Discovery & BDD Feature File

1. **Ask clarifying questions** using the AskUserQuestion tool until the feature is completely clear:
   - Which domain/bounded context? (e.g., "user", "order", "product")
   - What is the feature name? (e.g., "CreateUser", "PlaceOrder")
   - What is the business goal/user story?
   - What are the acceptance criteria?
   - What are the main scenarios (happy path and edge cases)?
   - What input data is needed?
   - What validations are required?
   - What should happen on success?
   - What should happen on failure (error cases)?
   - Are there any dependencies on other entities/aggregates?

2. **Create the Godog feature file** in `features/[domain]/[feature_name].feature`:
   - Use proper Gherkin syntax (Feature, Scenario, Given, When, Then)
   - Include multiple scenarios: happy path + edge cases
   - Make it human-readable (non-technical people should understand it)
   - Use concrete examples with real data

3. **Review the feature file with the user** before proceeding to code

## Phase 2: Domain Layer (TDD - Red Phase)

4. **Create domain entity** in `internal/[domain]/domain/entity.go`:
   - Define the aggregate root or entity
   - Include value objects if needed
   - Add business logic methods
   - Add domain validations
   - NO external dependencies

5. **Create repository interface** (port) in `internal/[domain]/domain/repository.go`:
   - Define contract for persistence
   - Use domain entities only
   - Return domain errors

6. **Write unit tests** in `test/unit/[domain]/entity_test.go`:
   - Test entity creation
   - Test business logic
   - Test domain validations
   - Tests should FAIL initially (Red phase)

## Phase 3: Application Layer (Use Cases)

7. **Create use case** in `internal/[domain]/application/usecase/[feature_name]_usecase.go`:
   - Define input/output DTOs (or in `internal/[domain]/application/dto/`)
   - Implement business workflow
   - Use repository interface
   - Handle errors appropriately

8. **Write use case tests** in `test/unit/[domain]/usecase_test.go`:
   - Mock repository
   - Test happy path
   - Test error cases
   - Tests should FAIL initially

## Phase 4: Infrastructure Layer (Adapters)

9. **Create repository implementation** in `internal/[domain]/infrastructure/persistence/[db_type]_repository.go`:
   - Implement domain repository interface
   - Map between domain entities and database models
   - Handle database operations

10. **Create HTTP handler** in `internal/[domain]/infrastructure/http/handler.go`:
    - Handle HTTP request/response
    - Validate input
    - Call use case
    - Map to HTTP status codes
    - Return appropriate JSON responses

11. **Add route** to router configuration (usually in `cmd/api/main.go` or `internal/shared/`)

## Phase 5: BDD Tests (Godog Step Definitions)

12. **Create step definitions** in `features/[domain]/[feature_name]_test.go`:
    - Implement Given/When/Then steps
    - Use real implementations (integration-style)
    - Setup test database/dependencies
    - Run full flow end-to-end

13. **Run BDD tests** - they should PASS now (Green phase)

## Phase 6: Integration Tests

14. **Create integration test** in `test/integration/[domain]/[feature_name]_test.go`:
    - Test HTTP endpoint
    - Test with real database (testcontainers or similar)
    - Test full request/response cycle

**Architecture Rules to Enforce:**

- Domain layer: pure business logic, no infrastructure
- Application layer: orchestrate domain, use interfaces
- Infrastructure layer: technical implementations
- Dependencies flow: Infrastructure → Application → Domain
- All layers must be tested independently
- Feature file comes FIRST, code comes AFTER

**File Naming Conventions:**

- Entities: `user.go`, `order.go` (singular)
- Repositories: `user_repository.go`
- Use cases: `create_user_usecase.go`, `place_order_usecase.go`
- Handlers: `user_handler.go`, `order_handler.go`
- Tests: `*_test.go`
- Features: `*.feature`

Be thorough but concise. Ask questions first, code later.
