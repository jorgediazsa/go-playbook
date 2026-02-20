package reflection

import (
	"reflect"
)

// Context: Custom Struct Validation via Tags
// You are building an API router. When a JSON payload is unmarshaled into a struct,
// you want to automatically validate it based on struct tags.
//
// Example: `type User struct { Name string \`validate:"required"\` }`
//
// Why this matters: This is the exact implementation pattern behind libraries like
// `go-playground/validator` and standard `encoding/json`.
//
// Requirements:
// 1. Refactor `ValidateStruct` to use `reflect.TypeOf()` and `reflect.ValueOf()`.
// 2. Iterate through all the fields of the struct.
// 3. For any field with the tag `validate:"required"`:
//    - If the field is a `string` and it is empty (`""`), return an error.
//    - If the field is an `int` and it is `0`, return an error.
// 4. If the input `s` is NOT a struct (or a pointer to a struct), return an error.

func ValidateStruct(s any) error {
	// BUG: It currently does nothing and always returns nil.
	// TODO: Use `reflect` to dynamically inspect the fields of `s`.

	val := reflect.ValueOf(s)

	// Hint: If s is a pointer, you need `val.Elem()` to get the underlying struct.
	// If it's not a struct, return an error.

	_ = val // Silence unused warning

	return nil
}
