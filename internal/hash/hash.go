package hash

import (
	"hash/fnv"
	"math"
)

// Types of Hashable inputs

func calculateHash(hashA, hashB, n, size int) int {
	return int(math.Abs(float64((hashA + n*hashB + int(math.Floor((math.Pow(float64(n), 3) - float64(n)) / 6))) % size)))
}

func Hash(data []byte, seed uint32) uint32 {
	h := fnv.New32a()
	h.Write(data)
	return h.Sum32() + seed
}
