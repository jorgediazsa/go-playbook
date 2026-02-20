package ctxexercises

// Context: Context Values and Layered Architecture
// You have a layered architecture: HTTP Handler -> Service -> Database.
//
// Why this matters: Tracing and observability require passing a "Trace ID"
// (or Request ID) through every single layer, from the incoming HTTP request
// down to the final SQL query so you can search logs effectively.
// `context.WithValue` is the idiomatic way to pass this metadata through
// layers that don't need to know about it.
//
// Rule: ALWAYS define custom, unexported types for your Context Keys to prevent
// collisions with other packages (e.g., `type contextKey int`, then `const TraceKey contextKey = 1`).
//
// Requirements:
// 1. Define a private custom type for context keys (e.g. `type traceKeyType string`).
// 2. Wrap `ctx` using `context.WithValue` in `Handler` to inject the `traceID`.
// 3. Extract the `traceID` from `ctx` in `DatabaseLayer` and return it.

import "context"

// TODO: Define a custom, unexported type for your Context Key here.
// e.g. type ctxKey string
// const TraceIDKey ctxKey = "traceID"

func Handler(traceID string) string {
	ctx := context.Background()

	// BUG: The context is currently empty.
	// TODO: Inject `traceID` into a new context using `context.WithValue`.
	// Be sure to use your custom key type!

	return ServiceLayer(ctx)
}

func ServiceLayer(ctx context.Context) string {
	// The service layer doesn't care about the traceID. It just passes the context down.
	return DatabaseLayer(ctx)
}

func DatabaseLayer(ctx context.Context) string {
	// BUG: Currently returns an empty string.
	// TODO: Extract the trace ID using `ctx.Value(YourKey)`.
	// You will need to type-assert the result to a string.

	return ""
}
