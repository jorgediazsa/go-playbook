# 20 - Database and IO Patterns (Streaming and Safe APIs)

This topic focuses on production patterns that show up constantly:
- `database/sql` usage patterns (without depending on a real driver)
- streaming large data through `io.Reader/io.Writer`
- JSON streaming with `json.Decoder`

---

## Mental model

- `database/sql` is an abstraction over drivers and connection pooling.
- Rows must be closed.
- Context must be propagated to IO-bound work.
- Streaming avoids loading everything into memory.

---

## 1) database/sql patterns

Key points:
- A `*sql.DB` is a pool, not a single connection.
- Always `defer rows.Close()`.
- Always check `rows.Err()` after iteration.
- Use context-aware methods (`QueryContext`, `ExecContext`).

Because this repo avoids external DB drivers, exercises model testable patterns by:
- defining small interfaces around query/exec
- injecting fakes/mocks

---

## 2) Transactions

Patterns:
- begin
- do work
- commit
- rollback on error

Pitfalls:
- forgetting rollback on early return
- committing partial state

---

## 3) Null handling

SQL NULLs require explicit handling (`sql.NullString`, etc.).

Pitfall:
- scanning NULL into a non-nullable Go type

---

## 4) Streaming IO

Prefer streaming when inputs can be large:
- transform via Reader → Writer
- avoid `io.ReadAll` unless size is bounded

---

## 5) Streaming JSON

Use `json.Decoder` for:
- large arrays/streams
- incremental processing

Pitfalls:
- ignoring unknown fields in strict APIs
- buffering entire payloads unnecessarily

---

## Common interview traps
- Treating `*sql.DB` as a connection
- Forgetting `rows.Close()` / `rows.Err()`
- Using `io.ReadAll` everywhere

---

## Production checklist
- Context passed into DB calls
- Rows always closed
- Transactions rollback on error
- Streaming used for large payloads

---

## Exercises
These exercises enforce:
- correct transaction semantics via early returns
- testable DB boundaries using small interfaces
- streaming transformations without full buffering
- json.Decoder usage patterns
