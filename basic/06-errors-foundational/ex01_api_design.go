package errorsfunc

import "errors"

// Context: Error API Design and Early Returns
// You are refactoring a legacy configuration loader. The previous engineer came from
// an object-oriented background and used `panic` to abort parsing when something
// went wrong, wrapping the "happy path" heavily.
//
// Why this matters: Go does not use exceptions for flow control. If this library
// is consumed by a web server, a `panic` here would crash the entire HTTP server
// if not recovered. We must handle errors as values and use early returns to
// maintain the horizontal "line of sight".
//
// Requirements:
// 1. Refactor `LoadConfig` to return `(Config, error)`.
// 2. Remove all panics. Return sentinel errors instead (`ErrEmptyPayload` or `ErrInvalidFormat`).
// 3. Un-nest the happy path. The final return should be the successful configuration.

type Config struct {
	Timeout int
	Retries int
}

var ErrEmptyPayload = errors.New("empty payload")
var ErrInvalidFormat = errors.New("invalid format")


// BUG: Legacy signature used panic and deep nesting.
// TODO: Refactor to return (Config, error), remove panics, and use early returns.
// NOTE: This stub compiles but intentionally fails tests until implemented.
func LoadConfig(payload string) (Config, error) {
	// TODO: Implement according to the spec above.
	return Config{}, errors.New("TODO: implement LoadConfig")
}
