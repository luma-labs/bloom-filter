package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Error message for buffer issues
var BufferError = errors.New("The buffer class must be available, consider using the bytes package in Go")

// TwoHashes holds the result of two hash functions applied to a value
type TwoHashes struct {
	First  uint64
	Second uint64
}

// TwoHashesTemplated allows templated types for hash results
type TwoHashesTemplated[T any] struct {
	First  T
	Second T
}

// TwoHashesIntAndString holds the result of hash functions in both integer and string formats
type TwoHashesIntAndString struct {
	Int    TwoHashesTemplated[uint64]
	String TwoHashesTemplated[string]
}

// AllocateArray creates a new array filled with a base value
func AllocateArray[T any](size int, defaultValue T) []T {
	arr := make([]T, size)
	for i := range arr {
		arr[i] = defaultValue
	}
	return arr
}

// AllocateArrayWithFunction creates a new array filled by invoking a function to get the default value
func AllocateArrayWithFunction[T any](size int, defaultValueFunc func() T) []T {
	arr := make([]T, size)
	for i := range arr {
		arr[i] = defaultValueFunc()
	}
	return arr
}

// NumberToHex converts a number to a hexadecimal string, padded with zeroes if necessary
func NumberToHex(num uint64) string {
	hexStr := fmt.Sprintf("%x", num)
	padding := 4 - (len(hexStr) % 4)
	if padding < 4 {
		hexStr = fmt.Sprintf("%0*s", padding+len(hexStr), hexStr)
	}
	return hexStr
}

// RandomInt generates a random integer between two bounds (inclusive)
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// XORBuffer returns the non-destructive XOR of two byte slices
func XORBuffer(a, b []byte) []byte {
	length := max(len(a), len(b))
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		if i < len(a) && i < len(b) {
			result[length-i-1] = a[len(a)-i-1] ^ b[len(b)-i-1]
		} else if i < len(a) {
			result[length-i-1] ^= a[len(a)-i-1]
		} else if i < len(b) {
			result[length-i-1] ^= b[len(b)-i-1]
		}
	}

	// Remove leading zeros
	for len(result) > 0 && result[0] == 0 {
		result = result[1:]
	}

	return result
}

// IsEmptyBuffer returns true if the buffer is empty (i.e., all values are zero)
func IsEmptyBuffer(buffer []byte) bool {
	for _, value := range buffer {
		if value != 0 {
			return false
		}
	}
	return true
}

// GetDefaultSeed returns a default seed
func GetDefaultSeed() uint64 {
	return 0x1234567890
}

// Helper function to get the max of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
