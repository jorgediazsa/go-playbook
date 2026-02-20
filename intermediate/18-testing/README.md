# 18 - Testing (Table-Driven, Deterministic, and Fast)

Go’s testing culture is pragmatic: table-driven tests, small interfaces for mocking, and benchmarks when performance matters.

---

## Mental model

- Tests are documentation of behavior.
- Prefer **deterministic** tests over time-based flakiness.
- Write code that is easy to test by designing boundaries.

---

## 1) Table-driven tests

Pattern:
- define a slice of cases
- iterate, run subtests

Benefits:
- consistent coverage
- easy to add edge cases

---

## 2) Subtests and parallelism

- `t.Run` creates nested test scopes.
- `t.Parallel()` can speed up tests, but can introduce shared-state races.

Pitfalls:
- capturing loop variables in subtests
- parallel tests that mutate shared globals

---

## 3) Mocking with small interfaces

Instead of heavy mocking frameworks:
- design a boundary interface with 1–3 methods
- inject it

Example:

```go
type Clock interface { Now() time.Time }
```

This enables deterministic tests.

---

## 4) Benchmarks

Benchmarks answer: “is this fast enough?”

- `go test -bench .`
- avoid allocations (`b.ReportAllocs()`)

Pitfalls:
- benchmarking unrealistic inputs
- mixing IO with CPU benchmarks

---

## 5) Fuzzing (overview)

Go fuzzing helps find edge cases by generating inputs:

- `go test -fuzz=.`

Use for:
- parsers
- encoders/decoders
- validation logic

---

## Common interview traps
- Writing brittle tests that depend on ordering of maps
- Using time.Sleep instead of synchronization
- Not knowing how to run a specific test (`-run`)

---

## Production checklist
- CI runs tests with `-race` where appropriate
- Tests are deterministic and fast
- Boundaries are designed for testability
- Benchmarks exist for hot paths

---

## Exercises
These exercises focus on:
- designing code to be testable
- refactoring for small interface boundaries
- deterministic concurrency tests
- adding benchmark/fuzz stubs where it makes sense
