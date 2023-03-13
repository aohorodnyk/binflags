package binflags

import (
	"unsafe"
)

// Dynamic is a dynamic size implementation of the binflags.
// This implementation uses a map under the hood.
// The map's key is a uint.
// Specify an expected size for the map to improve performance.
// To initialize the Dynamic type, use make function or usual instatination:
//
//	flags := make(Dynamic[uint], 3)
//
// or
//
//	flags := make(Dynamic[uint])
//
// or
//
//	flags := Dynamic[uint]{}
type Dynamic[T Integer] map[uint]T

// IsSet returns true if the flag is set.
// If the flag is not set, then it returns false.
// This function has O(1) complexity.
func (d Dynamic[T]) IsSet(bitNumber uint) bool {
	size := uint(unsafe.Sizeof(d[0])) * BitsInByte
	byteIdx := bitNumber / size

	val, ok := d[byteIdx]
	if !ok {
		return false
	}

	bitIdx := uint8(bitNumber % size)

	return IsSetBit(val, bitIdx)
}

// Set sets the flag.
// This function is idempotent.
// This function has O(1) complexity.
// It always returns true.
func (d Dynamic[T]) Set(bitNumber uint) bool {
	size := uint(unsafe.Sizeof(d[0])) * BitsInByte
	byteIdx := bitNumber / size
	bitIdx := uint8(bitNumber % size)

	d[byteIdx] = SetBit(d[byteIdx], bitIdx)

	return true
}

// Unset unsets the flag.
// This function is idempotent.
// This function has O(1) complexity.
// It always returns true.
func (d Dynamic[T]) Unset(bitNumber uint) bool {
	size := uint(unsafe.Sizeof(d[0])) * BitsInByte
	byteIdx := bitNumber / size

	val := d[byteIdx]
	if val == 0 {
		return true
	}

	bitIdx := uint8(bitNumber % size)

	val = UnsetBit(val, bitIdx)
	if val == 0 {
		delete(d, byteIdx)
	} else {
		d[byteIdx] = val
	}

	return true
}
