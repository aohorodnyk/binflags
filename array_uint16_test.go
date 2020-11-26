package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagArrayUint16(t *testing.T) {
	for idx, prov := range providerHasFlagArrayUint16() {
		t.Run(fmt.Sprintf("TestHasFlagArrayUint16_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagArrayUint16(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagArrayUint16(t *testing.T) {
	for idx, prov := range providerSetFlagArrayUint16() {
		t.Run(fmt.Sprintf("TestHasFlagArrayUint16_%d", idx), func(t *testing.T) {
			err := gobitflags.SetFlagArrayUint16(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagArrayUint16 struct {
	flags    []uint16
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayUint16 struct {
	flags    []uint16
	flag     uint64
	set      bool
	expected []uint16
	err      error
}

func providerHasFlagArrayUint16() []providerTypeHasFlagArrayUint16 {
	return []providerTypeHasFlagArrayUint16{
		{
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			flags:    []uint16{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			flags:    []uint16{115},
			flag:     5,
			expected: true,
		},
		{
			flags:    []uint16{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			flags:    []uint16{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			flags:    []uint16{0, 0, 1, 0},
			flag:     32,
			expected: true,
		},
		{
			flags:    []uint16{0, 0, 9, 0},
			flag:     35,
			expected: true,
		},
		{
			flags:    []uint16{0, 0, 217, 0},
			flag:     39,
			expected: true,
		},
		{
			flags:    []uint16{0, 0, 217, 0},
			flag:     38,
			expected: true,
		},
		{
			flags:    []uint16{0, 0, 217, 0},
			flag:     37,
			expected: false,
		},
		{
			flags:    []uint16{0, 0, 217, 0},
			flag:     40,
			expected: false,
		},
	}
}

func providerSetFlagArrayUint16() []providerTypeSetFlagArrayUint16 {
	return []providerTypeSetFlagArrayUint16{
		{
			flags:    []uint16{0},
			flag:     0,
			set:      true,
			expected: []uint16{1},
			err:      nil,
		},
		{
			flags:    []uint16{1},
			flag:     6,
			set:      true,
			expected: []uint16{65},
			err:      nil,
		},
		{
			flags:    []uint16{65},
			flag:     6,
			set:      false,
			expected: []uint16{1},
			err:      nil,
		},
		{
			flags:    []uint16{65},
			flag:     12,
			set:      true,
			expected: []uint16{4161},
			err:      nil,
		},
		{
			flags:    []uint16{65},
			flag:     15,
			set:      true,
			expected: []uint16{32833},
			err:      nil,
		},
		{
			flags:    []uint16{65, 0},
			flag:     16,
			set:      true,
			expected: []uint16{65, 1},
			err:      nil,
		},
		{
			flags:    []uint16{65, 0, 1235, 724, 635},
			flag:     50,
			set:      true,
			expected: []uint16{65, 0, 1235, 724, 635},
			err:      nil,
		},
		{
			flags:    []uint16{65, 0, 1235, 724, 635},
			flag:     50,
			set:      false,
			expected: []uint16{65, 0, 1235, 720, 635},
			err:      nil,
		},
		{
			flags:    []uint16{0},
			flag:     16,
			set:      true,
			expected: []uint16{0},
			err:      errors.New(gobitflags.ErrorMsgOutOfRange),
		},
		{
			flags:    []uint16{0, 0},
			flag:     32,
			set:      true,
			expected: []uint16{0, 0},
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
