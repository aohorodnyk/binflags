package binflags

import (
	"unsafe"
)

// Fixed is a fixed size implementation of the binflags.
// This is a slice that is allocated to collect flags in bits.
// The max amount of flags can be stored per slice can be calculated by the formula: `slice length` * `bits in byte`.
//
// This implementation can be used whenever you want to have a predictable size and performance.
// Also it can be saved in a database as an array type.
// To initialize the Dynamic type, use make function or usual instatination:
//
//	flags := make(Fixed[uint], 3)
//
// or
//
//	flags := Fixed[uint]{0, 0, 0}
type Fixed[T Integer] []T

// IsSet returns true if the flag is set.
// If the flag is not set, then it returns false.
// This function has O(1) complexity.
func (s Fixed[T]) IsSet(bitNumber uint) bool {
	if len(s) == 0 {
		return false
	}

	size := uint(unsafe.Sizeof(s[0])) * BitsInByte

	byteIdx := int(bitNumber / size)
	if byteIdx >= len(s) {
		// Bit number is out of range.
		return false
	}

	bitIdx := uint8(bitNumber % size)

	return IsSetBit(s[byteIdx], bitIdx)
}

// Set sets the flag.
// This function is idempotent.
// This function has O(1) complexity.
// It returns false if the flag is out of range.
func (s Fixed[T]) Set(bitNumber uint) bool {
	if len(s) == 0 {
		return false
	}

	size := uint(unsafe.Sizeof(s[0])) * BitsInByte

	byteIdx := int(bitNumber / size)
	if byteIdx >= len(s) {
		// Bit number is out of range.
		return false
	}

	bitIdx := uint8(bitNumber % size)
	s[byteIdx] = SetBit(s[byteIdx], bitIdx)

	return true
}

// Unset unsets the flag.
// This function is idempotent.
// This function has O(1) complexity.
// It returns false if the flag is out of range.
func (s Fixed[T]) Unset(bitNumber uint) bool {
	if len(s) == 0 {
		return false
	}

	size := uint(unsafe.Sizeof(s[0])) * BitsInByte

	byteIdx := int(bitNumber / size)
	if byteIdx >= len(s) {
		// Bit number is out of range.
		return false
	}

	bitIdx := uint8(bitNumber % size)
	s[byteIdx] = UnsetBit(s[byteIdx], bitIdx)

	return true
}
