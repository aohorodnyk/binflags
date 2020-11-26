package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagArrayInt8(t *testing.T) {
	for idx, prov := range providerHasFlagArrayInt8() {
		t.Run(fmt.Sprintf("TestHasFlagArrayInt8_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagArrayInt8(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagArrayInt8(t *testing.T) {
	for idx, prov := range providerSetFlagArrayInt8() {
		t.Run(fmt.Sprintf("TestHasFlagArrayInt8_%d", idx), func(t *testing.T) {
			err := gobitflags.SetFlagArrayInt8(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagArrayInt8 struct {
	flags    []int8
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayInt8 struct {
	flags    []int8
	flag     uint64
	set      bool
	expected []int8
	err      error
}

func providerHasFlagArrayInt8() []providerTypeHasFlagArrayInt8 {
	return []providerTypeHasFlagArrayInt8{
		{
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			flags:    []int8{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			flags:    []int8{115},
			flag:     5,
			expected: true,
		},
		{
			flags:    []int8{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			flags:    []int8{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			flags:    []int8{0, 0, 1, 0},
			flag:     16,
			expected: true,
		},
		{
			flags:    []int8{0, 0, 9, 0},
			flag:     19,
			expected: true,
		},
		{
			flags:    []int8{0, 0, -39, 0},
			flag:     23,
			expected: true,
		},
		{
			flags:    []int8{0, 0, -39, 0},
			flag:     24,
			expected: false,
		},
	}
}

func providerSetFlagArrayInt8() []providerTypeSetFlagArrayInt8 {
	return []providerTypeSetFlagArrayInt8{
		{
			flags:    []int8{0},
			flag:     0,
			set:      true,
			expected: []int8{1},
			err:      nil,
		},
		{
			flags:    []int8{1},
			flag:     6,
			set:      true,
			expected: []int8{65},
			err:      nil,
		},
		{
			flags:    []int8{65, 1, 0, 83},
			flag:     30,
			set:      true,
			expected: []int8{65, 1, 0, 83},
			err:      nil,
		},
		{
			flags:    []int8{65, 1, 0, 83},
			flag:     30,
			set:      false,
			expected: []int8{65, 1, 0, 19},
			err:      nil,
		},
		{
			flags:    []int8{0, 0},
			flag:     7,
			set:      true,
			expected: []int8{-128, 0},
			err:      nil,
		},
		{
			flags:    []int8{0},
			flag:     8,
			set:      true,
			expected: []int8{0},
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
