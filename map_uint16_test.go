package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagMapUint16(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagMapUint16() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagMapUint16(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagMapUint16(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagMapUint16() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagMapUint16(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagMapUint16 struct {
	name     string
	flags    map[uint64]uint16
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapUint16 struct {
	name     string
	flags    map[uint64]uint16
	flag     uint64
	set      bool
	expected map[uint64]uint16
	err      error
}

func providerHasFlagMapUint16() []providerTypeHasFlagMapUint16 {
	return []providerTypeHasFlagMapUint16{
		{
			name:     "Flags {}, flag 0",
			flags:    map[uint64]uint16{},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags {}, flag 12425",
			flags:    map[uint64]uint16{},
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
			flags:    map[uint64]uint16{},
			flag:     2356241245436,
			expected: false,
		},
		{
			name:     "Flags {0: 1}, flag 0",
			flags:    map[uint64]uint16{0: 1},
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags {0: 32768}, flag 0",
			flags:    map[uint64]uint16{0: 32768},
			flag:     15,
			expected: true,
		},
		{
			name:     "Flags {0: 32768}, flag 31",
			flags:    map[uint64]uint16{1: 32768},
			flag:     31,
			expected: true,
		},
		{
			name:     "Flags {6235: 32}, flag 99765",
			flags:    map[uint64]uint16{6235: 32},
			flag:     99765,
			expected: true,
		},
		{
			name:     "Flags {6235: 32}, flag 99764",
			flags:    map[uint64]uint16{6235: 32},
			flag:     99764,
			expected: false,
		},
		{
			name:     "Flags {6235: 32}, flag 99766",
			flags:    map[uint64]uint16{6235: 32},
			flag:     99766,
			expected: false,
		},
	}
}

func providerSetFlagMapUint16() []providerTypeSetFlagMapUint16 {
	return []providerTypeSetFlagMapUint16{
		{
			name:     "Flags {}, flag 0, set",
			flags:    map[uint64]uint16{},
			flag:     0,
			set:      true,
			expected: map[uint64]uint16{0: 1},
		},
		{
			name:     "Flags {}, flag 7, set",
			flags:    map[uint64]uint16{},
			flag:     7,
			set:      true,
			expected: map[uint64]uint16{0: 128},
		},
		{
			name:     "Flags {}, flag 15, set",
			flags:    map[uint64]uint16{},
			flag:     15,
			set:      true,
			expected: map[uint64]uint16{0: 32768},
		},
		{
			name:     "Flags {0: 92}, flag 7, set",
			flags:    map[uint64]uint16{0: 92},
			flag:     7,
			set:      true,
			expected: map[uint64]uint16{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 7, unset",
			flags:    map[uint64]uint16{0: 220},
			flag:     7,
			set:      false,
			expected: map[uint64]uint16{0: 92},
		},
		{
			name:     "Flags {0: 220}, flag 7, set",
			flags:    map[uint64]uint16{0: 220},
			flag:     7,
			set:      true,
			expected: map[uint64]uint16{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 1, unset",
			flags:    map[uint64]uint16{0: 220},
			flag:     1,
			set:      false,
			expected: map[uint64]uint16{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 14, set",
			flags:    map[uint64]uint16{0: 220},
			flag:     14,
			set:      true,
			expected: map[uint64]uint16{0: 16604},
		},
		{
			name:     "Flags {0: 16604, 6: 23562}, flag 95, set",
			flags:    map[uint64]uint16{0: 16604, 6: 23562},
			flag:     95,
			set:      true,
			expected: map[uint64]uint16{0: 16604, 5: 32768, 6: 23562},
		},
		{
			name:     "Flags {0: 16604, 5: 32768, 6: 23562}, flag 94, set",
			flags:    map[uint64]uint16{0: 16604, 5: 32768, 6: 23562},
			flag:     94,
			set:      true,
			expected: map[uint64]uint16{0: 16604, 5: 49152, 6: 23562},
		},
		{
			name:     "Flags {0: 16604, 5: 32768, 6: 23562}, flag 95, unset",
			flags:    map[uint64]uint16{0: 16604, 5: 32768, 6: 23562},
			flag:     95,
			set:      false,
			expected: map[uint64]uint16{0: 16604, 6: 23562},
		},
		{
			name:     "Flags {0: 16604, 5: 32768, 6: 23562}, flag 97, unset",
			flags:    map[uint64]uint16{0: 16604, 5: 32768, 6: 23562},
			flag:     97,
			set:      false,
			expected: map[uint64]uint16{0: 16604, 5: 32768, 6: 23560},
		},
		{
			name:     "Flags {0: 16604, 5: 32768, 6: 23562}, flag 25365262353, set",
			flags:    map[uint64]uint16{0: 16604, 5: 32768, 6: 23562},
			flag:     25365262353,
			set:      true,
			expected: map[uint64]uint16{0: 16604, 5: 32768, 6: 23562, 1585328897: 2},
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
