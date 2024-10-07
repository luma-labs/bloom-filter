package bloomfilter

import (
	"github.com/luma-labs/bloom-filter/internal/bits"
)

type Hashable interface {
	isHashable()
}

type StringInput string
type ByteInput []byte
type IntInput int32

func (StringInput) isHashable() {}
func (ByteInput) isHashable()   {}
func (IntInput) isHashable()    {}

type BloomFilter struct {
	size     int
	hashFunc int
	bitset   *bits.Bitset
}

func NewBloomFilter(size, hashFunc int) *BloomFilter {
	return &BloomFilter{
		size:     size,
		hashFunc: hashFunc,
		bitset:   bits.NewBitset(size),
	}
}
