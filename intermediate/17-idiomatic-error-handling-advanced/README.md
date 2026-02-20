# 17 - Idiomatic Error Handling (Advanced)

Go errors are values. At senior level, the important part is designing error behavior as part of your API contract.

---

## Mental model

- Errors are for expected failure modes.
- Panics are for programmer bugs / unrecoverable invariants.
- Errors should be:
  - machine-checkable when the caller must branch
  - human-readable when logged

---

## 1) Wrapping and inspection

- Use `%w` with `fmt.Errorf` to wrap.
- Use `errors.Is` to match sentinel errors.
- Use `errors.As` to extract typed errors.

Guideline:
> Wrap errors when adding context, not when changing meaning.

---

## 2) Sentinel errors

Sentinel errors are package-level vars:

```go
var ErrNotFound = errors.New("not found")
```

Use when:
- callers need stable identity checks (`errors.Is(err, ErrNotFound)`)

Pitfall:
- don’t compare errors by string

---

## 3) Typed errors

Use typed errors when callers need structured information:

```go
type ValidationError struct { Field string; Reason string }
```

Design:
- keep fields small
- implement `Error()`
- use `errors.As` for extraction

---

## 4) Error boundaries

At boundaries (HTTP handlers, CLIs):
- classify errors
- map to user-facing responses
- log with context

Avoid leaking internal error details to external clients.

---

## 5) Retriable vs fatal

Design errors to support retry logic:
- transient network errors
- rate limits
- optimistic concurrency

But never retry blindly without bounds and backoff.

---

## Common interview traps
- Comparing errors by string
- Not using `%w` → breaks Is/As
- Over-wrapping every error with no additional context

---

## Production checklist
- Errors are part of API design
- Wrapping preserves identity via `%w`
- Callers can reliably branch using Is/As
- Boundaries translate internal errors appropriately

---

## Exercises
These exercises require:
- correct wrapping chains
- sentinel + typed error design
- stable error classification
- preserving error identity across layers
