package syncprims

import (
	"errors"
	"sync"
)

// Context: sync.Once vs sync.OnceValue (Fallible Initialization)
// You are building a Database connection pool singleton. Hundreds of goroutines
// will ask for the DB connection at startup, but it should only be initialized once.
//
// Why this matters: Historically, `sync.Once.Do(initFunc)` was dangerous.
// If `initFunc` returned an error (e.g., dial timeout), `Once` still marked it
// as "successfully run", and all future callers got nil.
// Go 1.21 introduced `sync.OnceValues` precisely for this.
//
// Requirements:
// 1. Refactor `DBManager` to efficiently initialize the connection ONCE.
// 2. You MUST return the error if the initialization fails.
// 3. If initialization fails, SUBSEQUENT callers should ALSO receive the
//    initialization error, rather than attempting to reconnect or panicking.
//    (Hint: Use `sync.OnceValues` from Go 1.21).

type DBConnection struct{ Status string }

var ErrTimeout = errors.New("connection timeout")

type DBManager struct {
	// TODO: Replace Mutex with sync.OnceValues returning (*DBConnection, error)
	mu   sync.Mutex
	conn *DBConnection
	err  error
	done bool
}

// simulateConnect is a slow, fallible operation.
var mockFail = true

func simulateConnect() (*DBConnection, error) {
	if mockFail {
		return nil, ErrTimeout
	}
	return &DBConnection{Status: "OK"}, nil
}

func (m *DBManager) GetConnection() (*DBConnection, error) {
	// BUG: The Mutex blocks all readers, turning this into a massive bottleneck
	// for the entire lifecycle of the application!
	// TODO: Refactor using sync.OnceValues for thread-safe, fast, fallible caching.

	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.done {
		m.conn, m.err = simulateConnect()
		m.done = true
	}
	return m.conn, m.err
}

// NewDBManager constructor
func NewDBManager() *DBManager {
	// TODO: Initialize your sync.OnceValues closure here.
	return &DBManager{}
}
