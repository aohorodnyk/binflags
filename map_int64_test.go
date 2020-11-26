package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagMapInt64(t *testing.T) {
	for idx, prov := range providerHasFlagMapInt64() {
		t.Run(fmt.Sprintf("TestHasFlagMapInt64_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagMapInt64(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagMapInt64(t *testing.T) {
	for idx, prov := range providerSetFlagMapInt64() {
		t.Run(fmt.Sprintf("TestHasFlagMapInt64_%d", idx), func(t *testing.T) {
			err := gobitflags.SetFlagMapInt64(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagMapInt64 struct {
	flags    map[uint64]int64
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapInt64 struct {
	flags    map[uint64]int64
	flag     uint64
	set      bool
	expected map[uint64]int64
	err      error
}

func providerHasFlagMapInt64() []providerTypeHasFlagMapInt64 {
	return []providerTypeHasFlagMapInt64{
		{
			flags:    map[uint64]int64{},
			flag:     0,
			expected: false,
		},
		{
			flags:    map[uint64]int64{},
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
			flags:    map[uint64]int64{},
			flag:     2356241245436,
			expected: false,
		},
		{
			flags: map[uint64]int64{
				0: 1,
			},
			flag:     0,
			expected: true,
		},
		{
			flags: map[uint64]int64{
				0: 32768,
			},
			flag:     15,
			expected: true,
		},
		{
			flags: map[uint64]int64{
				0: 2147483648,
			},
			flag:     31,
			expected: true,
		},
		{
			flags: map[uint64]int64{
				0: 4503601774854144,
			},
			flag:     52,
			expected: true,
		},
		{
			flags: map[uint64]int64{
				0: 4503599627370496,
			},
			flag:     52,
			expected: true,
		},
		{
			flags: map[uint64]int64{
				0: -9218868473709551616,
			},
			flag:     63,
			expected: true,
		},
		{
			flags: map[uint64]int64{
				0:     -9218868473709551616,
				36644: 262144,
			},
			flag:     2345234,
			expected: true,
		},
	}
}

func providerSetFlagMapInt64() []providerTypeSetFlagMapInt64 {
	return []providerTypeSetFlagMapInt64{
		{
			flags: map[uint64]int64{},
			flag:  0,
			set:   true,
			expected: map[uint64]int64{
				0: 1,
			},
		},
		{
			flags: map[uint64]int64{},
			flag:  7,
			set:   true,
			expected: map[uint64]int64{
				0: 128,
			},
		},
		{
			flags: map[uint64]int64{},
			flag:  15,
			set:   true,
			expected: map[uint64]int64{
				0: 32768,
			},
		},
		{
			flags: map[uint64]int64{
				0: 92,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]int64{
				0: 220,
			},
		},
		{
			flags: map[uint64]int64{
				0: 220,
			},
			flag: 7,
			set:  false,
			expected: map[uint64]int64{
				0: 92,
			},
		},
		{
			flags: map[uint64]int64{
				0: 220,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]int64{
				0: 220,
			},
		},
		{
			flags: map[uint64]int64{
				0: 220,
			},
			flag: 1,
			set:  false,
			expected: map[uint64]int64{
				0: 220,
			},
		},
		{
			flags: map[uint64]int64{
				0: 220,
			},
			flag: 14,
			set:  true,
			expected: map[uint64]int64{
				0: 16604,
			},
		},
		{
			flags: map[uint64]int64{
				0: 16604,
			},
			flag: 31,
			set:  true,
			expected: map[uint64]int64{
				0: 2147500252,
			},
		},
		{
			flags: map[uint64]int64{
				0: 2147500252,
			},
			flag: 14,
			set:  false,
			expected: map[uint64]int64{
				0: 2147483868,
			},
		},
		{
			flags: map[uint64]int64{
				0: 2147500252,
			},
			flag: 60,
			set:  true,
			expected: map[uint64]int64{
				0: 1152921506754347228,
			},
		},
		{
			flags: map[uint64]int64{
				0: 2147500252,
			},
			flag: 63,
			set:  true,
			expected: map[uint64]int64{
				0: -9223372034707275556,
			},
		},
		{
			flags: map[uint64]int64{
				0: -9223372034707275556,
			},
			flag: 63,
			set:  false,
			expected: map[uint64]int64{
				0: 2147500252,
			},
		},
		{
			flags: map[uint64]int64{
				0: -9223372034707275556,
			},
			flag: 1245,
			set:  true,
			expected: map[uint64]int64{
				0:  -9223372034707275556,
				19: 536870912,
			},
		},
		{
			flags: map[uint64]int64{
				0:  -9223372034707275556,
				19: 536870912,
			},
			flag: 1242,
			set:  true,
			expected: map[uint64]int64{
				0:  -9223372034707275556,
				19: 603979776,
			},
		},
		{
			flags: map[uint64]int64{
				0:  -9223372034707275556,
				19: 603979776,
			},
			flag: 1242,
			set:  false,
			expected: map[uint64]int64{
				0:  -9223372034707275556,
				19: 536870912,
			},
		},
		{
			flags: map[uint64]int64{
				0:  -9223372034707275556,
				19: 603979776,
			},
			flag: 1242,
			set:  true,
			expected: map[uint64]int64{
				0:  -9223372034707275556,
				19: 603979776,
			},
		},
		{
			flags: map[uint64]int64{
				0:  -9223372034707275556,
				19: 536870912,
			},
			flag: 1242,
			set:  false,
			expected: map[uint64]int64{
				0:  -9223372034707275556,
				19: 536870912,
			},
		},
		{
			flags: map[uint64]int64{
				0: 436324,
				1: 1,
				2: 25623541,
			},
			flag: 64,
			set:  false,
			expected: map[uint64]int64{
				0: 436324,
				2: 25623541,
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
