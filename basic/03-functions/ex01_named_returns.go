package functions

import "fmt"

// Context: Named Returns and Defer
// You are writing a database transaction wrapper. If any panic occurs during
// the transaction, or if the transaction function returns an error, you must
// ensure that the returned error is properly wrapped with context indicating
// that the transaction failed.
//
// Why this matters: In Go, deferred functions execute after `return` is evaluated
// but before the function exits. They can read and modify named return values.
// This is the idiomatic way to add contextual information to errors universally
// across a function that has many early-return paths.
//
// Requirements:
// 1. `ExecuteTx` takes a callback that might return an error or panic.
// 2. We use a named return `err error`.
// 3. Implement a `defer` block that catches panics (translating them to errors)
//    AND wraps any non-nil error with "tx failed: %w".
// 4. Do not change the function signature or the direct return statements.

func ExecuteTx(fn func() error) (err error) {
	// TODO: Review this defer block, understand why it works, then re-implement it yourself.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("tx failed: panic: %v", r)
			return
		}
		if err != nil {
			err = fmt.Errorf("tx failed: %w", err)
		}
	}()

	// Simulated logic
	err = fn()

	// If fn() fails, err is returned. But it lacks the "tx failed:" context unless
	// the defer block decorates it.
	return err
}
