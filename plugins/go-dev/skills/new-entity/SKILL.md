---
description: Create domain entity with repository interface
---

Create a new domain entity following DDD principles with its repository interface.

**IMPORTANT - Code Quality:**
Before generating code, fetch the latest Uber Go Style Guide:
```
fetch_webpage("https://github.com/uber-go/guide/blob/master/style.md", "latest Go best practices")
```
This ensures all generated code follows current production standards.

**What to Create:**

1. **Ask the user:**
   - Which domain/bounded context? (e.g., "user", "order", "product")
   - Entity name (e.g., User, Product, Order)
   - Fields/attributes with types
   - Business rules/validations
   - Is it an aggregate root or entity?
   - Value objects needed?

2. **Create entity** in `internal/[domain]/domain/entity.go` (or `[entity_name].go` if multiple entities):
   - Define struct with fields
   - Add constructor (New[Entity]) with validation
   - Add business logic methods
   - Add domain validations
   - Use value objects where appropriate
   - NO infrastructure dependencies

3. **Create repository interface** in `internal/[domain]/domain/repository.go`:
   - Define persistence contract
   - Common methods: Save, FindByID, Update, Delete, List
   - Use domain types only
   - Return domain errors

4. **Create unit tests** in `test/unit/[domain]/entity_test.go`:
   - Test entity creation (valid and invalid)
   - Test business logic methods
   - Test validations
   - Use table-driven tests

5. **Create value objects** (if needed) in `internal/shared/valueobject/[value_object_name].go`:
   - Email, Money, Address, etc.
   - Immutable
   - Self-validating
   - With unit tests

**DDD Principles:**

- Entities have identity (ID field)
- Rich domain model (behavior, not just data)
- Encapsulation (use methods, not public setters)
- Invariants are always protected
- Use ubiquitous language
- Value objects for concepts without identity

**Example Entity Structure:**

```go
type User struct {
    id        uuid.UUID
    email     Email // value object
    name      string
    createdAt time.Time
}

func NewUser(email, name string) (*User, error) {
    // validation
    // business rules
    return &User{...}, nil
}

func (u *User) ChangeName(newName string) error {
    // validation
    // business rules
}
```

Be concise and follow Go best practices.
