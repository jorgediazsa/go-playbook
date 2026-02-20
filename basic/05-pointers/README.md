# 05 - Pointers

This module covers the semantics of pointers in Go, specifically around method receivers, mutation vs. copying, and the performance implications of Escape Analysis.

## 1. Value vs Pointer Receivers

When you attach a method to a struct in Go, you must choose whether the receiver is a value (`func (s Struct)`) or a pointer (`func (s *Struct)`).

### The Production Danger

If you use a **value receiver**, Go creates a full copy of the struct every time the method is called.
1. If the method mutates the struct, the caller **will not see the mutation** because the method mutated its own local copy.
2. If the struct is extremely large, or contains lock primitives (`sync.Mutex`), copying it is disastrous. Copying a mutex invalidates its lock state and triggers data races.

### Idiomatic Solution

- **Use Pointer Receivers (`*T`)** if the method needs to mutate the struct, or if the struct is large / contains a `sync.Mutex`.
- **Use Value Receivers (`T`)** if the struct is small (like a small configuration struct or a custom ID type) and you want to strongly guarantee immutability.

---

## 2. Escape Analysis (Performance Implications)

In languages like C or C++, returning a pointer to a local variable is a catastrophic bug (dangling pointer).
In Go, it is perfectly safe. If the compiler sees that a pointer outlives the function it was created in, it "escapes" the variable to the **heap** instead of allocating it on the **stack**.

### The Production Danger

Heap allocations are expensive. They require locks in the memory allocator and put pressure on the Garbage Collector (GC), which leads to CPU spikes and latency jitter in high-throughput applications. Stack allocations are essentially free (just moving a stack pointer) and instantly cleaned up when the function returns.

```go
// HEAP ALLOCATION (Escapes)
func NewUser() *User {
    u := User{Name: "Alice"} 
    return &u // Pointer escapes! "u" is allocated on the heap.
}

// STACK ALLOCATION (Does not escape)
func GetUser() User {
    u := User{Name: "Alice"}
    return u // Value is copied down the stack. "u" is allocated on the stack.
}
```

### Idiomatic Solution

For small, short-lived structs used entirely within a single request lifecycle, **pass them by value** (`User` instead of `*User`). The CPU cost of copying a 64-byte struct is typically cheaper than the GC cost of managing a heap allocation. Reserve pointers for structs that must be shared/mutated, or are undeniably massive.

---

## Exercises

- `ex01_receivers.go`
- `ex02_escape.go`
