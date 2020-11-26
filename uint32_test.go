package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagUint32(t *testing.T) {
	for idx, prov := range providerHasFlagUint32() {
		t.Run(fmt.Sprintf("TestHasFlagUint32_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagUint32(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagUint32(t *testing.T) {
	for idx, prov := range providerSetFlagUint32() {
		t.Run(fmt.Sprintf("TestSetFlagUint32_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.SetFlagUint32(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagUint32 struct {
	flags    uint32
	flag     uint8
	expected bool
}

type providerTypeSetFlagUint32 struct {
	flags    uint32
	flag     uint8
	set      bool
	expected uint32
	err      error
}

func providerHasFlagUint32() []providerTypeHasFlagUint32 {
	return []providerTypeHasFlagUint32{
		{
			flags:    0,
			flag:     0,
			expected: false,
		},
		{
			flags:    0,
			flag:     1,
			expected: false,
		},
		{
			flags:    0,
			flag:     3,
			expected: false,
		},
		{
			flags:    0,
			flag:     5,
			expected: false,
		},
		{
			flags:    0,
			flag:     7,
			expected: false,
		},
		{
			flags:    2,
			flag:     0,
			expected: false,
		},
		{
			flags:    1,
			flag:     0,
			expected: true,
		},
		{
			flags:    3,
			flag:     0,
			expected: true,
		},
		{
			flags:    26,
			flag:     1,
			expected: true,
		},
		{
			flags:    26,
			flag:     3,
			expected: true,
		},
		{
			flags:    26,
			flag:     4,
			expected: true,
		},
		{
			flags:    128,
			flag:     7,
			expected: true,
		},
		{
			flags:    64,
			flag:     6,
			expected: true,
		},
		{
			flags:    30000,
			flag:     4,
			expected: true,
		},
		{
			flags:    30000,
			flag:     14,
			expected: true,
		},
		{
			flags:    30000,
			flag:     3,
			expected: false,
		},
		{
			flags:    math.MaxUint32,
			flag:     0,
			expected: true,
		},
		{
			flags:    math.MaxUint32,
			flag:     1,
			expected: true,
		},
		{
			flags:    math.MaxUint32,
			flag:     3,
			expected: true,
		},
		{
			flags:    math.MaxUint32,
			flag:     5,
			expected: true,
		},
		{
			flags:    math.MaxUint32,
			flag:     7,
			expected: true,
		},
		{
			flags:    math.MaxUint32,
			flag:     31,
			expected: true,
		},
		{
			flags:    math.MaxUint32,
			flag:     32,
			expected: false,
		},
	}
}

func providerSetFlagUint32() []providerTypeSetFlagUint32 {
	return []providerTypeSetFlagUint32{
		{
			flags:    0,
			flag:     0,
			set:      true,
			expected: 1,
			err:      nil,
		},
		{
			flags:    0,
			flag:     1,
			set:      true,
			expected: 2,
			err:      nil,
		},
		{
			flags:    0,
			flag:     2,
			set:      true,
			expected: 4,
			err:      nil,
		},
		{
			flags:    1,
			flag:     2,
			set:      true,
			expected: 5,
			err:      nil,
		},
		{
			flags:    53,
			flag:     1,
			set:      true,
			expected: 55,
			err:      nil,
		},
		{
			flags:    53,
			flag:     0,
			set:      true,
			expected: 53,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     2,
			set:      true,
			expected: math.MaxUint32,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     2,
			set:      false,
			expected: math.MaxUint32 - 4,
			err:      nil,
		},
		{
			flags:    235,
			flag:     5,
			set:      false,
			expected: 203,
			err:      nil,
		},
		{
			flags:    62768,
			flag:     12,
			set:      false,
			expected: 58672,
			err:      nil,
		},
		{
			flags:    5,
			flag:     7,
			set:      true,
			expected: 133,
			err:      nil,
		},
		{
			flags:    128,
			flag:     7,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    32768,
			flag:     15,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    2147483648,
			flag:     31,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    3467463648,
			flag:     26,
			set:      false,
			expected: 3400354784,
			err:      nil,
		},
		{
			flags:    5,
			flag:     32,
			set:      true,
			expected: 5,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}
