package bloomfilter

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := Create(100, 0.01)

	bf.Add([]byte("example"))

	if !bf.Has([]byte("example")) {
		t.Errorf("Expected item to be present in the Bloom filter")
	}

	if bf.Has([]byte("missing")) {
		t.Errorf("Expected item to be absent from the Bloom filter")
	}
}
