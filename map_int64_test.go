package binflags_test

import (
	"reflect"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestHasFlagMapInt64(t *testing.T) {
	t.Parallel()

	for _, prov := range providerHasFlagMapInt64() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			actual := binflags.HasFlagMapInt64(prov.flags, prov.flag)

			if prov.expected != actual {
				t.Fatalf("Unexpected result.\nExpected: %t\nActual: %t", prov.expected, actual)
			}
		})
	}
}

func TestSetFlagMapInt64(t *testing.T) {
	t.Parallel()

	for _, prov := range providerSetFlagMapInt64() {
		prov := prov
		t.Run(prov.name, func(t *testing.T) {
			t.Parallel()

			err := binflags.SetFlagMapInt64(prov.flags, prov.flag, prov.set)

			if !reflect.DeepEqual(prov.err, err) {
				t.Fatalf("Unexpected error.\nExpected: %T(%v)\nActual: %T(%v)", prov.err, prov.err, err, err)
			}

			if !reflect.DeepEqual(prov.expected, prov.flags) {
				t.Fatalf("Unexpected flags.\nExpected: %+v\nActual: %+v", prov.expected, prov.flags)
			}
		})
	}
}

type providerTypeHasFlagMapInt64 struct {
	name     string
	flags    map[uint64]int64
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapInt64 struct {
	name     string
	flags    map[uint64]int64
	flag     uint64
	set      bool
	expected map[uint64]int64
	err      error
}

func providerHasFlagMapInt64() []providerTypeHasFlagMapInt64 {
	return []providerTypeHasFlagMapInt64{
		{
			name:     "Flags {}, flag 0",
			flags:    map[uint64]int64{},
			flag:     0,
			expected: false,
		},
		{
			name:     "Flags {}, flag 12425",
			flags:    map[uint64]int64{},
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
			flags:    map[uint64]int64{},
			flag:     2356241245436,
			expected: false,
		},
		{
			name:     "Flags {0: 1}, flag 0",
			flags:    map[uint64]int64{0: 1},
			flag:     0,
			expected: true,
		},
		{
			name:     "Flags {0: 32768}, flag 15",
			flags:    map[uint64]int64{0: 32768},
			flag:     15,
			expected: true,
		},
		{
			name:     "Flags {0: 2147483648}, flag 31",
			flags:    map[uint64]int64{0: 2147483648},
			flag:     31,
			expected: true,
		},
		{
			name:     "Flags {0: 4503601774854144}, flag 52",
			flags:    map[uint64]int64{0: 4503601774854144},
			flag:     52,
			expected: true,
		},
		{
			name:     "Flags {0: 4503599627370496}, flag 52",
			flags:    map[uint64]int64{0: 4503599627370496},
			flag:     52,
			expected: true,
		},
		{
			name:     "Flags {0: -9218868473709551616}, flag 63",
			flags:    map[uint64]int64{0: -9218868473709551616},
			flag:     63,
			expected: true,
		},
		{
			name:     "Flags {0: -9218868473709551616, 36644: 262144}, flag 2345234",
			flags:    map[uint64]int64{0: -9218868473709551616, 36644: 262144},
			flag:     2345234,
			expected: true,
		},
	}
}

func providerSetFlagMapInt64() []providerTypeSetFlagMapInt64 {
	return []providerTypeSetFlagMapInt64{
		{
			name:     "Flags {}, flag 0, set",
			flags:    map[uint64]int64{},
			flag:     0,
			set:      true,
			expected: map[uint64]int64{0: 1},
		},
		{
			name:     "Flags {}, flag 7, set",
			flags:    map[uint64]int64{},
			flag:     7,
			set:      true,
			expected: map[uint64]int64{0: 128},
		},
		{
			name:  "Flags {}, flag 15, set",
			flags: map[uint64]int64{},
			flag:  15,
			set:   true,
			expected: map[uint64]int64{
				0: 32768,
			},
		},
		{
			name:     "Flags {0: 92}, flag 7, set",
			flags:    map[uint64]int64{0: 92},
			flag:     7,
			set:      true,
			expected: map[uint64]int64{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 7, unset",
			flags:    map[uint64]int64{0: 220},
			flag:     7,
			set:      false,
			expected: map[uint64]int64{0: 92},
		},
		{
			name:     "Flags {0: 220}, flag 7, set",
			flags:    map[uint64]int64{0: 220},
			flag:     7,
			set:      true,
			expected: map[uint64]int64{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 1, unset",
			flags:    map[uint64]int64{0: 220},
			flag:     1,
			set:      false,
			expected: map[uint64]int64{0: 220},
		},
		{
			name:     "Flags {0: 220}, flag 14, set",
			flags:    map[uint64]int64{0: 220},
			flag:     14,
			set:      true,
			expected: map[uint64]int64{0: 16604},
		},
		{
			name:     "Flags {0: 16604}, flag 31, set",
			flags:    map[uint64]int64{0: 16604},
			flag:     31,
			set:      true,
			expected: map[uint64]int64{0: 2147500252},
		},
		{
			name:     "Flags {0: 2147500252}, flag 14, unset",
			flags:    map[uint64]int64{0: 2147500252},
			flag:     14,
			set:      false,
			expected: map[uint64]int64{0: 2147483868},
		},
		{
			name:     "Flags {0: 2147500252}, flag 60, set",
			flags:    map[uint64]int64{0: 2147500252},
			flag:     60,
			set:      true,
			expected: map[uint64]int64{0: 1152921506754347228},
		},
		{
			name:     "Flags {0: 2147500252}, flag 63, set",
			flags:    map[uint64]int64{0: 2147500252},
			flag:     63,
			set:      true,
			expected: map[uint64]int64{0: -9223372034707275556},
		},
		{
			name:     "Flags {0: -9223372034707275556}, flag 63, unset",
			flags:    map[uint64]int64{0: -9223372034707275556},
			flag:     63,
			set:      false,
			expected: map[uint64]int64{0: 2147500252},
		},
		{
			name:     "Flags {0: -9223372034707275556}, flag 1245, set",
			flags:    map[uint64]int64{0: -9223372034707275556},
			flag:     1245,
			set:      true,
			expected: map[uint64]int64{0: -9223372034707275556, 19: 536870912},
		},
		{
			name:     "Flags {0: -9223372034707275556, 19: 536870912}, flag 1242, set",
			flags:    map[uint64]int64{0: -9223372034707275556, 19: 536870912},
			flag:     1242,
			set:      true,
			expected: map[uint64]int64{0: -9223372034707275556, 19: 603979776},
		},
		{
			name:     "Flags {0:  -9223372034707275556, 19: 603979776}, flag 1242, unset",
			flags:    map[uint64]int64{0: -9223372034707275556, 19: 603979776},
			flag:     1242,
			set:      false,
			expected: map[uint64]int64{0: -9223372034707275556, 19: 536870912},
		},
		{
			name:     "Flags {0: -9223372034707275556, 19: 603979776}, flag 1242, set",
			flags:    map[uint64]int64{0: -9223372034707275556, 19: 603979776},
			flag:     1242,
			set:      true,
			expected: map[uint64]int64{0: -9223372034707275556, 19: 603979776},
		},
		{
			name:     "Flags {0: -9223372034707275556, 19: 536870912}, flag 1242, unset",
			flags:    map[uint64]int64{0: -9223372034707275556, 19: 536870912},
			flag:     1242,
			set:      false,
			expected: map[uint64]int64{0: -9223372034707275556, 19: 536870912},
		},
		{
			name:     "Flags {0: 436324, 1: 1, 2: 25623541}, flag 64, unset",
			flags:    map[uint64]int64{0: 436324, 1: 1, 2: 25623541},
			flag:     64,
			set:      false,
			expected: map[uint64]int64{0: 436324, 2: 25623541},
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
