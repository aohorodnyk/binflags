package gobitflags

import "errors"

func HasFlagMapByte(flags map[uint64]uint8, flag uint64) bool {
	return HasFlagMapUint8(flags, flag)
}

func SetFlagMapByte(flags map[uint64]uint8, flag uint64, val bool) error {
	return SetFlagMapUint8(flags, flag, val)
}

func HasFlagMapUint8(flags map[uint64]uint8, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	flagBits := flags[idx]

	conv := uint8(1 << bit)
	return flagBits&conv == conv
}

func SetFlagMapUint8(flags map[uint64]uint8, flag uint64, set bool) error {
	if flags == nil {
		return errors.New(ErrorMsgFlagsMapNil)
	}

	if HasFlagMapUint8(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	flagBits, ok := flags[idx]

	conv := uint8(1 << bit)
	ret := flagBits ^ conv
	if ret != 0 {
		flags[idx] = ret
	} else if ok {
		delete(flags, idx)
	}

	return nil
}
