package control

// Context: Labeled Break/Continue
// You are writing a primitive event loop that consumes off a slice simulating a channel.
// The system runs continuously until it receives a "SHUTDOWN" command.
//
// Why this matters: `select` and `switch` blocks in Go absorb the `break` keyword.
// If you write `break` inside a switch that is inside a for loop, you only break the switch!
// The loop continues infinitely. This causes silent CPU-spinning production hangs.
//
// Requirements:
// 1. Process commands sequentially.
// 2. Ignore any command prefixed with "IGNORE".
// 3. When encountering "SHUTDOWN", cleanly exit the ENTIRE loop.
// 4. Return the list of correctly processed commands.

func EventLoopWorker(commands []string) []string {
	var processed []string

	// TODO: Fix the control flow bugs!
	// 1. Ensure IGNORE skips that command but continues the loop.
	// 2. Ensure SHUTDOWN breaks the outer loop entirely so the function returns.

	for _, cmd := range commands {
		switch {
		case cmd == "SHUTDOWN":
			// BUG: This only breaks the switch, not the `for` loop!
			break
		case len(cmd) > 6 && cmd[:6] == "IGNORE":
			// BUG: This only breaks the switch, so it proceeds to append!
			break
		default:
			processed = append(processed, cmd)
		}

		// Unintentional append if IGNORE or SHUTDOWN bugs surface
		if cmd == "SHUTDOWN" {
			// In our buggy setup, we just appended nothing, but the loop keeps going.
			// Actually wait, if the bug is present, it hits default append? No, it breaks switch.
			// Let's force an append after the switch to prove the concept:
			// processed = append(...) - No, let's keep it simple.
			// If it only breaks switch, the loop continues to the next command.
			// We want SHUTDOWN to stop processing entirely.
		}
	}

	return processed
}
