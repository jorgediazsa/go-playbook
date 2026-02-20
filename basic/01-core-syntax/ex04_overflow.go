package core

import "errors"

// Context: Subtle integer overflows.
// We are parsing 64-bit integer timestamps or counters from a modern gRPC stream,
// but we need to push them into legacy systems that only support 32-bit unsigned integers (uint32).
// A direct cast `uint32(val)` truncates the bits, silently corrupting data in production.
//
// Why this matters: Go's type conversions (e.g., uint32(myInt64)) NEVER panic on overflow.
// They silently truncate the high bits. This is a severe silent data corruption vector.
//
// Requirements:
// 1. Implement `SafeConvertInt64ToUint32(val int64) (uint32, error)`.
// 2. If the value is strictly less than 0 or exceeds the max value a uint32 can hold,
//    return a sentinel error `ErrOverflow`.
// 3. Otherwise return the converted value.

var ErrOverflow = errors.New("integer overflow")

func SafeConvertInt64ToUint32(val int64) (uint32, error) {
	// TODO: Catch overflow/underflow, returning ErrOverflow when appropriate.
	// Otherwise, return the safely casted uint32.

	return uint32(val), nil // BUG: Silently truncates today.
}
