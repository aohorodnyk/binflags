package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagMapUint8(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagMapUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagMapUint8(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestHasFlagMapByte(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagMapUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagMapByte(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagMapUint8(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagMapUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagMapUint8(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

func TestSetFlagMapByte(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagMapUint8() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagMapByte(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagMapUint8 struct {
	name     string
	flags    map[uint64]uint8
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapUint8 struct {
	name     string
	flags    map[uint64]uint8
	flag     uint64
	set      bool
	expected map[uint64]uint8
	err      error
}

func providerHasFlagMapUint8() []providerTypeHasFlagMapUint8 {
	return []providerTypeHasFlagMapUint8{
		{
			name:     "Flags {}, flag 0",
			flags:    map[uint64]uint8{},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags {}, flag 12425",
			flags:    map[uint64]uint8{},
			flag:     12425,
			expected: false,
		},
		{
			name:     "Flags nil, flag 0",
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags nil, flag 12425",
			flags:    nil,
			flag:     12425,
			expected: false,
		},
		{
			name:     "Flags {}, flag 2356241245436",
			flags:    map[uint64]uint8{},
			flag:     2356241245436,
			expected: false,
		},
		{
			name:     "Flags {0: 1}, flag 0",
			flags:    map[uint64]uint8{0: 1},
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags {234125: 128}, flag 1873007",
			flags:    map[uint64]uint8{234125: 128},
			flag:     1873007,
			expected: true,
		},
		{
			name:     "Flags {234125: 128}, flag 1873006",
			flags:    map[uint64]uint8{234125: 128},
			flag:     1873006,
			expected: false,
		},
		{
			name:     "Flags {234125: 128}, flag 1873008",
			flags:    map[uint64]uint8{234125: 128},
			flag:     1873008,
			expected: false,
		},
		{
			name:     "Flags {0: 128, 234125: 128}, flag 7",
			flags:    map[uint64]uint8{0: 128, 234125: 128},
			flag:     7,
			expected: true,
		},
		{
			name:     "Flags {0: 255, 5: 255, 63: 255, 234125: 128}, flag 7",
			flags:    map[uint64]uint8{0: 255, 5: 255, 63: 255, 234125: 128},
			flag:     7,
			expected: true,
		},
		{
			name:     "Flags {0: 255, 5: 64, 63: 128, 234125: 128}, flag 7",
			flags:    map[uint64]uint8{0: 255, 5: 64, 63: 128, 234125: 128},
			flag:     7,
			expected: true,
		},
		{
			name:     "Flags {0: 255, 5: 64, 63: 128, 234125: 128}, flag 46",
			flags:    map[uint64]uint8{0: 255, 5: 64, 63: 128, 234125: 128},
			flag:     46,
			expected: true,
		},
		{
			name:     "Flags {0: 255, 5: 64, 63: 128, 234125: 128}, flag 511",
			flags:    map[uint64]uint8{0: 255, 5: 64, 63: 128, 234125: 128},
			flag:     511,
			expected: true,
		},
	}
}

func providerSetFlagMapUint8() []providerTypeSetFlagMapUint8 {
	return []providerTypeSetFlagMapUint8{
		{
			name:     "Flags {}, flag 0, set",
			flags:    map[uint64]uint8{},
			flag:     0,
			set:      true,
			expected: map[uint64]uint8{0: 1},
		},
		{
			name:     "Flags {}, flag 7, set",
			flags:    map[uint64]uint8{},
			flag:     7,
			set:      true,
			expected: map[uint64]uint8{0: 128},
		},
		{
			name:     "Flags {}, flag 2346, set",
			flags:    map[uint64]uint8{},
			flag:     2346,
			set:      true,
			expected: map[uint64]uint8{293: 4},
		},
		{
			name:     "Flags {0: 92}, flag 7, set",
			flags:    map[uint64]uint8{0: 92},
			flag:     7,
			set:      true,
			expected: map[uint64]uint8{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 7, unset",
			flags:    map[uint64]uint8{0: 220},
			flag:     7,
			set:      false,
			expected: map[uint64]uint8{0: 92},
		},
		{
			name:     "Flags {0: 220, 1: 56, 6: 32, 6235: 123}, flag 48, set",
			flags:    map[uint64]uint8{0: 220, 1: 56, 6: 32, 6235: 123},
			flag:     48,
			set:      true,
			expected: map[uint64]uint8{0: 220, 1: 56, 6: 33, 6235: 123},
		},
		{
			name:     "Flags {0: 220, 1: 56, 6: 32, 6235: 123}, flag 53, unset",
			flags:    map[uint64]uint8{0: 220, 1: 56, 6: 32, 6235: 123},
			flag:     53,
			set:      false,
			expected: map[uint64]uint8{0: 220, 1: 56, 6235: 123},
		},
		{
			name:     "Flags {0: 220, 1: 56, 6235: 123}, flag 53, set",
			flags:    map[uint64]uint8{0: 220, 1: 56, 6235: 123},
			flag:     53,
			set:      true,
			expected: map[uint64]uint8{0: 220, 1: 56, 6: 32, 6235: 123},
		},
		{
			name:     "Flags {0: 220, 1: 56, 6235: 123}, flag 53, unset",
			flags:    map[uint64]uint8{0: 220, 1: 56, 6235: 123},
			flag:     53,
			set:      false,
			expected: map[uint64]uint8{0: 220, 1: 56, 6235: 123},
		},
		{
			name:     "Flags nil, flag 0, unset",
			flags:    nil,
			flag:     0,
			set:      false,
			expected: nil,
			err:      binflags.ErrorFlagsMapNil(binflags.ErrorMsgFlagsMapNil),
		},
		{
			name:     "Flags nil, flag 0, set",
			flags:    nil,
			flag:     0,
			set:      true,
			expected: nil,
			err:      binflags.ErrorFlagsMapNil(binflags.ErrorMsgFlagsMapNil),
		},
	}
}
