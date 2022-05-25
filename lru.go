package lru

import (
	"errors"
	"sync"

	"github.com/binjamil/go-lru/list"
)

// LRU is a thread-safe fixed size LRU cache.
type LRU[K comparable, V any] struct {
	capacity  uint
	evictList *list.List[entry[K, V]]
	items     map[K]*list.Element[entry[K, V]]
	lock      sync.RWMutex
}

type entry[K comparable, V any] struct {
	key K
	val V
}

// New creates an LRU cache with the given capacity.
func New[K comparable, V any](capacity uint) (*LRU[K, V], error) {
	if capacity == 0 {
		return nil, errors.New("capacity must be a positive integer")
	}
	lru := &LRU[K, V]{
		capacity:  capacity,
		evictList: list.New[entry[K, V]](),
		items:     make(map[K]*list.Element[entry[K, V]]),
	}
	return lru, nil
}

// Add adds a value to the cache and returns true if an eviction occurred.
func (c *LRU[K, V]) Add(key K, val V) (evicted bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	// Check if item already exists
	if el, ok := c.items[key]; ok {
		c.evictList.MoveToFront(el)
		el.Value.val = val
		return false
	}

	// Add the new item
	ent := entry[K, V]{key, val}
	el := c.evictList.PushFront(ent)
	c.items[key] = el

	// Evict oldest item if capacity overflows
	evicted = c.evictList.Len() > c.capacity
	if evicted {
		oldest := c.evictList.Back()
		c.evictList.Remove(oldest)
		delete(c.items, oldest.Value.key)
	}
	return
}

// Get looks up a key's value and returns (value, true) if it exists.
// If the value doesn't exist, it returns (nil, false).
func (c *LRU[K, V]) Get(key K) (val V, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	el, ok := c.items[key]
	if ok {
		c.evictList.MoveToFront(el)
		val = el.Value.val
	}
	return
}
