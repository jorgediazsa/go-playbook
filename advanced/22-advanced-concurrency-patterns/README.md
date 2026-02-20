# 22 - Advanced Concurrency Patterns

While basic goroutines and channels are easy to write, composing them into resilient, production-grade systems requires rigorous structural patterns.

---

## 1. Worker Pools and Bounded Concurrency

Unbounded concurrency (`go process(item)` in a massive loop) works until you exhaust file descriptors, memory, or the target API's rate limit. A **Bounded Worker Pool** limits the maximum number of active goroutines.

### The Production Concept (Draining and Cancellation)
A production worker pool must handle three things perfectly:
1. **Cancellation:** If the parent context is canceled, workers must abort in-flight work immediately.
2. **Error Propagation:** If one worker encounters a fatal error, it must be able to cancel the *others* and return the error to the caller.
3. **Draining (Safe Shutdown):** When the work queue is closed, workers must finish their current tasks and exit cleanly without leaking.

---

## 2. Pipelines and Backpressure

A pipeline is a series of stages connected by channels. Data flows from a Generator -> Processor -> Sink.

### The Production Danger (Cascading Deadlocks)
If the Sink is slow, the Processor's output channel fills up. The Processor blocks. Then the Generator's output channel fills up. The Generator blocks. This is **Backpressure**.
But what if the Sink encounters an error and returns early? The Processor and Generator are now permanently blocked (leaked) trying to send to unread channels!

### Idiomatic Solution
Every stage of a pipeline MUST select on a `ctx.Done()` channel. If any downstream stage fails, it must cancel the context, which instantly cascades up the pipeline, unblocking all senders so they can exit gracefully.

---

## 3. Semaphores and Rate Limiting

Sometimes you don't need a dedicated pool of workers; you just need to ensure no more than *N* goroutines access a resource simultaneously.

- **Semaphore:** A buffered channel (`make(chan struct{}, N)`) used as a token bucket. Acquire a token before doing work; release it after.
- **Rate Limiting:** `golang.org/x/time/rate` provides a robust token-bucket rate limiter. (Though in these exercises, we will explore the concepts manually via channels/tickers).

---

## Validating Concurrency

- Always run `go test -race ./...` to ensure no hidden data races.
- If a test hangs indefinitely, you have a leak or a deadlock.
- **Common Interview Trap:** Failing to close a channel when the sender is done, causing the receiver to range over it forever.

---

## Exercises

- `ex01_worker_pool.go`
- `ex02_pipeline.go`
- `ex03_rate_limit.go`
