# 28 - Go in Production

Building a Go application for production requires moving beyond standard library tutorials and focusing on operational resilience: structured logging, dynamic configuration, graceful shutdown, and backward compatibility.

---

## 1. Project Layout and Dependency Injection

While Go does not force a specific folder structure, idiomatic production apps separate their domain logic from their delivery mechanisms (HTTP/gRPC/CLI) and storage layers.

**Dependency Injection (DI):** 
Do not use global database connections (`var DB *sql.DB`). You must pass dependencies into your handlers.
- *Manual DI:* `NewHandler(db *sql.DB, logger *slog.Logger)`
- *Automated DI:* Uber's `fx` or Google's `wire` are fantastic for massive codebases, but start with manual DI.

---

## 2. Structured Logging (`log/slog`)

`fmt.Printf` and `log.Println` are useless in heavy production environments. Logs must be easily parseable by aggregators like Datadog, Elasticsearch, or Splunk.
Go 1.21 introduced `log/slog`, native structured logging!

```go
slog.Info("user created", slog.String("user_id", "123"), slog.Duration("latency", 5*time.Millisecond))
// Output: {"time":"...", "level":"INFO", "msg":"user created", "user_id":"123", "latency":5000000}
```
*(Third-party alternatives: `uber-go/zap` or `rs/zerolog` for absolute maximum performance).*

---

## 3. Graceful Shutdown

When you deploy a new version of your app (e.g., via Kubernetes), the old app receives a `SIGTERM` signal.
If your app dies instantly, inflight HTTP requests are severed, and users see `502 Bad Gateway` errors.

### The Production Concept
You must intercept `SIGTERM`, tell the HTTP server to stop accepting *new* connections, and wait for *existing* active requests to finish (with a timeout) before finally exiting.

---

## 4. API Versioning & Backward Compatibility

APIs evolve. If you delete a field from a JSON response, you will break mobile apps that haven't updated yet.
- Never remove fields. Add new ones.
- When behavior radically changes, route to a new version: `/v1/users` -> `/v2/users`.

---

## Exercises

- `ex01_shutdown.go`
- `ex02_structured.go`
- `ex03_compatibility.go`
