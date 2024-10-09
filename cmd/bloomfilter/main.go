package main

import (
	"fmt"
	"github.com/luma-labs/bloom-filter/pkg/bloomfilter"
)

func main() {
	// Create a Bloom filter for 100 items with a false positive rate of 0.01
	filter := bloomfilter.Create(15, 0.9)
	// Add some elements
	filter.Add([]byte("alice"))
	filter.Add([]byte("bob"))
	filter.Add([]byte("alice"))
	filter.Add([]byte("alisdfce"))
	filter.Add([]byte("alices"))

	fmt.Println("size:",filter.Size())
	// fmt.Println("False positive rate:",filter.Size())

	// anotherFilter := bloomfilter.NewBloomFilter(10, 3)
	// anotherFilter.Add([]byte("alice"))
	// anotherFilter.Add([]byte("bob"))
	// fmt.Println("Are the filters equal?", filter.Equals(anotherFilter)) // false
}