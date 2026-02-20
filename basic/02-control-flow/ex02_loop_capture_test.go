package control

import (
	"sort"
	"testing"
)

func TestCollectTaskPointers(t *testing.T) {
	names := []string{"DataSync", "LogRotate", "CacheClear"}
	tasks := CollectTaskPointers(names)

	if len(tasks) != 3 {
		t.Fatalf("Expected 3 tasks, got %d", len(tasks))
	}

	for i, task := range tasks {
		if task.Name != names[i] {
			t.Errorf("Expected task %d to be %s, got %s. You likely hit the loop capture bug.", i, names[i], task.Name)
		}
	}

	// Double check memory addresses to ensure they are distinct
	if &tasks[0] == &tasks[1] || &tasks[1] == &tasks[2] {
		t.Errorf("Multiple pointers point to the exact same memory address!")
	}
}

func TestProcessTasksConcurrently(t *testing.T) {
	names := []string{"JobA", "JobB", "JobC", "JobD", "JobE", "JobF", "JobG"}

	// We run it multiple times to aggravate the race condition if the bug exists.
	for i := 0; i < 50; i++ {
		processed := ProcessTasksConcurrently(names)
		if len(processed) != len(names) {
			t.Fatalf("Expected %d processed names, got %d", len(names), len(processed))
		}

		sort.Strings(processed)
		expected := []string{}
		for _, n := range names {
			expected = append(expected, n+"_processed")
		}
		sort.Strings(expected)

		for j := range processed {
			if processed[j] != expected[j] {
				t.Fatalf("Iteration %d: Context capture bug. Expected %v, got %v", i, expected, processed)
			}
		}
	}
}
