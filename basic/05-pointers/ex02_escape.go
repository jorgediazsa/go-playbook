package pointers

// Context: Escape Implications
// You are building an ultra-fast event emitter that creates and discards millions
// of events per second. The GC CPU usage is spiking to 40%.
//
// Why this matters: Returning a pointer to a struct created inside a function
// forces the Go compiler to allocate that struct on the Heap. The GC must later
// sweep and free it. Returning a struct by value keeps it on the Stack, which
// is automatically cleaned up essentially for free.
//
// Requirements:
// 1. You are modifying a critical path function.
// 2. Refactor `CreateEvent` to return a value instead of a pointer.
// 3. Fix the `ProcessStream` function to accommodate the new value-based return type.

type Event struct {
	ID      string
	Payload string
}

// BUG: Returning a pointer forces this small struct onto the heap.
// At 50,000 requests per second, this creates massive GC pressure.
// TODO: Refactor `CreateEvent` to return `Event` (the value), not `*Event`.
func CreateEvent(id, payload string) *Event {
	e := Event{
		ID:      id,
		Payload: payload,
	}
	return &e // Pointers to local variables escape to the heap!
}

// TODO: Refactor this processing function to accept the Value, or take the address of the returned value.
func ProcessStream() int {
	eventsProcessed := 0

	for i := 0; i < 1000; i++ {
		// If you changed CreateEvent to return a value, this next line will error!
		evt := CreateEvent("id-123", "data")

		// Simulate doing something that requires pointer access, e.g. a legacy func
		// that we can't change. You can take the address of `evt` safely here because
		// it's contained inside this loop's stack frame.
		legacyEmit(evt)
		eventsProcessed++
	}

	return eventsProcessed
}

// Simulate a legacy function we cannot change.
func legacyEmit(e *Event) {
	_ = e.ID
}
