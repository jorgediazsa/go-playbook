# 13 - Goroutines, Lifecycle, and Leaks

Goroutines are cheap, lightweight threads managed by the Go runtime. Cheap is not free: **leaks** and **races** are two of the most common production failure modes in Go.

---

## Mental model

- A goroutine is started with `go f()`.
- You cannot forcibly stop a goroutine from the outside.
- Therefore: **every goroutine needs a lifecycle plan**.

Rule:
> Never start a goroutine unless you can explain how it stops.

---

## 1) Concurrency vs parallelism

- **Concurrency**: structuring a program to make progress on multiple tasks (interleaving).
- **Parallelism**: executing multiple tasks at the same time on multiple CPU cores.

Go gives you concurrency by default; parallelism depends on `GOMAXPROCS` and runtime scheduling. Correct concurrent code must not rely on ordering, timing, or “it will run in parallel”.

Production implication:
- Concurrency is how you keep systems responsive under IO waits.
- Parallelism is how you scale CPU-bound work — but it’s not guaranteed by writing `go`.

---

## 2) Goroutine leaks

A goroutine leak happens when you start a goroutine that never terminates. Common causes:

- waiting forever (blocking read, lock not released, cond never signaled)
- writing to an unbuffered channel with no receiver
- `time.After` in loops without draining timers (can retain resources)
- forgetting to stop tickers

### Detection strategy
- Use timeouts in tests to avoid hangs.
- Use `go test -run ... -count=100` for stress.
- Use `pprof` goroutine profile in production.

### Prevention patterns
- cancellation (usually via `context.Context`)
- bounded work + timeouts
- explicit stop functions + `WaitGroup`

---

## 3) Shared memory hazards (data races)

A **data race** occurs when:
- two goroutines access the same memory concurrently
- at least one of them writes
- and there is no synchronization

Symptoms range from “wrong sum” to impossible states and crashes (`fatal error: concurrent map writes`).

### The race detector
Always run:

```bash
go test -race ./...
```

When you see races:
- protect shared state with `sync.Mutex` / `RWMutex`
- use `sync/atomic` for simple counters/flags
- avoid sharing mutable state when possible

---

## Common interview traps
- Starting goroutines in handlers without shutdown/cancellation
- Assuming goroutines execute in parallel
- Using maps concurrently without synchronization

---

## Production checklist
- Every background goroutine has a stop path
- Tickers are stopped (`defer ticker.Stop()`)
- Worker pools or executors have bounded queues and shutdown semantics
- CI runs `go test -race` on concurrency packages

---

## Exercises
The exercises in this folder focus on:
- lifecycle correctness (start/stop)
- leak prevention under error paths
- race-proof shared state updates
- deterministic concurrent behavior without relying on CPU count
