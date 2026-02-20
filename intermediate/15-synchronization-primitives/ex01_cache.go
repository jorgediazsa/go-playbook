package syncprims

import "sync"

// Context: Mutex vs RWMutex
// You are building an in-memory DNS cache. Millions of requests constantly read
// `Resolve(domain)`, but writes `Set(domain, ip)` happen only a few times an hour.
//
// Why this matters: Using a standard `sync.Mutex` forces millions of readers into
// a single-file line (contention bottleneck). Since reads don't mutate state, they
// are safe to occur entirely in parallel.
//
// Requirements:
// 1. Refactor `DNSCache` to use `sync.RWMutex`.
// 2. Update `Resolve` to use the Reader lock (allows parallel reads).
// 3. Update `Set` to use the Writer lock (exclusive).

type DNSCache struct {
	// BUG: Using standard Mutex restricts concurrency.
	// TODO: Use RWMutex instead.
	mu    sync.Mutex
	store map[string]string
}

func NewDNSCache() *DNSCache {
	return &DNSCache{
		store: make(map[string]string),
	}
}

func (c *DNSCache) Resolve(domain string) (string, bool) {
	// BUG: This locks out other readers!
	// TODO: Use RLock() instead.
	c.mu.Lock()
	defer c.mu.Unlock()

	ip, exists := c.store[domain]
	return ip, exists
}

func (c *DNSCache) Set(domain, ip string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[domain] = ip
}
