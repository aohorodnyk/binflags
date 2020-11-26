package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagInt16(t *testing.T) {
	for idx, prov := range providerHasFlagInt16() {
		t.Run(fmt.Sprintf("TestHasFlagInt16_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.HasFlagInt16(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

func TestSetFlagInt16(t *testing.T) {
	for idx, prov := range providerSetFlagInt16() {
		t.Run(fmt.Sprintf("TestSetFlagInt16_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.SetFlagInt16(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagInt16 struct {
	flags    int16
	flag     uint8
	expected bool
	err      error
}

type providerTypeSetFlagInt16 struct {
	flags    int16
	flag     uint8
	set      bool
	expected int16
	err      error
}

func providerHasFlagInt16() []providerTypeHasFlagInt16 {
	return []providerTypeHasFlagInt16{
		{
			flags:    0,
			flag:     0,
			expected: false,
			err:      nil,
		},
		{
			flags:    52,
			flag:     7,
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
			flags:    math.MinInt16,
			flag:     15,
			expected: true,
			err:      nil,
		},
		{
			flags:    math.MinInt16,
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
			flag:     16,
			expected: false,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}

func providerSetFlagInt16() []providerTypeSetFlagInt16 {
	return []providerTypeSetFlagInt16{
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
			flag:     15,
			set:      false,
			expected: 32767,
			err:      nil,
		},
		{
			flags:    math.MinInt16,
			flag:     15,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    math.MinInt16,
			flag:     16,
			set:      false,
			expected: math.MinInt16,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}
