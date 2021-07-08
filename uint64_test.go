package binflags_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagUint64(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagUint64() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagUint64(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagUint64(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagUint64() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual, err := binflags.SetFlagUint64(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %d\nActual: %d", prov.expected, actual)
			}
		})
	}
}

type providerTypeHasFlagUint64 struct {
	name     string
	flags    uint64
	flag     uint8
	expected bool
}

type providerTypeSetFlagUint64 struct {
	name     string
	flags    uint64
	flag     uint8
	set      bool
	expected uint64
	err      error
}

func providerHasFlagUint64() []providerTypeHasFlagUint64 {
	return []providerTypeHasFlagUint64{
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
			name:     "Flags 822372036854775808, flag 59",
			flags:    822372036854775808,
			flag:     59,
			expected: true,
		},
		{
			name:     "Flags 18446744073709551615, flag 0",
			flags:    math.MaxUint64,
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags 18446744073709551615, flag 1",
			flags:    math.MaxUint64,
			flag:     1,
			expected: true,
		},
		{
			name:     "Flags 18446744073709551615, flag 3",
			flags:    math.MaxUint64,
			flag:     3,
			expected: true,
		},
		{
			name:     "Flags 18446744073709551615, flag 5",
			flags:    math.MaxUint64,
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags 18446744073709551615, flag 7",
			flags:    math.MaxUint64,
			flag:     7,
			expected: true,
		},
		{
			name:     "Flags 18446744073709551615, flag 15",
			flags:    math.MaxUint64,
			flag:     15,
			expected: true,
		},
		{
			name:     "Flags 18446744073709551615, flag 31",
			flags:    math.MaxUint64,
			flag:     31,
			expected: true,
		},
		{
			name:     "Flags 18446744073709551615, flag 53",
			flags:    math.MaxUint64,
			flag:     53,
			expected: true,
		},
		{
			name:     "Flags 18446744073709551615, flag 63",
			flags:    math.MaxUint64,
			flag:     63,
			expected: true,
		},
		{
			name:     "Flags 18446744073709551615, flag 64",
			flags:    math.MaxUint64,
			flag:     64,
			expected: false,
		},
	}
}

func providerSetFlagUint64() []providerTypeSetFlagUint64 {
	return []providerTypeSetFlagUint64{
		{
			name:     "Flags 0, flag 0, set",
			flags:    0,
			flag:     0,
			set:      true,
			expected: 1,
			err:      nil,
		},
		{
			name:     "Flags 0, flag 1, set",
			flags:    0,
			flag:     1,
			set:      true,
			expected: 2,
			err:      nil,
		},
		{
			name:     "Flags 0, flag 2, set",
			flags:    0,
			flag:     2,
			set:      true,
			expected: 4,
			err:      nil,
		},
		{
			name:     "Flags 1, flag 2, set",
			flags:    1,
			flag:     2,
			set:      true,
			expected: 5,
			err:      nil,
		},
		{
			name:     "Flags 53, flag 1, set",
			flags:    53,
			flag:     1,
			set:      true,
			expected: 55,
			err:      nil,
		},
		{
			name:     "Flags 1, flag 0, set",
			flags:    53,
			flag:     0,
			set:      true,
			expected: 53,
			err:      nil,
		},
		{
			name:     "Flags 18446744073709551615, flag 2, set",
			flags:    math.MaxUint64,
			flag:     2,
			set:      true,
			expected: math.MaxUint64,
			err:      nil,
		},
		{
			name:     "Flags 18446744073709551615, flag 2, unset",
			flags:    math.MaxUint64,
			flag:     2,
			set:      false,
			expected: math.MaxUint64 - 4,
			err:      nil,
		},
		{
			name:     "Flags 5, flag 7, set",
			flags:    5,
			flag:     7,
			set:      true,
			expected: 133,
			err:      nil,
		},
		{
			name:     "Flags 128, flag 7, unset",
			flags:    128,
			flag:     7,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags 32768, flag 15, unset",
			flags:    32768,
			flag:     15,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags 235, flag 5, unset",
			flags:    235,
			flag:     5,
			set:      false,
			expected: 203,
			err:      nil,
		},
		{
			name:     "Flags 62768, flag 12, unset",
			flags:    62768,
			flag:     12,
			set:      false,
			expected: 58672,
			err:      nil,
		},
		{
			name:     "Flags 3467463648, flag 26, unset",
			flags:    3467463648,
			flag:     26,
			set:      false,
			expected: 3400354784,
			err:      nil,
		},
		{
			name:     "Flags 2147483648, flag 31, unset",
			flags:    2147483648,
			flag:     31,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags 2147483648, flag 31, unset",
			flags:    2147483648,
			flag:     31,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags 822372036854775808, flag 59, unset",
			flags:    822372036854775808,
			flag:     59,
			set:      false,
			expected: 245911284551352320,
			err:      nil,
		},
		{
			name:     "Flags 9223372036854775808, flag 63, unset",
			flags:    9223372036854775808,
			flag:     63,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags 235234, flag 64, unset",
			flags:    235234,
			flag:     64,
			set:      true,
			expected: 235234,
			err:      binflags.ErrorOutOfRange(binflags.ErrorMsgOutOfRange),
		},
	}
}
