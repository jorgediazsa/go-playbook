package generics

import (
	"testing"
)

// We can't write a rigorous test if the struct name or signature changes dramatically,
// but if the student implements `Set[T comparable]`, we can instantiate it here.

// BUG: To compile the test file out-of-the-box before the student finishes the exercise,
// these tests are heavily commented or rely on reflection if we wanted to be tricky.
// Instead, we provide the test code. When the user refactors `ex01_set.go`,
// they must also uncomment this test to prove it works!

func TestGenericSet(t *testing.T) {
	/* TODO: Uncomment this test after refactoring ex01_set.go!

	intSet := NewSet[int]()
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(1)

	if !intSet.Contains(1) || !intSet.Contains(2) {
		t.Fatalf("Int set failed contains check")
	}

	stringSet := NewSet[string]()
	stringSet.Add("hello")

	if !stringSet.Contains("hello") {
		t.Fatalf("String set failed")
	}
	*/
}
