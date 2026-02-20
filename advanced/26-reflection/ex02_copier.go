package reflection

import (
	"errors"
	"reflect"
)

// Context: Safe Reflection Copying
// You are building an ORM (Object Relational Mapper). When you query the DB,
// you receive a generic `map[string]interface{}` and you need to copy its
// values into a provided target struct pointer.
//
// Why this matters: `encoding/json` does this exact thing. You MUST ensure:
// 1. The target is a Pointer.
// 2. The target points to a Struct.
// 3. You only set fields that are Exported (uppercase).
//
// Requirements:
// 1. Refactor `CopyMapToStruct`.
// 2. Ensure `dst` is a pointer to a struct. If not, return an error.
// 3. Iterate through `src` map keys.
// 4. For each key, if `dst` has a field with that exact name, and the field
//    `CanSet()` (is exported), use reflection to set its value: `field.Set(reflect.ValueOf(mapValue))`.
//    (Assume the types match strictly for this exercise).

func CopyMapToStruct(src map[string]any, dst any) error {
	// BUG: Does nothing.
	// TODO: Handle pointer dereferencing with `reflect.ValueOf(dst).Elem()`
	// TODO: Check if it's a struct `val.Kind() == reflect.Struct`
	// TODO: Iterate over map, lookup `val.FieldByName(k)`, check `.IsValid()` and `.CanSet()`.

	val := reflect.ValueOf(dst)
	_ = val // unused

	return errors.New("unimplemented")
}
