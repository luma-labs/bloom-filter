package bloomfilter

import (
	"github.com/luma-labs/bloom-filter/internal/bits"
	"github.com/luma-labs/bloom-filter/internal/hash"
	"github.com/luma-labs/bloom-filter/internal/utils"
	"math"
	"fmt"
)

// BloomFilter represents the Bloom filter structure
type BloomFilter struct {
	size     uint
	nbHashes uint
	filter   *bits.Bitset
	seed     uint64
}

// NewBloomFilter creates a new Bloom filter with the given size and number of hash functions
func NewBloomFilter(size uint, nbHashes uint) *BloomFilter {
	if nbHashes < 1 {
		panic(fmt.Sprintf("A BloomFilter cannot use less than one hash function, but got %d", nbHashes))
	}
	return &BloomFilter{
		size:     size,
		nbHashes: nbHashes,
		filter:   bits.NewBitset(int(size)),
		seed:     utils.GetDefaultSeed(),
	}
}

// Create generates an optimal Bloom filter based on the number of items and desired error rate
func Create(nbItems int, errorRate float64) *BloomFilter {
	size := OptimalFilterSize(nbItems, errorRate)
	hashes := OptimalHashes(size, nbItems)
	return NewBloomFilter(size, hashes)
}

// From creates a new Bloom filter from an iterable of items with a fixed error rate
func From(items []string, errorRate float64, seed *uint64) *BloomFilter {
	filter := Create(len(items), errorRate)
	if seed != nil {
		filter.seed = *seed
	} else {
		filter.seed = utils.GetDefaultSeed()
	}
	for _, item := range items {
		filter.Add([]byte(item))
	}
	return filter
}

// OptimalFilterSize computes the optimal size for the Bloom filter
func OptimalFilterSize(nbItems int, errorRate float64) uint {
	return uint(math.Ceil(float64(nbItems) * math.Abs(math.Log(errorRate)) / (math.Ln2 * math.Ln2)))
}

// OptimalHashes computes the optimal number of hash functions
func OptimalHashes(size uint, nbItems int) uint {
	return uint(math.Ceil((float64(size) / float64(nbItems)) * math.Ln2))
}

// Add adds an element to the Bloom filter
func (bf *BloomFilter) Add(element []byte) {
	hashing := hash.Hashing{}
	indexes := hashing.GetIndexes(element, int(bf.size), int(bf.nbHashes), &bf.seed)
	for _, index := range indexes {
		bf.filter.Add(index)
	}
}

// Has checks if an element might be in the filter
func (bf *BloomFilter) Has(element []byte) bool {
	hashing := hash.Hashing{}
	indexes := hashing.GetIndexes(element, int(bf.size), int(bf.nbHashes), &bf.seed)
	for _, index := range indexes {
		if !bf.filter.Contains(index) {
			return false
		}
	}
	return true
}

// Rate calculates the current false positive rate of the Bloom filter
func (bf *BloomFilter) Rate() float64 {
	return math.Pow(1-math.Exp(-float64(bf.filter.NumOfTrueBits())/float64(bf.size)), float64(bf.nbHashes))
}

// Equals checks if another Bloom filter is equal to this one
func (bf *BloomFilter) Equals(other *BloomFilter) bool {
	if bf.size != other.size || bf.nbHashes != other.nbHashes {
		return false
	}
	return bf.filter.Equals(other.filter)
}
