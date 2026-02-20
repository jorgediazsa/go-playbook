# 24 - The Garbage Collector

Go uses a concurrent, non-generational, mark-and-sweep Garbage Collector (GC). Understanding its behavior is critical for building systems with predictable latency.

---

## 1. Stop The World (STW)

During a GC cycle, Go must briefly pause all running goroutines to turn on the "write barrier" and verify root pointers. This is the **Stop The World** (STW) phase.
Modern Go STW pauses are incredibly fast (usually < 1ms), but if you allocate memory constantly, you trigger the GC constantly, burning up to 25% of your total CPU budget just on garbage collection!

---

## 2. Object Retention (Memory Leaks in Go)

In C/C++, you leak memory by forgetting to call `free()`.
In Go, you leak memory by accidentally keeping a reference to an object, preventing the GC from ever cleaning it up.

### The Production Danger: Slice Retention
If you read a 1GB file into a `[]byte`, and then create a subslice of the first 10 bytes (`file[:10]`), and save that subslice in a global cache, the **entire 1GB backing array cannot be garbage collected**. The small subslice retains the huge underlying array.

### Idiomatic Solution
If you need a small piece of a huge slice, `copy()` it into a brand new, identically sized slice, and let the huge slice fall out of scope.

### The Production Danger: Unbounded Caches
If you put objects into a `map[string]interface{}` indefinitely, they will never be GC'd. Production caches MUST have eviction policies (TTL, LRU) and maximum capacity limits.

---

## 3. High-Throughput Pooling (`sync.Pool`)

If your server handles 10,000 JSON requests per second, and each request allocates a `bytes.Buffer`, you will trigger the GC frantically.

`sync.Pool` allows you to reuse allocated objects across goroutines. 

### The Production Concept
When a server finishes a request, it `Put()`s the buffer back into the pool.
When a new request starts, it `Get()`s a buffer. If the pool is empty, it allocates a new one.
**Crucially**, the Go Garbage Collector natively understands `sync.Pool`. During a GC sweep, it might safely wipe out unused pool objects if memory is low.

### The Production Danger
When you `Get()` an object from a pool, it is dirty (it contains the data from the previous request). You MUST `Reset()` the object before using it, otherwise you will leak data (or PII) between different users' requests!

---

## Validating GC Behavior

- You cannot deterministically trigger the GC in unit tests without forcing it via `runtime.GC()`.
- Real debugging uses `GODEBUG=gctrace=1 go run main.go` to print GC cycle statistics to stdout.
- Real profiling uses `pprof` heap profiles to see *where* memory was allocated.

---

## Exercises

- `ex01_retention.go`
- `ex02_pooling.go`
- `ex03_eviction.go`
