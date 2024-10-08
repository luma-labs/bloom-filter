package main

import (
	"fmt"
	"github.com/luma-labs/bloom-filter/pkg/bloomfilter"
)

func main() {
	// Create a Bloom filter for 100 items with a false positive rate of 0.01
	filter := bloomfilter.NewBloomFilter(10, 3)

	// Add some elements
	filter.Add([]byte("alice"))
	filter.Add([]byte("bob"))

	fmt.Println("False positive rate:", filter.Rate())

	anotherFilter := bloomfilter.NewBloomFilter(10, 3)
	anotherFilter.Add([]byte("alice"))
	anotherFilter.Add([]byte("bob"))
	fmt.Println("Are the filters equal?", filter.Equals(anotherFilter)) // false
}