package control

import (
	"strings"
)

// Context: Switch without expression
// You're building a highly-custom HTTP routing middleware.
// The incoming requests are complex and must be bucketed into priorities based on
// a mix of header presence, payload size, and user tiers.
//
// Why this matters: Deep nested if/else blocks are an anti-pattern in Go.
// We use `switch` with no expression to create clean, linear evaluation chains.
// Wait till you try to maintain an 8-level `if/else` block from a junior engineer!
//
// Requirements:
// 1. Evaluate the priority of a Request using a single `switch` block (no expression).
// 2. Return Priority based on these disjointed rules (evaluated in order):
//    - If the user is VIP AND payload is < 100 bytes, return "Priority-High"
//    - If the path starts with "/admin", return "Priority-Critical"
//    - If the path starts with "/api" and method is "POST", return "Priority-Medium"
//    - Otherwise, return "Priority-Low"

type Request struct {
	Path        string
	Method      string
	IsVIP       bool
	PayloadSize int
}

func RoutePriority(req Request) string {
	// TODO: Replace the single return below with an expression-less switch that
	// returns the correct priority string.
	// You may use strings.HasPrefix() from the standard library.

	return strings.ToUpper("unimplemented")
}
