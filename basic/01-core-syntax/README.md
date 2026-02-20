# 01 - Core Syntax

This module covers the core syntax of Go specifically tailored for senior engineers. Rather than explaining `if/else` or basic loops, this module focuses on **Go-specific pitfalls**, mental models, and production constraints that catch experienced developers off-guard when transitioning to Go.

## 1. Zero Values vs Missing Values

In Go, variables declared without an explicit initial value are given their **zero value** (e.g., `0` for ints, `false` for bools, `""` for strings, `nil` for pointers/slices/maps/interfaces).

### The Production Danger

In APIs (e.g., JSON handling) a missing field usually signifies "do not update", while an explicitly provided `false` or `0` means "update this field to the zero value." Go's standard unmarshaling assigns missing fields their zero values. This causes a devastating bug where a partial update request silently zeroes out other fields.

```go
// BAD PATTERN:
type PatchRequest struct {
    IsAdmin bool // An omission in JSON leaves this false. Did the user omit it, or explicitly say false?
    Age     int  // Omission gives 0. Is the user 0 years old, or was it omitted?
}
```

### Idiomatic Solutions

1. **Use Pointers:** `IsAdmin *bool`. A `nil` pointer means omitted. A non-nil pointer means provided. (Caveat: puts pressure on the Garbage Collector via heap allocations and escape analysis).
2. **Use SQL-Null style Wrappers:** `type OptionalBool struct { Value bool; Valid bool }`. This achieves the same goal safely completely on the stack.

---

## 2. Type Aliases vs Defined Types

Go 1.9 introduced **Type Aliases** (`type T1 = T2`), primarily to assist in code refactoring. Unfortunately, many developers confuse them with **Defined Types** (`type T1 T2`), using aliases for domain modeling.

### Core Mental Model

- **Defined Type:** `type UserID string`. Creates a *new, distinct* type. You cannot accidentally pass a `ProductID` to a function expecting a `UserID`.
- **Type Alias:** `type UserID = string`. This creates an alternative name, but the compiler still treats it as exactly `string`. A `UserID` and `ProductID` are 100% interchangeable and bypass all type safety.

### Common Production Bug

```go
type UserID = string
type OrgID = string

func DeleteOrg(u UserID, o OrgID) { ... }

// A developer calls this:
DeleteOrg(orgID, userID) // Compiles perfectly! Catastrophic bug.
```

Always use Defined Types for domain modeling identifiers.

---

## 3. Untyped Constants Behavior

Constants in Go (`const x = 5`) are conceptually different from variables. If you don't explicitly type a constant, it remains **untyped**. Go's compiler represents untyped numeric constants with at least 256 bits of precision.

### Why this matters

If you eagerly type a constant, you lock it into that type's boundaries.

```go
// WRONG
const Scale int64 = 1_000_000_000_000_000_000
const Mult  int64 = 5_000_000_000_000_000_000
// Computing (Scale * Mult) / Divisor will overflow an int64 during math!

// IDIOMATIC
const Scale = 1_000_000_000_000_000_000
const Mult  = 5_000_000_000_000_000_000
// Since they are untyped, the compiler does the intermediate multiplication using 256-bit math,
// then perfectly downsizes the final evaluated result to fit into whatever type you assign it to.
```

---

## 4. Subtle Integer Overflows in Conversions

Unlike languages that might throw exceptions on math overflows, Go **silently truncates**.

```go
var bigInt int64 = 5_000_000_000
var smallUint uint32 = uint32(bigInt) // This DOES NOT PANIC. It silently truncates.
```

### Common Interview Trap & Production Bug

When translating payloads across boundaries (e.g., reading an `int64` from gRPC and putting it into a legacy `uint32` database column), direct casting will lead to silent, hard-to-track data corruption. 

**Idiomatic Go:** Always explicitly check the boundaries of the target type *before* casting, returning an error if it falls outside the `<min, max>` threshold of that target type.

---

## Exercises

- `ex01_zero_values.go`
- `ex02_types.go`
- `ex03_constants.go`
- `ex04_overflow.go`
