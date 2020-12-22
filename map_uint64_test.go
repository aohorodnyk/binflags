package binflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/binflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagMapUint64(t *testing.T) {
	for idx, prov := range providerHasFlagMapUint64() {
		t.Run(fmt.Sprintf("TestHasFlagMapUint64_%d", idx), func(t *testing.T) {
			actual := binflags.HasFlagMapUint64(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagMapUint64(t *testing.T) {
	for idx, prov := range providerSetFlagMapUint64() {
		t.Run(fmt.Sprintf("TestHasFlagMapUint64_%d", idx), func(t *testing.T) {
			err := binflags.SetFlagMapUint64(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagMapUint64 struct {
	flags    map[uint64]uint64
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapUint64 struct {
	flags    map[uint64]uint64
	flag     uint64
	set      bool
	expected map[uint64]uint64
	err      error
}

func providerHasFlagMapUint64() []providerTypeHasFlagMapUint64 {
	return []providerTypeHasFlagMapUint64{
		{
			flags:    map[uint64]uint64{},
			flag:     0,
			expected: false,
		},
		{
			flags:    map[uint64]uint64{},
			flag:     12425,
			expected: false,
		},
		{
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			flags:    nil,
			flag:     12425,
			expected: false,
		},
		{
			flags:    map[uint64]uint64{},
			flag:     2356241245436,
			expected: false,
		},
		{
			flags: map[uint64]uint64{
				0: 1,
			},
			flag:     0,
			expected: true,
		},
		{
			flags: map[uint64]uint64{
				0: 32768,
			},
			flag:     15,
			expected: true,
		},
		{
			flags: map[uint64]uint64{
				0: 2147483648,
			},
			flag:     31,
			expected: true,
		},
		{
			flags: map[uint64]uint64{
				0: 4503601774854144,
			},
			flag:     52,
			expected: true,
		},
		{
			flags: map[uint64]uint64{
				0: 4503599627370496,
			},
			flag:     52,
			expected: true,
		},
		{
			flags: map[uint64]uint64{
				0: 9227875600000000000,
			},
			flag:     63,
			expected: true,
		},
		{
			flags: map[uint64]uint64{
				0:     9227875600000000000,
				36644: 262144,
			},
			flag:     2345234,
			expected: true,
		},
	}
}

func providerSetFlagMapUint64() []providerTypeSetFlagMapUint64 {
	return []providerTypeSetFlagMapUint64{
		{
			flags: map[uint64]uint64{},
			flag:  0,
			set:   true,
			expected: map[uint64]uint64{
				0: 1,
			},
		},
		{
			flags: map[uint64]uint64{},
			flag:  7,
			set:   true,
			expected: map[uint64]uint64{
				0: 128,
			},
		},
		{
			flags: map[uint64]uint64{},
			flag:  15,
			set:   true,
			expected: map[uint64]uint64{
				0: 32768,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 92,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]uint64{
				0: 220,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 220,
			},
			flag: 7,
			set:  false,
			expected: map[uint64]uint64{
				0: 92,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 220,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]uint64{
				0: 220,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 220,
			},
			flag: 1,
			set:  false,
			expected: map[uint64]uint64{
				0: 220,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 220,
			},
			flag: 14,
			set:  true,
			expected: map[uint64]uint64{
				0: 16604,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 16604,
			},
			flag: 31,
			set:  true,
			expected: map[uint64]uint64{
				0: 2147500252,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 2147500252,
			},
			flag: 14,
			set:  false,
			expected: map[uint64]uint64{
				0: 2147483868,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 2147500252,
			},
			flag: 60,
			set:  true,
			expected: map[uint64]uint64{
				0: 1152921506754347228,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 2147500252,
			},
			flag: 63,
			set:  true,
			expected: map[uint64]uint64{
				0: 9223372039002276060,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 9223372039002276060,
			},
			flag: 63,
			set:  false,
			expected: map[uint64]uint64{
				0: 2147500252,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 9223372039002276060,
			},
			flag: 1245,
			set:  true,
			expected: map[uint64]uint64{
				0:  9223372039002276060,
				19: 536870912,
			},
		},
		{
			flags: map[uint64]uint64{
				0:  9223372039002276060,
				19: 536870912,
			},
			flag: 1242,
			set:  true,
			expected: map[uint64]uint64{
				0:  9223372039002276060,
				19: 603979776,
			},
		},
		{
			flags: map[uint64]uint64{
				0:  9223372039002276060,
				19: 603979776,
			},
			flag: 1242,
			set:  false,
			expected: map[uint64]uint64{
				0:  9223372039002276060,
				19: 536870912,
			},
		},
		{
			flags: map[uint64]uint64{
				0:  9223372039002276060,
				19: 603979776,
			},
			flag: 1242,
			set:  true,
			expected: map[uint64]uint64{
				0:  9223372039002276060,
				19: 603979776,
			},
		},
		{
			flags: map[uint64]uint64{
				0:  9223372039002276060,
				19: 536870912,
			},
			flag: 1242,
			set:  false,
			expected: map[uint64]uint64{
				0:  9223372039002276060,
				19: 536870912,
			},
		},
		{
			flags: map[uint64]uint64{
				0: 436324,
				1: 1,
				2: 25623541,
			},
			flag: 64,
			set:  false,
			expected: map[uint64]uint64{
				0: 436324,
				2: 25623541,
			},
		},
		{
			flags:    nil,
			flag:     0,
			set:      false,
			expected: nil,
			err:      errors.New(binflags.ErrorMsgFlagsMapNil),
		},
		{
			flags:    nil,
			flag:     0,
			set:      true,
			expected: nil,
			err:      errors.New(binflags.ErrorMsgFlagsMapNil),
		},
	}
}
