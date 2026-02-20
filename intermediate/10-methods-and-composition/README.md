# 10 - Methods and Composition (Idiomatic Go OOP)

Go does not have classes or inheritance. It has **methods**, **method sets**, **composition via embedding**, and **implicit interfaces**. If you get these right, you’ll write Go that is easy to test, easy to refactor, and hard to misuse.

---

## Mental model

- A **method** is just a function with a receiver: `func (r T) M() {}` or `func (r *T) M() {}`.
- Receivers create a **method set**:
  - `T` has methods with receiver `T`.
  - `*T` has methods with receiver `T` **and** `*T`.
- **Interfaces are satisfied implicitly** by method sets.

This means receiver choice is not just about mutation/perf — it changes which interfaces a type can satisfy.

---

## 1) Value vs pointer receivers

### Rules of thumb
- Use a **pointer receiver** when:
  - the method mutates receiver state
  - the receiver contains a `sync.Mutex` / `sync.Once` / `atomic.*` or other non-copyable state
  - copying the receiver would be expensive or semantically wrong
- Use a **value receiver** when:
  - the receiver is small and immutable (or treated as such)
  - you want to preserve value semantics

### Pitfall: accidental copies
If you put a value receiver on a method that should mutate state, you’ll mutate a copy.

### Pitfall: interface satisfaction
If an interface requires `M()`, and `M` is defined on `*T`, then **`T` does not satisfy the interface**, only `*T` does.

---

## 2) Embedding (composition) and promotion

Embedding a type `X` into `Y`:

```go
type Y struct {
    X
}
```

- Promotes `X`’s methods to `Y` (you can call `y.M()` if `X` has `M`).
- `Y` can override promoted methods by defining its own `M`.

### Shadowing vs overriding
- If `Y` defines `M`, calls to `y.M()` use `Y.M`.
- You can still call the embedded method explicitly: `y.X.M()`.

This is commonly used to build decorators / adapters.

---

## 3) Implicit interfaces and small interfaces

Go’s most powerful abstraction is the **small interface**:

```go
type Closer interface { Close() error }
```

Design APIs around behavior, not concrete types.

### Pitfall: “interface pollution”
Large interfaces create tight coupling and make mocking painful.

---

## 4) io.Reader / io.Writer patterns

These are the core contracts for streaming:

- `io.Reader`: `Read(p []byte) (n int, err error)`
- `io.Writer`: `Write(p []byte) (n int, err error)`

### Common production patterns
- Decorators: wrap a reader/writer to add behavior (logging, redaction, counting)
- Transformers: normalize data while streaming (newline normalization, UTF-8 checks)
- Adapters: expose domain APIs as Reader/Writer to plug into stdlib

### Contract gotchas
- A Reader may return `n > 0` and `err == io.EOF` in the same call.
- A Writer may write `n < len(p)` with a non-nil error.

---

## Common interview traps
- Confusing method sets (why `T` doesn’t satisfy an interface but `*T` does)
- Embedding misconceptions (“inheritance”) vs composition realities
- Returning concrete types instead of small interfaces

---

## Production checklist
- Receiver choice: pointer for mutable/non-copyable state
- Prefer small interfaces for boundaries (testability)
- Use embedding intentionally (decorators/adapters)
- Prefer io.Reader/io.Writer for streaming (avoid loading everything into memory)

---

## Exercises
Solve the exercises in this folder. They are designed to force:
- correct receiver choice
- method set / interface satisfaction correctness
- embedding + promotion/shadowing behavior
- Reader/Writer contract correctness
