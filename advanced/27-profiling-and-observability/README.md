# 27 - Profiling and Observability

Performance engineering requires absolute proof, not guesswork. Go natively provides `pprof` (CPU, Head, Mutex, Block profiling) and `trace` (Goroutine scheduling and garbage collection).

---

## 1. Using `net/http/pprof`

For long-running web servers, you can attach the `pprof` handlers to your default HTTP multiplexer effortlessly.

```go
import _ "net/http/pprof"

func main() {
    // Starts an HTTP server on port 6060 exposing /debug/pprof/
    go http.ListenAndServe("localhost:6060", nil) 
}
```

### How to analyze (Commands to Run)

1. **CPU Profile (What functions are burning CPU?)**
   `go tool pprof -http=:8080 "http://localhost:6060/debug/pprof/profile?seconds=5"`
   - *Look for:* Wide boxes in the flame graph. Try to avoid optimizing functions that only take 1% of total time.

2. **Heap Profile (Where is memory being allocated / retained?)**
   `go tool pprof -http=:8080 "http://localhost:6060/debug/pprof/heap"`
   - *Look for:* `inuse_space` for memory leaks. `alloc_space` for high allocation pressure triggering the GC.

3. **Goroutine Profile (Are we leaking goroutines?)**
   `go tool pprof -http=:8080 "http://localhost:6060/debug/pprof/goroutine"`
   - *Look for:* A massive count of goroutines stuck in `runtime.gopark` or `select`.

---

## 2. Using `runtime/trace`

While `pprof` samples the program every ~10ms, `trace` records *every single scheduling event* (when a goroutine starts, stops, blocks, or yields).

```go
f, _ := os.Create("trace.out")
trace.Start(f)
defer trace.Stop()
```

### How to analyze (Commands to Run)

`go tool trace trace.out`
- *Look for:* Massive "Stop The World" (STW) gaps caused by Garbage Collection.
- *Look for:* Goroutines spending 90% of their time "Waiting" instead of "Executing" (indicating channel starvation or mutex contention).

---

## 3. Observability (Logs, Metrics, Traces)

In a production environment, you don't always have access to a live `pprof` endpoint. 

- **Metrics (Prometheus):** Track counts (requests) and distributions (latency).
- **Structured Logging (slog / zap):** Emitting logs as JSON so systems like Datadog/Splunk can query them.
- **Distributed Tracing (OpenTelemetry):** Injecting `trace_id`s into HTTP Headers and Go Contexts so you can track a request across 10 different microservices.

---

## Exercises

- `ex01_pprof.go` (A severely bottlenecked and leaky component)
- `cmd/loader/main.go` (A test runner you can execute and profile live!)
