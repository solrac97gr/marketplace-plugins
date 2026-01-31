---
description: Create application use case with tests
---

Create a new use case in the application layer following TDD principles.

**IMPORTANT - Code Quality:**
Before generating code, fetch the latest Uber Go Style Guide:
```
fetch_webpage("https://github.com/uber-go/guide/blob/master/style.md", "latest Go best practices")
```
This ensures all generated code follows current production standards.

**What to Create:**

1. **Ask the user:**
   - Which domain/bounded context? (e.g., "user", "order", "product")
   - Use case name (e.g., CreateUser, PlaceOrder, GetUserProfile)
   - What does it do? (business workflow)
   - What are the inputs?
   - What are the outputs?
   - What repositories/services does it need?
   - What validations are required?
   - What are the error cases?

2. **Create DTOs** (if not exists) in `internal/[domain]/application/dto/`:
   - Input DTO (request)
   - Output DTO (response)
   - Keep them simple, focused on the use case

3. **Create use case file** in `internal/[domain]/application/usecase/[usecase_name]_usecase.go`:
   - Define UseCase interface (if following interface-based approach)
   - Define struct with dependencies (repositories, services)
   - Implement Execute method
   - Handle business workflow
   - Validate inputs
   - Handle errors
   - Return DTOs

4. **Write unit tests FIRST** in `test/unit/[domain]/[usecase_name]_usecase_test.go`:
   - Mock all dependencies
   - Test happy path
   - Test validation errors
   - Test repository errors
   - Test business rule violations
   - Use table-driven tests
   - Tests should FAIL initially (Red phase TDD)

5. **Implement the use case** to make tests pass (Green phase)

**Use Case Structure Example:**

```go
type CreateUserUseCase struct {
    userRepo repository.UserRepository
}

func NewCreateUserUseCase(repo repository.UserRepository) *CreateUserUseCase {
    return &CreateUserUseCase{userRepo: repo}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
    // 1. Validate input
    // 2. Create domain entity
    // 3. Call repository
    // 4. Map to output DTO
    // 5. Return result
}
```

**Application Layer Principles:**

- Orchestrate domain objects
- Use domain repository interfaces (NOT implementations)
- No business logic here (that belongs in domain)
- Handle transaction boundaries
- Return application-level errors
- Use dependency injection
- Stateless (no instance variables that change)

**Testing Approach:**

- Mock all external dependencies
- Focus on workflow logic
- Test all error paths
- Use testify/mock or mockery for mocks
- Arrange-Act-Assert pattern

Be concise and follow clean architecture principles.
