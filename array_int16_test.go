package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagArrayInt16(t *testing.T) {
	for idx, prov := range providerHasFlagArrayInt16() {
		t.Run(fmt.Sprintf("TestHasFlagArrayInt16_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagArrayInt16(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagArrayInt16(t *testing.T) {
	for idx, prov := range providerSetFlagArrayInt16() {
		t.Run(fmt.Sprintf("TestHasFlagArrayInt16_%d", idx), func(t *testing.T) {
			err := gobitflags.SetFlagArrayInt16(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagArrayInt16 struct {
	flags    []int16
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayInt16 struct {
	flags    []int16
	flag     uint64
	set      bool
	expected []int16
	err      error
}

func providerHasFlagArrayInt16() []providerTypeHasFlagArrayInt16 {
	return []providerTypeHasFlagArrayInt16{
		{
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			flags:    []int16{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			flags:    []int16{115},
			flag:     5,
			expected: true,
		},
		{
			flags:    []int16{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			flags:    []int16{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			flags:    []int16{0, 0, 1, 0},
			flag:     32,
			expected: true,
		},
		{
			flags:    []int16{0, 0, 9, 0},
			flag:     35,
			expected: true,
		},
		{
			flags:    []int16{0, 0, 217, 0},
			flag:     39,
			expected: true,
		},
		{
			flags:    []int16{0, 0, 217, 0},
			flag:     38,
			expected: true,
		},
		{
			flags:    []int16{0, 0, 217, 0},
			flag:     37,
			expected: false,
		},
		{
			flags:    []int16{0, 0, 217, 0},
			flag:     40,
			expected: false,
		},
	}
}

func providerSetFlagArrayInt16() []providerTypeSetFlagArrayInt16 {
	return []providerTypeSetFlagArrayInt16{
		{
			flags:    []int16{0},
			flag:     0,
			set:      true,
			expected: []int16{1},
			err:      nil,
		},
		{
			flags:    []int16{1},
			flag:     6,
			set:      true,
			expected: []int16{65},
			err:      nil,
		},
		{
			flags:    []int16{65},
			flag:     6,
			set:      false,
			expected: []int16{1},
			err:      nil,
		},
		{
			flags:    []int16{65},
			flag:     12,
			set:      true,
			expected: []int16{4161},
			err:      nil,
		},
		{
			flags:    []int16{65},
			flag:     15,
			set:      true,
			expected: []int16{-32703},
			err:      nil,
		},
		{
			flags:    []int16{65, 0},
			flag:     16,
			set:      true,
			expected: []int16{65, 1},
			err:      nil,
		},
		{
			flags:    []int16{65, 0},
			flag:     31,
			set:      true,
			expected: []int16{65, -32768},
			err:      nil,
		},
		{
			flags:    []int16{65, 0, 1235, 724, 635},
			flag:     50,
			set:      true,
			expected: []int16{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			flags:    []int16{65, 0, 1235, 724, 635},
			flag:     50,
			set:      false,
			expected: []int16{65, 0, 1235, 720, 635},
			err:      nil,
		},
		{
			flags:    []int16{0},
			flag:     16,
			set:      true,
			expected: []int16{0},
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
		{
			flags:    []int16{0, 0},
			flag:     32,
			set:      true,
			expected: []int16{0, 0},
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
