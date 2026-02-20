# 07 - Defer, Panic, Recover

Go relies on three keywords to manage function cleanup and catastrophic failure: `defer`, `panic`, and `recover`.

## 1. Defer Order (LIFO)

When you `defer` a function call, Go pushes it onto a stack. When the surrounding function returns, the deferred functions are executed in **Last-In, First-Out (LIFO)** order.

### The Production Danger

If you acquire multiple resources (e.g., a database connection, then a file lock), you must release them in the exact reverse order. If you release the DB connection before the lock, any deferred cleanup logic that briefly relies on the DB will fail. 

**Idiomatic Go:** Immediately `defer` the cleanup of a resource on the very next line after successfully acquiring it. This naturally guarantees LIFO ordered cleanup without mental arithmetic.

---

## 2. Panic Boundaries and Recover

A `panic` is meant strictly for "impossible" conditions where the program cannot safely continue (e.g., out of memory, or a nil pointer dereference).

### The Production Danger

If a background goroutine panics, it crashes the **entire process**, including the main HTTP server and all other goroutines. One rogue goroutine takes down the ship.

### Idiomatic Solution

HTTP frameworks (like Gin or standard `net/http`) automatically recover from panics strictly within the boundary of a single request handler. However, if you spawn your own goroutines, you MUST define your own panic recovery boundaries at the top of that goroutine.

```go
go func() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Background worker panicked! Recovered: %v", r)
        }
    }()
    
    // Risk of panic here...
}()
```

### Recover Placement Rule

`recover()` ONLY works when called *directly* inside a deferred function. If you call `recover()` inside a nested function inside the defer, or outside a defer entirely, it will return `nil` and the panic will continue tearing down the process.

---

## Exercises

- `ex01_defer_order.go`
- `ex02_panic_boundaries.go`
