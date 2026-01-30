---
description: Create or update architecture tests with goarchtest
---

Create or update architecture tests using goarchtest to enforce hexagonal architecture and DDD constraints.

**What to Do:**

## Phase 1: Analyze Current Structure

1. Scan the project structure to identify:
   - All domains under `internal/`
   - Existing layers (domain, application, infrastructure)
   - Current architecture test file (if exists)

## Phase 2: Create/Update Architecture Tests

2. Create or update `test/architecture/architecture_test.go` with comprehensive tests:

### Test Structure Template

```go
package architecture_test

import (
    "path/filepath"
    "testing"

    "github.com/solrac97gr/goarchtest"
    "github.com/stretchr/testify/assert"
)

func TestHexagonalArchitecture(t *testing.T) {
    projectPath, err := filepath.Abs("../../")
    if err != nil {
        t.Fatal(err)
    }

    t.Run("Domain layer should not depend on infrastructure", func(t *testing.T) {
        result := goarchtest.InPath(projectPath).
            That().
            ResideInNamespace("internal/*/domain").
            ShouldNot().
            HaveDependencyOn("internal/*/infrastructure").
            GetResult()

        assert.True(t, result.IsSuccessful,
            "Domain layer should not depend on infrastructure layer")
    })

    t.Run("Domain layer should not depend on application", func(t *testing.T) {
        result := goarchtest.InPath(projectPath).
            That().
            ResideInNamespace("internal/*/domain").
            ShouldNot().
            HaveDependencyOn("internal/*/application").
            GetResult()

        assert.True(t, result.IsSuccessful,
            "Domain layer should not depend on application layer")
    })

    t.Run("Application layer should not depend on infrastructure", func(t *testing.T) {
        result := goarchtest.InPath(projectPath).
            That().
            ResideInNamespace("internal/*/application").
            ShouldNot().
            HaveDependencyOn("internal/*/infrastructure").
            GetResult()

        assert.True(t, result.IsSuccessful,
            "Application layer should not depend on infrastructure layer")
    })
}

func TestDomainIsolation(t *testing.T) {
    projectPath, _ := filepath.Abs("../../")

    // Add test for each domain pair to ensure no cross-domain dependencies
    // Example: user domain should not depend on order domain
    t.Run("User domain should not depend on Order domain", func(t *testing.T) {
        result := goarchtest.InPath(projectPath).
            That().
            ResideInNamespace("internal/user/").
            ShouldNot().
            HaveDependencyOn("internal/order/").
            GetResult()

        assert.True(t, result.IsSuccessful,
            "Domains should be isolated from each other")
    })

    // Domains CAN depend on shared
    t.Run("Domains can depend on shared kernel", func(t *testing.T) {
        result := goarchtest.InPath(projectPath).
            That().
            ResideInNamespace("internal/*/domain").
            Should().
            HaveDependencyOn("internal/shared/").
            Or().
            ShouldNot().
            HaveDependencyOn("internal/shared/").
            GetResult()

        // This should always pass - just verifying shared is accessible
        assert.True(t, result.IsSuccessful)
    })
}

func TestNamingConventions(t *testing.T) {
    projectPath, _ := filepath.Abs("../../")

    t.Run("Repository interfaces should end with Repository", func(t *testing.T) {
        result := goarchtest.InPath(projectPath).
            That().
            ResideInNamespace("internal/*/domain").
            And().
            HaveNameEndingWith("Repository").
            Should().
            BeInterfaces().
            GetResult()

        assert.True(t, result.IsSuccessful,
            "Repositories in domain layer should be interfaces")
    })

    t.Run("Use cases should end with UseCase", func(t *testing.T) {
        result := goarchtest.InPath(projectPath).
            That().
            ResideInNamespace("internal/*/application/usecase").
            Should().
            HaveNameEndingWith("UseCase").
            GetResult()

        assert.True(t, result.IsSuccessful,
            "Use cases should follow naming convention")
    })

    t.Run("Handlers should end with Handler", func(t *testing.T) {
        result := goarchtest.InPath(projectPath).
            That().
            ResideInNamespace("internal/*/infrastructure/http").
            And().
            HaveNameEndingWith("Handler").
            GetResult()

        // This is a soft check - not all files need to be handlers
        if !result.IsSuccessful {
            t.Log("Consider following Handler naming convention")
        }
    })
}

func TestInfrastructureConstraints(t *testing.T) {
    projectPath, _ := filepath.Abs("../../")

    t.Run("HTTP handlers should not contain business logic", func(t *testing.T) {
        // This checks that handlers don't import domain entities directly
        result := goarchtest.InPath(projectPath).
            That().
            ResideInNamespace("internal/*/infrastructure/http").
            ShouldNot().
            HaveDependencyOn("internal/*/domain/entity").
            GetResult()

        // Handlers should use DTOs, not domain entities
        if !result.IsSuccessful {
            t.Log("Warning: Handlers should use DTOs from application layer, not domain entities")
        }
    })
}
```

## Phase 3: Domain-Specific Tests

3. **Generate tests for each domain found** in the project:
   - For each domain under `internal/`, create isolation tests
   - Ensure domain A doesn't depend on domain B (except through shared)

4. **Add tests for specific architectural patterns:**
   - Repository pattern: interfaces in domain, implementations in infrastructure
   - Use case pattern: in application layer, uses domain interfaces
   - Handler pattern: in infrastructure, uses application use cases

## Phase 4: Custom Rules (if needed)

5. Ask the user if they have any **custom architectural rules**:
   - Specific naming conventions
   - Additional layer constraints
   - Custom patterns to enforce

6. Add custom predicates for domain-specific rules

## Phase 5: Integration

7. Ensure the architecture tests are integrated:
   - Add to `make test-arch` in Makefile
   - Add to CI/CD pipeline (if exists)
   - Document in README.md

8. Run the tests to verify they work:
   ```bash
   go test ./test/architecture/... -v
   ```

9. If tests fail, explain the violations and offer to fix them

**Important Notes:**

- Tests should be strict but realistic
- Start with core constraints, add more over time
- Each test should have a clear, descriptive name
- Use assertions with helpful error messages
- Consider generating dependency graphs for complex violations

**When to Update:**

- When adding a new domain
- When adding new architectural layers
- When discovering architectural drift
- When refining architectural rules
- After `/review-arch` finds issues

Be thorough and create comprehensive tests that enforce the hexagonal architecture.
