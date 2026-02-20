package collections

// Context: Map Concurrency Restrictions
// You're building a simple in-memory cache for user sessions.
// Thousands of goroutines will read from and write to this cache simultaneously.
//
// Why this matters: Maps in Go are completely unsafe for concurrent use.
// A concurrent read and write, or concurrent writes, will trigger a fatal panic
// that crashes the entire Go process (it cannot be recovered).
//
// Requirements:
// 1. Build a `SessionCache` struct that safely wraps a `map[string]string`.
// 2. Implement `Set` and `Get` methods that are safe for concurrent use.
// 3. Use `sync.RWMutex` to allow concurrent reads but exclusive writes.

// TODO: Define the struct properly.
type SessionCache struct {
	data map[string]string
	// Add a mutex here
}

func NewSessionCache() *SessionCache {
	return &SessionCache{
		data: make(map[string]string),
	}
}

func (c *SessionCache) Set(key, value string) {
	// BUG: Vulnerable to concurrent map writes
	// TODO: Make this safely lock for writing
	c.data[key] = value
}

func (c *SessionCache) Get(key string) (string, bool) {
	// BUG: Vulnerable to concurrent map read/writes
	// TODO: Make this safely lock for reading
	val, ok := c.data[key]
	return val, ok
}
