package core

// Context: Untyped Constants Behavior
//
// You are writing a precise rate calculation module for a crypto-exchange.
// You need to compute a scaling factor using massive multipliers.
// If you use typed constants (e.g., int64), intermediate calculations will overflow during
// compilation or execution, even if the final result fits nicely into the destination type.
//
// Why this matters: Go's untyped constants afford at least 256 bits of precision during
// compile time. By eagerly typing them, you rob the compiler of the ability to do high-precision
// intermediate math.
//
// Requirements:
// 1. Fix the constants so that they don't overflow during compilation when multiplied.
// 2. Implement `ComputeScale` to return the result of `(Base * Multiplier) / Divisor`.
// 3. Keep the return type of ComputeScale as int64.

// TODO: These typed constants prevent compilation of the math below, or cause overflow. Clean them up!
const Base int64 = 1000000000000000000
const Multiplier int64 = 5000000000000000000
const Divisor int64 = 1000000000000000000

func ComputeScale() int64 {
	// TODO: Return (Base * Multiplier) / Divisor
	// The test expects exactly 5000000000000000000.
	// You might see compilation overflow errors until you fix the consts.
	return 0
}
