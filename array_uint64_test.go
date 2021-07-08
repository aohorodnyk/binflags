package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagArrayUint64(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagArrayUint64() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagArrayUint64(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagArrayUint64(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagArrayUint64() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagArrayUint64(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagArrayUint64 struct {
	name     string
	flags    []uint64
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayUint64 struct {
	name     string
	flags    []uint64
	flag     uint64
	set      bool
	expected []uint64
	err      error
}

func providerHasFlagArrayUint64() []providerTypeHasFlagArrayUint64 {
	return []providerTypeHasFlagArrayUint64{
		{
			name:     "Flags nil, flag 0",
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [115, 12, 105], flag 36",
			flags:    []uint64{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			name:     "Flags [115], flag 5",
			flags:    []uint64{115},
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 0",
			flags:    []uint64{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 8",
			flags:    []uint64{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			name:     "Flags [9223372036854775808, 1, 0, 0], flag 63",
			flags:    []uint64{9223372036854775808, 1, 0, 0},
			flag:     63,
			expected: true,
		},
		{
			name:     "Flags [0, 1, 0, 0], flag 64",
			flags:    []uint64{0, 1, 0, 0},
			flag:     64,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 1, 0], flag 128",
			flags:    []uint64{0, 0, 1, 0},
			flag:     128,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 9, 0], flag 131",
			flags:    []uint64{0, 0, 9, 0},
			flag:     131,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 135",
			flags:    []uint64{0, 0, 217, 0},
			flag:     135,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 134",
			flags:    []uint64{0, 0, 217, 0},
			flag:     134,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 69",
			flags:    []uint64{0, 0, 217, 0},
			flag:     69,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 72",
			flags:    []uint64{0, 0, 217, 0},
			flag:     72,
			expected: false,
		},
	}
}

func providerSetFlagArrayUint64() []providerTypeSetFlagArrayUint64 {
	return []providerTypeSetFlagArrayUint64{
		{
			name:     "Flags [0], flag 0, set",
			flags:    []uint64{0},
			flag:     0,
			set:      true,
			expected: []uint64{1},
			err:      nil,
		},
		{
			name:     "Flags [1], flag 6, set",
			flags:    []uint64{1},
			flag:     6,
			set:      true,
			expected: []uint64{65},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 6, unset",
			flags:    []uint64{65},
			flag:     6,
			set:      false,
			expected: []uint64{1},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 12, set",
			flags:    []uint64{65},
			flag:     12,
			set:      true,
			expected: []uint64{4161},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 63, set",
			flags:    []uint64{65},
			flag:     63,
			set:      true,
			expected: []uint64{9223372036854775873},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 64, set",
			flags:    []uint64{65, 0},
			flag:     64,
			set:      true,
			expected: []uint64{65, 1},
			err:      nil,
		},
		{
			name:     "Flags [65, 0], flag 217, set",
			flags:    []uint64{65, 0},
			flag:     127,
			set:      true,
			expected: []uint64{65, 9223372036854775808},
			err:      nil,
		},
		{
			name:     "Flags [65, 0, 1235, 724, 635], flag 132, set",
			flags:    []uint64{65, 0, 1235, 724, 635},
			flag:     132,
			set:      true,
			expected: []uint64{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			name:     "Flags [65, 0, 1235, 724, 635], flag 132, unset",
			flags:    []uint64{65, 0, 1235, 724, 635},
			flag:     132,
			set:      false,
			expected: []uint64{65, 0, 1219, 724, 635},
			err:      nil,
		},
		{
			name:     "Flags [0], flag 64, set",
			flags:    []uint64{0},
			flag:     64,
			set:      true,
			expected: []uint64{0},
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
