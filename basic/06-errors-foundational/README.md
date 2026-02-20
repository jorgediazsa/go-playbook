# 06 - Errors (Foundational)

Go's error handling distinguishes it from most modern languages. Errors are just values that implement a simple interface. There are no exceptions (try/catch), which forces developers to treat errors as part of the normal program flow.

## 1. Designing Error-Returning APIs

In Go, an API that might fail should return `error` as its last return value.

### The Production Danger

Relying on `panic` for flow control (like throwing exceptions in Java/C#) is unidiomatic and dangerous in Go. A panic that isn't explicitly recovered will crash the entire Go process. Senior engineers must design APIs that return deterministic error values, forcing the caller to acknowledge the failure path.

```go
// BAD:
func ParseUser(payload []byte) User {
    if len(payload) == 0 {
        panic("empty payload") // Don't do this!
    }
    // ...
}

// IDIOMATIC:
func ParseUser(payload []byte) (User, error) {
    if len(payload) == 0 {
        return User{}, errors.New("empty payload")
    }
    // ...
}
```

---

## 2. Early Returns ("Line of Sight")

Go encourages the "Line of Sight" rule: the happy path is always left-aligned, and errors are handled immediately via early returns.

### The Production Danger

Deeply nesting the "happy path" inside `if err == nil` blocks makes code incredibly hard to read and test.

```go
// BAD (Nested Happy Path):
func Process() error {
    data, err := Fetch()
    if err == nil {
        parsed, parseErr := Parse(data)
        if parseErr == nil {
            return Save(parsed)
        } else {
            return parseErr
        }
    } else {
        return err
    }
}

// IDIOMATIC (Early Returns):
func Process() error {
    data, err := Fetch()
    if err != nil {
        return err
    }
    
    parsed, err := Parse(data)
    if err != nil {
        return err
    }
    
    return Save(parsed)
}
```

---

## 3. Sentinel Errors and Wrapping (Preview)

When an error occurs, simply returning `err` loses context about *where* it happened. Returning `fmt.Errorf("db fetch failed: %w", err)` adds textual context while preserving the original error type, allowing callers to inspect the cause using `errors.Is` and `errors.As`.

---

## Exercises

- `ex01_api_design.go`
- `ex02_wrapping.go`
