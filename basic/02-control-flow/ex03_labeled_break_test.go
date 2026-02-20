package control

import (
	"reflect"
	"testing"
)

func TestEventLoopWorker(t *testing.T) {
	commands := []string{
		"START",
		"IGNORE_THIS",
		"PROCESS_A",
		"SHUTDOWN",
		"SHOULD_NOT_PROCESS",
	}

	processed := EventLoopWorker(commands)

	expected := []string{
		"START",
		"PROCESS_A",
	}

	if !reflect.DeepEqual(processed, expected) {
		t.Fatalf("\nExpected: %v\nGot:      %v\nCheck labeled breaks and continues.", expected, processed)
	}
}
