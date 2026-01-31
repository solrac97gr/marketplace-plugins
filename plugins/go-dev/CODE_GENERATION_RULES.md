# Code Generation Rules

**Quick reference for code-generating skills. For complete and latest guidelines, always fetch from [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).**

## Fetching Latest Guidelines

Before generating significant code, use the `fetch_webpage` tool:
```
fetch_webpage("https://github.com/uber-go/guide/blob/master/style.md", "latest Go best practices")
```

This ensures you're using the most current patterns as the language evolves.

## Quick Checklist

When generating Go code, ALWAYS:

### Interfaces
- [ ] Pass interfaces as values, not pointers
- [ ] Add compile-time interface verification: `var _ Interface = (*Type)(nil)`

### Structs
- [ ] Initialize with field names: `User{Name: "x", Email: "y"}`
- [ ] Omit zero-value fields
- [ ] Use `var` for zero-value structs: `var user User`
- [ ] Use `&T{}` not `new(T)` for struct pointers

### Error Handling
- [ ] Prefix error vars with `Err`: `var ErrNotFound = errors.New("not found")`
- [ ] Suffix error types with `Error`: `type NotFoundError struct{}`
- [ ] Wrap errors with `%w`: `fmt.Errorf("operation: %w", err)`
- [ ] Add concise context: `"new store"` not `"failed to create new store"`
- [ ] Handle errors once (don't log AND return)

### Concurrency
- [ ] Don't embed mutexes in public structs
- [ ] Use zero-value mutexes: `var mu sync.Mutex` not `new(sync.Mutex)`
- [ ] Prefer `go.uber.org/atomic` for atomic operations

### Initialization
- [ ] Use `make()` for empty maps: `make(map[K]V)`
- [ ] Provide capacity hints when known: `make([]T, 0, size)`
- [ ] Use map literals for fixed elements

### Functions
- [ ] Use `defer` for cleanup (locks, files, etc.)
- [ ] Return errors, don't panic
- [ ] Exit only in `main()`

### Style
- [ ] Group imports: stdlib first, then external
- [ ] Prefix unexported package vars with `_`: `var _defaultPort = 8080`
- [ ] Minimize nesting (early returns)
- [ ] Minimize variable scope
- [ ] Use `:=` for explicit values, `var` when zero-value is clearer

### Types
- [ ] Use `time.Time` for instants, `time.Duration` for periods
- [ ] Start enums at 1 (unless zero is meaningful)
- [ ] Use comma-ok for type assertions: `v, ok := i.(string)`
- [ ] Don't shadow built-in names

### Testing
- [ ] Use table-driven tests: `tests := []struct{ give, want Type }{}`
- [ ] Name test slice `tests`, test case `tt`
- [ ] Use `give` prefix for inputs, `want` for outputs

### Performance
- [ ] Use `strconv.Itoa()` not `fmt.Sprint()` for primitives
- [ ] Avoid repeated []byte("string") conversions
- [ ] Specify capacity for frequently allocated containers

## Common Patterns

### Constructor
```go
func NewUser(email, name string) (*User, error) {
    // Validate
    if email == "" {
        return nil, ErrInvalidEmail
    }
    
    return &User{
        ID:    uuid.New(),
        Email: email,
        Name:  name,
    }, nil
}
```

### Error Variable
```go
var ErrNotFound = errors.New("not found")
```

### Custom Error
```go
type ValidationError struct {
    Field string
    Err   error
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on field %q: %v", e.Field, e.Err)
}
```

### Error Wrapping
```go
user, err := repo.GetUser(id)
if err != nil {
    return nil, fmt.Errorf("get user %q: %w", id, err)
}
```

### Defer Cleanup
```go
func process() error {
    mu.Lock()
    defer mu.Unlock()
    
    // ... critical section
    return nil
}
```

### Table-Driven Test
```go
func TestValidate(t *testing.T) {
    tests := []struct {
        name      string
        give      User
        wantError bool
    }{
        {
            name:      "valid user",
            give:      User{Email: "test@example.com"},
            wantError: false,
        },
        {
            name:      "invalid email",
            give:      User{Email: "invalid"},
            wantError: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.give.Validate()
            if tt.wantError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
```

### Functional Options
```go
type Option interface {
    apply(*config)
}

type config struct {
    timeout time.Duration
    retries int
}

type timeoutOption time.Duration

func (o timeoutOption) apply(c *config) {
    c.timeout = time.Duration(o)
}

func WithTimeout(d time.Duration) Option {
    return timeoutOption(d)
}

func NewClient(addr string, opts ...Option) *Client {
    cfg := config{
        timeout: 30 * time.Second,
        retries: 3,
    }
    
    for _, opt := range opts {
        opt.apply(&cfg)
    }
    
    return &Client{addr: addr, config: cfg}
}

// Usage
client := NewClient("localhost:8080", WithTimeout(10*time.Second))
```

### Interface Verification
```go
type Handler struct {
    logger *zap.Logger
}

// Verify at compile time
var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // ...
}
```

## Anti-Patterns to Avoid

### ❌ Pointer to Interface
```go
func process(r *io.Reader) {}  // NO!
func process(r io.Reader) {}   // YES
```

### ❌ Embedded Mutex in Public Struct
```go
type Data struct {
    sync.Mutex  // NO!
    items []string
}

type Data struct {
    mu    sync.Mutex  // YES
    items []string
}
```

### ❌ Log and Return Error
```go
if err != nil {
    log.Printf("error: %v", err)  // NO!
    return err
}

if err != nil {
    return fmt.Errorf("context: %w", err)  // YES
}
```

### ❌ Panic in Production
```go
if len(args) == 0 {
    panic("missing argument")  // NO!
}

if len(args) == 0 {
    return errors.New("missing argument")  // YES
}
```

### ❌ No Field Names
```go
user := User{"john", "doe", 25}  // NO!

user := User{
    FirstName: "john",
    LastName:  "doe", 
    Age:       25,
}  // YES
```

## Reference

For complete details, see:
- [UBER_GO_STYLE_GUIDE.md](UBER_GO_STYLE_GUIDE.md) - Full reference
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) - Official guide

## Enforcement

The Code Quality Reviewer agent automatically checks generated code against these rules.
