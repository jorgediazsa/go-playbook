# Go (Golang) Learning Roadmap - Senior Engineer Track

This roadmap is designed for experienced software engineers who want to learn Go deeply, idiomatically, and production-ready - not just syntactically.

It emphasizes:
- Correct mental models
- Go-specific pitfalls
- Concurrency safety
- Production hygiene
- Tooling mastery

---

# 🟢 BASIC - Language Fundamentals (With Go-Specific Gotchas)

## 1. Core Syntax
- Go program structure (`package main`, `func main`)
- Variable declaration (`var`, `:=`)
- Zero values
- Basic types: `int`, `float64`, `string`, `bool`
- Untyped constants
- Constants (`const`)
- Type conversions
- Type aliases vs defined types

---

## 2. Control Flow
- `if`, `else if`, `else`
- Expression-less `switch`
- `for` (the only loop)
- `for range` (slices, maps, strings)
- Loop variable capture pitfalls
- Labeled break / continue

---

## 3. Functions
- Multiple return values
- Named return parameters
- Variadic functions
- Anonymous functions
- Closures and variable capture semantics
- Function values and higher-order functions
- Defer evaluation timing

---

## 4. Standard Collections

### Arrays
- Value semantics
- When arrays are useful

### Slices
- Length vs capacity
- `make`, `append`, slicing
- Nil vs empty slices
- Append reallocation behavior
- Aliasing and shared backing arrays
- Memory retention via slicing (subslice leaks)

### Maps
- Declaration, assignment
- Safe lookup (`value, ok`)
- Nil maps behavior
- Iteration randomness
- Map concurrency restrictions

### Structs
- Definition
- Zero values
- Struct tags
- Comparability rules

---

## 5. Pointers
- `&`, `*`
- Pointer to struct
- Mutability vs copy
- Method sets (preview)
- When to use pointers vs values
- Escape implications (preview)

---

## 6. Errors (Foundational)
- `error` as a value
- `errors.New`, `fmt.Errorf`
- Value, err pattern
- Early return pattern
- Avoiding panic for flow control

---

## 7. Defer, Panic, Recover
- `defer` execution order (LIFO)
- Defer + named return interaction
- Resource safety patterns
- `panic` vs `error`
- `recover` boundaries
- Designing panic-safe libraries

---

## 8. Strings, Bytes, and Runes
- UTF-8 model
- `string` vs `[]byte`
- `rune`
- Iterating over strings
- Conversion costs
- Avoiding unnecessary allocations

---

## 9. Time
- `time.Duration`
- `time.Now()` monotonic behavior
- `time.After` vs `time.NewTimer`
- Timer and ticker leaks
- Parsing and formatting
- Deadlines vs timeouts

---

# 🟡 INTERMEDIATE - Idiomatic & Production Go

## 10. Methods and Composition
- Value vs pointer receivers
- Method sets and interface satisfaction
- Embedding (composition over inheritance)
- Method promotion and shadowing
- Implicit interfaces
- Small interfaces (single-method)
- `io.Reader` / `io.Writer` patterns
- Designing minimal contracts

---

## 11. Packages and Modules
- Code organization principles
- Internal packages
- Public API surface control
- `go mod init`, `go mod tidy`
- `replace`, `retract`
- Semantic versioning
- Private modules
- Workspaces (`go work`)
- Avoiding import cycles
- `GOMODCACHE`
- Dependency hygiene

---

## 12. Tooling & Static Analysis
- `go build`
- `go test ./...`
- `go vet`
- `-race`
- `go list`
- `go env`
- Build tags
- `go generate`
- `pprof` basics
- Static analysis tools (staticcheck overview)

---

## 13. Goroutines
- `go func() {}`
- Concurrency vs parallelism
- Scheduling model overview
- Goroutine lifecycle
- Avoiding goroutine leaks
- Proper shutdown patterns
- Shared memory hazards
- Data race detection

---

## 14. Channels
- `make(chan T)`
- Buffered vs unbuffered
- Send / receive
- Directional channels
- `select`
- Closing channels
- Ranging over channels
- Avoiding deadlocks
- Channel ownership rules

---

## 15. Synchronization Primitives
- `sync.WaitGroup`
- `sync.Mutex`
- `sync.RWMutex`
- `sync.Once`
- `sync.Map`
- `sync.Cond`
- `sync/atomic`
- Choosing the right primitive

---

## 16. Context
- `context.Background`
- `WithCancel`, `WithTimeout`, `WithDeadline`
- Cancellation propagation
- Cooperative cancellation
- Context in HTTP handlers
- Context in DB queries
- Context misuse patterns

---

## 17. Idiomatic Error Handling (Advanced)
- `errors.Is`
- `errors.As`
- Error wrapping
- Custom error types
- Sentinel errors
- Error design for packages
- Error boundaries
- Retriable vs fatal errors

---

## 18. Testing
- `testing` package
- Table-driven tests
- Subtests (`t.Run`)
- Benchmarks
- `go test -bench`
- Fuzzing (Go fuzz support)
- Mocking via small interfaces
- Deterministic concurrency tests
- Avoiding flaky tests

---

## 19. Filesystem & OS
- `os.Open`, `os.Create`
- `io/fs`
- Temp files
- Atomic file writes
- Path handling
- File permissions
- Signal handling basics

---

## 20. Database & IO Patterns
- `database/sql`
- Connection pooling
- Context-aware queries
- Transactions
- Null handling
- Streaming large result sets
- JSON streaming (`json.Decoder`)
- Memory-safe IO pipelines

---

## 21. HTTP and JSON
- `net/http` handlers
- Server lifecycle
- Middleware chaining
- JSON marshal/unmarshal
- `json:"omitempty"`
- Validation patterns
- Streaming responses
- Error envelope design

---

# 🔴 ADVANCED - Strong + Production Level

## 22. Advanced Concurrency Patterns
- Worker pools
- Fan-in / fan-out
- Pipelines
- Semaphore pattern
- Backpressure control
- Cascading cancellation
- Rate limiting

---

## 23. Memory Model
- Escape analysis (`-gcflags="-m"`)
- Stack vs heap
- Inlining
- Write barriers
- Memory visibility guarantees
- Happens-before relationships
- False sharing

---

## 24. Garbage Collector
- Tracing GC model
- STW phases
- Allocation patterns
- Reducing GC pressure
- Object lifetimes

---

## 25. Generics
- Type parameters
- Constraints
- Comparable
- Generic helpers
- Interfaces + generics interplay
- Performance implications

---

## 26. Reflection
- `reflect.Type`
- `reflect.Value`
- When reflection is justified
- Performance tradeoffs
- Alternatives to reflection

---

## 27. Profiling & Observability
- `pprof` (CPU, heap, goroutine)
- `trace`
- Exposing `/debug/pprof`
- Benchmark-driven optimization
- Identifying memory leaks
- Production observability basics

---

## 28. Go in Production
- Project layout
- Clean Architecture vs pragmatic layout
- Graceful shutdown
- Configuration (env, files)
- Structured logging
- Dependency injection (manual, fx, wire)
- Health checks
- Readiness/liveness probes
- Backward compatibility strategies

---

## 29. Go + Infrastructure
- Cross-compilation
- CGO basics
- Docker multi-stage builds
- Static binaries
- Distroless containers
- CI pipelines
- Versioning and release automation

---

# Outcome of This Track

After completing this roadmap, a senior engineer should be able to:

- Design idiomatic Go APIs
- Write concurrency-safe systems
- Avoid common Go memory and slice pitfalls
- Diagnose performance and race issues
- Build production-ready services
- Reason about Go’s runtime and memory model