package gobitflags_test

import (
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
			actual, err := gobitflags.SetFlagInt16(prov.flags, prov.flag, prov.val)

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
	val      bool
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
	}
}

func providerSetFlagInt16() []providerTypeSetFlagInt16 {
	return []providerTypeSetFlagInt16{
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
			flag:     15,
			val:      false,
			expected: 32767,
			err:      nil,
		},
		{
			flags:    math.MinInt16,
			flag:     15,
			val:      false,
			expected: 0,
			err:      nil,
		},
	}
}
