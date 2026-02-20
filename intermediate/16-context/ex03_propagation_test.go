package ctxexercises

import "testing"

func TestContextPropagation(t *testing.T) {
	result := Handler("req-abc-123")

	if result != "req-abc-123" {
		t.Fatalf("FAILED: Expected trace ID 'req-abc-123' to propagate to the DatabaseLayer, got %q", result)
	}

	// NOTE: Static analysis tools often warn if you use untyped strings as context keys,
	// because `ctx.Value("traceID")` in your package could collide with `ctx.Value("traceID")`
	// setup by a third-party auth middleware. Using a custom type guarantees uniqueness.
}
