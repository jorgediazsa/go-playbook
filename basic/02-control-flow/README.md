# 02 - Control Flow

This module covers Go's control flow mechanics. Because Go only has three core flow control mechanisms (`if`, `switch`, `for`), there are overloaded semantics to them that catch senior engineers off guard. 

## 1. Switch Without Expression

Unlike languages where `switch` exclusively matches a single variable against constants (e.g., C or Java), a `switch` statement in Go without a condition evaluates the `case` expressions for truthiness, executing the first one that evaluates to `true`.

### The Production Danger

Building deeply nested `if/else if/else` chains in Go is generally considered unidiomatic and visually taxing.

```go
// BAD PATTERN:
if conditionA {
    // ...
} else if conditionB && !conditionC {
    // ...
} else if conditionD {
    // ...
} else {
    // ...
}
```

### Idiomatic Solution

Use the "switch true" pattern (officially known as a switch without an expression).

```go
switch { // implicitly `switch true`
case conditionA:
    // ...
case conditionB && !conditionC:
    // ...
case conditionD:
    // ...
default:
    // ...
}
```
This aligns the left margin, making complex routing logic far more readable.

---

## 2. Loop Variable Capture (Go Pre-1.22 vs Post-1.22)

In Go, `for` loops declare a single loop variable that is reused on every iteration. 

### The Production Danger (Pre-Go 1.22)

If you capture the loop variable inside a closure (e.g., kicking off a goroutine) or take its memory address (`&val`), you capture the *reference* to the single reused memory location. When the goroutines run, the loop has typically finished, so all goroutines end up seeing the final value of the loop.

```go
// PRE-1.22 BUG:
var out []*string
for _, name := range []string{"Alice", "Bob"} {
    out = append(out, &name) // Both pointers in `out` point to "Bob"!
}
```

### The Fix

Historically, the fix was to shadow the variable: `name := name`. 
**Note:** As of Go 1.22, this behaves identically to languages like C#, where a new variable instance is created per iteration. However, for compatibility with legacy systems, understanding the shadowing idiom is still required of senior engineers.

---

## 3. Labeled Break and Continue

Go lacks an arbitrary `goto` for arbitrary jumps across scopes, but it heavily leverages **labeled statements** for jumping out of nested loops or `select` blocks.

### The Production Danger

In large applications (like nested parsers or retry mechanisms wrapped around `select` statements), a bare `break` only exits the immediate innermost block (e.g., the `switch` or `select`), throwing the application into an infinite outer loop.

```go
// BUG:
for {
    select {
    case msg := <-ch:
        if msg == "STOP" {
            break // Only breaks the `select`! The `for` loop continues forever!
        }
    }
}
```

### Idiomatic Solution

Use labeled breaks.

```go
Loop:
for {
    select {
    case msg := <-ch:
        if msg == "STOP" {
            break Loop // Exits the `for` loop entirely.
        }
    }
}
```

---

## 4. Range Over Maps and Strings

### Map Non-Determinism
Go intentionally randomizes the starting iteration order of a map. You cannot rely on keys coming out in the order they were inserted.

### String Iteration
Iterating over a string using a classical `for i := 0; i < len(str); i++` loops over *bytes*. Iterating using `for i, runeVal := range str` loops over *runes* (Unicode code points), decoding UTF-8 on the fly. 
`i` holds the byte offset (which might jump by up to 4 bytes per iteration!).

---

## Exercises

- `ex01_switch.go`
- `ex02_loop_capture.go`
- `ex03_labeled_break.go`
