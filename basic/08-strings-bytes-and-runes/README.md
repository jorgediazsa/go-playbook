# 08 - Strings, Bytes, and Runes

Go's string handling is deeply tied to UTF-8. Misunderstanding the difference between a `string`, a `[]byte`, and a `rune` is a guaranteed way to write code that perfectly corrupts multi-byte languages (like Japanese, Arabic, or Emojis) in production.

## 1. String vs Byte vs Rune

- A `string` in Go is a read-only slice of bytes. It is just arbitrary data. It is *assumed* to be UTF-8 encoded text, but it doesn't have to be.
- A `byte` is an alias for `uint8` (1 byte).
- A `rune` is an alias for `int32` (4 bytes). It represents a single Unicode Code Point.

### The Production Danger

If you index a string `s[0]`, you get the **first byte**, not the first character. If the string is "世界" (World), `s[0]` returns `228`, a garbage byte fragment, not '世'.

If you slice a string `s[:2]`, it slices **bytes**. If the characters take 3 bytes each, slicing at byte 2 rips a character in half, creating an invalid UTF-8 sequence that renders as ``.

### Idiomatic Solution

To safely manipulate characters (count them, slice them, reverse them), you must either:
1. Iterate using `for i, runeVal := range s` (which safely decodes UTF-8 on the fly).
2. Cast the string to a slice of runes: `runes := []rune(s)`. Slicing the rune array `runes[:2]` safely grabs the first two characters. Note that this requires a heap allocation!
3. Use the `unicode/utf8` standard package for zero-allocation counting and decoding.

---

## 2. Allocation Costs

Strings are immutable. If you build a string in a loop using `str += "a"`, Go allocates a brand new backing array on the heap, copies the old string, and appends the new character. 

### The Production Danger

Building a 1MB string using `+=` inside a loop will allocate **gigabytes** of garbage and kill the CPU. Furthermore, converting `[]byte(str)` or `string(byteSlice)` ALWAYS forces a heap allocation to preserve the immutability guarantee of strings.

### Idiomatic Solution

1. To build strings in a loop, always use `strings.Builder`. It minimizes allocations and allows `Grow(n)` to preallocate capacity.
2. Avoid bouncing between `string` and `[]byte` unless strictly necessary.

---

## Exercises

- `ex01_utf8.go`
- `ex02_allocations.go`
