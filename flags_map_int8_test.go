package gobitflags_test

import (
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasFlagMapInt8(t *testing.T) {
	for idx, prov := range providerHasFlagMapInt8() {
		t.Run(fmt.Sprintf("TestHasFlagMapInt8_%d", idx), func(t *testing.T) {
			actual := gobitflags.HasFlagMapInt8(prov.flags, prov.flag)

			assert.Equal(t, prov.expected, actual)
		})
	}
}

func TestSetFlagMapInt8(t *testing.T) {
	for idx, prov := range providerSetFlagMapInt8() {
		t.Run(fmt.Sprintf("TestHasFlagMapInt8_%d", idx), func(t *testing.T) {
			gobitflags.SetFlagMapInt8(prov.flags, prov.flag, prov.set)

			assert.Equal(t, prov.expected, prov.flags)
		})
	}
}

type providerTypeHasFlagMapInt8 struct {
	flags    map[uint64]int8
	flag     uint64
	expected bool
}

type providerTypeSetFlagMapInt8 struct {
	flags    map[uint64]int8
	flag     uint64
	set      bool
	expected map[uint64]int8
}

func providerHasFlagMapInt8() []providerTypeHasFlagMapInt8 {
	return []providerTypeHasFlagMapInt8{
		{
			flags:    map[uint64]int8{},
			flag:     0,
			expected: false,
		},
		{
			flags:    map[uint64]int8{},
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
			flags:    map[uint64]int8{},
			flag:     2356241245436,
			expected: false,
		},
		{
			flags: map[uint64]int8{
				0: 1,
			},
			flag:     0,
			expected: true,
		},
		{
			flags: map[uint64]int8{
				234125: -128,
			},
			flag:     1873007,
			expected: true,
		},
		{
			flags: map[uint64]int8{
				234125: -128,
			},
			flag:     1873006,
			expected: false,
		},
		{
			flags: map[uint64]int8{
				234125: -128,
			},
			flag:     1873008,
			expected: false,
		},
		{
			flags: map[uint64]int8{
				0:      -128,
				234125: -128,
			},
			flag:     7,
			expected: true,
		},
		{
			flags: map[uint64]int8{
				0:      -1,
				5:      -1,
				63:     -1,
				234125: -128,
			},
			flag:     7,
			expected: true,
		},
		{
			flags: map[uint64]int8{
				0:      -1,
				5:      64,
				63:     -128,
				234125: -128,
			},
			flag:     7,
			expected: true,
		},
		{
			flags: map[uint64]int8{
				0:      -1,
				5:      64,
				63:     -128,
				234125: -128,
			},
			flag:     46,
			expected: true,
		},
		{
			flags: map[uint64]int8{
				0:      -1,
				5:      64,
				63:     -128,
				234125: -128,
			},
			flag:     511,
			expected: true,
		},
	}
}

func providerSetFlagMapInt8() []providerTypeSetFlagMapInt8 {
	return []providerTypeSetFlagMapInt8{
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
			flags: map[uint64]int8{},
			flag:  0,
			set:   true,
			expected: map[uint64]int8{
				0: 1,
			},
		},
		{
			flags: map[uint64]int8{
				0: -128,
			},
			flag:     7,
			set:      false,
			expected: map[uint64]int8{},
		},
		{
			flags: map[uint64]int8{
				0: -128,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]int8{
				0: -128,
			},
		},
		{
			flags: map[uint64]int8{},
			flag:  7,
			set:   true,
			expected: map[uint64]int8{
				0: -128,
			},
		},
		{
			flags: map[uint64]int8{},
			flag:  2346,
			set:   true,
			expected: map[uint64]int8{
				293: 4,
			},
		},
		{
			flags: map[uint64]int8{
				0: 92,
			},
			flag: 7,
			set:  true,
			expected: map[uint64]int8{
				0: -36,
			},
		},
		{
			flags: map[uint64]int8{
				0: -36,
			},
			flag: 7,
			set:  false,
			expected: map[uint64]int8{
				0: 92,
			},
		},
		{
			flags: map[uint64]int8{
				0:    -36,
				1:    56,
				6:    32,
				6235: 123,
			},
			flag: 48,
			set:  true,
			expected: map[uint64]int8{
				0:    -36,
				1:    56,
				6:    33,
				6235: 123,
			},
		},
		{
			flags: map[uint64]int8{
				0:    -36,
				1:    56,
				6:    32,
				6235: 123,
			},
			flag: 53,
			set:  false,
			expected: map[uint64]int8{
				0:    -36,
				1:    56,
				6235: 123,
			},
		},
		{
			flags: map[uint64]int8{
				0:    -36,
				1:    56,
				6235: 123,
			},
			flag: 53,
			set:  true,
			expected: map[uint64]int8{
				0:    -36,
				1:    56,
				6:    32,
				6235: 123,
			},
		},
		{
			flags: map[uint64]int8{
				0:    -36,
				1:    56,
				6235: 123,
			},
			flag: 53,
			set:  false,
			expected: map[uint64]int8{
				0:    -36,
				1:    56,
				6235: 123,
			},
		},
	}
}
