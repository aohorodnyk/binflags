package binflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/binflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagUint8(t *testing.T) {
	for idx, prov := range providerHasFlagUint8() {
		t.Run(fmt.Sprintf("TestHasFlagUint8_%d", idx), func(t *testing.T) {
			actual := binflags.HasFlagUint8(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagUint8(t *testing.T) {
	for idx, prov := range providerSetFlagUint8() {
		t.Run(fmt.Sprintf("TestSetFlagUint8_%d", idx), func(t *testing.T) {
			actual, err := binflags.SetFlagUint8(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

func TestHasFlagByte(t *testing.T) {
	for idx, prov := range providerHasFlagUint8() {
		t.Run(fmt.Sprintf("TestHasFlagByte_%d", idx), func(t *testing.T) {
			actual := binflags.HasFlagByte(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagByte(t *testing.T) {
	for idx, prov := range providerSetFlagUint8() {
		t.Run(fmt.Sprintf("TestSetFlagByte_%d", idx), func(t *testing.T) {
			actual, err := binflags.SetFlagByte(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagUint8 struct {
	flags    uint8
	flag     uint8
	expected bool
}

type providerTypeSetFlagUint8 struct {
	flags    uint8
	flag     uint8
	set      bool
	expected uint8
	err      error
}

func providerHasFlagUint8() []providerTypeHasFlagUint8 {
	return []providerTypeHasFlagUint8{
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
			flags:    82,
			flag:     6,
			expected: true,
		},
		{
			flags:    math.MaxUint8,
			flag:     0,
			expected: true,
		},
		{
			flags:    math.MaxUint8,
			flag:     1,
			expected: true,
		},
		{
			flags:    math.MaxUint8,
			flag:     3,
			expected: true,
		},
		{
			flags:    math.MaxUint8,
			flag:     5,
			expected: true,
		},
		{
			flags:    math.MaxUint8,
			flag:     7,
			expected: true,
		},
		{
			flags:    math.MaxUint8,
			flag:     8,
			expected: false,
		},
	}
}

func providerSetFlagUint8() []providerTypeSetFlagUint8 {
	return []providerTypeSetFlagUint8{
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
			flags:    math.MaxUint8,
			flag:     2,
			set:      true,
			expected: math.MaxUint8,
			err:      nil,
		},
		{
			flags:    math.MaxUint8,
			flag:     2,
			set:      false,
			expected: math.MaxUint8 - 4,
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
			flags:    235,
			flag:     5,
			set:      false,
			expected: 203,
			err:      nil,
		},
		{
			flags:    5,
			flag:     8,
			set:      true,
			expected: 5,
			err:      errors.New(binflags.ErrorMsgOutOfRange),
		},
	}
}
