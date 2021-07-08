package binflags_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagMapInt16(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagMapInt16() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagMapInt16(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagMapInt16(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagMapInt16() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagMapInt16(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagMapInt16 struct {
	name     string
	flags    map[uint64]int16
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapInt16 struct {
	name     string
	flags    map[uint64]int16
	flag     uint64
	set      bool
	expected map[uint64]int16
	err      error
}

func providerHasFlagMapInt16() []providerTypeHasFlagMapInt16 {
	return []providerTypeHasFlagMapInt16{
		{
			name:     "Flags {}, flag 0",
			flags:    map[uint64]int16{},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags {}, flag 12425",
			flags:    map[uint64]int16{},
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
			flags:    map[uint64]int16{},
			flag:     2356241245436,
			expected: false,
		},
		{
			name:     "Flags {0: 1}, flag 0",
			flags:    map[uint64]int16{0: 1},
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags {0: -32768}, flag 15",
			flags:    map[uint64]int16{0: math.MinInt16},
			flag:     15,
			expected: true,
		},
		{
			name:     "Flags {1: -32768}, flag 31",
			flags:    map[uint64]int16{1: math.MinInt16},
			flag:     31,
			expected: true,
		},
		{
			name:     "Flags {6235: 32}, flag 99765",
			flags:    map[uint64]int16{6235: 32},
			flag:     99765,
			expected: true,
		},
		{
			name:     "Flags {6235: 32}, flag 99764",
			flags:    map[uint64]int16{6235: 32},
			flag:     99764,
			expected: false,
		},
		{
			name:     "Flags {6235: 32}, flag 99766",
			flags:    map[uint64]int16{6235: 32},
			flag:     99766,
			expected: false,
		},
	}
}

func providerSetFlagMapInt16() []providerTypeSetFlagMapInt16 {
	return []providerTypeSetFlagMapInt16{
		{
			name:     "Flags {}, flag 0, set",
			flags:    map[uint64]int16{},
			flag:     0,
			set:      true,
			expected: map[uint64]int16{0: 1},
		},
		{
			name:     "Flags {}, flag 7, set",
			flags:    map[uint64]int16{},
			flag:     7,
			set:      true,
			expected: map[uint64]int16{0: 128},
		},
		{
			name:     "Flags {}, flag 15, set",
			flags:    map[uint64]int16{},
			flag:     15,
			set:      true,
			expected: map[uint64]int16{0: -32768},
		},
		{
			name:     "Flags {0: 92}, flag 7, set",
			flags:    map[uint64]int16{0: 92},
			flag:     7,
			set:      true,
			expected: map[uint64]int16{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 7, unset",
			flags:    map[uint64]int16{0: 220},
			flag:     7,
			set:      false,
			expected: map[uint64]int16{0: 92},
		},
		{
			name:     "Flags {0: 220}, flag 7, set",
			flags:    map[uint64]int16{0: 220},
			flag:     7,
			set:      true,
			expected: map[uint64]int16{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 1, unset",
			flags:    map[uint64]int16{0: 220},
			flag:     1,
			set:      false,
			expected: map[uint64]int16{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 14, set",
			flags:    map[uint64]int16{0: 220},
			flag:     14,
			set:      true,
			expected: map[uint64]int16{0: 16604},
		},
		{
			name:     "Flags {0: 16604, 6: 23562}, flag 95, set",
			flags:    map[uint64]int16{0: 16604, 6: 23562},
			flag:     95,
			set:      true,
			expected: map[uint64]int16{0: 16604, 5: -32768, 6: 23562},
		},
		{
			name:     "Flags {0: 16604, 5: -32768, 6: 23562}, flag 92, set",
			flags:    map[uint64]int16{0: 16604, 5: -32768, 6: 23562},
			flag:     92,
			set:      true,
			expected: map[uint64]int16{0: 16604, 5: -28672, 6: 23562},
		},
		{
			name:     "Flags {0: 16604, 5: -32768, 6: 23562}, flag 95, unset",
			flags:    map[uint64]int16{0: 16604, 5: -32768, 6: 23562},
			flag:     95,
			set:      false,
			expected: map[uint64]int16{0: 16604, 6: 23562},
		},
		{
			name:     "Flags {0: 16604, 5: -32768, 6: 23562}, flag 97, unset",
			flags:    map[uint64]int16{0: 16604, 5: -32768, 6: 23562},
			flag:     97,
			set:      false,
			expected: map[uint64]int16{0: 16604, 5: -32768, 6: 23560},
		},
		{
			name:     "Flags {0: 16604, 5: -32768, 6: 23562}, flag 97, set",
			flags:    map[uint64]int16{0: 16604, 5: -32768, 6: 23562},
			flag:     25365262353,
			set:      true,
			expected: map[uint64]int16{0: 16604, 5: -32768, 6: 23562, 1585328897: 2},
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
