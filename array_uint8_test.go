package gobitflags_test

import (
	"errors"
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagArrayUint8(t *testing.T) {
	for idx, prov := range providerHasFlagArrayUint8() {
		t.Run(fmt.Sprintf("TestHasFlagArrayUint8_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagArrayUint8(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestHasFlagArrayByte(t *testing.T) {
	for idx, prov := range providerHasFlagArrayUint8() {
		t.Run(fmt.Sprintf("TestHasFlagArrayByte_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagArrayByte(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagArrayUint8(t *testing.T) {
	for idx, prov := range providerSetFlagArrayUint8() {
		t.Run(fmt.Sprintf("TestHasFlagArrayUint8_%d", idx), func(t *testing.T) {
			err := gobitflags.SetFlagArrayUint8(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

func TestSetFlagArrayByte(t *testing.T) {
	for idx, prov := range providerSetFlagArrayUint8() {
		t.Run(fmt.Sprintf("TestHasFlagArrayByte_%d", idx), func(t *testing.T) {
			err := gobitflags.SetFlagArrayByte(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
			assert.Equal(t, prov.err, err)
		})
	}
}

type providerTypeHasFlagArrayUint8 struct {
	flags    []uint8
	flag     uint64
	expected bool
}

type providerTypeSetFlagArrayUint8 struct {
	flags    []uint8
	flag     uint64
	set      bool
	expected []uint8
	err      error
}

func providerHasFlagArrayUint8() []providerTypeHasFlagArrayUint8 {
	return []providerTypeHasFlagArrayUint8{
		{
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			flags:    []uint8{115, 12, 105},
			flag:     36,
			expected: false,
		},
		{
			flags:    []uint8{115},
			flag:     5,
			expected: true,
		},
		{
			flags:    []uint8{0, 0, 0, 0},
			flag:     0,
			expected: false,
		},
		{
			flags:    []uint8{0, 0, 0, 0},
			flag:     8,
			expected: false,
		},
		{
			flags:    []uint8{0, 0, 1, 0},
			flag:     16,
			expected: true,
		},
		{
			flags:    []uint8{0, 0, 9, 0},
			flag:     19,
			expected: true,
		},
		{
			flags:    []uint8{0, 0, 217, 0},
			flag:     23,
			expected: true,
		},
		{
			flags:    []uint8{0, 0, 217, 0},
			flag:     24,
			expected: false,
		},
	}
}

func providerSetFlagArrayUint8() []providerTypeSetFlagArrayUint8 {
	return []providerTypeSetFlagArrayUint8{
		{
			flags:    []uint8{0},
			flag:     0,
			set:      true,
			expected: []uint8{1},
			err:      nil,
		},
		{
			flags:    []uint8{1},
			flag:     6,
			set:      true,
			expected: []uint8{65},
			err:      nil,
		},
		{
			flags:    []uint8{65},
			flag:     6,
			set:      false,
			expected: []uint8{1},
			err:      nil,
		},
		{
			flags:    []uint8{65, 1, 0, 83},
			flag:     30,
			set:      true,
			expected: []uint8{65, 1, 0, 83},
			err:      nil,
		},
		{
			flags:    []uint8{65, 1, 0, 83},
			flag:     30,
			set:      false,
			expected: []uint8{65, 1, 0, 19},
			err:      nil,
		},
		{
			flags:    []uint8{0, 0},
			flag:     7,
			set:      true,
			expected: []uint8{128, 0},
			err:      nil,
		},
		{
			flags:    []uint8{0},
			flag:     8,
			set:      true,
			expected: []uint8{0},
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
