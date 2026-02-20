package garbagecollector

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestExtractMetadataRetention(t *testing.T) {
	MetadataCache = nil // reset

	// 1. Create a massive fake image (1MB)
	massiveImage := make([]byte, 1024*1024)

	// 2. Run the extraction
	ExtractMetadata(massiveImage)

	if len(MetadataCache) != 1 {
		t.Fatalf("Cache should have 1 item")
	}

	cachedHeader := MetadataCache[0]

	if len(cachedHeader) != 64 {
		t.Fatalf("Expected 64 bytes, got %d", len(cachedHeader))
	}

	// 3. The true test of retention: Are they point to the same backing array?
	// We use `unsafe` and reflection to grab the raw memory pointers of the slices.

	ptrMassive := (*reflect.SliceHeader)(unsafe.Pointer(&massiveImage)).Data
	ptrCached := (*reflect.SliceHeader)(unsafe.Pointer(&cachedHeader)).Data

	// If the user used `massiveImage[:64]`, these pointers are mathematically identical!
	// If they used `copy`, `make([]byte, 64)` allocated memory cleanly elsewhere on the heap.

	if ptrMassive == ptrCached {
		t.Fatalf("FAILED: Memory Retention Detected! The cached subslice points directly to the 1MB underlying array! The GC can never free the image. You must explicitly allocate and copy().")
	}
}
