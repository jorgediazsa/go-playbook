package composition

// Title: Embedded Components with Promoted overriding Methods
//
// Context: You are building a metrics framework. You have a `BaseMetric` struct that
// handles ID generation and timestamping. You want a `CounterMetric` to gain those
// features without copying the code, so you embed `BaseMetric` inside `CounterMetric`.
//
// Why this matters: In Go, embedding is composition, NOT inheritance.
// While `CounterMetric` gains the `Record()` method of `BaseMetric` via promotion,
// if `CounterMetric` defines its OWN `Record()` method, it shadows (overrides) the base method.
// But critically: if `BaseMetric.Record()` calls another method internally, it will NEVER
// magically call `CounterMetric`'s method (no runtime polymorphism for embedded structs).
//
// Requirements:
// 1. `CounterMetric` currently embeds a pointer to `BaseMetric` (`*BaseMetric`).
// 2. Define a `Record()` method on `CounterMetric` that:
//      a) Increments its own `Count` field.
//      b) Explicitly calls the embedded `BaseMetric.Record()` to get the timestamp logging.
// 3. Fix the instantiation in `NewCounter` so it doesn't panic on a nil pointer dereference
//    when trying to access the embedded `*BaseMetric`.

import "time"

type BaseMetric struct {
	ID        string
	LastEvent time.Time
}

func (b *BaseMetric) Record() {
	b.LastEvent = time.Now()
}

type CounterMetric struct {
	*BaseMetric // Embedded pointer. Promotes all *BaseMetric methods.
	Count       int
}

// TODO: Define the `Record()` method for `*CounterMetric`.
// It should increment `Count`, and then call the underlying `BaseMetric.Record()`.

// NewCounter acts as a constructor.
func NewCounter(id string) *CounterMetric {
	// BUG: If we just return `&CounterMetric{}` leaving the embedded pointer nil,
	// calling `Record()` promoted from `BaseMetric` will panic!
	// TODO: Properly initialize the embedded struct.
	return &CounterMetric{}
}
