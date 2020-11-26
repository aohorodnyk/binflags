package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagArrayUint64(t *testing.T) {
	for idx, prov := range providerHasFlagArrayUint64() {
		t.Run(fmt.Sprintf("TestHasFlagArrayUint64_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagArrayUint64(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagArrayUint64(t *testing.T) {
	for idx, prov := range providerSetFlagArrayUint64() {
		t.Run(fmt.Sprintf("TestHasFlagArrayUint64_%d", idx), func(t *testing.T) {
			err := gobitflags.SetFlagArrayUint64(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagArrayUint64 struct {
	flags    []uint64
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayUint64 struct {
	flags    []uint64
	flag     uint64
	set      bool
	expected []uint64
	err      error
}

func providerHasFlagArrayUint64() []providerTypeHasFlagArrayUint64 {
	return []providerTypeHasFlagArrayUint64{
		{
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			flags:    []uint64{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			flags:    []uint64{115},
			flag:     5,
			expected: true,
		},
		{
			flags:    []uint64{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			flags:    []uint64{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			flags:    []uint64{9223372036854775808, 1, 0, 0},
			flag:     63,
			expected: true,
		},
		{
			flags:    []uint64{0, 1, 0, 0},
			flag:     64,
			expected: true,
		},
		{
			flags:    []uint64{0, 0, 1, 0},
			flag:     128,
			expected: true,
		},
		{
			flags:    []uint64{0, 0, 9, 0},
			flag:     131,
			expected: true,
		},
		{
			flags:    []uint64{0, 0, 217, 0},
			flag:     135,
			expected: true,
		},
		{
			flags:    []uint64{0, 0, 217, 0},
			flag:     134,
			expected: true,
		},
		{
			flags:    []uint64{0, 0, 217, 0},
			flag:     69,
			expected: false,
		},
		{
			flags:    []uint64{0, 0, 217, 0},
			flag:     72,
			expected: false,
		},
	}
}

func providerSetFlagArrayUint64() []providerTypeSetFlagArrayUint64 {
	return []providerTypeSetFlagArrayUint64{
		{
			flags:    []uint64{0},
			flag:     0,
			set:      true,
			expected: []uint64{1},
			err:      nil,
		},
		{
			flags:    []uint64{1},
			flag:     6,
			set:      true,
			expected: []uint64{65},
			err:      nil,
		},
		{
			flags:    []uint64{65},
			flag:     6,
			set:      false,
			expected: []uint64{1},
			err:      nil,
		},
		{
			flags:    []uint64{65},
			flag:     12,
			set:      true,
			expected: []uint64{4161},
			err:      nil,
		},
		{
			flags:    []uint64{65},
			flag:     63,
			set:      true,
			expected: []uint64{9223372036854775873},
			err:      nil,
		},
		{
			flags:    []uint64{65, 0},
			flag:     64,
			set:      true,
			expected: []uint64{65, 1},
			err:      nil,
		},
		{
			flags:    []uint64{65, 0},
			flag:     127,
			set:      true,
			expected: []uint64{65, 9223372036854775808},
			err:      nil,
		},
		{
			flags:    []uint64{65, 0, 1235, 724, 635},
			flag:     132,
			set:      true,
			expected: []uint64{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			flags:    []uint64{65, 0, 1235, 724, 635},
			flag:     132,
			set:      false,
			expected: []uint64{65, 0, 1219, 724, 635},
			err:      nil,
		},
		{
			flags:    []uint64{0},
			flag:     64,
			set:      true,
			expected: []uint64{0},
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
		{
			flags:    nil,
			flag:     0,
			set:      false,
			expected: nil,
			err:      errors.New(gobitflags.ErrorMsgFlagsArrayNil),
		},
	}
}
