package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagMapUint32(t *testing.T) {
	for idx, prov := range providerHasFlagMapUint32() {
		t.Run(fmt.Sprintf("TestHasFlagMapUint32_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagMapUint32(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagMapUint32(t *testing.T) {
	for idx, prov := range providerSetFlagMapUint32() {
		t.Run(fmt.Sprintf("TestHasFlagMapUint32_%d", idx), func(t *testing.T) {
			err := gobitflags.SetFlagMapUint32(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagMapUint32 struct {
	flags    map[uint64]uint32
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapUint32 struct {
	flags    map[uint64]uint32
	flag     uint64
	set      bool
	expected map[uint64]uint32
	err      error
}

func providerHasFlagMapUint32() []providerTypeHasFlagMapUint32 {
	return []providerTypeHasFlagMapUint32{
		{
			flags:    map[uint64]uint32{},
			flag:     0,
			expected: false,
		},
		{
			flags:    map[uint64]uint32{},
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
			flags:    map[uint64]uint32{},
			flag:     2356241245436,
			expected: false,
		},
		{
			flags: map[uint64]uint32{
				0: 1,
			},
			flag:     0,
			expected: true,
		},
		{
			flags: map[uint64]uint32{
				0: 32768,
			},
			flag:     15,
			expected: true,
		},
		{
			flags: map[uint64]uint32{
				0: 2147483648,
			},
			flag:     31,
			expected: true,
		},
		{
			flags: map[uint64]uint32{
				0: 2147483648,
				1: 1,
			},
			flag:     32,
			expected: true,
		},
		{
			flags: map[uint64]uint32{
				0: 2147483648,
				1: 1449509,
			},
			flag:     52,
			expected: true,
		},
		{
			flags: map[uint64]uint32{
				0:    2147483648,
				1323: 1449509,
			},
			flag:     42356,
			expected: true,
		},
	}
}

func providerSetFlagMapUint32() []providerTypeSetFlagMapUint32 {
	return []providerTypeSetFlagMapUint32{
		{
			flags: map[uint64]uint32{},
			flag:  0,
			set:   true,
			expected: map[uint64]uint32{
				0: 1,
			},
		},
		{
			flags: map[uint64]uint32{},
			flag:  7,
			set:   true,
			expected: map[uint64]uint32{
				0: 128,
			},
		},
		{
			flags: map[uint64]uint32{},
			flag:  15,
			set:   true,
			expected: map[uint64]uint32{
				0: 32768,
			},
		},
		{
			flags: map[uint64]uint32{
				0: 92,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]uint32{
				0: 220,
			},
		},
		{
			flags: map[uint64]uint32{
				0: 220,
			},
			flag: 7,
			set:  false,
			expected: map[uint64]uint32{
				0: 92,
			},
		},
		{
			flags: map[uint64]uint32{
				0: 220,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]uint32{
				0: 220,
			},
		},
		{
			flags: map[uint64]uint32{
				0: 220,
			},
			flag: 1,
			set:  false,
			expected: map[uint64]uint32{
				0: 220,
			},
		},
		{
			flags: map[uint64]uint32{
				0: 220,
			},
			flag: 14,
			set:  true,
			expected: map[uint64]uint32{
				0: 16604,
			},
		},
		{
			flags: map[uint64]uint32{
				0: 16604,
			},
			flag: 31,
			set:  true,
			expected: map[uint64]uint32{
				0: 2147500252,
			},
		},
		{
			flags: map[uint64]uint32{
				0: 2147500252,
			},
			flag: 14,
			set:  false,
			expected: map[uint64]uint32{
				0: 2147483868,
			},
		},
		{
			flags: map[uint64]uint32{
				0: 2147500252,
			},
			flag: 5234,
			set:  true,
			expected: map[uint64]uint32{
				0:   2147500252,
				163: 262144,
			},
		},
		{
			flags:    nil,
			flag:     0,
			set:      false,
			expected: nil,
			err:      errors.New(gobitflags.ErrorMsgFlagsMapNil),
		},
		{
			flags:    nil,
			flag:     0,
			set:      true,
			expected: nil,
			err:      errors.New(gobitflags.ErrorMsgFlagsMapNil),
		},
	}
}
