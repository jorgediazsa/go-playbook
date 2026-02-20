package garbagecollector

import (
	"strings"
	"testing"
)

func TestRenderTemplatePooling(t *testing.T) {
	// 1. Simulate User A
	resA := RenderTemplate("USER A CONTENT")

	if resA != "<html><body>USER A CONTENT</body></html>" {
		t.Fatalf("User A rendering failed: %v", resA)
	}

	// 2. Simulate User B happening right after.
	// Since RenderTemplate runs synchronously here, the pool will immediately
	// give User B the exact same Buffer pointer that User A just Returned!

	resB := RenderTemplate("USER B CONTENT")

	if resB != "<html><body>USER B CONTENT</body></html>" {
		// If they forgot to call `buf.Reset()` when pulling from the pool,
		// the string will contain BOTH User A and User B's content appended together!
		if strings.Contains(resB, "USER A") {
			t.Fatalf("CRITICAL SECURITY LEAK: You returned a dirty buffer from the pool without calling `buf.Reset()`! User B just received User A's private data! Output: %v", resB)
		}

		t.Fatalf("User B rendering failed: %v", resB)
	}

	// How to verify they used the Pool and didn't just ignore it?
	// It's hard to strictly enforce Pool usage without benchmarks, but the structure
	// forces them to confront `Reset()` if they do use it.
	// Standard benchmarks (not included) would show the allocation drop to 0.
}
