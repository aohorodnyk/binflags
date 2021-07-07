package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagArrayInt32(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagArrayInt32() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagArrayInt32(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagArrayInt32(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagArrayInt32() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagArrayInt32(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagArrayInt32 struct {
	name     string
	flags    []int32
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayInt32 struct {
	name     string
	flags    []int32
	flag     uint64
	set      bool
	expected []int32
	err      error
}

func providerHasFlagArrayInt32() []providerTypeHasFlagArrayInt32 {
	return []providerTypeHasFlagArrayInt32{
		{
			name:     "Flags nil, flag 0",
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [115, 12, 105], flag 36",
			flags:    []int32{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			name:     "Flags [115], flag 5",
			flags:    []int32{115},
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags [0,0,0,0], flag 0",
			flags:    []int32{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [0,0,0,0], flag 8",
			flags:    []int32{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			name:     "Flags {-2147483583, 1, 0, 0}, flag 31",
			flags:    []int32{-2147483583, 1, 0, 0},
			flag:     31,
			expected: true,
		},
		{
			name:     "Flags [0, 1, 0, 0], flag 32",
			flags:    []int32{0, 1, 0, 0},
			flag:     32,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 1, 0], flag 64",
			flags:    []int32{0, 0, 1, 0},
			flag:     64,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 9, 0], flag 67",
			flags:    []int32{0, 0, 9, 0},
			flag:     67,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 71",
			flags:    []int32{0, 0, 217, 0},
			flag:     71,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 70",
			flags:    []int32{0, 0, 217, 0},
			flag:     70,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 69",
			flags:    []int32{0, 0, 217, 0},
			flag:     69,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 72",
			flags:    []int32{0, 0, 217, 0},
			flag:     72,
			expected: false,
		},
	}
}

func providerSetFlagArrayInt32() []providerTypeSetFlagArrayInt32 {
	return []providerTypeSetFlagArrayInt32{
		{
			name:     "Flags [0], flag 0, set",
			flags:    []int32{0},
			flag:     0,
			set:      true,
			expected: []int32{1},
			err:      nil,
		},
		{
			name:     "Flags [1], flag 6, set",
			flags:    []int32{1},
			flag:     6,
			set:      true,
			expected: []int32{65},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 6, unset",
			flags:    []int32{65},
			flag:     6,
			set:      false,
			expected: []int32{1},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 12, set",
			flags:    []int32{65},
			flag:     12,
			set:      true,
			expected: []int32{4161},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 31, set",
			flags:    []int32{65},
			flag:     31,
			set:      true,
			expected: []int32{-2147483583},
			err:      nil,
		},
		{
			name:     "Flags [65, 0], flag 32, set",
			flags:    []int32{65, 0},
			flag:     32,
			set:      true,
			expected: []int32{65, 1},
			err:      nil,
		},
		{
			name:     "Flags [65, 0], flag 63, set",
			flags:    []int32{65, 0},
			flag:     63,
			set:      true,
			expected: []int32{65, -2147483648},
			err:      nil,
		},
		{
			name:     "Flags [65, 0, 1235, 724, 635], flag 68, set",
			flags:    []int32{65, 0, 1235, 724, 635},
			flag:     68,
			set:      true,
			expected: []int32{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			name:     "Flags [65, 0, 1235, 724, 635], flag 68, unset",
			flags:    []int32{65, 0, 1235, 724, 635},
			flag:     68,
			set:      false,
			expected: []int32{65, 0, 1219, 724, 635},
			err:      nil,
		},
		{
			name:     "Flags [0], flag 32, set",
			flags:    []int32{0},
			flag:     32,
			set:      true,
			expected: []int32{0},
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
