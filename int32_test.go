package binflags_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagInt32(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagInt32() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagInt32(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagInt32(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagInt32() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual, err := binflags.SetFlagInt32(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %d\nActual: %d", prov.expected, actual)
			}
		})
	}
}

type providerTypeHasFlagInt32 struct {
	name     string
	flags    int32
	flag     uint8
	expected bool
}

type providerTypeSetFlagInt32 struct {
	name     string
	flags    int32
	flag     uint8
	set      bool
	expected int32
	err      error
}

func providerHasFlagInt32() []providerTypeHasFlagInt32 {
	return []providerTypeHasFlagInt32{
		{
			name:     "Flags 0, flag 0",
			flags:    0,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags 52, flag 15",
			flags:    52,
			flag:     15,
			expected: false,
		},
		{
			name:     "Flags 52, flag 4",
			flags:    52,
			flag:     4,
			expected: true,
		},
		{
			name:     "Flags -2147483648, flag 31",
			flags:    math.MinInt32,
			flag:     31,
			expected: true,
		},
		{
			name:     "Flags -2147483648, flag 14",
			flags:    math.MinInt32,
			flag:     14,
			expected: false,
		},
		{
			name:     "Flags -1, flag 3",
			flags:    -1,
			flag:     3,
			expected: true,
		},
		{
			name:     "Flags -1, flag 7",
			flags:    -1,
			flag:     7,
			expected: true,
		},
		{
			name:     "Flags -1, flag 31",
			flags:    -1,
			flag:     31,
			expected: true,
		},
		{
			name:     "Flags -1, flag 32",
			flags:    -1,
			flag:     32,
			expected: false,
		},
	}
}

func providerSetFlagInt32() []providerTypeSetFlagInt32 {
	return []providerTypeSetFlagInt32{
		{
			name:     "Flags 0, flag 0, unset",
			flags:    0,
			flag:     0,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags 0, flag 0, set",
			flags:    0,
			flag:     0,
			set:      true,
			expected: 1,
			err:      nil,
		},
		{
			name:     "Flags 1, flag 1, set",
			flags:    1,
			flag:     1,
			set:      true,
			expected: 3,
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
			name:     "Flags 53, flag 0, unset",
			flags:    53,
			flag:     0,
			set:      false,
			expected: 52,
			err:      nil,
		},
		{
			name:     "Flags -1, flag 5, set",
			flags:    -1,
			flag:     5,
			set:      true,
			expected: -1,
			err:      nil,
		},
		{
			name:     "Flags 2567, flag 11, unset",
			flags:    2567,
			flag:     11,
			set:      false,
			expected: 519,
			err:      nil,
		},
		{
			name:     "Flags -1, flag 5, unset",
			flags:    -1,
			flag:     5,
			set:      false,
			expected: -33,
			err:      nil,
		},
		{
			name:     "Flags -1, flag 31, unset",
			flags:    -1,
			flag:     31,
			set:      false,
			expected: 2147483647,
			err:      nil,
		},
		{
			name:     "Flags -2147483648, flag 31, unset",
			flags:    math.MinInt32,
			flag:     31,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags -2147483648, flag 32, set",
			flags:    math.MinInt32,
			flag:     32,
			set:      true,
			expected: math.MinInt32,
			err:      binflags.ErrorOutOfRange(binflags.ErrorMsgOutOfRange),
		},
	}
}
