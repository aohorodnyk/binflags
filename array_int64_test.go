package binflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/binflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagArrayInt64(t *testing.T) {
	for idx, prov := range providerHasFlagArrayInt64() {
		t.Run(fmt.Sprintf("TestHasFlagArrayInt64_%d", idx), func(t *testing.T) {
			actual := binflags.HasFlagArrayInt64(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagArrayInt64(t *testing.T) {
	for idx, prov := range providerSetFlagArrayInt64() {
		t.Run(fmt.Sprintf("TestHasFlagArrayInt64_%d", idx), func(t *testing.T) {
			err := binflags.SetFlagArrayInt64(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagArrayInt64 struct {
	flags    []int64
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayInt64 struct {
	flags    []int64
	flag     uint64
	set      bool
	expected []int64
	err      error
}

func providerHasFlagArrayInt64() []providerTypeHasFlagArrayInt64 {
	return []providerTypeHasFlagArrayInt64{
		{
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			flags:    []int64{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			flags:    []int64{115},
			flag:     5,
			expected: true,
		},
		{
			flags:    []int64{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			flags:    []int64{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			flags:    []int64{-9223372036854775808, 1, 0, 0},
			flag:     63,
			expected: true,
		},
		{
			flags:    []int64{0, 1, 0, 0},
			flag:     64,
			expected: true,
		},
		{
			flags:    []int64{0, 0, 1, 0},
			flag:     128,
			expected: true,
		},
		{
			flags:    []int64{0, 0, 9, 0},
			flag:     131,
			expected: true,
		},
		{
			flags:    []int64{0, 0, 217, 0},
			flag:     135,
			expected: true,
		},
		{
			flags:    []int64{0, 0, 217, 0},
			flag:     134,
			expected: true,
		},
		{
			flags:    []int64{0, 0, 217, 0},
			flag:     69,
			expected: false,
		},
		{
			flags:    []int64{0, 0, 217, 0},
			flag:     72,
			expected: false,
		},
	}
}

func providerSetFlagArrayInt64() []providerTypeSetFlagArrayInt64 {
	return []providerTypeSetFlagArrayInt64{
		{
			flags:    []int64{0},
			flag:     0,
			set:      true,
			expected: []int64{1},
			err:      nil,
		},
		{
			flags:    []int64{1},
			flag:     6,
			set:      true,
			expected: []int64{65},
			err:      nil,
		},
		{
			flags:    []int64{65},
			flag:     6,
			set:      false,
			expected: []int64{1},
			err:      nil,
		},
		{
			flags:    []int64{65},
			flag:     12,
			set:      true,
			expected: []int64{4161},
			err:      nil,
		},
		{
			flags:    []int64{65},
			flag:     63,
			set:      true,
			expected: []int64{-9223372036854775743},
			err:      nil,
		},
		{
			flags:    []int64{65, 0},
			flag:     64,
			set:      true,
			expected: []int64{65, 1},
			err:      nil,
		},
		{
			flags:    []int64{65, 0},
			flag:     127,
			set:      true,
			expected: []int64{65, -9223372036854775808},
			err:      nil,
		},
		{
			flags:    []int64{65, 0, 1235, 724, 635},
			flag:     132,
			set:      true,
			expected: []int64{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			flags:    []int64{65, 0, 1235, 724, 635},
			flag:     132,
			set:      false,
			expected: []int64{65, 0, 1219, 724, 635},
			err:      nil,
		},
		{
			flags:    []int64{0},
			flag:     64,
			set:      true,
			expected: []int64{0},
			err:      errors.New(binflags.ErrorMsgOutOfRange),
		},
		{
			flags:    nil,
			flag:     0,
			set:      false,
			expected: nil,
			err:      errors.New(binflags.ErrorMsgFlagsArrayNil),
		},
	}
}
