package generics

import (
	"testing"
)

type FakeUser struct {
	Name string
	Age  int
}

func TestFilterGeneric(t *testing.T) {
	/* TODO: Uncomment this test after refactoring ex02_pipeline.go!

	// Test 1: Ints
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })

	if len(evens) != 3 || evens[0] != 2 {
		t.Fatalf("Failed to filter ints")
	}

	// Test 2: Structs
	users := []FakeUser{
		{"Alice", 25},
		{"Bob", 17},
		{"Charlie", 30},
	}

	adults := Filter(users, func(u FakeUser) bool { return u.Age >= 18 })

	if len(adults) != 2 || adults[0].Name != "Alice" {
		t.Fatalf("Failed to filter structs! Got: %v", adults)
	}
	*/
}
