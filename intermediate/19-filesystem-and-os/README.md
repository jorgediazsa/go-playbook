# 19 - Filesystem and OS (Correct IO and Portability)

Filesystem code is where “works on my machine” bugs are born. In Go, most IO is explicit and low-level — which is good, but you must be disciplined.

---

## Mental model

- Always handle errors.
- Always close resources.
- Design for portability (Windows vs Unix differences).

---

## 1) io/fs and testability

`io/fs` enables abstraction over real filesystems.

Benefits:
- can test using `fstest.MapFS`
- can run logic against embedded filesystems

---

## 2) Atomic writes

Common safe pattern:
- write to temp file
- fsync if required
- rename over target

Pitfalls:
- rename semantics differ across platforms
- atomicity is per-filesystem boundary

---

## 3) Path handling

- Use `path/filepath` for OS paths.
- Use `path` for slash-separated paths (URLs).

Pitfall:
- hardcoding `/` breaks on Windows.

---

## 4) Permissions

Be explicit about file modes and ownership assumptions.

---

## 5) Signals (design for testability)

Instead of calling `signal.Notify` everywhere:
- isolate signal handling in one place
- inject a channel into components

---

## Common interview traps
- Forgetting to close files
- Not handling partial writes
- Using `path` instead of `filepath`

---

## Production checklist
- Use atomic write patterns for configs/state
- Keep path handling portable
- Use io/fs abstractions for testability
- Avoid leaking file descriptors

---

## Exercises
These exercises validate:
- atomic write semantics
- portable path behavior
- filesystem abstraction via io/fs
- signal-aware shutdown patterns with injected channels
