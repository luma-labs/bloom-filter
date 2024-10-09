package main

import (
	"fmt"

	"github.com/luma-labs/bloom-filter/pkg/bloomfilter"
)

func main() {
	// Create a Bloom filter for 100 items with a false positive rate of 0.01
	filter := bloomfilter.Create(100, 0.01)

	filter.Add([]byte("alice"))
	filter.Add([]byte("bob"))

	fmt.Println("Has 'alice':", filter.Has([]byte("alice")))
	fmt.Println("Has 'bob':", filter.Has([]byte("bob")))
	fmt.Println("Has 'charlie':", filter.Has([]byte("charlie")))

	filter.PrintEveryByte()
}
