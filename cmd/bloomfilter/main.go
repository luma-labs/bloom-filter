package main

import (
	"fmt"
	"github.com/luma-labs/bloom-filter/pkg/bloomfilter"
)

func main() {
	// Create a Bloom filter for 100 items with a false positive rate of 0.01
	filter := bloomfilter.Create(100, 0.01)
	// Add some elements
	filter.Add([]byte("alice"))
	filter.Add([]byte("bob"))
	filter.Add([]byte("alice"))
	filter.Add([]byte("alisdf"))
	filter.Add([]byte("alsdfsdfdices"))
	filter.Add([]byte("alsdfsfices"))
	filter.Add([]byte("aldfsfdices"))
	filter.Add([]byte("aldffdices"))
	filter.Add([]byte("alicdses"))
	filter.Add([]byte("alicdeass"))
	filter.Add([]byte("aliceass"))
	filter.Add([]byte("alicasadsdsaeass"))
	filter.Add([]byte("aliceasdass"))
	filter.Add([]byte("alicedasdass"))



	fmt.Println("filter", filter.PrintEveryByte())
	fmt.Println("Has:",filter.Has([]byte("alicee")))
}