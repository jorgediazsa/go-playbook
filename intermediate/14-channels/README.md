# 14 - Channels (Correctness and Ownership)

Channels are Go’s communication primitive for coordinating goroutines. Most channel bugs in production are not “syntax errors” — they’re ownership and lifecycle mistakes.

---

## Mental model

- A channel is a typed conduit: `chan T`.
- Sends/receives can block.
- Closing is a **signal**: “no more values will ever be sent”.

Ownership rule:
> The goroutine that **sends** values is typically the one that **closes** the channel.

---

## 1) Buffered vs unbuffered

- Unbuffered channels synchronize sender/receiver (handoff).
- Buffered channels decouple sender/receiver up to capacity.

Production implications:
- buffering can reduce contention but can hide deadlocks until buffers fill
- unbuffered channels provide stronger backpressure semantics

---

## 2) Select

`select` lets you wait on multiple channel operations.

Patterns:
- cancellation channel (or context) + work channel
- timeouts using `time.After` (careful in loops)
- `default` case for non-blocking behavior

Pitfall:
- `select` fairness is not guaranteed; do not depend on it.

---

## 3) Closing and ranging

- Receiving from a closed channel yields the zero value plus `ok=false`.
- Ranging over a channel ends when it is closed and drained.

Pitfalls:
- sending on a closed channel panics
- closing a channel twice panics
- closing a channel you don’t own is a correctness bug

---

## 4) Nil channels

A nil channel blocks forever on send/receive. This is sometimes useful in `select` to “disable” a case.

---

## Common interview traps
- Misunderstanding who closes the channel
- Assuming buffered channels prevent deadlocks
- Forgetting to drain or close, causing goroutine leaks

---

## Production checklist
- Establish channel ownership (who closes)
- Ensure shutdown drains work and terminates workers
- Avoid timeouts that allocate endlessly (`time.After` in loops)
- Consider backpressure semantics intentionally (buffering)

---

## Exercises
These exercises emphasize:
- correct closing semantics
- select-based coordination
- avoiding deadlocks and leaks
- directional channels in APIs to enforce ownership
