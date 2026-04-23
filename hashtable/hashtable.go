package hashtable

import "sync"

type Node[K comparable, V any] struct {
	Key   K
	Value V
	Next  *Node[K, V]
}

type HashTable[K comparable, V any] struct {
	mu       sync.RWMutex
	buckets  []*Node[K, V]
	size     int
	capacity int
	hashFunc func(K) uint64
}

func NewHashTable[K comparable, V any](capacity int, hashFunc func(K) uint64) *HashTable[K, V] {
	if capacity <= 0 {
		capacity = 16
	}
	if hashFunc == nil {
		panic("нет hash func")
	}
	return &HashTable[K, V]{
		buckets: make([]*Node[K, V], capacity),
		capacity: capacity,
		hashFunc: hashFunc,
	}
}

func (ht *HashTable[K, V]) Insert(key K, value V) {
	ht.mu.Lock()
	defer ht.mu.Unlock()

	idx := ht.hashFunc(key) % uint64(ht.capacity)
	head := ht.buckets[idx]

	for curr := head; curr != nil; curr = curr.Next {
		if curr.Key == key {
			curr.Value = value
			return
		}
	}

	ht.buckets[idx] = &Node[K, V]{
		Key: key,
		Value: value,
		Next: head,
	}
	ht.size++
}

func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	ht.mu.RLock()
	defer ht.mu.RUnlock()

	idx := ht.hashFunc(key) % uint64(ht.capacity)
	var zero V

	for curr := ht.buckets[idx]; curr != nil; curr = curr.Next {
		if curr.Key == key {
			return curr.Value, true
		}
	}
	return zero, false
}

