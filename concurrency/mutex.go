package concurrency

import "sync"

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{v: make(map[string]int)}
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}
