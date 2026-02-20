package stringsbytes

// Context: UTF-8 Correctness
// You are writing a truncation middleware for a logging pipeline.
// If a log message exceeds a certain character limit, you must truncate it
// and append "..." so it limits storage costs.
//
// Why this matters: You must truncate based on actual CHARACTERS (runes),
// not bytes. If you truncate a Japanese log message blindly via bytes,
// you will slice a multi-byte character in half, causing invalid UTF-8
// that crashes down-stream Elasticsearch parsers.
//
// Requirements:
// 1. `Truncate` should safely return the first `maxChars` of the string.
// 2. If the string has fewer characters than `maxChars`, return it unchanged.
// 3. Otherwise, return the truncated string with "..." appended.
// 4. Do not corrupt multi-byte characters!

func Truncate(logMsg string, maxChars int) string {
	// BUG: len(logMsg) measures BYTES.
	// `logMsg[:maxChars]` slices BYTES.
	// TODO: Fix this to safely truncate based on RUNES (characters).

	if len(logMsg) <= maxChars {
		return logMsg
	}

	return logMsg[:maxChars] + "..."
}
