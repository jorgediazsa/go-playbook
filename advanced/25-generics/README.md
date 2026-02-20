# 25 - Generics (Type Parameters)

Go 1.18 introduced Generics (Type Parameters). Generics allow you to write functions and data structures that work across multiple types without sacrificing compile-time type safety or resorting to slow `interface{}` + reflection tricks.

---

## 1. Type Constraints and `comparable`

To use a generic type `T`, you must constrain it. 
- `[T any]`: The type can be literally anything. You cannot use operators like `==` or `+` on `any`.
- `[T comparable]`: The type must support `==` and `!=`. This is required if `T` is going to be used as a `map` key.
- `[T constraints.Ordered]`: (From `golang.org/x/exp/constraints`) The type supports `<`, `>`, `<=`, `>=`, etc.

### The Production Concept (Data Structures)
Before Generics, building a `Set` required writing `IntSet`, `StringSet`, `FloatSet`, or using `map[interface{}]bool` (which requires clunky type assertions). 
With Generics, you can build a single `Set[T comparable]` that is perfectly type-safe.

---

## 2. Generic Pipelines and Higher-Order Functions

Generics shine when building utility pipelines or algorithms (like `Map`, `Filter`, `Reduce`). 

### The Production Danger (Overengineering)
Do NOT use generics just to save 3 lines of code. If an `interface` naturally describes the *behavior* of an object (e.g. `io.Reader`), use the interface. 
Use Generics when you need to enforce strict *Data relationships* (e.g., "This function takes a slice of T and must return a slice of the exact same T").

---

## 3. The Result/Option Pattern

Go's multiple-return-value idiom (`res, err := DoWork()`) is standard and should almost always be respected. 
However, in highly concurrent pipelines (like channels), you cannot send two values through a single channel. You must send a wrapper struct. 

Generics allow us to build a strongly-typed `Result[T]` wrapper, similar to Rust's `Result` or Scala's `Try`, without resorting to `interface{}`.

---

## Validating Generics

- Compiling is 90% of the battle. If a generic function compiles, it is type-safe.
- There is a VERY slight performance overhead at compile time (monomorphization) and sometimes small runtime penalties depending on the Go version, but they are dramatically faster than reflection.

---

## Exercises

- `ex01_set.go`
- `ex02_pipeline.go`
- `ex03_result.go`
