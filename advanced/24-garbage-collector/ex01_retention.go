package garbagecollector

// Context: Object Retention via Subslicing
// You are building an image processing pipeline. A worker reads a 5MB image,
// extracts only the 64-byte EXIF metadata header, and stores the metadata
// in a long-lived cache.
//
// Why this matters: In Go, a subslice `largeSlice[:64]` points to the exact
// same underlying memory array as the large slice. If you store the 64-byte
// subslice in a cache that lasts for days, the GC is completely prohibited from
// freeing the 5MB image array, because there is still an active pointer to it!
// At thousands of images per second, memory rockets to 100% instantly.
//
// Requirements:
// 1. Refactor `ExtractMetadata` to return a completely independent byte slice.
// 2. You must allocate a new slice of the exact required length (64).
// 3. You must completely `copy()` the data into the new slice.

var MetadataCache [][]byte

func ExtractMetadata(massiveImage []byte) {
	if len(massiveImage) < 64 {
		return
	}

	// BUG: This subslice retains the ENTIRE underlying array of `massiveImage`.
	// TODO: Create a new slice `metadata` of length 64, and `copy` the first
	// 64 bytes of `massiveImage` into it. Then store `metadata`.

	metadata := massiveImage[:64]

	// Store it in the cache for the lifetime of the application
	MetadataCache = append(MetadataCache, metadata)
}
