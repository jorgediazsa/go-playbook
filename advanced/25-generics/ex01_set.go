package generics

// Context: Generic Data Structures (Set)
// You need a mathematical Set data structure that guarantees uniqueness.
// You need it to work for `int`, `string`, and `uuid.UUID`.
//
// Why this matters: Historically, Go developers either copy-pasted 3 versions
// of the Set, or used `interface{}` which lost type safety and required casting.
// Generics allow exactly ONE implementation that is 100% type-safe.
//
// Requirements:
// 1. Refactor `Set` to use a type parameter `T`.
// 2. Determine the correct constraint for `T`. (Hint: To be a map key, it must
//    be `comparable`).
// 3. Update `NewSet`, `Add`, and `Contains`.

// BUG: This set is hardcoded to strings.
// TODO: Make it Generic: `Set[T constraint]`
type Set struct {
	items map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[string]struct{}),
	}
}

func (s *Set) Add(item string) {
	s.items[item] = struct{}{}
}

func (s *Set) Contains(item string) bool {
	_, ok := s.items[item]
	return ok
}
