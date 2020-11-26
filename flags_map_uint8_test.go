package gobitflags_test

import (
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagMapUint8(t *testing.T) {
	for idx, prov := range providerHasFlagMapUint8() {
		t.Run(fmt.Sprintf("TestHasFlagMapUint8_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagMapUint8(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestHasFlagMapByte(t *testing.T) {
	for idx, prov := range providerHasFlagMapUint8() {
		t.Run(fmt.Sprintf("TestHasFlagMapUint8_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagMapByte(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

type providerTypeHasFlagMapUint8 struct {
	flags    map[uint64]uint8
	flag     uint64
	expected bool
}

func providerHasFlagMapUint8() []providerTypeHasFlagMapUint8 {
	return []providerTypeHasFlagMapUint8{
		{
			flags:    map[uint64]uint8{},
			flag:     0,
			expected: false,
		},
		{
			flags:    map[uint64]uint8{},
			flag:     2356241245436,
			expected: false,
		},
		{
			flags: map[uint64]uint8{
				0: 1,
			},
			flag:     0,
			expected: true,
		},
		{
			flags: map[uint64]uint8{
				234125: 128,
			},
			flag:     1873007,
			expected: true,
		},
		{
			flags: map[uint64]uint8{
				234125: 128,
			},
			flag:     1873006,
			expected: false,
		},
		{
			flags: map[uint64]uint8{
				234125: 128,
			},
			flag:     1873008,
			expected: false,
		},
		{
			flags: map[uint64]uint8{
				0:      128,
				234125: 128,
			},
			flag:     7,
			expected: true,
		},
		{
			flags: map[uint64]uint8{
				0:      255,
				5:      255,
				63:     255,
				234125: 128,
			},
			flag:     7,
			expected: true,
		},
		{
			flags: map[uint64]uint8{
				0:      255,
				5:      64,
				63:     128,
				234125: 128,
			},
			flag:     7,
			expected: true,
		},
		{
			flags: map[uint64]uint8{
				0:      255,
				5:      64,
				63:     128,
				234125: 128,
			},
			flag:     46,
			expected: true,
		},
		{
			flags: map[uint64]uint8{
				0:      255,
				5:      64,
				63:     128,
				234125: 128,
			},
			flag:     511,
			expected: true,
		},
	}
}
