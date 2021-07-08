package binflags_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagUint16(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagUint16() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagUint16(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagUint16(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagUint16() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual, err := binflags.SetFlagUint16(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %d\nActual: %d", prov.expected, actual)
			}
		})
	}
}

type providerTypeHasFlagUint16 struct {
	name     string
	flags    uint16
	flag     uint8
	expected bool
}

type providerTypeSetFlagUint16 struct {
	name     string
	flags    uint16
	flag     uint8
	set      bool
	expected uint16
	err      error
}

func providerHasFlagUint16() []providerTypeHasFlagUint16 {
	return []providerTypeHasFlagUint16{
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
			name:     "Flags 65535, flag 0",
			flags:    math.MaxUint16,
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags 65535, flag 1",
			flags:    math.MaxUint16,
			flag:     1,
			expected: true,
		},
		{
			name:     "Flags 65535, flag 3",
			flags:    math.MaxUint16,
			flag:     3,
			expected: true,
		},
		{
			name:     "Flags 65535, flag 5",
			flags:    math.MaxUint16,
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags 65535, flag 7",
			flags:    math.MaxUint16,
			flag:     7,
			expected: true,
		},
		{
			name:     "Flags 65535, flag 16",
			flags:    math.MaxUint16,
			flag:     16,
			expected: false,
		},
	}
}

func providerSetFlagUint16() []providerTypeSetFlagUint16 {
	return []providerTypeSetFlagUint16{
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
			name:     "Flags 53, flag 0, set",
			flags:    53,
			flag:     0,
			set:      true,
			expected: 53,
			err:      nil,
		},
		{
			name:     "Flags 65535, flag 2, set",
			flags:    math.MaxUint16,
			flag:     2,
			set:      true,
			expected: math.MaxUint16,
			err:      nil,
		},
		{
			name:     "Flags 65535, flag 2, unset",
			flags:    math.MaxUint16,
			flag:     2,
			set:      false,
			expected: math.MaxUint16 - 4,
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
			name:     "Flags 32768, flag 15, unset",
			flags:    32768,
			flag:     15,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags 5, flag 16, set",
			flags:    5,
			flag:     16,
			set:      true,
			expected: 5,
			err:      binflags.ErrorOutOfRange(binflags.ErrorMsgOutOfRange),
		},
	}
}
