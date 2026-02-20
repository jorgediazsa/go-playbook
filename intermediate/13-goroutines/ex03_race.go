package goroutines

// Context: Data Races
// You are building an Aggregator that receives stream chunks of logs from 100
// network connections concurrently. It tracks the total number of bytes received.
//
// Why this matters: `a.TotalBytes += bytes` looks atomic, but underneath it is:
// 1. Read TotalBytes from memory into CPU register.
// 2. Add count to register.
// 3. Write register back to memory.
// If two goroutines do this at the precise same moment, one of the updates will
// be silently overwritten (lost update), or the memory reading becomes garbled.
// This is called a Data Race.
//
// Testing: Run `go test -race ./...` to detect this bug!
//
// Requirements:
// 1. Refactor `Aggregator` to use a `sync.Mutex`.
// 2. Ensure `ReceiveChunk` safely increments `TotalBytes` without causing a data race.
// 3. Ensure `GetTotal` safely returns the value without reading while another
//    goroutine might be writing.

type Aggregator struct {
	// TODO: Add a Mutex
	TotalBytes  int
	TotalChunks int
}

func (a *Aggregator) ReceiveChunk(bytes int) {
	// BUG: This causes a data race!
	// TODO: Protect this critical section.

	a.TotalBytes += bytes
	a.TotalChunks++
}

func (a *Aggregator) GetTotal() int {
	// BUG: Reading while another goroutine is writing is ALSO a data race!
	// TODO: Protect this read.

	return a.TotalBytes
}
