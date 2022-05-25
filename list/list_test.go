package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Len(t *testing.T) {
	ll := New[int]()
	assert.Zero(t, ll.Len())
	ll.PushFront(10)
	assert.Equal(t, uint(1), ll.Len())
	ll.PushFront(11)
	assert.Equal(t, uint(2), ll.Len())
}

func TestList_Front(t *testing.T) {
	ll := New[int]()
	ll.PushFront(10)
	assert.Equal(t, 10, ll.Front().Value)
	ll.PushFront(11)
	assert.Equal(t, 11, ll.Front().Value)
	ll.PushBack(9)
	assert.Equal(t, 11, ll.Front().Value)
}

func TestList_Back(t *testing.T) {
	ll := New[int]()
	ll.PushBack(9)
	assert.Equal(t, 9, ll.Back().Value)
	ll.PushBack(8)
	assert.Equal(t, 8, ll.Back().Value)
	ll.PushFront(10)
	assert.Equal(t, 8, ll.Back().Value)
}

func TestList_PushFront(t *testing.T) {
	ll := New[int]()
	ll.PushFront(10)
	assert.Equal(t, 10, ll.root.next.Value)
	ll.PushFront(11)
	assert.Equal(t, 11, ll.root.next.Value)
	assert.Equal(t, 10, ll.root.next.next.Value)
}

func TestList_PushBack(t *testing.T) {
	ll := New[int]()
	ll.PushBack(9)
	assert.Equal(t, 9, ll.root.prev.Value)
	ll.PushBack(8)
	assert.Equal(t, 8, ll.root.prev.Value)
	assert.Equal(t, 9, ll.root.prev.prev.Value)
}

func TestList_Remove(t *testing.T) {
	ll := New[int]()
	e1 := ll.PushFront(10)
	e2 := ll.PushFront(11)
	e3 := ll.PushFront(12)
	v := ll.Remove(e2)
	assert.Equal(t, 11, v)
	assert.Equal(t, e3, ll.root.next)
	assert.Equal(t, e1, ll.root.prev)
}

func TestList_MoveToFront(t *testing.T) {
	ll := New[int]()
	e1 := ll.PushFront(10)
	e2 := ll.PushFront(11)
	assert.Equal(t, e2, ll.root.next)
	ll.MoveToFront(e1)
	assert.Equal(t, e1, ll.root.next)
}

func TestList_MoveToBack(t *testing.T) {
	ll := New[int]()
	e1 := ll.PushBack(9)
	e2 := ll.PushBack(8)
	assert.Equal(t, e2, ll.root.prev)
	ll.MoveToBack(e1)
	assert.Equal(t, e1, ll.root.prev)
}

func TestElement_Next(t *testing.T) {
	ll := New[int]()
	for i := 0; i < 5; i++ {
		ll.PushBack(i)
	}
	count := 0
	for e := ll.Front(); e != nil; e = e.Next() {
		assert.Equal(t, count, e.Value)
		count++
	}
}

func TestElement_Prev(t *testing.T) {
	ll := New[int]()
	for i := 0; i < 5; i++ {
		ll.PushBack(i)
	}
	count := 4
	for e := ll.Back(); e != nil; e = e.Prev() {
		println(e.Value)
		assert.Equal(t, count, e.Value)
		count--
	}
}
