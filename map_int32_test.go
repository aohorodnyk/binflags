package binflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/binflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagMapInt32(t *testing.T) {
	for idx, prov := range providerHasFlagMapInt32() {
		t.Run(fmt.Sprintf("TestHasFlagMapInt32_%d", idx), func(t *testing.T) {
			actual := binflags.HasFlagMapInt32(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagMapInt32(t *testing.T) {
	for idx, prov := range providerSetFlagMapInt32() {
		t.Run(fmt.Sprintf("TestHasFlagMapInt32_%d", idx), func(t *testing.T) {
			err := binflags.SetFlagMapInt32(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagMapInt32 struct {
	flags    map[uint64]int32
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapInt32 struct {
	flags    map[uint64]int32
	flag     uint64
	set      bool
	expected map[uint64]int32
	err      error
}

func providerHasFlagMapInt32() []providerTypeHasFlagMapInt32 {
	return []providerTypeHasFlagMapInt32{
		{
			flags:    map[uint64]int32{},
			flag:     0,
			expected: false,
		},
		{
			flags:    map[uint64]int32{},
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
			flags:    map[uint64]int32{},
			flag:     2356241245436,
			expected: false,
		},
		{
			flags: map[uint64]int32{
				0: 1,
			},
			flag:     0,
			expected: true,
		},
		{
			flags: map[uint64]int32{
				0: 32768,
			},
			flag:     15,
			expected: true,
		},
		{
			flags: map[uint64]int32{
				0: -2147483648,
			},
			flag:     31,
			expected: true,
		},
		{
			flags: map[uint64]int32{
				0: -2147483648,
				1: 1,
			},
			flag:     32,
			expected: true,
		},
		{
			flags: map[uint64]int32{
				0: -2147483648,
				1: 1449509,
			},
			flag:     52,
			expected: true,
		},
		{
			flags: map[uint64]int32{
				0:    -2147483648,
				1323: 1449509,
			},
			flag:     42356,
			expected: true,
		},
	}
}

func providerSetFlagMapInt32() []providerTypeSetFlagMapInt32 {
	return []providerTypeSetFlagMapInt32{
		{
			flags: map[uint64]int32{},
			flag:  0,
			set:   true,
			expected: map[uint64]int32{
				0: 1,
			},
		},
		{
			flags: map[uint64]int32{},
			flag:  7,
			set:   true,
			expected: map[uint64]int32{
				0: 128,
			},
		},
		{
			flags: map[uint64]int32{},
			flag:  15,
			set:   true,
			expected: map[uint64]int32{
				0: 32768,
			},
		},
		{
			flags: map[uint64]int32{
				0: 92,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]int32{
				0: 220,
			},
		},
		{
			flags: map[uint64]int32{
				0: 220,
			},
			flag: 7,
			set:  false,
			expected: map[uint64]int32{
				0: 92,
			},
		},
		{
			flags: map[uint64]int32{
				0: 220,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]int32{
				0: 220,
			},
		},
		{
			flags: map[uint64]int32{
				0: 220,
			},
			flag: 1,
			set:  false,
			expected: map[uint64]int32{
				0: 220,
			},
		},
		{
			flags: map[uint64]int32{
				0: 220,
			},
			flag: 14,
			set:  true,
			expected: map[uint64]int32{
				0: 16604,
			},
		},
		{
			flags: map[uint64]int32{
				0: 16604,
			},
			flag: 31,
			set:  true,
			expected: map[uint64]int32{
				0: -2147467044,
			},
		},
		{
			flags: map[uint64]int32{
				0: -2147467044,
			},
			flag: 14,
			set:  false,
			expected: map[uint64]int32{
				0: -2147483428,
			},
		},
		{
			flags: map[uint64]int32{
				0: -2147467044,
			},
			flag: 5234,
			set:  true,
			expected: map[uint64]int32{
				0:   -2147467044,
				163: 262144,
			},
		},
		{
			flags: map[uint64]int32{
				0: 436324,
				1: 1,
				2: 25623541,
			},
			flag: 32,
			set:  false,
			expected: map[uint64]int32{
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
