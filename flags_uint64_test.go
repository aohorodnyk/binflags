package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagUint64(t *testing.T) {
	for idx, prov := range providerHasFlagUint64() {
		t.Run(fmt.Sprintf("TestHasFlagUint64_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.HasFlagUint64(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

func TestSetFlagUint64(t *testing.T) {
	for idx, prov := range providerSetFlagUint64() {
		t.Run(fmt.Sprintf("TestSetFlagUint64_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.SetFlagUint64(prov.flags, prov.flag, prov.val)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagUint64 struct {
	flags    uint64
	flag     uint8
	expected bool
	err      error
}

type providerTypeSetFlagUint64 struct {
	flags    uint64
	flag     uint8
	val      bool
	expected uint64
	err      error
}

func providerHasFlagUint64() []providerTypeHasFlagUint64 {
	return []providerTypeHasFlagUint64{
		{
			flags:    0,
			flag:     0,
			expected: false,
			err:      nil,
		},
		{
			flags:    0,
			flag:     1,
			expected: false,
			err:      nil,
		},
		{
			flags:    0,
			flag:     3,
			expected: false,
			err:      nil,
		},
		{
			flags:    0,
			flag:     5,
			expected: false,
			err:      nil,
		},
		{
			flags:    0,
			flag:     7,
			expected: false,
			err:      nil,
		},
		{
			flags:    2,
			flag:     0,
			expected: false,
			err:      nil,
		},
		{
			flags:    1,
			flag:     0,
			expected: true,
			err:      nil,
		},
		{
			flags:    3,
			flag:     0,
			expected: true,
			err:      nil,
		},
		{
			flags:    26,
			flag:     1,
			expected: true,
			err:      nil,
		},
		{
			flags:    26,
			flag:     3,
			expected: true,
			err:      nil,
		},
		{
			flags:    26,
			flag:     4,
			expected: true,
			err:      nil,
		},
		{
			flags:    128,
			flag:     7,
			expected: true,
			err:      nil,
		},
		{
			flags:    64,
			flag:     6,
			expected: true,
			err:      nil,
		},
		{
			flags:    30000,
			flag:     4,
			expected: true,
			err:      nil,
		},
		{
			flags:    30000,
			flag:     14,
			expected: true,
			err:      nil,
		},
		{
			flags:    30000,
			flag:     3,
			expected: false,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     0,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     1,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     3,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     5,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     7,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     15,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     31,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     53,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     63,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     64,
			expected: false,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}

func providerSetFlagUint64() []providerTypeSetFlagUint64 {
	return []providerTypeSetFlagUint64{
		{
			flags:    0,
			flag:     0,
			val:      true,
			expected: 1,
			err:      nil,
		},
		{
			flags:    0,
			flag:     1,
			val:      true,
			expected: 2,
			err:      nil,
		},
		{
			flags:    0,
			flag:     2,
			val:      true,
			expected: 4,
			err:      nil,
		},
		{
			flags:    1,
			flag:     2,
			val:      true,
			expected: 5,
			err:      nil,
		},
		{
			flags:    53,
			flag:     1,
			val:      true,
			expected: 55,
			err:      nil,
		},
		{
			flags:    53,
			flag:     0,
			val:      true,
			expected: 53,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     2,
			val:      true,
			expected: math.MaxUint64,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     2,
			val:      false,
			expected: math.MaxUint64 - 4,
			err:      nil,
		},
		{
			flags:    5,
			flag:     7,
			val:      true,
			expected: 133,
			err:      nil,
		},
		{
			flags:    128,
			flag:     7,
			val:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    32768,
			flag:     15,
			val:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    2147483648,
			flag:     31,
			val:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    2147483648,
			flag:     31,
			val:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    9223372036854775808,
			flag:     63,
			val:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    5,
			flag:     64,
			val:      true,
			expected: 5,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}
