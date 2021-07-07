package binflags_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagInt8(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagInt8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagInt8(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagInt8(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagInt8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual, err := binflags.SetFlagInt8(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %d\nActual: %d", prov.expected, actual)
			}
		})
	}
}

type providerTypeHasFlagInt8 struct {
	name     string
	flags    int8
	flag     uint8
	expected bool
}

type providerTypeSetFlagInt8 struct {
	name     string
	flags    int8
	flag     uint8
	set      bool
	expected int8
	err      error
}

func providerHasFlagInt8() []providerTypeHasFlagInt8 {
	return []providerTypeHasFlagInt8{
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
			name:     "Flags -128, flag 7",
			flags:    math.MinInt8,
			flag:     7,
			expected: true,
		},
		{
			name:     "Flags -128, flag 3",
			flags:    math.MinInt8,
			flag:     3,
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
			name:     "Flags -1, flag 8",
			flags:    -1,
			flag:     8,
			expected: false,
		},
	}
}

func providerSetFlagInt8() []providerTypeSetFlagInt8 {
	return []providerTypeSetFlagInt8{
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
			name:     "Flags -1, flag 5, unset",
			flags:    -1,
			flag:     5,
			set:      false,
			expected: -33,
			err:      nil,
		},
		{
			name:     "Flags -1, flag 7, unset",
			flags:    -1,
			flag:     7,
			set:      false,
			expected: 127,
			err:      nil,
		},
		{
			name:     "Flags -128, flag 7, unset",
			flags:    math.MinInt8,
			flag:     7,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			name:     "Flags -128, flag 8, unset",
			flags:    math.MinInt8,
			flag:     8,
			set:      false,
			expected: math.MinInt8,
			err:      binflags.ErrorOutOfRange(binflags.ErrorMsgOutOfRange),
		},
	}
}
