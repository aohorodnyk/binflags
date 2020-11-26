package gobitflags

import "errors"

func HasFlagMapUint16(flags map[uint64]uint16, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	flagBits := flags[idx]

	conv := uint16(1 << bit)
	return flagBits&conv == conv
}

func SetFlagMapUint16(flags map[uint64]uint16, flag uint64, set bool) error {
	if flags == nil {
		return errors.New(ErrorMsgFlagsMapNil)
	}

	if HasFlagMapUint16(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt16)
	flagBits, ok := flags[idx]

	conv := uint16(1 << bit)
	ret := flagBits ^ conv
	if ret != 0 {
		flags[idx] = ret
	} else if ok {
		delete(flags, idx)
	}

	return nil
}
