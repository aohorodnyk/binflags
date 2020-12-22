package binflags

import "errors"

func HasFlagArrayInt64(flags []int64, flag uint64) bool {
	if len(flags) == 0 {
		return false
	}

	idx, bit := flagExt(flag, FlagMaxInt64)
	if len(flags) <= int(idx) {
		return false
	}

	return HasFlagInt64(flags[idx], bit)
}

func SetFlagArrayInt64(flags []int64, flag uint64, set bool) error {
	if flags == nil {
		return errors.New(ErrorMsgFlagsArrayNil)
	}

	if HasFlagArrayInt64(flags, flag) == set {
		return nil
	}

	idx, bit := flagExt(flag, FlagMaxInt64)
	if len(flags) <= int(idx) {
		return errors.New(ErrorMsgOutOfRange)
	}

	var err error
	flags[idx], err = SetFlagInt64(flags[idx], bit, set)
	return err
}
