package hash

import (
	"strconv"
	"github.com/cespare/xxhash/v2"
	"github.com/luma-labs/bloom-filter/internal/utils"
)

// TwoHashes represents the result of two hash functions applied to the same value
type TwoHashes struct {
	First  uint64
	Second uint64
}

// TwoHashesTemplated allows flexible types for hash results
type TwoHashesTemplated[T any] struct {
	First  T
	Second T
}

// TwoHashesIntAndString holds hash results in both integer and string formats
type TwoHashesIntAndString struct {
	Int    TwoHashesTemplated[uint64]
	String TwoHashesTemplated[string]
}

// HashableInput represents a hashable value, like string, []byte, etc.
type HashableInput []byte

// Hashing structure to perform hash operations
type Hashing struct{}

// numberToHex converts a number to hexadecimal string
func numberToHex(num uint64) string {
	return strconv.FormatUint(num, 16)
}

// DoubleHashing applies enhanced double hashing to produce an n-th hash
func (h *Hashing) DoubleHashing(n int, hashA, hashB, size uint64) uint64 {
	return uint64((hashA + uint64(n)*hashB + uint64((n*n*n-n)/6)) % size)
}

// GetDistinctIndexes generates a set of distinct indexes using double hashing
func (h *Hashing) GetDistinctIndexes(element HashableInput, size, number int, seed *uint64) []int {
	if seed == nil {
		defaultSeed := utils.GetDefaultSeed()
		seed = &defaultSeed
	}
	n := 0
	indexes := make(map[int]struct{})
	hashes := h.HashTwice(element, *seed)
	for len(indexes) < number {
		ind := int(hashes.First % uint64(size))
		if _, exists := indexes[ind]; !exists {
			indexes[ind] = struct{}{}
		}
		hashes.First = (hashes.First + hashes.Second) % uint64(size)
		hashes.Second = (hashes.Second + uint64(n)) % uint64(size)
		n++

		if n > size {
			*seed++
			hashes = h.HashTwice(element, *seed)
		}
	}

	var result []int
	for key := range indexes {
		result = append(result, key)
	}
	return result
}

// GetIndexes generates N indexes on range [0, size) using double hashing
func (h *Hashing) GetIndexes(element HashableInput, size, hashCount int, seed *uint64) []int {
	if seed == nil {
		defaultSeed := utils.GetDefaultSeed()
		seed = &defaultSeed
	}
	hashes := h.HashTwice(element, *seed)
	var result []int
	for i := 0; i < hashCount; i++ {
		result = append(result, int(h.DoubleHashing(i, hashes.First, hashes.Second, uint64(size))))
	}
	return result
}

// Serialize hashes an element into a uint64
func (h *Hashing) Serialize(element HashableInput, seed uint64) uint64 {
	if seed == 0 {
		seed = utils.GetDefaultSeed()
	}
	hasher := xxhash.NewWithSeed(seed)
	hasher.Write(element)
	return hasher.Sum64()
}

// HashTwice hashes a value into two values
func (h *Hashing) HashTwice(value HashableInput, seed uint64) TwoHashes {
	return TwoHashes{
		First:  h.Serialize(value, seed+1),
		Second: h.Serialize(value, seed+2),
	}
}

// HashTwiceAsString hashes a value and returns two hashes as hexadecimal strings
func (h *Hashing) HashTwiceAsString(value HashableInput, seed uint64) TwoHashesTemplated[string] {
	hashes := h.HashTwice(value, seed)
	return TwoHashesTemplated[string]{
		First:  numberToHex(hashes.First),
		Second: numberToHex(hashes.Second),
	}
}

// HashTwiceIntAndString hashes a value and returns the results as both integer and string
func (h *Hashing) HashTwiceIntAndString(value HashableInput, seed uint64) TwoHashesIntAndString {
	hashOne := h.HashIntAndString(value, seed+1)
	hashTwo := h.HashIntAndString(value, seed+2)
	return TwoHashesIntAndString{
		Int: TwoHashesTemplated[uint64]{
			First:  hashOne.First,
			Second: hashTwo.First,
		},
		String: TwoHashesTemplated[string]{
			First:  numberToHex(hashOne.First),
			Second: numberToHex(hashTwo.First),
		},
	}
}

// HashAsInt hashes an element and returns it as a uint64
func (h *Hashing) HashAsInt(elem HashableInput, seed uint64) uint64 {
	return h.Serialize(elem, seed)
}

// HashIntAndString hashes an element and returns both the integer and string (hex) representations
func (h *Hashing) HashIntAndString(elem HashableInput, seed uint64) TwoHashesTemplated[uint64] {
	hash := h.HashAsInt(elem, seed)
	return TwoHashesTemplated[uint64]{
		First:  hash,
		Second: hash,
	}
}

// func main() {
// 	hashing := Hashing{}
// 	element := []byte("example data")
// 	indexes := hashing.GetIndexes(element, 100, 5, nil)
// 	fmt.Println("Generated indexes:", indexes)

// 	twoHashes := hashing.HashTwice(element, utils.GetDefaultSeed())
// 	fmt.Printf("First hash: %d, Second hash: %d\n", twoHashes.First, twoHashes.Second)

// 	twoHashesString := hashing.HashTwiceAsString(element, utils.GetDefaultSeed())
// 	fmt.Printf("First hash (hex): %s, Second hash (hex): %s\n", twoHashesString.First, twoHashesString.Second)
// }
