package gobitflags_test

import (
	"fmt"
	"github.com/aohorodnyk/gobitflags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlagMap(t *testing.T) {
	for idx, prov := range providerFlagMap() {
		t.Run(fmt.Sprintf("TestFlagMap_%d", idx), func(t *testing.T) {
			actKey, actVal := gobitflags.FlagMap(prov.flag, prov.flagBytes)

			assert.Equal(t, prov.expKey, actKey)
			assert.Equal(t, prov.expVal, actVal)
		})
	}
}

type providerTypeFlagMap struct {
	flag      uint64
	flagBytes uint8
	expKey    uint64
	expVal    uint8
}

func providerFlagMap() []providerTypeFlagMap {
	return []providerTypeFlagMap{
		{
			flag:      0,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    0,
			expVal:    0,
		},
		{
			flag:      5,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    0,
			expVal:    5,
		},
		{
			flag:      7,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    0,
			expVal:    7,
		},
		{
			flag:      8,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    1,
			expVal:    0,
		},
		{
			flag:      15,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    1,
			expVal:    7,
		},
		{
			flag:      16,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    2,
			expVal:    0,
		},
		{
			flag:      32,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    4,
			expVal:    0,
		},
		{
			flag:      39,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    4,
			expVal:    7,
		},
		{
			flag:      125345424254123235,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    15668178031765404,
			expVal:    3,
		},
		{
			flag:      125345424254123239,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    15668178031765404,
			expVal:    7,
		},
		{
			flag:      125345424254123270,
			flagBytes: gobitflags.FlagMaxInt8,
			expKey:    15668178031765408,
			expVal:    6,
		},
		{
			flag:      63,
			flagBytes: gobitflags.FlagMaxInt64,
			expKey:    0,
			expVal:    63,
		},
		{
			flag:      64,
			flagBytes: gobitflags.FlagMaxInt64,
			expKey:    1,
			expVal:    0,
		},
		{
			flag:      127,
			flagBytes: gobitflags.FlagMaxInt64,
			expKey:    1,
			expVal:    63,
		},
		{
			flag:      128,
			flagBytes: gobitflags.FlagMaxInt64,
			expKey:    2,
			expVal:    0,
		},
		{
			flag:      180,
			flagBytes: gobitflags.FlagMaxInt64,
			expKey:    2,
			expVal:    52,
		},
		{
			flag:      125345424254123235,
			flagBytes: gobitflags.FlagMaxInt64,
			expKey:    1958522253970675,
			expVal:    35,
		},
		{
			flag:      125345424254123263,
			flagBytes: gobitflags.FlagMaxInt64,
			expKey:    1958522253970675,
			expVal:    63,
		},
		{
			flag:      125345424254123293,
			flagBytes: gobitflags.FlagMaxInt64,
			expKey:    1958522253970676,
			expVal:    29,
		},
	}
}
