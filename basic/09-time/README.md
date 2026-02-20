# 09 - Time

Handling time correctly in distributed systems is notoriously difficult. Go provides a robust `time` package, but it has several hidden performance cliffs and logical gotchas.

## 1. `time.After` vs `time.NewTimer` (Memory Leaks)

The `time.After(duration)` function returns a channel that emits the current time after the duration has elapsed.

### The Production Danger

Under the hood, `time.After` creates a Timer that **cannot be garbage collected** until the timer actually fires. If you use `time.After(1 * time.Hour)` inside a `select` block that usually resolves in 50 milliseconds via another channel, that 1-hour timer sits in memory for a full hour! In a high-throughput system (e.g., 5,000 requests per second), this will rapidly consume gigabytes of RAM and crash the server with an Out Of Memory (OOM) error.

### Idiomatic Solution

For long-running timeouts inside highly trafficked `select` blocks, ALWAYS use `time.NewTimer`. This allows you to explicitly `Stop()` the timer and allow the GC to reclaim it immediately.

```go
// CORRECT:
timer := time.NewTimer(1 * time.Hour)
defer timer.Stop() // Prevents the leak!

select {
case msg := <-ch:
    return msg
case <-timer.C:
    return timeoutErr
}
```

---

## 2. Monotonic Clocks vs Wall Clocks

A `time.Time` object in Go contains BOTH a "wall clock" reading (the actual time of day, which can jump forwards or backwards if the server's NTP daemon syncs) AND a "monotonic clock" reading (a constantly ticking hardware counter that NEVER jumps backwards).

### The Production Danger

If you measure the duration of a function by saving `start := time.Now()`, and the server's clock adjusts backwards by 5 seconds while the function is running, `time.Since(start)` will still accurately report the duration because it uses the *monotonic* clock reading embedded inside the `start` variable.

HOWEVER, if you serialize that `start` time (e.g., convert it to JSON or save it to a database), the **monotonic reading is stripped away**. When you load it back, doing `time.Since(loadedStart)` relies purely on the wall clock, and can result in negative durations!

### Idiomatic Solution

Never rely on loaded/deserialized timestamps for strict duration arithmetic if network/NTP syncs could occur in between.

---

## 3. Parsing and Formatting

Go uses a magical reference date for parsing and formatting: `Mon Jan 2 15:04:05 MST 2006`.
(Think of it as 1, 2, 3, 4, 5, 6).

If you want to format a date to `YYYY-MM-DD`, you must use `time.Format("2006-01-02")`.

---

## Exercises

- `ex01_timer_leaks.go`
- `ex02_monotonic_time.go`
