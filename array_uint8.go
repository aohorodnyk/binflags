package gobitflags

import "errors"

func HasFlagArrayByte(flags []byte, flag uint64) bool {
	return HasFlagArrayUint8(flags, flag)
}

func SetFlagArrayByte(flags []uint8, flag uint64, set bool) error {
	return SetFlagArrayUint8(flags, flag, set)
}

func HasFlagArrayUint8(flags []uint8, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	if len(flags) <= int(idx) {
		return false
	}

	return HasFlagUint8(flags[idx], bit)
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

	var err error
	flags[idx], err = SetFlagUint8(flags[idx], bit, set)
	return err
}
