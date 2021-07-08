package binflags_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagUint8(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagUint8(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagUint8(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual, err := binflags.SetFlagUint8(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %d\nActual: %d", prov.expected, actual)
			}
		})
	}
}

func TestHasFlagByte(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagByte(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagByte(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual, err := binflags.SetFlagByte(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if prov.expected != actual {
				t.Fatalf("Unexpected flags.\nExpected: %d\nActual: %d", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagUint8 struct {
	name     string
	flags    uint8
	flag     uint8
	expected bool
}

type providerTypeSetFlagUint8 struct {
	name     string
	flags    uint8
	flag     uint8
	set      bool
	expected uint8
	err      error
}

func providerHasFlagUint8() []providerTypeHasFlagUint8 {
	return []providerTypeHasFlagUint8{
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
			name:     "Flags 82, flag 6",
			flags:    82,
			flag:     6,
			expected: true,
		},
		{
			name:     "Flags 255, flag 0",
			flags:    math.MaxUint8,
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags 255, flag 1",
			flags:    math.MaxUint8,
			flag:     1,
			expected: true,
		},
		{
			name:     "Flags 255, flag 3",
			flags:    math.MaxUint8,
			flag:     3,
			expected: true,
		},
		{
			name:     "Flags 255, flag 5",
			flags:    math.MaxUint8,
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags 255, flag 7",
			flags:    math.MaxUint8,
			flag:     7,
			expected: true,
		},
		{
			name:     "Flags 255, flag 8",
			flags:    math.MaxUint8,
			flag:     8,
			expected: false,
		},
	}
}

func providerSetFlagUint8() []providerTypeSetFlagUint8 {
	return []providerTypeSetFlagUint8{
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
			name:     "Flags 255, flag 2, set",
			flags:    math.MaxUint8,
			flag:     2,
			set:      true,
			expected: math.MaxUint8,
			err:      nil,
		},
		{
			name:     "Flags 255, flag 2, unset",
			flags:    math.MaxUint8,
			flag:     2,
			set:      false,
			expected: math.MaxUint8 - 4,
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
			name:     "Flags 5, flag 8, set",
			flags:    5,
			flag:     8,
			set:      true,
			expected: 5,
			err:      binflags.ErrorOutOfRange(binflags.ErrorMsgOutOfRange),
		},
	}
}
