package lru

import "github.com/binjamil/go-lru/list"

type LRU[T any] struct {
	ll *list.List[T]
}
