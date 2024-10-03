package main

import (
    "flag"
    "fmt"
    "github.com/luma-labs/bloom-filter/pkg/bloomfilter"
)

func main() {
    element := flag.String("add", "", "Element to add to the Bloom filter")
    check := flag.String("check", "", "Element to check in the Bloom filter")
    flag.Parse()

    bf := bloomfilter.NewBloomFilter(1000, 3)

    if *element != "" {
        bf.Add([]byte(*element))
        fmt.Println("Element added.")
    }

    if *check != "" {
        if bf.Contains([]byte(*check)) {
            fmt.Println("Item is probably in the set")
        } else {
            fmt.Println("Item is definitely not in the set")
        }
    }
}
