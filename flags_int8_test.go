package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagInt8(t *testing.T) {
	for idx, prov := range providerHasFlagInt8() {
		t.Run(fmt.Sprintf("TestHasFlagInt8_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagInt8(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagInt8(t *testing.T) {
	for idx, prov := range providerSetFlagInt8() {
		t.Run(fmt.Sprintf("TestSetFlagInt8_%d", idx), func(t *testing.T) {
			actual, err := gobitflags.SetFlagInt8(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagInt8 struct {
	flags    int8
	flag     uint8
	expected bool
}

type providerTypeSetFlagInt8 struct {
	flags    int8
	flag     uint8
	set      bool
	expected int8
	err      error
}

func providerHasFlagInt8() []providerTypeHasFlagInt8 {
	return []providerTypeHasFlagInt8{
		{
			flags:    0,
			flag:     0,
			expected: false,
		},
		{
			flags:    52,
			flag:     7,
			expected: false,
		},
		{
			flags:    math.MinInt8,
			flag:     7,
			expected: true,
		},
		{
			flags:    math.MinInt8,
			flag:     3,
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
			flag:     8,
			expected: false,
		},
	}
}

func providerSetFlagInt8() []providerTypeSetFlagInt8 {
	return []providerTypeSetFlagInt8{
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
			flags:    -1,
			flag:     5,
			set:      false,
			expected: -33,
			err:      nil,
		},
		{
			flags:    -1,
			flag:     7,
			set:      false,
			expected: 127,
			err:      nil,
		},
		{
			flags:    math.MinInt8,
			flag:     7,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    math.MinInt8,
			flag:     8,
			set:      false,
			expected: math.MinInt8,
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
	}
}
