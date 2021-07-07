package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagArrayInt16(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagArrayInt16() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagArrayInt16(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagArrayInt16(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagArrayInt16() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagArrayInt16(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagArrayInt16 struct {
	name     string
	flags    []int16
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayInt16 struct {
	name     string
	flags    []int16
	flag     uint64
	set      bool
	expected []int16
	err      error
}

func providerHasFlagArrayInt16() []providerTypeHasFlagArrayInt16 {
	return []providerTypeHasFlagArrayInt16{
		{
			name:     "Flags nil, flag 0",
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [115, 12, 105], flag 36",
			flags:    []int16{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			name:     "Flags [115], flag 5",
			flags:    []int16{115},
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 0",
			flags:    []int16{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 8",
			flags:    []int16{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 1, 0], flag 32",
			flags:    []int16{0, 0, 1, 0},
			flag:     32,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 9, 0], flag 35",
			flags:    []int16{0, 0, 9, 0},
			flag:     35,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 39",
			flags:    []int16{0, 0, 217, 0},
			flag:     39,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 38",
			flags:    []int16{0, 0, 217, 0},
			flag:     38,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 37",
			flags:    []int16{0, 0, 217, 0},
			flag:     37,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 40",
			flags:    []int16{0, 0, 217, 0},
			flag:     40,
			expected: false,
		},
	}
}

func providerSetFlagArrayInt16() []providerTypeSetFlagArrayInt16 {
	return []providerTypeSetFlagArrayInt16{
		{
			name:     "Flags [0], flag 0, set",
			flags:    []int16{0},
			flag:     0,
			set:      true,
			expected: []int16{1},
			err:      nil,
		},
		{
			name:     "Flags [1], flag 6, set",
			flags:    []int16{1},
			flag:     6,
			set:      true,
			expected: []int16{65},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 6, unset",
			flags:    []int16{65},
			flag:     6,
			set:      false,
			expected: []int16{1},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 12, set",
			flags:    []int16{65},
			flag:     12,
			set:      true,
			expected: []int16{4161},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 15, set",
			flags:    []int16{65},
			flag:     15,
			set:      true,
			expected: []int16{-32703},
			err:      nil,
		},
		{
			name:     "Flags [65, 0], flag 16, set",
			flags:    []int16{65, 0},
			flag:     16,
			set:      true,
			expected: []int16{65, 1},
			err:      nil,
		},
		{
			name:     "Flags [65, 0], flag 31, set",
			flags:    []int16{65, 0},
			flag:     31,
			set:      true,
			expected: []int16{65, -32768},
			err:      nil,
		},
		{
			name:     "Flags [65, 0, 1235, 724, 635], flag 50, set",
			flags:    []int16{65, 0, 1235, 724, 635},
			flag:     50,
			set:      true,
			expected: []int16{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			name:     "Flags [65, 0, 1235, 724, 635], flag 50, unset",
			flags:    []int16{65, 0, 1235, 724, 635},
			flag:     50,
			set:      false,
			expected: []int16{65, 0, 1235, 720, 635},
			err:      nil,
		},
		{
			name:     "Flags [0], flag 16, set",
			flags:    []int16{0},
			flag:     16,
			set:      true,
			expected: []int16{0},
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
