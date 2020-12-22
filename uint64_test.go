package binflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/binflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagUint64(t *testing.T) {
	for idx, prov := range providerHasFlagUint64() {
		t.Run(fmt.Sprintf("TestHasFlagUint64_%d", idx), func(t *testing.T) {
			actual := binflags.HasFlagUint64(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagUint64(t *testing.T) {
	for idx, prov := range providerSetFlagUint64() {
		t.Run(fmt.Sprintf("TestSetFlagUint64_%d", idx), func(t *testing.T) {
			actual, err := binflags.SetFlagUint64(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, actual)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagUint64 struct {
	flags    uint64
	flag     uint8
	expected bool
}

type providerTypeSetFlagUint64 struct {
	flags    uint64
	flag     uint8
	set      bool
	expected uint64
	err      error
}

func providerHasFlagUint64() []providerTypeHasFlagUint64 {
	return []providerTypeHasFlagUint64{
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
			flags:    64,
			flag:     6,
			expected: true,
		},
		{
			flags:    30000,
			flag:     4,
			expected: true,
		},
		{
			flags:    30000,
			flag:     14,
			expected: true,
		},
		{
			flags:    30000,
			flag:     3,
			expected: false,
		},
		{
			flags:    822372036854775808,
			flag:     59,
			expected: true,
		},
		{
			flags:    math.MaxUint64,
			flag:     0,
			expected: true,
		},
		{
			flags:    math.MaxUint64,
			flag:     1,
			expected: true,
		},
		{
			flags:    math.MaxUint64,
			flag:     3,
			expected: true,
		},
		{
			flags:    math.MaxUint64,
			flag:     5,
			expected: true,
		},
		{
			flags:    math.MaxUint64,
			flag:     7,
			expected: true,
		},
		{
			flags:    math.MaxUint64,
			flag:     15,
			expected: true,
		},
		{
			flags:    math.MaxUint64,
			flag:     31,
			expected: true,
		},
		{
			flags:    math.MaxUint64,
			flag:     53,
			expected: true,
		},
		{
			flags:    math.MaxUint64,
			flag:     63,
			expected: true,
		},
		{
			flags:    math.MaxUint64,
			flag:     64,
			expected: false,
		},
	}
}

func providerSetFlagUint64() []providerTypeSetFlagUint64 {
	return []providerTypeSetFlagUint64{
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
			flags:    math.MaxUint64,
			flag:     2,
			set:      true,
			expected: math.MaxUint64,
			err:      nil,
		},
		{
			flags:    math.MaxUint64,
			flag:     2,
			set:      false,
			expected: math.MaxUint64 - 4,
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
			flags:    32768,
			flag:     15,
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
			flags:    62768,
			flag:     12,
			set:      false,
			expected: 58672,
			err:      nil,
		},
		{
			flags:    3467463648,
			flag:     26,
			set:      false,
			expected: 3400354784,
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
			flags:    2147483648,
			flag:     31,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    822372036854775808,
			flag:     59,
			set:      false,
			expected: 245911284551352320,
			err:      nil,
		},
		{
			flags:    9223372036854775808,
			flag:     63,
			set:      false,
			expected: 0,
			err:      nil,
		},
		{
			flags:    235234,
			flag:     64,
			set:      true,
			expected: 235234,
			err:      errors.New(binflags.ErrorMsgOutOfRange),
		},
	}
}
