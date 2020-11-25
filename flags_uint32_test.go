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
			actual, err := gobitflags.HasFlagUint32(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

func TestSetFlagUint32(t *testing.T) {
	for idx, prov := range providerSetFlagUint32() {
		t.Run(fmt.Sprintf("TestSetFlagUint32_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.SetFlagUint32(prov.flags, prov.flag, prov.val)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagUint32 struct {
	flags    uint32
	flag     uint8
	expected bool
	err      error
}

type providerTypeSetFlagUint32 struct {
	flags    uint32
	flag     uint8
	val      bool
	expected uint32
	err      error
}

func providerHasFlagUint32() []providerTypeHasFlagUint32 {
	return []providerTypeHasFlagUint32{
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
			flags:    math.MaxUint32,
			flag:     0,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     1,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     3,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     5,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     7,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     31,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     32,
			expected: false,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}

func providerSetFlagUint32() []providerTypeSetFlagUint32 {
	return []providerTypeSetFlagUint32{
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
			flags:    math.MaxUint32,
			flag:     2,
			val:      true,
			expected: math.MaxUint32,
			err:      nil,
		},
		{
			flags:    math.MaxUint32,
			flag:     2,
			val:      false,
			expected: math.MaxUint32 - 4,
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
			flags:    5,
			flag:     32,
			val:      true,
			expected: 5,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
		{
			flags:    53,
			flag:     35,
			val:      false,
			expected: 53,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
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
	}
}
