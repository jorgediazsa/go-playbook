package control

import (
	"sync"
)

// Context: Loop Variable Capture (and addressal)
// You are loading configuration items and creating worker configurations based on them.
// A common pattern is storing pointers to structs in a slice to avoid copying large objects.
//
// Why this matters: Historically, taking the address of a loop variable in Go
// (`&item`) took the address of the SINGLE instance of that variable used by the loop.
// By the end of the loop, all pointers in the slice pointed to the exact same struct (the last one).
// Even with Go 1.22 fixing `for _, v`, understanding variable shadowing and pointer safety is crucial.
//
// Requirements:
// 1. Process string tasks into pointers to `Task` structs.
// 2. Ensure that each pointer in the returned slice points to a distinct `Task`.
// 3. Do the same for processing via goroutines (ensure they don't all process the last task).

type Task struct {
	Name string
	Done bool
}

func CollectTaskPointers(taskNames []string) []*Task {
	var result []*Task
	for _, name := range taskNames {
		// BUG: We are capturing the loop variable by reference.
		// Pre-1.22, this causes all pointers in `result` to point to the memory
		// holding the final iterated string.
		// TODO: Fix the generation of pointers so each is distinct.
		t := Task{Name: name, Done: false}
		result = append(result, &t)
	}
	return result
}

func ProcessTasksConcurrently(taskNames []string) []string {
	var processed []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, name := range taskNames {
		wg.Add(1)
		// BUG: The goroutine closes over the `name` variable.
		// Pre-1.22, all goroutines might wake up after the loop finishes and only see the final task.
		// TODO: Fix the concurrency capture bug.
		go func() {
			defer wg.Done()
			mu.Lock()
			processed = append(processed, name+"_processed")
			mu.Unlock()
		}()
	}

	wg.Wait()
	return processed
}
