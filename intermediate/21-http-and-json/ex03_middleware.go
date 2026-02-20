package httpjson

import "net/http"

// Context: HTTP Middleware
// You are exposing a public API `/health`. Before any request reaches the handler,
// you want to ensure the request contains the header `X-API-Key: secret`.
//
// Why this matters: Instead of putting authentication checks inside EVERY single
// handler function, idiomatic Go wraps handlers in "Middlewares".
// A middleware is a function that takes an `http.Handler` and returns a *new*
// `http.Handler` (usually via `http.HandlerFunc`).
//
// Requirements:
// 1. Refactor `RequireAuth` to return an `http.Handler` that wraps `next`.
// 2. The returned handler should inspect the HTTP request.
// 3. If `r.Header.Get("X-API-Key")` equals `"secret"`, call `next.ServeHTTP(w, r)`.
// 4. If it does not, return an HTTP 401 Unauthorized status and STOP the chain.
//    (Use `http.Error(w, "unauthorized", http.StatusUnauthorized)`).

func RequireAuth(next http.Handler) http.Handler {
	// BUG: This middleware is currently a no-op that just immediately calls the next handler.
	// TODO: Wrap this in an `http.HandlerFunc` closure so you can intercept the request.

	return next
}

// Provided
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HEALTHY"))
}
