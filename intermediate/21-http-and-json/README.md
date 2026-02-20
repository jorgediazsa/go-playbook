# 21 - HTTP and JSON (Realistic Handler Design)

Most Go services are `net/http` servers. The mistakes that hurt in production are usually:
- inconsistent error responses
- poor request validation
- ignoring cancellation
- middleware ordering bugs

---

## Mental model

- Handlers are functions: `ServeHTTP(w, r)`.
- Requests carry a Context (`r.Context()`).
- Responses are streamed; headers must be written before the body.

---

## 1) JSON decoding and validation

Use `json.Decoder` instead of `json.Unmarshal` for request bodies:
- handles streams
- allows strict decoding

Strict decoding pattern:
- `dec.DisallowUnknownFields()`
- decode once
- ensure EOF

Validation:
- validate required fields
- validate numeric ranges
- return structured errors

---

## 2) Encoding responses

Use `json.Encoder` and set headers:
- `Content-Type: application/json`
- status code first

Pitfalls:
- writing headers after writing body
- writing multiple JSON objects without framing

---

## 3) Middleware

Middleware is just a function that wraps a handler.

Common middleware:
- request ID
- auth
- logging
- panic recovery

Pitfalls:
- order matters
- leaking internal errors to clients

---

## 4) Cancellation

If the client disconnects, `r.Context()` is cancelled.
Handlers that do expensive work must stop early when context is done.

---

## 5) Error envelopes

In production, standardize errors:

```json
{ "error": { "code": "invalid_argument", "message": "..." } }
```

Keep internal details out of external responses.

---

## Common interview traps
- Not knowing how to test handlers (`httptest`)
- Not handling unknown JSON fields
- Forgetting to set headers / status correctly

---

## Production checklist
- strict JSON decoding when API requires it
- consistent error schema
- middleware ordering documented
- handlers respect request cancellation

---

## Exercises
These exercises validate:
- strict decoding + validation
- middleware chaining and ordering
- consistent JSON error envelopes
- context-aware handler behavior
