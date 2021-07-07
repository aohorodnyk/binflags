package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagMapInt32(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagMapInt32() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagMapInt32(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagMapInt32(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagMapInt32() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagMapInt32(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagMapInt32 struct {
	name     string
	flags    map[uint64]int32
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapInt32 struct {
	name     string
	flags    map[uint64]int32
	flag     uint64
	set      bool
	expected map[uint64]int32
	err      error
}

func providerHasFlagMapInt32() []providerTypeHasFlagMapInt32 {
	return []providerTypeHasFlagMapInt32{
		{
			name:     "Flags {}, flag 0",
			flags:    map[uint64]int32{},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags {}, flag 12425",
			flags:    map[uint64]int32{},
			flag:     12425,
			expected: false,
		},
		{
			name:     "Flags nil, flag 0",
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags nil, flag 12425",
			flags:    nil,
			flag:     12425,
			expected: false,
		},
		{
			name:     "Flags nil, flag 2356241245436",
			flags:    map[uint64]int32{},
			flag:     2356241245436,
			expected: false,
		},
		{
			name:     "Flags {0: 1}, flag 0",
			flags:    map[uint64]int32{0: 1},
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags {0: 32768}, flag 15",
			flags:    map[uint64]int32{0: 32768},
			flag:     15,
			expected: true,
		},
		{
			name:     "Flags {0: -2147483648}, flag 31",
			flags:    map[uint64]int32{0: -2147483648},
			flag:     31,
			expected: true,
		},
		{
			name:     "Flags {0: -2147483648, 1: 1}, flag 32",
			flags:    map[uint64]int32{0: -2147483648, 1: 1},
			flag:     32,
			expected: true,
		},
		{
			name:     "Flags {0: -2147483648, 1: 1449509}, flag 52",
			flags:    map[uint64]int32{0: -2147483648, 1: 1449509},
			flag:     52,
			expected: true,
		},
		{
			name:     "Flags {0: -2147483648, 1323: 1449509}, flag 42356",
			flags:    map[uint64]int32{0: -2147483648, 1323: 1449509},
			flag:     42356,
			expected: true,
		},
	}
}

func providerSetFlagMapInt32() []providerTypeSetFlagMapInt32 {
	return []providerTypeSetFlagMapInt32{
		{
			name:     "Flags {}, flag 0, set",
			flags:    map[uint64]int32{},
			flag:     0,
			set:      true,
			expected: map[uint64]int32{0: 1},
		},
		{
			name:     "Flags {}, flag 7, set",
			flags:    map[uint64]int32{},
			flag:     7,
			set:      true,
			expected: map[uint64]int32{0: 128},
		},
		{
			name:     "Flags {}, flag 15, set",
			flags:    map[uint64]int32{},
			flag:     15,
			set:      true,
			expected: map[uint64]int32{0: 32768},
		},
		{
			name:     "Flags {0: 92}, flag 7, set",
			flags:    map[uint64]int32{0: 92},
			flag:     7,
			set:      true,
			expected: map[uint64]int32{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 7, unset",
			flags:    map[uint64]int32{0: 220},
			flag:     7,
			set:      false,
			expected: map[uint64]int32{0: 92},
		},
		{
			name:     "Flags {0: 220}, flag 7, set",
			flags:    map[uint64]int32{0: 220},
			flag:     7,
			set:      true,
			expected: map[uint64]int32{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 1, unset",
			flags:    map[uint64]int32{0: 220},
			flag:     1,
			set:      false,
			expected: map[uint64]int32{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 14, set",
			flags:    map[uint64]int32{0: 220},
			flag:     14,
			set:      true,
			expected: map[uint64]int32{0: 16604},
		},
		{
			name:     "Flags {0: 16604}, flag 31, set",
			flags:    map[uint64]int32{0: 16604},
			flag:     31,
			set:      true,
			expected: map[uint64]int32{0: -2147467044},
		},
		{
			name:     "Flags {0: -2147467044}, flag 14, set",
			flags:    map[uint64]int32{0: -2147467044},
			flag:     14,
			set:      false,
			expected: map[uint64]int32{0: -2147483428},
		},
		{
			name:     "Flags {0: -2147467044}, flag 5234, set",
			flags:    map[uint64]int32{0: -2147467044},
			flag:     5234,
			set:      true,
			expected: map[uint64]int32{0: -2147467044, 163: 262144},
		},
		{
			name:     "Flags {0: 436324, 1: 1, 2: 25623541}, flag 32, unset",
			flags:    map[uint64]int32{0: 436324, 1: 1, 2: 25623541},
			flag:     32,
			set:      false,
			expected: map[uint64]int32{0: 436324, 2: 25623541},
		},
		{
			name:     "Flags nil, flag 0, unset",
			flags:    nil,
			flag:     0,
			set:      false,
			expected: nil,
			err:      binflags.ErrorFlagsMapNil(binflags.ErrorMsgFlagsMapNil),
		},
		{
			name:     "Flags nil, flag 0, set",
			flags:    nil,
			flag:     0,
			set:      true,
			expected: nil,
			err:      binflags.ErrorFlagsMapNil(binflags.ErrorMsgFlagsMapNil),
		},
	}
}
