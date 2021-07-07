package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagArrayUint8(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagArrayUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagArrayUint8(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestHasFlagArrayByte(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagArrayUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagArrayByte(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagArrayUint8(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagArrayUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagArrayUint8(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

func TestSetFlagArrayByte(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagArrayUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagArrayByte(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagArrayUint8 struct {
	name     string
	flags    []uint8
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayUint8 struct {
	name     string
	flags    []uint8
	flag     uint64
	set      bool
	expected []uint8
	err      error
}

func providerHasFlagArrayUint8() []providerTypeHasFlagArrayUint8 {
	return []providerTypeHasFlagArrayUint8{
		{
			name:     "Flags nil, flag 0",
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [115, 12, 105], flag 36",
			flags:    []uint8{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			name:     "Flags [115], flag 5",
			flags:    []uint8{115},
			flag:     5,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 0",
			flags:    []uint8{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 0, 0], flag 8",
			flags:    []uint8{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			name:     "Flags [0, 0, 1, 0], flag 16",
			flags:    []uint8{0, 0, 1, 0},
			flag:     16,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 9, 0], flag 19",
			flags:    []uint8{0, 0, 9, 0},
			flag:     19,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 127, 0], flag 23",
			flags:    []uint8{0, 0, 217, 0},
			flag:     23,
			expected: true,
		},
		{
			name:     "Flags [0, 0, 217, 0], flag 24",
			flags:    []uint8{0, 0, 217, 0},
			flag:     24,
			expected: false,
		},
	}
}

func providerSetFlagArrayUint8() []providerTypeSetFlagArrayUint8 {
	return []providerTypeSetFlagArrayUint8{
		{
			name:     "Flags [0], flag 0, set",
			flags:    []uint8{0},
			flag:     0,
			set:      true,
			expected: []uint8{1},
			err:      nil,
		},
		{
			name:     "Flags [1], flag 6, set",
			flags:    []uint8{1},
			flag:     6,
			set:      true,
			expected: []uint8{65},
			err:      nil,
		},
		{
			name:     "Flags [65], flag 6, unset",
			flags:    []uint8{65},
			flag:     6,
			set:      false,
			expected: []uint8{1},
			err:      nil,
		},
		{
			name:     "Flags [65, 1, 0, 83], flag 30, set",
			flags:    []uint8{65, 1, 0, 83},
			flag:     30,
			set:      true,
			expected: []uint8{65, 1, 0, 83},
			err:      nil,
		},
		{
			name:     "Flags [65, 1, 0, 83], flag 30, unset",
			flags:    []uint8{65, 1, 0, 83},
			flag:     30,
			set:      false,
			expected: []uint8{65, 1, 0, 19},
			err:      nil,
		},
		{
			name:     "Flags [0, 0], flag 7, set",
			flags:    []uint8{0, 0},
			flag:     7,
			set:      true,
			expected: []uint8{128, 0},
			err:      nil,
		},
		{
			name:     "Flags [0], flag 8, set",
			flags:    []uint8{0},
			flag:     8,
			set:      true,
			expected: []uint8{0},
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
