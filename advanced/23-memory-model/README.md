# 23 - The Go Memory Model

To write high-performance, concurrent Go code, you must understand how Go manages memory natively (Stack vs Heap) and how CPU architectures interact with that memory (Caches, False Sharing, Visibility).

---

## 1. Escape Analysis (Stack vs Heap)

Go has two main areas of memory:
- **The Stack:** Extremely fast, self-cleaning. Variables are pushed/popped as functions are called.
- **The Heap:** Slower. Memory must be requested, managed, and eventually cleaned up by the Garbage Collector (GC).

**The Golden Rule:** If the compiler cannot mathematically prove that a variable's lifetime is entirely contained within the function that created it, the variable "escapes" to the Heap.

### What Causes Escapes?
1. Returning pointers from functions (`return &x`).
2. Storing pointers in structs or slices that live longer than the function.
3. Passing pointers to `interface{}` arguments (like `fmt.Println(&x)`).
4. Closures capturing variables by reference.

### How to Validate
Run `go build -gcflags="-m"` to see the compiler's escape analysis decisions.
Run benchmarks with `-benchmem` to observe `B/op` (bytes per operation) and `allocs/op` (heap allocations).

---

## 2. Visibility and Happens-Before

In modern multi-core CPUs, each core has its own L1/L2 cache. If Core 1 writes to a variable, Core 2 does **not** instantly see it.

Go's Memory Model defines a strict set of **"happens-before"** guarantees. 
For example:
- "A send on a channel *happens before* the corresponding receive from that channel completes."
- "The `Unlock` of a `sync.Mutex` *happens before* any execution of the corresponding `Lock` returns."

If you read and write the same variable from two goroutines *without* a happens-before edge (like a Mutex or a Channel), you have a **Data Race**, and the read is completely undefined behavior (you might read garbage, old data, or half-written data).

### How to Validate
Never guess. Always use `go test -race ./...`.

---

## 3. False Sharing

Modern CPUs read memory in blocks called **Cache Lines** (usually 64 bytes).
If two completely independent variables (`var A int64` and `var B int64`) happen to sit next to each other in memory, they will be loaded into the **same** cache line.

If Goroutine 1 constantly modifies `A` on Core 1, and Goroutine 2 constantly modifies `B` on Core 2, the CPU hardware will constantly invalidate and bounce that single cache line back and forth between the cores, destroying performance. This is called **False Sharing**.

### Idiomatic Solution
Pad your structs! Insert empty arrays (e.g., `_ [56]byte`) between heavily contended, independent fields to force them into different cache lines.

---

## Exercises

- `ex01_escape.go` (Bench test available)
- `ex02_visibility.go`
- `ex03_false_sharing.go` (Bench test available)
