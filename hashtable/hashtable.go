package hashtable

import (
	"fmt"
	"sync"
)

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
		buckets:  make([]*Node[K, V], capacity),
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
		Key:   key,
		Value: value,
		Next:  head,
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

func (ht *HashTable[K, V]) Delete(key K) bool {
	ht.mu.Lock()
	defer ht.mu.Unlock()

	idx := ht.hashFunc(key) % uint64(ht.capacity)
	head := ht.buckets[idx]

	if head == nil {
		return false
	}

	if head.Key == key {
		ht.buckets[idx] = head.Next
		ht.size--
		return true
	}

	for curr := head; curr.Next != nil; curr = curr.Next {
		if curr.Next.Key == key {
			curr.Next = curr.Next.Next
			ht.size--
			return true
		}
	}
	return false
}

func (ht *HashTable[K, V]) Size() int {
	ht.mu.RLock()
	defer ht.mu.RUnlock()
	return ht.size
}

func hashString(key string) uint64 {
	const (
		offset = 14695981039346656037
		prime  = 1099511628211
	)
	var hash uint64 = offset
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= prime
	}
	return hash
}

func StartHashTable() {
	ht := NewHashTable[string, int](4, hashString)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			ht.Insert(key, id*10)
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			if val, ok := ht.Get(key); ok {
				fmt.Printf("[Чтение] %s -> %d\n", key, val)
			}
		}(i)
	}
	wg.Wait()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", id)
			if ht.Delete(key) {
				fmt.Printf("[Удаление] %s удалён успешно\n", key)
			}
		}(i)
	}
	wg.Wait()

	fmt.Printf("\nИтоговый размер таблицы: %d\n", ht.Size())
}
