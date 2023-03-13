package binflags_test

import (
	"fmt"
	"testing"

	"github.com/aohorodnyk/binflags"
)

func TestSetBit_int8(t *testing.T) {
	t.Parallel()

	provider := []struct {
		value  uint8
		bitNum uint8
		exp    uint8
	}{
		{0, 0, 1},
		{1, 0, 1},
		{8, 7, 136},
		{127, 7, 255},
		{255, 7, 255},
		{8, 8, 8},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestSetBit_int8_%d", idx), func(t *testing.T) {
			t.Parallel()

			act := binflags.SetBit(prov.value, prov.bitNum)
			if act != prov.exp {
				t.Errorf("SetBit returned unexpected result, expected %d, got %d", prov.exp, act)
			}
		})
	}
}

func TestSetBit_uint8(t *testing.T) {
	t.Parallel()

	provider := []struct {
		value  int8
		bitNum uint8
		exp    int8
	}{
		{0, 0, 1},
		{1, 0, 1},
		{-1, 7, -1},
		{-128, 7, -128},
		{8, 7, -120},
		{8, 8, 8},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestSetBit_int8_%d", idx), func(t *testing.T) {
			t.Parallel()

			act := binflags.SetBit(prov.value, prov.bitNum)
			if act != prov.exp {
				t.Errorf("SetBit returned unexpected result, expected %d, got %d", prov.exp, act)
			}
		})
	}
}

func TestUnsetBit(t *testing.T) {
	t.Parallel()

	provider := []struct {
		value  int8
		bitNum uint8
		exp    int8
	}{
		{0, 0, 0},
		{1, 0, 0},
		{-1, 7, 127},
		{-121, 4, -121},
		{-125, 7, 3},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestUnsetBit_%d", idx), func(t *testing.T) {
			t.Parallel()

			act := binflags.UnsetBit(prov.value, prov.bitNum)
			if act != prov.exp {
				t.Errorf("Unexpected resilt, expected %d, got %d", prov.exp, act)
			}
		})
	}
}

func TestIsSetBit_int8(t *testing.T) {
	t.Parallel()

	provider := []struct {
		value  int8
		bitNum uint8
		exp    bool
	}{
		{0, 0, false},
		{1, 0, true},
		{-128, 7, true},
		{0b0100000, 5, true},
		{0b0100100, 2, true},
		{-1, 6, true},
		{-1, 7, true},
		{-1, 8, false},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestIsSetBit_int8_%d", idx), func(t *testing.T) {
			t.Parallel()

			act := binflags.IsSetBit(prov.value, prov.bitNum)
			if act != prov.exp {
				t.Errorf("Unexpected resilt, expected %t, got %t", prov.exp, act)
			}
		})
	}
}

func TestIsSetBit_uint8(t *testing.T) {
	t.Parallel()

	provider := []struct {
		value  uint8
		bitNum uint8
		exp    bool
	}{
		{0, 0, false},
		{1, 0, true},
		{0b0100000, 5, true},
		{0b0100100, 2, true},
		{255, 6, true},
		{255, 7, true},
		{255, 8, false},
	}

	for idx, prov := range provider {
		prov := prov

		t.Run(fmt.Sprintf("TestIsSetBit_uint8_%d", idx), func(t *testing.T) {
			t.Parallel()

			act := binflags.IsSetBit(prov.value, prov.bitNum)
			if act != prov.exp {
				t.Errorf("Unexpected result, expected %t, got %t", prov.exp, act)
			}
		})
	}
}
