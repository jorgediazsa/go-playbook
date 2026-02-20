package functions

// Context: Closure Variable Capture
// You are designing an event emitter where users can register callbacks to be executed later.
// A common bug occurs when users register callbacks inside a loop or reuse variables,
// but the callback executes asynchronously, capturing the *reference* to the variable,
// not the *value* it had at the time of registration.
//
// Why this matters: Asynchronous systems rely heavily on closures. Capturing mutable state
// by reference leads to race conditions and logical bugs where callbacks fire with the
// "last known state" rather than the state they were intended for.
//
// Requirements:
// 1. `CreateHandlers` simulates setting up 3 handlers.
// 2. Each handler receives an integer state.
// 3. Fix the state capture bug so that each handler returns its creation index (0, 1, 2)
//    rather than all returning the final mutated state.

func CreateHandlers() []func() int {
	var handlers []func() int

	state := 0

	for i := 0; i < 3; i++ {
		// BUG: The closure captures `state` by reference.
		// Since `state` is mutated below, all closures will see the final value (3)
		// when they are eventually executed.
		// TODO: Fix the closure so it captures the CURRENT value of `state`.

		handlers = append(handlers, func() int {
			return state
		})

		state++
	}

	return handlers
}
