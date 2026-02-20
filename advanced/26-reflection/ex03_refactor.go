package reflection

// Context: Replacing Reflection with Interfaces & Generics
// Your junior engineer wrote a massive logging utility using `reflect` to dynamically
// extract IDs from ANY struct passed into the logger.
//
// Why this matters: Using `reflect` here makes the logger 100x slower and circumvents
// compile-time type safety. If someone renames `ID` to `UserId`, `reflect` silently fails
// at runtime.
//
// Requirements:
// 1. You MUST completely DELETE `ExtractIDReflection` and never use it.
// 2. Define an interface `Identifiable` with a method `GetID() string`.
// 3. Write a new function `LogID` that takes an `Identifiable` (or a Generic `[T Identifiable]`).
// 4. Update the `Order` and `Device` structs to implement `Identifiable`.

import (
	"reflect"
)

type Order struct {
	ID     string
	Amount int
}

type Device struct {
	ID  string
	Mac string
}

// BUG: Terrible, slow, unsafe reflection code.
// TODO: Delete this entirely and replace with an interface-driven approach!
func ExtractIDReflection(obj any) string {
	val := reflect.ValueOf(obj)

	// Dereference pointer if necessary
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Struct {
		field := val.FieldByName("ID")
		if field.IsValid() && field.Kind() == reflect.String {
			return field.String()
		}
	}
	return "UNKNOWN"
}
