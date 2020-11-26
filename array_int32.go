package gobitflags

import "errors"

func HasFlagArrayInt32(flags []int32, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt32)
	if len(flags) <= int(idx) {
		return false
	}

	return HasFlagInt32(flags[idx], bit)
}

func SetFlagArrayInt32(flags []int32, flag uint64, set bool) error {
	if flags == nil {
		return errors.New(ErrorMsgFlagsArrayNil)
	}

	if HasFlagArrayInt32(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt32)
	if len(flags) <= int(idx) {
		return errors.New(ErrorMsgOutOfRange)
	}

	var err error
	flags[idx], err = SetFlagInt32(flags[idx], bit, set)
	return err
}
