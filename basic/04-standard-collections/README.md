# 04 - Standard Collections

This module covers Arrays, Slices, Maps, and Structs, with a severe focus on the memory mechanics and concurrency constraints that frequently cause production outages in Go.

## 1. Slice Aliasing and Reallocation

A slice in Go is a 24-byte header containing a pointer to a backing array, a length, and a capacity. 

### The Production Danger

When you pass a slice to a function, the 24-byte header is copied by value, but both headers point to the **same backing array**. If the function modifies existing elements, the caller sees the change.

However, if the function uses `append` and exceeds the capacity of the backing array, Go allocates a **new** backing array for that specific slice header. The caller's slice header still points to the old backing array. This is called "slice disjointing" and causes data loss.

### Idiomatic Solution

If a function intends to append to a slice, it MUST return the slice to the caller.

```go
// CORRECT:
func AddItem(items []string, item string) []string {
    return append(items, item)
}
// Caller: items = AddItem(items, "new")
```

---

## 2. Memory Retention from Subslicing (Memory Leaks)

When you slice a slice (`newSlice := bigSlice[:5]`), the new slice header points to the exact same backing array as `bigSlice`. 

### The Production Danger

If `bigSlice` was a 10MB file read into memory, and you only want to keep the first 50 bytes (`bigSlice[:50]`), the Garbage Collector CANNOT free the 10MB array because `newSlice` still holds a pointer to it. This causes massive memory leaks in long-running services (like parsers or network sniffers).

### Idiomatic Solution

Create a new slice and `copy` the data you want to keep. This allow the large array to be garbage collected.

```go
// CORRECT:
smallBuf := make([]byte, 5)
copy(smallBuf, bigSlice[:5])
// bigSlice goes out of scope and is GC'd. smallBuf is backed by a 5-byte array.
```

---

## 3. Map Concurrency Restrictions

Maps in Go are **not concurrency-safe**. If one goroutine writes to a map while another goroutine reads from or writes to the same map, the runtime will instantly crash with a fatal `fatal error: concurrent map read and map write`. You cannot recover from this panic.

### Idiomatic Solution

Use a `sync.RWMutex` to guard the map, or use `sync.Map` for highly specific read-heavy cache workloads.

---

## 4. Struct Comparability

Not all structs can be compared using `==`. A struct is only comparable if all of its fields are comparable. Examples of incomparable types: slices, maps, functions, and channels.

### The Production Danger

If you use a struct as a map key, it MUST be comparable. If you later add a slice field to that struct, your code will stop compiling. If you add a field with an interface holding a dynamic slice, it will compile but **panic at runtime** when used as a map key.

---

## Exercises

- `ex01_slices.go`
- `ex02_slice_leaks.go`
- `ex03_map_safety.go`
- `ex04_struct_comparability.go`
