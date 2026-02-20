package defers

import (
	"reflect"
	"testing"
)

func TestReleaseResources(t *testing.T) {
	log := ReleaseResources()

	// The defers execute when ReleaseResources returns.
	// Since the acquisitions were: Lock -> DB -> File
	// The LIFO cleanup MUST be: File -> DB -> Lock

	expected := []string{
		"temp_file.txt_Closed", // Was it "temp_file" or "wrong_file" ?
		"DB_Closed",
		"Lock_Released",
	}

	if !reflect.DeepEqual(log, expected) {
		t.Fatalf("\nExpected: %v\nGot:      %v", expected, log)
	}
}
