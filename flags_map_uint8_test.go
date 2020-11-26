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

func TestSetFlagMapUint8(t *testing.T) {
	for idx, prov := range providerSetFlagMapUint8() {
		t.Run(fmt.Sprintf("TestHasFlagMapUint8_%d", idx), func(t *testing.T) {
			gobitflags.SetFlagMapUint8(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
		})
	}
}

func TestSetFlagMapByte(t *testing.T) {
	for idx, prov := range providerSetFlagMapUint8() {
		t.Run(fmt.Sprintf("TestHasFlagMapUint8_%d", idx), func(t *testing.T) {
			gobitflags.SetFlagMapByte(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
		})
	}
}

type providerTypeHasFlagMapUint8 struct {
	flags    map[uint64]uint8
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapUint8 struct {
	flags    map[uint64]uint8
	flag     uint64
	set      bool
	expected map[uint64]uint8
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

func providerSetFlagMapUint8() []providerTypeSetFlagMapUint8 {
	return []providerTypeSetFlagMapUint8{
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
			flags: map[uint64]uint8{},
			flag:  0,
			set:   true,
			expected: map[uint64]uint8{
				0: 1,
			},
		},
		{
			flags: map[uint64]uint8{},
			flag:  7,
			set:   true,
			expected: map[uint64]uint8{
				0: 128,
			},
		},
		{
			flags: map[uint64]uint8{},
			flag:  2346,
			set:   true,
			expected: map[uint64]uint8{
				293: 4,
			},
		},
		{
			flags: map[uint64]uint8{
				0: 92,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]uint8{
				0: 220,
			},
		},
		{
			flags: map[uint64]uint8{
				0: 220,
			},
			flag: 7,
			set:  false,
			expected: map[uint64]uint8{
				0: 92,
			},
		},
		{
			flags: map[uint64]uint8{
				0:    220,
				1:    56,
				6:    32,
				6235: 123,
			},
			flag: 48,
			set:  true,
			expected: map[uint64]uint8{
				0:    220,
				1:    56,
				6:    33,
				6235: 123,
			},
		},
		{
			flags: map[uint64]uint8{
				0:    220,
				1:    56,
				6:    32,
				6235: 123,
			},
			flag: 53,
			set:  false,
			expected: map[uint64]uint8{
				0:    220,
				1:    56,
				6235: 123,
			},
		},
		{
			flags: map[uint64]uint8{
				0:    220,
				1:    56,
				6235: 123,
			},
			flag: 53,
			set:  true,
			expected: map[uint64]uint8{
				0:    220,
				1:    56,
				6:    32,
				6235: 123,
			},
		},
		{
			flags: map[uint64]uint8{
				0:    220,
				1:    56,
				6235: 123,
			},
			flag: 53,
			set:  false,
			expected: map[uint64]uint8{
				0:    220,
				1:    56,
				6235: 123,
			},
		},
	}
}
