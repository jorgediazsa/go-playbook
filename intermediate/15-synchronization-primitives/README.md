# 15 - Synchronization Primitives (Choosing the Right Tool)

Go gives you multiple synchronization tools. The hard part is not knowing they exist — it’s choosing the right one and using it correctly.

---

## Mental model

- Use **mutexes** to protect shared mutable state.
- Use **WaitGroup** to wait for goroutines to complete.
- Use **Cond** when you need to wait for a condition without busy-spinning.
- Use **atomic** for simple counters/flags where lock-free is appropriate.

Rule:
> Correctness first. Optimize only after measuring.

---

## 1) sync.Mutex and sync.RWMutex

- `Mutex` is simplest and often fastest in real systems.
- `RWMutex` helps only when reads dominate and contention is high.

Pitfalls:
- copying a struct containing a mutex is a bug
- forgetting to unlock (use `defer` carefully in hot paths)
- holding locks across IO

---

## 2) sync.Once

`Once` guarantees a function runs once. It does not naturally model “run once unless error” — for that you need your own state.

---

## 3) sync.Map

`sync.Map` is specialized for patterns with many readers and infrequent writes. It is not a general replacement for a map + mutex.

---

## 4) sync.Cond

Use a Cond when:
- goroutines need to wait until a condition becomes true
- you want to avoid polling

Pattern:
- condition checked in a loop
- signal/broadcast after state changes

---

## 5) sync/atomic

Atomics are for:
- counters
- state flags

Pitfalls:
- atomics do not make composite structures safe
- memory ordering matters (Go provides strong enough guarantees for typical patterns, but you still need disciplined design)

---

## Common interview traps
- Using RWMutex everywhere “because reads”
- Using sync.Map without understanding workload
- Using atomics for complex invariants

---

## Production checklist
- Protect maps with mutexes unless you have a proven reason not to
- Keep lock scopes small; avoid locks across IO
- Prefer simple primitives; avoid cleverness
- Run `-race` in CI for concurrency-heavy packages

---

## Exercises
The exercises focus on:
- correct locking and invariants
- choosing Mutex vs RWMutex vs atomic
- Cond-based coordination
- safe initialization patterns
