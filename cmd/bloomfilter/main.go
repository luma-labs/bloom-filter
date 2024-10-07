package main

import (
    "fmt"
    "github.com/luma-labs/bloom-filter/pkg/bloomfilter"
)



func main() {
    bf := bloomfilter.NewBloomFilter(1000, 3)
    bf.Add([]byte("example123"))
    bf.Add([]byte("example"))
	fmt.Println(bf)

    if bf.Contains([]byte("example123")) {
        fmt.Println("Item is probably in the set")
    } else {
        fmt.Println("Item is definitely not in the set")
    }
}
