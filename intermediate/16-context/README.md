# 16 - context.Context (Cancellation and Propagation)

`context.Context` is Go’s mechanism for cancellation, deadlines, and request-scoped values. It is central in servers and in any system that calls external resources.

---

## Mental model

- Context is **immutable**: you derive new contexts from parents.
- Context flows **down** the call stack.
- Cancellation is **cooperative**: code must check `Done()` or use context-aware APIs.

Rule:
> Accept a Context as your first parameter when the operation can block or be cancelled.

---

## 1) Creating contexts

- `context.Background()` for top-level roots
- `context.WithCancel(parent)` for manual cancellation
- `context.WithTimeout(parent, d)` for timeouts
- `context.WithDeadline(parent, t)` for fixed deadlines

Always call the returned cancel function to release resources:

```go
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
defer cancel()
```

---

## 2) Propagation

A derived context inherits:
- cancellation
- deadline
- values

This is how HTTP request cancellation propagates to DB calls and downstream services.

---

## 3) Context in net/http

- Use `r.Context()` inside handlers.
- If the client disconnects, the request context is cancelled.
- Long-running work must respect cancellation.

Pitfall:
- Starting goroutines in handlers that outlive the request without explicit lifecycle management.

---

## 4) Values

Context values are for request-scoped metadata (request IDs), not for passing business dependencies.

Pitfalls:
- storing large values
- using context as a parameter bag

---

## Common interview traps
- Forgetting to call cancel
- Not propagating context to downstream calls
- Misusing values for dependency injection

---

## Production checklist
- Context is first param in IO-bound APIs
- Cancel funcs are always called
- Handlers respect `r.Context()`
- Timeouts are explicit for external calls

---

## Exercises
These exercises validate:
- correct cancellation handling
- propagation across layered APIs
- correct behavior in HTTP handlers under client disconnect/timeout
