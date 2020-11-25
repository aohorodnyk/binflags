package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagUint8(t *testing.T) {
	for idx, prov := range providerHasFlagUint8() {
		t.Run(fmt.Sprintf("TestHasFlagUint8_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.HasFlagUint8(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

func TestSetFlagUint8(t *testing.T) {
	for idx, prov := range providerSetFlagUint8() {
		t.Run(fmt.Sprintf("TestSetFlagUint8_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.SetFlagUint8(prov.flags, prov.flag, prov.val)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

func TestHasFlagByte(t *testing.T) {
	for idx, prov := range providerHasFlagUint8() {
		t.Run(fmt.Sprintf("TestHasFlagByte_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.HasFlagByte(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

func TestSetFlagByte(t *testing.T) {
	for idx, prov := range providerSetFlagUint8() {
		t.Run(fmt.Sprintf("TestSetFlagByte_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.SetFlagByte(prov.flags, prov.flag, prov.val)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagUint8 struct {
	flags    uint8
	flag     uint8
	expected bool
	err      error
}

type providerTypeSetFlagUint8 struct {
	flags    uint8
	flag     uint8
	val      bool
	expected uint8
	err      error
}

func providerHasFlagUint8() []providerTypeHasFlagUint8 {
	return []providerTypeHasFlagUint8{
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
			flags:    math.MaxUint8,
			flag:     0,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint8,
			flag:     1,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint8,
			flag:     3,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint8,
			flag:     5,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint8,
			flag:     7,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MaxUint8,
			flag:     8,
			expected: false,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}

func providerSetFlagUint8() []providerTypeSetFlagUint8 {
	return []providerTypeSetFlagUint8{
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
			flags:    math.MaxUint8,
			flag:     2,
			val:      true,
			expected: math.MaxUint8,
			err:      nil,
		},
		{
			flags:    math.MaxUint8,
			flag:     2,
			val:      false,
			expected: math.MaxUint8 - 4,
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
			flag:     8,
			val:      true,
			expected: 5,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
		{
			flags:    53,
			flag:     10,
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
	}
}
