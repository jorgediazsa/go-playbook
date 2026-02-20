package httpjson

import (
	"net/http"
)

// Context: Safe HTTP Clients
// You are building a Webhook sender. It fires HTTP POST requests to third-party
// URLs provided by your users.
//
// Why this matters: `http.DefaultClient` has NO timeouts. If a malicious user
// provides a URL to a server that accepts the connection but never responds (a tarpit),
// your goroutine will block infinitely, leaking memory and exhausting file descriptors.
//
// Requirements:
// 1. Refactor `NewWebhookClient` to NOT return `http.DefaultClient`.
// 2. Return a custom `http.Client` with a strict `Timeout` of 5 seconds.
// 3. Configure its `Transport` (`*http.Transport`) to limit `MaxIdleConnsPerHost` to 5.
//    (This prevents connection pool bloat when spamming the same host).

func NewWebhookClient() *http.Client {
	// BUG: The default client is infinitely blocking!
	// TODO: Create a safe, bounded client with strict Timeouts and Transport limits.

	return http.DefaultClient
}
