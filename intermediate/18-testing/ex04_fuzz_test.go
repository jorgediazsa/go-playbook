package testingadv

import "testing"

// Context: Native Fuzzing
// You are testing the `ParseMetric` function from ex01_table_driven.go.
//
// Why this matters: You wrote 4 hardcoded test cases in your table. But what if
// a user sends `":::::"`, or `"\x00CPU"`, or a 1GB string?
// Go 1.18 introduces `Fuzz()`, which intelligently mutates corpus inputs to
// trigger panics, index-out-of-bounds, or infinite loops you never predicted.
//
// Requirements:
// 1. Look at this Fuzz target. It feeds random strings to `ParseMetric`.
// 2. We assert an INVARIANT: `ParseMetric` must NEVER panic, regardless
//    of what garbage string is fed into it.
// 3. To run it, you would execute: `go test -fuzz=FuzzParseMetric -fuzztime=10s`.
//
// Note: You do not need to implement anything here. This file serves purely
// to demonstrate structural fuzzing invariants for your learning.
// But you must ensure `ex01_table_driven.go` passes standard tests for this to run.

func FuzzParseMetric(f *testing.F) {
	// 1. Add known "seed" inputs to guide the fuzzer.
	f.Add("cpu:5")
	f.Add("mem:")
	f.Add(":100")
	f.Add(":::::")

	// 2. The Fuzz target gives us a structurally random string.
	// You cannot assert `if name == "cpu"` because you don't know the input!
	f.Fuzz(func(t *testing.T, input string) {
		// INVARIANT: This function must never panic (crash the process).
		name, val, err := ParseMetric(input)

		// INVARIANT: If there is no exact colon, it must return an error.
		// (Fuzzers are great at finding strings with 0 array length logic bombs).
		if err == nil {
			// If it succeeds, the name and length must theoretically make sense.
			if val < 0 {
				t.Fatalf("ParseMetric created a negative length from input: %q", input)
			}
			_ = name
		}
	})
}
