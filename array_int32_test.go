package binflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/binflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagArrayInt32(t *testing.T) {
	for idx, prov := range providerHasFlagArrayInt32() {
		t.Run(fmt.Sprintf("TestHasFlagArrayInt32_%d", idx), func(t *testing.T) {
			actual := binflags.HasFlagArrayInt32(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagArrayInt32(t *testing.T) {
	for idx, prov := range providerSetFlagArrayInt32() {
		t.Run(fmt.Sprintf("TestHasFlagArrayInt32_%d", idx), func(t *testing.T) {
			err := binflags.SetFlagArrayInt32(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagArrayInt32 struct {
	flags    []int32
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayInt32 struct {
	flags    []int32
	flag     uint64
	set      bool
	expected []int32
	err      error
}

func providerHasFlagArrayInt32() []providerTypeHasFlagArrayInt32 {
	return []providerTypeHasFlagArrayInt32{
		{
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			flags:    []int32{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			flags:    []int32{115},
			flag:     5,
			expected: true,
		},
		{
			flags:    []int32{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			flags:    []int32{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			flags:    []int32{-2147483583, 1, 0, 0},
			flag:     31,
			expected: true,
		},
		{
			flags:    []int32{0, 1, 0, 0},
			flag:     32,
			expected: true,
		},
		{
			flags:    []int32{0, 0, 1, 0},
			flag:     64,
			expected: true,
		},
		{
			flags:    []int32{0, 0, 9, 0},
			flag:     67,
			expected: true,
		},
		{
			flags:    []int32{0, 0, 217, 0},
			flag:     71,
			expected: true,
		},
		{
			flags:    []int32{0, 0, 217, 0},
			flag:     70,
			expected: true,
		},
		{
			flags:    []int32{0, 0, 217, 0},
			flag:     69,
			expected: false,
		},
		{
			flags:    []int32{0, 0, 217, 0},
			flag:     72,
			expected: false,
		},
	}
}

func providerSetFlagArrayInt32() []providerTypeSetFlagArrayInt32 {
	return []providerTypeSetFlagArrayInt32{
		{
			flags:    []int32{0},
			flag:     0,
			set:      true,
			expected: []int32{1},
			err:      nil,
		},
		{
			flags:    []int32{1},
			flag:     6,
			set:      true,
			expected: []int32{65},
			err:      nil,
		},
		{
			flags:    []int32{65},
			flag:     6,
			set:      false,
			expected: []int32{1},
			err:      nil,
		},
		{
			flags:    []int32{65},
			flag:     12,
			set:      true,
			expected: []int32{4161},
			err:      nil,
		},
		{
			flags:    []int32{65},
			flag:     31,
			set:      true,
			expected: []int32{-2147483583},
			err:      nil,
		},
		{
			flags:    []int32{65, 0},
			flag:     32,
			set:      true,
			expected: []int32{65, 1},
			err:      nil,
		},
		{
			flags:    []int32{65, 0},
			flag:     63,
			set:      true,
			expected: []int32{65, -2147483648},
			err:      nil,
		},
		{
			flags:    []int32{65, 0, 1235, 724, 635},
			flag:     68,
			set:      true,
			expected: []int32{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			flags:    []int32{65, 0, 1235, 724, 635},
			flag:     68,
			set:      false,
			expected: []int32{65, 0, 1219, 724, 635},
			err:      nil,
		},
		{
			flags:    []int32{0},
			flag:     32,
			set:      true,
			expected: []int32{0},
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
