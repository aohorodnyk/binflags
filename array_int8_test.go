package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagArrayInt8(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagArrayInt8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagArrayInt8(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagArrayInt8(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagArrayInt8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagArrayInt8(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagArrayInt8 struct {
	name     string
	flags    []int8
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayInt8 struct {
	name     string
	flags    []int8
	flag     uint64
	set      bool
	expected []int8
	err      error
}

func providerHasFlagArrayInt8() []providerTypeHasFlagArrayInt8 {
	return []providerTypeHasFlagArrayInt8{
		{
			name:     "Flags nil, flag 0",
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [115, 12, 105], flag 36",
			flags:    []int8{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			name:     "Flags [115], flag 5",
			flags:    []int8{115},
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 0",
			flags:    []int8{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 8",
			flags:    []int8{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 1, 0], flag 16",
			flags:    []int8{0, 0, 1, 0},
			flag:     16,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 9, 0], flag 19",
			flags:    []int8{0, 0, 9, 0},
			flag:     19,
			expected: true,
		},
		{
			name:     "Flags {0, 0, -39, 0}, flag 19",
			flags:    []int8{0, 0, -39, 0},
			flag:     23,
			expected: true,
		},
		{
			name:     "Flags {0, 0, -39, 0}, flag 24",
			flags:    []int8{0, 0, -39, 0},
			flag:     24,
			expected: false,
		},
	}
}

func providerSetFlagArrayInt8() []providerTypeSetFlagArrayInt8 {
	return []providerTypeSetFlagArrayInt8{
		{
			name:     "Flags [0], flag 0, set",
			flags:    []int8{0},
			flag:     0,
			set:      true,
			expected: []int8{1},
			err:      nil,
		},
		{
			name:     "Flags [1], flag 6, set",
			flags:    []int8{1},
			flag:     6,
			set:      true,
			expected: []int8{65},
			err:      nil,
		},
		{
			name:     "Flags [65, 1, 0, 83], flag 30, set",
			flags:    []int8{65, 1, 0, 83},
			flag:     30,
			set:      true,
			expected: []int8{65, 1, 0, 83},
			err:      nil,
		},
		{
			name:     "Flags [65, 1, 0, 83], flag 30, unset",
			flags:    []int8{65, 1, 0, 83},
			flag:     30,
			set:      false,
			expected: []int8{65, 1, 0, 19},
			err:      nil,
		},
		{
			name:     "Flags [0, 0], flag 7, set",
			flags:    []int8{0, 0},
			flag:     7,
			set:      true,
			expected: []int8{-128, 0},
			err:      nil,
		},
		{
			name:     "Flags [0], flag 8, set",
			flags:    []int8{0},
			flag:     8,
			set:      true,
			expected: []int8{0},
			err:      binflags.ErrorOutOfRange(binflags.ErrorMsgOutOfRange),
		},
		{
			name:     "Flags nil, flag 0, unset",
			flags:    nil,
			flag:     0,
			set:      false,
			expected: nil,
			err:      binflags.ErrorFlagsArrayNil(binflags.ErrorMsgFlagsArrayNil),
		},
	}
}
