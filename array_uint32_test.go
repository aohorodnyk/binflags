package binflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/binflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagArrayUint32(t *testing.T) {
	for idx, prov := range providerHasFlagArrayUint32() {
		t.Run(fmt.Sprintf("TestHasFlagArrayUint32_%d", idx), func(t *testing.T) {
			actual := binflags.HasFlagArrayUint32(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagArrayUint32(t *testing.T) {
	for idx, prov := range providerSetFlagArrayUint32() {
		t.Run(fmt.Sprintf("TestHasFlagArrayUint32_%d", idx), func(t *testing.T) {
			err := binflags.SetFlagArrayUint32(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagArrayUint32 struct {
	flags    []uint32
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayUint32 struct {
	flags    []uint32
	flag     uint64
	set      bool
	expected []uint32
	err      error
}

func providerHasFlagArrayUint32() []providerTypeHasFlagArrayUint32 {
	return []providerTypeHasFlagArrayUint32{
		{
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			flags:    []uint32{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			flags:    []uint32{115},
			flag:     5,
			expected: true,
		},
		{
			flags:    []uint32{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			flags:    []uint32{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			flags:    []uint32{2147483713, 1, 0, 0},
			flag:     31,
			expected: true,
		},
		{
			flags:    []uint32{0, 1, 0, 0},
			flag:     32,
			expected: true,
		},
		{
			flags:    []uint32{0, 0, 1, 0},
			flag:     64,
			expected: true,
		},
		{
			flags:    []uint32{0, 0, 9, 0},
			flag:     67,
			expected: true,
		},
		{
			flags:    []uint32{0, 0, 217, 0},
			flag:     71,
			expected: true,
		},
		{
			flags:    []uint32{0, 0, 217, 0},
			flag:     70,
			expected: true,
		},
		{
			flags:    []uint32{0, 0, 217, 0},
			flag:     69,
			expected: false,
		},
		{
			flags:    []uint32{0, 0, 217, 0},
			flag:     72,
			expected: false,
		},
	}
}

func providerSetFlagArrayUint32() []providerTypeSetFlagArrayUint32 {
	return []providerTypeSetFlagArrayUint32{
		{
			flags:    []uint32{0},
			flag:     0,
			set:      true,
			expected: []uint32{1},
			err:      nil,
		},
		{
			flags:    []uint32{1},
			flag:     6,
			set:      true,
			expected: []uint32{65},
			err:      nil,
		},
		{
			flags:    []uint32{65},
			flag:     6,
			set:      false,
			expected: []uint32{1},
			err:      nil,
		},
		{
			flags:    []uint32{65},
			flag:     12,
			set:      true,
			expected: []uint32{4161},
			err:      nil,
		},
		{
			flags:    []uint32{65},
			flag:     31,
			set:      true,
			expected: []uint32{2147483713},
			err:      nil,
		},
		{
			flags:    []uint32{65, 0},
			flag:     32,
			set:      true,
			expected: []uint32{65, 1},
			err:      nil,
		},
		{
			flags:    []uint32{65, 0},
			flag:     63,
			set:      true,
			expected: []uint32{65, 2147483648},
			err:      nil,
		},
		{
			flags:    []uint32{65, 0, 1235, 724, 635},
			flag:     68,
			set:      true,
			expected: []uint32{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			flags:    []uint32{65, 0, 1235, 724, 635},
			flag:     68,
			set:      false,
			expected: []uint32{65, 0, 1219, 724, 635},
			err:      nil,
		},
		{
			flags:    []uint32{0},
			flag:     32,
			set:      true,
			expected: []uint32{0},
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
