// Package lru provides a fixed size LRU cache
//
// - O(1) Add, Get, Contains and Remove
// - Generic implementation for better type-safety
// - Thread-safe operations
//
// LRU is implemented via a generic doubly linked list, available in package
// "github.com/binjamil/go-lru/list"
package lru
