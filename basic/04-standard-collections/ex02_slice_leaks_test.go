package collections

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestExtractTxID(t *testing.T) {
	// Simulate a 100MB payload (we'll just use a small one for the test, but the pointer logic holds)
	payload := make([]byte, 1024)
	copy(payload, []byte("TX-99887766-DATA-DATA-DATA"))

	txID := ExtractTxID(payload)

	expected := []byte("TX-9988776")
	if !reflect.DeepEqual(txID, expected) {
		t.Fatalf("Expected %q, got %q", expected, txID)
	}

	// Verify they do not share the exact same backing array.
	// We can do this by taking the address of the first element.
	ptrPayload := &payload[0]
	ptrTxID := &txID[0]

	if uintptr(unsafe.Pointer(ptrPayload)) == uintptr(unsafe.Pointer(ptrTxID)) {
		t.Fatalf("MEMORY LEAK: The returned txID slice points to the same backing array as the large payload!")
	}
}
