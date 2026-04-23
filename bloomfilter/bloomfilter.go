package bloomfilter

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	size uint64
	k    uint64
	bits []byte
}

func NewBloomFilter(sizeBits, k uint64) *BloomFilter {
	if sizeBits == 0 || k == 0 {
		panic("sizeBits и k должны быть больше 0")
	}
	return &BloomFilter{
		size: sizeBits,
		k:    k,
		bits: make([]byte, (sizeBits+7)/8),
	}
}

func (bf *BloomFilter) Add(data []byte) {
	h1, h2 := bf.gatBaseHashes(data)
	for i := uint64(0); i < bf.k; i++ {
		pos := (h1 + i*h2) % bf.size
		byteIdx := pos / 8
		bitIdx := pos % 8
		bf.bits[byteIdx] |= (1 << bitIdx)
	}
}

func (bf *BloomFilter) Test(data []byte) bool {
	h1, h2 := bf.gatBaseHashes(data)
	for i := uint64(0); i < bf.k; i++ {
		pos := (h1 + i*h2) % bf.size
		byteIdx := pos / 8
		bitIdx := pos % 8
		if bf.bits[byteIdx]&(1<<bitIdx) == 0 {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) gatBaseHashes(data []byte) (uint64, uint64) {
	h := fnv.New64a()
	_, _ = h.Write(data)
	hash := h.Sum64()
	return hash, hash >> 32
}

func StartBloomFilter() {
	bf := NewBloomFilter(1024, 3)

	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))
	bf.Add([]byte("golang"))
	bf.Add([]byte("filter"))

	fmt.Println("'hello' присутствует:", bf.Test([]byte("hello")))
	fmt.Println("'world' присутствует:", bf.Test([]byte("world")))
	fmt.Println("'golang' присутствует:", bf.Test([]byte("golang")))
	fmt.Println("'python' присутствует:", bf.Test([]byte("python")))
}
