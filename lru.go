package lru

import (
	"errors"
	"sync"

	"github.com/binjamil/go-lru/list"
)

// LRU is a thread-safe fixed size LRU cache.
type LRU[K comparable, V any] struct {
	capacity  uint
	evictList *list.List[V]
	items     map[K]*list.Element[V]
	lock      sync.RWMutex
}

// New creates an LRU cache with the given capacity.
func New[K comparable, V any](capacity uint) (*LRU[K, V], error) {
	if capacity == 0 {
		return nil, errors.New("capacity must be a positive integer")
	}
	lru := &LRU[K, V]{
		capacity:  capacity,
		evictList: list.New[V](),
		items:     make(map[K]*list.Element[V]),
	}
	return lru, nil
}

// Add adds a value to the cache and returns true if an eviction occurred.
func (lru *LRU[K, V]) Add(key K, val V) (evicted bool) {
	lru.lock.Lock()
	defer lru.lock.Unlock()
	// Check if item already exists
	if el, ok := lru.items[key]; ok {
		lru.evictList.MoveToFront(el)
		el.Value = val
		return false
	}

	// Add the new item
	el := lru.evictList.PushFront(val)
	lru.items[key] = el

	// Evict oldest item if capacity overflows
	evicted = lru.evictList.Len() > lru.capacity
	if evicted {
		oldest := lru.evictList.Back()
		lru.evictList.Remove(oldest)
	}
	return
}

// Get looks up a key's value and returns (value, true) if it exists.
// If the value doesn't exist, it returns (nil, false).
func (lru *LRU[K, V]) Get(key K) (val V, ok bool) {
	if el, ok := lru.items[key]; ok {
		lru.evictList.MoveToFront(el)
		val = el.Value
	}
	return
}
