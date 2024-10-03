package bloomfilter

import (
	"github.com/luma-labs/bloom-filter/internal/hash"
)

type BloomFilter struct {
	bitset []bool
	size   int
	hashes int
}

func NewBloomFilter(size int, hashes int) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, size),
		size:   size,
		hashes: hashes,
	}
}

func (bf *BloomFilter) Add(item []byte) {
	for i := 0; i < bf.hashes; i++ {
		position := int(hash.Hash(item, uint32(i)) % uint32(bf.size))
		bf.bitset[position] = true
	}
}

func (bf *BloomFilter) Contains(item []byte) bool {
	for i := 0; i < bf.hashes; i++ {
		position := int(hash.Hash(item, uint32(i)) % uint32(bf.size))
		if !bf.bitset[position] {
			return false
		}
	}
	return true
}
