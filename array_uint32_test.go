package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagArrayUint32(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagArrayUint32() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagArrayUint32(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagArrayUint32(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagArrayUint32() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagArrayUint32(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagArrayUint32 struct {
	name     string
	flags    []uint32
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayUint32 struct {
	name     string
	flags    []uint32
	flag     uint64
	set      bool
	expected []uint32
	err      error
}

func providerHasFlagArrayUint32() []providerTypeHasFlagArrayUint32 {
	return []providerTypeHasFlagArrayUint32{
		{
			name:     "Flags nil, flag 0",
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [115, 12, 105], flag 36",
			flags:    []uint32{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			name:     "Flags [115], flag 5",
			flags:    []uint32{115},
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 0",
			flags:    []uint32{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 8",
			flags:    []uint32{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			name:     "Flags [2147483713, 1, 0, 0], flag 31",
			flags:    []uint32{2147483713, 1, 0, 0},
			flag:     31,
			expected: true,
		},
		{
			name:     "Flags [0, 1, 0, 0], flag 32",
			flags:    []uint32{0, 1, 0, 0},
			flag:     32,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 1, 0], flag 64",
			flags:    []uint32{0, 0, 1, 0},
			flag:     64,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 9, 0], flag 67",
			flags:    []uint32{0, 0, 9, 0},
			flag:     67,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 71",
			flags:    []uint32{0, 0, 217, 0},
			flag:     71,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 70",
			flags:    []uint32{0, 0, 217, 0},
			flag:     70,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 69",
			flags:    []uint32{0, 0, 217, 0},
			flag:     69,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 72",
			flags:    []uint32{0, 0, 217, 0},
			flag:     72,
			expected: false,
		},
	}
}

func providerSetFlagArrayUint32() []providerTypeSetFlagArrayUint32 {
	return []providerTypeSetFlagArrayUint32{
		{
			name:     "Flags [0], flag 0, set",
			flags:    []uint32{0},
			flag:     0,
			set:      true,
			expected: []uint32{1},
			err:      nil,
		},
		{
			name:     "Flags [1], flag 6, set",
			flags:    []uint32{1},
			flag:     6,
			set:      true,
			expected: []uint32{65},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 6, set",
			flags:    []uint32{65},
			flag:     6,
			set:      false,
			expected: []uint32{1},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 12, set",
			flags:    []uint32{65},
			flag:     12,
			set:      true,
			expected: []uint32{4161},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 31, set",
			flags:    []uint32{65},
			flag:     31,
			set:      true,
			expected: []uint32{2147483713},
			err:      nil,
		},
		{
			name:     "Flags [65, 0], flag 32, set",
			flags:    []uint32{65, 0},
			flag:     32,
			set:      true,
			expected: []uint32{65, 1},
			err:      nil,
		},
		{
			name:     "Flags [65, 0], flag 63, set",
			flags:    []uint32{65, 0},
			flag:     63,
			set:      true,
			expected: []uint32{65, 2147483648},
			err:      nil,
		},
		{
			name:     "Flags [65, 0, 1235, 724, 635], flag 68, set",
			flags:    []uint32{65, 0, 1235, 724, 635},
			flag:     68,
			set:      true,
			expected: []uint32{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			name:     "Flags [65, 0, 1235, 724, 635], flag 68, unset",
			flags:    []uint32{65, 0, 1235, 724, 635},
			flag:     68,
			set:      false,
			expected: []uint32{65, 0, 1219, 724, 635},
			err:      nil,
		},
		{
			name:     "Flags [0], flag 32, set",
			flags:    []uint32{0},
			flag:     32,
			set:      true,
			expected: []uint32{0},
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
