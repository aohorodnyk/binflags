package binflags_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagInt16(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagInt16() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagInt16(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagInt16(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagInt16() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual, err := binflags.SetFlagInt16(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %d\nActual: %d", prov.expected, actual)
			}
		})
	}
}

type providerTypeHasFlagInt16 struct {
	name     string
	flags    int16
	flag     uint8
	expected bool
}

type providerTypeSetFlagInt16 struct {
	name     string
	flags    int16
	flag     uint8
	set      bool
	expected int16
	err      error
}

func providerHasFlagInt16() []providerTypeHasFlagInt16 {
	return []providerTypeHasFlagInt16{
		{
			name:     "Flags 0, flag 0",
			flags:    0,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags 52, flag 7",
			flags:    52,
			flag:     7,
			expected: false,
		},
		{
			name:     "Flags 52, flag 15",
			flags:    52,
			flag:     15,
			expected: false,
		},
		{
			name:     "Flags -32768, flag 15",
			flags:    math.MinInt16,
			flag:     15,
			expected: true,
		},
		{
			name:     "Flags -32768, flag 14",
			flags:    math.MinInt16,
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
			name:     "Flags -1, flag 16",
			flags:    -1,
			flag:     16,
			expected: false,
		},
	}
}

func providerSetFlagInt16() []providerTypeSetFlagInt16 {
	return []providerTypeSetFlagInt16{
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
			name:     "Flags -1, flag 15, unset",
			flags:    -1,
			flag:     15,
			set:      false,
			expected: 32767,
			err:      nil,
		},
		{
			name:     "Flags -32768, flag 15, unset",
			flags:    math.MinInt16,
			flag:     15,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags -32768, flag 16, unset",
			flags:    math.MinInt16,
			flag:     16,
			set:      false,
			expected: math.MinInt16,
			err:      binflags.ErrorOutOfRange(binflags.ErrorMsgOutOfRange),
		},
	}
}
