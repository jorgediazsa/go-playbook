package garbagecollector

import (
	"bytes"
	"sync"
)

// Context: High-Throughput Pooling (`sync.Pool`)
// You are building a blazing fast web server. Every request needs a `bytes.Buffer`
// to asynchronously render HTML templates before writing to the network.
//
// Why this matters: Allocating a `bytes.Buffer` 10,000 times a second creates
// tremendous GC pressure. A `sync.Pool` safely reuses these buffers across
// Goroutines smoothly handling spikes without crashing the GC.
//
// Requirements:
// 1. Initialize `bufferPool` with a `New` function that returns a `*bytes.Buffer`.
// 2. Refactor `RenderTemplate` to get a buffer from the pool, use it, and then
//    `defer bufferPool.Put(buf)` to return it.
// 3. CRITICAL: You MUST call `buf.Reset()` before doing anything with it!
//    If you don't reset, User B will receive the HTML from User A appended to theirs!

var bufferPool = sync.Pool{
	// BUG: Needs a New constructor
	// TODO: Add `New: func() any { return new(bytes.Buffer) }`
}

func RenderTemplate(content string) string {
	// BUG: Allocates a brand new buffer every single time.
	// TODO: 1. buf := bufferPool.Get().(*bytes.Buffer)
	// TODO: 2. defer bufferPool.Put(buf)
	// TODO: 3. buf.Reset() // CRITICAL: Erase old data!

	buf := new(bytes.Buffer)

	buf.WriteString("<html><body>")
	buf.WriteString(content)
	buf.WriteString("</body></html>")

	return buf.String()
}
