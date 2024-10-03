package benchmarks
// package bloomfilter

import (
	"testing"
	"github.com/luma-labs/bloom-filter/pkg/bloomfilter"
)

func BenchmarkBloomFilterAdd(b *testing.B) {
	bf := bloomfilter.NewBloomFilter(1000, 3)
	for i := 0; i < b.N; i++ {
		bf.Add([]byte("example"))
	}
}

func BenchmarkBloomFilterContains(b *testing.B) {
	bf := bloomfilter.NewBloomFilter(1000, 3)
	bf.Add([]byte("example"))
	for i := 0; i < b.N; i++ {
		bf.Contains([]byte("example"))
	}
}
