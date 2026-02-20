package memorymodel

import (
	"strings"
)

// Context: Escape Analysis
// You are building a high-throughput log formatter. It parses millions of
// raw log lines per second into structured objects.
//
// Why this matters: Every time a variable escapes to the heap, the Garbage
// Collector has to track it and eventually clean it up. At 1,000,000 logs/sec,
// small allocations destroy CPU cache locality and trigger constant GC pauses.
//
// Requirements:
// 1. Run the benchmark: `go test -bench BenchmarkParseLog -benchmem`
//    Notice that it allocates heavily per operation.
// 2. You can also run `go build -gcflags="-m" ex01_escape.go` to see the compiler
//    decisions (e.g., "escapes to heap").
// 3. Refactor `ParseLog` and `FormatLog` so that they execute with ZERO allocations
//    (0 allocs/op).
//    - Hint: Do not return pointers to small structs if you don't need to.
//    - Hint: Return values directly, passing them by value on the stack.

type LogRecord struct {
	Level   string
	Message string
}

func ParseLog(raw string) *LogRecord {
	// BUG: Returning a pointer to a locally declared struct forces it to escape
	// to the heap, because the struct outlives the function's stack frame.
	// TODO: Refactor everything to use value types (`LogRecord` instead of `*LogRecord`)

	parts := strings.SplitN(raw, ":", 2)
	if len(parts) != 2 {
		return &LogRecord{Level: "UNKNOWN", Message: raw}
	}

	rec := &LogRecord{
		Level:   parts[0],
		Message: parts[1],
	}
	return rec
}

func FormatLog(rec *LogRecord) string {
	// Let's assume this formatting avoids allocations through magic, but the
	// struct allocation itself in `ParseLog` is what we are fixing.
	return "[" + rec.Level + "] " + rec.Message
}
