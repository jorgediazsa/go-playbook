package generics

// Context: Generic Functional Pipelines
// You are building a data-processing framework. Customers need to filter slices made
// of different structs (e.g. `Filter([]User)`, `Filter([]Order)`).
//
// Why this matters: You cannot pass `[]User` to a function expecting `[]any`.
// Slices are strictly monomorphic. Generics allow you to write a single `Filter`
// function that works on any slice type effortlessly.
//
// Requirements:
// 1. Refactor `Filter` to accept a slice of ANY type `T`.
// 2. The `predicate` function must accept `T` and return a `bool`.
// 3. `Filter` must return a new slice of `T`.

// BUG: This only works for integers.
// TODO: Refactor to func Filter[T any](input []T, predicate func(T) bool) []T
func Filter(input []int, predicate func(int) bool) []int {
	var results []int
	for _, v := range input {
		if predicate(v) {
			results = append(results, v)
		}
	}
	return results
}
