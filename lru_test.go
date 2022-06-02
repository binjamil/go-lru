package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	l, err := New[string, int](2)
	assert.Nil(t, err)

	assert.False(t, l.Add("one", 1))
	assert.False(t, l.Add("two", 2))
	assert.False(t, l.Add("one", 1))
	assert.True(t, l.Add("three", 3))

	val, ok := l.Get("one")
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	_, ok = l.Get("two")
	assert.False(t, ok)
	val, ok = l.Get("three")
	assert.True(t, ok)
	assert.Equal(t, 3, val)

	assert.Equal(t, uint(2), l.Len())
	l.Remove("one")
	assert.Equal(t, uint(1), l.Len())
	l.Add("thirty", 30)
	assert.Equal(t, uint(2), l.Len())

	ok = l.Contains("three")
	assert.True(t, ok)
	l.Add("one", 1)
	ok = l.Contains("three")
	assert.False(t, ok)
}

func TestLRU_Add(t *testing.T) {
	l, err := New[string, int](1)
	assert.Nil(t, err)

	evicted := l.Add("one", 1)
	assert.False(t, evicted)
	el, ok := l.items["one"]
	got := el.Value
	assert.True(t, ok)
	assert.Equal(t, got.key, "one")
	assert.Equal(t, got.val, 1)
	assert.Equal(t, 1, len(l.items))
	assert.Equal(t, uint(1), l.evictList.Len())

	evicted = l.Add("two", 2)
	assert.True(t, evicted)
	_, ok = l.items["one"]
	assert.False(t, ok)
	el, ok = l.items["two"]
	got = el.Value
	assert.True(t, ok)
	assert.Equal(t, got.key, "two")
	assert.Equal(t, got.val, 2)
	assert.Equal(t, 1, len(l.items))
	assert.Equal(t, uint(1), l.evictList.Len())
}

func TestLRU_Get(t *testing.T) {
	l, err := New[string, int](1)
	assert.Nil(t, err)

	_, ok := l.Get("one")
	assert.False(t, ok)
	l.items["one"] = l.evictList.PushFront(entry[string, int]{"one", 1})
	val, ok := l.Get("one")
	assert.True(t, ok)
	assert.Equal(t, 1, val)
}

func TestLRU_Contains(t *testing.T) {
	l, err := New[string, int](1)
	assert.Nil(t, err)

	assert.False(t, l.Contains("one"))
	l.items["one"] = l.evictList.PushFront(entry[string, int]{"one", 1})
	assert.True(t, l.Contains("one"))
}

func TestLRU_Len(t *testing.T) {
	l, err := New[string, int](1)
	assert.Nil(t, err)

	assert.Equal(t, uint(0), l.Len())
	l.items["one"] = l.evictList.PushFront(entry[string, int]{"one", 1})
	assert.Equal(t, uint(1), l.Len())
}

func TestLRU_Remove(t *testing.T) {
	l, err := New[string, int](1)
	assert.Nil(t, err)

	l.items["one"] = l.evictList.PushFront(entry[string, int]{"one", 1})
	assert.True(t, l.Remove("one"))
	assert.Equal(t, 0, len(l.items))
	assert.Equal(t, uint(0), l.evictList.Len())
	assert.False(t, l.Remove("one"))
}
