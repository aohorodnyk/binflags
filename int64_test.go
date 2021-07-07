package binflags_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagInt64(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagInt64() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagInt64(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagInt64(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagInt64() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual, err := binflags.SetFlagInt64(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %d\nActual: %d", prov.expected, actual)
			}
		})
	}
}

type providerTypeHasFlagInt64 struct {
	name     string
	flags    int64
	flag     uint8
	expected bool
}

type providerTypeSetFlagInt64 struct {
	name     string
	flags    int64
	flag     uint8
	set      bool
	expected int64
	err      error
}

func providerHasFlagInt64() []providerTypeHasFlagInt64 {
	return []providerTypeHasFlagInt64{
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
			name:     "Flags -9223372036854775808, flag 31",
			flags:    math.MinInt64,
			flag:     31,
			expected: false,
		},
		{
			name:     "Flags -9223372036854775808, flag 63",
			flags:    math.MinInt64,
			flag:     63,
			expected: true,
		},
		{
			name:     "Flags -9223372036854775808, flag 14",
			flags:    math.MinInt64,
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
			name:     "Flags -1, flag 64",
			flags:    -1,
			flag:     64,
			expected: false,
		},
	}
}

func providerSetFlagInt64() []providerTypeSetFlagInt64 {
	return []providerTypeSetFlagInt64{
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
			expected: -2147483649,
			err:      nil,
		},
		{
			name:     "Flags -1, flag 63, unset",
			flags:    -1,
			flag:     63,
			set:      false,
			expected: 9223372036854775807,
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
			name:     "Flags -9223372036854775808, flag 63, unset",
			flags:    math.MinInt64,
			flag:     63,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags -9223372036854775808, flag 64, unset",
			flags:    math.MinInt64,
			flag:     64,
			set:      true,
			expected: math.MinInt64,
			err:      binflags.ErrorOutOfRange(binflags.ErrorMsgOutOfRange),
		},
	}
}
