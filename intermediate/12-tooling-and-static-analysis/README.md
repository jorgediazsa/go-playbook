# 12 - Tooling and Static Analysis (How Go Stays Maintainable)

Go’s tooling is part of the language. In production, you don’t “just write code” — you rely on `go test`, `-race`, `go vet`, build tags, and reproducible builds.

---

## Mental model

- `go` is a build system, test runner, formatter, and dependency manager.
- The default toolchain encourages consistency and avoids “configuration explosion”.

---

## Core commands (practical)

### Build
- `go build ./...` compiles packages.
- `go test ./...` compiles + runs tests.

### Inspect
- `go env` shows toolchain configuration.
- `go list ./...` lists packages in module.
- `go list -deps` helps understand dependency graphs.

### Vet and race
- `go vet` catches suspicious constructs (printf mismatch, unreachable code patterns, shadowing in some cases).
- `go test -race` detects data races dynamically.

---

## Build tags

Build tags allow multiple implementations behind the same API:

```go
//go:build linux
```

Use cases:
- OS-specific logic
- “safe vs fast” implementations
- optional debug instrumentation

Pitfalls:
- You can accidentally exclude all files for a package.
- Tests must be explicit about tags when needed.

---

## go:generate

`go generate` is not magic — it runs commands you specify.

Use cases:
- embedding static assets
- generating small lookup tables
- generating version metadata

Pitfalls:
- generators must be deterministic
- keep generated files small and reviewable

---

## pprof basics

`pprof` is for answering:
- where CPU time goes
- where memory is allocated
- which goroutines are blocked

Typical workflow:
- instrument program / server with pprof endpoints
- capture profiles under load
- inspect with `go tool pprof`

Pitfall: optimize only after measuring.

---

## staticcheck (overview)

Staticcheck is a commonly used linter suite in Go teams. This repo does not require it, but you should know what it flags:
- ineffective assignments
- misuse of time APIs
- subtle correctness issues

---

## Common interview traps
- Not knowing how to run “all tests” (`go test ./...`)
- Forgetting `-race` for concurrent code
- Misunderstanding build tags selection

---

## Production checklist
- CI runs `go test ./...` and `go test -race` for concurrency-heavy packages
- `go vet` runs in CI
- Build tags are documented and minimal
- Generators are deterministic and checked in appropriately

---

## Exercises
These exercises enforce:
- correct usage of build tags and generation
- tooling-driven correctness constraints
- practical profiling literacy (without fragile assertions on profile output)
