package binflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/binflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagInt64(t *testing.T) {
	for idx, prov := range providerHasFlagInt64() {
		t.Run(fmt.Sprintf("TestHasFlagInt64_%d", idx), func(t *testing.T) {
			actual := binflags.HasFlagInt64(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagInt64(t *testing.T) {
	for idx, prov := range providerSetFlagInt64() {
		t.Run(fmt.Sprintf("TestSetFlagInt64_%d", idx), func(t *testing.T) {
			actual, err := binflags.SetFlagInt64(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagInt64 struct {
	flags    int64
	flag     uint8
	expected bool
}

type providerTypeSetFlagInt64 struct {
	flags    int64
	flag     uint8
	set      bool
	expected int64
	err      error
}

func providerHasFlagInt64() []providerTypeHasFlagInt64 {
	return []providerTypeHasFlagInt64{
		{
			flags:    0,
			flag:     0,
			expected: false,
		},
		{
			flags:    52,
			flag:     15,
			expected: false,
		},
		{
			flags:    52,
			flag:     4,
			expected: true,
		},
		{
			flags:    math.MinInt64,
			flag:     31,
			expected: false,
		},
		{
			flags:    math.MinInt64,
			flag:     63,
			expected: true,
		},
		{
			flags:    math.MinInt64,
			flag:     14,
			expected: false,
		},
		{
			flags:    -1,
			flag:     3,
			expected: true,
		},
		{
			flags:    -1,
			flag:     7,
			expected: true,
		},
		{
			flags:    -1,
			flag:     64,
			expected: false,
		},
	}
}

func providerSetFlagInt64() []providerTypeSetFlagInt64 {
	return []providerTypeSetFlagInt64{
		{
			flags:    0,
			flag:     0,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    0,
			flag:     0,
			set:      true,
			expected: 1,
			err:      nil,
		},
		{
			flags:    1,
			flag:     1,
			set:      true,
			expected: 3,
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
			set:      false,
			expected: 52,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     5,
			set:      true,
			expected: -1,
			err:      nil,
		},
		{
			flags:    2567,
			flag:     11,
			set:      false,
			expected: 519,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     5,
			set:      false,
			expected: -33,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     31,
			set:      false,
			expected: -2147483649,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     63,
			set:      false,
			expected: 9223372036854775807,
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
			flags:    math.MinInt64,
			flag:     63,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    math.MinInt64,
			flag:     64,
			set:      true,
			expected: math.MinInt64,
			err:      errors.New(binflags.ErrorMsgOutOfRange),
		},
	}
}
