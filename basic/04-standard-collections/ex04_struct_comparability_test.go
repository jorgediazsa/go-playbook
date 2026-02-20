package collections

import "testing"

func TestIsDuplicate(t *testing.T) {
	ResetCache()

	// If the user hasn't fixed the struct, this will not compile (if they fixed the map type)
	// OR it will panic (if they left it as interface{}).
	// We use an explicit defer recover to catch the panic if they didn't fix the struct
	// but tried to run the test.

	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Code Panicked! You tried to use an incomparable struct as a map key: %v", r)
		}
	}()

	// The user must modify RequestSignature to not use []string.
	// Wait, if we change the struct definition, this test file won't compile unless we also
	// change how we construct `RequestSignature` here.
	// Therefore, the test MUST construct it in a way that matches the expected solution,
	// or we provide a constructor function that the user has to implement.
	//
	// Let's change the exercise constraints slightly via the test:
	// We expect the student to change Headers from `[]string` to `string`.

	// To make this test compile regardless of what the user does (before they fix it),
	// we won't directly construct the struct in the test if it causes a compiler error.
	// But `Headers []string` vs `Headers string` WILL cause a compiler error.
	// We will supply a `CreateSignature(method, path string, headers []string) RequestSignature`
	// function in the main file that the user must implement.
}
