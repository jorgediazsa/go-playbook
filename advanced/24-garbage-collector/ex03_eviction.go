package garbagecollector

import (
	"sync"
)

// Context: Unbounded Caches and OOM
// You are building a session store. Whenever a user logs in, you save their
// profile in a global `map[string]User{}` so you don't have to hit the DB again.
//
// Why this matters: A map in Go never shrinks. If you add 10,000,000 users
// over a month, the map will consume gigabytes of memory. Because the map
// holds active pointers to the `User` structs, the Garbage Collector can NEVER
// free them. Your server will inevitably OOM (Out Of Memory) crash.
//
// Requirements:
// 1. Refactor `SessionCache` into a bounded cache.
// 2. Simplest approach: Add a `MaxKeys int` field. If `Set()` is called and
//    `len(c.store) >= c.MaxKeys`, you MUST evict an item.
// 3. For this exercise, simple random eviction is acceptable (just use a `for k := range c.store { delete(c.store, k); break }`).
//    In the real world, you would use an LRU or a TTL.
// 4. Ensure it remains safe for concurrent access (Mutex).

type User struct {
	ID   string
	Data [1024]byte // 1KB of data per user
}

type SessionCache struct {
	mu    sync.Mutex
	store map[string]User

	// TODO: Add MaxKeys to enforce a cap
}

func NewSessionCache() *SessionCache {
	return &SessionCache{
		store: make(map[string]User),
	}
}

func (c *SessionCache) Set(id string, u User) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// BUG: Unbounded growth!
	// TODO: If len >= MaxKeys, evict something before inserting!

	c.store[id] = u
}

func (c *SessionCache) Get(id string) (User, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	u, ok := c.store[id]
	return u, ok
}
