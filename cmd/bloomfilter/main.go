package main

import (
	"fmt"

	"github.com/luma-labs/bloom-filter/pkg/bloomfilter"
)

func main() {
	// Create a Bloom filter for 100 items with a false positive rate of 0.01
	nbItems := 1000
	errorRate := 0.01
	filter := bloomfilter.Create(nbItems, errorRate)
	bloomfilter.DisplaySize(nbItems, errorRate)

	filter.Add([]byte("alice"))
	filter.Add([]byte("alice1"))
	filter.Add([]byte("alice2"))
	filter.Add([]byte("bob"))
	filter.Add([]byte("bob1"))
	filter.Add([]byte("bob2"))
	filter.Add([]byte("bob3"))

	fmt.Println("Has 'alice':", filter.Has([]byte("alice")))
	fmt.Println("Has 'alice':", filter.Has([]byte("alice1")))
	fmt.Println("Has 'alice':", filter.Has([]byte("alice2")))
	fmt.Println("Has 'bob':", filter.Has([]byte("bob")))
	fmt.Println("Has 'charlie':", filter.Has([]byte("bob1")))
	fmt.Println("Has 'charlie':", filter.Has([]byte("bob2")))
	fmt.Println("Has 'charlie':", filter.Has([]byte("bob3")))
	fmt.Println("Has 'charlie':", filter.Has([]byte("bob4")))

	filter.PrintEveryByte()
}
