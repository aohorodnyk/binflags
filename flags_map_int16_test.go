package gobitflags_test

import (
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHasFlagMapUint16(t *testing.T) {
	for idx, prov := range providerHasFlagMapUint16() {
		t.Run(fmt.Sprintf("TestHasFlagMapUint16_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagMapInt16(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagMapUint16(t *testing.T) {
	for idx, prov := range providerSetFlagMapUint16() {
		t.Run(fmt.Sprintf("TestHasFlagMapUint16_%d", idx), func(t *testing.T) {
			gobitflags.SetFlagMapInt16(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
		})
	}
}

type providerTypeHasFlagMapUint16 struct {
	flags    map[uint64]int16
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapUint16 struct {
	flags    map[uint64]int16
	flag     uint64
	set      bool
	expected map[uint64]int16
}

func providerHasFlagMapUint16() []providerTypeHasFlagMapUint16 {
	return []providerTypeHasFlagMapUint16{
		{
			flags:    map[uint64]int16{},
			flag:     0,
			expected: false,
		},
		{
			flags:    map[uint64]int16{},
			flag:     12425,
			expected: false,
		},
		{
			flags:    nil,
			flag:     0,
			expected: false,
		},
		{
			flags:    nil,
			flag:     12425,
			expected: false,
		},
		{
			flags:    map[uint64]int16{},
			flag:     2356241245436,
			expected: false,
		},
		{
			flags: map[uint64]int16{
				0: 1,
			},
			flag:     0,
			expected: true,
		},
		{
			flags: map[uint64]int16{
				0: math.MinInt16,
			},
			flag:     15,
			expected: true,
		},
		{
			flags: map[uint64]int16{
				1: math.MinInt16,
			},
			flag:     31,
			expected: true,
		},
		{
			flags: map[uint64]int16{
				6235: 32,
			},
			flag:     99765,
			expected: true,
		},
		{
			flags: map[uint64]int16{
				6235: 32,
			},
			flag:     99764,
			expected: false,
		},
		{
			flags: map[uint64]int16{
				6235: 32,
			},
			flag:     99766,
			expected: false,
		},
	}
}

func providerSetFlagMapUint16() []providerTypeSetFlagMapUint16 {
	return []providerTypeSetFlagMapUint16{
		{
			flags:    nil,
			flag:     0,
			set:      false,
			expected: nil,
		},
		{
			flags:    nil,
			flag:     0,
			set:      true,
			expected: nil,
		},
		{
			flags: map[uint64]int16{},
			flag:  0,
			set:   true,
			expected: map[uint64]int16{
				0: 1,
			},
		},
		{
			flags: map[uint64]int16{},
			flag:  7,
			set:   true,
			expected: map[uint64]int16{
				0: 128,
			},
		},
		{
			flags: map[uint64]int16{},
			flag:  15,
			set:   true,
			expected: map[uint64]int16{
				0: -32768,
			},
		},
		{
			flags: map[uint64]int16{
				0: 92,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]int16{
				0: 220,
			},
		},
		{
			flags: map[uint64]int16{
				0: 220,
			},
			flag: 7,
			set:  false,
			expected: map[uint64]int16{
				0: 92,
			},
		},
		{
			flags: map[uint64]int16{
				0: 220,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]int16{
				0: 220,
			},
		},
		{
			flags: map[uint64]int16{
				0: 220,
			},
			flag: 1,
			set:  false,
			expected: map[uint64]int16{
				0: 220,
			},
		},
		{
			flags: map[uint64]int16{
				0: 220,
			},
			flag: 14,
			set:  true,
			expected: map[uint64]int16{
				0: 16604,
			},
		},
		{
			flags: map[uint64]int16{
				0: 16604,
				6: 23562,
			},
			flag: 95,
			set:  true,
			expected: map[uint64]int16{
				0: 16604,
				5: -32768,
				6: 23562,
			},
		},
		{
			flags: map[uint64]int16{
				0: 16604,
				5: -32768,
				6: 23562,
			},
			flag: 95,
			set:  false,
			expected: map[uint64]int16{
				0: 16604,
				6: 23562,
			},
		},
		{
			flags: map[uint64]int16{
				0: 16604,
				5: -32768,
				6: 23562,
			},
			flag: 97,
			set:  false,
			expected: map[uint64]int16{
				0: 16604,
				5: -32768,
				6: 23560,
			},
		},
		{
			flags: map[uint64]int16{
				0: 16604,
				5: -32768,
				6: 23562,
			},
			flag: 25365262353,
			set:  true,
			expected: map[uint64]int16{
				0:          16604,
				5:          -32768,
				6:          23562,
				1585328897: 2,
			},
		},
	}
}
