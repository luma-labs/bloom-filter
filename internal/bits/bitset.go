package bits

import (
	"fmt"
	"math"
)

type Bitset struct {
	size  int
	array []uint8
}

const bitsPerWord = 8

func NewBitset(size int) *Bitset {
	if size % bitsPerWord != 0 {
		differeceToroundOff := bitsPerWord - (size % bitsPerWord)
		size += differeceToroundOff
	}
	return &Bitset{
		size:  size,
		array: make([]uint8, int(math.Ceil(float64(size)/float64(bitsPerWord)))),
	}
}


func (b *Bitset) Add(index int) {
	wordIndex := index / bitsPerWord
	bitIndex := 1 << (index % bitsPerWord)
	b.array[wordIndex] |= uint8(bitIndex)
}

func (b *Bitset) Contains(index int) bool {
	wordIndex := index / bitsPerWord
	bitIndex := 1 << (index % bitsPerWord)
	return b.array[wordIndex]&uint8(bitIndex) != 0
}

func (b *Bitset) MaxTrueBit() int {
	for i := len(b.array) - 1; i >= 0; i-- {
		if b.array[i] != 0 {
			for j := bitsPerWord - 1; j >= 0; j-- {
				if b.array[i]&(1<<uint(j)) != 0 {
					return i*bitsPerWord + j
				}
			}
		}
	}
	return -1
}

func (b *Bitset) countTrueBits(bits uint8) int {
	count := 0
	for bits != 0 {
		count++
		bits &= bits - 1
	}
	return count
}

func (b *Bitset) NumOfTrueBits() int {
	count := 0
	for _, bits := range b.array {
		count += b.countTrueBits(bits)
	}
	return count
}

func (b *Bitset) Equals(other *Bitset) bool {
	if b.size != other.size {
		return false
	}
	for i := 0; i < len(b.array); i++ {
		if b.array[i] != other.array[i] {
			return false
		}
	}
	return true
}

func (b *Bitset) PrintEverything() int {
	for i := 0; i < len(b.array); i++ {
		fmt.Printf("%08b ", b.array[i])
	}
	fmt.Println()
	return 0
}