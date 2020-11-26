package gobitflags

import "errors"

func HasFlagArrayUint8(flags []uint8, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	if len(flags) <= int(idx) {
		return false
	}

	flagBits := flags[idx]

	conv := uint8(1 << bit)
	return flagBits&conv == conv
}

func SetFlagArrayUint8(flags []uint8, flag uint64, set bool) error {
	if flags == nil {
		return errors.New(ErrorMsgFlagsArrayNil)
	}

	if HasFlagArrayUint8(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	if len(flags) <= int(idx) {
		return errors.New(ErrorMsgOutOfRange)
	}

	flagBits := flags[idx]

	conv := uint8(1 << bit)
	flags[idx] = flagBits ^ conv

	return nil
}
