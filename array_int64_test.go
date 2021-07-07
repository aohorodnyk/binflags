package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagArrayInt64(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagArrayInt64() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagArrayInt64(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagArrayInt64(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagArrayInt64() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagArrayInt64(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagArrayInt64 struct {
	name     string
	flags    []int64
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayInt64 struct {
	name     string
	flags    []int64
	flag     uint64
	set      bool
	expected []int64
	err      error
}

func providerHasFlagArrayInt64() []providerTypeHasFlagArrayInt64 {
	return []providerTypeHasFlagArrayInt64{
		{
			name:     "Flags nil, flag 0",
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [115, 12, 105], flag 36",
			flags:    []int64{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			name:     "Flags [115], flag 5",
			flags:    []int64{115},
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 0",
			flags:    []int64{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 8",
			flags:    []int64{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			name:     "Flags {-9223372036854775808, 1, 0, 0}, flag 63",
			flags:    []int64{-9223372036854775808, 1, 0, 0},
			flag:     63,
			expected: true,
		},
		{
			name:     "Flags [0, 1, 0, 0], flag 64",
			flags:    []int64{0, 1, 0, 0},
			flag:     64,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 1, 0], flag 128",
			flags:    []int64{0, 0, 1, 0},
			flag:     128,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 9, 0], flag 131",
			flags:    []int64{0, 0, 9, 0},
			flag:     131,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 135",
			flags:    []int64{0, 0, 217, 0},
			flag:     135,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 134",
			flags:    []int64{0, 0, 217, 0},
			flag:     134,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 69",
			flags:    []int64{0, 0, 217, 0},
			flag:     69,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 72",
			flags:    []int64{0, 0, 217, 0},
			flag:     72,
			expected: false,
		},
	}
}

func providerSetFlagArrayInt64() []providerTypeSetFlagArrayInt64 {
	return []providerTypeSetFlagArrayInt64{
		{
			name:     "Flags [0], flag 0, set",
			flags:    []int64{0},
			flag:     0,
			set:      true,
			expected: []int64{1},
			err:      nil,
		},
		{
			name:     "Flags [1], flag 6, set",
			flags:    []int64{1},
			flag:     6,
			set:      true,
			expected: []int64{65},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 6, unset",
			flags:    []int64{65},
			flag:     6,
			set:      false,
			expected: []int64{1},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 12, set",
			flags:    []int64{65},
			flag:     12,
			set:      true,
			expected: []int64{4161},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 63, set",
			flags:    []int64{65},
			flag:     63,
			set:      true,
			expected: []int64{-9223372036854775743},
			err:      nil,
		},
		{
			name:     "Flags [65, 0], flag 64, set",
			flags:    []int64{65, 0},
			flag:     64,
			set:      true,
			expected: []int64{65, 1},
			err:      nil,
		},
		{
			name:     "Flags [65, 0], flag 127, set",
			flags:    []int64{65, 0},
			flag:     127,
			set:      true,
			expected: []int64{65, -9223372036854775808},
			err:      nil,
		},
		{
			name:     "Flags [65, 0, 1235, 724, 635], flag 132, set",
			flags:    []int64{65, 0, 1235, 724, 635},
			flag:     132,
			set:      true,
			expected: []int64{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			name:     "Flags [65, 0, 1235, 724, 635], flag 132, unset",
			flags:    []int64{65, 0, 1235, 724, 635},
			flag:     132,
			set:      false,
			expected: []int64{65, 0, 1219, 724, 635},
			err:      nil,
		},
		{
			name:     "Flags [0], flag 64, set",
			flags:    []int64{0},
			flag:     64,
			set:      true,
			expected: []int64{0},
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
