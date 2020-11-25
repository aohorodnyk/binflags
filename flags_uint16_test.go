package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagUint16(t *testing.T) {
	for idx, prov := range providerHasFlagUint16() {
		t.Run(fmt.Sprintf("TestHasFlagUint16_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.HasFlagUint16(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

func TestSetFlagUint16(t *testing.T) {
	for idx, prov := range providerSetFlagUint16() {
		t.Run(fmt.Sprintf("TestSetFlagUint16_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.SetFlagUint16(prov.flags, prov.flag, prov.val)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagUint16 struct {
	flags    uint16
	flag     uint8
	expected bool
	err      error
}

type providerTypeSetFlagUint16 struct {
	flags    uint16
	flag     uint8
	val      bool
	expected uint16
	err      error
}

func providerHasFlagUint16() []providerTypeHasFlagUint16 {
	return []providerTypeHasFlagUint16{
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
			flags:    math.MaxUint16,
			flag:     0,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint16,
			flag:     1,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint16,
			flag:     3,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint16,
			flag:     5,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint16,
			flag:     7,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint16,
			flag:     16,
			expected: false,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}

func providerSetFlagUint16() []providerTypeSetFlagUint16 {
	return []providerTypeSetFlagUint16{
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
			flags:    math.MaxUint16,
			flag:     2,
			val:      true,
			expected: math.MaxUint16,
			err:      nil,
		},
		{
			flags:    math.MaxUint16,
			flag:     2,
			val:      false,
			expected: math.MaxUint16 - 4,
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
			flag:     16,
			val:      true,
			expected: 5,
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
	}
}
