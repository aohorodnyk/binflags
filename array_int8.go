package gobitflags

import "errors"

func HasFlagArrayInt8(flags []int8, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	if len(flags) <= int(idx) {
		return false
	}

	return HasFlagInt8(flags[idx], bit)
}

func SetFlagArrayInt8(flags []int8, flag uint64, set bool) error {
	if flags == nil {
		return errors.New(ErrorMsgFlagsArrayNil)
	}

	if HasFlagArrayInt8(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt8)
	if len(flags) <= int(idx) {
		return errors.New(ErrorMsgOutOfRange)
	}

	var err error
	flags[idx], err = SetFlagInt8(flags[idx], bit, set)
	return err
}
