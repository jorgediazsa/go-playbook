package vet

import "fmt"

// Context: The `go vet` tool
// You are writing a logging formatter.
//
// Why this matters: `go build` ensures the syntax is correct. `go vet` catches
// logic errors that correctly compile but are almost certainly bugs (like bad
// Printf formats, unreachable code, or accidental shadowing of variables).
//
// Requirements:
// 1. The code below compiles perfectly! But it has a formatting bug.
// 2. Fix the `fmt.Sprintf` so `go vet` stops complaining.
// 3. (Optional) Run `go vet .` to see the warning before you fix it.

func FormatLog(level string, id int) string {
	// BUG: The verbs provided do not match the arguments!
	// `%d` expects an integer, but gets a string.
	// `%s` expects a string, but gets an integer.

	// TODO: Fix the verbs so `go vet` passes.
	return fmt.Sprintf("Level: %d | RequestID: %s", level, id)
}
