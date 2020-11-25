package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagInt32(t *testing.T) {
	for idx, prov := range providerHasFlagInt32() {
		t.Run(fmt.Sprintf("TestHasFlagInt32_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.HasFlagInt32(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

func TestSetFlagInt32(t *testing.T) {
	for idx, prov := range providerSetFlagInt32() {
		t.Run(fmt.Sprintf("TestSetFlagInt32_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.SetFlagInt32(prov.flags, prov.flag, prov.val)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagInt32 struct {
	flags    int32
	flag     uint8
	expected bool
	err      error
}

type providerTypeSetFlagInt32 struct {
	flags    int32
	flag     uint8
	val      bool
	expected int32
	err      error
}

func providerHasFlagInt32() []providerTypeHasFlagInt32 {
	return []providerTypeHasFlagInt32{
		{
			flags:    0,
			flag:     0,
			expected: false,
			err:      nil,
		},
		{
			flags:    52,
			flag:     15,
			expected: false,
			err:      nil,
		},
		{
			flags:    math.MinInt32,
			flag:     31,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MinInt32,
			flag:     14,
			expected: false,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     3,
			expected: true,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     7,
			expected: true,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     32,
			expected: false,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}

func providerSetFlagInt32() []providerTypeSetFlagInt32 {
	return []providerTypeSetFlagInt32{
		{
			flags:    0,
			flag:     0,
			val:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    0,
			flag:     0,
			val:      true,
			expected: 1,
			err:      nil,
		},
		{
			flags:    1,
			flag:     1,
			val:      true,
			expected: 3,
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
			val:      false,
			expected: 52,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     5,
			val:      true,
			expected: -1,
			err:      nil,
		},
		{
			flags:    2567,
			flag:     11,
			val:      false,
			expected: 519,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     5,
			val:      false,
			expected: -33,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     31,
			val:      false,
			expected: 2147483647,
			err:      nil,
		},
		{
			flags:    math.MinInt32,
			flag:     31,
			val:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    math.MinInt32,
			flag:     32,
			val:      true,
			expected: math.MinInt32,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}
