package binflags_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagUint32(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagUint32() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagUint32(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagUint32(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagUint32() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual, err := binflags.SetFlagUint32(prov.flags, prov.flag, prov.set)
			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %d\nActual: %d", prov.expected, actual)
			}
		})
	}
}

type providerTypeHasFlagUint32 struct {
	name     string
	flags    uint32
	flag     uint8
	expected bool
}

type providerTypeSetFlagUint32 struct {
	name     string
	flags    uint32
	flag     uint8
	set      bool
	expected uint32
	err      error
}

func providerHasFlagUint32() []providerTypeHasFlagUint32 {
	return []providerTypeHasFlagUint32{
		{
			name:     "Flags 0, flag 0",
			flags:    0,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags 0, flag 1",
			flags:    0,
			flag:     1,
			expected: false,
		},
		{
			name:     "Flags 0, flag 3",
			flags:    0,
			flag:     3,
			expected: false,
		},
		{
			name:     "Flags 0, flag 5",
			flags:    0,
			flag:     5,
			expected: false,
		},
		{
			name:     "Flags 0, flag 7",
			flags:    0,
			flag:     7,
			expected: false,
		},
		{
			name:     "Flags 2, flag 0",
			flags:    2,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags 1, flag 0",
			flags:    1,
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags 3, flag 0",
			flags:    3,
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags 26, flag 1",
			flags:    26,
			flag:     1,
			expected: true,
		},
		{
			name:     "Flags 26, flag 3",
			flags:    26,
			flag:     3,
			expected: true,
		},
		{
			name:     "Flags 26, flag 4",
			flags:    26,
			flag:     4,
			expected: true,
		},
		{
			name:     "Flags 128, flag 7",
			flags:    128,
			flag:     7,
			expected: true,
		},
		{
			name:     "Flags 64, flag 6",
			flags:    64,
			flag:     6,
			expected: true,
		},
		{
			name:     "Flags 30000, flag 4",
			flags:    30000,
			flag:     4,
			expected: true,
		},
		{
			name:     "Flags 30000, flag 14",
			flags:    30000,
			flag:     14,
			expected: true,
		},
		{
			name:     "Flags 30000, flag 3",
			flags:    30000,
			flag:     3,
			expected: false,
		},
		{
			name:     "Flags 4294967295, flag 0",
			flags:    math.MaxUint32,
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags 4294967295, flag 1",
			flags:    math.MaxUint32,
			flag:     1,
			expected: true,
		},
		{
			name:     "Flags 4294967295, flag 3",
			flags:    math.MaxUint32,
			flag:     3,
			expected: true,
		},
		{
			name:     "Flags 4294967295, flag 5",
			flags:    math.MaxUint32,
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags 4294967295, flag 7",
			flags:    math.MaxUint32,
			flag:     7,
			expected: true,
		},
		{
			name:     "Flags 4294967295, flag 31",
			flags:    math.MaxUint32,
			flag:     31,
			expected: true,
		},
		{
			name:     "Flags 4294967295, flag 32",
			flags:    math.MaxUint32,
			flag:     32,
			expected: false,
		},
	}
}

func providerSetFlagUint32() []providerTypeSetFlagUint32 {
	return []providerTypeSetFlagUint32{
		{
			flags:    0,
			flag:     0,
			set:      true,
			expected: 1,
			err:      nil,
		},
		{
			flags:    0,
			flag:     1,
			set:      true,
			expected: 2,
			err:      nil,
		},
		{
			flags:    0,
			flag:     2,
			set:      true,
			expected: 4,
			err:      nil,
		},
		{
			flags:    1,
			flag:     2,
			set:      true,
			expected: 5,
			err:      nil,
		},
		{
			flags:    53,
			flag:     1,
			set:      true,
			expected: 55,
			err:      nil,
		},
		{
			flags:    53,
			flag:     0,
			set:      true,
			expected: 53,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     2,
			set:      true,
			expected: math.MaxUint32,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     2,
			set:      false,
			expected: math.MaxUint32 - 4,
			err:      nil,
		},
		{
			flags:    235,
			flag:     5,
			set:      false,
			expected: 203,
			err:      nil,
		},
		{
			flags:    62768,
			flag:     12,
			set:      false,
			expected: 58672,
			err:      nil,
		},
		{
			flags:    5,
			flag:     7,
			set:      true,
			expected: 133,
			err:      nil,
		},
		{
			flags:    128,
			flag:     7,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    32768,
			flag:     15,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    2147483648,
			flag:     31,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    3467463648,
			flag:     26,
			set:      false,
			expected: 3400354784,
			err:      nil,
		},
		{
			flags:    5,
			flag:     32,
			set:      true,
			expected: 5,
			err:      binflags.ErrorOutOfRange(binflags.ErrorMsgOutOfRange),
		},
	}
}
